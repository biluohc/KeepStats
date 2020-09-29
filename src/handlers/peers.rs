use actix_web::{get, web, Responder};
// use serde_json::Value;

use crate::{api::ApiResult, models::peer::Peer, state::AppState};

#[serde(rename_all = "camelCase")]
#[derive(Serialize, Deserialize, Debug)]
pub struct Form {
    pub kind: String,
    pub netid: i16,
    pub last_active_hours: u16,
}

// curl  -v  'localhost:8080/api/peers?netid=3&kind=keep_core&lastActiveHours=25' | jq .
#[get("/peers")]
async fn peers(state: AppState, form: web::Query<Form>) -> impl Responder {
    let mut form = form.into_inner();
    form.kind = form.kind.to_lowercase();

    if form.netid > 3 {
        return ApiResult::new().code(400).with_msg("invalid ethereum netid");
    }

    let peers = match sqlx::query_as!(
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
    debug!("get peers by {:?} ok: {}", form, peers.len());

    let json = peers.iter().map(|p| state.json_with_location(&p)).collect::<Vec<_>>();

    ApiResult::new().with_data(json)
}
