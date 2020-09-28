use crate::models::DateTime;

// PBKDF2 < bcrypt < scrypt
fn passhash(name: &str, pass: &str) -> String {
    let namedpass = format!("{}{}", name, pass);
    let hash = bcrypt::hash(namedpass.as_bytes(), bcrypt::DEFAULT_COST).unwrap();
    // info!("{}{}: {}", name, pass, hash);
    hash
}
fn passhash_verify(name: &str, pass: &str, hash: &str) -> bool {
    let namedpass = format!("{}{}", name, pass);
    bcrypt::verify(namedpass.as_bytes(), hash).unwrap()
}

#[derive(FromRow, Serialize, Deserialize, Debug)]
pub struct User {
    pub id: i64,
    pub name: String,
    // pub phone: String,
    pub email: String,
    // not return password
    #[serde(skip_serializing)]
    pub pass: String,
    pub create_dt: DateTime,
    pub update_dt: DateTime,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Login {
    pub name: String,
    pub password: String,
    #[serde(default)]
    pub rememberme: bool,
}

impl Login {
    pub fn verify(&self, hash: &str) -> bool {
        passhash_verify(&self.name, &self.password, hash)
    }
}

#[derive(Debug, Default, Serialize, Deserialize)]
pub struct Claims {
    // username
    pub sub: String,
    pub exp: usize,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Register {
    pub name: String,
    pub email: String,
    pub password: String,
}

impl Register {
    pub fn passhash(&self) -> String {
        passhash(&self.name, &self.password)
    }
}
