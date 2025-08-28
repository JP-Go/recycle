CREATE TABLE company (
    id uuid primary key,
    name varchar not null,
    company_owner_id uuid not null,

    created_at timestampz not null default CURRENT_TIMESTAMP,
    updated_at timestampz not null default CURRENT_TIMESTAMP,
    deleted_at timestampz
    created_by uuid not null,
    updated_by uuid not null
);

CREATE TABLE user(
    id uuid primary key,
    email varchar not null,
    hashed_password varchar not null,
    confirmed_at timestampz,
    role varchar not null,

    created_at timestampz not null default CURRENT_TIMESTAMP,
    updated_at timestampz not null default CURRENT_TIMESTAMP,
    deleted_at timestampz
    created_by uuid not null,
    updated_by uuid not null
);

---- create above / drop below ----

drop table company;
drop table user;
