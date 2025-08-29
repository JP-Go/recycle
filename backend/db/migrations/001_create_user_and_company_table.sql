CREATE SCHEMA {{.schema}};

CREATE TABLE {{.schema}}.user(
    id uuid primary key,
    email varchar not null,
    hashed_password varchar not null,
    confirmed_at timestamp with time zone,
    role varchar not null,

    created_at timestamp with time zone not null default CURRENT_TIMESTAMP,
    updated_at timestamp with time zone not null default CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    created_by uuid not null,
    updated_by uuid not null
);

CREATE TABLE {{.schema}}.company (
    id uuid primary key,
    name varchar not null,
    company_owner_id uuid references {{.schema}}.user(id) not null,

    created_at timestamp with time zone not null default CURRENT_TIMESTAMP,
    updated_at timestamp with time zone not null default CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    created_by uuid not null,
    updated_by uuid not null
);

---- create above / drop below ----

drop table {{.schema}}.company;
drop table {{.schema}}.user;
drop schema {{.schema}} cascade;
