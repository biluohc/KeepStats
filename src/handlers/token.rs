use actix_web::{get, web, Responder};
// use serde_json::Value;

use crate::{api::ApiResult, models::token::TokenStats, state::AppState, token::*};

#[serde(rename_all = "camelCase")]
#[derive(Serialize, Deserialize, Debug)]
pub struct Form {
    pub token: String,
    pub netid: i16,
    pub days: u16,
}

impl Form {
    fn fix_and_check<T: serde::Serialize>(&mut self) -> Result<(), ApiResult<Vec<T>>> {
        self.token = self.token.to_uppercase();

        let (code, msg) = if ![TOKEN_KEEP, TOKEN_TBTC].contains(&self.token.as_str()) {
            (400, "invalid token")
        } else if self.netid > 3 {
            (400, "invalid ethereum netid")
        } else if self.days <= 0 || self.days > 90 {
            (400, "invalid days")
        } else {
            return Ok(());
        };

        Err(ApiResult::new().code(code).with_msg(msg))
    }
}

// curl  -v  'localhost:8080/api/tokenstats?netid=3&token=tbtc&days=10' | jq .
#[get("/tokenstats")]
async fn tokenstats(state: AppState, form: web::Query<Form>) -> impl Responder {
    let mut form = form.into_inner();
    if let Err(e) = form.fix_and_check() {
        return e;
    }

    let ps = match sqlx::query_as!(
        TokenStats,
        r#"
    SELECT id, netid, token, date, create_dt, total_supply
    FROM tokenstats
    where netid = $1 and token = $2 and date >= now() - $3 * INTERVAL '1day'
            ;"#,
        form.netid,
        form.token,
        form.days as f64
    )
    .fetch_all(&state.pg)
    .await
    {
        Ok(ps) => ps,
        Err(e) => {
            error!("get tokenstats by {:?} failed: {}", form, e);
            return ApiResult::new().code(404).with_msg(e.to_string());
        }
    };
    debug!("get tokenstats by {:?} ok: {}", form, ps.len());

    ApiResult::new().with_data(ps)
}
