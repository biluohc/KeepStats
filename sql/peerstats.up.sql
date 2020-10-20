create table peerstats (
    id serial8 primary key,
    netid int2 not NULL,
    kind varchar(10) not null,
    date timestamptz not NULL,
    online int not null DEFAULT 0,
    create_dt timestamptz not null default current_timestamp, -- 'create datetime'
    CONSTRAINT dnk UNIQUE (date, netid, kind)
);

-- insert into peerstats (netid, kind, date, online) values(3, 'keep_core',  '2020-09-28', 0) on conflict (date, netid, kind) do nothing

-- insert into peerstats (netid, kind, date, online) select netid, kind, '2020-09-29' as date, count(distinct(ethereum_address)) as online from peers where date(create_dt) <= date('2020-09-29') and date(update_dt) >= date('2020-09-29') group by netid, kind on conflict (date, netid, kind) do nothing

-- insert into peerstats (netid, kind, date, online) select netid, kind, (select max(date) from peerstats) as date, count(distinct(ethereum_address)) as online from peers where date(create_dt) <= (select max(date) from peerstats as date) and date(update_dt) >= (select max(date) from peerstats as date) group by netid, kind on conflict (date, netid, kind) do nothing

-- select netid, kind, s.date, count(distinct(p.ethereum_address)) as online from peers as p join (select (max(date) + interval '1day') as date from peerstats ) as s on s.date >= date(p.create_dt) and s.date <= date(p.update_dt) and s.date < date(now()) GROUP BY netid, kind, s.date
