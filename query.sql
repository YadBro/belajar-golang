-- Active: 1664468138770@@127.0.0.1@3306@belajar_golang_database

CREATE TABLE
    customer (
        id VARCHAR(100) NOT NULL,
        name VARCHAR(100) NOT NULL,
        PRIMARY KEY(id)
    ) ENGINE = InnoDb;

SELECT * FROM customer;

DELETE FROM customer;

ALTER TABLE customer
ADD COLUMN email VARCHAR(100),
ADD
    COLUMN balance INT DEFAULT 0,
ADD
    COLUMN rating DOUBLE DEFAULT 0.0,
ADD
    COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN birth_date DATE,
ADD
    COLUMN married BOOLEAN DEFAULT false;

DESC customer;

INSERT INTO
    customer(
        id,
        name,
        email,
        balance,
        rating,
        birth_date,
        married
    )
VALUES (
        "bambang",
        "BAMBANG",
        "bambang@gmail.com",
        1000000,
        90.0,
        "2022-30-09",
        true
    ), (
        "budi",
        "BUDI",
        "budi@gmail.com",
        5000000,
        85.5,
        "1986-06-10",
        true
    ), (
        "yadi",
        "YADI",
        "yadi@gmail.com",
        8500000,
        100.0,
        "2003-04-11",
        false
    );

UPDATE customer
SET
    email = NULL,
    birth_date = NULL
WHERE id = "bambang";

CREATE TABLE
    user (
        username VARCHAR(100) NOT NULL,
        password VARCHAR(100) NOT NULL,
        PRIMARY KEY(username)
    ) ENGINE = InnoDB;

SELECT * FROM user;

INSERT INTO user(username, password) VALUES('admin', 'admin');

CREATE TABLE
    comments(
        id INT NOT NULL AUTO_INCREMENT,
        email VARCHAR(100) NOT NULL,
        comment TEXT,
        PRIMARY KEY(id)
    ) ENGINE = InnoDB;

DESC comments;

SELECT * FROM comments;

SELECT count(*) FROM comments;