-- MySQL dump 10.13  Distrib 8.0.33, for Win64 (x86_64)
--
-- Host: localhost    Database: db_game
-- ------------------------------------------------------
-- Server version	8.0.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `categories` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `nama` varchar(255) DEFAULT NULL,
  `deskripsi` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES (1,'Adventure','Petualangan seru'),(2,'Horror','Seram parah cuy asli'),(3,'Puzzle','latih otak'),(4,'Survival','Bertahan hidup'),(5,'Action','Melawan musuh dengan pesona');
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comments`
--

DROP TABLE IF EXISTS `comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comments` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `review_id` bigint DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `content` text,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `likes` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_users_comments` (`user_id`),
  KEY `fk_reviews_comments` (`review_id`),
  CONSTRAINT `fk_reviews_comments` FOREIGN KEY (`review_id`) REFERENCES `reviews` (`id`),
  CONSTRAINT `fk_users_comments` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comments`
--

LOCK TABLES `comments` WRITE;
/*!40000 ALTER TABLE `comments` DISABLE KEYS */;
INSERT INTO `comments` VALUES (1,7,1,'asli bro epic parah','2023-07-27 21:53:17.880','2023-07-27 21:53:17.880',14114),(2,7,5,'Hmm menurut gw gamenya agak susah si...','2023-07-28 16:39:47.415','2023-07-28 16:39:47.415',141),(3,15,5,'Lunya aja kali bro yang ga jelas','2023-07-28 16:41:08.277','2023-07-28 16:41:08.277',19499),(4,10,5,'Reviewnya sangat jelas terima kasih','2023-07-28 16:42:29.802','2023-07-28 16:42:29.802',0),(5,20,5,'Setuju','2023-07-28 16:43:00.840','2023-07-28 16:43:00.840',1),(6,20,7,'Menurut saya bisa ini bagus bagus aja','2023-07-28 16:43:29.632','2023-07-28 16:43:29.632',0),(7,10,7,'Review paling bagus','2023-07-28 16:43:55.745','2023-07-28 16:43:55.745',1);
/*!40000 ALTER TABLE `comments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `games`
--

DROP TABLE IF EXISTS `games`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `games` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `category_id` bigint DEFAULT NULL,
  `nama` varchar(255) DEFAULT NULL,
  `deskripsi` text,
  `developer` varchar(255) DEFAULT NULL,
  `year_published` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_categories_games` (`category_id`),
  CONSTRAINT `fk_categories_games` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `games`
--

LOCK TABLES `games` WRITE;
/*!40000 ALTER TABLE `games` DISABLE KEYS */;
INSERT INTO `games` VALUES (1,1,'Elden Ring','Open World Souls Game','FROMSOFT','2022'),(2,1,'God of War','Cerita Kratos menebar abu istrinya','Santa Monica Studio','2018'),(3,5,'Grand Theft Auto V','Keseharian kriminal','Rockstar','2013'),(4,2,'Resident Evil 2','Petualangan seram melawan zombie','Capcom','2019'),(5,5,'Batman: Arkham Knight','I\'m Batman','Rocksteady Studio','2015'),(6,4,'Minecraft','Bertahan di dunia kotak','Mojang Studios','2011'),(7,1,'Far Cry New Dawn','Petualangan setelah terjadi nuklir','Ubisoft','2019'),(8,5,'Assassin\'s Creed Chronicles: China','Aksi menjadi Assassin','Ubisoft','2015'),(9,3,'Fallout 76','Bertahan setelah tragedi nuklir','Bethesda Game Studios','2018'),(10,5,'WWE 2K20','AND HIS NAME IS JOHNN CENA','2k Games','2019');
/*!40000 ALTER TABLE `games` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ratings`
--

DROP TABLE IF EXISTS `ratings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ratings` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `score` varchar(10) DEFAULT NULL,
  `deskripsi` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ratings`
--

LOCK TABLES `ratings` WRITE;
/*!40000 ALTER TABLE `ratings` DISABLE KEYS */;
INSERT INTO `ratings` VALUES (1,'10','Masterpiece'),(2,'9','Amazing'),(3,'8','Great'),(4,'7','Good'),(5,'6','Okay'),(6,'5','Mediocre'),(7,'4','Bad'),(8,'3','Awful'),(9,'2','Painful'),(10,'1','Unbearable');
/*!40000 ALTER TABLE `ratings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reviews`
--

DROP TABLE IF EXISTS `reviews`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `reviews` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `game_id` bigint DEFAULT NULL,
  `rating_id` bigint DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `content` text,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_games_reviews` (`game_id`),
  KEY `fk_ratings_reviews` (`rating_id`),
  KEY `fk_users_reviews` (`user_id`),
  CONSTRAINT `fk_games_reviews` FOREIGN KEY (`game_id`) REFERENCES `games` (`id`),
  CONSTRAINT `fk_ratings_reviews` FOREIGN KEY (`rating_id`) REFERENCES `ratings` (`id`),
  CONSTRAINT `fk_users_reviews` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reviews`
--

LOCK TABLES `reviews` WRITE;
/*!40000 ALTER TABLE `reviews` DISABLE KEYS */;
INSERT INTO `reviews` VALUES (7,1,1,'Elden Ring: Sebuah Boss Battle Masterpiece','Boss Battle dan musik dari game ini sangat fenomenal','2023-07-27 21:46:47.981','2023-07-27 21:46:47.981',4),(8,1,10,'Jangan Maenin Elden Ring','Game susah parah ga worth it buat dimaenin','2023-07-27 21:59:23.515','2023-07-27 21:59:23.515',6),(9,2,2,'God of War: Story, not Action','Story God of War sangat menakjubkan, pantas menjadi game yang bagus','2023-07-28 16:09:00.267','2023-07-28 16:09:00.267',5),(10,3,3,'Boss Maen GTA V','GTA V yang dimainkan anak jaman sekaran seru juga ya','2023-07-28 16:10:35.164','2023-07-28 16:10:35.164',7),(11,3,1,'GTA V, Chaos Asli','GTA V masterpiece asli, bisa ngancurin mobil - mobil orang :v','2023-07-28 16:11:32.838','2023-07-28 16:11:32.838',6),(12,4,8,'Resident Evil Serem','Serem parah gw masih stuck di early game ga mau gerak','2023-07-28 16:12:22.374','2023-07-28 16:12:22.374',6),(13,4,3,'Resident Evil Horror Asli','Ga terlalu suka game horror, tapi ini bagus','2023-07-28 16:13:15.032','2023-07-28 16:13:15.032',4),(14,5,3,'Batman: Arkham Knight Could be Better','Konsepnya bagus, tapi combatnya masi kurang kerasa','2023-07-28 16:14:31.782','2023-07-28 16:14:31.782',5),(15,6,6,'Minecraft Ga Jelas','Grafiknya jelek kotak kotak jadi biasa aja ah','2023-07-28 16:16:07.531','2023-07-28 16:16:07.531',6),(16,6,1,'Minecraft is a Timeless Masterpiece','Minecraft sangat bagus untuk melatih kreativitas','2023-07-28 16:16:51.259','2023-07-28 16:16:51.259',4),(17,7,7,'Far Cry New Dawn: If the word bad was a game','Far Cry ini cukup buruk dibanding Far Cry sebelumnya','2023-07-28 16:18:46.958','2023-07-28 16:18:46.958',4),(18,8,8,'The Worst Assassin\'s Creed','Assassin\'s Creed ini terlalu membosankan','2023-07-28 16:19:38.666','2023-07-28 16:19:38.666',5),(19,10,9,'WWE 2k20 ga seru...','Walaupun maen sama temen ini lebih ke nyiksa aja si combatnya','2023-07-28 16:21:08.123','2023-07-28 16:21:08.123',6),(20,9,10,'Fallout 4, tapi jauh lebih buruk','Ini game sudah bukan buruk lagi, tapi paling buruk','2023-07-28 16:23:46.279','2023-07-28 16:23:46.279',5);
/*!40000 ALTER TABLE `reviews` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(191) NOT NULL,
  `email` varchar(191) NOT NULL,
  `password` longtext NOT NULL,
  `role` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'admin','daffaraygama55','$2a$10$YPgpQp7Iv1DimxhteueE2uUEm/eepR0qhv4OtTCr5YOzhLhbn8Dii','admin'),(3,'user','user@gmail.com','$2a$10$xx15l2Yz3p.qb4Hyh377wej1fl7FSXhLJXbHr/rxXxehal52x7GX6','user'),(4,'raygama','raygama@gmail.com','$2a$10$xf.D.WSVvsRB/g5BKjkqR.NVN.0nF0XqVWoCQp7LJKZq7pwch8AuK','user'),(5,'BIY','abiyyu@gmail.com','$2a$10$WdjreQcK3SYkKQAJCCXIEuxYiWzCrpT7ZOxw2evmQQ8phgLDdOyIW','user'),(6,'imissher','XXX_gwganteng_XXX@gmail.com','$2a$10$LmjUQpxTGhym.jDCd0BLXe6iRcV4LowtPui.XbkR3cKQ0p3NwhFSW','user'),(7,'boss','boss@gmail.com','$2a$10$ltHoUiFQmxbBUBaHH6JFROerBf9QzwKR3PEtfvLR7v5WELevMDBL6','admin');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-07-28 17:00:11
