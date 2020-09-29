use crate::how::AnyResult;
use crate::state::AppStateRaw as State;

use crate::sqlx::{Done, Executor};
use reqwest::Client;
use std::net::SocketAddr;
use std::{collections::HashMap as Map, time::Duration};
use tokio::task::spawn;
use tokio::time::delay_for as sleep;

pub fn poll_keepstats(state: &State) {
    let keep = &state.config.keep;
    let client = Client::builder()
        .connect_timeout(Duration::from_secs(keep.request_timeout))
        .timeout(Duration::from_secs(keep.request_timeout))
        .build()
        .expect("new reqwest::Client");

    for keep_info in &keep.urls {
        spawn(loop_poll_peers(
            state.clone(),
            client.clone(),
            keep_info.keep_core.clone(),
            keep_info.netid,
            "keep_core",
        ));
        spawn(loop_poll_peers(
            state.clone(),
            client.clone(),
            keep_info.keep_ecdsa.clone(),
            keep_info.netid,
            "keep_ecdsa",
        ));
    }
}

async fn loop_poll_peers(
    state: State,
    client: Client,
    urls: Vec<String>,
    netid: u64,
    kind: &'static str,
) {
    if urls.is_empty() {
        return;
    }

    let sleep_interval = Duration::from_secs(state.config.keep.poll_interval);
    loop {
        for url in &urls {
            poll_peers(&state, &client, url.as_str(), netid, kind)
                .await
                .map(|(n, rows)| {
                    info!(
                        "poll_peers({}, {}) from {} got {} peers, {} rows affected",
                        netid, kind, url, n, rows
                    )
                })
                .map_err(|e| error!("poll_peers({}, {}) from {} failed {}", netid, kind, url, e))
                .ok();
        }
        sleep(sleep_interval).await
    }
}

#[derive(Serialize, Deserialize, Debug)]
pub struct KeepClientInfo {
    datetime: String,
    network_id: String,
    network_addrs: Vec<String>,
    ethereum_address: String,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct KeepPeer {
    network_id: String,
    network_addr: String,
    #[serde(default)]
    network_ip: String,
    #[serde(default)]
    network_port: u16,
    ethereum_address: String,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct KeepPeers {
    client_info: KeepClientInfo,
    connected_peers: Vec<KeepPeer>,
}

async fn poll_peers(
    state: &State,
    client: &Client,
    url: &str,
    netid: u64,
    kind: &str,
) -> AnyResult<(u64, u64)> {
    let peersif = client.get(url).send().await?.json::<KeepPeers>().await?;
    let mut peers = Vec::with_capacity(peersif.connected_peers.len() + 1);

    let sa_check = |a: &str, who: &str| {
        a.parse::<SocketAddr>()
            .map_err(|e| {
                error!(
                    "{}.{}'s {}.network_addr {} isn't SocketAddr: {}",
                    kind, url, who, a, e
                )
            })
            .ok()
            .and_then(|sa| {
                if sa.ip().is_global() {
                    Some(sa)
                } else {
                    error!(
                        "{}.{}'s {}.network_addr: {} isn't global SocketAddr",
                        kind, url, who, a
                    );
                    None
                }
            })
    };

    if let Some(sa) = peersif
        .client_info
        .network_addrs
        .iter()
        .filter_map(|a| sa_check(a, "client_info"))
        .next()
    {
        peers.push(KeepPeer {
            network_addr: sa.to_string(),
            network_ip: sa.ip().to_string(),
            network_port: sa.port(),
            network_id: peersif.client_info.network_id.clone(),
            ethereum_address: peersif.client_info.ethereum_address.to_lowercase(),
        })
    }

    let mut wallet_ip_nodes_map = Map::with_capacity(peers.len());
    for mut p in peersif.connected_peers {
        if let Some(sa) = sa_check(&p.network_addr, &p.network_id) {
            p.ethereum_address = p.ethereum_address.to_lowercase();
            p.network_ip = sa.ip().to_string();
            p.network_port = sa.port();

            // Multiple nodes using the same wallet are multiple independent nodes?
            let wallet_ip_nodes = wallet_ip_nodes_map
                .entry(format!("{}.{}", p.ethereum_address, p.network_ip))
                .or_insert(vec![]);
            wallet_ip_nodes.push(p.network_port);

            if wallet_ip_nodes.len() > 1 {
                warn!(
                    "{}.{}'s ethereum_address {} running {} nodes on {} will skip it's[1..]: {:?}",
                    kind,
                    url,
                    p.ethereum_address,
                    wallet_ip_nodes.len(),
                    p.network_ip,
                    wallet_ip_nodes
                );
            } else {
                peers.push(p);
            }
        }
    }

    // INSERT into peers(netid, kind, network_id, ethereum_address, create_dt, update_dt) (),() on conflict (e2k) do update set update_dt=, network_id=
    let mut sql = "INSERT into peers(netid, kind, network_id, network_ip, network_port, ethereum_address, create_dt, update_dt) values ".to_string();
    for i in 0..peers.len() {
        let peer = &peers[i];

        let value = format!(
            "({}, '{}', '{}', '{}', {},'{}', '{}', '{}')",
            netid,
            kind,
            peer.network_id,
            peer.network_ip,
            peer.network_port,
            peer.ethereum_address,
            peersif.client_info.datetime,
            peersif.client_info.datetime
        );
        sql.push_str(&value);

        if i + 1 < peers.len() {
            sql.push(',')
        }
    }

    sql.push_str(&format!(
        " on conflict (ethereum_address, network_ip, netid, kind) do update set update_dt='{}', network_id=excluded.network_id, network_port=least(peers.network_port, excluded.network_port)",
        peersif.client_info.datetime
    ));

    let rows_affected = state.pg.execute(sql.as_str()).await?.rows_affected();

    Ok((peers.len() as _, rows_affected))
}
