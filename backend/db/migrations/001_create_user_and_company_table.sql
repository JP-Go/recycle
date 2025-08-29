CREATE SCHEMA recycle;

CREATE TABLE recycle.user(
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

CREATE TABLE recycle.company (
    id uuid primary key,
    name varchar not null,
    company_owner_id uuid references recycle.user(id) not null,

    created_at timestamp with time zone not null default CURRENT_TIMESTAMP,
    updated_at timestamp with time zone not null default CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    created_by uuid not null,
    updated_by uuid not null
);

---- create above / drop below ----

drop table recycle.company;
drop table recycle.user;
drop schema recycle cascade;
