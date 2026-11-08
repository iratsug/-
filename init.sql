CREATE TABLE IF NOT EXISTS roles (
    id   TEXT PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS users (
    id                TEXT PRIMARY KEY,
    name              TEXT NOT NULL,
    email             TEXT NOT NULL UNIQUE,
    phone             TEXT,
    role_id           TEXT NOT NULL REFERENCES roles(id),
    registration_date DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS products (
    id          TEXT PRIMARY KEY,
    name        TEXT NOT NULL,
    type        TEXT,
    gender      TEXT,
    size        TEXT,
    color       TEXT,
    description TEXT,
    brand       TEXT,
    season      TEXT,
    price       REAL NOT NULL
);

CREATE TABLE IF NOT EXISTS photos (
    id         TEXT PRIMARY KEY,
    filename   TEXT NOT NULL,
    product_id TEXT NOT NULL REFERENCES products(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS cart (
    id         TEXT PRIMARY KEY,
    user_id    TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    product_id TEXT NOT NULL REFERENCES products(id),
    count      INTEGER NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS orders (
    id        TEXT PRIMARY KEY,
    user_id   TEXT NOT NULL REFERENCES users(id),
    date_time DATETIME NOT NULL,
    status    TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS order_items (
    id         TEXT PRIMARY KEY,
    order_id   TEXT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    product_id TEXT NOT NULL REFERENCES products(id),
    price      REAL NOT NULL
);