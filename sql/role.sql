use `operation`;

CREATE TABLE `tb_operation_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `role_name` varchar(50) NOT NULL DEFAULT '' COMMENT '角色名称',
  `parent_id` int(11) unsigned not null default 0 comment '父级角色id',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '状态  1:启用  2:禁用',
  `is_delete` tinyint(3) unsigned NOT NULL DEFAULT '2' COMMENT '是否删除  1:是  2:否',
  `insert_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '写入时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最近一次更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT="角色表";