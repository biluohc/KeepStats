use ethabi::RawLog;
use rust_decimal::Decimal;
use web3::contract::{Contract, Options};
use web3::types::{H160, H256, U256};
use web3::{
    transports::Http,
    types::{BlockId, FilterBuilder, Log as Web3Log},
    Web3,
};

use std::time::Duration;
use tokio::task::spawn;
use tokio::time::delay_for as sleep;

use crate::how::AnyResult;
use crate::models::{contract_event::ContractEvent, stats::DbStats, DateTime};
use crate::state::AppStateRaw as State;
use crate::util::*;

// stakingAbi: keep-core/pkg/chain/gen/abi/TokenStaking.go:31
const STAKING_ABI: &str = include_str!("token.abi.staking.keep.json");
const BONDING_ABI: &str = include_str!("token.abi.bonding.eth.json");

use_contract!(token_staking, "src/token.abi.staking.keep.json");
use_contract!(keep_bonding, "src/token.abi.bonding.eth.json");

/*  Thanks
mainnet:    https://gist.github.com/knarz/034654b56096f99aa857bdaebe6c8710
ropsten:    https://gist.github.com/knarz/b3b906bf8b3f7255b5dc90f96d1c0165
*/
pub async fn poll_stakingstats(state: &State) {
    let config = &state.config;

    let p = |s| crate::util::parse::<_, H160>(s).unwrap();
    for token in &config.token {
        let rpc = match token.netid {
            // 1 => &config.mainnet_rpc, // disable now
            3 => &config.ropsten_rpc,
            other => {
                error!("unsupported netid: {}, skip", other);
                continue;
            }
        };

        token_staking::Contract::start(
            token.netid,
            vec![p(&token.contract_address_token_staking)],
            rpc,
            state,
        );
        keep_bonding::Contract::start(
            token.netid,
            vec![p(&token.contract_address_keep_bonding)],
            rpc,
            state,
        );
    }

    info!(
        "{:?}",
        poll(
            3,
            "0x234d2182B29c6a64ce3ab6940037b5C8FdAB608e",
            "0x502a2d67c6efb0a4e1a41257f8386b3e65b44920",
            &config.ropsten_rpc,
            state
        )
        .await
    );
}

async fn poll(netid: u64, ca: &str, oa: &str, rpc: &str, _state: &State) -> AnyResult<u64> {
    let web3 = Web3::new(Http::new(rpc)?);
    let (amount, created_at, undelegated_at) = get_gelegation_info(&web3, ca, oa).await?;
    info!(
        "{}.KEEP.{}: amount: {}, created_at: {}, undelegated: {}",
        netid, oa, amount, created_at, undelegated_at
    );

    let supply =
        get_unbonded_value(&web3, "0x60535a59b4e71f908f3feb0116f450703fb35ed8", oa).await?;
    info!("{}.ETH unbondedValue: {}", netid, supply);

    Ok(0)
}

async fn get_gelegation_info(
    web3: &Web3<Http>,
    ca: &str,
    oa: &str,
) -> AnyResult<(Decimal, u64, u64)> {
    let ca: H160 = parse(ca)?;
    let oa: H160 = parse(oa)?;
    let contract = Contract::from_json(web3.eth(), ca, STAKING_ABI.as_bytes())?;
    let (amount, created_at, undelegated_at): (U256, U256, U256) = contract
        .query("getDelegationInfo", (oa,), None, Options::default(), None)
        .await?;
    Ok((
        amount.to_string().parse()?,
        created_at.as_u64(),
        undelegated_at.as_u64(),
    ))
}

async fn get_unbonded_value(web3: &Web3<Http>, ca: &str, oa: &str) -> AnyResult<Decimal> {
    let ca: H160 = parse(ca)?;
    let oa: H160 = parse(oa)?;
    let contract = Contract::from_json(web3.eth(), ca, BONDING_ABI.as_bytes())?;

    let balance: U256 = contract
        .query("unbondedValue", (oa,), None, Options::default(), None)
        .await?;
    Ok(balance.to_string().parse()?)
}

async fn get_height(web3: &Web3<Http>) -> AnyResult<u64> {
    Ok(web3.eth().block_number().await?.as_u64())
}

// [from, to]
async fn get_logs(
    addresses: Vec<H160>,
    from: u64,
    to: u64,
    topic0: Vec<H256>,
    web3: &Web3<Http>,
) -> AnyResult<Vec<Web3Log>> {
    let filter = FilterBuilder::default()
        .address(addresses)
        .from_block(from.into())
        .to_block(to.into())
        .topics(Some(topic0), None, None, None)
        .build();

    let logs = web3.eth().logs(filter).await?;
    // let str = serde_json::to_string(&logs).unwrap();
    // let mut f = std::fs::OpenOptions::new().create(true).write(true).open("log.8968548-.json").unwrap();
    // use std::io::Write;
    // f.write(str.as_bytes()).unwrap();

    Ok(logs)
}

const HSAFE: u64 = 10;
const HSTEP_MAX: u64 = 10_0000;
const HSTEP_MIN: u64 = 10;

/// (height, datetime, json)
pub type AnalyzeRes = (String, i32, DateTime, String);
pub type LogWithInfo = (u32, u32, u32, String, &'static str, u64, String);

#[async_trait]
pub trait ContractLogHandle: Copy + Send + 'static {
    fn name() -> &'static str;
    fn event_hashs() -> Vec<(&'static str, ethabi::Hash)>;
    fn parse_log(rawlog: RawLog) -> ethabi::Result<(&'static str, serde_json::Value)>;

    async fn analyze_addr(
        netid: u64,
        _contract_addrs: &Vec<H160>,
        addr: &str,
        rpc: &str,
        state: &State,
    ) -> AnyResult<()> {
        let web3 = Web3::new(Http::new(&rpc)?);
        let events = Self::get_events(netid, &chrono::Utc::now(), addr, state).await?;
        let res = Self::analyze_events(netid, addr, events, &web3, state).await?;
        println!("{}.{}'s {}.stats: {:?}", netid, Self::name(), addr, res);
        Ok(())
    }

    fn start(netid: u64, contract_addrs: Vec<H160>, rpc: &String, state: &State) {
        spawn(Self::try_loop_call(
            netid,
            contract_addrs.clone(),
            rpc.clone(),
            state.clone(),
            "events",
        ));
        spawn(Self::try_loop_call(
            netid,
            contract_addrs.clone(),
            rpc.clone(),
            state.clone(),
            "stats",
        ));
    }

    async fn try_loop_call(
        netid: u64,
        contract_addrs: Vec<H160>,
        rpc: String,
        state: State,
        name: &'static str,
    ) {
        info!(
            "try_loop_handle {} {}.{}'s contract_addrs: {:?}, rpc: {}",
            name,
            netid,
            Self::name(),
            contract_addrs,
            rpc
        );

        // stagger
        sleep(Duration::from_secs((rand::random::<u8>() % 10) as _)).await;
        // return info!(
        //     "analyze_addr: {:?}",
        //     Self::analyze_addr(netid, &contract_addrs, "0x7295dbd0449aca361640c4084c3fdb0f30b720d8", &rpc, &state).await
        // );

        loop {
            let res = if name == "stats" {
                Self::loop_analyze_addrs(netid, &contract_addrs, &rpc, &state).await
            } else {
                Self::loop_handle_events(netid, &contract_addrs, &rpc, &state).await
            };
            warn!(
                "try_loop_handle {} {}.{}'s contract_addrs: {:?}, finished: {:?}",
                name,
                netid,
                Self::name(),
                contract_addrs,
                res
            );
            sleep(Duration::from_secs(60 * 1)).await;
        }
    }

    async fn loop_analyze_addrs(
        netid: u64,
        _contract_addrs: &Vec<H160>,
        rpc: &str,
        state: &State,
    ) -> AnyResult<()> {
        let next = Self::get_time_db(netid, state).await?;

        if next.should_update_normal.unwrap_or_default() && next.stats_date_next.is_some() {
            let next_date = next.stats_date_next.unwrap();
            let current_date = next.stats_date.unwrap();

            let addrs = Self::select_addrs(netid, &current_date, &next_date, state).await?;
            info!(
                "loop_analyze_addrs {}.{}'s events, {} < time < {}, addrs: {}",
                netid,
                Self::name(),
                current_date,
                next_date,
                addrs.len()
            );

            let web3 = Web3::new(Http::new(&rpc)?);

            let mut stats = vec![];
            for addr in &addrs {
                let events = Self::get_events(netid, &next_date, &addr, state).await?;
                let res = Self::analyze_events(netid, &addr, events, &web3, state).await?;
                stats.push(res);
            }

            let stats_len = stats.len();
            if stats_len == 0 {
                stats.push((
                    "".to_string(),
                    0,
                    next_date - chrono::Duration::hours(1),
                    "{}".to_string(),
                ));
            }

            let rows = Self::save_analyze_stats(netid, stats, &web3, state).await?;
            warn!(
                "loop_analyze_addrs {}.{}'s events, stats: {}, rows: {}",
                netid,
                Self::name(),
                stats_len,
                rows
            );
        } else {
            sleep(Duration::from_secs(60 * 60)).await;
        }

        Ok(())
    }

    async fn loop_handle_events(
        netid: u64,
        contract_addrs: &Vec<H160>,
        rpc: &str,
        state: &State,
    ) -> AnyResult<()> {
        info!(
            "loop_handle_events {}.{}'s events, rpc: {}, contract_addrs: {:?}",
            netid,
            Self::name(),
            rpc,
            contract_addrs
        );

        let web3 = Web3::new(Http::new(&rpc)?);
        let topic0 = Self::event_hashs().iter().map(|x| x.1).collect::<Vec<_>>();
        let mut height_db = Self::get_height_db(netid, state).await?;

        loop {
            let height_chain = get_height(&web3).await?;

            let height_chain_safe = height_chain - HSAFE;
            let blocks = if height_db + HSTEP_MAX < height_chain_safe {
                HSTEP_MAX
            } else if height_db + HSTEP_MIN < height_chain_safe {
                height_chain_safe - height_db
            } else {
                sleep(Duration::from_secs(60 * 10)).await;
                continue;
            };

            let from = height_db;
            let to = height_db + blocks;
            let logs = get_logs(contract_addrs.clone(), from, to, topic0.clone(), &web3).await?;
            info!(
                "{}.contract-{}, get_logs [{}, {}]: {}",
                netid,
                Self::name(),
                from,
                to,
                logs.len()
            );

            if logs.len() == 0 {
                height_db = to;
            } else {
                Self::save_logs(netid, logs, &web3, &state).await?;
                height_db = to;
            }
        }
    }

    async fn get_height_db(netid: u64, state: &State) -> AnyResult<u64> {
        Ok(sqlx::query!(
            r#"select greatest(max(height), 0) as h from contract_events where netid=$1 and contract = $2 ;"#,
            netid as i16,
            Self::name()
        )
        .fetch_one(&state.pg)
        .await?
        .h
        .unwrap_or(0) as _)
    }

    async fn get_time_db(netid: u64, state: &State) -> AnyResult<DbStats> {
        let time = sqlx::query_as!(DbStats,
            r#"
            select
            event_time, stats_time, stats_date, stats_date_next, date(event_time) > stats_date as should_update, -- select date_part('hour', current_time)  > 1
            now() > (stats_date + INTERVAL '10minute') and stats_date_next < now() as should_update_normal
          from (
          select
            event_time, stats_time,  
            cast(date(stats_time) as timestamptz) + INTERVAL '1day'  as stats_date, 
            cast(date(stats_time) + INTERVAL '2day' as timestamptz) as stats_date_next
          from 
          (
            select event_time, coalesce(stats_time_current, event_time_init - INTERVAL '1day') as stats_time
          from
            (
              select
                max(time) as event_time,
                min(time) as event_time_init
              from
                contract_events
              where
                netid = $1
                and contract = $2
            ) a
            cross join (
              select
                max(time) as stats_time_current
              from
                operatorstats
              where
                netid = $1
                and contract = $2
            ) b
          ) c
          )d
  ;"#,
            netid as i16,
            Self::name()
        )
        .fetch_one(&state.pg)
        .await?;

        info!(
            "get {}.{} next date: {}",
            netid,
            Self::name(),
            serde_json::to_string(&time).unwrap()
        );

        Ok(time)
    }

    async fn save_logs(
        netid: u64,
        logs: Vec<Web3Log>,
        web3: &Web3<Http>,
        state: &State,
    ) -> AnyResult<()> {
        let mut block;
        let mut batch = vec![];
        let mut height = 0;

        let logp = |batch: &mut Vec<LogWithInfo>, rows: u64, nexth: u64| {
            warn!(
                "{}.{} save {}'s {} logs to db: {} rows, next_height: {}",
                netid,
                Self::name(),
                batch[0].0,
                batch.len(),
                rows,
                nexth,
            );
            batch.clear();
            batch.shrink_to_fit();
        };

        for log in logs.into_iter().filter(|l| !l.is_removed()) {
            if let Some(block_number) = log.block_number {
                assert!(height <= block_number.as_u64(), "logs unsorted");

                // store a batch
                if height < block_number.as_u64() && !batch.is_empty() {
                    let n = Self::save_logs_to_db(netid, &batch, state).await?;
                    logp(&mut batch, n, block_number.as_u64());
                }

                if let Some(b) = web3
                    .eth()
                    .block(BlockId::Number(block_number.into()))
                    .await?
                {
                    if b.hash == log.block_hash {
                        block = Some(b);
                        height = block_number.as_u64();
                        let b = block.as_ref().unwrap();

                        let rawlog = ethabi::RawLog {
                            topics: log.topics,
                            data: log.data.0,
                        };
                        match Self::parse_log(rawlog) {
                            Ok((name, json)) => {
                                if log.transaction_index.is_some()
                                    && log.log_index.is_some()
                                    && b.hash.is_some()
                                {
                                    batch.push((
                                        block_number.as_u32(),
                                        log.transaction_index.unwrap().as_u32(),
                                        log.log_index.unwrap().as_u32(),
                                        b.hash.unwrap().to_string(),
                                        name,
                                        b.timestamp.as_u64(),
                                        serde_json::to_string(&json).unwrap(),
                                    ));
                                } else {
                                    error!(
                                        "block {}-{:?} log's some field is null, giveup: txidx: {:?}, logidx: {:?}",
                                        height, b.hash, log.transaction_index, log.log_index
                                    );
                                }
                            }
                            Err(e) => error!("parse_log failed: {}", e),
                        }
                    }
                }
            }
        }

        let n = Self::save_logs_to_db(netid, &batch, state).await?;
        logp(&mut batch, n, 0);

        Ok(())
    }

    async fn save_logs_to_db(netid: u64, logs: &Vec<LogWithInfo>, state: &State) -> AnyResult<u64> {
        use crate::sqlx::{Done, Executor};

        let mut sql = "insert into contract_events (netid, height, txidx, logidx, blockhash, contract, name, time, data) values ".to_string();
        for i in 0..logs.len() {
            let log = &logs[i];

            let value = format!(
                "({}, {}, {}, {}, '{}', '{}', '{}', to_timestamp({}), '{}')",
                netid,
                log.0,
                log.1,
                log.2,
                log.3,
                Self::name(),
                log.4,
                log.5,
                log.6
            );
            sql.push_str(&value);

            if i + 1 < logs.len() {
                sql.push(',')
            }
        }

        sql.push_str(" on conflict (netid, height, txidx, logidx) do nothing");

        let rows_affected = state.pg.execute(sql.as_str()).await?.rows_affected();

        Ok(rows_affected)
    }

    async fn select_addrs(
        netid: u64,
        date: &DateTime,
        next_date: &DateTime,
        state: &State,
    ) -> AnyResult<Vec<String>> {
        let logs = sqlx::query!(
            "select data ->> 'operator' as addr, count(id) from contract_events where netid = $1 and contract = $2 and time > $3 and time < $4 group by data ->> 'operator' limit 100000;",
            netid as i16,
            Self::name(),
            date,
            next_date
        )
        .fetch_all(&state.pg)
        .await?;
        Ok(logs.into_iter().filter_map(|r| r.addr).collect())
    }
    async fn get_events(
        netid: u64,
        date: &DateTime,
        addr: &str,
        state: &State,
    ) -> AnyResult<Vec<ContractEvent>> {
        let addr = addr.to_lowercase();
        let logs = sqlx::query_as!(
            ContractEvent,
            "select * from contract_events where netid = $1 and contract = $2 and time < $3 and data ->> 'operator' = $4  order by height, txidx, logidx;",
            netid as i16,
            Self::name(),
            date,
            &addr
        )
        .fetch_all(&state.pg)
        .await?;

        Ok(logs)
    }
    async fn analyze_events(
        netid: u64,
        addr: &str,
        events: Vec<ContractEvent>,
        web3: &Web3<Http>,
        state: &State,
    ) -> AnyResult<AnalyzeRes>;
    async fn save_analyze_stats(
        netid: u64,
        events: Vec<AnalyzeRes>,
        _web3: &Web3<Http>,
        state: &State,
    ) -> AnyResult<u64> {
        use crate::sqlx::{Done, Executor};

        let mut sql =
            "insert into operatorstats (netid, height, contract, addr, time, data) values "
                .to_string();
        for i in 0..events.len() {
            let e = &events[i];

            let value = format!(
                "({}, {}, '{}', '{}', '{}', '{}')",
                netid,
                e.1,
                Self::name(),
                e.0,
                e.2,
                e.3
            );
            sql.push_str(&value);

            if i + 1 < events.len() {
                sql.push(',')
            }
        }

        sql.push_str(" on conflict (netid, contract, time) do nothing");

        let rows_affected = state.pg.execute(sql.as_str()).await?.rows_affected();

        Ok(rows_affected)
    }
}

#[async_trait]
impl ContractLogHandle for keep_bonding::Contract {
    fn name() -> &'static str {
        Self::name()
    }
    fn parse_log(rawlog: RawLog) -> ethabi::Result<(&'static str, serde_json::Value)> {
        keep_bonding::parse_log(rawlog)
    }
    fn event_hashs() -> Vec<(&'static str, ethabi::Hash)> {
        keep_bonding::event_hashs()
    }

    async fn analyze_events(
        _netid: u64,
        addr: &str,
        events: Vec<ContractEvent>,
        _web3: &Web3<Http>,
        _state: &State,
    ) -> AnyResult<AnalyzeRes> {
        let mut ec = std::collections::HashMap::new();
        let mut rmap = std::collections::BTreeMap::new();
        let mut ina = 0i128;
        let mut outa = 0i128;
        let mut uba = 0i128;
        let mut ba = 0i128;
        let mut sa = 0i128;

        let p = |s: &str| crate::util::parse::<_, U256>(s);
        for l in &events {
            *ec.entry(&l.name).or_insert(0usize) += 1;

            let lo = l.data["operator"].as_str().unwrap();
            let is = lo == addr;
            let la = l.data["amount"].as_str();

            if l.name == "UnbondedValueDeposited" && is {
                let a = p(la.unwrap()).unwrap();
                uba += a.as_u128() as i128;
                ina += a.as_u128() as i128;
            }

            if l.name == "UnbondedValueWithdrawn" && is {
                let a = p(la.unwrap()).unwrap();
                uba -= a.as_u128() as i128;
                ina -= a.as_u128() as i128;
                outa += a.as_u128() as i128;
            }
            if l.name == "BondCreated" && is {
                let a = p(la.unwrap()).unwrap();
                let j = l.data["reference_id"].as_str().unwrap().to_string();
                rmap.insert(j, a);

                uba -= a.as_u128() as i128;
                ba += a.as_u128() as i128;

                // if l.data["sortition_pool"].as_str().unwrap() != "0x20F1f14a42135d3944fEd1AeD2bE13b01c152054".to_lowercase() {
                //     println!("BondCreated with other sortition_pool: {}", a.as_u128());
                // }
            }

            if l.name == "BondSeized" && is {
                let a = p(la.unwrap()).unwrap();
                let j = l.data["reference_id"].as_str().unwrap().to_string();
                ba -= a.as_u128() as i128;
                sa += a.as_u128() as i128;
                rmap.get_mut(&j).map(|ra| *ra -= a);
            }

            if l.name == "BondReleased" && is {
                let j = l.data["reference_id"].as_str().unwrap().to_string();
                let a = rmap.get_mut(&j).unwrap();
                ba -= a.as_u128() as i128;
                uba += a.as_u128() as i128;
                *a = 0.into();
            }
        }

        let last = events.last().unwrap();

        let json = serde_json::to_string(&json!({
            "deposited": ina.to_string(),
            "withdrawn": outa.to_string(),
            "unbondedValue": uba.to_string(),
            "bondedValue": ba.to_string(),
            "SeizedValue": sa.to_string(),
        }))
        .unwrap();

        // println!("{}.{}'s {}.stats: {:?}", _netid, Self::name(), addr, json);

        Ok((addr.to_owned(), last.height, last.time, json))
    }
}

#[async_trait]
impl ContractLogHandle for token_staking::Contract {
    fn name() -> &'static str {
        Self::name()
    }
    fn parse_log(rawlog: RawLog) -> ethabi::Result<(&'static str, serde_json::Value)> {
        token_staking::parse_log(rawlog)
    }
    fn event_hashs() -> Vec<(&'static str, ethabi::Hash)> {
        token_staking::event_hashs()
    }

    async fn analyze_events(
        _netid: u64,
        addr: &str,
        events: Vec<ContractEvent>,
        _web3: &Web3<Http>,
        _state: &State,
    ) -> AnyResult<AnalyzeRes> {
        let mut ec = std::collections::HashMap::new();
        let mut amount = 0i128;
        let mut inita = 0i128;
        // let mut outa = 0i128;
        let mut seizeda = 0i128;
        let mut slasheda = 0i128;
        let mut undelegated_at = None;
        let mut created_at = 0;

        let p = |s: &str| crate::util::parse::<_, U256>(s);
        for l in &events {
            *ec.entry(&l.name).or_insert(0usize) += 1;

            let lo = l.data["operator"].as_str().unwrap();
            let is = lo == addr;
            let la = l.data["amount"].as_str();
            let la_updated = l.data["new_amount"].as_str();

            if l.name == "OperatorStaked" && is {
                let a = p(l.data["value"].as_str().unwrap()).unwrap();
                amount = a.as_u128() as i128;
                inita = a.as_u128() as i128;
                created_at = l.time.timestamp();
            }

            if l.name == "TopUpCompleted" && is {
                let a = p(la_updated.unwrap()).unwrap();
                amount = a.as_u128() as i128;
            }

            if l.name == "TokensSeized" && is {
                let a = p(la.unwrap()).unwrap();
                amount -= a.as_u128() as i128;
                seizeda += a.as_u128() as i128;
            }

            if l.name == "TokensSlashed" && is {
                let a = p(la.unwrap()).unwrap();
                amount -= a.as_u128() as i128;
                slasheda += a.as_u128() as i128;
            }

            if l.name == "Undelegated" && is {
                let at = p(l.data["undelegated_at"].as_str().unwrap()).unwrap();
                undelegated_at = Some(at.as_u64());
            }
        }

        let last = events.last().unwrap();

        let json = serde_json::to_string(&json!({
            "amount": amount.to_string(),
            "initValue": inita.to_string(),
            "seizedValue": seizeda.to_string(),
            "slashedValue": slasheda.to_string(),
            "createdAt": created_at,
            "undelegatedAt": undelegated_at,
        }))
        .unwrap();

        // println!("{}.{}'s {}.stats: {:?}", _netid, Self::name(), addr, json);

        Ok((addr.to_owned(), last.height, last.time, json))
    }
}
