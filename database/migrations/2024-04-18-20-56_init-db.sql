CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABlE "order"
(
    id           UUID primary key not null default uuid_generate_v4(),
    total_amount decimal          not null,
    created_at   timestamptz      not null default now(),
    status       varchar(50)      not null
);

DROP TABLE IF EXISTS "product";
CREATE TABLE "product"
(
    id          UUID primary key not null default uuid_generate_v4(),
    name        varchar(255)     not null,
    description varchar(1000)    not null,
    category    varchar(50)      not null,
    price       decimal          not null,
    created_at  timestamptz      not null default now(),
    deleted_at  timestamptz               default null
)

