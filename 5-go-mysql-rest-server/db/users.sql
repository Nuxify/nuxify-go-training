/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `contact_number` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UNIQUE` (`email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `users` (`id`, `email`, `first_name`, `last_name`, `contact_number`, `created_at`, `updated_at`) VALUES
(1, 'dogie@alma.com', 'Dogie', 'Alma', '+639456042882', '2020-11-15 08:48:46', '2020-11-15 08:48:46');
INSERT INTO `users` (`id`, `email`, `first_name`, `last_name`, `contact_number`, `created_at`, `updated_at`) VALUES
(6, 'updatedemail@example.com', 'Lecty', 'Eisenach', '+639456042882', '2020-11-15 09:03:04', '2020-11-15 10:43:02');
INSERT INTO `users` (`id`, `email`, `first_name`, `last_name`, `contact_number`, `created_at`, `updated_at`) VALUES
(7, '5thexample@gmail.com', 'SecondTest', 'User', '+639456042882', '2020-11-15 09:03:25', '2020-11-15 09:03:25');
INSERT INTO `users` (`id`, `email`, `first_name`, `last_name`, `contact_number`, `created_at`, `updated_at`) VALUES
(8, 'GarzAlma@example.com', 'Garz', 'Alma', '', '2020-11-17 22:31:23', '2020-11-17 22:31:23'),
(11, 'GarzAlma2@example.com', 'Garz', 'Alma', '+639456042882', '2020-11-17 22:32:00', '2020-11-17 22:32:00'),
(12, 'idolmitoy@example.com', 'Idol Mitoy Lang', 'Malakas', '+639456042882', '2020-11-18 18:23:08', '2020-11-19 11:49:09'),
(15, 'deleteyq1@example.com', 'Test', 'Updated', '+639456042882', '2020-11-20 21:01:03', '2020-11-22 13:32:15'),
(16, 'deleteyq1wq@example.com', 'Sample User', 'To be Updated', '+639456042882', '2020-11-22 20:42:40', '2020-11-22 20:42:40');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;