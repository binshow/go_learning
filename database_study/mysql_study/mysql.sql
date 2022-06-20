DROP TABLE IF EXISTS `t_user_info`;
CREATE TABLE `t_user_info` (
                               `user_id` bigint(20) NOT NULL COMMENT '用户Id',
                               `username` varchar(64) COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '用户名',
                               `account_id` bigint(20) NOT NULL,
                               `account_name` varchar(64) COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '用户名',
                               `phone` varchar(11) COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '手机号',
                               `gender` char(1) COLLATE utf8_bin NOT NULL COMMENT '性别 male 1； female 2；un_known or hide 0',
                               `avatar` varchar(255) COLLATE utf8_bin NOT NULL COMMENT '头像',
                               `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                               `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               PRIMARY KEY (`user_id`),
                               UNIQUE KEY `uniq_phone` (`phone`),
                               UNIQUE KEY `uniq_user_id` (`user_id`),
                               UNIQUE KEY `uniq_username` (`username`),
                               UNIQUE KEY `uniq_account_id` (`account_id`),
                               UNIQUE KEY `uniq_account_name` (`account_name`),
                               KEY `idx_create_time` (`create_time`),
                               KEY `idx_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='用户信息表';
