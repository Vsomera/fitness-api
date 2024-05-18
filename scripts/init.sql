
-- create users table if it does not exist
DO $$ BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM information_schema.tables
        WHERE table_schema = 'public'
        AND table_name = 'users'
    ) THEN
        CREATE TABLE users (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            email VARCHAR(255) NOT NULL,
            password VARCHAR(255) NOT NULL,
            created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
            updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp
        );
    END IF;
END $$;

-- create measurements table if it does not exist
DO $$ BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM information_schema.tables
        WHERE table_schema = 'public'
        AND table_name = 'measurements'
    ) THEN
        CREATE TABLE measurements (
            id SERIAL PRIMARY KEY,
            user_id INTEGER NOT NULL,
            weight FLOAT NOT NULL,
            height FLOAT NOT NULL,
            body_fat FLOAT NOT NULL,
            created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
            FOREIGN KEY (user_id) REFERENCES users(id) -- ref to user
        );
    END IF;
END $$;
