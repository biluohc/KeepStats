pub fn clean0x(s: &str) -> &str {
    if s.starts_with("0x") {
        return &s[2..];
    }
    s
}
pub fn parse<E, T: std::str::FromStr<Err = E>>(s: &str) -> Result<T, E> {
    let s = clean0x(s.trim());
    s.parse()
}
