CREATE TABLE `custom_page` (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `custom_url` varchar(50) NOT NULL,
  `page_content` text NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT NULL,
  `deleted_at` TIMESTAMP DEFAULT NULL
);