use super::*;
#[serde(rename_all = "camelCase")]
#[derive(FromRow, Serialize, Deserialize, Debug)]
pub struct OperatorStats {
    #[serde(skip_serializing)]
    pub id: i64,
    pub netid: i16,
    pub height: i32,
    pub contract: String,
    pub addr: String,
    pub time: DateTime,
    #[serde(skip_serializing)]
    pub create_dt: DateTime,
    pub data: serde_json::Value,
}
