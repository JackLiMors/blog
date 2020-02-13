/*
SQLyog Community v12.4.0 (64 bit)
MySQL - 5.7.26 : Database - blog
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`blog` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;

USE `blog`;

/*Table structure for table `article` */

DROP TABLE IF EXISTS `article`;

CREATE TABLE `article` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '文章ID',
  `category_id` bigint(20) NOT NULL COMMENT '分类ID',
  `content` longtext COLLATE utf8_unicode_ci NOT NULL COMMENT '文章内容',
  `title` varchar(1024) COLLATE utf8_unicode_ci NOT NULL COMMENT '文章标题',
  `view_count` int(255) NOT NULL COMMENT '阅读次数',
  `comment_count` int(255) NOT NULL COMMENT '评论次数',
  `username` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '作者',
  `status` int(10) NOT NULL DEFAULT '1' COMMENT '状态，正常为1',
  `summary` varchar(256) COLLATE utf8_unicode_ci NOT NULL COMMENT '文章摘要',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '发布时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

/*Data for the table `article` */

insert  into `article`(`id`,`category_id`,`content`,`title`,`view_count`,`comment_count`,`username`,`status`,`summary`,`create_time`,`update_time`) values 
(1,1,'abcdefg','test title',0,0,'jacklimors',1,'abcd','2020-02-08 22:21:15','2020-02-08 22:21:15'),
(2,1,'abcdefg','test title2',0,0,'jacklimors',1,'abcd','2020-02-08 22:22:57','2020-02-08 22:22:57'),
(3,1,'abcdefg','test title3',0,0,'jacklimors',1,'abcd','2020-02-08 22:30:28','2020-02-08 22:30:28');

/*Table structure for table `category` */

DROP TABLE IF EXISTS `category`;

CREATE TABLE `category` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `category_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '分类名字',
  `category_no` int(10) NOT NULL COMMENT '分类排序',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

/*Data for the table `category` */

insert  into `category`(`id`,`category_name`,`category_no`,`create_time`,`update_time`) values 
(1,'test1',1,'2020-02-08 20:16:15','2020-02-08 20:16:15'),
(2,'test2',2,'2020-02-08 21:03:19','2020-02-08 21:03:19'),
(3,'test3',3,'2020-02-08 21:03:27','2020-02-08 21:03:27');

/*Table structure for table `comment` */

DROP TABLE IF EXISTS `comment`;

CREATE TABLE `comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `content` text COLLATE utf8_unicode_ci NOT NULL COMMENT '评论内容',
  `username` varchar(64) COLLATE utf8_unicode_ci NOT NULL COMMENT '评论作者',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论发布时间',
  `status` int(255) NOT NULL DEFAULT '1' COMMENT '评论状态:0 删除：1，。。。',
  `article_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

/*Data for the table `comment` */

/*Table structure for table `leave` */

DROP TABLE IF EXISTS `leave`;

CREATE TABLE `leave` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `cotent` text COLLATE utf8_unicode_ci NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

/*Data for the table `leave` */

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
