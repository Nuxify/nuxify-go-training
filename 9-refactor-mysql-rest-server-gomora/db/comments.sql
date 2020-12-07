/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `post_id` int NOT NULL,
  `author_id` int NOT NULL,
  `content` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `post_id` (`post_id`),
  KEY `author_id` (`author_id`),
  CONSTRAINT `comment_ibfk_1` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`),
  CONSTRAINT `comment_ibfk_2` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `comments` (`id`, `post_id`, `author_id`, `content`, `created_at`, `updated_at`) VALUES
(16, 21, 15, 'Dougiee\'s test post number 12', '2020-11-22 20:43:59', '2020-11-22 20:43:59');
INSERT INTO `comments` (`id`, `post_id`, `author_id`, `content`, `created_at`, `updated_at`) VALUES
(18, 3, 6, 'I can do all things...', '2020-11-22 20:45:01', '2020-11-22 20:45:01');
INSERT INTO `comments` (`id`, `post_id`, `author_id`, `content`, `created_at`, `updated_at`) VALUES
(19, 4, 1, 'Philippians 4:13 new comment.....', '2020-11-22 20:45:23', '2020-11-22 20:45:23');
INSERT INTO `comments` (`id`, `post_id`, `author_id`, `content`, `created_at`, `updated_at`) VALUES
(20, 5, 1, 'Philippians 4:13', '2020-11-22 20:45:58', '2020-11-22 20:45:58'),
(21, 25, 15, 'Sample comment', '2020-11-22 20:47:21', '2020-11-22 20:47:21'),
(22, 25, 15, '4:13 sample 123', '2020-11-22 20:47:23', '2020-11-22 20:47:23'),
(23, 25, 15, '4:13.', '2020-11-22 20:47:23', '2020-11-22 20:47:23'),
(24, 25, 15, 'Philippians 4:13 new comment.....', '2020-11-22 20:47:23', '2020-11-22 20:47:23'),
(25, 25, 15, 'Sample comment', '2020-11-22 20:47:23', '2020-11-22 20:47:23'),
(26, 25, 15, 'Sample comment 312', '2020-11-22 20:54:16', '2020-11-22 20:54:16'),
(27, 25, 15, 'Sample commenasdsadas', '2020-12-07 16:34:16', '2020-12-07 16:34:16'),
(28, 25, 15, 'Sample asda', '2020-12-07 16:34:33', '2020-12-07 16:34:33');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;