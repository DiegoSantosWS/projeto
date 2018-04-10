-- MySQL dump 10.13  Distrib 5.7.21, for Linux (x86_64)
--
-- Host: localhost    Database: cms
-- ------------------------------------------------------
-- Server version	5.7.21-0ubuntu0.16.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `categorys`
--

DROP TABLE IF EXISTS `categorys`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `categorys` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `grupo` int(11) DEFAULT NULL,
  `categoria` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categorys`
--

LOCK TABLES `categorys` WRITE;
/*!40000 ALTER TABLE `categorys` DISABLE KEYS */;
INSERT INTO `categorys` VALUES (1,1,'Quem sou'),(2,1,'Curriculo'),(3,1,'Experiência'),(4,1,'Paginas'),(6,2,'Consultoria');
/*!40000 ALTER TABLE `categorys` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `content`
--

DROP TABLE IF EXISTS `content`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `content` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(50) DEFAULT NULL,
  `description` varchar(50) DEFAULT NULL,
  `text` text,
  `slug` varchar(45) DEFAULT NULL,
  `date_ini` date DEFAULT NULL,
  `date_end` date DEFAULT NULL,
  `group` int(11) DEFAULT NULL,
  `category_content` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `content`
--

LOCK TABLES `content` WRITE;
/*!40000 ALTER TABLE `content` DISABLE KEYS */;
INSERT INTO `content` VALUES (1,'Curriculo diego','descricao','ashsdaksdhasdahdas\r\nalshdajsdlkdaljsdklasdjksd\r\nsadçjsdkasldasd','teste-teste','2018-01-01','2018-07-01',1,'2'),(2,'teste 2 teste','teste 2 teste','teste 2 teste','teste-2-teste','2017-01-28','2017-03-31',1,'2'),(3,'teste teste','teste teste teste teste','teste teste experiencia','teste-teste','2018-01-26','2018-06-30',1,'3'),(4,'teste diego asdas','testando alteração','teste teste teste teste teste alterado','teste-teste','2018-01-01','2018-07-01',2,'6'),(8,'Meus Dados de conteudo','teste teste teste','teste teste teste','teste-teste','2000-01-01','2019-01-01',1,'1'),(9,'DIEGO DOS SANTOS TESTE','TESTE TESTE','LÇDKASKDÇLSDÇLSDAKÇLASDKÇLSD TEXMDNSDKJSAKDAJSKL','teste-teste','2018-02-23','2018-02-28',1,'2');
/*!40000 ALTER TABLE `content` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ged`
--

DROP TABLE IF EXISTS `ged`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ged` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `idref` int(11) DEFAULT NULL,
  `nome` varchar(255) DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  `comentario` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ged`
--

LOCK TABLES `ged` WRITE;
/*!40000 ALTER TABLE `ged` DISABLE KEYS */;
INSERT INTO `ged` VALUES (3,1,'29012018192749ima1.jpg','uploadfile/29012018192749ima1.jpg','teste 3'),(5,1,'290120181955282.jpg','uploadfile/290120181955282.jpg','teste 4'),(6,8,'30012018123627Captura de tela_2018-01-29_19-56-56.png','uploadfile/30012018123627Captura de tela_2018-01-29_19-56-56.png','Cassada #123 ç @345'),(7,8,'30012018140226Captura de tela_2018-01-04_04-36-09.png','uploadfile/30012018140226Captura de tela_2018-01-04_04-36-09.png','@Palhaço R$12.23 reais'),(8,1,'07022018165953Captura de tela_2018-01-04_04-36-09.png','uploadfile/07022018165953Captura de tela_2018-01-04_04-36-09.png','capa'),(9,1,'23022018164444WxH.jpg','uploadfile/23022018164444WxH.jpg','teste'),(10,9,'23022018164933WxH.jpg','uploadfile/23022018164933WxH.jpg','COMPARIO'),(11,9,'23022018164958Captura de tela_2018-01-29_19-56-46.png','uploadfile/23022018164958Captura de tela_2018-01-29_19-56-46.png','COMENTARIO 2'),(13,1,'23022018181647Captura de tela_2018-01-29_19-56-46.png','uploadfile/23022018181647Captura de tela_2018-01-29_19-56-46.png',NULL);
/*!40000 ALTER TABLE `ged` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `groups`
--

DROP TABLE IF EXISTS `groups`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `groups` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `groups`
--

LOCK TABLES `groups` WRITE;
/*!40000 ALTER TABLE `groups` DISABLE KEYS */;
INSERT INTO `groups` VALUES (1,'Conteúdo'),(2,'Blog');
/*!40000 ALTER TABLE `groups` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `log_access`
--

DROP TABLE IF EXISTS `log_access`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `log_access` (
  `id` int(11) NOT NULL,
  `user` varchar(45) DEFAULT NULL,
  `tipo` varchar(45) DEFAULT NULL,
  `status` enum('Sucesso','Falha') DEFAULT NULL,
  `data` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log_access`
--

LOCK TABLES `log_access` WRITE;
/*!40000 ALTER TABLE `log_access` DISABLE KEYS */;
/*!40000 ALTER TABLE `log_access` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `email` varchar(45) NOT NULL,
  `login` varchar(45) NOT NULL,
  `pass` varchar(255) NOT NULL,
  `type` enum('Admin','User') DEFAULT NULL,
  `hash` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'DIEGO DOS SANTOS','tec.infor321@gmail.com','diego','$2a$14$vcu9xHGWotjQFcYDtYoYee3poF893Wr8JEV9BZ/XxIt.yoGt5.1i6','Admin','dlklsdkldksdkldklsksd'),(2,'WANDER DOUGLAS','wander.gouglas@gmail.com','wander','1234','Admin','dlklsdkldksdkldklsksd'),(4,'josá da silva mentes','jose@gmail.com','jose','12345','User','dlklsdkldksdkldklsksd'),(6,'TESTE','teste@teste.com.br','teste2','321','Admin','dlklsdkldksdkldklsksd'),(7,'flavio','flavio@teste.com.br','flavio','$2a$14$/Gy/uGUrvadjB62YF4aiHOiwVblSdFcBRJC9.E4pIAuk9/.O80BzO','Admin','dlklsdkldksdkldklsksd'),(8,'saulo','saulo@wsitebrasil.com.br','saulo','$2a$14$B/pMwBbvAgo4U1Esw3VRV.yQWC25/K1J1733koP.Y4idFdpxFrddG','User','dlklsdkldksdkldklsksd');
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

-- Dump completed on 2018-03-15 22:41:57
