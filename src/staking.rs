use rust_decimal::Decimal;
use web3::contract::{Contract, Options};
use web3::types::{H160, U256};
use web3::{
    contract::tokens::Detokenize,
    transports::Http,
    types::{BlockNumber, FilterBuilder},
    Web3,
};

use std::time::Duration;
use tokio::task::spawn;
use tokio::time::delay_for as sleep;

use crate::how::AnyResult;
use crate::state::AppStateRaw as State;
use crate::token::*;
use crate::util::*;

// stakingAbi: keep-core/pkg/chain/gen/abi/TokenStaking.go:31
const STAKING_ABI: &str = include_str!("token.abi.staking.keep.json");
const BONDING_ABI: &str = include_str!("token.abi.bonding.eth.json");

/*
# mainnet
curl 'https://api.etherscan.io/api?module=stats&action=tokensupply&contractaddress=0x85eee30c52b0b379b046fb0f85f4f3dc3009afec'
curl 'https://api.etherscan.io/api?module=stats&action=tokensupply&contractaddress=0x8dAEBADE922dF735c38C80C7eBD708Af50815fAa'
# ropsten
curl 'https://api-ropsten.etherscan.io/api?module=stats&action=tokensupply&contractaddress=0x343d3dda00415289cdd4e8030f63a4a5a2548ff9'
curl 'https://api-ropsten.etherscan.io/api?module=stats&action=tokensupply&contractaddress=0x7c07c42973047223f80c4a69bb62d5195460eb5f'
*/
pub async fn poll_stakingstats(state: &State) {
    let config = &state.config;
    // for token in &config.token {
    //     let rpc = match token.netid {
    //         1 => &config.mainnet_rpc,
    //         3 => &config.ropsten_rpc,
    //         other => {
    //             error!("unsupported netid: {}, skip", other);
    //             continue;
    //         }
    //     };

    //     spawn(loop_poll(TOKEN_KEEP, token.netid, token.contract_address_keep.to_owned(), rpc.to_string(), state.clone()));
    //     spawn(loop_poll(TOKEN_TBTC, token.netid, token.contract_address_tbtc.to_owned(), rpc.to_string(), state.clone()));
    // }
    info!(
        "{:?}",
        poll(
            TOKEN_KEEP,
            3,
            "0x234d2182B29c6a64ce3ab6940037b5C8FdAB608e",
            "0x859e07CeA4E6F88767b000e222A8276e6Bd13fA9", // Seized
            // "0x502a2D67c6EfB0A4e1A41257f8386b3E65b44920", // 大户
            &config.ropsten_rpc,
            state
        )
        .await
    );
}

// async fn loop_poll(token: &'static str, netid: u64, ca: String, rpc: String, state: State) {
//     loop {
//         match poll(token, netid, &ca, &rpc, &state).await {
//             Ok(rows) => {
//                 info!("poll({}, {}, {}) totalSupply rows_affected: {}", rpc, token, ca, rows);
//                 sleep(Duration::from_secs(60 * 60)).await
//             }
//             Err(e) => {
//                 error!("poll({}, {}, {}) totalSupply failed: {}", rpc, token, ca, e);
//                 sleep(Duration::from_secs(60 * 1)).await;
//             }
//         }
//     }
// }

async fn poll(token: &str, netid: u64, ca: &str, oa: &str, rpc: &str, state: &State) -> AnyResult<u64> {
    use crate::sqlx::Done;

    let web3 = Web3::new(Http::new(rpc)?);
    let (amount, created_at, undelegated_at) = get_gelegation_info(&web3, ca, oa).await?;
    info!("{}.{}.{}: amount: {}, created_at: {}, undelegated: {}", netid, token, oa, amount, created_at, undelegated_at);

    let supply = get_unbonded_value(&web3, "0x60535a59b4e71f908f3feb0116f450703fb35ed8", oa).await?;
    info!("{}.{} ETH unbondedValue: {}", netid, token, supply);

    let supply = get_bonded_value(&web3, "0x60535a59b4e71f908f3feb0116f450703fb35ed8", oa).await?;
    info!("{}.{} ETH bondedValue: {}", netid, token, supply);

    Ok(0)
}

async fn get_gelegation_info(web3: &Web3<Http>, ca: &str, oa: &str) -> AnyResult<(Decimal, u64, u64)> {
    let ca: H160 = parse(ca)?;
    let oa: H160 = parse(oa)?;
    let contract = Contract::from_json(web3.eth(), ca, STAKING_ABI.as_bytes())?;
    let (amount, created_at, undelegated_at): (U256, U256, U256) = contract.query("getDelegationInfo", (oa,), None, Options::default(), None).await?;
    Ok(((amount / 10u128.pow(15)).to_string().parse()?, created_at.as_u64(), undelegated_at.as_u64()))
}

async fn get_unbonded_value(web3: &Web3<Http>, ca: &str, oa: &str) -> AnyResult<Decimal> {
    let ca: H160 = parse(ca)?;
    let oa: H160 = parse(oa)?;
    let contract = Contract::from_json(web3.eth(), ca, BONDING_ABI.as_bytes())?;

    let balance: U256 = contract.query("unbondedValue", (oa,), None, Options::default(), None).await?;
    Ok((balance / 10u128.pow(15)).to_string().parse()?)
}

async fn get_bonded_value(web3: &Web3<Http>, ca: &str, oa: &str) -> AnyResult<Decimal> {
    let ca: H160 = parse(ca)?;
    let oa: H160 = parse(oa)?;
    let sortition: H160 = parse("0x20F1f14a42135d3944fEd1AeD2bE13b01c152054")?;
    let contract = Contract::from_json(web3.eth(), ca, BONDING_ABI.as_bytes())?;

    let event_abi = contract.abi().event("BondCreated").unwrap();
    let filter = FilterBuilder::default()
        .topics(Some(vec![event_abi.signature()]), None, None, None)
        .from_block(BlockNumber::Earliest)
        .build();

    type R = (H160, H160, H160, U256, U256);
    let logs = web3.eth().logs(filter).await?;
    let nlogs = logs.len();

    let mut amount = U256::from(0);
    let mut amount2 = U256::from(0);
    let mut count = 0;
    for log in logs {
        let event = R::from_tokens(
            event_abi
                .parse_log(ethabi::RawLog { topics: log.topics, data: log.data.0 })?
                .params
                .into_iter()
                .map(|x| x.value)
                .collect::<Vec<_>>(),
        )?;
        if event.0 == oa && event.2 == sortition {
            let balance: U256 = contract.query("bondAmount", (oa, event.1, event.3), None, Options::default(), None).await?;
            amount += balance;
            amount2 = event.4;
            count += 1;
        }
    }

    info!("oa-{}: {} logs, amount: {}, amount2: {}, count: {}", oa, nlogs, amount, amount2, count);

    let event_abi = contract.abi().event("BondReassigned").unwrap();
    let filter = FilterBuilder::default()
        .topics(Some(vec![event_abi.signature()]), None, None, None)
        .from_block(BlockNumber::Earliest)
        .build();

    type R2 = (H160, U256, H160, U256);
    let logs = web3.eth().logs(filter).await?;
    let nlogs = logs.len();

    let mut amount1 = U256::from(0);
    for log in logs {
        let event = R2::from_tokens(
            event_abi
                .parse_log(ethabi::RawLog { topics: log.topics, data: log.data.0 })?
                .params
                .into_iter()
                .map(|x| x.value)
                .collect::<Vec<_>>(),
        )?;
        if event.0 == oa && event.2 == sortition {
            let balance: U256 = contract.query("bondAmount", (oa, event.2, event.3), None, Options::default(), None).await?;
            amount1 += balance;
        }
    }

    info!("oa-{}: {} logs, amount: {}, amount2: {}", oa, nlogs, amount1, amount2);

    Ok((amount / 10u128.pow(15)).to_string().parse()?)
}
