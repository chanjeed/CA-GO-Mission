DROP DATABASE IF EXISTS game;
CREATE DATABASE game;
USE game;
DROP TABLE IF EXISTS Users;
 
CREATE TABLE Users (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
name TEXT NOT NULL,
token TEXT NOT NULL
)DEFAULT CHARACTER SET=utf8;
 