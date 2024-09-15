CREATE TABLE users (
guid UUID PRIMARY KEY,
first_name VARCHAR(50) NOT NULL,
last_name VARCHAR(50) NOT NULL,
email VARCHAR(100) UNIQUE NOT NULL,
password VARCHAR(255) NOT NULL
);

CREATE TABLE refresh_tokens (
user_guid UUID REFERENCES users(guid),
ip_address VARCHAR(45) NOT NULL,
refresh_token_hash VARCHAR(255) NOT NULL,
token_id UUID PRIMARY KEY
);