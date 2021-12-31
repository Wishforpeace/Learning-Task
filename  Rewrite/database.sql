CREATE DATABASE IF NOT EXISTS football;

USE football;

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS games;
DROP TABLE IF EXISTS players;
DROP TABLE IF EXISTS teams;

CREATE TABLE users(
    `user_id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_name` VARCHAR(100)  NULL,
    `user_password` VARCHAR(100)  NULL,
    `identity` INT(10)  NULL,
    PRIMARY KEY(`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE games (
    `game_id` BIGINT NOT NULL AUTO_INCREMENT,
    `game_name` VARCHAR(100) NOT NULL ,
    `game_date` VARCHAR(100) NOT NULL,
    `place` VARCHAR(100) NOT NULL,
    `Info` TEXT NULL,
    `appointment` BIGINT NOT NULL,
    `TEAMA` VARCHAR(100) NOT NULL,
    `TEAMB` VARCHAR(100) NOT NULL,
    PRIMARY KEY (`game_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE players (
    `player_ID` BIGINT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL,
    `team_name` VARCHAR(100) NULL,
    `img` VARCHAR(100) NULL,
    `Info` TEXT NULL,
    PRIMARY KEY(`player_ID`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE teams (
    `team_ID` BIGINT NOT NULL AUTO_INCREMENT,
    `T_name` VARCHAR(100) NOT NULL,
    `logo` VARCHAR(100) NULL,
    `Info` TEXT NULL,
    PRIMARY KEY(`team_ID`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE reservation(
    `user_name` VARCHAR(100) NOT NULL,
    `game` VARCHAR(100) NOT NUll
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE authentication_tokens (
           token_id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
           user_id BIGINT,
           auth_token VARCHAR(255),
           generated_at DATETIME,
           expires_at   DATETIME
) ENGINE = InnoDB;

