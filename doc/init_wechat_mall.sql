/*
 Navicat Premium Data Transfer

 Source Server         : 47.98.199.80
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : 47.98.199.80:3306
 Source Schema         : wechat_mall

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 05/05/2020 10:57:25
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for wechat_mall_banner
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_banner`;
CREATE TABLE `wechat_mall_banner` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `picture` varchar(200) NOT NULL DEFAULT '' COMMENT '图片地址',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '名称',
  `business_type` tinyint(2) NOT NULL DEFAULT '0' COMMENT '业务类型：1-商品',
  `business_id` int(1) NOT NULL DEFAULT '0' COMMENT '业务主键',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否显示：0-否 1-是',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='小程序Banner表';

-- ----------------------------
-- Table structure for wechat_mall_category
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_category`;
CREATE TABLE `wechat_mall_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父级分类ID',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '分类名称',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `online` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否上线：0-否 1-是',
  `picture` varchar(200) NOT NULL DEFAULT '' COMMENT '图片地址',
  `description` varchar(50) NOT NULL DEFAULT '' COMMENT '分类描述',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城-分类表';

-- ----------------------------
-- Table structure for wechat_mall_cms_user
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_cms_user`;
CREATE TABLE `wechat_mall_cms_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
  `email` varchar(20) NOT NULL DEFAULT '' COMMENT '邮箱',
  `mobile` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `avatar` varchar(200) NOT NULL DEFAULT '' COMMENT '头像',
  `group_id` int(11) NOT NULL DEFAULT '0' COMMENT '分组ID',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_email` (`email`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='CMS后台用户表';

-- ----------------------------
-- Records of wechat_mall_cms_user
-- ----------------------------
BEGIN;
INSERT INTO `wechat_mall_cms_user` VALUES (1, 'admin', '4d6a203f65ffa22dc78a8abd71822422', '', '', 'http://i1.sleeve.7yue.pro/af72b2f7-c889-47dd-82a1-80d8a796b044.jpg', 0, 0, '2020-05-05 10:54:35', '2020-05-05 10:54:35');
COMMIT;

-- ----------------------------
-- Table structure for wechat_mall_coupon
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_coupon`;
CREATE TABLE `wechat_mall_coupon` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(30) NOT NULL DEFAULT '' COMMENT '标题',
  `full_money` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '满减额',
  `minus` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '优惠额',
  `rate` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '折扣',
  `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '券类型：1-满减券 2-折扣券 3-代金券 4-满金额折扣券',
  `grant_num` int(11) NOT NULL DEFAULT '0' COMMENT '发券数量',
  `limit_num` int(11) NOT NULL DEFAULT '0' COMMENT '单人限领',
  `start_time` datetime NOT NULL COMMENT '开始时间',
  `end_time` datetime NOT NULL COMMENT '结束时间',
  `description` varchar(30) NOT NULL DEFAULT '' COMMENT '描述',
  `online` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否上架: 0-下架 1-上架',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城-优惠券表';

-- ----------------------------
-- Table structure for wechat_mall_coupon_log
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_coupon_log`;
CREATE TABLE `wechat_mall_coupon_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `coupon_id` int(11) NOT NULL DEFAULT '0' COMMENT '优惠券ID',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `use_time` datetime DEFAULT NULL COMMENT '使用时间',
  `expire_time` datetime NOT NULL COMMENT '过期时间',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '状态：0-未使用 1-已使用 2-已过期',
  `code` varchar(30) NOT NULL DEFAULT '' COMMENT '兑换码',
  `order_no` varchar(30) NOT NULL DEFAULT '' COMMENT '核销的订单号',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_coupon_id` (`coupon_id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城-优惠券领取记录表';

-- ----------------------------
-- Table structure for wechat_mall_goods
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_goods`;
CREATE TABLE `wechat_mall_goods` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `brand_name` varchar(30) NOT NULL DEFAULT '' COMMENT '品牌名称',
  `title` varchar(80) NOT NULL DEFAULT '' COMMENT '标题',
  `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '价格',
  `discount_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '折扣',
  `category_id` int(11) NOT NULL DEFAULT '0' COMMENT '分类ID',
  `online` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否上架：0-下架 1-上架',
  `picture` varchar(200) NOT NULL DEFAULT '' COMMENT '主图',
  `banner_picture` text COMMENT '轮播图',
  `detail_picture` text COMMENT '详情图',
  `tags` varchar(100) NOT NULL DEFAULT '' COMMENT '标签，示例：包邮$热门',
  `sale_num` int(11) NOT NULL DEFAULT '0' COMMENT '商品销量',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_title` (`title`),
  KEY `idx_category_id` (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城-商品表';

-- ----------------------------
-- Table structure for wechat_mall_goods_browse_record
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_goods_browse_record`;
CREATE TABLE `wechat_mall_goods_browse_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `goods_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品ID',
  `picture` varchar(200) NOT NULL DEFAULT '' COMMENT '商品图片',
  `title` varchar(80) NOT NULL DEFAULT '' COMMENT '商品名称',
  `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '商品价格',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_goods_id` (`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城-商品浏览记录';

-- ----------------------------
-- Table structure for wechat_mall_goods_spec
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_goods_spec`;
CREATE TABLE `wechat_mall_goods_spec` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `goods_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品ID',
  `spec_id` int(11) NOT NULL DEFAULT '0' COMMENT '规格ID',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_goods_id` (`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城-商品规格表';

-- ----------------------------
-- Table structure for wechat_mall_grid_category
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_grid_category`;
CREATE TABLE `wechat_mall_grid_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(30) NOT NULL DEFAULT '' COMMENT '宫格标题',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '宫格名',
  `category_id` int(11) NOT NULL DEFAULT '0' COMMENT '顶级分类ID',
  `picture` varchar(200) NOT NULL DEFAULT '' COMMENT '图片地址',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_category_id` (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='小程序-首页宫格表';

-- ----------------------------
-- Table structure for wechat_mall_group_page_permission
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_group_page_permission`;
CREATE TABLE `wechat_mall_group_page_permission` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `group_id` int(11) NOT NULL DEFAULT '0' COMMENT '分组ID',
  `page_id` int(11) NOT NULL DEFAULT '0' COMMENT '页面ID',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='CMS-用户分组-页面权限表';

-- ----------------------------
-- Table structure for wechat_mall_module
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_module`;
CREATE TABLE `wechat_mall_module` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(10) NOT NULL DEFAULT '' COMMENT '模块名称',
  `description` varchar(30) NOT NULL DEFAULT '' COMMENT '描述',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='CMS-组成模块';

-- ----------------------------
-- Records of wechat_mall_module
-- ----------------------------
BEGIN;
INSERT INTO `wechat_mall_module` VALUES (1, '商城', '商城管理', 0, '2020-03-14 08:50:49', '2020-03-14 08:50:49');
INSERT INTO `wechat_mall_module` VALUES (2, '商品', '商品管理', 0, '2020-03-14 08:50:59', '2020-03-14 08:50:59');
INSERT INTO `wechat_mall_module` VALUES (3, '订单', '订单管理', 0, '2020-03-14 08:51:09', '2020-03-14 08:51:09');
INSERT INTO `wechat_mall_module` VALUES (4, '营销', '营销管理', 0, '2020-03-14 08:51:17', '2020-03-14 08:51:17');
COMMIT;

-- ----------------------------
-- Table structure for wechat_mall_module_page
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_module_page`;
CREATE TABLE `wechat_mall_module_page` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `module_id` int(11) NOT NULL DEFAULT '0' COMMENT '模块ID',
  `name` varchar(10) NOT NULL DEFAULT '' COMMENT '页面名称',
  `description` varchar(30) NOT NULL DEFAULT '' COMMENT '描述',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_module_id` (`module_id`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COMMENT='CMS-模块页面';

-- ----------------------------
-- Records of wechat_mall_module_page
-- ----------------------------
BEGIN;
INSERT INTO `wechat_mall_module_page` VALUES (1, 1, '查看Banner', '', 0, '2020-03-14 08:51:40', '2020-03-14 08:51:40');
INSERT INTO `wechat_mall_module_page` VALUES (2, 1, '管理Banner', '', 0, '2020-03-14 08:52:02', '2020-03-14 08:52:02');
INSERT INTO `wechat_mall_module_page` VALUES (3, 1, '查看宫格', '', 0, '2020-03-14 08:52:14', '2020-03-14 08:52:14');
INSERT INTO `wechat_mall_module_page` VALUES (4, 1, '管理宫格', '', 0, '2020-03-14 08:52:25', '2020-03-14 08:52:25');
INSERT INTO `wechat_mall_module_page` VALUES (5, 2, '查看商品分类', '', 0, '2020-03-14 08:52:42', '2020-03-14 08:52:42');
INSERT INTO `wechat_mall_module_page` VALUES (6, 2, '管理商品分类', '', 0, '2020-03-14 08:52:53', '2020-03-14 08:52:53');
INSERT INTO `wechat_mall_module_page` VALUES (7, 2, '查看规格', '', 0, '2020-03-14 08:53:05', '2020-03-14 08:53:05');
INSERT INTO `wechat_mall_module_page` VALUES (8, 2, '管理规格', '', 0, '2020-03-14 08:53:13', '2020-03-14 08:53:13');
INSERT INTO `wechat_mall_module_page` VALUES (9, 2, '查看商品', '', 0, '2020-05-04 16:52:27', '2020-05-04 16:52:27');
INSERT INTO `wechat_mall_module_page` VALUES (10, 2, '管理商品', ' ', 0, '2020-05-04 16:52:27', '2020-05-04 16:52:27');
INSERT INTO `wechat_mall_module_page` VALUES (11, 2, '查看库存', ' ', 0, '2020-05-04 16:52:27', '2020-05-04 16:52:27');
INSERT INTO `wechat_mall_module_page` VALUES (12, 2, '管理库存', ' ', 0, '2020-05-04 16:52:27', '2020-05-04 16:52:27');
INSERT INTO `wechat_mall_module_page` VALUES (13, 3, '查看订单', ' ', 0, '2020-05-04 16:52:27', '2020-05-04 16:52:27');
INSERT INTO `wechat_mall_module_page` VALUES (14, 3, '管理订单', ' ', 0, '2020-05-04 16:52:27', '2020-05-04 16:52:27');
INSERT INTO `wechat_mall_module_page` VALUES (15, 4, '查看优惠券', ' ', 0, '2020-05-04 16:52:27', '2020-05-04 16:52:27');
INSERT INTO `wechat_mall_module_page` VALUES (16, 4, '管理优惠券', ' ', 0, '2020-05-04 16:52:27', '2020-05-04 16:52:27');
COMMIT;

-- ----------------------------
-- Table structure for wechat_mall_order
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_order`;
CREATE TABLE `wechat_mall_order` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `order_no` varchar(32) NOT NULL DEFAULT '' COMMENT '订单号',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `pay_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '订单金额（商品金额 + 运费 - 优惠金额）',
  `goods_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '商品小计金额',
  `discount_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '优惠金额',
  `dispatch_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '运费',
  `pay_time` datetime NOT NULL DEFAULT '2006-01-02 15:04:05' COMMENT '支付时间',
  `deliver_time` datetime NOT NULL DEFAULT '2006-01-02 15:04:05' COMMENT '发货时间',
  `finish_time` datetime NOT NULL DEFAULT '2006-01-02 15:04:05' COMMENT '成交时间',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态 -1 已取消 0-待付款 1-待发货 2-待收货 3-已完成 4-（待发货）退款申请 5-已退款\n',
  `address_id` int(11) NOT NULL DEFAULT '0' COMMENT '收货地址ID',
  `address_snapshot` varchar(200) NOT NULL DEFAULT '' COMMENT '收货地址快照',
  `wxapp_prepay_id` varchar(50) NOT NULL DEFAULT '' COMMENT '微信预支付ID',
  `transaction_id` varchar(50) NOT NULL DEFAULT '' COMMENT '微信支付单号',
  `remark` varchar(100) NOT NULL DEFAULT '' COMMENT '订单备注',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_order_no` (`order_no`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城订单表';

-- ----------------------------
-- Table structure for wechat_mall_order_goods
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_order_goods`;
CREATE TABLE `wechat_mall_order_goods` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `order_no` varchar(30) NOT NULL DEFAULT '' COMMENT '订单号',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `goods_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品ID',
  `sku_id` int(11) NOT NULL DEFAULT '0' COMMENT 'sku ID',
  `picture` varchar(200) NOT NULL DEFAULT '' COMMENT '商品图片',
  `title` varchar(80) NOT NULL DEFAULT '' COMMENT '商品标题',
  `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '价格',
  `specs` varchar(500) NOT NULL DEFAULT '' COMMENT 'sku规格属性',
  `num` int(11) NOT NULL DEFAULT '0' COMMENT '数量',
  `lock_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '锁定状态：0-锁定 1-解锁',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_order_no` (`order_no`),
  KEY `idx_goods_id` (`goods_id`),
  KEY `idx_sku_id` (`sku_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城订单-商品表';

-- ----------------------------
-- Table structure for wechat_mall_order_refund
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_order_refund`;
CREATE TABLE `wechat_mall_order_refund` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `refund_no` varchar(30) NOT NULL DEFAULT '' COMMENT '退款编号',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '平台用户ID',
  `order_no` varchar(30) NOT NULL DEFAULT '' COMMENT '订单号',
  `reason` varchar(30) NOT NULL DEFAULT '' COMMENT '退款原因',
  `refund_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '退款金额',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态：0-退款申请 1-退款完成 2-撤销申请',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `refund_time` datetime NOT NULL DEFAULT '2006-01-02 15:04:05' COMMENT '退款时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城-订单退款申请表';

-- ----------------------------
-- Table structure for wechat_mall_sku
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_sku`;
CREATE TABLE `wechat_mall_sku` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(20) NOT NULL DEFAULT '' COMMENT '标题',
  `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '价格',
  `code` varchar(30) NOT NULL DEFAULT '' COMMENT '编码',
  `stock` int(11) NOT NULL DEFAULT '0' COMMENT '库存量',
  `goods_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属商品',
  `online` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否上架: 0-下架 1-上架',
  `picture` varchar(200) NOT NULL DEFAULT '' COMMENT '图片',
  `specs` varchar(500) NOT NULL DEFAULT '' COMMENT '规格属性',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_title` (`title`),
  KEY `idx_code` (`code`),
  KEY `idx_goods_id` (`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城-SKU表';

-- ----------------------------
-- Table structure for wechat_mall_sku_spec_attr
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_sku_spec_attr`;
CREATE TABLE `wechat_mall_sku_spec_attr` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `sku_id` int(11) NOT NULL DEFAULT '0' COMMENT 'sku表主键',
  `spec_id` int(11) NOT NULL DEFAULT '0' COMMENT '规格ID',
  `attr_id` int(11) NOT NULL DEFAULT '0' COMMENT '规格-属性ID',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_sku_id` (`sku_id`),
  KEY `idx_spec_id` (`spec_id`),
  KEY `idx_attr_id` (`attr_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城-SKU关联的规格属性';

-- ----------------------------
-- Table structure for wechat_mall_specification
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_specification`;
CREATE TABLE `wechat_mall_specification` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '规格名名称',
  `description` varchar(30) NOT NULL DEFAULT '' COMMENT '规格名描述',
  `unit` varchar(10) NOT NULL DEFAULT '' COMMENT '单位',
  `standard` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否标准: 0-非标准 1-标准',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城-规格表';

-- ----------------------------
-- Table structure for wechat_mall_specification_attr
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_specification_attr`;
CREATE TABLE `wechat_mall_specification_attr` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `spec_id` int(11) NOT NULL DEFAULT '0' COMMENT '规格ID',
  `value` varchar(20) NOT NULL DEFAULT '' COMMENT '属性值',
  `extend` varchar(30) NOT NULL DEFAULT '' COMMENT '扩展',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_spec_id` (`spec_id`),
  KEY `idx_value` (`value`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城-规格属性表';

-- ----------------------------
-- Table structure for wechat_mall_user
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_user`;
CREATE TABLE `wechat_mall_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `openid` varchar(50) NOT NULL DEFAULT '' COMMENT '微信openid',
  `nickname` varchar(30) NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(200) NOT NULL DEFAULT '' COMMENT '微信头像',
  `mobile` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `city` varchar(30) NOT NULL DEFAULT '' COMMENT '城市',
  `province` varchar(30) NOT NULL DEFAULT '' COMMENT '省份',
  `country` varchar(30) NOT NULL DEFAULT '' COMMENT '国家',
  `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别 0：未知、1：男、2：女',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_openid` (`openid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='小程序-用户表';

-- ----------------------------
-- Table structure for wechat_mall_user_address
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_user_address`;
CREATE TABLE `wechat_mall_user_address` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `contacts` varchar(15) NOT NULL DEFAULT '' COMMENT '联系人',
  `mobile` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `province_id` varchar(10) NOT NULL DEFAULT '' COMMENT '省份编码',
  `city_id` varchar(10) NOT NULL DEFAULT '' COMMENT '城市编码',
  `area_id` varchar(10) NOT NULL DEFAULT '' COMMENT '地区编码',
  `province_str` varchar(10) NOT NULL DEFAULT '' COMMENT '省份',
  `city_str` varchar(10) NOT NULL DEFAULT '' COMMENT '城市',
  `area_str` varchar(10) NOT NULL DEFAULT '' COMMENT '地区',
  `address` varchar(30) NOT NULL DEFAULT '' COMMENT '详细地址',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT '默认收货地址：0-否 1-是',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城-用户收货地址表';

-- ----------------------------
-- Table structure for wechat_mall_user_cart
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_user_cart`;
CREATE TABLE `wechat_mall_user_cart` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `goods_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品ID',
  `sku_id` int(11) NOT NULL DEFAULT '0' COMMENT 'sku ID',
  `num` int(11) NOT NULL DEFAULT '0' COMMENT '数量',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_goods_id` (`goods_id`),
  KEY `idx_sku_id` (`sku_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城-购物车表';

-- ----------------------------
-- Table structure for wechat_mall_user_group
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_user_group`;
CREATE TABLE `wechat_mall_user_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(10) NOT NULL DEFAULT '' COMMENT '名称',
  `description` varchar(30) NOT NULL DEFAULT '' COMMENT '分组描述',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='CMS-后台用户分组表';

-- ----------------------------
-- Table structure for wechat_mall_visitor_record
-- ----------------------------
DROP TABLE IF EXISTS `wechat_mall_visitor_record`;
CREATE TABLE `wechat_mall_visitor_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '平台用户ID',
  `ip` varchar(20) NOT NULL DEFAULT '' COMMENT '独立IP',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城-访客记录表';

SET FOREIGN_KEY_CHECKS = 1;
