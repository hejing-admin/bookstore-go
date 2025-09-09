CREATE TABLE `bookstore`.`user` (
     `id` INT AUTO_INCREMENT PRIMARY KEY COMMENT 'User ID',
     `username` VARCHAR(50) NOT NULL UNIQUE COMMENT 'Username',
     `password` VARCHAR(255) NOT NULL COMMENT 'Password (encrypted)',
     `email` VARCHAR(100) NOT NULL UNIQUE COMMENT 'Email address',
     `phone` VARCHAR(20) COMMENT 'Phone number',
     `avatar` VARCHAR(255) DEFAULT '' COMMENT 'Avatar URL',
     `is_admin` TINYINT(1) DEFAULT 0 COMMENT 'Is administrator (0: no, 1: yes)',
     `create_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
     `update_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update time',
     INDEX `idx_username` (`username`),
     INDEX `idx_phone` (`phone`),
     INDEX `idx_email` (`email`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='User information table';