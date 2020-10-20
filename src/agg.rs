use crate::state::AppStateRaw as State;

use crate::sqlx::{Done, Executor};
use std::time::Duration;
use tokio::task::spawn;
use tokio::time::delay_for as sleep;

pub fn agg_keepstats(state: &State) {
    spawn(agg_peerstats(state.clone()));
}

async fn agg_peerstats(state: State) {
    loop {
        match state.pg.execute(r#"insert into peerstats (netid, kind, date, online) 
        select netid, kind, s.date, count(distinct(p.ethereum_address)) as online from peers as p join (select (max(date) + interval '1day') as date from peerstats ) as s on s.date >= date(p.create_dt) and s.date <= date(p.update_dt) and s.date < date(now()) GROUP BY netid, kind, s.date
        on conflict (date, netid, kind) do nothing;"#).await.map(|a| a.rows_affected()) {
            Ok(rows) => {
                warn!("agg_peerstats.rows_affected: {}", rows);
                if rows == 0 {
                    sleep(Duration::from_secs(60 * 60 * 1)).await;
                } else {
                    let lasts = sqlx::query!("select * from peerstats where date = (select max(date) as date from peerstats)")
                        .fetch_all(&state.pg)
                        .await;
                    warn!("agg_peerstats.last: {:?}", lasts);
                }
            }
            Err(e) => {
                warn!("agg_peerstats failed: {}", e);
                sleep(Duration::from_secs(10)).await;
            }
        }
    }
}
