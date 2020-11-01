use crate::state::*;
use crate::state::{redis::Client, KvPool, RedisConnectionManager};

use std::path::PathBuf;
use std::sync::Arc;

#[derive(Serialize, Deserialize, Debug, Clone, Default)]
pub struct Config {
    pub pg: String,
    pub redis: String,
    pub listen: String,
    pub jwt_priv: String,
    pub maxminddb: String,
    pub request_timeout: u64,
    pub mainnet_rpc: String,
    pub ropsten_rpc: String,
    pub keep: KeepConfig,
    pub token: Vec<TokenConfig>,
}

#[derive(Serialize, Deserialize, Debug, Clone, Default)]
pub struct TokenConfig {
    pub netid: u64,
    pub contract_address_keep: String,
    pub contract_address_tbtc: String,
    pub contract_address_token_staking: String,
    pub contract_address_keep_bonding: String,
}

#[derive(Serialize, Deserialize, Debug, Clone, Default)]
pub struct KeepConfig {
    pub urls: Vec<KeepInfo>,
    pub poll_interval: u64,
}

#[derive(Serialize, Deserialize, Debug, Clone, Default)]
pub struct KeepInfo {
    pub netid: u64,
    pub keep_core: Vec<String>,
    pub keep_ecdsa: Vec<String>,
}

#[serde(rename_all = "camelCase")]
#[derive(Debug, Deserialize, PartialEq, Serialize)]
struct DbOptions {
    timeout: u64,
    #[serde(default)]
    server_timezone: String,
}

impl Config {
    pub fn parse_from_file(file: &PathBuf) -> Self {
        use std::fs::read_to_string;

        info!("confp: {}", file.display());
        let confstr = read_to_string(file).expect("confile read");
        json5::from_str(&confstr).expect("confile deser")
    }
    pub async fn into_state(self) -> AppStateRaw {
        info!("config: {:?}", self);
        let mut pool_options = PoolOptions::new();

        if let Some(opstr) = url::Url::parse(&self.pg).expect("Invalid PG URL").query() {
            if let Some(ops) = serde_qs::from_str::<DbOptions>(opstr)
                .map_err(|e| error!("serde_qs::from_str::<DbOptions> failed: {}", e))
                .ok()
            {
                pool_options =
                    pool_options.connect_timeout(std::time::Duration::from_secs(ops.timeout));

                if !ops.server_timezone.is_empty() {
                    // UTC, +00:00, HongKong, etc
                    let set = format!("SET TIME ZONE '{}'", ops.server_timezone.clone());

                    // cannot move out of `set_str`, a captured variable in an `Fn` closure
                    let set_str = unsafe { std::mem::transmute::<_, &'static str>(set.as_str()) };
                    std::mem::forget(set);
                    pool_options = pool_options.after_connect(move |conn| {
                        Box::pin(async move {
                            use crate::sqlx::Executor;
                            conn.execute(set_str).await.map(|_| ())
                        })
                    })
                }
            }
        }

        let pg = pool_options.connect(&self.pg).await.expect("pg open");
        let kvm =
            RedisConnectionManager::new(Client::open(self.redis.clone()).expect("redis open"));
        let maxminddb = maxminddb::Reader::open_readfile(&self.maxminddb).expect("maxminddb open");
        let kv = KvPool::builder().build(kvm);

        Arc::new(State {
            config: self,
            pg,
            kv,
            maxminddb,
        })
    }
    // generate and show config string
    pub fn show() {
        let de: Self = Default::default();
        println!("{}", serde_json::to_string_pretty(&de).unwrap())
    }
}

pub fn version_with_gitif() -> &'static str {
    concat!(
        env!("CARGO_PKG_VERSION"),
        " ",
        env!("VERGEN_COMMIT_DATE"),
        ": ",
        env!("VERGEN_SHA_SHORT")
    )
}

#[derive(structopt::StructOpt, Debug)]
// #[structopt(name = "template")]
#[structopt(version = version_with_gitif())]
pub struct Opt {
    // /// Activate debug mode
    // #[structopt(short, long)]
    // debug: bool,

    // The number of occurrences of the `v/verbose` flag
    /// Verbose mode (-v, -vv, -vvv, etc.)
    #[structopt(short, long, parse(from_occurrences))]
    pub verbose: u8,

    /// Output file
    #[structopt(
        short = "c",
        long = "config",
        parse(from_os_str),
        default_value = "keepstats.json"
    )]
    pub config: PathBuf,
}

impl Opt {
    pub fn parse_from_args() -> (JoinHandle, Self) {
        use structopt::StructOpt;

        let opt: Self = Opt::from_args();

        let level = match opt.verbose {
            0 => LevelFilter::Warn,
            1 => LevelFilter::Info,
            2 => LevelFilter::Debug,
            _more => LevelFilter::Trace,
        };

        let formater = BaseFormater::new()
            .local(true)
            .color(true)
            .level(4)
            .formater(format);
        let filter = BaseFilter::new()
            .starts_with(true)
            .notfound(true)
            .max_level(level)
            .chain("sqlx", LevelFilter::Warn);

        let handle = NonblockLogger::new()
            .filter(filter)
            .unwrap()
            .formater(formater)
            .log_to_stdout()
            .map_err(|e| eprintln!("failed to init nonblock_logger: {:?}", e))
            .unwrap();

        info!("opt: {:?}", opt);

        (handle, opt)
    }
}

use nonblock_logger::{
    log::{LevelFilter, Record},
    BaseFilter, BaseFormater, FixedLevel, JoinHandle, NonblockLogger,
};

pub fn format(base: &BaseFormater, record: &Record) -> String {
    let level = FixedLevel::with_color(record.level(), base.color_get())
        .length(base.level_get())
        .into_colored()
        .into_coloredfg();

    format!(
        "[{} {}#{}:{} {}] {}\n",
        chrono::Local::now().format("%Y-%m-%d %H:%M:%S.%3f"),
        level,
        record.module_path().unwrap_or("*"),
        // record.file().unwrap_or("*"),
        record.line().unwrap_or(0),
        nonblock_logger::current_thread_name(),
        record.args()
    )
}
