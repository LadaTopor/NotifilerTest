CREATE TABLE users (
                       id VARCHAR PRIMARY KEY,
                       name TEXT,
                       email VARCHAR(255) UNIQUE ,
                       hashed_password TEXT,
                       token TEXT,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE note (
                       id SERIAL PRIMARY KEY,
                       user_id VARCHAR REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
                       title VARCHAR(50),
                       body TEXT,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (email, hashed_password) VALUES
                                           ('Привет', 'Hello')
