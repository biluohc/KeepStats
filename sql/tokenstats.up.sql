create table tokenstats (
    id serial8 primary key,
    netid int2 not NULL,
    token varchar(10) not null,
    date timestamptz not NULL,
    total_supply NUMERIC not null DEFAULT 0,
    create_dt timestamptz not null default current_timestamp, -- 'create datetime'
    CONSTRAINT dnt UNIQUE (date, netid, token)
);

-- insert into tokenstats (netid, token, date, total_supply) values(3, 'TBTC',  '2020-09-28', 0) on conflict (date, netid, token) do nothing
-- insert into tokenstats (netid, token, date, total_supply) values(3, 'KEEP',  '2020-09-28', 1000000000000000000000000000) on conflict (date, netid, token) do nothing
