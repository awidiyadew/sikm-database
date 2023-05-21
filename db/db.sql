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

INSERT INTO categories (name)
VALUES 
    ('Snack'),
    ('Minuman'),
    ('Fashion');

INSERT INTO products (name, price, category_id)
VALUES 
    ('Beng-beng', 2500, 1),
    ('Top', 1500, 1),
    ('Larutan Kaki Tiga', 5000, 2);
