CREATE TABLE `news` (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `category_id` int NOT NULL,
  `news_content` text NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT NULL,
  `deleted_at` TIMESTAMP DEFAULT NULL
);


