-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products(
    id BIGSERIAL PRIMARY KEY,
    type VARCHAR(255) UNIQUE NOT NULL,
    price BIGINT NOT NULL CHECK (price > 0)
);

CREATE TABLE IF NOT EXISTS  users(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    balance BIGINT CHECK (balance >= 0) DEFAULT 1000
);

CREATE TABLE IF NOT EXISTS products_users(
    user_name VARCHAR(255) NOT NULL REFERENCES users(name),
    product_type VARCHAR(255) NOT NULL REFERENCES products(type),
    quantity BIGINT DEFAULT 1,
    PRIMARY KEY (user_name, product_type)
);

CREATE TABLE IF NOT EXISTS  transfers(
    id BIGSERIAL PRIMARY KEY,
    from_user_name TEXT NOT NULL REFERENCES users(name),
    to_user_name TEXT NOT NULL REFERENCES users(name),
    amount BIGINT NOT NULL CHECK (amount > 0),
    CHECK (from_user_name != to_user_name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS products_users CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS products CASCADE;
DROP TABLE IF EXISTS transfers CASCADE;
-- +goose StatementEnd
