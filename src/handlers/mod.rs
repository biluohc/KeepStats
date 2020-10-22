pub mod peers;
pub mod token;

pub fn init(cfg: &mut actix_web::web::ServiceConfig) {
    cfg.service(peers::peers);
    cfg.service(peers::peerstats);
    cfg.service(token::tokenstats);
}
