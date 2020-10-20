-- local: timestamp
-- UTC: timestamptz
create table users (
    id serial8 primary key,
    name varchar(10) UNIQUE not null,
    email varchar(20) UNIQUE not null,
    pass varchar(65) not null, -- 'passwd hash'
    create_dt timestamptz not null default current_timestamp, -- 'create datetime'
    update_dt timestamptz not null default current_timestamp -- 'udpate datetime'
    -- status char(10) not null default 'nomal' -- 'status: normal, blocked',
);

-- COMMENT ON COLUMN "users"."pass" IS 'passwd hash'

-- Must to use your own username and password
-- docker run -it -d --name postgresql -e POSTGRES_USER=template -e POSTGRES_DB=templatedb  -e POSTGRES_PASSWORD=MTcwNzUyNzIzMDY4Nzk2MzQ3Mjg= -p 5432:5432  postgres:12

-- libpg-dev/postgresql-devel
-- pip3 install pgcli
-- pgcli postgres://template:MTcwNzUyNzIzMDY4Nzk2MzQ3Mjg=@localhost:5432/templatedb
-- create table users
-- \db; \l+; \di; \d users;

-- DATABASE_URL="postgres://template:MTcwNzUyNzIzMDY4Nzk2MzQ3Mjg=@localhost:5432/templatedb"
