-- SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
-- SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
-- SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
SET CHARSET UTF8;
CREATE SCHEMA IF NOT EXISTS `resns_app` DEFAULT CHARACTER SET utf8;
USE `resns_app`;

CREATE TABLE IF NOT EXISTS `resns_app`.`users` (
`user_id` VARCHAR(16) NOT NULL COMMENT 'ユーザID',
`token`VARCHAR(64) NOT NULL COMMENT 'トークン',
`user_name` VARCHAR(16) NOT NULL COMMENT 'ユーザ名',
`yaer` int NOT NULL COMMENT '生年月日(年)',
`month` int NOT NULL COMMENT '生年月日年代(月)',
`gender` int NOT NULL COMMENT '性別',
`genre1` int NOT NULL COMMENT 'ジャンル1',
`genre2` int NOT NULL COMMENT 'ジャンル2',
`genre3` int NOT NULL COMMENT 'ジャンル3',
`genre4` int NOT NULL COMMENT 'ジャンル4',
PRIMARY KEY (`user_id`),
INDEX `idx_auth_token` (`user_id` ASC)
)
ENGINE = InnoDB
COMMENT = 'ユーザプロフィール';

CREATE TABLE IF NOT EXISTS `resns_app`.`users_info` (
`token`VARCHAR(64) NOT NULL COMMENT 'トークン',
`user_mail_adress` VARCHAR(16) NOT NULL COMMENT 'メールアドレス',
`user_password` VARCHAR(16) NOT NULL COMMENT 'パスワード',
PRIMARY KEY (`token`)
)
ENGINE = InnoDB
COMMENT = 'ユーザ情報';

CREATE TABLE IF NOT EXISTS `resns_app`.`users_list` (
`user_id` VARCHAR(16) NOT NULL COMMENT 'ユーザID',
`user_mail_adress` VARCHAR(16) NOT NULL COMMENT 'メールアドレス',
PRIMARY KEY (`user_id`)
)
ENGINE = InnoDB
COMMENT = 'リスト';

CREATE TABLE IF NOT EXISTS `resns_app`.`articles_contents` (
`article_id` VARCHAR(16) NOT NULL COMMENT '記事識別ID',
`image_path` VARCHAR(256) NOT NULL COMMENT '画像のパス',
`title` VARCHAR(32) NOT NULL COMMENT '記事のタイトル',
`context` VARCHAR(256) NOT NULL COMMENT '記事の内容',
`genre` int NOT NULL COMMENT '記事のジャンル',
`nice` int NOT NULL COMMENT 'いいね数',
`ageYaer` int NOT NULL COMMENT '年代(年)',
`ageMonth` int NOT NULL COMMENT '年代(月)',
PRIMARY KEY (`article_id`)
)
ENGINE = InnoDB
COMMENT = '記事の内容';

CREATE TABLE IF NOT EXISTS `resns_app`.`articles_tags` (
`article_id` VARCHAR(16) NOT NULL COMMENT '記事識別ID',
`article_tag1` VARCHAR(32) NOT NULL COMMENT '記事のタグ1',
`article_tag2` VARCHAR(32) NOT NULL COMMENT '記事のタグ2',
`article_tag3` VARCHAR(32) NOT NULL COMMENT '記事のタグ3',
`article_tag4` VARCHAR(32) NOT NULL COMMENT '記事のタグ4',
PRIMARY KEY (`article_id`)
)
ENGINE = InnoDB
COMMENT = '記事のタグ';

CREATE TABLE IF NOT EXISTS `resns_app`.`articles_comments` (
`article_id` VARCHAR(16) NOT NULL COMMENT '記事識別ID',
`comments_id` VARCHAR(16) NOT NULL COMMENT 'コメントのID',
`comments_contents` VARCHAR(32) NOT NULL COMMENT 'コメントの内容',
`user_id` VARCHAR(16) NOT NULL COMMENT 'ユーザID',
`user_image` VARCHAR(32) NOT NULL COMMENT 'ユーザーの画像',
PRIMARY KEY (`article_id`)
)
ENGINE = InnoDB
COMMENT = '記事へのコメント';

CREATE TABLE IF NOT EXISTS `resns_app`.`articles_nice_status` (
`article_id` VARCHAR(16) NOT NULL COMMENT '記事識別ID',
`user_id` VARCHAR(16) NOT NULL COMMENT 'ユーザID',
PRIMARY KEY (`article_id`)
)
ENGINE = InnoDB
COMMENT = '記事にいいねした人';

---- insert ----
INSERT INTO `articles_contents` VALUES ('1', 'https://initial-practice.s3-ap-northeast-1.amazonaws.com/IMG_7004.jpeg', 'Title_sample1','五稜郭で飲んだ帰りsample1sample1sample1sample1sample1sample1sample1sample1',3,46,2020,6);
INSERT INTO `articles_contents` VALUES ('2', 'https://initial-practice.s3-ap-northeast-1.amazonaws.com/sample.jpg','Title_sample2', 'sample2sample2sample2sample2sample2sample2sample2sample2sample2sample2sample2sample2sample2',4,100,2000,5 );
INSERT INTO `articles_contents` VALUES ('3', 'https://initial-practice.s3-ap-northeast-1.amazonaws.com/sample.jpg','Title_sample3', 'sample3sample3sample3sample3sample3sample3sample3sample3sample3sample3sample3sample3sample3',1,2,1999,10 );
INSERT INTO `articles_contents` VALUES ('4', 'https://initial-practice.s3-ap-northeast-1.amazonaws.com/sample.jpg', 'Title_sample4','sample4sample4sample4sample4sample4sample4sample4sample4sample4sample4sample4sample4sample4',2,0,1940,6);
INSERT INTO `articles_contents` VALUES ('5', 'https://initial-practice.s3-ap-northeast-1.amazonaws.com/sample.jpg', 'Title_sample5','sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5',5,2006,2 );

INSERT INTO `users_info` VALUES ('aaaaaaa', 'b1018085@fun.ac.jp', 'Title_sample5','sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5',5,2006,2 );

