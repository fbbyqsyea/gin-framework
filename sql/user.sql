-- 在operation库中创建tb_operation_user用户表
use `operation`;

-- 新建用户表
CREATE TABLE `tb_operation_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `account` varchar(50) NOT NULL DEFAULT '' COMMENT '账号',
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(16) NOT NULL DEFAULT '' COMMENT '密码加密盐值',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '状态  1:启用  2:禁用',
  `is_delete` tinyint(3) unsigned NOT NULL DEFAULT '2' COMMENT '是否删除  1:是  2:否',
  `insert_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '写入时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最近一次更新时间',
  PRIMARY KEY (`id`),
  KEY `tb_operation_user_account_IDX` (`account`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT="用户表";

-- 写入初始化数据
INSERT INTO operation.tb_operation_user
(id, account, username, password, salt, status, is_delete)
VALUES(1, 'admin', '超级管理员', 'a4ba57c03c548a762c4ab0bb03c0b377', 'sJtKCEpSaZrPjxcF', 1, 2);