-- Таблица пользователей
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       login VARCHAR(255) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL
);

-- Таблица заказов
CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        order_id VARCHAR(255) NOT NULL UNIQUE,
                        user_id INTEGER NOT NULL,
                        status VARCHAR(50) NOT NULL,
                        accrual INTEGER NOT NULL DEFAULT 0,
                        uploaded_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Таблица балансов
CREATE TABLE balances (
                          id SERIAL PRIMARY KEY,
                          user_id INTEGER NOT NULL UNIQUE,
                          balance INTEGER NOT NULL DEFAULT 0,
                          with_drawn INTEGER NOT NULL DEFAULT 0,
                          FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Таблица списаний
CREATE TABLE withdrawals (
                             id SERIAL PRIMARY KEY,
                             withdraw_id INTEGER NOT NULL,
                             user_id INTEGER NOT NULL,
                             sum INTEGER NOT NULL,
                             processed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
