CREATE DATABASE qbit_db;

CREATE TABLE users (
    id serial primary key,
    username varchar(200) not null,
    email varchar (200) not null,
    password varchar(200) not null,
    address text not null,
    city varchar(100) not null,
    full_name varchar(200) not null,
    phone_number varchar(15) not null,
    role_id int not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    FOREIGN KEY (role_id) REFERENCES roles(id)
)

CREATE TABLE roles (
    id serial primary key,
    name varchar(100) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
)

CREATE TABLE merchants (
    id serial primary key,
    name varchar(200) not null,
    address varchar(200) not null,
    city varchar(100) not null,
    user_id int not null,
    balance numeric(11, 2) not null default 0,
    FOREIGN KEY (user_id) REFERENCES users(id)
)

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  merchant_id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  price DECIMAL(10, 2) NOT NULL,
  final_price DECIMAL(10, 2) NOT NULL,
  discount DECIMAL(4, 1),
  stock INT NOT NULL,
  category_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT current_timestamp,
  updated_at TIMESTAMP DEFAULT current_timestamp,
    FOREIGN KEY (merchant_id) REFERENCES merchants(id),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE shopping_cart (
   id SERIAL PRIMARY KEY,
   user_id INT NOT NULL,
   created_at TIMESTAMP DEFAULT current_timestamp,
   updated_at TIMESTAMP DEFAULT current_timestamp,
   FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE cart_items (
   id SERIAL PRIMARY KEY,
   cart_id INT NOT NULL,
   product_id INT NOT NULL,
   quantity INT NOT NULL,
   price DECIMAL(10, 2) NOT NULL,
   subtotal DECIMAL(10, 2) NOT NULL,
   created_at TIMESTAMP DEFAULT current_timestamp,
   updated_at TIMESTAMP DEFAULT current_timestamp,
   FOREIGN KEY (cart_id) REFERENCES shopping_cart(id),
   FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE orders (
     id SERIAL PRIMARY KEY,
     user_id INT NOT NULL,
     order_number VARCHAR(200) NOT NULL,
     total_amount DECIMAL(10, 2) NOT NULL,
     status VARCHAR(50) NOT NULL,
     payment_status VARCHAR(50) NOT NULL,
    payment_date TIMESTAMP,
     order_date TIMESTAMP DEFAULT current_timestamp,
     delivery_address TEXT NOT NULL,
     created_at TIMESTAMP DEFAULT current_timestamp,
     updated_at TIMESTAMP DEFAULT current_timestamp,
     FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE order_items (
    order_item_id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    subtotal DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

drop table orders


create function generate_new_order_number() returns text
    language plpgsql
as
$$
DECLARE
    last_id INT;
    new_id INT;
    transaction_code TEXT;
BEGIN
    SELECT COALESCE(MAX(id), 0) INTO last_id FROM orders;

    new_id := last_id + 1;

    transaction_code := 'ORD-' || new_id;

    RETURN transaction_code;
END;
$$;







