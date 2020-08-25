create table categories(
 `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name`    VARCHAR(160),
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL
)ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;