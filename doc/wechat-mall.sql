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

-- 小程序Banner表
CREATE TABLE `wxapp_mall_banner` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `picture` varchar(100) NOT NULL DEFAULT '' COMMENT '图片地址',
    `name` varchar(30) NOT NULL DEFAULT '' COMMENT '名称',
    `title` varchar(30) NOT NULL DEFAULT '' COMMENT '标题',
    `description` varchar(50) NOT NULL DEFAULT '' COMMENT '描述',
    `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_name`(`name`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '小程序Banner表';

-- 小程序Banner子表
CREATE TABLE `wxapp_mall_banner_item` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `banner_id` int(11) NOT NULL DEFAULT '0' COMMENT 'Banner表主键',
    `name` varchar(30) NOT NULL DEFAULT '' COMMENT '名称',
    `picture` varchar(100) NOT NULL DEFAULT '' COMMENT '图片地址',
    `keyword` varchar(10) NOT NULL DEFAULT '' COMMENT '关键字',
    `type` varchar(10) NOT NULL DEFAULT '' COMMENT '类型',
    `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_name`(`name`)
) ENGINE  = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '小程序Banner子表';

-- 商城-SPU分类表
CREATE TABLE `wxapp_mall_spu_category` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父级分类ID',
    `name` varchar(30) NOT NULL DEFAULT '' COMMENT '分类名称',
    `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
    `online` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否上线：0-否 1-是',
    `picture` varchar(100) NOT NULL DEFAULT '' COMMENT '图片地址',
    `description` varchar(50) NOT NULL DEFAULT '' COMMENT '分类描述',
    `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`id`),
    KEY `idx_parent_id`(`parent_id`),
    KEY `idx_name`(`name`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '商城-SPU分类表';

-- 商城-首页宫格表
CREATE TABLE `wxapp_mall_grid_category` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title` varchar(30) NOT NULL DEFAULT '' COMMENT '宫格标题',
    `name` varchar(30) NOT NULL DEFAULT '' COMMENT '宫格名',
    `category_id` int(11) NOT NULL DEFAULT '0' COMMENT '顶级分类ID',
    `picture` varchar(100) NOT NULL DEFAULT '' COMMENT '图片地址',
    `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`id`),
    KEY `idx_category_id`(`category_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '商城-首页宫格表';

-- 商城-规格表
CREATE TABLE `wxapp_mall_specification` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name` varchar(30) NOT NULL DEFAULT '' COMMENT '规格名名称',
    `description` varchar(30) NOT NULL DEFAULT '' COMMENT '规格名描述',
    `unit` varchar(10) NOT NULL DEFAULT '' COMMENT '单位',
    `standard` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否标准: 0-非标准 1-标准',
    `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`id`),
    KEY `idx_name`(`name`)
) ENGINE = InnoDB auto_increment = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '商城-规格表';

-- 商城-规格属性表
CREATE TABLE `wxapp_mall_specification_attr` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `spec_id` int(11) NOT NULL DEFAULT '' COMMENT '规格ID',
    `value` varchar(10) NOT NULL DEFAULT '' COMMENT '属性值',
    `extend` varchar(30) NOT NULL DEFAULT '' COMMENT '扩展',
    `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`id`),
    KEY `idx_spec_id`(`spec_id`),
    KEY `idx_value`(`value`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '商城-规格属性表';

-- 商城-SPU表
CREATE TABLE `wxapp_mall_spu` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `brand_name` varchar(30) NOT NULL DEFAULT '' COMMENT '品牌名称',
    `title` varchar(30) NOT NULL DEFAULT '' COMMENT '标题',
    `subtitle` varchar(30) NOT NULL DEFAULT '' COMMENT '副标题',
    `price` decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '价格',
    `discount_price` decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '折扣',
    `category_id` int(11) NOT NULL DEFAULT '0' COMMENT '分类ID',
    `default_sku_id` int(11) NOT NULL DEFAULT '0' COMMENT '默认sku',
    `online` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否上架：0-下架 1-上架',
    `picture` varchar(100) NOT NULL DEFAULT '' COMMENT '主图',
    `for_theme_picture` varchar(100) NOT NULL DEFAULT '' COMMENT '主题图',
    `banner_picture` text COMMENT '轮播图',
    `detail_picture` text COMMENT '详情图',
    `tags` varchar(100) NOT NULL DEFAULT '' COMMENT '标签，示例：包邮$热门',
    `sketch_spec_id` int(11) NOT NULL DEFAULT '0' COMMENT '可视规格',
    `description` varchar(30) NOT NULL DEFAULT '' COMMENT '描述',
    `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`id`),
    KEY `idx_title`(`title`),
    KEY `idx_category_id`(`category_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '商城-SPU表';

-- 商城-SPU-规格表
CREATE TABLE `wxapp_mall_spu_spec` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `spu_id` int(11) NOT NULL DEFAULT '0' COMMENT 'SPU ID',
    `spec_id` int(11) NOT NULL DEFAULT '0' COMMENT '规格ID',
    `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`id`),
    KEY `idx_spu_id`(`spu_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '商城-SPU-规格表';

-- 商城-SKU表
CREATE TABLE `wxapp_mall_sku` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title` varchar(10) NOT NULL DEFAULT '' COMMENT '标题',
    `price` decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '价格',
    `code` varchar(30) NOT NULL DEFAULT '' COMMENT '编码',
    `stock` int(11) NOT NULL DEFAULT '0' COMMENT '库存量',
    `spu_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属SPU',
    `online` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否上架: 0-下架 1-上架',
    `picture` varchar(100) NOT NULL DEFAULT '' COMMENT '图片',
    `specs` varchar(100) NOT NULL DEFAULT '' COMMENT '规格属性',
    `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`id`),
    KEY `idx_spu_id`(`spu_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '商城-SKU表';

-- 商城-活动表
CREATE TABLE `wxapp_mall_activity` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主题',
    `title` varchar(30) NOT NULL DEFAULT '' COMMENT '标题',
    `name` varchar(30) NOT NULL DEFAULT '' COMMENT '活动名称',
    `remark` varchar(30) NOT NULL DEFAULT '' COMMENT '提示',
    `online` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否上线：0-否 1-是',
    `start_time` datetime NOT NULL COMMENT '开始时间',
    `end_time` datetime NOT NULL COMMENT '结束时间',
    `description` varchar(50) NOT NULL DEFAULT '' COMMENT '描述',
    `entrance_picture` varchar(100) NOT NULL DEFAULT '' COMMENT '主图',
    `internal_top_picture` varchar(100) NOT NULL DEFAULT '' COMMENT '顶部图',
    `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '商城-活动表';

-- 商城-优惠券表
CREATE TABLE `wxapp_mall_coupon` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `activity_id` int(11) NOT NULL DEFAULT '0' COMMENT '活动ID',
    `title` varchar(30) NOT NULL DEFAULT '' COMMENT '标题',
    `full_money` decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '满减额',
    `minus` decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '优惠额',
    `rate` decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '折扣',
    `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '券类型：1-满减券 2-折扣券 3-无门槛券 4-满金额折扣券',
    `start_time` datetime NOT NULL COMMENT '开始时间',
    `end_time` datetime NOT NULL COMMENT '结束时间',
    `description` VARCHAR(30) NOT NULL DEFAULT '' COMMENT '描述',
    `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY(`id`),
    KEY `idx_activity_id`(`activity_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '商城-优惠券表';

-- 商城-购物车表
CREATE TABLE `wxapp_mall_user_cart` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
    `spu_id` int(11) NOT NULL DEFAULT '0' COMMENT 'SPU ID',
    `num` int(11) NOT NULL DEFAULT '0' COMMENT '数量',
    `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除：0-否 1-是',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`),
    KEY `idx_user_id`(`user_id`),
    KEY `idx_spu_id`(`spu_id`)
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



















