use actix_web::{get, web, Responder};
// use serde_json::Value;

use crate::{api::ApiResult, models::operator::OperatorStats, peers::*, state::AppState};

#[serde(rename_all = "camelCase")]
#[derive(Serialize, Deserialize, Debug)]
pub struct Form {
    pub kind: String,
    pub netid: i16,
    #[serde(default)]
    pub days: u16,
}

impl Form {
    fn fix_and_check<T: serde::Serialize>(&mut self) -> Result<(), ApiResult<Vec<T>>> {
        self.days = 91;

        let (code, msg) = if ![KEEP_CORE, KEEP_ECDSA].contains(&self.kind.as_str()) {
            (400, "invalid kind")
        } else if self.netid > 3 {
            (400, "invalid ethereum netid")
        } else if self.days <= 0 || self.days > 91 {
            (400, "invalid days")
        } else {
            // kind to contract name
            if self.kind == KEEP_CORE {
                self.kind = crate::staking::token_staking::Contract::name().to_string();
            } else if self.kind == KEEP_ECDSA {
                self.kind = crate::staking::keep_bonding::Contract::name().to_string();
            }

            return Ok(());
        };

        Err(ApiResult::new().code(code).with_msg(msg))
    }
}

// curl  -v  'localhost:8080/api/operatorstats?netid=3&kind=keep_core' | jq .
#[get("/operatorstats")]
async fn operatorstats(state: AppState, form: web::Query<Form>) -> impl Responder {
    let mut form = form.into_inner();
    if let Err(e) = form.fix_and_check() {
        return e;
    }

    let ps = match sqlx::query_as!(
        OperatorStats,
        r#"
        select *
        from (
                select distinct on (addr) *
                from (
                        select *
                        from operatorstats
                        where netid = $1
                            and contract = $2
                    ) a
                    order by addr, time desc
            ) b
        ;"#,
        form.netid,
        form.kind,
    )
    .fetch_all(&state.pg)
    .await
    {
        Ok(ps) => ps,
        Err(e) => {
            error!("get operatorstats by {:?} failed: {}", form, e);
            return ApiResult::new().code(404).with_msg(e.to_string());
        }
    };
    info!("get operatorstats by {:?} ok: {}", form, ps.len());

    ApiResult::new().with_data(ps)
}
