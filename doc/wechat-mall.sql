-- 建库
CREATE DATABASE IF NOT EXISTS `wechat_mall` CHARACTER SET = utf8mb4;

-- 小程序-用户表
CREATE TABLE `wxapp_mall_user` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `openid` varchar(50) NOT NULL DEFAULT '' COMMENT '微信openid',
    `nickname` varchar(30) NOT NULL DEFAULT '' COMMENT '昵称',
    `avatar` varchar(100) NOT NULL DEFAULT '' COMMENT '微信头像',
    `mobile` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
    `city` varchar(30) NOT NULL DEFAULT '' COMMENT '城市',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_openid`(`openid`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '小程序-用户表';

-- 小程序-首页Banner表
CREATE TABLE `wxapp_mall_home_banner` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `pic_url` varchar(100) NOT NULL DEFAULT '' COMMENT 'OSS资源',
    `business_type` int(11) NOT NULL DEFAULT '0' COMMENT '业务类型',
    `business_id` int(11) NOT NULL DEFAULT '0' COMMENT '业务主键',
    `sort` tinyint(2) NOT NULL DEFAULT '0' COMMENT '序号',
    `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态：1-显示 2-隐藏',
    `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '小程序-首页Banner表';

-- 小程序-公告表
CREATE TABLE `wxapp_mall_ext_notice` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title` varchar(30) NOT NULL DEFAULT '' COMMENT '标题',
    `content` text COMMENT '公告内容',
    `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态：1-显示 2-隐藏',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '小程序-公告表';

-- 商城-商品分类表
CREATE TABLE `wxapp_mall_goods_category` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name` varchar(5) NOT NULL DEFAULT '' COMMENT '分类名称',
    `icon` varchar(100) NOT NULL DEFAULT '' COMMENT 'icon资源',
    `sort` tinyint(2) NOT NULL DEFAULT '0' COMMENT '排序',
    `pid` int(11) NOT NULL DEFAULT '0' COMMENT '父级分类',
    `is_use` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否使用：1-是 2-否',
    `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_pid`(`pid`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '商城-商品分类表';

-- 商城-商品表
CREATE TABLE `wxapp_mall_goods` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `goods_name` varchar(100) DEFAULT NULL COMMENT '商品名称',
    `original_price` decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '原价',
    `sale_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '售价',
    `goods_num` int(11) NOT NULL DEFAULT '0' COMMENT '商品数量',
    `picture_url` varchar(255) NOT NULL DEFAULT '' COMMENT '商品图片（首张）',
    `multiple_picture` text COMMENT '商品图片（多张），数组json',
    `video_url` varchar(255) NOT NULL DEFAULT '' COMMENT '视频地址',
    `desc_content` text COMMENT '商品详情，富文本',
    `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '上下架状态：1-上架 2-下架',
    `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`),
    KEY `idx_goods_name`(`goods_name`)
) ENGINE=InnoDB AUTO_INCREMENT = 0 DEFAULT CHARSET=utf8mb4 COMMENT='商城-商品表';

-- 商城-购物车表
CREATE TABLE `wxapp_mall_user_cart` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
    `goods_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品ID',
    `num` int(11) NOT NULL DEFAULT '0' COMMENT '数量',
    `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`),
    KEY `idx_user_id`(`user_id`),
    KEY `idx_goods_id`(`goods_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 0 DEFAULT CHARSET = utf8mb4 COMMENT = '商城-购物车表';

-- 商城-用户收货地址表
CREATE TABLE `wxapp_mall_user_address` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
    `province_id` varchar(10) NOT NULL DEFAULT '' COMMENT '省份编码',
    `city_id` varchar(10) NOT NULL DEFAULT '' COMMENT '城市编码',
    `area_id` varchar(10) NOT NULL DEFAULT '' COMMENT '地区编码',
    `province_str` varchar(10) NOT NULL DEFAULT '' COMMENT '省份',
    `city_str` varchar(10) NOT NULL DEFAULT '' COMMENT '城市',
    `area_str` varchar(10) NOT NULL DEFAULT '' COMMENT '地区',
    `address` varchar(30) NOT NULL DEFAULT '' COMMENT '详细地址',
    `id_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT '默认收货地址：0-否 1-是',
    `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`),
    KEY `idx_user_id`(`user_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 0 DEFAULT CHARSET = utf8mb4 COMMENT = '商城-用户收货地址表';

-- CMS后台用户表
CREATE TABLE `wxapp_mall_cms_user` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `username` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
    `password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
    `email` varchar(20) NOT NULL DEFAULT '' COMMENT '邮箱',
    `mobile` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
    `avatar` varchar(100) NOT NULL DEFAULT '' COMMENT '头像',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY(`id`),
    KEY `idx_email`(`email`),
    KEY `idx_mobile`(`mobile`)
) ENGINE = InnoDB AUTO_INCREMENT = 0 DEFAULT CHARSET = utf8mb4 COMMENT = 'CMS后台用户表';



















