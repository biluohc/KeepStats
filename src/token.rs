use rust_decimal::Decimal;
use web3::contract::{Contract, Options};
use web3::types::{H160, U256};
use web3::{transports::Http, Web3};

use std::time::Duration;
use tokio::task::spawn;
use tokio::time::delay_for as sleep;

use crate::how::AnyResult;
use crate::state::AppStateRaw as State;
use crate::util::*;

// erc20abi: keep-core/pkg/chain/gen/abi/KeepToken.go
const TOKEN_ABI: &str = include_str!("token.abi.keep.json");

pub const TOKEN_KEEP: &str = "KEEP";
pub const TOKEN_TBTC: &str = "TBTC";
pub const TOKEN_ETH: &str = "ETH";

/*
# mainnet
curl 'https://api.etherscan.io/api?module=stats&action=tokensupply&contractaddress=0x85eee30c52b0b379b046fb0f85f4f3dc3009afec'
curl 'https://api.etherscan.io/api?module=stats&action=tokensupply&contractaddress=0x8dAEBADE922dF735c38C80C7eBD708Af50815fAa'
# ropsten
curl 'https://api-ropsten.etherscan.io/api?module=stats&action=tokensupply&contractaddress=0x343d3dda00415289cdd4e8030f63a4a5a2548ff9'
curl 'https://api-ropsten.etherscan.io/api?module=stats&action=tokensupply&contractaddress=0x7c07c42973047223f80c4a69bb62d5195460eb5f'
*/
pub fn poll_tokenstats(state: &State) {
    let config = &state.config;
    for token in &config.token {
        let rpc = match token.netid {
            1 => &config.mainnet_rpc,
            3 => &config.ropsten_rpc,
            other => {
                error!("unsupported netid: {}, skip", other);
                continue;
            }
        };

        spawn(loop_poll(
            TOKEN_KEEP,
            token.netid,
            token.contract_address_keep.to_owned(),
            rpc.to_string(),
            state.clone(),
        ));
        spawn(loop_poll(
            TOKEN_TBTC,
            token.netid,
            token.contract_address_tbtc.to_owned(),
            rpc.to_string(),
            state.clone(),
        ));
    }
}

async fn loop_poll(token: &'static str, netid: u64, ca: String, rpc: String, state: State) {
    loop {
        match poll(token, netid, &ca, &rpc, &state).await {
            Ok(rows) => {
                info!(
                    "poll({}, {}, {}) totalSupply rows_affected: {}",
                    rpc, token, ca, rows
                );
                sleep(Duration::from_secs(60 * 60)).await
            }
            Err(e) => {
                error!("poll({}, {}, {}) totalSupply failed: {}", rpc, token, ca, e);
                sleep(Duration::from_secs(60 * 1)).await;
            }
        }
    }
}

async fn poll(token: &str, netid: u64, ca: &str, rpc: &str, state: &State) -> AnyResult<u64> {
    use crate::sqlx::Done;

    let web3 = Web3::new(Http::new(rpc)?);
    let supply = token_supply(&web3, ca).await?;
    info!("{}.{} totalSupply: {}", netid, token, supply);

    sqlx::query!(
        r#"
insert into tokenstats (date, netid, token, total_supply) values(date(current_timestamp), $1, $2, $3) on conflict (date, netid, token) do nothing
        ;"#,
        netid as i16,
        token,
        supply
    )
    .execute(&state.pg)
    .await
    .map(|d| d.rows_affected())
    .map_err(Into::into)
}

async fn token_supply(web3: &Web3<Http>, ca: &str) -> AnyResult<Decimal> {
    let ca: H160 = parse(ca)?;
    let contract = Contract::from_json(web3.eth(), ca, TOKEN_ABI.as_bytes())?;
    let result: U256 = contract
        .query("totalSupply", (), None, Options::default(), None)
        .await?;
    Ok(result.to_string().parse()?)
}
