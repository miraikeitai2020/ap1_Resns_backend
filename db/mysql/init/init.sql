-- SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
-- SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
-- SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
SET CHARSET UTF8;
CREATE SCHEMA IF NOT EXISTS `resns_app` DEFAULT CHARACTER SET utf8;
USE `resns_app`;


CREATE TABLE IF NOT EXISTS `resns_app`.`articles` (
`article_id` VARCHAR(16) NOT NULL COMMENT '記事識別ID',
`image_path` VARCHAR(256) NOT NULL COMMENT '画像のパス',
`title` VARCHAR(32) NOT NULL COMMENT '記事のタイトル',
`context` VARCHAR(256) NOT NULL COMMENT '記事の内容',
`genre` int NOT NULL COMMENT '記事のジャンル',
`nice` int NOT NULL COMMENT 'いいね数',
`ageYaer` int NOT NULL COMMENT '年代(年)',
`ageMonth` int NOT NULL COMMENT '年代(月)',
-- `tag` VARCHAR(256) NOT NULL COMMENT 'タグ',
PRIMARY KEY (`article_id`),
INDEX `idx_auth_token` (`article_id` ASC))
DEFAULT CHARSET=utf8 COLLATE=utf8_bin
ENGINE = InnoDB
COMMENT = '記事';


---- insert ----
INSERT INTO `articles` VALUES ('1', 'https://initial-practice.s3-ap-northeast-1.amazonaws.com/IMG_7004.jpeg', 'Title_sample1','五稜郭で飲んだ帰りsample1sample1sample1sample1sample1sample1sample1sample1',3,46,2020,6);
INSERT INTO `articles` VALUES ('2', 'https://initial-practice.s3-ap-northeast-1.amazonaws.com/sample.jpg','Title_sample2', 'sample2sample2sample2sample2sample2sample2sample2sample2sample2sample2sample2sample2sample2',4,100,2000,5 );
INSERT INTO `articles` VALUES ('3', 'https://initial-practice.s3-ap-northeast-1.amazonaws.com/sample.jpg','Title_sample3', 'sample3sample3sample3sample3sample3sample3sample3sample3sample3sample3sample3sample3sample3',1,2,1999,10 );
INSERT INTO `articles` VALUES ('4', 'https://initial-practice.s3-ap-northeast-1.amazonaws.com/sample.jpg', 'Title_sample4','sample4sample4sample4sample4sample4sample4sample4sample4sample4sample4sample4sample4sample4',2,0,1940,6);
INSERT INTO `articles` VALUES ('5', 'https://initial-practice.s3-ap-northeast-1.amazonaws.com/sample.jpg', 'Title_sample5','sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5',5,2006,2 );

