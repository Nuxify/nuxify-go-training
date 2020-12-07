/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `author_id` int NOT NULL,
  `content` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `author_id` (`author_id`),
  CONSTRAINT `posts_ibfk_1` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `posts` (`id`, `author_id`, `content`, `created_at`, `updated_at`) VALUES
(1, 1, 'Another post', '2020-11-15 11:38:03', '2020-11-15 11:38:03');
INSERT INTO `posts` (`id`, `author_id`, `content`, `created_at`, `updated_at`) VALUES
(3, 6, 'Philippians 4:13: I can do all Things...', '2020-11-15 11:57:43', '2020-11-22 13:40:52');
INSERT INTO `posts` (`id`, `author_id`, `content`, `created_at`, `updated_at`) VALUES
(4, 1, 'My First Post', '2020-11-19 10:28:29', '2020-11-19 10:28:29');
INSERT INTO `posts` (`id`, `author_id`, `content`, `created_at`, `updated_at`) VALUES
(5, 1, 'My Second Post', '2020-11-19 10:31:27', '2020-11-19 10:31:27'),
(7, 6, 'Philippians 4:13', '2020-11-19 10:31:53', '2020-11-19 11:51:08'),
(8, 7, 'My Fist Post', '2020-11-19 10:31:58', '2020-11-19 10:31:58'),
(9, 8, 'I can do all things', '2020-11-19 10:32:01', '2020-11-19 11:39:59'),
(14, 1, 'Dougiee\'s test post', '2020-11-20 21:13:28', '2020-11-20 21:13:28'),
(16, 1, 'Dougiee\'s test post number 12', '2020-11-22 13:27:10', '2020-11-22 13:27:10'),
(17, 1, 'Dougiee\'s test post number 12', '2020-11-22 14:14:37', '2020-11-22 14:14:37'),
(18, 1, 'Dougiee\'s test post number 12', '2020-11-22 14:15:25', '2020-11-22 14:15:25'),
(19, 1, 'Dougiee\'s test post number 12', '2020-11-22 14:35:05', '2020-11-22 14:35:05'),
(20, 1, 'Dougiee\'s test post number 12', '2020-11-22 14:43:17', '2020-11-22 14:43:17'),
(21, 15, 'Dougiee\'s test post number 12', '2020-11-22 14:49:33', '2020-11-22 14:49:33'),
(22, 15, 'Dougiee\'s test post number 12', '2020-11-22 19:38:54', '2020-11-22 19:38:54'),
(23, 15, 'Dougiee\'s test post number 12232', '2020-11-22 19:39:18', '2020-11-22 19:39:18'),
(24, 15, 'Dougiee\'s test post number 12232123', '2020-11-22 19:39:40', '2020-11-22 19:39:40'),
(25, 15, 'Philippians 4:13: I can do all Things...', '2020-11-22 19:40:30', '2020-11-22 20:12:14'),
(26, 15, 'Dougiee\'s test post number 1223', '2020-12-07 12:54:19', '2020-12-07 12:54:19');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;