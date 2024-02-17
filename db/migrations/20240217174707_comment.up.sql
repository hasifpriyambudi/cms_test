CREATE TABLE `comment` (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `news_id` INT NOT NULL,
  `name` varchar(50) NOT NULL,
  `comment` text NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT NULL,
  `deleted_at` TIMESTAMP DEFAULT NULL
);