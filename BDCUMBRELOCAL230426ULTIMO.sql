CREATE DATABASE  IF NOT EXISTS `db_cumbre_ppln` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `db_cumbre_ppln`;
-- MySQL dump 10.13  Distrib 8.0.41, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: db_cumbre_ppln
-- ------------------------------------------------------
-- Server version	8.4.9

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `document_types`
--

DROP TABLE IF EXISTS `document_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `document_types` (
  `id` varchar(36) NOT NULL,
  `number` varchar(5) NOT NULL,
  `description` varchar(200) NOT NULL,
  `abbreviated_description` varchar(50) NOT NULL,
  `enable` tinyint(1) NOT NULL,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `document_types`
--

LOCK TABLES `document_types` WRITE;
/*!40000 ALTER TABLE `document_types` DISABLE KEYS */;
INSERT INTO `document_types` VALUES ('00a58296-93b4-11ee-a040-0242ac11000e','01','DOCUMENTO NACIONAL DE IDENTIDAD','DNI',1,'2023-12-05 15:49:56',NULL),('00a584ae-93b4-11ee-a040-0242ac11000e','04','CARNÉ DE EXTRANJERÍA','CARNÉ EXT.',1,'2023-12-05 15:49:56',NULL),('00a58522-93b4-11ee-a040-0242ac11000e','06','REG. ÚNICO DE CONTRIBUYENTES (1)','RUC',1,'2023-12-05 15:49:56',NULL),('00a58572-93b4-11ee-a040-0242ac11000e','07','PASAPORTE','PASAPORTE',1,'2023-12-05 15:49:56',NULL),('00a585c3-93b4-11ee-a040-0242ac11000e','09','CARNÉ DE SOLICIT DE REFUGIO','CARNÉ SOLIC REFUGIO',1,'2023-12-05 15:49:56',NULL),('00a58610-93b4-11ee-a040-0242ac11000e','11','PARTIDA DE NACIMIENTO (2)','PART. NAC.',1,'2023-12-05 15:49:56',NULL),('00a58659-93b4-11ee-a040-0242ac11000e','22','CARNÉ DE IDENTIDAD - RELACIONES EXTERIORES','C.IDENT.-RREE',1,'2023-12-05 15:49:56',NULL),('00a586a3-93b4-11ee-a040-0242ac11000e','23','PERM.TEMP.PERMANENCIA','PTP',1,'2023-12-05 15:49:56',NULL),('00a586f0-93b4-11ee-a040-0242ac11000e','24','DOC. DE IDENTIDAD EXTRANJERO (3)','DOC.ID.EXTR.',1,'2023-12-05 15:49:56',NULL),('00a58739-93b4-11ee-a040-0242ac11000e','26','CARNÉ DE PERMISO TEMP DE PERMANENCIA','CPP',1,'2023-12-05 15:49:56',NULL);
/*!40000 ALTER TABLE `document_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `event_types`
--

DROP TABLE IF EXISTS `event_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `event_types` (
  `id` varchar(36) NOT NULL,
  `code` varchar(15) NOT NULL,
  `description` text NOT NULL,
  `enable` tinyint(1) NOT NULL,
  `created_at` timestamp NOT NULL,
  `created_by` varchar(36) NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `deleted_by` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `event_types`
--

LOCK TABLES `event_types` WRITE;
/*!40000 ALTER TABLE `event_types` DISABLE KEYS */;
INSERT INTO `event_types` VALUES ('2fe413cd-3d91-11f1-bd7e-0242ac110002','0001','EVENTO',1,'2026-04-21 09:48:57','30e42728-fb67-11ee-a6a0-0242ac110013',NULL,NULL);
/*!40000 ALTER TABLE `event_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `events`
--

DROP TABLE IF EXISTS `events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `events` (
  `id` varchar(36) NOT NULL,
  `type_id` varchar(36) NOT NULL,
  `code` varchar(15) DEFAULT NULL,
  `name` varchar(36) NOT NULL,
  `description` varchar(256) NOT NULL,
  `enable` tinyint(1) NOT NULL,
  `total_reg` int NOT NULL,
  `total_pay` int NOT NULL,
  `total_pres` int NOT NULL,
  `created_at` timestamp NOT NULL,
  `created_by` varchar(36) NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `deleted_by` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `events_event_types__fk` (`type_id`),
  KEY `events_users_id_fk` (`created_by`),
  KEY `events_users_id_fk_2` (`deleted_by`),
  CONSTRAINT `events_event_types__fk` FOREIGN KEY (`type_id`) REFERENCES `event_types` (`id`),
  CONSTRAINT `events_users_id_fk` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`),
  CONSTRAINT `events_users_id_fk_2` FOREIGN KEY (`deleted_by`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events`
--

LOCK TABLES `events` WRITE;
/*!40000 ALTER TABLE `events` DISABLE KEYS */;
INSERT INTO `events` VALUES ('45e2176c-3d91-11f1-bd7e-0242ac110002','2fe413cd-3d91-11f1-bd7e-0242ac110002','0001','EVENTO','EVENTO',1,0,0,0,'2026-04-21 09:49:41','30e42728-fb67-11ee-a6a0-0242ac110013',NULL,NULL);
/*!40000 ALTER TABLE `events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `modules`
--

DROP TABLE IF EXISTS `modules`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `modules` (
  `id` varchar(36) NOT NULL,
  `name` varchar(256) NOT NULL,
  `description` varchar(256) DEFAULT NULL,
  `code` varchar(200) NOT NULL,
  `icon` varchar(100) NOT NULL,
  `position` int NOT NULL,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `modules`
--

LOCK TABLES `modules` WRITE;
/*!40000 ALTER TABLE `modules` DISABLE KEYS */;
INSERT INTO `modules` VALUES ('5ebd7b4e-3e8a-11f1-8bb5-0242ac110002','Configuraciones','Configuraciones de evento','event.settings','ri-settings-3-line o ri-key-line text-danger',6,'2026-04-22 15:33:59',NULL),('92106652-3e8b-11f1-8bb5-0242ac110002','Pago Inscripcion','Pago de inscripcion','event.payment','ri-money-line o ri-exchange-dollar-line',4,'2026-04-22 15:42:10',NULL),('9f367713-3e8a-11f1-8bb5-0242ac110002','Sesiones','Sessiones','event.sessions','ri-file-list-2-line o ri-checkbox-multiple-line',1,'2026-04-22 15:35:31',NULL),('c1a8df87-3e8b-11f1-8bb5-0242ac110002','Reportes','Reportes','event.report','ri-file-list-2-line o ri-checkbox-multiple-line',5,'2026-04-22 15:42:52',NULL),('d293e523-3e8a-11f1-8bb5-0242ac110002','Inscripciones','Inscripciones','event.registrations','ri-table-line o ri-bar-chart-2-line',2,'2026-04-22 15:36:27',NULL),('f34d5eec-3e8a-11f1-8bb5-0242ac110002','Entrega Material','Entrega y recojo de Materiales','event.receive','ri-truck-line o ri-shipping-box-line',3,'2026-04-22 15:40:53',NULL);
/*!40000 ALTER TABLE `modules` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payments`
--

DROP TABLE IF EXISTS `payments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `payments` (
  `id` varchar(36) NOT NULL,
  `amount` decimal(10,3) NOT NULL,
  `created_at` timestamp NOT NULL,
  `created_by` varchar(36) NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `deleted_by` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `payments_users_id_fk` (`created_by`),
  KEY `payments_users_id_fk_2` (`deleted_by`),
  CONSTRAINT `payments_users_id_fk` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`),
  CONSTRAINT `payments_users_id_fk_2` FOREIGN KEY (`deleted_by`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payments`
--

LOCK TABLES `payments` WRITE;
/*!40000 ALTER TABLE `payments` DISABLE KEYS */;
/*!40000 ALTER TABLE `payments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `people`
--

DROP TABLE IF EXISTS `people`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `people` (
  `id` varchar(36) NOT NULL,
  `user_id` varchar(36) DEFAULT NULL,
  `type_document_id` varchar(250) NOT NULL,
  `document` varchar(250) NOT NULL,
  `names` varchar(250) NOT NULL,
  `surname` varchar(250) NOT NULL,
  `last_name` varchar(250) DEFAULT NULL,
  `phone` varchar(250) DEFAULT NULL,
  `email` varchar(250) DEFAULT NULL,
  `gender` varchar(255) DEFAULT NULL,
  `enable` tinyint(1) NOT NULL,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `people_document_types_id_fk` (`type_document_id`),
  KEY `people_users_id_fk` (`user_id`),
  CONSTRAINT `people_document_types_id_fk` FOREIGN KEY (`type_document_id`) REFERENCES `document_types` (`id`),
  CONSTRAINT `people_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `people`
--

LOCK TABLES `people` WRITE;
/*!40000 ALTER TABLE `people` DISABLE KEYS */;
INSERT INTO `people` VALUES ('01307ba2-726c-4316-b3df-b2a14745d684','30e42728-fb67-11ee-a6a0-0242ac110013','00a58296-93b4-11ee-a040-0242ac11000e','00000000','SMART','ONE','','','','VARON',1,'2024-02-21 14:45:24',NULL),('038d1d8a-3ea1-11f1-8bb5-0242ac110002','2191771e-3e8f-11f1-8bb5-0242ac110002','00a58296-93b4-11ee-a040-0242ac11000e','11111111','PERSON 1','ONE',NULL,NULL,NULL,NULL,1,'2026-04-22 18:15:11',NULL);
/*!40000 ALTER TABLE `people` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `registration_payments`
--

DROP TABLE IF EXISTS `registration_payments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `registration_payments` (
  `id` varchar(36) NOT NULL,
  `registration_id` varchar(36) NOT NULL,
  `payment_id` varchar(36) NOT NULL,
  `created_at` timestamp NOT NULL,
  `created_by` varchar(36) NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `deleted_by` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `registration_payments_payments_id_fk` (`payment_id`),
  KEY `registration_payments_registrations_id_fk` (`registration_id`),
  KEY `registration_payments_users_id_fk` (`created_by`),
  KEY `registration_payments_users_id_fk_2` (`deleted_by`),
  CONSTRAINT `registration_payments_payments_id_fk` FOREIGN KEY (`payment_id`) REFERENCES `payments` (`id`),
  CONSTRAINT `registration_payments_registrations_id_fk` FOREIGN KEY (`registration_id`) REFERENCES `registrations` (`id`),
  CONSTRAINT `registration_payments_users_id_fk` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`),
  CONSTRAINT `registration_payments_users_id_fk_2` FOREIGN KEY (`deleted_by`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `registration_payments`
--

LOCK TABLES `registration_payments` WRITE;
/*!40000 ALTER TABLE `registration_payments` DISABLE KEYS */;
/*!40000 ALTER TABLE `registration_payments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `registrations`
--

DROP TABLE IF EXISTS `registrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `registrations` (
  `id` varchar(36) NOT NULL,
  `session_id` varchar(36) NOT NULL,
  `user_id` varchar(36) NOT NULL,
  `created_at` timestamp NOT NULL,
  `created_by` varchar(36) NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `deleted_by` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `registrations_sessions_id_fk` (`session_id`),
  KEY `registrations_users_id_fk` (`user_id`),
  KEY `registrations_users_id_fk_2` (`created_by`),
  KEY `registrations_users_id_fk_3` (`deleted_by`),
  CONSTRAINT `registrations_sessions_id_fk` FOREIGN KEY (`session_id`) REFERENCES `sessions` (`id`),
  CONSTRAINT `registrations_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `registrations_users_id_fk_2` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`),
  CONSTRAINT `registrations_users_id_fk_3` FOREIGN KEY (`deleted_by`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `registrations`
--

LOCK TABLES `registrations` WRITE;
/*!40000 ALTER TABLE `registrations` DISABLE KEYS */;
INSERT INTO `registrations` VALUES ('b6cc2af6-3d91-11f1-bd7e-0242ac110002','88939fe5-3d91-11f1-bd7e-0242ac110002','30e42728-fb67-11ee-a6a0-0242ac110013','2026-04-21 09:52:50','30e42728-fb67-11ee-a6a0-0242ac110013',NULL,NULL);
/*!40000 ALTER TABLE `registrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_views`
--

DROP TABLE IF EXISTS `role_views`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role_views` (
  `id` varchar(36) NOT NULL,
  `view_id` varchar(36) NOT NULL,
  `role_id` varchar(36) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_role_views_role_id_view_id_unique` (`role_id`,`view_id`,`deleted_at`),
  KEY `idx_role_views_roles_null_fk` (`role_id`),
  KEY `idx_role_views_views_null_fk` (`view_id`),
  CONSTRAINT `role_views_roles_null_fk` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`),
  CONSTRAINT `role_views_views_null_fk` FOREIGN KEY (`view_id`) REFERENCES `views` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_views`
--

LOCK TABLES `role_views` WRITE;
/*!40000 ALTER TABLE `role_views` DISABLE KEYS */;
INSERT INTO `role_views` VALUES ('10902291-3e95-11f1-8bb5-0242ac110002','ccf60206-3e8d-11f1-8bb5-0242ac110002','9fdf1211-3e90-11f1-8bb5-0242ac110002','2026-04-22 16:49:45',NULL),('2f51363e-3e95-11f1-8bb5-0242ac110002','fd5d83a7-3e8d-11f1-8bb5-0242ac110002','9fdf1211-3e90-11f1-8bb5-0242ac110002','2026-04-22 16:50:06',NULL),('3c65beb6-3e95-11f1-8bb5-0242ac110002','2e5ed8a1-3e8e-11f1-8bb5-0242ac110002','9fdf1211-3e90-11f1-8bb5-0242ac110002','2026-04-22 16:50:30',NULL),('68775e45-3e93-11f1-8bb5-0242ac110002','693fd0de-3e8c-11f1-8bb5-0242ac110002','9fdf1211-3e90-11f1-8bb5-0242ac110002','2026-04-22 16:44:32',NULL),('74d38759-3e94-11f1-8bb5-0242ac110002','df936f2c-3e8c-11f1-8bb5-0242ac110002','9fdf1211-3e90-11f1-8bb5-0242ac110002','2026-04-22 16:44:59',NULL),('895948a2-3e95-11f1-8bb5-0242ac110002','df936f2c-3e8c-11f1-8bb5-0242ac110002','a378f375-3e8f-11f1-8bb5-0242ac110002','2026-04-22 16:53:03',NULL),('8cd4934e-3e94-11f1-8bb5-0242ac110002','121e1288-3e8d-11f1-8bb5-0242ac110002','9fdf1211-3e90-11f1-8bb5-0242ac110002','2026-04-22 16:46:26',NULL),('b9ce1b24-3e94-11f1-8bb5-0242ac110002','48d5d812-3e8d-11f1-8bb5-0242ac110002','9fdf1211-3e90-11f1-8bb5-0242ac110002','2026-04-22 16:46:56',NULL),('caf69e6e-3e94-11f1-8bb5-0242ac110002','520276ee-3e8d-11f1-8bb5-0242ac110002','9fdf1211-3e90-11f1-8bb5-0242ac110002','2026-04-22 16:47:40',NULL),('f424d91e-3e94-11f1-8bb5-0242ac110002','8e2b30fa-3e8d-11f1-8bb5-0242ac110002','9fdf1211-3e90-11f1-8bb5-0242ac110002','2026-04-22 16:48:52',NULL);
/*!40000 ALTER TABLE `role_views` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `roles` (
  `id` varchar(36) NOT NULL,
  `name` varchar(256) DEFAULT NULL,
  `description` varchar(256) DEFAULT NULL,
  `enable` tinyint(1) DEFAULT NULL,
  `role_default_id` varchar(36) DEFAULT NULL,
  `event_id` varchar(36) DEFAULT NULL,
  `workshop_id` varchar(36) DEFAULT NULL,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `roles_merchants_id_fk` (`event_id`),
  KEY `roles_roles_defaults_id_fk` (`role_default_id`),
  KEY `roles_stores_id_fk` (`workshop_id`),
  CONSTRAINT `roles_merchants_id_fk` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`),
  CONSTRAINT `roles_roles_defaults_id_fk` FOREIGN KEY (`role_default_id`) REFERENCES `roles_defaults` (`id`),
  CONSTRAINT `roles_stores_id_fk` FOREIGN KEY (`workshop_id`) REFERENCES `workshops` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES ('556ac338-3e91-11f1-8bb5-0242ac110002','Registrador Pago','Registrador Pago',1,'2322f5a5-3e90-11f1-8bb5-0242ac110002','45e2176c-3d91-11f1-bd7e-0242ac110002',NULL,'2026-04-22 16:22:45',NULL),('9fdf1211-3e90-11f1-8bb5-0242ac110002','Administrador del Sistema','Administrador del Sistema',1,'7b4da544-3e8f-11f1-8bb5-0242ac110002',NULL,NULL,'2026-04-22 16:18:24',NULL),('a378f375-3e8f-11f1-8bb5-0242ac110002','Registrador Inscripcion','Registrador de Inscripcion',1,'a378f375-3e8f-11f1-8bb5-0242ac110002','45e2176c-3d91-11f1-bd7e-0242ac110002',NULL,'2026-04-22 16:21:53',NULL),('e82702de-3e90-11f1-8bb5-0242ac110002','Registrador Asistencia','Registrador de Asistencia',1,'d071caaf-3e8f-11f1-8bb5-0242ac110002','45e2176c-3d91-11f1-bd7e-0242ac110002',NULL,'2026-04-22 16:21:19',NULL);
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles_defaults`
--

DROP TABLE IF EXISTS `roles_defaults`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `roles_defaults` (
  `id` varchar(36) NOT NULL,
  `name` varchar(256) NOT NULL,
  `description` varchar(256) NOT NULL,
  `to_event` tinyint(1) DEFAULT NULL,
  `to_workshop` tinyint(1) DEFAULT NULL,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles_defaults`
--

LOCK TABLES `roles_defaults` WRITE;
/*!40000 ALTER TABLE `roles_defaults` DISABLE KEYS */;
INSERT INTO `roles_defaults` VALUES ('2322f5a5-3e90-11f1-8bb5-0242ac110002','Registrador Pago','Registrador de Pago',1,0,'2026-04-22 16:14:02',NULL),('7b4da544-3e8f-11f1-8bb5-0242ac110002','Administrador General','Administrador General Evento',0,0,'2026-04-22 16:09:53',NULL),('a378f375-3e8f-11f1-8bb5-0242ac110002','Registrador Inscripcion','Registrador de Inscripcion',1,0,'2026-04-22 16:11:17',NULL),('d071caaf-3e8f-11f1-8bb5-0242ac110002','Registrador Asistencia','Registrador de Asistencia',1,0,'2026-04-22 16:13:23',NULL);
/*!40000 ALTER TABLE `roles_defaults` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sessions`
--

DROP TABLE IF EXISTS `sessions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sessions` (
  `id` varchar(36) NOT NULL,
  `workshop_id` varchar(36) NOT NULL,
  `start_date` timestamp NOT NULL,
  `end_date` timestamp NOT NULL,
  `total_reg` int NOT NULL,
  `total_pay` int NOT NULL,
  `total_pres` int NOT NULL,
  `created_at` timestamp NOT NULL,
  `created_by` varchar(36) NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `deleted_by` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sessions_users_id_fk` (`created_by`),
  KEY `sessions_users_id_fk_2` (`deleted_by`),
  KEY `sessions_workshops_id_fk` (`workshop_id`),
  CONSTRAINT `sessions_users_id_fk` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`),
  CONSTRAINT `sessions_users_id_fk_2` FOREIGN KEY (`deleted_by`) REFERENCES `users` (`id`),
  CONSTRAINT `sessions_workshops_id_fk` FOREIGN KEY (`workshop_id`) REFERENCES `workshops` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sessions`
--

LOCK TABLES `sessions` WRITE;
/*!40000 ALTER TABLE `sessions` DISABLE KEYS */;
INSERT INTO `sessions` VALUES ('02687e8c-d267-4842-b974-3fb2a1873467','0464efbd-3d91-11f1-bd7e-0242ac110002','2024-09-06 08:10:00','2024-09-06 08:10:00',0,0,0,'2026-04-23 15:41:42','30e42728-fb67-11ee-a6a0-0242ac110013','2026-04-23 15:42:13','30e42728-fb67-11ee-a6a0-0242ac110013'),('88939fe5-3d91-11f1-bd7e-0242ac110002','0464efbd-3d91-11f1-bd7e-0242ac110002','2026-04-21 09:51:23','2026-04-21 11:00:00',0,0,0,'2026-04-21 09:51:57','30e42728-fb67-11ee-a6a0-0242ac110013',NULL,NULL),('8b4806f7-91d3-4c63-9b2f-c9c583fbe34c','0464efbd-3d91-11f1-bd7e-0242ac110002','2025-09-06 08:10:00','2025-09-06 08:10:00',0,0,0,'2026-04-23 15:40:41','30e42728-fb67-11ee-a6a0-0242ac110013',NULL,NULL);
/*!40000 ALTER TABLE `sessions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_roles`
--

DROP TABLE IF EXISTS `user_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_roles` (
  `id` varchar(36) NOT NULL,
  `user_id` varchar(36) NOT NULL,
  `role_id` varchar(36) NOT NULL,
  `enable` tinyint(1) NOT NULL,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_roles_roles_null_fk` (`role_id`),
  KEY `user_roles_users_null_fk` (`user_id`),
  CONSTRAINT `user_roles_roles_null_fk` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`),
  CONSTRAINT `user_roles_users_null_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_roles`
--

LOCK TABLES `user_roles` WRITE;
/*!40000 ALTER TABLE `user_roles` DISABLE KEYS */;
INSERT INTO `user_roles` VALUES ('0635049d-3e92-11f1-8bb5-0242ac110002','30e42728-fb67-11ee-a6a0-0242ac110013','9fdf1211-3e90-11f1-8bb5-0242ac110002',1,'2026-04-22 16:27:25',NULL),('b2dd58fb-3e92-11f1-8bb5-0242ac110002','2191771e-3e8f-11f1-8bb5-0242ac110002','e82702de-3e90-11f1-8bb5-0242ac110002',1,'2026-04-22 16:32:11',NULL);
/*!40000 ALTER TABLE `user_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_types`
--

DROP TABLE IF EXISTS `user_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_types` (
  `id` varchar(36) NOT NULL,
  `description` varchar(200) NOT NULL,
  `code` varchar(100) NOT NULL,
  `enable` tinyint(1) NOT NULL,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_types`
--

LOCK TABLES `user_types` WRITE;
/*!40000 ALTER TABLE `user_types` DISABLE KEYS */;
INSERT INTO `user_types` VALUES ('2b25d1a7-fb68-11ee-a6a0-0242ac110013','Usuario externo','USER_EXTERNAL',1,'2024-01-01 09:00:00',NULL),('6d449be9-fb67-11ee-a6a0-0242ac110013','Usuario interno','USER_INTERNAL',1,'2024-01-01 09:00:00',NULL);
/*!40000 ALTER TABLE `user_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` varchar(36) NOT NULL,
  `username` varchar(256) NOT NULL,
  `password_hash` varchar(256) NOT NULL,
  `type_id` varchar(36) NOT NULL,
  `theme` varchar(25) DEFAULT NULL,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `users_user_types_id_fk` (`type_id`),
  CONSTRAINT `users_user_types_id_fk` FOREIGN KEY (`type_id`) REFERENCES `user_types` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('2191771e-3e8f-11f1-8bb5-0242ac110002','user.cumbre','smart2024','6d449be9-fb67-11ee-a6a0-0242ac110013','LIGHT','2026-04-22 16:07:53',NULL),('30e42728-fb67-11ee-a6a0-0242ac110013','admin.smart','smart2024','6d449be9-fb67-11ee-a6a0-0242ac110013','LIGHT','2024-01-01 09:00:00',NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `views`
--

DROP TABLE IF EXISTS `views`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `views` (
  `id` varchar(36) NOT NULL,
  `name` varchar(250) NOT NULL,
  `description` text NOT NULL,
  `position` int DEFAULT NULL,
  `url` varchar(250) NOT NULL,
  `icon` varchar(250) NOT NULL,
  `module_id` varchar(36) NOT NULL,
  `created_at` timestamp NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `views_modules_id_fk` (`module_id`),
  CONSTRAINT `views_modules_id_fk` FOREIGN KEY (`module_id`) REFERENCES `modules` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `views`
--

LOCK TABLES `views` WRITE;
/*!40000 ALTER TABLE `views` DISABLE KEYS */;
INSERT INTO `views` VALUES ('121e1288-3e8d-11f1-8bb5-0242ac110002','Entrega Material','Entrega Material',3,'/event/received','ri-truck-line','f34d5eec-3e8a-11f1-8bb5-0242ac110002','2026-04-22 15:52:16',NULL),('2e5ed8a1-3e8e-11f1-8bb5-0242ac110002','Talleres','Mantenimiento de Talleres',4,'/event/configs/workshops','ri-user-2-fill','5ebd7b4e-3e8a-11f1-8bb5-0242ac110002','2026-04-22 16:00:25',NULL),('48d5d812-3e8d-11f1-8bb5-0242ac110002','Pago Inscripcion','Pago de inscripcion',4,'/event/payment','ri-money-line','92106652-3e8b-11f1-8bb5-0242ac110002','2026-04-22 15:53:03',NULL),('520276ee-3e8d-11f1-8bb5-0242ac110002','Reportes','Reportes',5,'/event/reports','ri-file-list-2-line','c1a8df87-3e8b-11f1-8bb5-0242ac110002','2026-04-22 15:54:26',NULL),('693fd0de-3e8c-11f1-8bb5-0242ac110002','Sesiones','Sesiones',1,'/event/sessions','ri-file-list-2-line','9f367713-3e8a-11f1-8bb5-0242ac110002','2026-04-22 15:48:43',NULL),('8e2b30fa-3e8d-11f1-8bb5-0242ac110002','Usuarios','Mantenimiento de Usuarios',1,'/event/configs/users','ri-user-2-fill','5ebd7b4e-3e8a-11f1-8bb5-0242ac110002','2026-04-22 15:56:48',NULL),('ccf60206-3e8d-11f1-8bb5-0242ac110002','Registrados','Personas registradas',2,'/event/configs/people','ri-user-2-fill','5ebd7b4e-3e8a-11f1-8bb5-0242ac110002','2026-04-22 15:58:06',NULL),('df936f2c-3e8c-11f1-8bb5-0242ac110002','Inscripciones','Inscripciones',2,'/event/asistentes','ri-table-line','d293e523-3e8a-11f1-8bb5-0242ac110002','2026-04-22 15:51:30',NULL),('fd5d83a7-3e8d-11f1-8bb5-0242ac110002','Eventos','Mantenimiento de Eventos',3,'/event/configs/events','ri-user-2-fill','5ebd7b4e-3e8a-11f1-8bb5-0242ac110002','2026-04-22 15:59:06',NULL);
/*!40000 ALTER TABLE `views` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `workshop_types`
--

DROP TABLE IF EXISTS `workshop_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `workshop_types` (
  `id` varchar(36) NOT NULL,
  `code` varchar(15) NOT NULL,
  `description` text NOT NULL,
  `enable` tinyint(1) NOT NULL,
  `created_at` timestamp NOT NULL,
  `created_by` varchar(36) NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `deleted_by` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `workshop_types_users_id_fk` (`created_by`),
  KEY `workshop_types_users_id_fk_2` (`deleted_by`),
  CONSTRAINT `workshop_types_users_id_fk` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`),
  CONSTRAINT `workshop_types_users_id_fk_2` FOREIGN KEY (`deleted_by`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `workshop_types`
--

LOCK TABLES `workshop_types` WRITE;
/*!40000 ALTER TABLE `workshop_types` DISABLE KEYS */;
INSERT INTO `workshop_types` VALUES ('a70f46f3-3d90-11f1-bd7e-0242ac110002','0001','TIPO DE TALLER',1,'2026-04-21 09:46:24','30e42728-fb67-11ee-a6a0-0242ac110013',NULL,NULL);
/*!40000 ALTER TABLE `workshop_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `workshops`
--

DROP TABLE IF EXISTS `workshops`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `workshops` (
  `id` varchar(36) NOT NULL,
  `type_id` varchar(36) NOT NULL,
  `name` text NOT NULL,
  `shortname` text,
  `code` varchar(15) DEFAULT NULL,
  `capacity` int NOT NULL,
  `total_reg` int NOT NULL,
  `total_pay` int NOT NULL,
  `total_pres` int NOT NULL,
  `event_id` varchar(36) NOT NULL,
  `created_at` timestamp NOT NULL,
  `created_by` varchar(36) NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `deleted_by` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `workshops_events_id_fk` (`event_id`),
  KEY `workshops_users_id_fk` (`created_by`),
  KEY `workshops_users_id_fk_2` (`deleted_by`),
  KEY `workshops_workshop_types_id_fk` (`type_id`),
  CONSTRAINT `workshops_events_id_fk` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`),
  CONSTRAINT `workshops_users_id_fk` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`),
  CONSTRAINT `workshops_users_id_fk_2` FOREIGN KEY (`deleted_by`) REFERENCES `users` (`id`),
  CONSTRAINT `workshops_workshop_types_id_fk` FOREIGN KEY (`type_id`) REFERENCES `workshop_types` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `workshops`
--

LOCK TABLES `workshops` WRITE;
/*!40000 ALTER TABLE `workshops` DISABLE KEYS */;
INSERT INTO `workshops` VALUES ('0464efbd-3d91-11f1-bd7e-0242ac110002','a70f46f3-3d90-11f1-bd7e-0242ac110002','TALLERES - 1ER BLOQUE','1ER BLOQUE','0001',0,0,0,0,'45e2176c-3d91-11f1-bd7e-0242ac110002','2026-04-21 09:50:04','30e42728-fb67-11ee-a6a0-0242ac110013',NULL,NULL),('b3c5cde4-5a58-4841-a998-d2bd5d40cfa2','a70f46f3-3d90-11f1-bd7e-0242ac110002','TALLERES - PRUEBA #2','2ER BLOQUE','0002',40,0,0,0,'45e2176c-3d91-11f1-bd7e-0242ac110002','2026-04-23 15:48:25','30e42728-fb67-11ee-a6a0-0242ac110013','2026-04-23 15:51:56','30e42728-fb67-11ee-a6a0-0242ac110013');
/*!40000 ALTER TABLE `workshops` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-04-23 19:00:26
