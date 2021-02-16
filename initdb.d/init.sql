DROP DATABASE IF EXISTS game;
CREATE DATABASE game;
USE game;
DROP TABLE IF EXISTS Users;
 
CREATE TABLE Users (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
name TEXT NOT NULL,
token TEXT NOT NULL
)DEFAULT CHARACTER SET=utf8;


DROP TABLE IF EXISTS Characters;
  
CREATE TABLE Characters (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
name TEXT NOT NULL
)DEFAULT CHARACTER SET=utf8;

INSERT INTO Characters (name) VALUES ("Totoro");
INSERT INTO Characters (name) VALUES ("Chihiro");
INSERT INTO Characters (name) VALUES ("Kiki");
INSERT INTO Characters (name) VALUES ("Marnie");


DROP TABLE IF EXISTS UserCharacters; 

CREATE TABLE UserCharacters (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
userId INT NOT NULL,
characterId INT NOT NULL,
ts TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)DEFAULT CHARACTER SET=utf8;
 
INSERT INTO UserCharacters (userId,characterId) VALUES (1,1);
INSERT INTO UserCharacters (userId,characterId) VALUES (1,2);
INSERT INTO UserCharacters (userId,characterId) VALUES (1,3);

DROP TABLE IF EXISTS Gachas;
  
CREATE TABLE Gachas (
characterId INT NOT NULL PRIMARY KEY,
number INT NOT NULL
)DEFAULT CHARACTER SET=utf8;

INSERT INTO Gachas (characterId,number) VALUES (1,10);
INSERT INTO Gachas (characterId,number) VALUES (2,10);
INSERT INTO Gachas (characterId,number) VALUES (3,10);
INSERT INTO Gachas (characterId,number) VALUES (4,10);
 

