use douyin;
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `user_name` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `register_at` datetime NOT NULL COMMENT '注册时间',
  `last_login` datetime NOT NULL COMMENT '最后一次登录时间',
  `login_num` int NOT NULL COMMENT '登录次数',
  `follow_count` int NOT NULL DEFAULT '0' COMMENT '关注总数',
  `follower_count` int NOT NULL DEFAULT '0' COMMENT '粉丝总数',
  `is_admin` tinyint NOT NULL DEFAULT '0' COMMENT '是否是管理员 0否  1是',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `created_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=10001 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';