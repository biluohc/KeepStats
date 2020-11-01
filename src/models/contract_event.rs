use super::*;
#[serde(rename_all = "camelCase")]
#[derive(FromRow, Serialize, Deserialize, Debug)]
pub struct ContractEvent {
    pub id: i64,
    pub netid: i16,
    pub height: i32,
    pub txidx: i16,
    pub logidx: i16,
    pub blockhash: String,
    pub contract: String,
    pub name: String,
    pub time: DateTime,
    pub create_dt: DateTime,
    pub data: serde_json::Value,
}
