/*
Navicat MySQL Data Transfer

Source Server         : 李长全
Source Server Version : 50527
Source Host           : 127.0.0.1:3306
Source Database       : yoawo

Target Server Type    : MYSQL
Target Server Version : 50527
File Encoding         : 65001

Date: 2019-03-13 17:43:24
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for banner
-- ----------------------------
DROP TABLE IF EXISTS `banner`;
CREATE TABLE `banner` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `area_id` varchar(255) DEFAULT NULL COMMENT '县id',
  `pay_merchant_id` varchar(255) DEFAULT NULL COMMENT '购买商家id',
  `merchant_id` varchar(255) DEFAULT NULL COMMENT '广告商家id',
  `dad` int(11) DEFAULT NULL COMMENT '表id',
  `banner_site` varchar(255) DEFAULT NULL COMMENT '广告位 (轮播1,轮播2,轮播3,轮播4...)',
  `banner_price` int(11) DEFAULT NULL COMMENT '广告位价格(100,200,500)',
  `today_times` int(11) DEFAULT NULL COMMENT '今日点击次数',
  `tick_outs` int(11) DEFAULT NULL COMMENT '累计点击次数',
  `remains` int(11) DEFAULT NULL COMMENT '剩余点击次数',
  `total_times` int(11) DEFAULT NULL COMMENT '总共点击次数',
  `pay_time` int(11) DEFAULT NULL COMMENT '付款时间',
  `show_time` int(11) DEFAULT NULL COMMENT '上架时间',
  `banner_status` int(11) DEFAULT NULL COMMENT '状态  1:等待中 2:上架中 3:已下架 4:删除',
  `banner_url` varchar(255) DEFAULT NULL COMMENT '图片 地址',
  `title` varchar(255) DEFAULT NULL COMMENT '标题',
  `title_sec` varchar(255) DEFAULT NULL COMMENT '副标题',
  `content` varchar(255) DEFAULT NULL COMMENT '内容',
  `show_end` varchar(255) DEFAULT NULL COMMENT '结束时间',
  `pay_status` int(11) DEFAULT NULL COMMENT '支付方式 1:支付宝|默认 2:诺 3:其他',
  `pay_id` varchar(255) DEFAULT NULL COMMENT '订单id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;
