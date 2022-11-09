CREATE DATABASE persons;
CREATE ROLE program WITH PASSWORD 'test';
GRANT ALL PRIVILEGES ON DATABASE persons TO program;
ALTER ROLE program WITH LOGIN;

CREATE TABLE persons (
    "Id" serial PRIMARY KEY,
    "Name" varchar(50) NOT NULL,
    "Age" int,
    "Address" varchar(100),
    "Work" varchar(50)
)