/*
 Navicat Premium Data Transfer

 Source Server         : Local DB
 Source Server Type    : MySQL
 Source Server Version : 100418
 Source Host           : localhost:3306
 Source Schema         : game_currency

 Target Server Type    : MySQL
 Target Server Version : 100418
 File Encoding         : 65001

 Date: 28/11/2021 08:45:25
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for conversions
-- ----------------------------
DROP TABLE IF EXISTS `conversions`;
CREATE TABLE `conversions`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `currency_id_from` int NOT NULL,
  `currency_id_to` int NULL DEFAULT NULL,
  `rate` int NULL DEFAULT NULL,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_currency_id_from`(`currency_id_from`) USING BTREE,
  INDEX `fk_currency_id_to`(`currency_id_to`) USING BTREE,
  CONSTRAINT `fk_currency_id_from` FOREIGN KEY (`currency_id_from`) REFERENCES `currency` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_currency_id_to` FOREIGN KEY (`currency_id_to`) REFERENCES `currency` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of conversions
-- ----------------------------
INSERT INTO `conversions` VALUES (16, 2, 1, 29, NULL, NULL);

-- ----------------------------
-- Table structure for currency
-- ----------------------------
DROP TABLE IF EXISTS `currency`;
CREATE TABLE `currency`  (
  `id` int NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of currency
-- ----------------------------
INSERT INTO `currency` VALUES (1, 'Knut');
INSERT INTO `currency` VALUES (2, 'Sickle');
INSERT INTO `currency` VALUES (3, 'Galleon');

SET FOREIGN_KEY_CHECKS = 1;
