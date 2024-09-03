-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- 主机： mysql:3306
-- 生成日期： 2024-09-04 07:35:15
-- 服务器版本： 8.0.34
-- PHP 版本： 8.2.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `go_blog`
--

-- --------------------------------------------------------

--
-- 表的结构 `gf_category`
--

CREATE TABLE `gf_category` (
                               `id` int UNSIGNED NOT NULL COMMENT '分类ID，自增主键',
                               `content_type` char(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '内容类型：topic, ask, article, reply',
                               `key` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '栏目唯一键名，用于程序部分场景硬编码，一般不会用得到',
                               `parent_id` int UNSIGNED DEFAULT '0' COMMENT '父级分类ID，用于层级管理',
                               `user_id` int UNSIGNED NOT NULL COMMENT '创建的用户ID',
                               `name` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '分类名称',
                               `sort` int UNSIGNED DEFAULT '0' COMMENT '排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶',
                               `thumb` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '封面图',
                               `brief` varchar(9000) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '简述',
                               `content` mediumtext COLLATE utf8mb4_unicode_ci COMMENT '详细介绍',
                               `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                               `updated_at` datetime DEFAULT NULL COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='栏目列表' ROW_FORMAT=DYNAMIC;

--
-- 转存表中的数据 `gf_category`
--

INSERT INTO `gf_category` (`id`, `content_type`, `key`, `parent_id`, `user_id`, `name`, `sort`, `thumb`, `brief`, `content`, `created_at`, `updated_at`) VALUES
                                                                                                                                                             (1, 'topic', '', 0, 1, '随聊', 0, NULL, NULL, '', '2020-06-04 00:03:48', '2020-06-04 00:03:55'),
                                                                                                                                                             (2, 'ask', '', 0, 1, 'BUG反馈', 0, NULL, NULL, NULL, '2020-06-04 00:04:37', '2020-06-04 00:04:41'),
                                                                                                                                                             (3, 'topic', '', 1, 1, 'TopicCat2', 0, NULL, NULL, NULL, NULL, NULL),
                                                                                                                                                             (4, 'topic', '', 3, 1, 'TopicCat3', 0, NULL, NULL, NULL, NULL, NULL),
                                                                                                                                                             (5, 'ask', NULL, 2, 1, 'AskCat2', 0, NULL, NULL, NULL, NULL, NULL),
                                                                                                                                                             (6, 'topic', NULL, 0, 1, '翻译', 0, NULL, NULL, NULL, NULL, NULL),
                                                                                                                                                             (7, 'topic', NULL, 0, 1, '分享', 0, NULL, NULL, NULL, NULL, NULL),
                                                                                                                                                             (9, 'article', NULL, 0, 1, '文章分享', 0, NULL, NULL, NULL, NULL, NULL),
                                                                                                                                                             (10, 'article', NULL, 0, 1, '源码解读', 0, NULL, NULL, NULL, NULL, NULL),
                                                                                                                                                             (11, 'article', NULL, 0, 1, '技术架构', 0, NULL, NULL, NULL, NULL, NULL),
                                                                                                                                                             (12, 'topic', NULL, 0, 1, '招聘', 0, NULL, NULL, NULL, NULL, NULL),
                                                                                                                                                             (13, 'ask', NULL, 0, 1, '我要问问', 0, NULL, NULL, NULL, NULL, NULL),
                                                                                                                                                             (14, 'ask', NULL, 0, 1, 'VIP问答', 0, NULL, NULL, NULL, NULL, NULL),
                                                                                                                                                             (15, 'blog', NULL, 0, 1, '技术分享', 0, NULL, NULL, NULL, NULL, NULL),
                                                                                                                                                             (16, 'blog', NULL, 0, 1, '生活日常', 0, NULL, NULL, NULL, NULL, NULL),
                                                                                                                                                             (17, 'blog', NULL, 0, 1, '开发笔记', 0, NULL, NULL, NULL, NULL, NULL);

-- --------------------------------------------------------

--
-- 表的结构 `gf_content`
--

CREATE TABLE `gf_content` (
                              `id` int UNSIGNED NOT NULL COMMENT '自增ID',
                              `key` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '唯一键名，用于程序硬编码，一般不常用',
                              `type` char(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '内容模型: topic, ask, article等，具体由程序定义',
                              `category_id` int UNSIGNED NOT NULL COMMENT '栏目ID',
                              `user_id` int UNSIGNED NOT NULL COMMENT '用户ID',
                              `adopted_reply_id` int UNSIGNED DEFAULT NULL COMMENT '采纳的回复ID，问答模块有效',
                              `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '标题',
                              `content` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '内容',
                              `sort` int UNSIGNED DEFAULT '0' COMMENT '排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶',
                              `brief` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '摘要',
                              `thumb` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '缩略图',
                              `tags` varchar(900) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '标签名称列表，以JSON存储',
                              `referer` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '内容来源，例如github/gitee',
                              `status` smallint UNSIGNED DEFAULT '0' COMMENT '状态 0: 正常, 1: 禁用',
                              `reply_count` int UNSIGNED DEFAULT '0' COMMENT '回复数量',
                              `view_count` int UNSIGNED DEFAULT '0' COMMENT '浏览数量',
                              `zan_count` int UNSIGNED DEFAULT '0' COMMENT '赞',
                              `cai_count` int UNSIGNED DEFAULT '0' COMMENT '踩',
                              `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                              `updated_at` datetime DEFAULT NULL COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='问答列表' ROW_FORMAT=DYNAMIC;



-- --------------------------------------------------------

--
-- 表的结构 `gf_file`
--

CREATE TABLE `gf_file` (
                           `id` int UNSIGNED NOT NULL COMMENT '自增ID',
                           `name` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件名称',
                           `src` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '本地文件存储路径',
                           `url` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'URL地址，可能为空',
                           `user_id` int UNSIGNED NOT NULL COMMENT '操作用户',
                           `created_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='文件列表' ROW_FORMAT=DYNAMIC;


-- --------------------------------------------------------

--
-- 表的结构 `gf_interact`
--

CREATE TABLE `gf_interact` (
                               `id` int UNSIGNED NOT NULL COMMENT '自增ID',
                               `type` tinyint NOT NULL COMMENT '操作类型。0:赞，1:踩。',
                               `user_id` int UNSIGNED NOT NULL COMMENT '操作用户',
                               `target_id` int UNSIGNED NOT NULL COMMENT '对应内容ID，该内容可能是content, reply',
                               `target_type` char(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '内容模型: content, reply, 具体由程序定义',
                               `count` int UNSIGNED DEFAULT NULL COMMENT '操作数据值',
                               `created_at` datetime DEFAULT NULL,
                               `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='交互管理表' ROW_FORMAT=DYNAMIC;

--
-- 转存表中的数据 `gf_interact`
--

INSERT INTO `gf_interact` (`id`, `type`, `user_id`, `target_id`, `target_type`, `count`, `created_at`, `updated_at`) VALUES
                                                                                                                         (69, 1, 15, 10803, 'content', 0, '2020-11-17 21:24:33', '2020-11-17 21:24:33'),
                                                                                                                         (70, 0, 15, 46925, 'content', 0, '2020-11-18 22:03:31', '2020-11-18 22:03:31'),
                                                                                                                         (75, 0, 17, 10803, 'content', 0, '2020-12-29 23:04:00', '2020-12-29 23:04:00'),
                                                                                                                         (77, 0, 17, 174, 'content', 0, '2020-12-29 23:04:18', '2020-12-29 23:04:18'),
                                                                                                                         (79, 1, 17, 174, 'content', 0, '2020-12-29 23:04:21', '2020-12-29 23:04:21'),
                                                                                                                         (82, 1, 17, 10793, 'content', 0, '2020-12-30 01:09:10', '2020-12-30 01:09:10'),
                                                                                                                         (83, 0, 17, 10793, 'content', 0, '2020-12-30 01:09:11', '2020-12-30 01:09:11'),
                                                                                                                         (85, 0, 17, 10794, 'content', 0, '2020-12-30 01:09:19', '2020-12-30 01:09:19'),
                                                                                                                         (87, 1, 17, 10803, 'content', 0, '2020-12-30 23:33:54', '2020-12-30 23:33:54'),
                                                                                                                         (88, 0, 17, 10801, 'content', 0, '2020-12-30 23:33:58', '2020-12-30 23:33:58'),
                                                                                                                         (89, 1, 17, 10801, 'content', 0, '2020-12-30 23:33:58', '2020-12-30 23:33:58');

-- --------------------------------------------------------

--
-- 表的结构 `gf_page`
--

CREATE TABLE `gf_page` (
                           `id` int NOT NULL COMMENT 'id',
                           `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题',
                           `url_path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述',
                           `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '介绍',
                           `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '内容',
                           `template` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '页面模版',
                           `reply` tinyint NOT NULL DEFAULT '0' COMMENT '是否开启评论',
                           `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- 转存表中的数据 `gf_page`
--

INSERT INTO `gf_page` (`id`, `title`, `url_path`, `description`, `content`, `template`, `reply`, `created_at`, `updated_at`) VALUES
    (1, 'about me', 'about', '关于本人', '### 关于我\r\n九一年生，坐标成都。喜欢数码科技，对新鲜事物充满兴趣，会写一点点代码，什么都想学，什么都不精通。偶尔玩玩游戏，听听音乐。', 'default', 1, '2024-08-30 03:51:56', '2024-08-30 20:02:31');

-- --------------------------------------------------------

--
-- 表的结构 `gf_reply`
--

CREATE TABLE `gf_reply` (
                            `id` int UNSIGNED NOT NULL COMMENT '回复ID',
                            `parent_id` int UNSIGNED DEFAULT '0' COMMENT '回复对应的上一级回复ID(没有的话默认为0)',
                            `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '回复标题',
                            `content` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '回复内容',
                            `target_type` char(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '评论类型: content, reply',
                            `target_id` int UNSIGNED NOT NULL COMMENT '对应内容ID，可能回复的是另一个回复，所以这里没有使用content_id',
                            `user_id` int UNSIGNED NOT NULL COMMENT '网站用户ID',
                            `zan_count` int UNSIGNED DEFAULT '0' COMMENT '赞',
                            `cai_count` int UNSIGNED DEFAULT '0' COMMENT '踩',
                            `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                            `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='回复列表' ROW_FORMAT=DYNAMIC;

--
-- 转存表中的数据 `gf_reply`
--

INSERT INTO `gf_reply` (`id`, `parent_id`, `title`, `content`, `target_type`, `target_id`, `user_id`, `zan_count`, `cai_count`, `created_at`, `updated_at`) VALUES
                                                                                                                                                                (32, 0, '', '<p>支持一下Goframe</p>\n', 'topic', 10785, 1, 0, 0, '2021-01-01 22:40:27', '2021-01-01 22:40:27'),
                                                                                                                                                                (33, 0, '', '<p>支持一下GoFrame</p>\n', 'topic', 20822, 1, 0, 0, '2021-01-01 22:44:13', '2021-01-01 22:44:13');

-- --------------------------------------------------------

--
-- 表的结构 `gf_setting`
--

CREATE TABLE `gf_setting` (
                              `k` char(45) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '键名',
                              `v` varchar(9000) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '键值',
                              `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                              `updated_at` datetime DEFAULT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='键值对设置表' ROW_FORMAT=DYNAMIC;

--
-- 转存表中的数据 `gf_setting`
--

INSERT INTO `gf_setting` (`k`, `v`, `created_at`, `updated_at`) VALUES
    ('TopMenus', '[{\"name\":\"首页\",\"url\":\"[[main]]/\",\"target\":\"\",\"items\":null},{\"name\":\"博客\",\"url\":\"[[blog]]\",\"target\":\"\",\"items\":null}]', '2020-11-11 23:11:30', '2020-11-11 23:13:03');

-- --------------------------------------------------------

--
-- 表的结构 `gf_user`
--

CREATE TABLE `gf_user` (
                           `id` int UNSIGNED NOT NULL COMMENT 'UID',
                           `passport` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '账号',
                           `password` char(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'MD5密码',
                           `nickname` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '昵称',
                           `avatar` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '头像地址',
                           `status` tinyint DEFAULT '1' COMMENT '状态 0:启用 1:禁用',
                           `gender` tinyint(1) DEFAULT '0' COMMENT '性别 0: 未设置 1: 男 2: 女',
                           `created_at` datetime DEFAULT NULL COMMENT '注册时间',
                           `updated_at` datetime DEFAULT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户基础表' ROW_FORMAT=DYNAMIC;

--
-- 转存表中的数据 `gf_user`
--

INSERT INTO `gf_user` VALUES (1, 'goframe', 'b741a2d80cc7320cccc56303cf0fba0a', 'goframe', 'https://i.loli.net/2020/06/13/4XaPUnEjpeVGfzt.jpg', 1, 0, '2020-06-03 21:04:11', '2020-06-13 21:20:59');


--
-- 转储表的索引
--

--
-- 表的索引 `gf_category`
--
ALTER TABLE `gf_category`
    ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `key` (`key`) USING BTREE,
  ADD KEY `content_type` (`content_type`) USING BTREE;

--
-- 表的索引 `gf_content`
--
ALTER TABLE `gf_content`
    ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `key` (`key`) USING BTREE,
  ADD KEY `updated_at` (`updated_at`) USING BTREE,
  ADD KEY `view_count` (`view_count`) USING BTREE,
  ADD KEY `zan_count` (`zan_count`) USING BTREE,
  ADD KEY `type_category_id` (`type`,`category_id`) USING BTREE,
  ADD KEY `idx_content_user_id` (`user_id`) USING BTREE;

--
-- 表的索引 `gf_file`
--
ALTER TABLE `gf_file`
    ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `user_id` (`user_id`) USING BTREE;

--
-- 表的索引 `gf_interact`
--
ALTER TABLE `gf_interact`
    ADD PRIMARY KEY (`id`) USING BTREE,
  ADD UNIQUE KEY `unique` (`user_id`,`target_id`,`target_type`,`type`) USING BTREE;

--
-- 表的索引 `gf_page`
--
ALTER TABLE `gf_page`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `gf_page_unique_path` (`url_path`);

--
-- 表的索引 `gf_reply`
--
ALTER TABLE `gf_reply`
    ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `content_type_target_id` (`target_type`,`target_id`) USING BTREE,
  ADD KEY `user_id` (`user_id`) USING BTREE;

--
-- 表的索引 `gf_setting`
--
ALTER TABLE `gf_setting`
    ADD PRIMARY KEY (`k`) USING BTREE;

--
-- 表的索引 `gf_user`
--
ALTER TABLE `gf_user`
    ADD PRIMARY KEY (`id`) USING BTREE,
  ADD UNIQUE KEY `uni_user_passport` (`passport`) USING BTREE,
  ADD UNIQUE KEY `uni_user_nickname` (`nickname`) USING BTREE;

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `gf_category`
--
ALTER TABLE `gf_category`
    MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类ID，自增主键', AUTO_INCREMENT=18;

--
-- 使用表AUTO_INCREMENT `gf_content`
--
ALTER TABLE `gf_content`
    MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID', AUTO_INCREMENT=46940;

--
-- 使用表AUTO_INCREMENT `gf_file`
--
ALTER TABLE `gf_file`
    MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID', AUTO_INCREMENT=128;

--
-- 使用表AUTO_INCREMENT `gf_interact`
--
ALTER TABLE `gf_interact`
    MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID', AUTO_INCREMENT=90;

--
-- 使用表AUTO_INCREMENT `gf_page`
--
ALTER TABLE `gf_page`
    MODIFY `id` int NOT NULL AUTO_INCREMENT COMMENT 'id', AUTO_INCREMENT=5;

--
-- 使用表AUTO_INCREMENT `gf_reply`
--
ALTER TABLE `gf_reply`
    MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '回复ID', AUTO_INCREMENT=34;

--
-- 使用表AUTO_INCREMENT `gf_user`
--
ALTER TABLE `gf_user`
    MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'UID', AUTO_INCREMENT=19;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
