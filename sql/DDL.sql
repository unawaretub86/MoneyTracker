CREATE TABLE MT_USERS (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    username VARCHAR(30) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    birth_date varchar(10) NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE MT_CATEGORIES (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    description VARCHAR(30) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES MT_USERS(id)
);

CREATE TABLE MT_BUDGETS (
    id SERIAL PRIMARY KEY,
    category_id INTEGER NOT NULL,
    amount NUMERIC(15, 2) NOT NULL,
    is_last_update BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES MT_CATEGORIES(id)
);

CREATE TABLE MT_EXPENSES (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    month INTEGER NOT NULL,
    year INTEGER not null,
    amount NUMERIC(15, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES MT_USERS(id),
    FOREIGN KEY (category_id) REFERENCES MT_CATEGORIES(id)
);

CREATE TABLE MT_INCOMES (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    amount NUMERIC(15, 2) NOT NULL,
    month INTEGER NOT NULL,
    year INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES MT_USERS(id)
);

CREATE TABLE MT_BALANCE_SNAPSHOT (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    income NUMERIC(15, 2) NOT NULL,
    category_1 VARCHAR(30) not null,
    category_2 VARCHAR(30) null,
    category_3 VARCHAR(30) null,
    category_4 VARCHAR(30) null,
    category_5 VARCHAR(30) null,
    category_6 VARCHAR(30) null,
    category_7 VARCHAR(30) null,
    category_8 VARCHAR(30) null,
    category_9 VARCHAR(30) null,
    category_10 VARCHAR(30) null,
    expense_1 NUMERIC(15, 2) NOT NULL,
    expense_2 NUMERIC(15, 2) NULL,
    expense_3 NUMERIC(15, 2) NULL,
    expense_4 NUMERIC(15, 2) NULL,
    expense_5 NUMERIC(15, 2) NULL,
    expense_6 NUMERIC(15, 2) NULL,
    expense_7 NUMERIC(15, 2) NULL,
    expense_8 NUMERIC(15, 2) NULL,
    expense_9 NUMERIC(15, 2) NULL,
    expense_10 NUMERIC(15, 2) NULL,
    available NUMERIC(15, 2) not NULL,
    month INTEGER NOT NULL,
    year INTEGER NOT NULL,
    is_last_update BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES MT_USERS(id)
);

CREATE TABLE MT_RESIDUAL (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    month INTEGER NOT NULL,
    year INTEGER NOT NULL,
    amount NUMERIC(15, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES MT_USERS(id)
);

CREATE TABLE MT_RESIDUAL_DETAIL (
    id SERIAL PRIMARY KEY,
    residual_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    amount NUMERIC(15, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (residual_id) REFERENCES MT_RESIDUAL(id),
    FOREIGN KEY (category_id) REFERENCES MT_CATEGORIES(id)
);