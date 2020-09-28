create table peers (
    id serial8 primary key,
    netid int2 not NULL,
    kind varchar(10) not null,
    network_id varchar(60) not null,
    network_addr varchar(60) not null,
    ethereum_address varchar(50) not null,
    create_dt timestamptz not null default current_timestamp, -- 'create datetiem'
    update_dt timestamptz not null default current_timestamp, -- 'udpate datetime'
    CONSTRAINT e2k UNIQUE (ethereum_address, netid, kind)
);

-- INSERT into peers(netid, kind, network_id, network_addr, ethereum_address, create_dt, update_dt) values (),() on conflict (ethereum_address, ethnetid, kind) do update set update_dt=, network_id=, network_addr=

