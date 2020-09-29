use super::*;

#[serde(rename_all = "camelCase")]
#[derive(FromRow, Serialize, Deserialize, Debug)]
pub struct Peer {
    // #[serde(default)]
    pub id: i64,
    // #[serde(skip_serializing)]
    pub netid: i16,
    pub kind: String,
    pub network_id: String,
    pub network_ip: String,
    pub network_port: i32,
    pub ethereum_address: String,
    pub create_dt: DateTime,
    pub update_dt: DateTime,
}

impl AsRef<str> for Peer {
    fn as_ref(&self) -> &str {
        &self.network_ip
    }
}
