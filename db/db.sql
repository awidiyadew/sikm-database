CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10) NOT NULL,
    category_id INT NOT NULL,
    CONSTRAINT fk_categories_id FOREIGN KEY(category_id) REFERENCES categories(id)
);

INSERT INTO categories (name, created_at, updated_at)
VALUES 
    ('Snack', now(), now()),
    ('Minuman', now(), now()),
    ('Fashion', now(), now());

INSERT INTO products (name, price, category_id, created_at, updated_at)
VALUES 
    ('Beng-beng', 2500, 1, now(), now()),
    ('Top', 1500, 1, now(), now()),
    ('Larutan Kaki Tiga', 5000, 2, now(), now());
