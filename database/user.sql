CREATE TABLE `users` (`id` bigint AUTO_INCREMENT,`first_name` varchar(32),`last_name` varchar(32),`email` varchar(32),`password` varchar(255),`citizenship` varchar(32),`created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,`updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,PRIMARY KEY (`id`));