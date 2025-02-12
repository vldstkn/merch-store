CREATE TYPE product_type AS ENUM ('t-shirt', 'cup', 'book',
    'pen', 'powerbank', 'hoody', 'umbrella', 'socks', 'wallet', 'pink-hoody');
CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,
    price BIGINT NOT NULL CHECK (price > 0),
    type product_type NOT NULL
);
CREATE TABLE users(
    id BIGSERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    balance BIGINT CHECK (balance >= 0) DEFAULT 1000
);
CREATE TABLE products_users(
    user_id BIGINT NOT NULL REFERENCES users(id),
    product_id BIGINT NOT NULL REFERENCES products(id),
    quantity BIGINT DEFAULT 0,
    PRIMARY KEY (user_id, product_id)
);
CREATE TABLE transfers(
    id BIGSERIAL PRIMARY KEY,
    from_user_id BIGINT NOT NULL REFERENCES users(id),
    to_user_id BIGINT NOT NULL REFERENCES users(id),
    amount BIGINT NOT NULL CHECK (amount > 0),
    CHECK (from_user_id != transfers.to_user_id)
);