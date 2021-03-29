/*
 Navicat Premium Data Transfer

 Source Server         : 本机
 Source Server Type    : MySQL
 Source Server Version : 80019
 Source Host           : localhost:3306
 Source Schema         : audio

 Target Server Type    : MySQL
 Target Server Version : 80019
 File Encoding         : 65001

 Date: 29/03/2021 22:15:14
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for audio_file
-- ----------------------------
DROP TABLE IF EXISTS `audio_file`;
CREATE TABLE `aduio_file` (
  `id` int NOT NULL AUTO_INCREMENT,
  `telephone_kind` varchar(255) NOT NULL,
  `from_telephone_number` varchar(255) NOT NULL,
  `to_telephone_number` varchar(255) NOT NULL,
  `happen_timestamp` int NOT NULL,
  `total_duration` int NOT NULL,
  `file_name` varchar(255) NOT NULL,
  `size` int NOT NULL,
  `file_type` varchar(255) NOT NULL,
  `md5` varchar(255) NOT NULL,
  `file_path` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `from_telephone_number` (`from_telephone_number`),
  KEY `to_telephone_number` (`to_telephone_number`),
  KEY `happen_timestamp` (`happen_timestamp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
