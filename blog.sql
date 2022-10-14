/*
 Navicat Premium Data Transfer

 Source Server         : mysql8_spxzx
 Source Server Type    : MySQL
 Source Server Version : 50733
 Source Host           : localhost:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 50733
 File Encoding         : 65001

 Date: 14/10/2022 13:08:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL COMMENT '作者',
  `category_id` int(11) NULL DEFAULT NULL COMMENT '文章分类',
  `article_cover` varchar(1024) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文章缩略图',
  `article_title` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标题',
  `article_content` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '内容',
  `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '文章类型 1原创 2转载 3翻译',
  `original_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '原文链接',
  `is_top` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否置顶 0否 1是',
  `is_delete` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除  0否 1是',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态值 1公开 2私密 3评论可见',
  `create_time` datetime NOT NULL COMMENT '发表时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES (1, 1, 1, 'https://blog-1311853727.cos.ap-guangzhou.myqcloud.com/article/anime.jpg', '测试文章', '# 测试文章内容\n## 测试文章内容\n### 测试文章内容\n#### 测试文章内容\n##### 测试文章内容\n###### 测试文章内容\n![anime.jpg](https://blog-1311853727.cos.ap-guangzhou.myqcloud.com/article/anime.jpg)', 1, '', 0, 0, 1, '2022-10-14 11:51:09', '2022-10-14 11:51:09');

-- ----------------------------
-- Table structure for article_tag
-- ----------------------------
DROP TABLE IF EXISTS `article_tag`;
CREATE TABLE `article_tag`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NOT NULL COMMENT '文章id',
  `tag_id` int(11) NOT NULL COMMENT '标签id',
  PRIMARY KEY (`id`, `article_id`, `tag_id`) USING BTREE,
  INDEX `fk_article_tag_1`(`article_id`) USING BTREE,
  INDEX `fk_article_tag_2`(`tag_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of article_tag
-- ----------------------------
INSERT INTO `article_tag` VALUES (1, 1, 1);

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1235 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (140, 'g', 'admin', 'anonymous', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (665, 'g', 'admin', 'logout', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (441, 'g', 'admin', 'upload', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (142, 'g', 'test', 'anonymous', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (667, 'g', 'test', 'logout', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (141, 'g', 'user', 'anonymous', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (666, 'g', 'user', 'logout', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1116, 'p', 'admin', '/admin', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1117, 'p', 'admin', '/admin/about', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1121, 'p', 'admin', '/admin/articles', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1118, 'p', 'admin', '/admin/articles', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1119, 'p', 'admin', '/admin/articles', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1120, 'p', 'admin', '/admin/articles', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1124, 'p', 'admin', '/admin/articles/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1122, 'p', 'admin', '/admin/articles/images', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1123, 'p', 'admin', '/admin/articles/top', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1127, 'p', 'admin', '/admin/categories', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1125, 'p', 'admin', '/admin/categories', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1126, 'p', 'admin', '/admin/categories', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1128, 'p', 'admin', '/admin/categories/search', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1130, 'p', 'admin', '/admin/comments', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1129, 'p', 'admin', '/admin/comments', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1131, 'p', 'admin', '/admin/comments/review', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1134, 'p', 'admin', '/admin/links', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1132, 'p', 'admin', '/admin/links', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1133, 'p', 'admin', '/admin/links', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1135, 'p', 'admin', '/admin/menus', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1136, 'p', 'admin', '/admin/menus', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1137, 'p', 'admin', '/admin/menus/:menuId', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1139, 'p', 'admin', '/admin/messages', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1138, 'p', 'admin', '/admin/messages', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1140, 'p', 'admin', '/admin/messages/review', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1142, 'p', 'admin', '/admin/operation/logs', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1141, 'p', 'admin', '/admin/operation/logs', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1143, 'p', 'admin', '/admin/pages', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1144, 'p', 'admin', '/admin/pages', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1145, 'p', 'admin', '/admin/pages/:pageId', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1149, 'p', 'admin', '/admin/photos', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1146, 'p', 'admin', '/admin/photos', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1147, 'p', 'admin', '/admin/photos', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1148, 'p', 'admin', '/admin/photos', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1150, 'p', 'admin', '/admin/photos/album', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1151, 'p', 'admin', '/admin/photos/albums', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1152, 'p', 'admin', '/admin/photos/albums', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1155, 'p', 'admin', '/admin/photos/albums/:albumId', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1156, 'p', 'admin', '/admin/photos/albums/:albumId/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1153, 'p', 'admin', '/admin/photos/albums/cover', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1154, 'p', 'admin', '/admin/photos/albums/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1157, 'p', 'admin', '/admin/photos/delete', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1158, 'p', 'admin', '/admin/resources', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1159, 'p', 'admin', '/admin/resources', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1161, 'p', 'admin', '/admin/resources/:resourceId', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1160, 'p', 'admin', '/admin/resources/import/swagger', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1162, 'p', 'admin', '/admin/role', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1163, 'p', 'admin', '/admin/role/menus', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1164, 'p', 'admin', '/admin/role/resources', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1166, 'p', 'admin', '/admin/roles', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1165, 'p', 'admin', '/admin/roles', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1169, 'p', 'admin', '/admin/tags', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1167, 'p', 'admin', '/admin/tags', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1168, 'p', 'admin', '/admin/tags', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1170, 'p', 'admin', '/admin/tags/search', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1191, 'p', 'admin', '/admin/talks', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1192, 'p', 'admin', '/admin/talks', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1190, 'p', 'admin', '/admin/talks', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1193, 'p', 'admin', '/admin/talks/:talkId', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1189, 'p', 'admin', '/admin/talks/images', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1171, 'p', 'admin', '/admin/user/menus', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1172, 'p', 'admin', '/admin/users', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1178, 'p', 'admin', '/admin/users/:userInfoId/online', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (1187, 'p', 'admin', '/admin/users/area', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1173, 'p', 'admin', '/admin/users/disable', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1174, 'p', 'admin', '/admin/users/online', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1175, 'p', 'admin', '/admin/users/password', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1176, 'p', 'admin', '/admin/users/role', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1177, 'p', 'admin', '/admin/users/role', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1179, 'p', 'admin', '/admin/website/config', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1180, 'p', 'admin', '/admin/website/config', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1181, 'p', 'admin', '/articles/:articleId/like', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1182, 'p', 'admin', '/comments', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1183, 'p', 'admin', '/comments/:commentId/like', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1188, 'p', 'admin', '/talks/:talkId/like', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1184, 'p', 'admin', '/users/avatar', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1185, 'p', 'admin', '/users/email', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1186, 'p', 'admin', '/users/info', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (2, 'p', 'anonymous', '/', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (3, 'p', 'anonymous', '/about', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (13, 'p', 'anonymous', '/albums/*/photos', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (891, 'p', 'anonymous', '/albums/:albumId/photos', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (7, 'p', 'anonymous', '/articles', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (11, 'p', 'anonymous', '/articles/*', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (8, 'p', 'anonymous', '/articles/archives', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (9, 'p', 'anonymous', '/articles/condition', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (10, 'p', 'anonymous', '/articles/search', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1, 'p', 'anonymous', '/categories', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (22, 'p', 'anonymous', '/comments', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (23, 'p', 'anonymous', '/comments/*/replies', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1115, 'p', 'anonymous', '/comments/:commentId/replies', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (24, 'p', 'anonymous', '/home/talks', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (6, 'p', 'anonymous', '/links', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (19, 'p', 'anonymous', '/messages', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (20, 'p', 'anonymous', '/messages', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (21, 'p', 'anonymous', '/photos/albums', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (14, 'p', 'anonymous', '/register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (5, 'p', 'anonymous', '/report', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (12, 'p', 'anonymous', '/tags', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (25, 'p', 'anonymous', '/talks', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (26, 'p', 'anonymous', '/talks/*', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1234, 'p', 'anonymous', '/users/avatar', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (15, 'p', 'anonymous', '/users/code', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (16, 'p', 'anonymous', '/users/oauth/qq', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (17, 'p', 'anonymous', '/users/oauth/weibo', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (18, 'p', 'anonymous', '/users/password', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (4, 'p', 'anonymous', '/voice', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (668, 'p', 'logout', '/logout', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1201, 'p', 'test', '/admin', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1202, 'p', 'test', '/admin/articles', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1203, 'p', 'test', '/admin/articles/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1204, 'p', 'test', '/admin/categories', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1205, 'p', 'test', '/admin/categories/search', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1206, 'p', 'test', '/admin/comments', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1207, 'p', 'test', '/admin/links', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1208, 'p', 'test', '/admin/menus', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1209, 'p', 'test', '/admin/messages', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1210, 'p', 'test', '/admin/operation/logs', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1211, 'p', 'test', '/admin/pages', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1212, 'p', 'test', '/admin/photos', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1213, 'p', 'test', '/admin/photos/albums', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1215, 'p', 'test', '/admin/photos/albums/:albumId/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1214, 'p', 'test', '/admin/photos/albums/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1216, 'p', 'test', '/admin/resources', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1217, 'p', 'test', '/admin/role/menus', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1218, 'p', 'test', '/admin/role/resources', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1219, 'p', 'test', '/admin/roles', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1220, 'p', 'test', '/admin/tags', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1221, 'p', 'test', '/admin/tags/search', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1232, 'p', 'test', '/admin/talks', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1233, 'p', 'test', '/admin/talks/:talkId', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1222, 'p', 'test', '/admin/user/menus', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1223, 'p', 'test', '/admin/users', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1230, 'p', 'test', '/admin/users/area', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1224, 'p', 'test', '/admin/users/online', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1225, 'p', 'test', '/admin/users/role', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1226, 'p', 'test', '/admin/website/config', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1227, 'p', 'test', '/articles/:articleId/like', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1228, 'p', 'test', '/comments', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1229, 'p', 'test', '/comments/:commentId/like', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1231, 'p', 'test', '/talks/:talkId/like', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (442, 'p', 'upload', '/admin/config/images', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1194, 'p', 'user', '/articles/:articleId/like', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1195, 'p', 'user', '/comments', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1196, 'p', 'user', '/comments/:commentId/like', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1200, 'p', 'user', '/talks/:talkId/like', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1197, 'p', 'user', '/users/avatar', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1198, 'p', 'user', '/users/email', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1199, 'p', 'user', '/users/info', 'PUT', '', '', '');

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `category_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '分类名',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES (1, '测试分类', '2022-10-14 11:49:46', '2022-10-14 11:49:46');

-- ----------------------------
-- Table structure for chat_record
-- ----------------------------
DROP TABLE IF EXISTS `chat_record`;
CREATE TABLE `chat_record`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(11) NULL DEFAULT NULL COMMENT '用户id',
  `nickname` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '昵称',
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '头像',
  `content` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '聊天内容',
  `ip_address` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'ip地址',
  `ip_source` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'ip来源',
  `type` tinyint(4) NOT NULL COMMENT '类型',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2990 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of chat_record
-- ----------------------------

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(11) NOT NULL COMMENT '评论用户Id',
  `topic_id` int(11) NULL DEFAULT NULL COMMENT '评论主题id',
  `comment_content` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '评论内容',
  `reply_user_id` int(11) NULL DEFAULT NULL COMMENT '回复用户id',
  `parent_id` int(11) NULL DEFAULT NULL COMMENT '父评论id',
  `type` tinyint(4) NOT NULL COMMENT '评论类型 1.文章 2.友链 3.说说',
  `is_delete` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否删除  0否 1是',
  `is_review` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否审核',
  `create_time` datetime NOT NULL COMMENT '评论时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_comment_user`(`user_id`) USING BTREE,
  INDEX `fk_comment_parent`(`parent_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES (1, 1, 1, '<h1 style=\"color=#FF0000\">评论测试\n<h2>评论测试\n<h3>评论测试\n<img src= \'https://static.talkxj.com/emoji/zhichi.jpg\' width=\'24\'height=\'24\' style=\'margin: 0 1px;vertical-align: text-bottom\'/><img src= \'https://static.talkxj.com/emoji/zhichi.jpg\' width=\'24\'height=\'24\' style=\'margin: 0 1px;vertical-align: text-bottom\'/><img src= \'https://static.talkxj.com/emoji/zhichi.jpg\' width=\'24\'height=\'24\' style=\'margin: 0 1px;vertical-align: text-bottom\'/>', NULL, NULL, 1, 0, 1, '2022-10-14 11:55:23', NULL);

-- ----------------------------
-- Table structure for friend_link
-- ----------------------------
DROP TABLE IF EXISTS `friend_link`;
CREATE TABLE `friend_link`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `link_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '链接名',
  `link_avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '链接头像',
  `link_address` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '链接地址',
  `link_intro` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '链接介绍',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_friend_link_user`(`link_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of friend_link
-- ----------------------------
INSERT INTO `friend_link` VALUES (1, '风丶宇的个人博客', 'https://static.talkxj.com/photos/b553f564f81a80dc338695acb1b475d2.jpg', 'https://www.talkxj.com', '往事不随风', '2022-01-18 00:26:46', NULL);

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '菜单名',
  `path` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '菜单路径',
  `component` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '组件',
  `icon` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '菜单icon',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `order_num` tinyint(1) NOT NULL COMMENT '排序',
  `parent_id` int(11) NULL DEFAULT NULL COMMENT '父id',
  `is_hidden` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否隐藏  0否1是',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 220 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES (1, '首页', '/', '/home/Home.vue', 'el-icon-myshouye', '2021-01-26 17:06:51', '2021-01-26 17:06:53', 1, NULL, 0);
INSERT INTO `menu` VALUES (2, '文章管理', '/article-submenu', 'Layout', 'el-icon-mywenzhang-copy', '2021-01-25 20:43:07', '2021-01-25 20:43:09', 2, NULL, 0);
INSERT INTO `menu` VALUES (3, '消息管理', '/message-submenu', 'Layout', 'el-icon-myxiaoxi', '2021-01-25 20:44:17', '2021-01-25 20:44:20', 3, NULL, 0);
INSERT INTO `menu` VALUES (4, '系统管理', '/system-submenu', 'Layout', 'el-icon-myshezhi', '2021-01-25 20:45:57', '2022-10-10 17:08:10', 9, NULL, 0);
INSERT INTO `menu` VALUES (5, '个人中心', '/setting', '/setting/Setting.vue', 'el-icon-myuser', '2021-01-26 17:22:38', '2022-10-10 17:06:45', 10, NULL, 0);
INSERT INTO `menu` VALUES (6, '发布文章', '/articles', '/article/Article.vue', 'el-icon-myfabiaowenzhang', '2021-01-26 14:30:48', '2021-01-26 14:30:51', 1, 2, 0);
INSERT INTO `menu` VALUES (7, '修改文章', '/articles/*', '/article/Article.vue', 'el-icon-myfabiaowenzhang', '2021-01-26 14:31:32', '2021-01-26 14:31:34', 2, 2, 1);
INSERT INTO `menu` VALUES (8, '文章列表', '/article-list', '/article/ArticleList.vue', 'el-icon-mywenzhangliebiao', '2021-01-26 14:32:13', '2021-01-26 14:32:16', 3, 2, 0);
INSERT INTO `menu` VALUES (9, '分类管理', '/categories', '/category/Category.vue', 'el-icon-myfenlei', '2021-01-26 14:33:42', '2021-01-26 14:33:43', 4, 2, 0);
INSERT INTO `menu` VALUES (10, '标签管理', '/tags', '/tag/Tag.vue', 'el-icon-myicontag', '2021-01-26 14:34:33', '2021-01-26 14:34:36', 5, 2, 0);
INSERT INTO `menu` VALUES (11, '评论管理', '/comments', '/comment/Comment.vue', 'el-icon-mypinglunzu', '2021-01-26 14:35:31', '2021-01-26 14:35:34', 1, 3, 0);
INSERT INTO `menu` VALUES (12, '留言管理', '/messages', '/message/Message.vue', 'el-icon-myliuyan', '2021-01-26 14:36:09', '2021-01-26 14:36:13', 2, 3, 0);
INSERT INTO `menu` VALUES (13, '用户列表', '/users', '/user/User.vue', 'el-icon-myyonghuliebiao', '2021-01-26 14:38:09', '2021-01-26 14:38:12', 1, 202, 0);
INSERT INTO `menu` VALUES (14, '角色管理', '/roles', '/role/Role.vue', 'el-icon-myjiaoseliebiao', '2021-01-26 14:39:01', '2022-10-10 17:17:30', 3, 213, 0);
INSERT INTO `menu` VALUES (15, '接口管理', '/resources', '/resource/Resource.vue', 'el-icon-myjiekouguanli', '2021-01-26 14:40:14', '2021-08-07 20:00:28', 2, 213, 0);
INSERT INTO `menu` VALUES (16, '菜单管理', '/menus', '/menu/Menu.vue', 'el-icon-mycaidan', '2021-01-26 14:40:54', '2022-10-10 17:17:20', 1, 213, 0);
INSERT INTO `menu` VALUES (17, '友链管理', '/links', '/friendLink/FriendLink.vue', 'el-icon-mydashujukeshihuaico-', '2021-01-26 14:41:35', '2021-01-26 14:41:37', 3, 4, 0);
INSERT INTO `menu` VALUES (18, '关于我', '/about', '/about/About.vue', 'el-icon-myguanyuwo', '2021-01-26 14:42:05', '2021-01-26 14:42:10', 4, 4, 0);
INSERT INTO `menu` VALUES (19, '日志管理', '/log-submenu', 'Layout', 'el-icon-myguanyuwo', '2021-01-31 21:33:56', '2022-10-12 21:53:39', 8, NULL, 0);
INSERT INTO `menu` VALUES (20, '操作日志', '/operation/log', '/log/Operation.vue', 'el-icon-myguanyuwo', '2021-01-31 15:53:21', '2021-01-31 15:53:25', 1, 19, 0);
INSERT INTO `menu` VALUES (201, '在线用户', '/online/users', '/user/Online.vue', 'el-icon-myyonghuliebiao', '2021-02-05 14:59:51', '2022-10-12 15:24:10', 7, 202, 0);
INSERT INTO `menu` VALUES (202, '用户管理', '/users-submenu', 'Layout', 'el-icon-myyonghuliebiao', '2021-02-06 23:44:59', '2022-10-10 17:08:33', 6, NULL, 0);
INSERT INTO `menu` VALUES (205, '相册管理', '/album-submenu', 'Layout', 'el-icon-myimage-fill', '2021-08-03 15:10:54', '2022-10-10 17:08:24', 4, NULL, 0);
INSERT INTO `menu` VALUES (206, '相册列表', '/albums', '/album/Album.vue', 'el-icon-myzhaopian', '2021-08-03 20:29:19', '2021-08-04 11:45:47', 1, 205, 0);
INSERT INTO `menu` VALUES (208, '照片管理', '/albums/:albumId', '/album/Photo.vue', 'el-icon-myzhaopian', '2021-08-03 21:37:47', '2021-08-05 10:24:08', 1, 205, 1);
INSERT INTO `menu` VALUES (209, '页面管理', '/pages', '/page/Page.vue', 'el-icon-myyemianpeizhi', '2021-08-04 11:36:27', '2021-08-07 20:01:26', 2, 4, 0);
INSERT INTO `menu` VALUES (210, '照片回收站', '/photos/delete', '/album/Delete.vue', 'el-icon-myhuishouzhan', '2021-08-05 13:55:19', NULL, 3, 205, 1);
INSERT INTO `menu` VALUES (213, '权限管理', '/permission-submenu', 'Layout', 'el-icon-mydaohanglantubiao_quanxianguanli', '2021-08-07 19:56:55', '2022-10-10 17:07:56', 7, NULL, 0);
INSERT INTO `menu` VALUES (214, '网站管理', '/website', '/website/Website.vue', 'el-icon-myxitong', '2021-08-07 20:06:41', NULL, 1, 4, 0);
INSERT INTO `menu` VALUES (215, '说说管理', '/talk-submenu', 'Layout', 'el-icon-mypinglun', '2022-01-23 20:17:59', '2022-01-23 20:38:06', 5, NULL, 0);
INSERT INTO `menu` VALUES (216, '发布说说', '/talks', '/talk/Talk.vue', 'el-icon-myfabusekuai', '2022-01-23 20:18:43', '2022-10-11 13:16:22', 2, 215, 0);
INSERT INTO `menu` VALUES (217, '说说列表', '/talk-list', '/talk/TalkList.vue', 'el-icon-myiconfontdongtaidianji', '2022-01-23 23:28:24', '2022-10-11 13:16:18', 1, 215, 0);
INSERT INTO `menu` VALUES (218, '修改说说', '/talks/:talkId', '/talk/Talk.vue', 'el-icon-myshouye', '2022-01-24 00:10:44', NULL, 3, 215, 1);

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `nickname` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '昵称',
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '头像',
  `message_content` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '留言内容',
  `ip_address` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户ip',
  `ip_source` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户地址',
  `time` tinyint(1) NULL DEFAULT NULL COMMENT '弹幕速度',
  `is_review` tinyint(4) NOT NULL DEFAULT 1 COMMENT '是否审核',
  `create_time` datetime NOT NULL COMMENT '发布时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of message
-- ----------------------------
INSERT INTO `message` VALUES (1, '水平线之下', 'https://blog-1311853727.cos.ap-guangzhou.myqcloud.com/user/admin.jpg', '留言测试', '110.52.119.165', '中国|0|湖南省|湘潭市|联通', 8, 1, '2022-10-14 12:04:05', '2022-10-14 12:04:05');

-- ----------------------------
-- Table structure for operation_log
-- ----------------------------
DROP TABLE IF EXISTS `operation_log`;
CREATE TABLE `operation_log`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `opt_module` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '操作模块',
  `opt_type` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '操作类型',
  `opt_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '操作url',
  `opt_method` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '操作方法',
  `opt_desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '操作描述',
  `request_param` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '请求参数',
  `request_method` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '请求方式',
  `response_data` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '返回数据',
  `user_id` int(11) NOT NULL COMMENT '用户id',
  `nickname` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户昵称',
  `ip_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '操作ip',
  `ip_source` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '操作地址',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 36 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of operation_log
-- ----------------------------
INSERT INTO `operation_log` VALUES (9, '用户', '删除', '/api/admin/users/1006/online', 'myblog/api/v1.User.ForceOffline-fm', '删除', '{\"ipAddress\":\"110.53.129.156\",\"browser\":\"Chrome 92.0.4515\",\"os\":\"Windows 10.0.0\"}', 'DELETE', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 13:14:48', '2022-10-13 13:14:48');
INSERT INTO `operation_log` VALUES (10, '操作日志', '删除', '/api/admin/operation/logs', 'myblog/api/v1.OperationLog.Delete-fm', '删除', '[8,7]', 'DELETE', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 13:16:38', '2022-10-13 13:16:38');
INSERT INTO `operation_log` VALUES (11, '说说', '新增或修改', '/api/admin/talks', 'myblog/api/v1.Talk.SaveOrUpdate-fm', '新增或修改', '{\"id\":52,\"nickname\":\"水平线之下\",\"avatar\":\"https://blog-1311853727.cos.ap-guangzhou.myqcloud.com/user/admin.jpg\",\"content\":\"rrr\",\"images\":\"[\\\"https://blog-1311853727.cos.ap-guangzhou.myqcloud.com/talk/anime.jpg\\\"]\",\"imgList\":[\"https://blog-1311853727.cos.ap-guangzhou.myqcloud.com/talk/anime.jpg\"],\"isTop\":0,\"status\":1,\"createTime\":\"2022-10-11T14:20:12+08:00\"}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 13:39:00', '2022-10-13 13:39:00');
INSERT INTO `operation_log` VALUES (12, '说说', '新增或修改', '/api/admin/talks', 'myblog/api/v1.Talk.SaveOrUpdate-fm', '新增或修改', '{\"id\":49,\"nickname\":\"水平线之下\",\"avatar\":\"https://blog-1311853727.cos.ap-guangzhou.myqcloud.com/user/admin.jpg\",\"content\":\"ttt<img src=\\\"https://static.talkxj.com/emoji/smile.jpg\\\" width=\\\"24\\\" height=\\\"24\\\" alt=\\\"[微笑]\\\" style=\\\"margin: 0 1px;vertical-align: text-bottom\\\"><img src=\\\" https://static.talkxj.com/emoji/dx.jpg\\\" width=\\\"24\\\" height=\\\"24\\\" alt=\\\"[笑]\\\" style=\\\"margin: 0 1px;vertical-align: text-bottom\\\"><img src=\\\"https://static.talkxj.com/emoji/cy.jpg\\\" width=\\\"24\\\" height=\\\"24\\\" alt=\\\"[呲牙]\\\" style=\\\"margin: 0 1px;vertical-align: text-bottom\\\"><img src=\\\"https://static.talkxj.com/emoji/ok.jpg\\\" width=\\\"24\\\" height=\\\"24\\\" alt=\\\"[OK]\\\" style=\\\"margin: 0 1px;vertical-align: text-bottom\\\">\",\"images\":\"\",\"imgList\":null,\"isTop\":0,\"status\":1,\"createTime\":\"2022-01-24T23:34:59+08:00\"}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 13:40:04', '2022-10-13 13:40:04');
INSERT INTO `operation_log` VALUES (13, '说说', '新增或修改', '/api/admin/talks', 'myblog/api/v1.Talk.SaveOrUpdate-fm', '新增或修改', '{\"id\":null,\"content\":\"test\",\"isTop\":0,\"status\":1,\"images\":null}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 13:40:22', '2022-10-13 13:40:22');
INSERT INTO `operation_log` VALUES (14, '资源权限', '新增或修改', '/api/admin/resources', 'myblog/api/v1.Resource.SaveOrUpdate-fm', '新增或修改', '{\"id\":282,\"resourceName\":\"点赞说说\",\"url\":\"/talks/:talkId/like\",\"requestMethod\":\"POST\",\"isAnonymous\":0,\"createTime\":\"2022-01-24T01:30:30+08:00\",\"children\":null}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 13:59:59', '2022-10-13 13:59:59');
INSERT INTO `operation_log` VALUES (15, '角色', '新增或修改', '/api/admin/role', 'myblog/api/v1.Role.SaveOrUpdate-fm', '新增或修改', '{\"id\":1,\"roleName\":\"管理员\",\"roleLabel\":\"admin\",\"createTime\":\"2021-03-22T14:10:21+08:00\",\"isDisable\":0,\"resourceIdList\":[165,192,193,194,195,166,183,184,246,247,167,199,200,201,168,185,186,187,188,189,190,191,254,169,208,209,170,234,235,236,237,171,213,214,215,216,217,224,172,240,241,244,245,267,269,270,173,239,242,276,174,205,206,207,175,218,219,220,221,222,223,176,202,203,204,230,238,177,229,232,233,243,178,196,197,198,257,258,179,225,226,227,228,231,180,210,211,212,278,282,283,284,285,286,287],\"menuIdList\":null}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 14:00:29', '2022-10-13 14:00:29');
INSERT INTO `operation_log` VALUES (16, '角色', '新增或修改', '/api/admin/role', 'myblog/api/v1.Role.SaveOrUpdate-fm', '新增或修改', '{\"id\":3,\"roleName\":\"测试\",\"roleLabel\":\"test\",\"createTime\":\"2021-03-22T14:42:23+08:00\",\"isDisable\":\"0\",\"resourceIdList\":[192,195,183,246,199,185,191,254,208,234,237,213,241,239,276,205,218,221,223,202,230,238,232,243,196,257,258,225,231,210,282,286,287],\"menuIdList\":null}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 14:00:40', '2022-10-13 14:00:40');
INSERT INTO `operation_log` VALUES (17, '资源权限', '新增或修改', '/api/admin/resources', 'myblog/api/v1.Resource.SaveOrUpdate-fm', '新增或修改', '{\"id\":248,\"resourceName\":\"根据相册id查看照片列表\",\"url\":\"/albums/:albumId/photos\",\"requestMethod\":\"GET\",\"isAnonymous\":1,\"createTime\":\"2021-08-11T21:04:22+08:00\",\"children\":null}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 15:21:46', '2022-10-13 15:21:46');
INSERT INTO `operation_log` VALUES (18, '角色', '新增或修改', '/api/admin/role', 'myblog/api/v1.Role.SaveOrUpdate-fm', '新增或修改', '{\"id\":3,\"roleName\":\"测试\",\"roleLabel\":\"test\",\"createTime\":\"2021-03-22T14:42:23+08:00\",\"isDisable\":0,\"resourceIdList\":[192,195,183,246,199,185,191,254,208,234,237,213,241,239,276,205,218,221,223,202,230,238,232,243,196,257,258,225,231,210,282,286,287],\"menuIdList\":null}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 15:21:55', '2022-10-13 15:21:55');
INSERT INTO `operation_log` VALUES (19, '角色', '新增或修改', '/api/admin/role', 'myblog/api/v1.Role.SaveOrUpdate-fm', '新增或修改', '{\"id\":1,\"roleName\":\"管理员\",\"roleLabel\":\"admin\",\"createTime\":\"2021-03-22T14:10:21+08:00\",\"isDisable\":\"0\",\"resourceIdList\":[165,192,193,194,195,166,183,184,246,247,167,199,200,201,168,185,186,187,188,189,190,191,254,169,208,209,170,234,235,236,237,171,213,214,215,216,217,224,172,240,241,244,245,267,269,270,173,239,242,276,174,205,206,207,175,218,219,220,221,222,223,176,202,203,204,230,238,177,229,232,233,243,178,196,197,198,257,258,179,225,226,227,228,231,180,210,211,212,278,282,283,284,285,286,287],\"menuIdList\":null}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 15:21:57', '2022-10-13 15:21:57');
INSERT INTO `operation_log` VALUES (20, '角色', '新增或修改', '/api/admin/role', 'myblog/api/v1.Role.SaveOrUpdate-fm', '新增或修改', '{\"id\":2,\"roleName\":\"用户\",\"roleLabel\":\"user\",\"createTime\":\"2021-03-22T14:25:25+08:00\",\"isDisable\":\"0\",\"resourceIdList\":[],\"menuIdList\":null}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 15:22:00', '2022-10-13 15:22:00');
INSERT INTO `operation_log` VALUES (21, '资源权限', '新增或修改', '/api/admin/resources', 'myblog/api/v1.Resource.SaveOrUpdate-fm', '新增或修改', '{\"id\":254,\"resourceName\":\"点赞文章\",\"url\":\"/articles/:articleId/like\",\"requestMethod\":\"POST\",\"isAnonymous\":0,\"createTime\":\"2021-08-11T21:04:22+08:00\",\"children\":null}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 17:22:45', '2022-10-13 17:22:45');
INSERT INTO `operation_log` VALUES (22, '角色', '新增或修改', '/api/admin/role', 'myblog/api/v1.Role.SaveOrUpdate-fm', '新增或修改', '{\"id\":1,\"roleName\":\"管理员\",\"roleLabel\":\"admin\",\"createTime\":\"2021-03-22T14:10:21+08:00\",\"isDisable\":0,\"resourceIdList\":[165,192,193,194,195,166,183,184,246,247,167,199,200,201,168,185,186,187,188,189,190,191,254,169,208,209,170,234,235,236,237,171,213,214,215,216,217,224,172,240,241,244,245,267,269,270,173,239,242,276,174,205,206,207,175,218,219,220,221,222,223,176,202,203,204,230,238,177,229,232,233,243,178,196,197,198,257,258,179,225,226,227,228,231,180,210,211,212,278,282,283,284,285,286,287],\"menuIdList\":null}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 17:22:49', '2022-10-13 17:22:49');
INSERT INTO `operation_log` VALUES (23, '角色', '新增或修改', '/api/admin/role', 'myblog/api/v1.Role.SaveOrUpdate-fm', '新增或修改', '{\"id\":2,\"roleName\":\"用户\",\"roleLabel\":\"user\",\"createTime\":\"2021-03-22T14:25:25+08:00\",\"isDisable\":\"0\",\"resourceIdList\":[254],\"menuIdList\":null}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 17:22:59', '2022-10-13 17:22:59');
INSERT INTO `operation_log` VALUES (24, '角色', '新增或修改', '/api/admin/role', 'myblog/api/v1.Role.SaveOrUpdate-fm', '新增或修改', '{\"id\":3,\"roleName\":\"测试\",\"roleLabel\":\"test\",\"createTime\":\"2021-03-22T14:42:23+08:00\",\"isDisable\":\"0\",\"resourceIdList\":[192,195,183,246,199,185,191,254,208,234,237,213,241,239,276,205,218,221,223,202,230,238,232,243,196,257,258,225,231,210,282,286,287],\"menuIdList\":null}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 17:23:03', '2022-10-13 17:23:03');
INSERT INTO `operation_log` VALUES (25, '资源权限', '新增或修改', '/api/admin/resources', 'myblog/api/v1.Resource.SaveOrUpdate-fm', '新增或修改', '{\"id\":258,\"resourceName\":\"评论点赞\",\"url\":\"/comments/:commentId/like\",\"requestMethod\":\"POST\",\"isAnonymous\":0,\"createTime\":\"2021-08-11T21:04:22+08:00\",\"children\":null}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 21:31:24', '2022-10-13 21:31:24');
INSERT INTO `operation_log` VALUES (26, '资源权限', '新增或修改', '/api/admin/resources', 'myblog/api/v1.Resource.SaveOrUpdate-fm', '新增或修改', '{\"id\":259,\"resourceName\":\"查询评论下的回复\",\"url\":\"/comments/:commentId/replies\",\"requestMethod\":\"GET\",\"isAnonymous\":1,\"createTime\":\"2021-08-11T21:04:22+08:00\",\"children\":null}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 21:31:31', '2022-10-13 21:31:31');
INSERT INTO `operation_log` VALUES (27, '角色', '新增或修改', '/api/admin/role', 'myblog/api/v1.Role.SaveOrUpdate-fm', '新增或修改', '{\"id\":1,\"roleName\":\"管理员\",\"roleLabel\":\"admin\",\"createTime\":\"2021-03-22T14:10:21+08:00\",\"isDisable\":0,\"resourceIdList\":[165,192,193,194,195,166,183,184,246,247,167,199,200,201,168,185,186,187,188,189,190,191,254,169,208,209,170,234,235,236,237,171,213,214,215,216,217,224,172,240,241,244,245,267,269,270,173,239,242,276,174,205,206,207,175,218,219,220,221,222,223,176,202,203,204,230,238,177,229,232,233,243,178,196,197,198,257,258,179,225,226,227,228,231,180,210,211,212,278,282,283,284,285,286,287],\"menuIdList\":null}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 21:31:37', '2022-10-13 21:31:37');
INSERT INTO `operation_log` VALUES (28, '角色', '新增或修改', '/api/admin/role', 'myblog/api/v1.Role.SaveOrUpdate-fm', '新增或修改', '{\"id\":2,\"roleName\":\"用户\",\"roleLabel\":\"user\",\"createTime\":\"2021-03-22T14:25:25+08:00\",\"isDisable\":\"0\",\"resourceIdList\":[254,267,269,270,257,258,282],\"menuIdList\":null}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 21:32:52', '2022-10-13 21:32:52');
INSERT INTO `operation_log` VALUES (29, '角色', '新增或修改', '/api/admin/role', 'myblog/api/v1.Role.SaveOrUpdate-fm', '新增或修改', '{\"id\":3,\"roleName\":\"测试\",\"roleLabel\":\"test\",\"createTime\":\"2021-03-22T14:42:23+08:00\",\"isDisable\":\"0\",\"resourceIdList\":[192,195,183,246,199,185,191,254,208,234,237,213,241,239,276,205,218,221,223,202,230,238,232,243,196,257,258,225,231,210,282,286,287],\"menuIdList\":null}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 21:32:59', '2022-10-13 21:32:59');
INSERT INTO `operation_log` VALUES (30, '评论', '修改', '/api/admin/comments/review', 'myblog/api/v1.Comment.UpdateReview-fm', '修改', '{\"idList\":[756],\"isReview\":1}', 'PUT', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 21:51:27', '2022-10-13 21:51:27');
INSERT INTO `operation_log` VALUES (31, '评论', '删除', '/api/admin/comments', 'myblog/api/v1.Comment.Delete-fm', '删除', '[766]', 'DELETE', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-13 22:13:42', '2022-10-13 22:13:42');
INSERT INTO `operation_log` VALUES (32, '资源权限', '新增或修改', '/api/admin/resources', 'myblog/api/v1.Resource.SaveOrUpdate-fm', '新增或修改', '{\"id\":267,\"resourceName\":\"更新用户头像\",\"url\":\"/users/avatar\",\"requestMethod\":\"POST\",\"isAnonymous\":1,\"createTime\":\"2021-08-11T21:04:22+08:00\",\"children\":null}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-14 09:59:06', '2022-10-14 09:59:06');
INSERT INTO `operation_log` VALUES (33, '标签', '新增或修改', '/api/admin/tags', 'myblog/api/v1.Tag.SaveOrUpdate-fm', '新增或修改', '{\"id\":null,\"tagName\":\"测试标签\"}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-14 11:49:37', '2022-10-14 11:49:37');
INSERT INTO `operation_log` VALUES (34, '分类', '新增或修改', '/api/admin/categories', 'myblog/api/v1.Category.SaveOrUpdate-fm', '新增或修改', '{\"id\":null,\"categoryName\":\"测试分类\"}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-14 11:49:46', '2022-10-14 11:49:46');
INSERT INTO `operation_log` VALUES (35, '文章', '新增或修改', '/api/admin/articles', 'myblog/api/v1.Article.SaveOrUpdate-fm', '新增或修改', '{\"id\":null,\"articleTitle\":\"测试文章\",\"articleContent\":\"# 测试文章内容\\n## 测试文章内容\\n### 测试文章内容\\n#### 测试文章内容\\n##### 测试文章内容\\n###### 测试文章内容\\n![anime.jpg](https://blog-1311853727.cos.ap-guangzhou.myqcloud.com/article/anime.jpg)\",\"articleCover\":\"https://blog-1311853727.cos.ap-guangzhou.myqcloud.com/article/anime.jpg\",\"categoryName\":\"测试分类\",\"tagNameList\":[\"测试标签\"],\"originalUrl\":\"\",\"isTop\":0,\"type\":1,\"status\":1}', 'POST', '{\"code\":20000,\"flag\":true,\"message\":\"操作成功\"}', 1, '水平线之下', '119.39.148.45', '湖南省湘潭市 联通', '2022-10-14 11:51:09', '2022-10-14 11:51:09');

-- ----------------------------
-- Table structure for page
-- ----------------------------
DROP TABLE IF EXISTS `page`;
CREATE TABLE `page`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '页面id',
  `page_name` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '页面名',
  `page_label` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '页面标签',
  `page_cover` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '页面封面',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 906 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '页面' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of page
-- ----------------------------
INSERT INTO `page` VALUES (1, '首页', 'home', 'https://static.talkxj.com/config/0bee7ba5ac70155766648e14ae2a821f.jpg', '2021-08-07 10:32:36', '2021-12-27 12:19:01');
INSERT INTO `page` VALUES (2, '归档', 'archive', 'https://static.talkxj.com/config/643f28683e1c59a80ccfc9cb19735a9c.jpg', '2021-08-07 10:32:36', '2021-10-04 15:43:14');
INSERT INTO `page` VALUES (3, '分类', 'category', 'https://static.talkxj.com/config/83be0017d7f1a29441e33083e7706936.jpg', '2021-08-07 10:32:36', '2021-10-04 15:43:31');
INSERT INTO `page` VALUES (4, '标签', 'tag', 'https://static.talkxj.com/config/a6f141372509365891081d755da963a1.png', '2021-08-07 10:32:36', '2021-10-04 15:43:38');
INSERT INTO `page` VALUES (5, '相册', 'album', 'https://blog-1311853727.cos.ap-guangzhou.myqcloud.com/user/anime.jpg', '2021-08-07 10:32:36', '2022-10-11 12:26:14');
INSERT INTO `page` VALUES (6, '友链', 'link', 'https://static.talkxj.com/config/9034edddec5b8e8542c2e61b0da1c1da.jpg', '2021-08-07 10:32:36', '2021-10-04 15:44:02');
INSERT INTO `page` VALUES (7, '关于', 'about', 'https://static.talkxj.com/config/2a56d15dd742ff8ac238a512d9a472a1.jpg', '2021-08-07 10:32:36', '2021-10-04 15:44:08');
INSERT INTO `page` VALUES (8, '留言', 'message', 'https://static.talkxj.com/config/acfeab8379508233fa7e4febf90c2f2e.png', '2021-08-07 10:32:36', '2021-10-04 16:11:45');
INSERT INTO `page` VALUES (9, '个人中心', 'user', 'https://static.talkxj.com/config/ebae4c93de1b286a8d50aa62612caa59.jpeg', '2021-08-07 10:32:36', '2021-10-04 15:45:17');
INSERT INTO `page` VALUES (10, '文章列表', 'articleList', 'https://static.talkxj.com/config/924d65cc8312e6cdad2160eb8fce6831.jpg', '2021-08-10 15:36:19', '2021-10-04 15:45:45');
INSERT INTO `page` VALUES (904, '说说', 'talk', 'https://static.talkxj.com/config/a741b0656a9a3db2e2ba5c2f4140eb6c.jpg', '2022-01-23 00:51:24', '2022-01-23 03:01:21');

-- ----------------------------
-- Table structure for photo
-- ----------------------------
DROP TABLE IF EXISTS `photo`;
CREATE TABLE `photo`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `album_id` int(11) NOT NULL COMMENT '相册id',
  `photo_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '照片名',
  `photo_desc` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '照片描述',
  `photo_src` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '照片地址',
  `is_delete` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '照片' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of photo
-- ----------------------------

-- ----------------------------
-- Table structure for photo_album
-- ----------------------------
DROP TABLE IF EXISTS `photo_album`;
CREATE TABLE `photo_album`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `album_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '相册名',
  `album_desc` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '相册描述',
  `album_cover` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '相册封面',
  `is_delete` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态值 1公开 2私密',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '相册' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of photo_album
-- ----------------------------

-- ----------------------------
-- Table structure for resource
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `resource_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '资源名',
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '权限路径',
  `request_method` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '请求方式',
  `parent_id` int(11) NULL DEFAULT NULL COMMENT '父权限id',
  `is_anonymous` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否匿名访问 0否 1是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 288 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of resource
-- ----------------------------
INSERT INTO `resource` VALUES (165, '分类模块', NULL, NULL, NULL, 0, '2021-08-11 21:04:21', NULL);
INSERT INTO `resource` VALUES (166, '博客信息模块', NULL, NULL, NULL, 0, '2021-08-11 21:04:21', NULL);
INSERT INTO `resource` VALUES (167, '友链模块', NULL, NULL, NULL, 0, '2021-08-11 21:04:21', NULL);
INSERT INTO `resource` VALUES (168, '文章模块', NULL, NULL, NULL, 0, '2021-08-11 21:04:21', NULL);
INSERT INTO `resource` VALUES (169, '日志模块', NULL, NULL, NULL, 0, '2021-08-11 21:04:21', NULL);
INSERT INTO `resource` VALUES (170, '标签模块', NULL, NULL, NULL, 0, '2021-08-11 21:04:21', NULL);
INSERT INTO `resource` VALUES (171, '照片模块', NULL, NULL, NULL, 0, '2021-08-11 21:04:21', NULL);
INSERT INTO `resource` VALUES (172, '用户信息模块', NULL, NULL, NULL, 0, '2021-08-11 21:04:21', NULL);
INSERT INTO `resource` VALUES (173, '用户账号模块', NULL, NULL, NULL, 0, '2021-08-11 21:04:21', NULL);
INSERT INTO `resource` VALUES (174, '留言模块', NULL, NULL, NULL, 0, '2021-08-11 21:04:21', NULL);
INSERT INTO `resource` VALUES (175, '相册模块', NULL, NULL, NULL, 0, '2021-08-11 21:04:21', NULL);
INSERT INTO `resource` VALUES (176, '菜单模块', NULL, NULL, NULL, 0, '2021-08-11 21:04:21', NULL);
INSERT INTO `resource` VALUES (177, '角色模块', NULL, NULL, NULL, 0, '2021-08-11 21:04:21', NULL);
INSERT INTO `resource` VALUES (178, '评论模块', NULL, NULL, NULL, 0, '2021-08-11 21:04:21', NULL);
INSERT INTO `resource` VALUES (179, '资源模块', NULL, NULL, NULL, 0, '2021-08-11 21:04:21', NULL);
INSERT INTO `resource` VALUES (180, '页面模块', NULL, NULL, NULL, 0, '2021-08-11 21:04:21', NULL);
INSERT INTO `resource` VALUES (181, '查看博客信息', '/', 'GET', 166, 1, '2021-08-11 21:04:22', '2022-10-11 09:36:12');
INSERT INTO `resource` VALUES (182, '查看关于我信息', '/about', 'GET', 166, 1, '2021-08-11 21:04:22', '2022-10-11 09:36:14');
INSERT INTO `resource` VALUES (183, '查看后台信息', '/admin', 'GET', 166, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (184, '修改关于我信息', '/admin/about', 'PUT', 166, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (185, '查看后台文章', '/admin/articles', 'GET', 168, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (186, '添加或修改文章', '/admin/articles', 'POST', 168, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (187, '恢复或删除文章', '/admin/articles', 'PUT', 168, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (188, '物理删除文章', '/admin/articles', 'DELETE', 168, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (189, '上传文章图片', '/admin/articles/images', 'POST', 168, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (190, '修改文章置顶', '/admin/articles/top', 'PUT', 168, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (191, '根据id查看后台文章', '/admin/articles/:id', 'GET', 168, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (192, '查看后台分类列表', '/admin/categories', 'GET', 165, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (193, '添加或修改分类', '/admin/categories', 'POST', 165, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (194, '删除分类', '/admin/categories', 'DELETE', 165, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (195, '搜索文章分类', '/admin/categories/search', 'GET', 165, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (196, '查询后台评论', '/admin/comments', 'GET', 178, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (197, '删除评论', '/admin/comments', 'DELETE', 178, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (198, '审核评论', '/admin/comments/review', 'PUT', 178, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (199, '查看后台友链列表', '/admin/links', 'GET', 167, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (200, '保存或修改友链', '/admin/links', 'POST', 167, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (201, '删除友链', '/admin/links', 'DELETE', 167, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (202, '查看菜单列表', '/admin/menus', 'GET', 176, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (203, '新增或修改菜单', '/admin/menus', 'POST', 176, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (204, '删除菜单', '/admin/menus/:menuId', 'DELETE', 176, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (205, '查看后台留言列表', '/admin/messages', 'GET', 174, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (206, '删除留言', '/admin/messages', 'DELETE', 174, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (207, '审核留言', '/admin/messages/review', 'PUT', 174, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (208, '查看操作日志', '/admin/operation/logs', 'GET', 169, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (209, '删除操作日志', '/admin/operation/logs', 'DELETE', 169, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (210, '获取页面列表', '/admin/pages', 'GET', 180, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (211, '保存或更新页面', '/admin/pages', 'POST', 180, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (212, '删除页面', '/admin/pages/:pageId', 'DELETE', 180, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (213, '根据相册id获取照片列表', '/admin/photos', 'GET', 171, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (214, '保存照片', '/admin/photos', 'POST', 171, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (215, '更新照片信息', '/admin/photos', 'PUT', 171, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (216, '删除照片', '/admin/photos', 'DELETE', 171, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (217, '移动照片相册', '/admin/photos/album', 'PUT', 171, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (218, '查看后台相册列表', '/admin/photos/albums', 'GET', 175, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (219, '保存或更新相册', '/admin/photos/albums', 'POST', 175, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (220, '上传相册封面', '/admin/photos/albums/cover', 'POST', 175, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (221, '获取后台相册列表信息', '/admin/photos/albums/info', 'GET', 175, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (222, '根据id删除相册', '/admin/photos/albums/:albumId', 'DELETE', 175, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (223, '根据id获取后台相册信息', '/admin/photos/albums/:albumId/info', 'GET', 175, 0, '2021-08-11 21:04:22', '2022-10-11 16:35:48');
INSERT INTO `resource` VALUES (224, '更新照片删除状态', '/admin/photos/delete', 'PUT', 171, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (225, '查看资源列表', '/admin/resources', 'GET', 179, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (226, '新增或修改资源', '/admin/resources', 'POST', 179, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (227, '导入swagger接口', '/admin/resources/import/swagger', 'GET', 179, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (228, '删除资源', '/admin/resources/:resourceId', 'DELETE', 179, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (229, '保存或更新角色', '/admin/role', 'POST', 177, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (230, '查看角色菜单选项', '/admin/role/menus', 'GET', 176, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (231, '查看角色资源选项', '/admin/role/resources', 'GET', 179, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (232, '查询角色列表', '/admin/roles', 'GET', 177, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (233, '删除角色', '/admin/roles', 'DELETE', 177, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (234, '查询后台标签列表', '/admin/tags', 'GET', 170, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (235, '添加或修改标签', '/admin/tags', 'POST', 170, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (236, '删除标签', '/admin/tags', 'DELETE', 170, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (237, '搜索文章标签', '/admin/tags/search', 'GET', 170, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (238, '查看当前用户菜单', '/admin/user/menus', 'GET', 176, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (239, '查询后台用户列表', '/admin/users', 'GET', 173, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (240, '修改用户禁用状态', '/admin/users/disable', 'PUT', 172, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (241, '查看在线用户', '/admin/users/online', 'GET', 172, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (242, '修改管理员密码', '/admin/users/password', 'PUT', 173, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (243, '查询用户角色选项', '/admin/users/role', 'GET', 177, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (244, '修改用户角色', '/admin/users/role', 'PUT', 172, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (245, '下线用户', '/admin/users/:userInfoId/online', 'DELETE', 172, 0, '2021-08-11 21:04:22', '2022-10-12 15:44:29');
INSERT INTO `resource` VALUES (246, '获取网站配置', '/admin/website/config', 'GET', 166, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (247, '更新网站配置', '/admin/website/config', 'PUT', 166, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (248, '根据相册id查看照片列表', '/albums/:albumId/photos', 'GET', 171, 1, '2021-08-11 21:04:22', '2022-10-13 15:21:46');
INSERT INTO `resource` VALUES (249, '查看首页文章', '/articles', 'GET', 168, 1, '2021-08-11 21:04:22', '2022-10-11 09:36:33');
INSERT INTO `resource` VALUES (250, '查看文章归档', '/articles/archives', 'GET', 168, 1, '2021-08-11 21:04:22', '2022-10-11 09:36:34');
INSERT INTO `resource` VALUES (251, '根据条件查询文章', '/articles/condition', 'GET', 168, 1, '2021-08-11 21:04:22', '2022-10-11 09:36:35');
INSERT INTO `resource` VALUES (252, '搜索文章', '/articles/search', 'GET', 168, 1, '2021-08-11 21:04:22', '2022-10-11 09:36:36');
INSERT INTO `resource` VALUES (253, '根据id查看文章', '/articles/:articleId', 'GET', 168, 1, '2021-08-11 21:04:22', '2022-10-11 09:36:37');
INSERT INTO `resource` VALUES (254, '点赞文章', '/articles/:articleId/like', 'POST', 168, 0, '2021-08-11 21:04:22', '2022-10-13 17:22:45');
INSERT INTO `resource` VALUES (255, '查看分类列表', '/categories', 'GET', 165, 1, '2021-08-11 21:04:22', '2022-10-11 09:35:51');
INSERT INTO `resource` VALUES (256, '查询评论', '/comments', 'GET', 178, 1, '2021-08-11 21:04:22', '2022-10-11 09:37:28');
INSERT INTO `resource` VALUES (257, '添加评论', '/comments', 'POST', 178, 0, '2021-08-11 21:04:22', '2021-08-11 21:10:05');
INSERT INTO `resource` VALUES (258, '评论点赞', '/comments/:commentId/like', 'POST', 178, 0, '2021-08-11 21:04:22', '2022-10-13 21:31:24');
INSERT INTO `resource` VALUES (259, '查询评论下的回复', '/comments/:commentId/replies', 'GET', 178, 1, '2021-08-11 21:04:22', '2022-10-13 21:31:31');
INSERT INTO `resource` VALUES (260, '查看友链列表', '/links', 'GET', 167, 1, '2021-08-11 21:04:22', '2022-10-11 09:36:29');
INSERT INTO `resource` VALUES (261, '查看留言列表', '/messages', 'GET', 174, 1, '2021-08-11 21:04:22', '2022-10-11 09:37:13');
INSERT INTO `resource` VALUES (262, '添加留言', '/messages', 'POST', 174, 1, '2021-08-11 21:04:22', '2022-10-11 09:37:14');
INSERT INTO `resource` VALUES (263, '获取相册列表', '/photos/albums', 'GET', 175, 1, '2021-08-11 21:04:22', '2022-10-11 09:37:19');
INSERT INTO `resource` VALUES (264, '用户注册', '/register', 'POST', 173, 1, '2021-08-11 21:04:22', '2022-10-11 09:37:03');
INSERT INTO `resource` VALUES (265, '查询标签列表', '/tags', 'GET', 170, 1, '2021-08-11 21:04:22', '2022-10-11 09:36:51');
INSERT INTO `resource` VALUES (267, '更新用户头像', '/users/avatar', 'POST', 172, 1, '2021-08-11 21:04:22', '2022-10-14 09:59:06');
INSERT INTO `resource` VALUES (268, '发送邮箱验证码', '/users/code', 'GET', 173, 1, '2021-08-11 21:04:22', '2022-10-11 09:37:04');
INSERT INTO `resource` VALUES (269, '绑定用户邮箱', '/users/email', 'POST', 172, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (270, '更新用户信息', '/users/info', 'PUT', 172, 0, '2021-08-11 21:04:22', NULL);
INSERT INTO `resource` VALUES (271, 'qq登录', '/users/oauth/qq', 'POST', 173, 1, '2021-08-11 21:04:22', '2022-10-11 09:37:05');
INSERT INTO `resource` VALUES (272, '微博登录', '/users/oauth/weibo', 'POST', 173, 1, '2021-08-11 21:04:22', '2022-10-11 09:37:07');
INSERT INTO `resource` VALUES (273, '修改密码', '/users/password', 'PUT', 173, 1, '2021-08-11 21:04:22', '2022-10-11 09:37:08');
INSERT INTO `resource` VALUES (274, '上传语音', '/voice', 'POST', 166, 1, '2021-08-11 21:04:22', '2022-10-11 09:36:15');
INSERT INTO `resource` VALUES (275, '上传访客信息', '/report', 'POST', 166, 1, '2021-08-24 00:32:05', '2022-10-11 09:36:16');
INSERT INTO `resource` VALUES (276, '获取用户区域分布', '/admin/users/area', 'GET', 173, 0, '2021-08-24 00:32:35', '2021-09-24 16:25:34');
INSERT INTO `resource` VALUES (278, '说说模块', NULL, NULL, NULL, 0, '2022-01-24 01:29:13', NULL);
INSERT INTO `resource` VALUES (279, '查看首页说说', '/home/talks', 'GET', 278, 1, '2022-01-24 01:29:29', '2022-10-11 09:37:41');
INSERT INTO `resource` VALUES (280, '查看说说列表', '/talks', 'GET', 278, 1, '2022-01-24 01:29:52', '2022-10-11 09:37:42');
INSERT INTO `resource` VALUES (281, '根据id查看说说', '/talks/:talkId', 'GET', 278, 1, '2022-01-24 01:30:10', '2022-10-11 09:37:43');
INSERT INTO `resource` VALUES (282, '点赞说说', '/talks/:talkId/like', 'POST', 278, 0, '2022-01-24 01:30:30', '2022-10-13 13:59:59');
INSERT INTO `resource` VALUES (283, '上传说说图片', '/admin/talks/images', 'POST', 278, 0, '2022-01-24 01:30:46', NULL);
INSERT INTO `resource` VALUES (284, '保存或修改说说', '/admin/talks', 'POST', 278, 0, '2022-01-24 01:31:04', NULL);
INSERT INTO `resource` VALUES (285, '删除说说', '/admin/talks', 'DELETE', 278, 0, '2022-01-24 01:31:22', NULL);
INSERT INTO `resource` VALUES (286, '查看后台说说', '/admin/talks', 'GET', 278, 0, '2022-01-24 01:31:38', NULL);
INSERT INTO `resource` VALUES (287, '根据id查看后台说说', '/admin/talks/:talkId', 'GET', 278, 0, '2022-01-24 01:31:53', '2022-01-24 01:33:14');

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `role_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色名',
  `role_label` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色描述',
  `is_disable` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否禁用  0否 1是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1, '管理员', 'admin', 0, '2021-03-22 14:10:21', '2022-10-13 21:31:37');
INSERT INTO `role` VALUES (2, '用户', 'user', 0, '2021-03-22 14:25:25', '2022-10-13 21:32:52');
INSERT INTO `role` VALUES (3, '测试', 'test', 0, '2021-03-22 14:42:23', '2022-10-13 21:32:59');

-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `role_id` int(11) NULL DEFAULT NULL COMMENT '角色id',
  `menu_id` int(11) NULL DEFAULT NULL COMMENT '菜单id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2587 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of role_menu
-- ----------------------------
INSERT INTO `role_menu` VALUES (2461, 1, 1);
INSERT INTO `role_menu` VALUES (2462, 1, 2);
INSERT INTO `role_menu` VALUES (2463, 1, 6);
INSERT INTO `role_menu` VALUES (2464, 1, 7);
INSERT INTO `role_menu` VALUES (2465, 1, 8);
INSERT INTO `role_menu` VALUES (2466, 1, 9);
INSERT INTO `role_menu` VALUES (2467, 1, 10);
INSERT INTO `role_menu` VALUES (2468, 1, 3);
INSERT INTO `role_menu` VALUES (2469, 1, 11);
INSERT INTO `role_menu` VALUES (2470, 1, 12);
INSERT INTO `role_menu` VALUES (2471, 1, 202);
INSERT INTO `role_menu` VALUES (2472, 1, 13);
INSERT INTO `role_menu` VALUES (2473, 1, 201);
INSERT INTO `role_menu` VALUES (2474, 1, 213);
INSERT INTO `role_menu` VALUES (2475, 1, 14);
INSERT INTO `role_menu` VALUES (2476, 1, 15);
INSERT INTO `role_menu` VALUES (2477, 1, 16);
INSERT INTO `role_menu` VALUES (2478, 1, 4);
INSERT INTO `role_menu` VALUES (2479, 1, 214);
INSERT INTO `role_menu` VALUES (2480, 1, 209);
INSERT INTO `role_menu` VALUES (2481, 1, 17);
INSERT INTO `role_menu` VALUES (2482, 1, 18);
INSERT INTO `role_menu` VALUES (2483, 1, 205);
INSERT INTO `role_menu` VALUES (2484, 1, 206);
INSERT INTO `role_menu` VALUES (2485, 1, 208);
INSERT INTO `role_menu` VALUES (2486, 1, 210);
INSERT INTO `role_menu` VALUES (2487, 1, 215);
INSERT INTO `role_menu` VALUES (2488, 1, 216);
INSERT INTO `role_menu` VALUES (2489, 1, 217);
INSERT INTO `role_menu` VALUES (2490, 1, 218);
INSERT INTO `role_menu` VALUES (2491, 1, 19);
INSERT INTO `role_menu` VALUES (2492, 1, 20);
INSERT INTO `role_menu` VALUES (2493, 1, 5);
INSERT INTO `role_menu` VALUES (2494, 3, 1);
INSERT INTO `role_menu` VALUES (2495, 3, 2);
INSERT INTO `role_menu` VALUES (2496, 3, 6);
INSERT INTO `role_menu` VALUES (2497, 3, 7);
INSERT INTO `role_menu` VALUES (2498, 3, 8);
INSERT INTO `role_menu` VALUES (2499, 3, 9);
INSERT INTO `role_menu` VALUES (2500, 3, 10);
INSERT INTO `role_menu` VALUES (2501, 3, 3);
INSERT INTO `role_menu` VALUES (2502, 3, 11);
INSERT INTO `role_menu` VALUES (2503, 3, 12);
INSERT INTO `role_menu` VALUES (2504, 3, 202);
INSERT INTO `role_menu` VALUES (2505, 3, 13);
INSERT INTO `role_menu` VALUES (2506, 3, 201);
INSERT INTO `role_menu` VALUES (2507, 3, 213);
INSERT INTO `role_menu` VALUES (2508, 3, 14);
INSERT INTO `role_menu` VALUES (2509, 3, 15);
INSERT INTO `role_menu` VALUES (2510, 3, 16);
INSERT INTO `role_menu` VALUES (2511, 3, 4);
INSERT INTO `role_menu` VALUES (2512, 3, 214);
INSERT INTO `role_menu` VALUES (2513, 3, 209);
INSERT INTO `role_menu` VALUES (2514, 3, 17);
INSERT INTO `role_menu` VALUES (2515, 3, 18);
INSERT INTO `role_menu` VALUES (2516, 3, 205);
INSERT INTO `role_menu` VALUES (2517, 3, 206);
INSERT INTO `role_menu` VALUES (2518, 3, 208);
INSERT INTO `role_menu` VALUES (2519, 3, 210);
INSERT INTO `role_menu` VALUES (2520, 3, 215);
INSERT INTO `role_menu` VALUES (2521, 3, 216);
INSERT INTO `role_menu` VALUES (2522, 3, 217);
INSERT INTO `role_menu` VALUES (2523, 3, 218);
INSERT INTO `role_menu` VALUES (2524, 3, 19);
INSERT INTO `role_menu` VALUES (2525, 3, 20);
INSERT INTO `role_menu` VALUES (2526, 3, 5);

-- ----------------------------
-- Table structure for role_resource
-- ----------------------------
DROP TABLE IF EXISTS `role_resource`;
CREATE TABLE `role_resource`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NULL DEFAULT NULL COMMENT '角色id',
  `resource_id` int(11) NULL DEFAULT NULL COMMENT '权限id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8541 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of role_resource
-- ----------------------------
INSERT INTO `role_resource` VALUES (8406, 1, 165);
INSERT INTO `role_resource` VALUES (8407, 1, 192);
INSERT INTO `role_resource` VALUES (8408, 1, 193);
INSERT INTO `role_resource` VALUES (8409, 1, 194);
INSERT INTO `role_resource` VALUES (8410, 1, 195);
INSERT INTO `role_resource` VALUES (8411, 1, 166);
INSERT INTO `role_resource` VALUES (8412, 1, 183);
INSERT INTO `role_resource` VALUES (8413, 1, 184);
INSERT INTO `role_resource` VALUES (8414, 1, 246);
INSERT INTO `role_resource` VALUES (8415, 1, 247);
INSERT INTO `role_resource` VALUES (8416, 1, 167);
INSERT INTO `role_resource` VALUES (8417, 1, 199);
INSERT INTO `role_resource` VALUES (8418, 1, 200);
INSERT INTO `role_resource` VALUES (8419, 1, 201);
INSERT INTO `role_resource` VALUES (8420, 1, 168);
INSERT INTO `role_resource` VALUES (8421, 1, 185);
INSERT INTO `role_resource` VALUES (8422, 1, 186);
INSERT INTO `role_resource` VALUES (8423, 1, 187);
INSERT INTO `role_resource` VALUES (8424, 1, 188);
INSERT INTO `role_resource` VALUES (8425, 1, 189);
INSERT INTO `role_resource` VALUES (8426, 1, 190);
INSERT INTO `role_resource` VALUES (8427, 1, 191);
INSERT INTO `role_resource` VALUES (8428, 1, 254);
INSERT INTO `role_resource` VALUES (8429, 1, 169);
INSERT INTO `role_resource` VALUES (8430, 1, 208);
INSERT INTO `role_resource` VALUES (8431, 1, 209);
INSERT INTO `role_resource` VALUES (8432, 1, 170);
INSERT INTO `role_resource` VALUES (8433, 1, 234);
INSERT INTO `role_resource` VALUES (8434, 1, 235);
INSERT INTO `role_resource` VALUES (8435, 1, 236);
INSERT INTO `role_resource` VALUES (8436, 1, 237);
INSERT INTO `role_resource` VALUES (8437, 1, 171);
INSERT INTO `role_resource` VALUES (8438, 1, 213);
INSERT INTO `role_resource` VALUES (8439, 1, 214);
INSERT INTO `role_resource` VALUES (8440, 1, 215);
INSERT INTO `role_resource` VALUES (8441, 1, 216);
INSERT INTO `role_resource` VALUES (8442, 1, 217);
INSERT INTO `role_resource` VALUES (8443, 1, 224);
INSERT INTO `role_resource` VALUES (8444, 1, 172);
INSERT INTO `role_resource` VALUES (8445, 1, 240);
INSERT INTO `role_resource` VALUES (8446, 1, 241);
INSERT INTO `role_resource` VALUES (8447, 1, 244);
INSERT INTO `role_resource` VALUES (8448, 1, 245);
INSERT INTO `role_resource` VALUES (8449, 1, 267);
INSERT INTO `role_resource` VALUES (8450, 1, 269);
INSERT INTO `role_resource` VALUES (8451, 1, 270);
INSERT INTO `role_resource` VALUES (8452, 1, 173);
INSERT INTO `role_resource` VALUES (8453, 1, 239);
INSERT INTO `role_resource` VALUES (8454, 1, 242);
INSERT INTO `role_resource` VALUES (8455, 1, 276);
INSERT INTO `role_resource` VALUES (8456, 1, 174);
INSERT INTO `role_resource` VALUES (8457, 1, 205);
INSERT INTO `role_resource` VALUES (8458, 1, 206);
INSERT INTO `role_resource` VALUES (8459, 1, 207);
INSERT INTO `role_resource` VALUES (8460, 1, 175);
INSERT INTO `role_resource` VALUES (8461, 1, 218);
INSERT INTO `role_resource` VALUES (8462, 1, 219);
INSERT INTO `role_resource` VALUES (8463, 1, 220);
INSERT INTO `role_resource` VALUES (8464, 1, 221);
INSERT INTO `role_resource` VALUES (8465, 1, 222);
INSERT INTO `role_resource` VALUES (8466, 1, 223);
INSERT INTO `role_resource` VALUES (8467, 1, 176);
INSERT INTO `role_resource` VALUES (8468, 1, 202);
INSERT INTO `role_resource` VALUES (8469, 1, 203);
INSERT INTO `role_resource` VALUES (8470, 1, 204);
INSERT INTO `role_resource` VALUES (8471, 1, 230);
INSERT INTO `role_resource` VALUES (8472, 1, 238);
INSERT INTO `role_resource` VALUES (8473, 1, 177);
INSERT INTO `role_resource` VALUES (8474, 1, 229);
INSERT INTO `role_resource` VALUES (8475, 1, 232);
INSERT INTO `role_resource` VALUES (8476, 1, 233);
INSERT INTO `role_resource` VALUES (8477, 1, 243);
INSERT INTO `role_resource` VALUES (8478, 1, 178);
INSERT INTO `role_resource` VALUES (8479, 1, 196);
INSERT INTO `role_resource` VALUES (8480, 1, 197);
INSERT INTO `role_resource` VALUES (8481, 1, 198);
INSERT INTO `role_resource` VALUES (8482, 1, 257);
INSERT INTO `role_resource` VALUES (8483, 1, 258);
INSERT INTO `role_resource` VALUES (8484, 1, 179);
INSERT INTO `role_resource` VALUES (8485, 1, 225);
INSERT INTO `role_resource` VALUES (8486, 1, 226);
INSERT INTO `role_resource` VALUES (8487, 1, 227);
INSERT INTO `role_resource` VALUES (8488, 1, 228);
INSERT INTO `role_resource` VALUES (8489, 1, 231);
INSERT INTO `role_resource` VALUES (8490, 1, 180);
INSERT INTO `role_resource` VALUES (8491, 1, 210);
INSERT INTO `role_resource` VALUES (8492, 1, 211);
INSERT INTO `role_resource` VALUES (8493, 1, 212);
INSERT INTO `role_resource` VALUES (8494, 1, 278);
INSERT INTO `role_resource` VALUES (8495, 1, 282);
INSERT INTO `role_resource` VALUES (8496, 1, 283);
INSERT INTO `role_resource` VALUES (8497, 1, 284);
INSERT INTO `role_resource` VALUES (8498, 1, 285);
INSERT INTO `role_resource` VALUES (8499, 1, 286);
INSERT INTO `role_resource` VALUES (8500, 1, 287);
INSERT INTO `role_resource` VALUES (8501, 2, 254);
INSERT INTO `role_resource` VALUES (8502, 2, 267);
INSERT INTO `role_resource` VALUES (8503, 2, 269);
INSERT INTO `role_resource` VALUES (8504, 2, 270);
INSERT INTO `role_resource` VALUES (8505, 2, 257);
INSERT INTO `role_resource` VALUES (8506, 2, 258);
INSERT INTO `role_resource` VALUES (8507, 2, 282);
INSERT INTO `role_resource` VALUES (8508, 3, 192);
INSERT INTO `role_resource` VALUES (8509, 3, 195);
INSERT INTO `role_resource` VALUES (8510, 3, 183);
INSERT INTO `role_resource` VALUES (8511, 3, 246);
INSERT INTO `role_resource` VALUES (8512, 3, 199);
INSERT INTO `role_resource` VALUES (8513, 3, 185);
INSERT INTO `role_resource` VALUES (8514, 3, 191);
INSERT INTO `role_resource` VALUES (8515, 3, 254);
INSERT INTO `role_resource` VALUES (8516, 3, 208);
INSERT INTO `role_resource` VALUES (8517, 3, 234);
INSERT INTO `role_resource` VALUES (8518, 3, 237);
INSERT INTO `role_resource` VALUES (8519, 3, 213);
INSERT INTO `role_resource` VALUES (8520, 3, 241);
INSERT INTO `role_resource` VALUES (8521, 3, 239);
INSERT INTO `role_resource` VALUES (8522, 3, 276);
INSERT INTO `role_resource` VALUES (8523, 3, 205);
INSERT INTO `role_resource` VALUES (8524, 3, 218);
INSERT INTO `role_resource` VALUES (8525, 3, 221);
INSERT INTO `role_resource` VALUES (8526, 3, 223);
INSERT INTO `role_resource` VALUES (8527, 3, 202);
INSERT INTO `role_resource` VALUES (8528, 3, 230);
INSERT INTO `role_resource` VALUES (8529, 3, 238);
INSERT INTO `role_resource` VALUES (8530, 3, 232);
INSERT INTO `role_resource` VALUES (8531, 3, 243);
INSERT INTO `role_resource` VALUES (8532, 3, 196);
INSERT INTO `role_resource` VALUES (8533, 3, 257);
INSERT INTO `role_resource` VALUES (8534, 3, 258);
INSERT INTO `role_resource` VALUES (8535, 3, 225);
INSERT INTO `role_resource` VALUES (8536, 3, 231);
INSERT INTO `role_resource` VALUES (8537, 3, 210);
INSERT INTO `role_resource` VALUES (8538, 3, 282);
INSERT INTO `role_resource` VALUES (8539, 3, 286);
INSERT INTO `role_resource` VALUES (8540, 3, 287);

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tag_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标签名',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of tag
-- ----------------------------
INSERT INTO `tag` VALUES (1, '测试标签', '2022-10-14 11:49:37', '2022-10-14 11:49:37');

-- ----------------------------
-- Table structure for talk
-- ----------------------------
DROP TABLE IF EXISTS `talk`;
CREATE TABLE `talk`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '说说id',
  `user_id` int(11) NOT NULL COMMENT '用户id',
  `content` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '说说内容',
  `images` varchar(2500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '图片',
  `is_top` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否置顶',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态 1.公开 2.私密',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of talk
-- ----------------------------

-- ----------------------------
-- Table structure for unique_view
-- ----------------------------
DROP TABLE IF EXISTS `unique_view`;
CREATE TABLE `unique_view`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `views_count` int(11) NOT NULL COMMENT '访问量',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 534 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of unique_view
-- ----------------------------

-- ----------------------------
-- Table structure for user_auth
-- ----------------------------
DROP TABLE IF EXISTS `user_auth`;
CREATE TABLE `user_auth`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_info_id` int(11) NOT NULL COMMENT '用户信息id',
  `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `login_type` tinyint(1) NOT NULL COMMENT '登录类型',
  `ip_address` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户登录ip',
  `ip_source` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'ip来源',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `last_login_time` datetime NULL DEFAULT NULL COMMENT '上次登录时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_auth
-- ----------------------------
INSERT INTO `user_auth` VALUES (2, 2, 'test@qq.com', 'c4ca4238a0b923820dcc509a6f75849b', 1, '110.52.119.165', '湖南省湘潭市 联通', '2022-10-09 15:42:45', '2022-10-14 09:04:55', '2022-10-14 09:04:55');

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱号',
  `nickname` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户昵称',
  `avatar` varchar(1024) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `intro` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户简介',
  `web_site` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '个人网站',
  `is_disable` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否禁用',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_info
-- ----------------------------
INSERT INTO `user_info` VALUES (2, 'test@qq.com', '测试用户', 'https://blog-1311853727.cos.ap-guangzhou.myqcloud.com/user/user.jpeg', NULL, NULL, 0, '2022-10-09 15:42:11', '2022-10-12 10:27:12');

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL COMMENT '用户id',
  `role_id` int(11) NOT NULL COMMENT '角色id',
  PRIMARY KEY (`id`, `user_id`, `role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_role
-- ----------------------------
INSERT INTO `user_role` VALUES (2, 2, 3);

-- ----------------------------
-- Table structure for website_config
-- ----------------------------
DROP TABLE IF EXISTS `website_config`;
CREATE TABLE `website_config`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `config` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '配置信息',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of website_config
-- ----------------------------
INSERT INTO `website_config` VALUES (1, '{\"websiteAvatar\":\"https://blog-1311853727.cos.ap-guangzhou.myqcloud.com/user/admin.jpg\",\"websiteName\":\"水平线之下的个人博客\",\"websiteAuthor\":\"水平线之下\",\"websiteIntro\":\"网站简介\",\"websiteNotice\":\"请前往后台管理-\\u003e系统管理-\\u003e网站管理处修改信息\",\"websiteCreateTime\":\"2022-10-08\",\"websiteRecordNo\":\"湘ICP备2022010883号-1\",\"socialLoginList\":[],\"socialUrlList\":[\"gitee\",\"github\",\"qq\"],\"qq\":\"\",\"github\":\"\",\"gitee\":\"\",\"touristAvatar\":\"https://static.talkxj.com/photos/0bca52afdb2b9998132355d716390c9f.png\",\"userAvatar\":\"https://static.talkxj.com/config/2cd793c8744199053323546875655f32.jpg\",\"isCommentReview\":0,\"isMessageReview\":0,\"isEmailNotice\":0,\"isReward\":0,\"weiXinQRCode\":\"https://static.talkxj.com/photos/4f767ef84e55ab9ad42b2d20e51deca1.png\",\"alipayQRCode\":\"https://static.talkxj.com/photos/13d83d77cc1f7e4e0437d7feaf56879f.png\",\"articleCover\":\"\",\"isChatRoom\":0,\"websocketUrl\":\"ws://127.0.0.1:8080/websocket\",\"isMusicPlayer\":0}', '2021-08-09 19:37:30', '2022-10-12 15:13:59');

SET FOREIGN_KEY_CHECKS = 1;
