CREATE TABLE IF NOT EXISTS "cart"
(
    id         UUID primary key not null default uuid_generate_v4(),
    client_id  UUID             not null,
    created_at timestamptz      not null default now()
);

CREATE TABLE IF NOT EXISTS "cart_products"
(
    id         UUID primary key not null default uuid_generate_v4(),
    cart_id    UUID             not null,
    product_id UUID             not null,
    quantity   int              not null,
    comments   varchar,
    created_at timestamptz      not null default now(),
    CONSTRAINT fk_product
        FOREIGN KEY (product_id)
            REFERENCES product (id)
)

