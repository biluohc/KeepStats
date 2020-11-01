create table contract_events (
    id serial8 primary key,
    netid int2 not NULL,
    height int4 not NULL,
    txidx int2 not NULL,
    logidx int2 not NULL,
    blockhash VARCHAR(66) not NULL,
    contract varchar(20) not null,
    name varchar(30) not null,
    time timestamptz not NULL,
    create_dt timestamptz not null default current_timestamp, -- 'create datetime'
    data jsonb NOT NULL,
    CONSTRAINT cep UNIQUE (netid, height, txidx, logidx)
);

-- insert into contract_events (netid, height, txidx, logidx, blockhash, contract, name, time, data) values(9, 0, 1, 2, '0xffgg', 'try_test', 'Test0', '2020-09-28', '{"operator":"0xd354634ee2e7dd45f3fbeaa08e65445b75f0fd59","reference_id":"0xe098ad559c0ad369e64e67a4cb88e6994af01d62"}') on conflict (netid, height, txidx, logidx) do nothing
-- select count(*) from (select data ->> 'operator', count(id) from contract_events where netid = 3 and contract = 'keep_bonding' group by  data ->> 'operator' limit 10000) as a
-- select count(*) from (select data ->> 'operator', count(id) from contract_events where netid = 3 and contract = 'token_staking' group by  data ->> 'operator' limit 10000) as a

