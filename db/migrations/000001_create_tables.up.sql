CREATE TABLE urls (
    id integer primary key,
    full_url varchar(100) not null unique,
    short_url varchar(100) not null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
