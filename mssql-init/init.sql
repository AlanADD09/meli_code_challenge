-- mssql-init/init.sql
IF NOT EXISTS (SELECT * FROM sys.databases WHERE name = 'Meli')
BEGIN
    CREATE DATABASE Meli;
END
GO

USE Meli;
GO

IF NOT EXISTS (SELECT * FROM sys.tables WHERE name = 'tbl_data')
BEGIN
    CREATE TABLE tbl_data (
        site_id VARCHAR(10),
        id VARCHAR(50),
        price DECIMAL(18,2),
        date_created DATETIME,
        category_id VARCHAR(50),
        currency_id VARCHAR(10),
        seller_id BIGINT,
        Name VARCHAR(MAX),
        Description VARCHAR(255),
        Nickname VARCHAR(100),
        error VARCHAR(MAX)
    );
END
GO