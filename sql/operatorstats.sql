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

-- PG's distinct liks MySQL/SQLite's any_value
select *
from (
        select distinct on (addr) addr,
            time,
            data
        from (
                select *
                from operatorstats
                where netid = 3
                    and contract = 'keep_bonding'
                order by time desc
                limit 1000000
            ) a
    ) b
order by t ime
limit 10
