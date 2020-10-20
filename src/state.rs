pub use mobc_redis::{redis, Connection, RedisConnectionManager};
pub type KvPool = mobc::Pool<RedisConnectionManager>;

pub use sqlx::PgPool;
pub type PoolOptions = sqlx::postgres::PgPoolOptions;

use crate::config::Config;

pub type AppStateRaw = std::sync::Arc<State>;
pub type AppState = actix_web::web::Data<AppStateRaw>;

// #[derive(Clone)]
pub struct State {
    pub maxminddb: maxminddb::Reader<Vec<u8>>,
    pub config: Config,
    pub pg: PgPool,
    pub kv: KvPool,
}

use maxminddb::geoip2::model;
#[derive(Deserialize, Serialize, Clone, Debug)]
pub struct CityIsp<'a> {
    pub city: Option<model::City<'a>>,
    #[serde(borrow)]
    pub country: Option<model::Country<'a>>,
    pub location: Option<model::Location<'a>>,
    pub isp: Option<&'a str>,
}

#[derive(Deserialize, Serialize, Clone, Debug)]
pub struct Location {
    pub country: Option<String>,
    pub city: Option<String>,
    pub latitude: Option<f64>,
    pub longitude: Option<f64>,
}

impl State {
    pub fn lookup(&self, ip: &str) -> crate::how::AnyResult<Location> {
        let ip = ip.trim().parse::<std::net::IpAddr>()?;
        let ifs: CityIsp = self.maxminddb.lookup(ip)?;

        Ok(Location {
            country: ifs.country.as_ref().and_then(|c| c.names.as_ref()).and_then(|ns| ns.get("en")).map(|s| s.to_string()),
            city: ifs.city.as_ref().and_then(|c| c.names.as_ref()).and_then(|ns| ns.get("en")).map(|s| s.to_string()),
            latitude: ifs.location.as_ref().and_then(|loc| loc.latitude).clone(),
            longitude: ifs.location.as_ref().and_then(|loc| loc.longitude).clone(),
        })
    }
    pub fn json_with_location<T>(&self, data: &T) -> serde_json::Value
    where
        T: AsRef<str> + serde::Serialize,
    {
        let mut json = json!(data);
        json["location"] = json!(self.lookup(data.as_ref()).map_err(|e| error!("lookup location for {} failed: {:?}", data.as_ref(), e)).ok());

        json
    }
}
