-- SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
-- SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
-- SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
SET CHARSET UTF8;
CREATE SCHEMA IF NOT EXISTS `resns_app` DEFAULT CHARACTER SET utf8;
USE `resns_app`;

CREATE TABLE IF NOT EXISTS `resns_app`.`users` (
`token`VARCHAR(64) NOT NULL COMMENT 'トークン',
`user_id` VARCHAR(16) NOT NULL COMMENT 'ユーザID',
`user_name` VARCHAR(16) NOT NULL COMMENT 'ユーザ名',
`user_image` VARCHAR(128) NOT NULL COMMENT 'プロフィール画像',
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
`user_mail_adress` VARCHAR(32) NOT NULL COMMENT 'メールアドレス',
`user_password` VARCHAR(64) NOT NULL COMMENT 'パスワード',
PRIMARY KEY (`token`)
)
ENGINE = InnoDB
COMMENT = 'ユーザ情報';

CREATE TABLE IF NOT EXISTS `resns_app`.`users_list` (
`user_id` VARCHAR(16) NOT NULL COMMENT 'ユーザID',
`article_id` VARCHAR(64) NOT NULL COMMENT '記事識別ID',
PRIMARY KEY (`user_id`)
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
`ageYaer` int NOT NULL COMMENT '年代(年)',
`ageMonth` int NOT NULL COMMENT '年代(月)',
PRIMARY KEY (`article_id`)
)
ENGINE = InnoDB
COMMENT = '記事の内容';

CREATE TABLE IF NOT EXISTS `resns_app`.`articles_tags` (
`article_id` VARCHAR(64) NOT NULL COMMENT '記事識別ID',
`article_tag1` VARCHAR(32) NOT NULL COMMENT '記事のタグ1',
`article_tag2` VARCHAR(32) NOT NULL COMMENT '記事のタグ2',
`article_tag3` VARCHAR(32) NOT NULL COMMENT '記事のタグ3',
`article_tag4` VARCHAR(32) NOT NULL COMMENT '記事のタグ4',
PRIMARY KEY (`article_id`)
)
ENGINE = InnoDB
COMMENT = '記事のタグ';

CREATE TABLE IF NOT EXISTS `resns_app`.`articles_comments` (
`article_id` VARCHAR(64) NOT NULL COMMENT '記事識別ID',
`comments_id` VARCHAR(64) NOT NULL COMMENT 'コメントのID',
`comments_contents` VARCHAR(64) NOT NULL COMMENT 'コメントの内容',
`user` VARCHAR(16) NOT NULL COMMENT 'ユーザID',
`user_image` VARCHAR(128) NOT NULL COMMENT 'ユーザーの画像',
PRIMARY KEY (`article_id`)
)
ENGINE = InnoDB
COMMENT = '記事へのコメント';

CREATE TABLE IF NOT EXISTS `resns_app`.`articles_nice_status` (
`article_id` VARCHAR(64) NOT NULL COMMENT '記事識別ID',
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
INSERT INTO `articles_contents` VALUES ('5', 'https://initial-practice.s3-ap-northeast-1.amazonaws.com/sample.jpg', 'Title_sample5','sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5',5,123,2006,2 );

INSERT INTO `users_info` VALUES ('1c22d6ef-b83a-47d9-a63d-b5d78cac4a87', 'b1018085@fun.ac.jp', 'ed968e840d10d2d313a870bc131a4e2c311d7ad09bdf32b3418147221f51a6e2');

INSERT INTO `users` VALUES ('1c22d6ef-b83a-47d9-a63d-b5d78cac4a87','taketo1234','taketo','https://initial-practice.s3-ap-northeast-1.amazonaws.com/sample.jp',1999,8,0,1,2,3,4);

INSERT INTO `articles_tags` VALUES ('1', '1_sampleTag1', '1_sampleTag2', '1_sampleTag3', '1_sampleTag4');
INSERT INTO `articles_tags` VALUES ('2', '2_sampleTag1', '2_sampleTag2', '2_sampleTag3', '2_sampleTag4');
INSERT INTO `articles_tags` VALUES ('3', '3_sampleTag1', '3_sampleTag2', '3_sampleTag3', '3_sampleTag4');
INSERT INTO `articles_tags` VALUES ('4', '4_sampleTag1', '4_sampleTag2', '4_sampleTag3', '4_sampleTag4');
INSERT INTO `articles_tags` VALUES ('5', '5_sampleTag1', '5_sampleTag2', '5_sampleTag3', '5_sampleTag4');

INSERT INTO `articles_comments` VALUES ('1','2v43d6ef-a83d-57d0-f33d-b5d78ncj4a58','面白い','taketo1234','https://initial-practice.s3-ap-northeast-1.amazonaws.com/sample.jp');

INSERT INTO `users_list` VALUES ('1','taketo1234');
INSERT INTO `users_list` VALUES ('5','taketo1234');

