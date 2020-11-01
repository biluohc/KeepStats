use super::*;
#[serde(rename_all = "camelCase")]
#[derive(FromRow, Serialize, Deserialize, Debug)]
pub struct DbStats {
    pub event_time: Option<DateTime>,
    pub stats_time: Option<DateTime>,
    pub stats_date: Option<DateTime>,
    pub stats_date_next: Option<DateTime>,
    pub should_update: Option<bool>,
    pub should_update_normal: Option<bool>,
}
