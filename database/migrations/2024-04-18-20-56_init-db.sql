CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABlE "order"
(
    id           UUID primary key not null default uuid_generate_v4(),
    total_amount decimal          not null,
    createdAt    timestamptz      not null default now(),
    status       varchar(50)      not null
);