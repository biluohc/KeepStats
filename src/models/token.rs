use super::*;
#[serde(rename_all = "camelCase")]
#[derive(FromRow, Serialize, Deserialize, Debug)]
pub struct TokenStats {
    // #[serde(default)]
    pub id: i64,
    pub netid: i16,
    pub token: String,
    pub total_supply: Decimal,
    pub date: DateTime,
    pub create_dt: DateTime,
}
