-- MySQL dump 10.13  Distrib 8.0.17, for Win64 (x86_64)
--
-- Host: localhost    Database: question
-- ------------------------------------------------------
-- Server version	8.0.17

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `answer`
--

DROP TABLE IF EXISTS `answer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `answer`
(
  `id` bigint
(20) NOT NULL AUTO_INCREMENT,
  `answer_id` bigint
(20) unsigned NOT NULL,
  `content` text CHARACTER
SET utf8mb4
COLLATE utf8mb4_general_ci NOT NULL,
  `comment_count` int
(10) unsigned NOT NULL DEFAULT '0',
  `voteup_count` int
(11) NOT NULL DEFAULT '0',
  `author_id` bigint
(20) NOT NULL,
  `status` tinyint
(3) unsigned NOT NULL DEFAULT '1',
  `can_comment` tinyint
(3) unsigned NOT NULL DEFAULT '1',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON
UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY
(`id`),
  UNIQUE KEY `idx_answer_id`
(`answer_id`),
  KEY `idx_author_Id`
(`author_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `answer`
--

LOCK TABLES `answer` WRITE;
/*!40000 ALTER TABLE `answer` DISABLE KEYS */;
INSERT INTO `
answer`
VALUES
  (2, 289147270031474688, '你好，我是zz', 0, 0, 287398135439818752, 1, 1, '2020-02-16 17:40:03', '2020-02-16 17:40:03'),
  (3, 289147914209460224, '你好，hello ', 0, 0, 287398135439818752, 1, 1, '2020-02-16 17:46:27', '2020-02-16 17:46:27'),
  (4, 289148456801402880, 'shazi', 0, 0, 287398135439818752, 1, 1, '2020-02-16 17:51:50', '2020-02-16 17:51:50'),
  (5, 289231738079543296, '老李老李', 0, 0, 287398135439818752, 1, 1, '2020-02-17 07:39:10', '2020-02-17 07:39:10'),
  (6, 289232703977422848, '真的炸', 0, 0, 287398135439818752, 1, 1, '2020-02-17 07:48:45', '2020-02-17 07:48:45'),
  (7, 289234203105558528, '说的是个锤子', 0, 0, 287398135439818752, 1, 1, '2020-02-17 08:03:39', '2020-02-17 08:03:39');
/*!40000 ALTER TABLE `answer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `category`
(
  `id` int
(11) NOT NULL AUTO_INCREMENT,
  `category_id` int
(10) unsigned NOT NULL,
  `category_name` varchar
(128) CHARACTER
SET utf8mb4
COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON
UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY
(`id`),
  UNIQUE KEY `idx_category_id`
(`category_id`),
  UNIQUE KEY `idx_category_name`
(`category_name`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `
category`
VALUES
  (1, 1, '技术', '2019-01-01 00:30:40', '2019-01-01 00:30:40'),
  (2, 2, '情感', '2019-01-01 00:31:07', '2019-01-01 00:31:07'),
  (3, 3, '王者荣耀', '2019-01-01 00:31:25', '2019-01-01 00:31:25'),
  (4, 4, '吃鸡', '2019-01-01 07:45:13', '2019-01-01 07:45:13'),
  (5, 5, '科幻', '2019-01-05 15:02:43', '2019-01-05 15:02:43');
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment`
--

DROP TABLE IF EXISTS `comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comment`
(
  `id` bigint
(20) NOT NULL AUTO_INCREMENT,
  `comment_id` bigint
(20) unsigned NOT NULL,
  `content` text CHARACTER
SET utf8mb4
COLLATE utf8mb4_general_ci NOT NULL,
  `comment_count` int
(10) unsigned NOT NULL DEFAULT '0',
  `like_count` int
(10) unsigned NOT NULL DEFAULT '0',
  `author_id` bigint
(20) NOT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON
UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY
(`id`),
  UNIQUE KEY `idx_comment_id`
(`comment_id`),
  KEY `idx_author_Id`
(`author_id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
INSERT INTO `
comment`
VALUES
  (6, 288415481323323392, 'hello world!', 0, 0, 100, '2020-02-11 16:30:23', '2020-02-11 16:30:23'),
  (13, 289235315132989440, 'asdddddddddddddddddddddddd', 0, 0, 287398135439818752, '2020-02-17 08:14:42', '2020-02-17 08:14:42');
/*!40000 ALTER TABLE `comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment_rel`
--

DROP TABLE IF EXISTS `comment_rel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comment_rel`
(
  `id` bigint
(20) NOT NULL AUTO_INCREMENT,
  `comment_id` bigint
(20) unsigned NOT NULL,
  `parent_id` bigint
(20) unsigned NOT NULL,
  `level` int
(10) unsigned NOT NULL,
  `question_id` bigint
(20) unsigned NOT NULL,
  `reply_comment_id` bigint
(20) unsigned NOT NULL DEFAULT '0',
  `reply_author_id` bigint
(20) unsigned NOT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON
UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY
(`id`),
  KEY `idx_level`
(`level`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment_rel`
--

LOCK TABLES `comment_rel` WRITE;
/*!40000 ALTER TABLE `comment_rel` DISABLE KEYS */;
INSERT INTO `
comment_rel`
VALUES
  (1, 288415481323323392, 0, 1, 100, 0, 10001, '2020-02-11 16:30:23', '2020-02-11 16:30:23'),
  (3, 289235315132989440, 0, 1, 287544070828457984, 0, 0, '2020-02-17 08:14:42', '2020-02-17 08:14:42');
/*!40000 ALTER TABLE `comment_rel` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `question`
--

DROP TABLE IF EXISTS `question`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `question`
(
  `id` bigint
(20) NOT NULL AUTO_INCREMENT,
  `question_id` bigint
(20) NOT NULL COMMENT '问题id',
  `caption` varchar
(128) CHARACTER
SET utf8mb4
COLLATE utf8mb4_general_ci NOT NULL COMMENT '问题标题',
  `content` varchar
(8192) CHARACTER
SET utf8mb4
COLLATE utf8mb4_general_ci NOT NULL COMMENT '问题内容',
  `author_id` bigint
(20) NOT NULL COMMENT '作者的用户id',
  `category_id` bigint
(20) NOT NULL COMMENT '所属栏目',
  `status` tinyint
(4) NOT NULL DEFAULT '1' COMMENT '问题状态',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON
UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY
(`id`),
  KEY `idx_author_id`
(`author_id`),
  KEY `idx_question_id`
(`question_id`),
  KEY `idx_category_id`
(`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `question`
--

LOCK TABLES `question` WRITE;
/*!40000 ALTER TABLE `question` DISABLE KEYS */;
INSERT INTO `
question`
VALUES
  (1, 229544174527971329, '上岛咖啡斯拉夫斯卡就开始了是开始实施上课试试看是', '开始了打发时间浪费时间浪费多少反倒是浪费时间开放时间发生的雷锋精神离开房间是否是否', 224533709452214273, 1, 1, '2019-01-01 13:16:29', '2019-01-01 13:16:29'),
  (2, 230134441849126913, '未来三十年内，哪些行业的工作人员可能会被人工智能取代？', '人工智能这些年的快速发展，在某些领域已经开始渐渐取代人类的工作岗位了。未来这种情况是否会越来越严重？以后的人类会进入空虚的享乐时代么？', 224533709452214273, 1, 1, '2019-01-05 15:00:16', '2019-01-05 15:00:16'),
  (3, 230134511709454337, '你见过最渣的渣女有多渣？', '我一个玩的挺好的舍友，在一家旁边很多酒吧的电玩电玩城上班，酒吧多，帅哥，男生也就多了，喜欢她的男生就也还挺多的，她说她自己是个渣女，因为同时跟四五个男生暧昧，出去玩都是跟这个男生玩完，又去跟另一个男生玩。她长得一般，但身材可以，你们觉得她渣吗？\n\n\n', 224533709452214273, 1, 1, '2019-01-05 15:00:58', '2019-01-05 15:00:58'),
  (4, 230134710704013313, '你觉得《三体》中最残忍的一句话是什么？', '你觉得《三体》中最残忍的一句话是什么？', 224533709452214273, 5, 1, '2019-01-05 15:02:56', '2019-01-05 15:02:56'),
  (6, 287544070828457984, 'hello', 'c和指针', 287398135439818752, 1, 1, '2020-02-05 16:13:41', '2020-02-05 16:13:41');
/*!40000 ALTER TABLE `question` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `question_answer_rel`
--

DROP TABLE IF EXISTS `question_answer_rel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `question_answer_rel`
(
  `id` bigint
(20) NOT NULL AUTO_INCREMENT,
  `question_id` bigint
(20) NOT NULL,
  `answer_id` bigint
(20) NOT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY
(`id`),
  UNIQUE KEY `idx_question_answer`
(`question_id`,`answer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `question_answer_rel`
--

LOCK TABLES `question_answer_rel` WRITE;
/*!40000 ALTER TABLE `question_answer_rel` DISABLE KEYS */;
INSERT INTO `
question_answer_rel`
VALUES
  (2, 229544174527971329, 289147270031474688, '2020-02-16 17:40:03'),
  (3, 229544174527971329, 289147914209460224, '2020-02-16 17:46:27'),
  (4, 229544174527971329, 289148456801402880, '2020-02-16 17:51:50'),
  (5, 230134441849126913, 289231738079543296, '2020-02-17 07:39:10'),
  (6, 287544070828457984, 289232703977422848, '2020-02-17 07:48:45'),
  (7, 287544070828457984, 289234203105558528, '2020-02-17 08:03:39');
/*!40000 ALTER TABLE `question_answer_rel` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user`
(
  `id` bigint
(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint
(20) NOT NULL,
  `username` varchar
(64) CHARACTER
SET utf8mb4
COLLATE utf8mb4_general_ci NOT NULL,
  `nickname` varchar
(64) CHARACTER
SET utf8mb4
COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `password` varchar
(64) CHARACTER
SET utf8mb4
COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar
(64) CHARACTER
SET utf8mb4
COLLATE utf8mb4_general_ci NOT NULL,
  `sex` tinyint
(4) NOT NULL DEFAULT '0',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON
UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY
(`id`),
  UNIQUE KEY `idx_username`
(`username`) USING BTREE,
  UNIQUE KEY `idx_user_id`
(`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user`
VALUES
  (3, 223181233645944833, 'admin2', '', '21968b082e3af16563cad001184ddfa8', 'admin', 0, '2018-11-18 15:46:21', '2018-11-18 15:46:21'),
  (4, 224533709452214273, 'admin', '', '21968b082e3af16563cad001184ddfa8', 'admin', 0, '2018-11-27 23:41:59', '2018-11-27 23:41:59'),
  (5, 224533754784251905, 'admin22', '', 'c8814aea2aa8a4feae0569b18a97cd00', 'admin22', 0, '2018-11-27 23:42:26', '2018-11-27 23:42:26'),
  (6, 224533951312560129, 'admin222', '', 'c8814aea2aa8a4feae0569b18a97cd00', 'admin222', 0, '2018-11-27 23:44:24', '2018-11-27 23:44:24'),
  (7, 224534423356309505, 'admin2221', '', 'c8814aea2aa8a4feae0569b18a97cd00', 'admin2221', 0, '2018-11-27 23:49:05', '2018-11-27 23:49:05'),
  (8, 224534702126530561, 'admin22212', '', 'c8814aea2aa8a4feae0569b18a97cd00', 'admin22212', 0, '2018-11-27 23:51:51', '2018-11-27 23:51:51'),
  (9, 224534879344263169, 'admin222121', '', 'c8814aea2aa8a4feae0569b18a97cd00', 'admin222121', 0, '2018-11-27 23:53:37', '2018-11-27 23:53:37'),
  (10, 224535034466402305, 'admin2221212', '', 'c8814aea2aa8a4feae0569b18a97cd00', 'admin2221212', 0, '2018-11-27 23:55:09', '2018-11-27 23:55:09'),
  (11, 224535766808657921, 'admin22212121', '', 'c8814aea2aa8a4feae0569b18a97cd00', 'admin22212121', 0, '2018-11-28 00:02:26', '2018-11-28 00:02:26'),
  (12, 224536026788397057, 'admin111', '', '21968b082e3af16563cad001184ddfa8', 'admin111', 0, '2018-11-28 00:05:01', '2018-11-28 00:05:01'),
  (13, 224536207143469057, 'admin1111', '', '21968b082e3af16563cad001184ddfa8', 'admin1111', 0, '2018-11-28 00:06:48', '2018-11-28 00:06:48'),
  (14, 224536320255459329, 'admin11112', '', '21968b082e3af16563cad001184ddfa8', 'admin11112', 0, '2018-11-28 00:07:56', '2018-11-28 00:07:56'),
  (15, 224536455865696257, 'admin111122', '', '21968b082e3af16563cad001184ddfa8', 'admin111122', 0, '2018-11-28 00:09:16', '2018-11-28 00:09:16'),
  (16, 224536664054169601, 'admin1111223', '', '21968b082e3af16563cad001184ddfa8', 'admin1111223', 0, '2018-11-28 00:11:21', '2018-11-28 00:11:21'),
  (17, 224536736212975617, 'admin11112234', '', '21968b082e3af16563cad001184ddfa8', 'admin11112234', 0, '2018-11-28 00:12:04', '2018-11-28 00:12:04'),
  (18, 225180933609750529, 'admi', '', '46f19871f7a3088f42ecf44102d2e6e9', 'admin', 0, '2018-12-02 10:51:35', '2018-12-02 10:51:35'),
  (19, 225181104049487873, 'admi2', '', '46f19871f7a3088f42ecf44102d2e6e9', 'admin2', 0, '2018-12-02 10:53:17', '2018-12-02 10:53:17'),
  (20, 225181127923466241, 'admi22', '', '46f19871f7a3088f42ecf44102d2e6e9', 'admin2', 0, '2018-12-02 10:53:31', '2018-12-02 10:53:31'),
  (21, 225182107310227457, 'admin2222', '', '21968b082e3af16563cad001184ddfa8', 'admin', 0, '2018-12-02 11:03:15', '2018-12-02 11:03:15'),
  (22, 225183383855038465, 'admin1111222', '', '6a74dfafca685f959141472ffa1d907b', 'admin', 0, '2018-12-02 11:15:56', '2018-12-02 11:15:56'),
  (23, 225183813502763009, 'nadmi', '', '8b6a48f398cf7b5b9ae62adf75791e88', 'admi', 0, '2018-12-02 11:20:12', '2018-12-02 11:20:12'),
  (24, 225184375774380033, 'admin434444', '', '21968b082e3af16563cad001184ddfa8', 'admin', 0, '2018-12-02 11:25:47', '2018-12-02 11:25:47'),
  (25, 225186212074225665, 'DKDK', '', '9c0ac25434be0cf1ceeafaa11c80a578', 'SMIN', 0, '2018-12-02 11:44:01', '2018-12-02 11:44:01'),
  (26, 225190952694710273, 'admin222222222', '', '21968b082e3af16563cad001184ddfa8', 'admin', 0, '2018-12-02 12:31:07', '2018-12-02 12:31:07'),
  (27, 225191497299918849, 'sherlockhua', '拈花湾', 'e46445d7555d22b48bea1997cd607420', 'sherlockhua@163.com', 1, '2018-12-02 12:36:32', '2018-12-02 12:36:32'),
  (28, 225191580145811457, 'kala', 'kalo', '76bac7d7ee3687f8c583769d410a2c0f', 'kala@qq.com', 2, '2018-12-02 12:37:21', '2018-12-02 12:37:21'),
  (29, 287398135439818752, 'yiluhuakai', 'yiluhuakai', '45c2dc86118d020ed45f848f9f9f7861', '1272680782@qq.com', 1, '2020-02-04 16:03:57', '2020-02-04 16:03:57');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-02-21  1:15:30
