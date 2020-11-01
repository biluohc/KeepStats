use actix_web::{get, web, Responder};
// use serde_json::Value;

use crate::{api::ApiResult, models::peer::*, peers::*, state::AppState};

#[serde(rename_all = "camelCase")]
#[derive(Serialize, Deserialize, Debug)]
pub struct Form {
    pub kind: String,
    pub netid: i16,
    #[serde(default)]
    pub last_active_hours: u16,
    #[serde(default)]
    pub days: u16,
}

impl Form {
    fn fix_and_check<T: serde::Serialize>(&mut self, days: bool) -> Result<(), ApiResult<Vec<T>>> {
        self.kind = self.kind.to_lowercase();

        let (code, msg) = if ![KEEP_CORE, KEEP_ECDSA].contains(&self.kind.as_str()) {
            (400, "invalid kind")
        } else if self.netid > 3 {
            (400, "invalid ethereum netid")
        } else if days && (self.days <= 0 || self.days > 91) {
            (400, "invalid days")
        } else if (!days) && (self.last_active_hours <= 0 || self.last_active_hours > 25) {
            (400, "invalid last_active_hours")
        } else {
            return Ok(());
        };

        Err(ApiResult::new().code(code).with_msg(msg))
    }
}

// curl  -v  'localhost:8080/api/peers?netid=3&kind=keep_core&lastActiveHours=2' | jq .
#[get("/peers")]
async fn peers(state: AppState, form: web::Query<Form>) -> impl Responder {
    let mut form = form.into_inner();
    if let Err(e) = form.fix_and_check(false) {
        return e;
    }

    let ps = match sqlx::query_as!(
        Peer,
        r#"
    SELECT id, netid, kind, network_id, network_ip, network_port, ethereum_address, create_dt, update_dt
    FROM peers
    where netid = $1 and kind = $2 and update_dt > now() - $3 * INTERVAL '1hour'
            ;"#,
        form.netid,
        form.kind,
        form.last_active_hours as f64
    )
    .fetch_all(&state.pg)
    .await
    {
        Ok(ps) => ps,
        Err(e) => {
            error!("get peers by {:?} failed: {}", form, e);
            return ApiResult::new().code(404).with_msg(e.to_string());
        }
    };
    debug!("get peers by {:?} ok: {}", form, ps.len());

    let json = ps
        .iter()
        .map(|p| state.json_with_location(&p))
        .collect::<Vec<_>>();

    ApiResult::new().with_data(json)
}

// curl  -v  'localhost:8080/api/peerstats?netid=3&kind=keep_core&days=30' | jq .
#[get("/peerstats")]
async fn peerstats(state: AppState, form: web::Query<Form>) -> impl Responder {
    let mut form = form.into_inner();
    if let Err(e) = form.fix_and_check(true) {
        return e;
    }

    let ps = match sqlx::query_as!(
        PeerStats,
        r#"
    SELECT id, netid, kind, date, online, create_dt
    FROM peerstats
    where netid = $1 and kind = $2 and date > now() - $3 * INTERVAL '1day'
            ;"#,
        form.netid,
        form.kind,
        form.days as f64
    )
    .fetch_all(&state.pg)
    .await
    {
        Ok(ps) => ps,
        Err(e) => {
            error!("get peerstats by {:?} failed: {}", form, e);
            return ApiResult::new().code(404).with_msg(e.to_string());
        }
    };
    debug!("get peerstats by {:?} ok: {}", form, ps.len());

    ApiResult::new().with_data(ps)
}
