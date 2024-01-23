\c selltech;

-- Создание таблицы individuals, если она не существует
CREATE TABLE IF NOT EXISTS
    individuals (
                    uid SERIAL PRIMARY KEY,
                    first_name VARCHAR(255),
                    last_name VARCHAR(255),
                    UNIQUE(uid)
);