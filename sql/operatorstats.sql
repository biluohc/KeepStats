create table operatorstats (
    id serial8 primary key,
    netid int2 not NULL,
    height int4 not NULL,
    contract varchar(20) not null,
    addr VARCHAR(66) not NULL,
    time timestamptz not NULL,
    create_dt timestamptz not null default current_timestamp,
    -- 'create datetime'
    data jsonb NOT NULL,
    CONSTRAINT osp UNIQUE (netid, contract, time)
);

-- PG's distinct on liks MySQL/SQLite's any_value
select *
from (
        select distinct on (addr) addr,
            time,
            data
            from(
                select *
                from operatorstats
                where netid = 3
                    and contract = 'keep_bonding'
            )a
            order by addr, time desc
    ) b
order by addr,time desc
limit 10
