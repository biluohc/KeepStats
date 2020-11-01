#![feature(ip)]
// #![feature(associated_type_defaults)]

#[macro_use]
extern crate nonblock_logger;
#[macro_use]
extern crate async_trait;
#[macro_use]
extern crate serde_json;
#[macro_use]
extern crate sqlx;
#[macro_use]
extern crate serde;

use actix_files::Files;
use actix_web::{middleware, web, App, HttpServer};

pub mod agg;
pub mod api;
pub mod config;
pub mod handlers;
pub mod how;
pub mod middlewares;
#[macro_use]
pub mod macros;
pub mod models;
pub mod peers;
pub mod staking;
pub mod state;
pub mod token;
pub mod users;
pub mod util;

use config::{Config, Opt};

#[actix_rt::main]
async fn main() -> std::io::Result<()> {
    // Config::show();
    let (_handle, opt) = Opt::parse_from_args();
    let state = Config::parse_from_file(&opt.config).into_state().await;

    let state2 = state.clone();
    staking::poll_stakingstats(&state).await;
    agg::agg_keepstats(&state2);
    token::poll_tokenstats(&state2);
    peers::poll_peers(&state2);

    HttpServer::new(move || {
        App::new()
            .data(state2.clone())
            .app_data(state2.clone())
            .app_data(web::PathConfig::default().error_handler(api::json_error_handler))
            .app_data(web::JsonConfig::default().error_handler(api::json_error_handler))
            .app_data(web::QueryConfig::default().error_handler(api::json_error_handler))
            .app_data(web::FormConfig::default().error_handler(api::json_error_handler))
            // .wrap(middleware::Compress::new(actix_web::http::ContentEncoding::Br))
            .wrap(middleware::Logger::default())
            .default_service(web::route().to(api::notfound))
            .service(web::scope("/api").configure(handlers::init))
            // .service(web::scope("/user").configure(users::routes::init))
            .service(
                Files::new("/static", "assets")
                    .redirect_to_slash_directory()
                    .show_files_listing()
                    .use_last_modified(true),
            )
    })
    .keep_alive(300)
    .bind(&state.config.listen)?
    .run()
    .await
}
