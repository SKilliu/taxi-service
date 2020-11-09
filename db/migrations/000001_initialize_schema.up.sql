CREATE TABLE IF NOT EXISTS users (
    id varchar(36) primary key,
    name varchar(255) not null,
    hashed_password varchar(255) not null,
    email varchar(255) unique,
    profile_image_url varchar(255) not null
);