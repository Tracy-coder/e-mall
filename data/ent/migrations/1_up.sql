CREATE TABLE `users` (
    `id` bigint unsigned NOT NULL COMMENT "primary key" AUTO_INCREMENT, 
    `created_at` timestamp NOT NULL COMMENT "created time", 
    `updated_at` timestamp NOT NULL COMMENT "last update time", 
    `status` tinyint unsigned NULL DEFAULT 1 COMMENT "status 1 normal 0 ban | 状态 1 正常 0 禁用",
    `username` varchar(255) NOT NULL COMMENT "user's login name | 登录名", 
    `password` varchar(255) NOT NULL COMMENT "password | 密码", 
    `nickname` varchar(255) NOT NULL COMMENT "nickname | 昵称", 
    `email` varchar(255) NULL COMMENT "email | 邮箱号", 
    `avatar` varchar(512) NULL DEFAULT '' COMMENT "avatar | 头像路径", 
    PRIMARY KEY (`id`), 
    UNIQUE INDEX `username` (`username`), 
    UNIQUE INDEX `nickname` (`nickname`), 
    UNIQUE INDEX `user_username_email` (`username`, `email`)) 
    CHARSET utf8mb4 COLLATE utf8mb4_bin;