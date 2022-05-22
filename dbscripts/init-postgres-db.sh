#!/bin/bash
set -e

echo "creating postgres db"
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE gographql;
    GRANT ALL PRIVILEGES ON DATABASE gographql TO root;

    \c gographql

    CREATE SCHEMA acy;

    SET SCHEMA 'acy';

    CREATE TABLE IF NOT EXISTS acy.books (
        id serial NOT NULL,
        title VARCHAR(255),
        author VARCHAR(255),
        publisher VARCHAR(255),
        CONSTRAINT "PK_tbl_books" PRIMARY KEY (id)
    );

    INSERT INTO acy.books ("title", "author", "publisher") VALUES
    ('Blue Train','John Coltrane','publisher 01')
    ,('Jeru','Gerry Mulligan','publisher 02')
    ,('Sarah Vaughan and Clifford Brown','Sarah Vaughan','publisher 03')
    ;
EOSQL