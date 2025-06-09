/*
 Navicat Premium Dump SQL

 Source Server         : mysql_connect
 Source Server Type    : MySQL
 Source Server Version : 80403 (8.4.3)
 Source Host           : localhost:3306
 Source Schema         : msproject

 Target Server Type    : MySQL
 Target Server Version : 80403 (8.4.3)
 File Encoding         : 65001

 Date: 08/06/2025 19:09:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ms_department
-- ----------------------------
DROP TABLE IF EXISTS `ms_department`;
CREATE TABLE `ms_department`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `organization_code` bigint NULL DEFAULT NULL COMMENT '组织编号',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `sort` int NULL DEFAULT 0 COMMENT '排序',
  `pcode` bigint NULL DEFAULT NULL COMMENT '上级编号',
  `icon` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标',
  `create_time` bigint NULL DEFAULT NULL COMMENT '创建时间',
  `path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '上级路径',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '部门表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_department
-- ----------------------------
INSERT INTO `ms_department` VALUES (5, 14, 'D1', 0, 0, NULL, 1747198941019, '');
INSERT INTO `ms_department` VALUES (6, 14, 'D2', 0, 0, NULL, 1747199104859, '');
INSERT INTO `ms_department` VALUES (7, 14, 'D3', 0, 0, NULL, 1747199455215, '');
INSERT INTO `ms_department` VALUES (8, 14, 'D4', 0, 0, NULL, 1747222350786, '');
INSERT INTO `ms_department` VALUES (9, 14, 'DD1', 0, 8, NULL, 1747222442234, '');

-- ----------------------------
-- Table structure for ms_department_member
-- ----------------------------
DROP TABLE IF EXISTS `ms_department_member`;
CREATE TABLE `ms_department_member`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `department_code` bigint NULL DEFAULT NULL COMMENT '部门id',
  `organization_code` bigint NULL DEFAULT NULL COMMENT '组织id',
  `account_code` bigint NULL DEFAULT NULL COMMENT '成员id',
  `join_time` bigint NULL DEFAULT NULL COMMENT '加入时间',
  `is_principal` tinyint(1) NULL DEFAULT NULL COMMENT '是否负责人',
  `is_owner` tinyint(1) NULL DEFAULT 0 COMMENT '拥有者',
  `authorize` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '角色',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 38 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '部门-成员表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_department_member
-- ----------------------------

-- ----------------------------
-- Table structure for ms_file
-- ----------------------------
DROP TABLE IF EXISTS `ms_file`;
CREATE TABLE `ms_file`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `path_name` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '相对路径',
  `title` char(90) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '名称',
  `extension` char(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '扩展名',
  `size` int UNSIGNED NULL DEFAULT 0 COMMENT '文件大小',
  `object_type` char(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '对象类型',
  `organization_code` bigint NULL DEFAULT NULL COMMENT '组织编码',
  `task_code` bigint NULL DEFAULT NULL COMMENT '任务编码',
  `project_code` bigint NULL DEFAULT NULL COMMENT '项目编码',
  `create_by` bigint NULL DEFAULT NULL COMMENT '上传人',
  `create_time` bigint NULL DEFAULT NULL COMMENT '创建时间',
  `downloads` mediumint UNSIGNED NULL DEFAULT 0 COMMENT '下载次数',
  `extra` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '额外信息',
  `deleted` tinyint(1) NULL DEFAULT 0 COMMENT '删除标记',
  `file_url` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '完整地址',
  `file_type` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '文件类型',
  `deleted_time` bigint NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 54 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '文件表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_file
-- ----------------------------
INSERT INTO `ms_file` VALUES (51, 'upload/e08cf4b297/e08df7b093/2025-05-08/WeGameMiniLoader.LOL.4.11.14.1542.exe', 'WeGameMiniLoader.LOL.4.11.14.1542.exe', '.exe', 6945064, '', 14, 12363, 13047, 1015, 1746634393480, 0, '', 0, 'http://localhost/upload/e08cf4b297/e08df7b093/2025-05-08/WeGameMiniLoader.LOL.4.11.14.1542.exe', 'application/x-msdownload', 0);
INSERT INTO `ms_file` VALUES (52, 'upload/e08cf4b297/e08df7b093/2025-05-08/TortoiseGit-LanguagePack-2.17.0.0-64bit-zh_CN.msi', 'TortoiseGit-LanguagePack-2.17.0.0-64bit-zh_CN.msi', '.msi', 4485120, '', 14, 12363, 13047, 1015, 1746634616839, 0, '', 0, 'http://localhost/upload/e08cf4b297/e08df7b093/2025-05-08/TortoiseGit-LanguagePack-2.17.0.0-64bit-zh_CN.msi', 'application/octet-stream', 0);
INSERT INTO `ms_file` VALUES (53, 'upload/e08cf4b297/e08df7b093/2025-05-08/TortoiseGit-2.17.0.2-64bit.msi', 'TortoiseGit-2.17.0.2-64bit.msi', '.msi', 22618112, '', 14, 12363, 13047, 1015, 1746634716867, 0, '', 0, 'http://localhost/upload/e08cf4b297/e08df7b093/2025-05-08/TortoiseGit-2.17.0.2-64bit.msi', 'application/octet-stream', 0);

-- ----------------------------
-- Table structure for ms_member
-- ----------------------------
DROP TABLE IF EXISTS `ms_member`;
CREATE TABLE `ms_member`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '系统前台用户表',
  `account` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '用户登陆账号',
  `password` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '登陆密码',
  `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '用户昵称',
  `mobile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机',
  `realname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '真实姓名',
  `create_time` bigint NULL DEFAULT NULL COMMENT '创建时间',
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '状态',
  `last_login_time` bigint NULL DEFAULT NULL COMMENT '上次登录时间',
  `sex` tinyint NULL DEFAULT 0 COMMENT '性别',
  `avatar` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '头像',
  `idcard` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '身份证',
  `province` int NULL DEFAULT 0 COMMENT '省',
  `city` int NULL DEFAULT 0 COMMENT '市',
  `area` int NULL DEFAULT 0 COMMENT '区',
  `address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '所在地址',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '备注',
  `email` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `dingtalk_openid` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '钉钉openid',
  `dingtalk_unionid` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '钉钉unionid',
  `dingtalk_userid` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '钉钉用户id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1016 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '用户表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_member
-- ----------------------------
INSERT INTO `ms_member` VALUES (1014, 'Tom', '3d9188577cc9bfe9291ac66b5cc872b7', 'Tom', '13003810726', '', 1743998095070, 1, 1743998095070, 0, '', '', 0, 0, 0, '', '', 'nlp_irefe@163.com', '', '', '');
INSERT INTO `ms_member` VALUES (1015, 'Test', 'b7b31cfd0811f80675cbb850567a66ba', 'Test', '18759766318', '', 1744167903827, 1, 1744167903827, 0, '', '', 0, 0, 0, '', '', 'test@163.com', '', '', '');

-- ----------------------------
-- Table structure for ms_member_account
-- ----------------------------
DROP TABLE IF EXISTS `ms_member_account`;
CREATE TABLE `ms_member_account`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `member_code` bigint NULL DEFAULT NULL COMMENT '所属账号id',
  `organization_code` bigint NULL DEFAULT NULL COMMENT '所属组织',
  `department_code` bigint NULL DEFAULT NULL COMMENT '部门编号',
  `authorize` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色',
  `is_owner` tinyint(1) NULL DEFAULT 0 COMMENT '是否主账号',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '姓名',
  `mobile` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号码',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮件',
  `create_time` bigint NULL DEFAULT NULL COMMENT '创建时间',
  `last_login_time` bigint NULL DEFAULT NULL COMMENT '上次登录时间',
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '状态0禁用 1使用中',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像',
  `position` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '职位',
  `department` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '部门',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 36 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '组织账号表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_member_account
-- ----------------------------
INSERT INTO `ms_member_account` VALUES (35, 1015, NULL, NULL, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for ms_organization
-- ----------------------------
DROP TABLE IF EXISTS `ms_organization`;
CREATE TABLE `ms_organization`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `avatar` varchar(600) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `member_id` bigint NULL DEFAULT NULL COMMENT '拥有者',
  `create_time` bigint NULL DEFAULT NULL COMMENT '创建时间',
  `personal` tinyint(1) NULL DEFAULT 0 COMMENT '是否个人项目',
  `address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址',
  `province` int NULL DEFAULT 0 COMMENT '省',
  `city` int NULL DEFAULT 0 COMMENT '市',
  `area` int NULL DEFAULT 0 COMMENT '区',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '组织表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_organization
-- ----------------------------
INSERT INTO `ms_organization` VALUES (13, 'Tom个人组织', 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fc-ssl.dtstatic.com%2Fuploads%2Fblog%2F202103%2F31%2F20210331160001_9a852.thumb.1000_0.jpg&refer=http%3A%2F%2Fc-ssl.dtstatic.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1673017724&t=ced22fc74624e6940fd6a89a21d30cc5', '', 1014, 1743998095071, 1, '', 0, 0, 0);
INSERT INTO `ms_organization` VALUES (14, 'Test个人组织', 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fc-ssl.dtstatic.com%2Fuploads%2Fblog%2F202103%2F31%2F20210331160001_9a852.thumb.1000_0.jpg&refer=http%3A%2F%2Fc-ssl.dtstatic.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1673017724&t=ced22fc74624e6940fd6a89a21d30cc5', '', 1015, 1744167903833, 1, '', 0, 0, 0);

-- ----------------------------
-- Table structure for ms_project
-- ----------------------------
DROP TABLE IF EXISTS `ms_project`;
CREATE TABLE `ms_project`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `cover` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '封面',
  `name` varchar(90) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '名称',
  `description` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '描述',
  `access_control_type` tinyint NULL DEFAULT 0 COMMENT '访问控制l类型',
  `white_list` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '可以访问项目的权限组（白名单）',
  `order` int UNSIGNED NULL DEFAULT 0 COMMENT '排序',
  `deleted` tinyint(1) NULL DEFAULT 0 COMMENT '删除标记',
  `template_code` int NULL DEFAULT NULL COMMENT '项目类型',
  `schedule` double(5, 2) NULL DEFAULT 0.00 COMMENT '进度',
  `create_time` bigint NULL DEFAULT NULL COMMENT '创建时间',
  `organization_code` bigint NULL DEFAULT NULL COMMENT '组织id',
  `deleted_time` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '删除时间',
  `private` tinyint(1) NULL DEFAULT 1 COMMENT '是否私有',
  `prefix` varchar(10) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '项目前缀',
  `open_prefix` tinyint(1) NULL DEFAULT 0 COMMENT '是否开启项目前缀',
  `archive` tinyint(1) NULL DEFAULT 0 COMMENT '是否归档',
  `archive_time` bigint NULL DEFAULT NULL COMMENT '归档时间',
  `open_begin_time` tinyint(1) NULL DEFAULT 0 COMMENT '是否开启任务开始时间',
  `open_task_private` tinyint(1) NULL DEFAULT 0 COMMENT '是否开启新任务默认开启隐私模式',
  `task_board_theme` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT 'default' COMMENT '看板风格',
  `begin_time` bigint NULL DEFAULT NULL COMMENT '项目开始日期',
  `end_time` bigint NULL DEFAULT NULL COMMENT '项目截止日期',
  `auto_update_schedule` tinyint(1) NULL DEFAULT 0 COMMENT '自动更新项目进度',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `project`(`order` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13050 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '项目表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_project
-- ----------------------------
INSERT INTO `ms_project` VALUES (13043, 'https://img2.baidu.com/it/u=792555388,2449797505&fm=253&fmt=auto&app=138&f=JPEG?w=667&h=500', '我的项目12', '测试', 0, '', 0, 0, 11, 45.00, 1744888400490, 14, '', 1, '', 0, 0, 0, 0, 0, 'simple', 0, 0, 0);
INSERT INTO `ms_project` VALUES (13044, 'https://img2.baidu.com/it/u=792555388,2449797505&fm=253&fmt=auto&app=138&f=JPEG?w=667&h=500', 'project2', '测试', 0, '', 0, 0, 19, 2.00, 1744899261526, 14, '', 0, '', 0, 0, 0, 0, 0, 'simple', 0, 0, 0);
INSERT INTO `ms_project` VALUES (13045, 'https://img2.baidu.com/it/u=792555388,2449797505&fm=253&fmt=auto&app=138&f=JPEG?w=667&h=500', '我的项目', '介绍', 0, '', 0, 0, 19, 45.00, 1744958143283, 14, '', 1, '', 0, 0, 0, 0, 0, 'default', 0, 0, 0);
INSERT INTO `ms_project` VALUES (13046, 'https://img2.baidu.com/it/u=792555388,2449797505&fm=253&fmt=auto&app=138&f=JPEG?w=667&h=500', '123', '123', 0, '', 0, 0, 13, 0.00, 1744958683221, 14, '', 0, '', 0, 0, 0, 0, 0, 'simple', 0, 0, 0);
INSERT INTO `ms_project` VALUES (13047, 'https://img2.baidu.com/it/u=792555388,2449797505&fm=253&fmt=auto&app=138&f=JPEG?w=667&h=500', '测试任务步骤', '测试任务步骤', 0, '', 0, 0, 12, 0.00, 1745489790018, 14, '', 0, '', 0, 0, 0, 0, 0, 'simple', 0, 0, 0);
INSERT INTO `ms_project` VALUES (13048, 'https://img2.baidu.com/it/u=792555388,2449797505&fm=253&fmt=auto&app=138&f=JPEG?w=667&h=500', '测试项目', '测试项目111111', 0, '', 0, 0, 13, 0.00, 1747408522054, 14, '', 0, '', 0, 0, 0, 0, 0, 'simple', 0, 0, 0);
INSERT INTO `ms_project` VALUES (13049, 'https://img2.baidu.com/it/u=792555388,2449797505&fm=253&fmt=auto&app=138&f=JPEG?w=667&h=500', '测试项目2', '测试项目2简介', 0, '', 0, 0, 12, 0.00, 1747408830839, 14, '', 0, '', 0, 0, 0, 0, 0, 'simple', 0, 0, 0);

-- ----------------------------
-- Table structure for ms_project_auth
-- ----------------------------
DROP TABLE IF EXISTS `ms_project_auth`;
CREATE TABLE `ms_project_auth`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '权限名称',
  `status` tinyint UNSIGNED NULL DEFAULT 1 COMMENT '状态(0:禁用,1:启用)',
  `sort` smallint UNSIGNED NULL DEFAULT 0 COMMENT '排序权重',
  `desc` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '备注说明',
  `create_by` bigint UNSIGNED NULL DEFAULT 0 COMMENT '创建人',
  `create_at` bigint NULL DEFAULT NULL COMMENT '创建时间',
  `organization_code` bigint NULL DEFAULT NULL COMMENT '所属组织',
  `is_default` tinyint(1) NULL DEFAULT 0 COMMENT '是否默认',
  `type` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '权限类型',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '项目权限表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_project_auth
-- ----------------------------
INSERT INTO `ms_project_auth` VALUES (1, '管理员', 1, 0, '管理员', 0, 1746887957126, 14, 0, 'admin');
INSERT INTO `ms_project_auth` VALUES (2, '成员', 1, 0, '成员', 0, 1746887957205, 14, 1, 'member');

-- ----------------------------
-- Table structure for ms_project_auth_node
-- ----------------------------
DROP TABLE IF EXISTS `ms_project_auth_node`;
CREATE TABLE `ms_project_auth_node`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `auth` bigint UNSIGNED NULL DEFAULT NULL COMMENT '角色ID',
  `node` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '节点路径',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `index_system_auth_auth`(`auth` ASC) USING BTREE,
  INDEX `index_system_auth_node`(`node` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6467 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '项目角色与节点绑定' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_project_auth_node
-- ----------------------------
INSERT INTO `ms_project_auth_node` VALUES (3219, 2, 'project/account/index');
INSERT INTO `ms_project_auth_node` VALUES (3220, 2, 'project/auth/index');
INSERT INTO `ms_project_auth_node` VALUES (3221, 2, 'project/index/index');
INSERT INTO `ms_project_auth_node` VALUES (3222, 2, 'project/index');
INSERT INTO `ms_project_auth_node` VALUES (3223, 2, 'project/index/changecurrentorganization');
INSERT INTO `ms_project_auth_node` VALUES (3224, 2, 'project/index/systemconfig');
INSERT INTO `ms_project_auth_node` VALUES (3225, 2, 'project/index/info');
INSERT INTO `ms_project_auth_node` VALUES (3226, 2, 'project/index/editpersonal');
INSERT INTO `ms_project_auth_node` VALUES (3227, 2, 'project/index/editpassword');
INSERT INTO `ms_project_auth_node` VALUES (3228, 2, 'project/index/uploadimg');
INSERT INTO `ms_project_auth_node` VALUES (3229, 2, 'project/index/uploadavatar');
INSERT INTO `ms_project_auth_node` VALUES (3230, 2, 'project/menu/menu');
INSERT INTO `ms_project_auth_node` VALUES (3231, 2, 'project/node/index');
INSERT INTO `ms_project_auth_node` VALUES (3232, 2, 'project/node/alllist');
INSERT INTO `ms_project_auth_node` VALUES (3233, 2, 'project/notify/index');
INSERT INTO `ms_project_auth_node` VALUES (3234, 2, 'project/notify');
INSERT INTO `ms_project_auth_node` VALUES (3235, 2, 'project/notify/noreads');
INSERT INTO `ms_project_auth_node` VALUES (3236, 2, 'project/notify/setreadied');
INSERT INTO `ms_project_auth_node` VALUES (3237, 2, 'project/notify/batchdel');
INSERT INTO `ms_project_auth_node` VALUES (3238, 2, 'project/notify/read');
INSERT INTO `ms_project_auth_node` VALUES (3239, 2, 'project/notify/delete');
INSERT INTO `ms_project_auth_node` VALUES (3240, 2, 'project/organization/index');
INSERT INTO `ms_project_auth_node` VALUES (3241, 2, 'project/organization');
INSERT INTO `ms_project_auth_node` VALUES (3242, 2, 'project/organization/save');
INSERT INTO `ms_project_auth_node` VALUES (3243, 2, 'project/organization/read');
INSERT INTO `ms_project_auth_node` VALUES (3244, 2, 'project/organization/edit');
INSERT INTO `ms_project_auth_node` VALUES (3245, 2, 'project/organization/delete');
INSERT INTO `ms_project_auth_node` VALUES (3246, 2, 'project/project/index');
INSERT INTO `ms_project_auth_node` VALUES (3247, 2, 'project/project/read');
INSERT INTO `ms_project_auth_node` VALUES (3248, 2, 'project/project_collect/collect');
INSERT INTO `ms_project_auth_node` VALUES (3249, 2, 'project/project_collect');
INSERT INTO `ms_project_auth_node` VALUES (3250, 2, 'project/project_member/index');
INSERT INTO `ms_project_auth_node` VALUES (3251, 2, 'project/project_template/index');
INSERT INTO `ms_project_auth_node` VALUES (3252, 2, 'project/task/index');
INSERT INTO `ms_project_auth_node` VALUES (3253, 2, 'project/task/read');
INSERT INTO `ms_project_auth_node` VALUES (3254, 2, 'project/task/save');
INSERT INTO `ms_project_auth_node` VALUES (3255, 2, 'project/task/taskdone');
INSERT INTO `ms_project_auth_node` VALUES (3256, 2, 'project/task/assigntask');
INSERT INTO `ms_project_auth_node` VALUES (3257, 2, 'project/task/sort');
INSERT INTO `ms_project_auth_node` VALUES (3258, 2, 'project/task/createcomment');
INSERT INTO `ms_project_auth_node` VALUES (3259, 2, 'project/task/like');
INSERT INTO `ms_project_auth_node` VALUES (3260, 2, 'project/task/star');
INSERT INTO `ms_project_auth_node` VALUES (3261, 2, 'project/task_log/index');
INSERT INTO `ms_project_auth_node` VALUES (3262, 2, 'project/task_log');
INSERT INTO `ms_project_auth_node` VALUES (3263, 2, 'project/task_log/getlistbyselfproject');
INSERT INTO `ms_project_auth_node` VALUES (3264, 2, 'project/task_member/index');
INSERT INTO `ms_project_auth_node` VALUES (3265, 2, 'project/task_member/searchinvitemember');
INSERT INTO `ms_project_auth_node` VALUES (3266, 2, 'project/task_stages/index');
INSERT INTO `ms_project_auth_node` VALUES (3267, 2, 'project/task_stages/tasks');
INSERT INTO `ms_project_auth_node` VALUES (3268, 2, 'project/task_stages/sort');
INSERT INTO `ms_project_auth_node` VALUES (3269, 2, 'project/task_stages_template/index');
INSERT INTO `ms_project_auth_node` VALUES (3270, 2, 'project/department/index');
INSERT INTO `ms_project_auth_node` VALUES (3271, 2, 'project/department/read');
INSERT INTO `ms_project_auth_node` VALUES (3272, 2, 'project/department_member/index');
INSERT INTO `ms_project_auth_node` VALUES (3273, 2, 'project/department_member/searchinvitemember');
INSERT INTO `ms_project_auth_node` VALUES (3274, 2, 'project/project/selflist');
INSERT INTO `ms_project_auth_node` VALUES (3275, 2, 'project/project/save');
INSERT INTO `ms_project_auth_node` VALUES (3276, 2, 'project/task/selflist');
INSERT INTO `ms_project_auth_node` VALUES (6294, 1, 'project/index');
INSERT INTO `ms_project_auth_node` VALUES (6295, 1, 'project/index/info');
INSERT INTO `ms_project_auth_node` VALUES (6296, 1, 'project/index/index');
INSERT INTO `ms_project_auth_node` VALUES (6297, 1, 'project/index/systemconfig');
INSERT INTO `ms_project_auth_node` VALUES (6298, 1, 'project/index/editpersonal');
INSERT INTO `ms_project_auth_node` VALUES (6299, 1, 'project/index/uploadavatar');
INSERT INTO `ms_project_auth_node` VALUES (6300, 1, 'project/index/changecurrentorganization');
INSERT INTO `ms_project_auth_node` VALUES (6301, 1, 'project/index/editpassword');
INSERT INTO `ms_project_auth_node` VALUES (6302, 1, 'project/index/uploadimg');
INSERT INTO `ms_project_auth_node` VALUES (6303, 1, 'project/account');
INSERT INTO `ms_project_auth_node` VALUES (6304, 1, 'project/account/index');
INSERT INTO `ms_project_auth_node` VALUES (6305, 1, 'project/account/auth');
INSERT INTO `ms_project_auth_node` VALUES (6306, 1, 'project/account/add');
INSERT INTO `ms_project_auth_node` VALUES (6307, 1, 'project/account/edit');
INSERT INTO `ms_project_auth_node` VALUES (6308, 1, 'project/account/del');
INSERT INTO `ms_project_auth_node` VALUES (6309, 1, 'project/account/forbid');
INSERT INTO `ms_project_auth_node` VALUES (6310, 1, 'project/account/resume');
INSERT INTO `ms_project_auth_node` VALUES (6311, 1, 'project/account/read');
INSERT INTO `ms_project_auth_node` VALUES (6312, 1, 'project/organization');
INSERT INTO `ms_project_auth_node` VALUES (6313, 1, 'project/organization/index');
INSERT INTO `ms_project_auth_node` VALUES (6314, 1, 'project/organization/save');
INSERT INTO `ms_project_auth_node` VALUES (6315, 1, 'project/organization/read');
INSERT INTO `ms_project_auth_node` VALUES (6316, 1, 'project/organization/edit');
INSERT INTO `ms_project_auth_node` VALUES (6317, 1, 'project/organization/delete');
INSERT INTO `ms_project_auth_node` VALUES (6318, 1, 'project/auth');
INSERT INTO `ms_project_auth_node` VALUES (6319, 1, 'project/auth/index');
INSERT INTO `ms_project_auth_node` VALUES (6320, 1, 'project/auth/add');
INSERT INTO `ms_project_auth_node` VALUES (6321, 1, 'project/auth/edit');
INSERT INTO `ms_project_auth_node` VALUES (6322, 1, 'project/auth/forbid');
INSERT INTO `ms_project_auth_node` VALUES (6323, 1, 'project/auth/resume');
INSERT INTO `ms_project_auth_node` VALUES (6324, 1, 'project/auth/del');
INSERT INTO `ms_project_auth_node` VALUES (6325, 1, 'project/auth/apply');
INSERT INTO `ms_project_auth_node` VALUES (6326, 1, 'project/auth/setdefault');
INSERT INTO `ms_project_auth_node` VALUES (6327, 1, 'project/notify');
INSERT INTO `ms_project_auth_node` VALUES (6328, 1, 'project/notify/index');
INSERT INTO `ms_project_auth_node` VALUES (6329, 1, 'project/notify/noreads');
INSERT INTO `ms_project_auth_node` VALUES (6330, 1, 'project/notify/read');
INSERT INTO `ms_project_auth_node` VALUES (6331, 1, 'project/notify/delete');
INSERT INTO `ms_project_auth_node` VALUES (6332, 1, 'project/notify/setreadied');
INSERT INTO `ms_project_auth_node` VALUES (6333, 1, 'project/notify/batchdel');
INSERT INTO `ms_project_auth_node` VALUES (6334, 1, 'project/department_member');
INSERT INTO `ms_project_auth_node` VALUES (6335, 1, 'project/department_member/index');
INSERT INTO `ms_project_auth_node` VALUES (6336, 1, 'project/department_member/searchinvitemember');
INSERT INTO `ms_project_auth_node` VALUES (6337, 1, 'project/department_member/invitemember');
INSERT INTO `ms_project_auth_node` VALUES (6338, 1, 'project/department_member/removemember');
INSERT INTO `ms_project_auth_node` VALUES (6339, 1, 'project/department_member/detail');
INSERT INTO `ms_project_auth_node` VALUES (6340, 1, 'project/department_member/uploadfile');
INSERT INTO `ms_project_auth_node` VALUES (6341, 1, 'project/menu');
INSERT INTO `ms_project_auth_node` VALUES (6342, 1, 'project/menu/menu');
INSERT INTO `ms_project_auth_node` VALUES (6343, 1, 'project/menu/menuadd');
INSERT INTO `ms_project_auth_node` VALUES (6344, 1, 'project/menu/menuedit');
INSERT INTO `ms_project_auth_node` VALUES (6345, 1, 'project/menu/menuforbid');
INSERT INTO `ms_project_auth_node` VALUES (6346, 1, 'project/menu/menuresume');
INSERT INTO `ms_project_auth_node` VALUES (6347, 1, 'project/menu/menudel');
INSERT INTO `ms_project_auth_node` VALUES (6348, 1, 'project/node');
INSERT INTO `ms_project_auth_node` VALUES (6349, 1, 'project/node/index');
INSERT INTO `ms_project_auth_node` VALUES (6350, 1, 'project/node/alllist');
INSERT INTO `ms_project_auth_node` VALUES (6351, 1, 'project/node/clear');
INSERT INTO `ms_project_auth_node` VALUES (6352, 1, 'project/node/save');
INSERT INTO `ms_project_auth_node` VALUES (6353, 1, 'project/project');
INSERT INTO `ms_project_auth_node` VALUES (6354, 1, 'project/project/index');
INSERT INTO `ms_project_auth_node` VALUES (6355, 1, 'project/project/selflist');
INSERT INTO `ms_project_auth_node` VALUES (6356, 1, 'project/project/save');
INSERT INTO `ms_project_auth_node` VALUES (6357, 1, 'project/project/read');
INSERT INTO `ms_project_auth_node` VALUES (6358, 1, 'project/project/edit');
INSERT INTO `ms_project_auth_node` VALUES (6359, 1, 'project/project/uploadcover');
INSERT INTO `ms_project_auth_node` VALUES (6360, 1, 'project/project/recycle');
INSERT INTO `ms_project_auth_node` VALUES (6361, 1, 'project/project/recovery');
INSERT INTO `ms_project_auth_node` VALUES (6362, 1, 'project/project/archive');
INSERT INTO `ms_project_auth_node` VALUES (6363, 1, 'project/project/recoveryarchive');
INSERT INTO `ms_project_auth_node` VALUES (6364, 1, 'project/project/quit');
INSERT INTO `ms_project_auth_node` VALUES (6365, 1, 'project/project/getlogbyselfproject');
INSERT INTO `ms_project_auth_node` VALUES (6366, 1, 'project/project_collect');
INSERT INTO `ms_project_auth_node` VALUES (6367, 1, 'project/project_collect/collect');
INSERT INTO `ms_project_auth_node` VALUES (6368, 1, 'project/project_member');
INSERT INTO `ms_project_auth_node` VALUES (6369, 1, 'project/project_member/index');
INSERT INTO `ms_project_auth_node` VALUES (6370, 1, 'project/project_member/searchinvitemember');
INSERT INTO `ms_project_auth_node` VALUES (6371, 1, 'project/project_member/invitemember');
INSERT INTO `ms_project_auth_node` VALUES (6372, 1, 'project/project_member/removemember');
INSERT INTO `ms_project_auth_node` VALUES (6373, 1, 'project/project_template');
INSERT INTO `ms_project_auth_node` VALUES (6374, 1, 'project/project_template/index');
INSERT INTO `ms_project_auth_node` VALUES (6375, 1, 'project/project_template/save');
INSERT INTO `ms_project_auth_node` VALUES (6376, 1, 'project/project_template/uploadcover');
INSERT INTO `ms_project_auth_node` VALUES (6377, 1, 'project/project_template/edit');
INSERT INTO `ms_project_auth_node` VALUES (6378, 1, 'project/project_template/delete');
INSERT INTO `ms_project_auth_node` VALUES (6379, 1, 'project/task');
INSERT INTO `ms_project_auth_node` VALUES (6380, 1, 'project/task/index');
INSERT INTO `ms_project_auth_node` VALUES (6381, 1, 'project/task/selflist');
INSERT INTO `ms_project_auth_node` VALUES (6382, 1, 'project/task/read');
INSERT INTO `ms_project_auth_node` VALUES (6383, 1, 'project/task/save');
INSERT INTO `ms_project_auth_node` VALUES (6384, 1, 'project/task/taskdone');
INSERT INTO `ms_project_auth_node` VALUES (6385, 1, 'project/task/assigntask');
INSERT INTO `ms_project_auth_node` VALUES (6386, 1, 'project/task/sort');
INSERT INTO `ms_project_auth_node` VALUES (6387, 1, 'project/task/createcomment');
INSERT INTO `ms_project_auth_node` VALUES (6388, 1, 'project/task/edit');
INSERT INTO `ms_project_auth_node` VALUES (6389, 1, 'project/task/like');
INSERT INTO `ms_project_auth_node` VALUES (6390, 1, 'project/task/star');
INSERT INTO `ms_project_auth_node` VALUES (6391, 1, 'project/task/recycle');
INSERT INTO `ms_project_auth_node` VALUES (6392, 1, 'project/task/recovery');
INSERT INTO `ms_project_auth_node` VALUES (6393, 1, 'project/task/delete');
INSERT INTO `ms_project_auth_node` VALUES (6394, 1, 'project/task/datetotalforproject');
INSERT INTO `ms_project_auth_node` VALUES (6395, 1, 'project/task/tasksources');
INSERT INTO `ms_project_auth_node` VALUES (6396, 1, 'project/task/tasklog');
INSERT INTO `ms_project_auth_node` VALUES (6397, 1, 'project/task/recyclebatch');
INSERT INTO `ms_project_auth_node` VALUES (6398, 1, 'project/task/setprivate');
INSERT INTO `ms_project_auth_node` VALUES (6399, 1, 'project/task/batchassigntask');
INSERT INTO `ms_project_auth_node` VALUES (6400, 1, 'project/task/tasktotags');
INSERT INTO `ms_project_auth_node` VALUES (6401, 1, 'project/task/settag');
INSERT INTO `ms_project_auth_node` VALUES (6402, 1, 'project/task/getlistbytasktag');
INSERT INTO `ms_project_auth_node` VALUES (6403, 1, 'project/task/savetaskworktime');
INSERT INTO `ms_project_auth_node` VALUES (6404, 1, 'project/task/edittaskworktime');
INSERT INTO `ms_project_auth_node` VALUES (6405, 1, 'project/task/deltaskworktime');
INSERT INTO `ms_project_auth_node` VALUES (6406, 1, 'project/task/uploadfile');
INSERT INTO `ms_project_auth_node` VALUES (6407, 1, 'project/task_member');
INSERT INTO `ms_project_auth_node` VALUES (6408, 1, 'project/task_member/index');
INSERT INTO `ms_project_auth_node` VALUES (6409, 1, 'project/task_member/searchinvitemember');
INSERT INTO `ms_project_auth_node` VALUES (6410, 1, 'project/task_member/invitemember');
INSERT INTO `ms_project_auth_node` VALUES (6411, 1, 'project/task_member/invitememberbatch');
INSERT INTO `ms_project_auth_node` VALUES (6412, 1, 'project/task_stages');
INSERT INTO `ms_project_auth_node` VALUES (6413, 1, 'project/task_stages/index');
INSERT INTO `ms_project_auth_node` VALUES (6414, 1, 'project/task_stages/tasks');
INSERT INTO `ms_project_auth_node` VALUES (6415, 1, 'project/task_stages/sort');
INSERT INTO `ms_project_auth_node` VALUES (6416, 1, 'project/task_stages/save');
INSERT INTO `ms_project_auth_node` VALUES (6417, 1, 'project/task_stages/edit');
INSERT INTO `ms_project_auth_node` VALUES (6418, 1, 'project/task_stages/delete');
INSERT INTO `ms_project_auth_node` VALUES (6419, 1, 'project/task_stages_template');
INSERT INTO `ms_project_auth_node` VALUES (6420, 1, 'project/task_stages_template/index');
INSERT INTO `ms_project_auth_node` VALUES (6421, 1, 'project/task_stages_template/save');
INSERT INTO `ms_project_auth_node` VALUES (6422, 1, 'project/task_stages_template/edit');
INSERT INTO `ms_project_auth_node` VALUES (6423, 1, 'project/task_stages_template/delete');
INSERT INTO `ms_project_auth_node` VALUES (6424, 1, 'project/file');
INSERT INTO `ms_project_auth_node` VALUES (6425, 1, 'project/file/index');
INSERT INTO `ms_project_auth_node` VALUES (6426, 1, 'project/file/read');
INSERT INTO `ms_project_auth_node` VALUES (6427, 1, 'project/file/uploadfiles');
INSERT INTO `ms_project_auth_node` VALUES (6428, 1, 'project/file/edit');
INSERT INTO `ms_project_auth_node` VALUES (6429, 1, 'project/file/recycle');
INSERT INTO `ms_project_auth_node` VALUES (6430, 1, 'project/file/recovery');
INSERT INTO `ms_project_auth_node` VALUES (6431, 1, 'project/file/delete');
INSERT INTO `ms_project_auth_node` VALUES (6432, 1, 'project/source_link');
INSERT INTO `ms_project_auth_node` VALUES (6433, 1, 'project/source_link/delete');
INSERT INTO `ms_project_auth_node` VALUES (6434, 1, 'project/invite_link');
INSERT INTO `ms_project_auth_node` VALUES (6435, 1, 'project/invite_link/save');
INSERT INTO `ms_project_auth_node` VALUES (6436, 1, 'project/task_tag');
INSERT INTO `ms_project_auth_node` VALUES (6437, 1, 'project/task_tag/index');
INSERT INTO `ms_project_auth_node` VALUES (6438, 1, 'project/task_tag/save');
INSERT INTO `ms_project_auth_node` VALUES (6439, 1, 'project/task_tag/edit');
INSERT INTO `ms_project_auth_node` VALUES (6440, 1, 'project/task_tag/delete');
INSERT INTO `ms_project_auth_node` VALUES (6441, 1, 'project/project_features');
INSERT INTO `ms_project_auth_node` VALUES (6442, 1, 'project/project_features/index');
INSERT INTO `ms_project_auth_node` VALUES (6443, 1, 'project/project_features/save');
INSERT INTO `ms_project_auth_node` VALUES (6444, 1, 'project/project_features/edit');
INSERT INTO `ms_project_auth_node` VALUES (6445, 1, 'project/project_features/delete');
INSERT INTO `ms_project_auth_node` VALUES (6446, 1, 'project/project_version');
INSERT INTO `ms_project_auth_node` VALUES (6447, 1, 'project/project_version/index');
INSERT INTO `ms_project_auth_node` VALUES (6448, 1, 'project/project_version/save');
INSERT INTO `ms_project_auth_node` VALUES (6449, 1, 'project/project_version/edit');
INSERT INTO `ms_project_auth_node` VALUES (6450, 1, 'project/project_version/changestatus');
INSERT INTO `ms_project_auth_node` VALUES (6451, 1, 'project/project_version/read');
INSERT INTO `ms_project_auth_node` VALUES (6452, 1, 'project/project_version/addversiontask');
INSERT INTO `ms_project_auth_node` VALUES (6453, 1, 'project/project_version/removeversiontask');
INSERT INTO `ms_project_auth_node` VALUES (6454, 1, 'project/project_version/delete');
INSERT INTO `ms_project_auth_node` VALUES (6455, 1, 'project/task_workflow/delete');
INSERT INTO `ms_project_auth_node` VALUES (6456, 1, 'project/task_workflow');
INSERT INTO `ms_project_auth_node` VALUES (6457, 1, 'project/task_workflow/index');
INSERT INTO `ms_project_auth_node` VALUES (6458, 1, 'project/task_workflow/save');
INSERT INTO `ms_project_auth_node` VALUES (6459, 1, 'project/task_workflow/edit');
INSERT INTO `ms_project_auth_node` VALUES (6460, 1, 'project/department');
INSERT INTO `ms_project_auth_node` VALUES (6461, 1, 'project/department/index');
INSERT INTO `ms_project_auth_node` VALUES (6462, 1, 'project/department/read');
INSERT INTO `ms_project_auth_node` VALUES (6463, 1, 'project/department/save');
INSERT INTO `ms_project_auth_node` VALUES (6464, 1, 'project/department/edit');
INSERT INTO `ms_project_auth_node` VALUES (6465, 1, 'project/department/delete');
INSERT INTO `ms_project_auth_node` VALUES (6466, 1, 'project');

-- ----------------------------
-- Table structure for ms_project_collection
-- ----------------------------
DROP TABLE IF EXISTS `ms_project_collection`;
CREATE TABLE `ms_project_collection`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `project_code` bigint NULL DEFAULT 0 COMMENT '项目id',
  `member_code` bigint NULL DEFAULT 0 COMMENT '成员id',
  `create_time` bigint NULL DEFAULT 0 COMMENT '加入时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 51 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '项目-收藏表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_project_collection
-- ----------------------------

-- ----------------------------
-- Table structure for ms_project_log
-- ----------------------------
DROP TABLE IF EXISTS `ms_project_log`;
CREATE TABLE `ms_project_log`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `member_code` bigint NULL DEFAULT 0 COMMENT '操作人id',
  `content` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '操作内容',
  `remark` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `type` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT 'create' COMMENT '操作类型',
  `create_time` bigint NULL DEFAULT NULL COMMENT '添加时间',
  `source_code` bigint NULL DEFAULT 0 COMMENT '任务id',
  `action_type` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '场景类型',
  `to_member_code` bigint NULL DEFAULT 0,
  `is_comment` tinyint(1) NULL DEFAULT 0 COMMENT '是否评论，0：否',
  `project_code` bigint NULL DEFAULT NULL,
  `icon` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `is_robot` tinyint(1) NULL DEFAULT 0 COMMENT '是否机器人',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `member_code`(`member_code` ASC) USING BTREE,
  INDEX `source_code`(`source_code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5089 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '项目日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ms_project_log
-- ----------------------------
INSERT INTO `ms_project_log` VALUES (5086, 1015, '研发1', '创建任务', 'create', 1745944330005, 12376, 'task', 0, 0, 13047, 'plus', 0);
INSERT INTO `ms_project_log` VALUES (5087, 1015, 'wqeq', '创建任务', 'create', 1745947016779, 12377, 'task', 1015, 0, 13047, 'plus', 0);
INSERT INTO `ms_project_log` VALUES (5088, 1015, '完美！', '完美！', 'createComment', 1746769407276, 12366, 'task', 0, 1, 13047, 'plus', 0);

-- ----------------------------
-- Table structure for ms_project_member
-- ----------------------------
DROP TABLE IF EXISTS `ms_project_member`;
CREATE TABLE `ms_project_member`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `project_code` bigint NULL DEFAULT NULL COMMENT '项目id',
  `member_code` bigint NULL DEFAULT NULL COMMENT '成员id',
  `join_time` bigint NULL DEFAULT NULL COMMENT '加入时间',
  `is_owner` bigint NULL DEFAULT 0 COMMENT '拥有者',
  `authorize` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '角色',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unique`(`project_code` ASC, `member_code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 44 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '项目-成员表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_project_member
-- ----------------------------
INSERT INTO `ms_project_member` VALUES (37, 13043, 1015, 1744888400499, 1015, '');
INSERT INTO `ms_project_member` VALUES (38, 13044, 1015, 1744899261531, 1015, '');
INSERT INTO `ms_project_member` VALUES (39, 13045, 1015, 1744958143289, 1015, '');
INSERT INTO `ms_project_member` VALUES (40, 13046, 1015, 1744958683253, 1015, '');
INSERT INTO `ms_project_member` VALUES (41, 13047, 1015, 1745489790020, 1015, '');
INSERT INTO `ms_project_member` VALUES (42, 13048, 1015, 1747408522072, 1015, '');
INSERT INTO `ms_project_member` VALUES (43, 13049, 1015, 1747408830841, 1015, '');

-- ----------------------------
-- Table structure for ms_project_menu
-- ----------------------------
DROP TABLE IF EXISTS `ms_project_menu`;
CREATE TABLE `ms_project_menu`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '父id',
  `title` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `icon` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '菜单图标',
  `url` varchar(400) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '链接',
  `file_path` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '文件路径',
  `params` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '链接参数',
  `node` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '#' COMMENT '权限节点',
  `sort` int UNSIGNED NULL DEFAULT 0 COMMENT '菜单排序',
  `status` tinyint UNSIGNED NULL DEFAULT 1 COMMENT '状态(0:禁用,1:启用)',
  `create_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
  `is_inner` tinyint(1) NULL DEFAULT 0 COMMENT '是否内页',
  `values` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '参数默认值',
  `show_slider` tinyint(1) NULL DEFAULT 1 COMMENT '是否显示侧栏',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 176 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '项目菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ms_project_menu
-- ----------------------------
INSERT INTO `ms_project_menu` VALUES (120, 0, '工作台', 'appstore-o', 'home', 'home', ':org', '#', 0, 1, 0, 0, '', 0);
INSERT INTO `ms_project_menu` VALUES (121, 0, '项目管理', 'project', '#', '#', '', '#', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (122, 121, '项目列表', 'branches', '#', '#', '', '#', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (124, 0, '系统设置', 'setting', '#', '#', '', '#', 100, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (125, 124, '成员管理', 'unlock', '#', '#', '', '#', 10, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (126, 125, '账号列表', '', 'system/account', 'system/account', '', 'project/account/index', 10, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (127, 122, '我的组织', '', 'organization', 'organization', '', 'project/organization/index', 30, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (130, 125, '访问授权', '', 'system/account/auth', 'system/account/auth', '', 'project/auth/index', 20, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (131, 125, '授权页面', '', 'system/account/apply', 'system/account/apply', ':id', 'project/auth/apply', 30, 1, 0, 1, '', 1);
INSERT INTO `ms_project_menu` VALUES (138, 121, '消息提醒', 'info-circle-o', '#', '#', '', '#', 30, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (139, 138, '站内消息', '', 'notify/notice', 'notify/notice', '', 'project/notify/index', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (140, 138, '系统公告', '', 'notify/system', 'notify/system', '', 'project/notify/index', 10, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (143, 124, '系统管理', 'appstore', '#', '#', '', '#', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (144, 143, '菜单路由', '', 'system/config/menu', 'system/config/menu', '', 'project/menu/menuadd', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (145, 143, '访问节点', '', 'system/config/node', 'system/config/node', '', 'project/node/save', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (148, 124, '个人管理', 'user', '#', '#', '', '#', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (149, 148, '个人设置', '', 'account/setting/base', 'account/setting/base', '', 'project/index/editpersonal', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (150, 148, '安全设置', '', 'account/setting/security', 'account/setting/security', '', 'project/index/editpersonal', 0, 1, 0, 1, '', 1);
INSERT INTO `ms_project_menu` VALUES (151, 122, '我的项目', '', 'project/list', 'project/list', ':type', 'project/project/index', 0, 1, 0, 0, 'my', 1);
INSERT INTO `ms_project_menu` VALUES (152, 122, '回收站', '', 'project/recycle', 'project/recycle', '', 'project/project/index', 20, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (153, 121, '项目空间', 'heat-map', 'project/space/task', 'project/space/task', ':code', '#', 20, 1, 0, 1, '', 1);
INSERT INTO `ms_project_menu` VALUES (154, 153, '任务详情', '', 'project/space/task/:code/detail', 'project/space/taskdetail', ':code', 'project/task/read', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (155, 122, '我的收藏', '', 'project/list', 'project/list', ':type', 'project/project/index', 10, 1, 0, 0, 'collect', 1);
INSERT INTO `ms_project_menu` VALUES (156, 121, '基础设置', 'experiment', '#', '#', '', '#', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (157, 156, '项目模板', '', 'project/template', 'project/template', '', 'project/project_template/index', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (158, 156, '项目列表模板', '', 'project/template/taskStages', 'project/template/taskStages', ':code', 'project/task_stages_template/index', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (159, 122, '已归档项目', '', 'project/archive', 'project/archive', '', 'project/project/index', 10, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (160, 0, '团队成员', 'team', '#', '#', '', '#', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (161, 153, '项目概况', '', 'project/space/overview', 'project/space/overview', ':code', 'project/index/info', 20, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (162, 153, '项目文件', '', 'project/space/files', 'project/space/files', ':code', 'project/index/info', 10, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (163, 122, '项目分析', '', 'project/analysis', 'project/analysis', '', 'project/index/info', 5, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menu` VALUES (164, 160, '团队成员', '', '#', '#', '', '#', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (166, 164, '团队成员', '', 'members', 'members', '', 'project/department/index', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (167, 164, '成员信息', '', 'members/profile', 'members/profile', ':code', 'project/department/read', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menu` VALUES (168, 153, '版本管理', '', 'project/space/features', 'project/space/features', ':code', 'project/index/info', 20, 1, 0, 1, '', 0);

-- ----------------------------
-- Table structure for ms_project_node
-- ----------------------------
DROP TABLE IF EXISTS `ms_project_node`;
CREATE TABLE `ms_project_node`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `node` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '节点代码',
  `title` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '节点标题',
  `is_menu` tinyint UNSIGNED NULL DEFAULT 0 COMMENT '是否可设置为菜单',
  `is_auth` tinyint UNSIGNED NULL DEFAULT 1 COMMENT '是否启动RBAC权限控制',
  `is_login` tinyint UNSIGNED NULL DEFAULT 1 COMMENT '是否启动登录控制',
  `create_at` bigint NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `index_system_node_node`(`node` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 641 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '项目端节点表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_project_node
-- ----------------------------
INSERT INTO `ms_project_node` VALUES (360, 'project', '项目管理模块', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (361, 'project/index/info', '详情', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (362, 'project/index', '基础版块', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (363, 'project/index/index', '框架布局', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (364, 'project/index/systemconfig', '系统信息', 0, 0, 0, 1673277965322);
INSERT INTO `ms_project_node` VALUES (365, 'project/index/editpersonal', '修改个人资料', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (366, 'project/index/uploadavatar', '上传头像', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (370, 'project/account', '账号管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (371, 'project/account/index', '账号列表', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (372, 'project/organization/index', '组织列表', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (373, 'project/organization/save', '创建组织', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (374, 'project/organization/read', '组织信息', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (375, 'project/organization/edit', '编辑组织', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (376, 'project/organization/delete', '删除组织', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (377, 'project/organization', '组织管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (388, 'project/auth/index', '权限列表', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (389, 'project/auth/add', '添加权限角色', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (390, 'project/auth/edit', '编辑权限', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (391, 'project/auth/forbid', '禁用权限', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (392, 'project/auth/resume', '启用权限', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (393, 'project/auth/del', '删除权限', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (394, 'project/auth', '访问授权', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (395, 'project/auth/apply', '应用权限', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (396, 'project/notify/index', '通知列表', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (397, 'project/notify/noreads', '未读通知', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (399, 'project/notify/read', '通知信息', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (401, 'project/notify/delete', '删除通知', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (402, 'project/notify', '通知管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (434, 'project/account/auth', '授权管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (435, 'project/account/add', '添加账号', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (436, 'project/account/edit', '编辑账号', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (437, 'project/account/del', '删除账号', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (438, 'project/account/forbid', '禁用账号', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (439, 'project/account/resume', '启用账号', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (498, 'project/notify/setreadied', '设置已读', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (499, 'project/notify/batchdel', '批量删除', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (500, 'project/auth/setdefault', '设置默认权限', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (501, 'project/department', '部门管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (502, 'project/department/index', '部门列表', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (503, 'project/department/read', '部门信息', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (504, 'project/department/save', '创建部门', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (505, 'project/department/edit', '编辑部门', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (506, 'project/department/delete', '删除部门', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (507, 'project/department_member', '部门成员管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (508, 'project/department_member/index', '部门成员列表', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (509, 'project/department_member/searchinvitemember', '搜索部门成员', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (510, 'project/department_member/invitemember', '添加部门成员', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (511, 'project/department_member/removemember', '移除部门成员', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (512, 'project/index/changecurrentorganization', '切换当前组织', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (513, 'project/index/editpassword', '修改密码', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (514, 'project/index/uploadimg', '上传图片', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (515, 'project/menu', '菜单管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (516, 'project/menu/menu', '菜单列表', 0, 0, 0, 1673277965322);
INSERT INTO `ms_project_node` VALUES (517, 'project/menu/menuadd', '添加菜单', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (518, 'project/menu/menuedit', '编辑菜单', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (519, 'project/menu/menuforbid', '禁用菜单', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (520, 'project/menu/menuresume', '启用菜单', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (521, 'project/menu/menudel', '删除菜单', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (522, 'project/node', '节点管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (523, 'project/node/index', '节点列表', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (524, 'project/node/alllist', '全部节点列表', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (525, 'project/node/clear', '清理节点', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (526, 'project/node/save', '编辑节点', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (527, 'project/project', '项目管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (528, 'project/project/index', '项目列表', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (529, 'project/project/selflist', '个人项目列表', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (530, 'project/project/save', '创建项目', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (531, 'project/project/read', '项目信息', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (532, 'project/project/edit', '编辑项目', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (533, 'project/project/uploadcover', '上传项目封面', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (534, 'project/project/recycle', '项目放入回收站', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (535, 'project/project/recovery', '恢复项目', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (536, 'project/project/archive', '归档项目', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (537, 'project/project/recoveryarchive', '取消归档项目', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (538, 'project/project/quit', '退出项目', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (539, 'project/project_collect', '项目收藏管理', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (540, 'project/project_collect/collect', '收藏项目', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (541, 'project/project_member', '项目成员管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (542, 'project/project_member/index', '项目成员列表', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (543, 'project/project_member/searchinvitemember', '搜索项目成员', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (544, 'project/project_member/invitemember', '邀请项目成员', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (545, 'project/project_template', '项目模板管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (546, 'project/project_template/index', '项目模板列表', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (547, 'project/project_template/save', '创建项目模板', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (548, 'project/project_template/uploadcover', '上传项目模板封面', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (549, 'project/project_template/edit', '编辑项目模板', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (550, 'project/project_template/delete', '删除项目模板', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (551, 'project/task/index', '任务列表', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (552, 'project/task/selflist', '个人任务列表', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (553, 'project/task/read', '任务信息', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (554, 'project/task/save', '创建任务', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (555, 'project/task/taskdone', '更改任务状态', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (556, 'project/task/assigntask', '指派任务执行者', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (557, 'project/task/sort', '任务排序', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (558, 'project/task/createcomment', '发表任务评论', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (559, 'project/task/edit', '编辑任务', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (560, 'project/task/like', '点赞任务', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (561, 'project/task/star', '收藏任务', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (562, 'project/task/recycle', '移动任务到回收站', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (563, 'project/task/recovery', '恢复任务', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (564, 'project/task/delete', '删除任务', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (565, 'project/task', '任务管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (569, 'project/task_member/index', '任务成员列表', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (570, 'project/task_member/searchinvitemember', '搜索任务成员', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (571, 'project/task_member/invitemember', '添加任务成员', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (572, 'project/task_member/invitememberbatch', '批量添加任务成员', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (573, 'project/task_member', '任务成员管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (574, 'project/task_stages', '任务分组管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (575, 'project/task_stages/index', '任务分组列表', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (576, 'project/task_stages/tasks', '任务分组任务列表', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (577, 'project/task_stages/sort', '任务分组排序', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (578, 'project/task_stages/save', '添加任务分组', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (579, 'project/task_stages/edit', '编辑任务分组', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (580, 'project/task_stages/delete', '删除任务分组', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (581, 'project/task_stages_template/index', '任务分组模板列表', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (582, 'project/task_stages_template/save', '创建任务分组模板', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (583, 'project/task_stages_template/edit', '编辑任务分组模板', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (584, 'project/task_stages_template/delete', '删除任务分组模板', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (585, 'project/task_stages_template', '任务分组模板管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (587, 'project/project_member/removemember', '移除项目成员', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (588, 'project/task/datetotalforproject', '任务统计', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (589, 'project/task/tasksources', '任务资源列表', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (590, 'project/file', '文件管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (591, 'project/file/index', '文件列表', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (592, 'project/file/read', '文件详情', 0, 0, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (593, 'project/file/uploadfiles', '上传文件', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (594, 'project/file/edit', '编辑文件', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (595, 'project/file/recycle', '文件移至回收站', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (596, 'project/file/recovery', '恢复文件', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (597, 'project/file/delete', '删除文件', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (598, 'project/project/getlogbyselfproject', '项目概况', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (599, 'project/source_link', '资源关联管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (600, 'project/source_link/delete', '取消关联', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (601, 'project/task/tasklog', '任务动态', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (602, 'project/task/recyclebatch', '批量移动任务到回收站', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (603, 'project/invite_link', '邀请链接管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (604, 'project/invite_link/save', '创建邀请链接', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (605, 'project/task/setprivate', '设置任务隐私模式', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (606, 'project/account/read', '账号信息', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (607, 'project/task/batchassigntask', '批量指派任务', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (608, 'project/task/tasktotags', '任务标签', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (609, 'project/task/settag', '设置任务标签', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (610, 'project/task_tag', '任务标签管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (611, 'project/task_tag/index', '任务标签列表', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (612, 'project/task_tag/save', '创建任务标签', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (613, 'project/task_tag/edit', '编辑任务标签', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (614, 'project/task_tag/delete', '删除任务标签', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (615, 'project/project_features', '项目版本库管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (616, 'project/project_features/index', '版本库列表', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (617, 'project/project_features/save', '添加版本库', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (618, 'project/project_features/edit', '编辑版本库', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (619, 'project/project_features/delete', '删除版本库', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (620, 'project/project_version', '项目版本管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (621, 'project/project_version/index', '项目版本列表', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (622, 'project/project_version/save', '添加项目版本', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (623, 'project/project_version/edit', '编辑项目版本', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (624, 'project/project_version/changestatus', '更改项目版本状态', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (625, 'project/project_version/read', '项目版本详情', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (626, 'project/project_version/addversiontask', '关联项目版本任务', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (627, 'project/project_version/removeversiontask', '移除项目版本任务', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (628, 'project/project_version/delete', '删除项目版本', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (629, 'project/task/getlistbytasktag', '标签任务列表', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (630, 'project/task_workflow', '任务流转管理', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (631, 'project/task_workflow/index', '任务流转列表', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (632, 'project/task_workflow/save', '添加任务流转', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (633, 'project/task_workflow/edit', '编辑任务流转', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (634, 'project/task_workflow/delete', '删除任务流转', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (635, 'project/department_member/detail', '部门成员详情', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (636, 'project/department_member/uploadfile', '上传头像', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (637, 'project/task/savetaskworktime', '保存任务流转', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (638, 'project/task/edittaskworktime', '编辑任务流转', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (639, 'project/task/deltaskworktime', '删除任务流转', 0, 1, 1, 1673277965322);
INSERT INTO `ms_project_node` VALUES (640, 'project/task/uploadfile', '上传文件', 0, 1, 1, 1673277965322);

-- ----------------------------
-- Table structure for ms_project_template
-- ----------------------------
DROP TABLE IF EXISTS `ms_project_template`;
CREATE TABLE `ms_project_template`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '类型名称',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '备注',
  `sort` tinyint NULL DEFAULT 0,
  `create_time` bigint NULL DEFAULT 0,
  `organization_code` bigint NULL DEFAULT NULL COMMENT '组织id',
  `cover` varchar(511) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '封面',
  `member_code` bigint NULL DEFAULT NULL COMMENT '创建人',
  `is_system` tinyint(1) NULL DEFAULT 0 COMMENT '系统默认',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '项目类型表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_project_template
-- ----------------------------
INSERT INTO `ms_project_template` VALUES (11, '产品进展', '适用于互联网产品人员对产品计划、跟进及发布管理', 0, 1670904236057, 17, 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fbpic.51yuansu.com%2Fpic3%2Fcover%2F01%2F91%2F92%2F5982adf6c88ea_610.jpg&refer=http%3A%2F%2Fbpic.51yuansu.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1673496114&t=956c5614481fedea97794e161deddb00', NULL, 1);
INSERT INTO `ms_project_template` VALUES (12, '需求管理', '适用于产品部门对需求的收集、评估及反馈管理', 0, 1670904236057, 17, 'https://img0.baidu.com/it/u=437485064,4277010738&fm=253&fmt=auto&app=138&f=JPEG?w=610&h=491', NULL, 1);
INSERT INTO `ms_project_template` VALUES (13, '机械制造', '适用于制造商对图纸设计及制造安装的工作流程管理', 0, 1670904236057, 17, 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fbpic.51yuansu.com%2Fpic2%2Fcover%2F00%2F38%2F93%2F5812ca7a24020_610.jpg&refer=http%3A%2F%2Fbpic.51yuansu.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1673496114&t=6d03fb91b230058fc43f1b7ae00f73e3', NULL, 1);
INSERT INTO `ms_project_template` VALUES (19, 'OKR 管理', '适用于团队的 OKR 管理', 0, 1670904236057, 17, 'https://img2.baidu.com/it/u=2241642503,1613686234&fm=253&fmt=auto&app=138&f=JPEG?w=603&h=500', 1015, 0);

-- ----------------------------
-- Table structure for ms_source_link
-- ----------------------------
DROP TABLE IF EXISTS `ms_source_link`;
CREATE TABLE `ms_source_link`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `source_type` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '资源类型',
  `source_code` bigint NULL DEFAULT NULL COMMENT '资源编号',
  `link_type` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '关联类型',
  `link_code` bigint NULL DEFAULT NULL COMMENT '关联编号',
  `organization_code` bigint NULL DEFAULT NULL COMMENT '组织编码',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `create_time` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建时间',
  `sort` int NULL DEFAULT 0 COMMENT '排序',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '资源关联表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_source_link
-- ----------------------------
INSERT INTO `ms_source_link` VALUES (11, 'file', 51, 'task', 12363, 14, 1015, '1746634393480', 0);
INSERT INTO `ms_source_link` VALUES (12, 'file', 52, 'task', 12363, 14, 1015, '1746634616839', 0);
INSERT INTO `ms_source_link` VALUES (13, 'file', 53, 'task', 12363, 14, 1015, '1746634716867', 0);

-- ----------------------------
-- Table structure for ms_task
-- ----------------------------
DROP TABLE IF EXISTS `ms_task`;
CREATE TABLE `ms_task`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `project_code` bigint NOT NULL COMMENT '项目编号',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `pri` tinyint UNSIGNED NULL DEFAULT 0 COMMENT '紧急程度',
  `execute_status` tinyint NULL DEFAULT NULL COMMENT '执行状态',
  `description` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '详情',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建人',
  `done_by` bigint NULL DEFAULT NULL COMMENT '完成人',
  `done_time` bigint NULL DEFAULT NULL COMMENT '完成时间',
  `create_time` bigint NULL DEFAULT NULL COMMENT '创建日期',
  `assign_to` bigint NULL DEFAULT NULL COMMENT '指派给谁',
  `deleted` tinyint(1) NULL DEFAULT 0 COMMENT '回收站',
  `stage_code` int NULL DEFAULT NULL COMMENT '任务列表',
  `task_tag` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '任务标签',
  `done` tinyint NULL DEFAULT 0 COMMENT '是否完成',
  `begin_time` bigint NULL DEFAULT NULL COMMENT '开始时间',
  `end_time` bigint NULL DEFAULT NULL COMMENT '截止时间',
  `remind_time` bigint NULL DEFAULT NULL COMMENT '提醒时间',
  `pcode` bigint NULL DEFAULT NULL COMMENT '父任务id',
  `sort` int NULL DEFAULT 0 COMMENT '排序',
  `like` int NULL DEFAULT 0 COMMENT '点赞数',
  `star` int NULL DEFAULT 0 COMMENT '收藏数',
  `deleted_time` bigint NULL DEFAULT NULL COMMENT '删除时间',
  `private` tinyint(1) NULL DEFAULT 0 COMMENT '是否隐私模式',
  `id_num` int NULL DEFAULT 1 COMMENT '任务id编号',
  `path` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '上级任务路径',
  `schedule` int NULL DEFAULT 0 COMMENT '进度百分比',
  `version_code` bigint NULL DEFAULT 0 COMMENT '版本id',
  `features_code` bigint NULL DEFAULT 0 COMMENT '版本库id',
  `work_time` int NULL DEFAULT 0 COMMENT '预估工时',
  `status` tinyint NULL DEFAULT 0 COMMENT '执行状态。0：未开始，1：已完成，2：进行中，3：挂起，4：测试中',
  PRIMARY KEY (`id`, `project_code`) USING BTREE,
  INDEX `stage_code`(`stage_code` ASC) USING BTREE,
  INDEX `project_code`(`project_code` ASC) USING BTREE,
  INDEX `pcode`(`pcode` ASC) USING BTREE,
  INDEX `sort`(`sort` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12378 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '任务表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_task
-- ----------------------------
INSERT INTO `ms_task` VALUES (12363, 13047, '任务1', 0, 0, '', 1015, 0, 0, 1745736914355, 1015, 0, 77, '', 0, 1745736914355, 1745909714355, 0, 0, 1, 0, 0, 0, 0, 1, '', 0, 0, 0, 0, 0);
INSERT INTO `ms_task` VALUES (12364, 13047, '111', 0, 0, '', 1015, 0, 0, 1745836572090, 1015, 0, 77, '', 0, 1745836572090, 1746009372090, 0, 0, 2, 0, 0, 0, 0, 2, '', 0, 0, 0, 0, 0);
INSERT INTO `ms_task` VALUES (12365, 13047, '222', 0, 0, '', 1015, 0, 0, 1745836599046, 1015, 0, 77, '', 0, 1745836599046, 1746009399046, 0, 0, 4, 0, 0, 0, 0, 3, '', 0, 0, 0, 0, 0);
INSERT INTO `ms_task` VALUES (12366, 13047, '111', 0, 0, '', 1015, 0, 0, 1745836644416, 1015, 0, 77, '', 0, 1745836644416, 1746009444416, 0, 0, 5, 0, 0, 0, 0, 4, '', 0, 0, 0, 0, 0);
INSERT INTO `ms_task` VALUES (12367, 13047, '任务1', 0, 0, '', 1015, 0, 0, 1745837895317, 1015, 0, 79, '', 0, 1745837895317, 1746010695317, 0, 0, 3, 0, 0, 0, 0, 5, '', 0, 0, 0, 0, 0);
INSERT INTO `ms_task` VALUES (12368, 13047, '任务2', 0, 0, '', 1015, 0, 0, 1745837915484, 1015, 0, 79, '', 0, 1745837915484, 1746010715484, 0, 0, 9, 0, 0, 0, 0, 6, '', 0, 0, 0, 0, 0);
INSERT INTO `ms_task` VALUES (12369, 13047, '任务3', 0, 0, '', 1015, 0, 0, 1745837941271, 1015, 0, 79, '', 0, 1745837941271, 1746010741271, 0, 0, 10, 0, 0, 0, 0, 7, '', 0, 0, 0, 0, 0);
INSERT INTO `ms_task` VALUES (12370, 13047, '任务111', 0, 0, '', 1015, 0, 0, 1745838208212, 1015, 0, 79, '', 0, 1745838208212, 1746011008212, 0, 0, 2, 0, 0, 0, 0, 8, '', 0, 0, 0, 0, 0);
INSERT INTO `ms_task` VALUES (12371, 13047, '研发1', 0, 0, '', 1015, 0, 0, 1745838639015, 1015, 0, 82, '', 0, 1745838639015, 1746011439015, 0, 0, 1, 0, 0, 0, 0, 9, '', 0, 0, 0, 0, 0);
INSERT INTO `ms_task` VALUES (12372, 13047, '内测1', 0, 0, '', 1015, 0, 0, 1745842791948, 1015, 0, 81, '', 0, 1745842791948, 1746015591948, 0, 0, 1, 0, 0, 0, 0, 10, '', 0, 0, 0, 0, 0);
INSERT INTO `ms_task` VALUES (12373, 13047, '内测2', 0, 0, '', 1015, 0, 0, 1745842798458, 1015, 0, 81, '', 0, 1745842798458, 1746015598458, 0, 0, 2, 0, 0, 0, 0, 11, '', 0, 0, 0, 0, 0);
INSERT INTO `ms_task` VALUES (12374, 13047, '内测3', 0, 0, '', 1015, 0, 0, 1745842807800, 1015, 0, 81, '', 0, 1745842807800, 1746015607800, 0, 0, 5, 0, 0, 0, 0, 12, '', 0, 0, 0, 0, 0);
INSERT INTO `ms_task` VALUES (12375, 13047, '评估确认1', 0, 0, '', 1015, 0, 0, 1745851908722, 1015, 0, 78, '', 0, 1745851908722, 1746024708722, 0, 0, 1, 0, 0, 0, 0, 13, '', 0, 0, 0, 0, 0);
INSERT INTO `ms_task` VALUES (12376, 13047, '研发1', 0, 0, '', 1015, 0, 0, 1745944329948, 1015, 0, 80, '', 0, 1745944329948, 1746117129948, 0, 0, 1, 0, 0, 0, 0, 14, '', 0, 0, 0, 0, 0);
INSERT INTO `ms_task` VALUES (12377, 13047, 'wqeq', 0, 0, '', 1015, 0, 0, 1745947016714, 1015, 0, 78, '', 0, 1745947016714, 1746119816714, 0, 0, 2, 0, 0, 0, 0, 15, '', 0, 0, 0, 0, 0);

-- ----------------------------
-- Table structure for ms_task_member
-- ----------------------------
DROP TABLE IF EXISTS `ms_task_member`;
CREATE TABLE `ms_task_member`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `task_code` bigint NULL DEFAULT 0 COMMENT '任务ID',
  `is_executor` tinyint(1) NULL DEFAULT 0 COMMENT '执行者',
  `member_code` bigint NULL DEFAULT NULL COMMENT '成员id',
  `join_time` bigint NULL DEFAULT NULL,
  `is_owner` tinyint(1) NULL DEFAULT 0 COMMENT '是否创建人',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 288 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '任务-成员表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_task_member
-- ----------------------------
INSERT INTO `ms_task_member` VALUES (273, 12363, 1, 1015, 1745736914355, 1);
INSERT INTO `ms_task_member` VALUES (274, 12364, 1, 1015, 1745836572090, 1);
INSERT INTO `ms_task_member` VALUES (275, 12365, 1, 1015, 1745836599046, 1);
INSERT INTO `ms_task_member` VALUES (276, 12366, 1, 1015, 1745836644416, 1);
INSERT INTO `ms_task_member` VALUES (277, 12367, 1, 1015, 1745837895317, 1);
INSERT INTO `ms_task_member` VALUES (278, 12368, 1, 1015, 1745837915484, 1);
INSERT INTO `ms_task_member` VALUES (279, 12369, 1, 1015, 1745837941271, 1);
INSERT INTO `ms_task_member` VALUES (280, 12370, 1, 1015, 1745838208212, 1);
INSERT INTO `ms_task_member` VALUES (281, 12371, 1, 1015, 1745838639015, 1);
INSERT INTO `ms_task_member` VALUES (282, 12372, 1, 1015, 1745842791948, 1);
INSERT INTO `ms_task_member` VALUES (283, 12373, 1, 1015, 1745842798458, 1);
INSERT INTO `ms_task_member` VALUES (284, 12374, 1, 1015, 1745842807800, 1);
INSERT INTO `ms_task_member` VALUES (285, 12375, 1, 1015, 1745851908722, 1);
INSERT INTO `ms_task_member` VALUES (286, 12376, 1, 1015, 1745944329948, 1);
INSERT INTO `ms_task_member` VALUES (287, 12377, 1, 1015, 1745947016714, 1);

-- ----------------------------
-- Table structure for ms_task_stage
-- ----------------------------
DROP TABLE IF EXISTS `ms_task_stage`;
CREATE TABLE `ms_task_stage`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '类型名称',
  `project_code` bigint NULL DEFAULT NULL COMMENT '项目id',
  `sort` int NULL DEFAULT 0 COMMENT '排序',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '备注',
  `create_time` bigint NULL DEFAULT NULL COMMENT '创建时间',
  `deleted` tinyint(1) NULL DEFAULT 0 COMMENT '删除标记',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 98 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '任务列表表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_task_stage
-- ----------------------------
INSERT INTO `ms_task_stage` VALUES (77, '需求收集', 13047, 1, '', 1745489790021, 0);
INSERT INTO `ms_task_stage` VALUES (78, '评估确认', 13047, 2, '', 1745489790022, 0);
INSERT INTO `ms_task_stage` VALUES (79, '需求暂缓', 13047, 3, '', 1745489790022, 0);
INSERT INTO `ms_task_stage` VALUES (80, '研发中', 13047, 4, '', 1745489790023, 0);
INSERT INTO `ms_task_stage` VALUES (81, '内测中', 13047, 5, '', 1745489790024, 0);
INSERT INTO `ms_task_stage` VALUES (82, '通知用户', 13047, 6, '', 1745489790024, 0);
INSERT INTO `ms_task_stage` VALUES (83, '已完成&归档', 13047, 7, '', 1745489790025, 0);
INSERT INTO `ms_task_stage` VALUES (84, '协议签订', 13048, 1, '', 1747408522073, 0);
INSERT INTO `ms_task_stage` VALUES (85, '图纸设计', 13048, 2, '', 1747408522075, 0);
INSERT INTO `ms_task_stage` VALUES (86, '评审及打样', 13048, 3, '', 1747408522075, 0);
INSERT INTO `ms_task_stage` VALUES (87, '构件采购', 13048, 4, '', 1747408522080, 0);
INSERT INTO `ms_task_stage` VALUES (88, '制造安装', 13048, 5, '', 1747408522080, 0);
INSERT INTO `ms_task_stage` VALUES (89, '内部检验', 13048, 6, '', 1747408522080, 0);
INSERT INTO `ms_task_stage` VALUES (90, '验收', 13048, 7, '', 1747408522081, 0);
INSERT INTO `ms_task_stage` VALUES (91, '需求收集', 13049, 1, '', 1747408830842, 0);
INSERT INTO `ms_task_stage` VALUES (92, '评估确认', 13049, 2, '', 1747408830843, 0);
INSERT INTO `ms_task_stage` VALUES (93, '需求暂缓', 13049, 3, '', 1747408830843, 0);
INSERT INTO `ms_task_stage` VALUES (94, '研发中', 13049, 4, '', 1747408830844, 0);
INSERT INTO `ms_task_stage` VALUES (95, '内测中', 13049, 5, '', 1747408830844, 0);
INSERT INTO `ms_task_stage` VALUES (96, '通知用户', 13049, 6, '', 1747408830844, 0);
INSERT INTO `ms_task_stage` VALUES (97, '已完成&归档', 13049, 7, '', 1747408830845, 0);

-- ----------------------------
-- Table structure for ms_task_stages_template
-- ----------------------------
DROP TABLE IF EXISTS `ms_task_stages_template`;
CREATE TABLE `ms_task_stages_template`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '类型名称',
  `project_template_code` int NULL DEFAULT 0 COMMENT '项目id',
  `create_time` bigint NULL DEFAULT NULL,
  `sort` int NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 84 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '任务列表模板表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_task_stages_template
-- ----------------------------
INSERT INTO `ms_task_stages_template` VALUES (61, '待处理', 19, 1670904236057, 1);
INSERT INTO `ms_task_stages_template` VALUES (62, '进行中', 19, 1670904236057, 0);
INSERT INTO `ms_task_stages_template` VALUES (63, '已完成', 19, 1670904236057, 0);
INSERT INTO `ms_task_stages_template` VALUES (65, '协议签订', 13, 1670904236057, 0);
INSERT INTO `ms_task_stages_template` VALUES (66, '图纸设计', 13, 1670904236057, 0);
INSERT INTO `ms_task_stages_template` VALUES (67, '评审及打样', 13, 1670904236057, 0);
INSERT INTO `ms_task_stages_template` VALUES (68, '构件采购', 13, 1670904236057, 0);
INSERT INTO `ms_task_stages_template` VALUES (69, '制造安装', 13, 1670904236057, 0);
INSERT INTO `ms_task_stages_template` VALUES (70, '内部检验', 13, 1670904236057, 0);
INSERT INTO `ms_task_stages_template` VALUES (71, '验收', 13, 1670904236057, 0);
INSERT INTO `ms_task_stages_template` VALUES (72, '需求收集', 12, 1670904236057, 0);
INSERT INTO `ms_task_stages_template` VALUES (73, '评估确认', 12, 1670904236057, 0);
INSERT INTO `ms_task_stages_template` VALUES (74, '需求暂缓', 12, 1670904236057, 0);
INSERT INTO `ms_task_stages_template` VALUES (75, '研发中', 12, 1670904236057, 0);
INSERT INTO `ms_task_stages_template` VALUES (76, '内测中', 12, 1670904236057, 0);
INSERT INTO `ms_task_stages_template` VALUES (77, '通知用户', 12, 1670904236057, 0);
INSERT INTO `ms_task_stages_template` VALUES (78, '已完成&归档', 12, 1670904236057, 0);
INSERT INTO `ms_task_stages_template` VALUES (79, '产品计划', 11, 1670904236057, 0);
INSERT INTO `ms_task_stages_template` VALUES (80, '即将发布', 11, 1670904236057, 0);
INSERT INTO `ms_task_stages_template` VALUES (81, '测试', 11, 1670904236057, 0);
INSERT INTO `ms_task_stages_template` VALUES (82, '准备发布', 11, 1670904236057, 0);
INSERT INTO `ms_task_stages_template` VALUES (83, '发布成功', 11, 1670904236057, 0);

-- ----------------------------
-- Table structure for ms_task_work_time
-- ----------------------------
DROP TABLE IF EXISTS `ms_task_work_time`;
CREATE TABLE `ms_task_work_time`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `task_code` bigint NULL DEFAULT 0 COMMENT '任务ID',
  `member_code` bigint NULL DEFAULT NULL COMMENT '成员id',
  `create_time` bigint NULL DEFAULT NULL,
  `content` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '描述',
  `begin_time` bigint NULL DEFAULT NULL COMMENT '开始时间',
  `num` int NULL DEFAULT 0 COMMENT '工时',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '任务工时表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of ms_task_work_time
-- ----------------------------
INSERT INTO `ms_task_work_time` VALUES (1, 12363, 1015, 0, '一些内容', 1745675700000, 12);
INSERT INTO `ms_task_work_time` VALUES (2, 12377, 1015, 0, '内容', 1745588340000, 12);
INSERT INTO `ms_task_work_time` VALUES (3, 12377, 1015, 0, 'dsad', 1744292400000, 5);
INSERT INTO `ms_task_work_time` VALUES (4, 12364, 1015, 0, '工作内容1', 1746912240000, 6);

SET FOREIGN_KEY_CHECKS = 1;
