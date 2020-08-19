/*
 Navicat MySQL Data Transfer

 Source Server         : weixin
 Source Server Type    : MySQL
 Source Server Version : 50505
 Source Host           : localhost
 Source Database       : test

 Target Server Type    : MySQL
 Target Server Version : 50505
 File Encoding         : utf-8

 Date: 08/19/2020 23:25:13 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `account_logs`
-- ----------------------------
DROP TABLE IF EXISTS `account_logs`;
CREATE TABLE `account_logs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `trade_no` varchar(32) DEFAULT '' COMMENT '交易单号，全局唯一',
  `log_no` varchar(32) DEFAULT '' COMMENT '流水编号，全局唯一',
  `account_no` varchar(32) DEFAULT '' COMMENT '账户编号',
  `user_id` varchar(40) DEFAULT '' COMMENT '用户编号',
  `username` varchar(64) DEFAULT '' COMMENT '用户名称',
  `target_user_id` varchar(40) DEFAULT '' COMMENT '目标用户编号',
  `target_username` varchar(64) DEFAULT '' COMMENT '目标用户名称',
  `target_account_no` varchar(32) DEFAULT '' COMMENT '目标账户编号',
  `amount` decimal(30,6) DEFAULT NULL COMMENT '交易金额',
  `balance` decimal(30,6) DEFAULT NULL COMMENT '交易之后的余额',
  `change_type` tinyint(2) DEFAULT NULL COMMENT '交易类型：0创建账户 >0收入',
  `change_flag` tinyint(2) DEFAULT NULL COMMENT '交易标识：-1出账1进账',
  `status` tinyint(2) DEFAULT NULL COMMENT '交易状态',
  `desc` varchar(128) DEFAULT '' COMMENT '交易描述',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
--  Table structure for `accounts`
-- ----------------------------
DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `account_no` varchar(32) DEFAULT '' COMMENT '账户编号，账户唯一标识',
  `account_name` varchar(64) DEFAULT '' COMMENT '账户名称',
  `account_type` tinyint(2) DEFAULT NULL COMMENT '账户类型',
  `currency_code` char(3) DEFAULT NULL COMMENT '货币类型：CNY人民币',
  `user_id` varchar(40) DEFAULT '' COMMENT '用户编号',
  `username` varchar(64) DEFAULT '' COMMENT '用户名称',
  `balance` decimal(30,6) DEFAULT NULL COMMENT '账户可用余额',
  `status` tinyint(2) DEFAULT NULL COMMENT '账户状态',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

SET FOREIGN_KEY_CHECKS = 1;
