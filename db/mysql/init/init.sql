-- SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
-- SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
-- SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
SET CHARSET UTF8;
DROP SCHEMA IF EXISTS `resns_app`;
CREATE SCHEMA IF NOT EXISTS `resns_app` DEFAULT CHARACTER SET utf8;
USE `resns_app`;

---- drop ----
DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `users_genres`;
DROP TABLE IF EXISTS `users_info`;
DROP TABLE IF EXISTS `users_list`;
DROP TABLE IF EXISTS `articles_contents`;
DROP TABLE IF EXISTS `articles_tag`;
DROP TABLE IF EXISTS `articles_comments`;
DROP TABLE IF EXISTS `articles_nice_status`;

---- create ----
CREATE TABLE IF NOT EXISTS `resns_app`.`users` (
`id` VARCHAR(32) NOT NULL COMMENT 'ユーザID',
`name` VARCHAR(32) NOT NULL COMMENT 'ユーザ名',
`image` VARCHAR(128) NOT NULL COMMENT 'プロフィール画像',
`year` int NOT NULL COMMENT '生年月日(年)',
`month` int NOT NULL COMMENT '生年月日(月)',
`gender` int NOT NULL COMMENT '性別',
PRIMARY KEY (`id`),
INDEX `idx_auth_token` (`id` ASC)
)
ENGINE = InnoDB
COMMENT = 'ユーザプロフィール';

CREATE TABLE IF NOT EXISTS `resns_app`.`users_genres` (
`id` VARCHAR(64) NOT NULL COMMENT 'ユーザID',
`genre_1` int NOT NULL COMMENT '1つ目のジャンル',
`genre_2` int NOT NULL COMMENT '2つ目のジャンル',
`genre_3` int NOT NULL COMMENT '3つ目のジャンル',
`genre_4` int NOT NULL COMMENT '4つ目のジャンル',
PRIMARY KEY (`id`)
)
ENGINE = InnoDB
COMMENT = 'ユーザ選択ジャンル';

CREATE TABLE IF NOT EXISTS `resns_app`.`users_info` (
`id` VARCHAR(64) NOT NULL COMMENT 'ユーザID',
`password` VARCHAR(64) NOT NULL COMMENT 'パスワード',
PRIMARY KEY (`id`)
)
ENGINE = InnoDB
COMMENT = 'ユーザ情報';

CREATE TABLE IF NOT EXISTS `resns_app`.`users_list` (
`list_id` VARCHAR(64) NOT NULL COMMENT '識別ID',
`user_id` VARCHAR(64) NOT NULL COMMENT 'ユーザID',
`article_id` VARCHAR(64) NOT NULL COMMENT '記事識別ID',
PRIMARY KEY (`list_id`),
FOREIGN KEY (`user_id`)
    REFERENCES `resns_app`.`users` (`id`)
)
ENGINE = InnoDB
COMMENT = 'リスト';

CREATE TABLE IF NOT EXISTS `resns_app`.`articles_contents` (
`article_id` VARCHAR(64) NOT NULL COMMENT '記事識別ID',
`image_path` VARCHAR(128) NOT NULL COMMENT '画像のパス',
`title` VARCHAR(32) NOT NULL COMMENT '記事のタイトル',
`context` VARCHAR(256) NOT NULL COMMENT '記事の内容',
`genre` int NOT NULL COMMENT '記事のジャンル',
`nice` int NOT NULL COMMENT 'いいね数',
`age_year` int NOT NULL COMMENT '年代(年)',
`age_month` int NOT NULL COMMENT '年代(月)',
PRIMARY KEY (`article_id`)
)
ENGINE = InnoDB
COMMENT = '記事の内容';

CREATE TABLE IF NOT EXISTS `resns_app`.`articles_tag` (
`tag_id` VARCHAR(64) NOT NULL COMMENT 'タグID',
`article_id` VARCHAR(64) NOT NULL COMMENT '記事識別ID',
`article_tag` VARCHAR(32) NOT NULL COMMENT '記事のタグ',
PRIMARY KEY (`tag_id`)
)
ENGINE = InnoDB
COMMENT = '記事のタグ';

CREATE TABLE IF NOT EXISTS `resns_app`.`articles_comments` (
`article_id` VARCHAR(64) NOT NULL COMMENT '記事識別ID',
`comments_id` VARCHAR(64) NOT NULL COMMENT 'コメントのID',
`comments_contents` VARCHAR(64) NOT NULL COMMENT 'コメントの内容',
`user_id` VARCHAR(64) NOT NULL COMMENT 'ユーザID',
`user_image` VARCHAR(128) NOT NULL COMMENT 'ユーザーの画像',
PRIMARY KEY (`comments_id`),
FOREIGN KEY (`user_id`)
    REFERENCES `resns_app`.`users` (`id`)
)
ENGINE = InnoDB
COMMENT = '記事へのコメント';

CREATE TABLE IF NOT EXISTS `resns_app`.`articles_nice_status` (
`nice_id` VARCHAR(64) NOT NULL COMMENT 'NiceID',
`article_id` VARCHAR(64) NOT NULL COMMENT '記事識別ID',
`user_id` VARCHAR(16) NOT NULL COMMENT 'ユーザID',
PRIMARY KEY (`nice_id`),
FOREIGN KEY (`user_id`)
    REFERENCES `resns_app`.`users` (`id`)
)
ENGINE = InnoDB
COMMENT = '記事にいいねした人';

---- insert ----
INSERT INTO `articles_contents` VALUES ('1', 'https://initial-practice.s3-ap-northeast-1.amazonaws.com/IMG_7004.jpeg', 'Title_sample1','五稜郭で飲んだ帰りsample1sample1sample1sample1sample1sample1sample1sample1',3,46,2020,6);
INSERT INTO `articles_contents` VALUES ('2', 'https://initial-practice.s3-ap-northeast-1.amazonaws.com/sample.jpg','Title_sample2', 'sample2sample2sample2sample2sample2sample2sample2sample2sample2sample2sample2sample2sample2',4,100,2000,5 );
INSERT INTO `articles_contents` VALUES ('3', 'https://initial-practice.s3-ap-northeast-1.amazonaws.com/sample.jpg','Title_sample3', 'sample3sample3sample3sample3sample3sample3sample3sample3sample3sample3sample3sample3sample3',1,2,1999,10 );
INSERT INTO `articles_contents` VALUES ('4', 'https://initial-practice.s3-ap-northeast-1.amazonaws.com/sample.jpg', 'Title_sample4','sample4sample4sample4sample4sample4sample4sample4sample4sample4sample4sample4sample4sample4',2,0,1940,6);
INSERT INTO `articles_contents` VALUES ('5', 'https://initial-practice.s3-ap-northeast-1.amazonaws.com/sample.jpg', 'Title_sample5','sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5',5,123,2006,2 );

INSERT INTO `users_info` VALUES ('b1018085@fun.ac.jp', 'ed968e840d10d2d313a870bc131a4e2c311d7ad09bdf32b3418147221f51a6e2');

INSERT INTO `users` VALUES ('b1018085@fun.ac.jp','taketo','https://initial-practice.s3-ap-northeast-1.amazonaws.com/sample.jp',1999,8,1);

INSERT INTO `articles_tag` VALUES ('1','1', '1_sampleTag1');
INSERT INTO `articles_tag` VALUES ('2','1', '1_sampleTag2');
INSERT INTO `articles_tag` VALUES ('3','1', '1_sampleTag3');
INSERT INTO `articles_tag` VALUES ('4','1', '1_sampleTag4');
INSERT INTO `articles_tag` VALUES ('5','2', '2_sampleTag1');
INSERT INTO `articles_tag` VALUES ('6','2', '2_sampleTag2');
INSERT INTO `articles_tag` VALUES ('7','2', '2_sampleTag3');
INSERT INTO `articles_tag` VALUES ('8','2', '2_sampleTag4');
INSERT INTO `articles_tag` VALUES ('9','3', '3_sampleTag1');
INSERT INTO `articles_tag` VALUES ('10','3', '3_sampleTag2');
INSERT INTO `articles_tag` VALUES ('11','3', '3_sampleTag3');
INSERT INTO `articles_tag` VALUES ('12','3', '3_sampleTag4');
INSERT INTO `articles_tag` VALUES ('13','4', '4_sampleTag1');
INSERT INTO `articles_tag` VALUES ('14','4', '4_sampleTag2');
INSERT INTO `articles_tag` VALUES ('15','4', '4_sampleTag3');
INSERT INTO `articles_tag` VALUES ('16','4', '4_sampleTag4');
INSERT INTO `articles_tag` VALUES ('17','5', '5_sampleTag1');
INSERT INTO `articles_tag` VALUES ('18','5', '5_sampleTag2');
INSERT INTO `articles_tag` VALUES ('19','5', '5_sampleTag3');
INSERT INTO `articles_tag` VALUES ('20','5', '5_sampleTag4');

INSERT INTO `articles_comments` VALUES ('1','2v43d6ef-a83d-57d0-f33d-b5d78ncj4a58','面白い','b1018085@fun.ac.jp','https://initial-practice.s3-ap-northeast-1.amazonaws.com/sample.jp');

INSERT INTO `users_list` VALUES ('1','b1018085@fun.ac.jp','1');
INSERT INTO `users_list` VALUES ('2','b1018085@fun.ac.jp','4');