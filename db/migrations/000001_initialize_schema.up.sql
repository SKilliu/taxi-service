CREATE TABLE IF NOT EXISTS users (
    id varchar(36) primary key,
    name varchar(255) not null,
    account_type varchar(50) not null,
    hashed_password varchar(255) not null,
    email varchar(255) unique,
    profile_image_url varchar(255) not null
);

CREATE TABLE IF NOT EXISTS cars (
    id varchar(36) primary key,
    model varchar(255) not null,
    series varchar(255) not null,
    number varchar(255) not null,
    status varchar(50) not null
);

CREATE TABLE IF NOT EXISTS trips (
    id varchar(36) primary key,
    starting_point_location varchar(255) not null,
    starting_point_longitude float not null,
    starting_point_latitude float not null,
    destination_location varchar(255) not null,
    destination_point_longitude float not null,
    destination_point_latitude float not null,
    distance float not null
);

CREATE TABLE IF NOT EXISTS orders (
    id varchar(36) primary key,
    driver_id varchar(36) not null,
    client_id varchar(36) references users (id) on delete cascade,
    car_id varchar(36) not null,
    trip_id varchar(36) not null,
    price float not null,
    status varchar(50) not null,
    car_arrival_time timestamp default current_timestamp,
    created_at timestamp default current_timestamp
);

CREATE TABLE IF NOT EXISTS driver_cars (
    id varchar(36) primary key,
    driver_id varchar(36) references users (id) on delete cascade,
    car_id varchar(36) references cars (id) on delete cascade
);
