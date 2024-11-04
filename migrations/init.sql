-- Drop and create `supplier` table
DROP TABLE IF EXISTS "suppliers";
CREATE TABLE "suppliers" (
    id serial PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    leadtime_max INT DEFAULT 0,
    leadtime_avg INT DEFAULT 0
);

-- Drop and create `user` table
DROP TABLE IF EXISTS "users";
CREATE TABLE "users" (
    id serial PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Drop and create `product` table
DROP TABLE IF EXISTS "products";
CREATE TABLE "products" (
    id serial PRIMARY KEY,
    supplier_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    stock INT DEFAULT 0,
    price INT DEFAULT 0
);

-- Drop and create `transaction` table
DROP TABLE IF EXISTS "transactions";
CREATE TABLE "transactions" (
    id serial PRIMARY KEY,
    user_id INT NOT NULL,
    product_id INT NOT NULL,
    total_price INT NOT NULL,
    quantity INT NOT NULL
);