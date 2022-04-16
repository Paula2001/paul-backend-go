-- MySQL dump 10.13  Distrib 8.0.28, for Linux (x86_64)
--
-- Host: localhost    Database: test
-- ------------------------------------------------------
-- Server version	8.0.28

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
-- Table structure for table `countries`
--

DROP TABLE IF EXISTS `countries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `countries` (
  `id` int NOT NULL AUTO_INCREMENT,
  `iso2` varchar(2) DEFAULT NULL,
  `iso3` varchar(3) DEFAULT NULL,
  `numcode` int DEFAULT NULL,
  `long_name` varchar(100) DEFAULT NULL,
  `short_name` varchar(75) DEFAULT NULL,
  `is_supported` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `numcode` (`numcode`)
) ENGINE=InnoDB AUTO_INCREMENT=251 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `countries`
--

LOCK TABLES `countries` WRITE;
/*!40000 ALTER TABLE `countries` DISABLE KEYS */;
INSERT INTO `countries` VALUES (1,'AF','AFG',4,'Islamic Republic of Afghanistan','Afghanistan',1),(2,'AX','ALA',248,'&Aring;land Islands','Aland Islands',0),(3,'AL','ALB',8,'Republic of Albania','Albania',1),(4,'DZ','DZA',12,'People\'s Democratic Republic of Algeria','Algeria',1),(5,'AS','ASM',16,'American Samoa','American Samoa',0),(6,'AD','AND',20,'Principality of Andorra','Andorra',1),(7,'AO','AGO',24,'Republic of Angola','Angola',1),(8,'AI','AIA',660,'Anguilla','Anguilla',0),(9,'AQ','ATA',10,'Antarctica','Antarctica',0),(10,'AG','ATG',28,'Antigua and Barbuda','Antigua and Barbuda',1),(11,'AR','ARG',32,'Argentine Republic','Argentina',1),(12,'AM','ARM',51,'Republic of Armenia','Armenia',1),(13,'AW','ABW',533,'Aruba','Aruba',0),(14,'AU','AUS',36,'Commonwealth of Australia','Australia',1),(15,'AT','AUT',40,'Republic of Austria','Austria',1),(16,'AZ','AZE',31,'Republic of Azerbaijan','Azerbaijan',1),(17,'BS','BHS',44,'Commonwealth of The Bahamas','Bahamas',1),(18,'BH','BHR',48,'Kingdom of Bahrain','Bahrain',1),(19,'BD','BGD',50,'People\'s Republic of Bangladesh','Bangladesh',1),(20,'BB','BRB',52,'Barbados','Barbados',1),(21,'BY','BLR',112,'Republic of Belarus','Belarus',1),(22,'BE','BEL',56,'Kingdom of Belgium','Belgium',1),(23,'BZ','BLZ',84,'Belize','Belize',1),(24,'BJ','BEN',204,'Republic of Benin','Benin',1),(25,'BM','BMU',60,'Bermuda Islands','Bermuda',0),(26,'BT','BTN',64,'Kingdom of Bhutan','Bhutan',1),(27,'BO','BOL',68,'Plurinational State of Bolivia','Bolivia',1),(28,'BQ','BES',535,'Bonaire, Sint Eustatius and Saba','Bonaire, Sint Eustatius and Saba',0),(29,'BA','BIH',70,'Bosnia and Herzegovina','Bosnia and Herzegovina',1),(30,'BW','BWA',72,'Republic of Botswana','Botswana',1),(31,'BV','BVT',74,'Bouvet Island','Bouvet Island',0),(32,'BR','BRA',76,'Federative Republic of Brazil','Brazil',1),(33,'IO','IOT',86,'British Indian Ocean Territory','British Indian Ocean Territory',0),(34,'BN','BRN',96,'Brunei Darussalam','Brunei',1),(35,'BG','BGR',100,'Republic of Bulgaria','Bulgaria',1),(36,'BF','BFA',854,'Burkina Faso','Burkina Faso',1),(37,'BI','BDI',108,'Republic of Burundi','Burundi',1),(38,'KH','KHM',116,'Kingdom of Cambodia','Cambodia',1),(39,'CM','CMR',120,'Republic of Cameroon','Cameroon',1),(40,'CA','CAN',124,'Canada','Canada',1),(41,'CV','CPV',132,'Republic of Cape Verde','Cape Verde',1),(42,'KY','CYM',136,'The Cayman Islands','Cayman Islands',0),(43,'CF','CAF',140,'Central African Republic','Central African Republic',1),(44,'TD','TCD',148,'Republic of Chad','Chad',1),(45,'CL','CHL',152,'Republic of Chile','Chile',1),(46,'CN','CHN',156,'People\'s Republic of China','China',1),(47,'CX','CXR',162,'Christmas Island','Christmas Island',0),(48,'CC','CCK',166,'Cocos (Keeling) Islands','Cocos (Keeling) Islands',0),(49,'CO','COL',170,'Republic of Colombia','Colombia',1),(50,'KM','COM',174,'Union of the Comoros','Comoros',1),(51,'CG','COG',178,'Republic of the Congo','Congo',1),(52,'CK','COK',184,'Cook Islands','Cook Islands',0),(53,'CR','CRI',188,'Republic of Costa Rica','Costa Rica',1),(54,'CI','CIV',384,'Republic of C&ocirc;te D\'Ivoire (Ivory Coast)','Cote d\'ivoire (Ivory Coast)',1),(55,'HR','HRV',191,'Republic of Croatia','Croatia',1),(56,'CU','CUB',192,'Republic of Cuba','Cuba',1),(57,'CW','CUW',531,'Cura&ccedil;ao','Curacao',0),(58,'CY','CYP',196,'Republic of Cyprus','Cyprus',1),(59,'CZ','CZE',203,'Czech Republic','Czech Republic',1),(60,'CD','COD',180,'Democratic Republic of the Congo','Democratic Republic of the Congo',1),(61,'DK','DNK',208,'Kingdom of Denmark','Denmark',1),(62,'DJ','DJI',262,'Republic of Djibouti','Djibouti',1),(63,'DM','DMA',212,'Commonwealth of Dominica','Dominica',1),(64,'DO','DOM',214,'Dominican Republic','Dominican Republic',1),(65,'EC','ECU',218,'Republic of Ecuador','Ecuador',1),(66,'EG','EGY',818,'Arab Republic of Egypt','Egypt',1),(67,'SV','SLV',222,'Republic of El Salvador','El Salvador',1),(68,'GQ','GNQ',226,'Republic of Equatorial Guinea','Equatorial Guinea',1),(69,'ER','ERI',232,'State of Eritrea','Eritrea',1),(70,'EE','EST',233,'Republic of Estonia','Estonia',1),(71,'ET','ETH',231,'Federal Democratic Republic of Ethiopia','Ethiopia',1),(72,'FK','FLK',238,'The Falkland Islands (Malvinas)','Falkland Islands (Malvinas)',0),(73,'FO','FRO',234,'The Faroe Islands','Faroe Islands',0),(74,'FJ','FJI',242,'Republic of Fiji','Fiji',1),(75,'FI','FIN',246,'Republic of Finland','Finland',1),(76,'FR','FRA',250,'French Republic','France',1),(77,'GF','GUF',254,'French Guiana','French Guiana',0),(78,'PF','PYF',258,'French Polynesia','French Polynesia',0),(79,'TF','ATF',260,'French Southern Territories','French Southern Territories',0),(80,'GA','GAB',266,'Gabonese Republic','Gabon',1),(81,'GM','GMB',270,'Republic of The Gambia','Gambia',1),(82,'GE','GEO',268,'Georgia','Georgia',1),(83,'DE','DEU',276,'Federal Republic of Germany','Germany',1),(84,'GH','GHA',288,'Republic of Ghana','Ghana',1),(85,'GI','GIB',292,'Gibraltar','Gibraltar',0),(86,'GR','GRC',300,'Hellenic Republic','Greece',1),(87,'GL','GRL',304,'Greenland','Greenland',0),(88,'GD','GRD',308,'Grenada','Grenada',1),(89,'GP','GLP',312,'Guadeloupe','Guadaloupe',0),(90,'GU','GUM',316,'Guam','Guam',0),(91,'GT','GTM',320,'Republic of Guatemala','Guatemala',1),(92,'GG','GGY',831,'Guernsey','Guernsey',0),(93,'GN','GIN',324,'Republic of Guinea','Guinea',1),(94,'GW','GNB',624,'Republic of Guinea-Bissau','Guinea-Bissau',1),(95,'GY','GUY',328,'Co-operative Republic of Guyana','Guyana',1),(96,'HT','HTI',332,'Republic of Haiti','Haiti',1),(97,'HM','HMD',334,'Heard Island and McDonald Islands','Heard Island and McDonald Islands',0),(98,'HN','HND',340,'Republic of Honduras','Honduras',1),(99,'HK','HKG',344,'Hong Kong','Hong Kong',0),(100,'HU','HUN',348,'Hungary','Hungary',1),(101,'IS','ISL',352,'Republic of Iceland','Iceland',1),(102,'IN','IND',356,'Republic of India','India',1),(103,'ID','IDN',360,'Republic of Indonesia','Indonesia',1),(104,'IR','IRN',364,'Islamic Republic of Iran','Iran',1),(105,'IQ','IRQ',368,'Republic of Iraq','Iraq',1),(106,'IE','IRL',372,'Ireland','Ireland',1),(107,'IM','IMN',833,'Isle of Man','Isle of Man',0),(108,'IL','ISR',376,'State of Israel','Israel',1),(109,'IT','ITA',380,'Italian Republic','Italy',1),(110,'JM','JAM',388,'Jamaica','Jamaica',1),(111,'JP','JPN',392,'Japan','Japan',1),(112,'JE','JEY',832,'The Bailiwick of Jersey','Jersey',0),(113,'JO','JOR',400,'Hashemite Kingdom of Jordan','Jordan',1),(114,'KZ','KAZ',398,'Republic of Kazakhstan','Kazakhstan',1),(115,'KE','KEN',404,'Republic of Kenya','Kenya',1),(116,'KI','KIR',296,'Republic of Kiribati','Kiribati',1),(117,'XK','---',0,'Republic of Kosovo','Kosovo',0),(118,'KW','KWT',414,'State of Kuwait','Kuwait',1),(119,'KG','KGZ',417,'Kyrgyz Republic','Kyrgyzstan',1),(120,'LA','LAO',418,'Lao People\'s Democratic Republic','Laos',1),(121,'LV','LVA',428,'Republic of Latvia','Latvia',1),(122,'LB','LBN',422,'Republic of Lebanon','Lebanon',1),(123,'LS','LSO',426,'Kingdom of Lesotho','Lesotho',1),(124,'LR','LBR',430,'Republic of Liberia','Liberia',1),(125,'LY','LBY',434,'Libya','Libya',1),(126,'LI','LIE',438,'Principality of Liechtenstein','Liechtenstein',1),(127,'LT','LTU',440,'Republic of Lithuania','Lithuania',1),(128,'LU','LUX',442,'Grand Duchy of Luxembourg','Luxembourg',1),(129,'MO','MAC',446,'The Macao Special Administrative Region','Macao',0),(130,'MK','MKD',807,'The Former Yugoslav Republic of Macedonia','Macedonia',1),(131,'MG','MDG',450,'Republic of Madagascar','Madagascar',1),(132,'MW','MWI',454,'Republic of Malawi','Malawi',1),(133,'MY','MYS',458,'Malaysia','Malaysia',1),(134,'MV','MDV',462,'Republic of Maldives','Maldives',1),(135,'ML','MLI',466,'Republic of Mali','Mali',1),(136,'MT','MLT',470,'Republic of Malta','Malta',1),(137,'MH','MHL',584,'Republic of the Marshall Islands','Marshall Islands',1),(138,'MQ','MTQ',474,'Martinique','Martinique',0),(139,'MR','MRT',478,'Islamic Republic of Mauritania','Mauritania',1),(140,'MU','MUS',480,'Republic of Mauritius','Mauritius',1),(141,'YT','MYT',175,'Mayotte','Mayotte',0),(142,'MX','MEX',484,'United Mexican States','Mexico',1),(143,'FM','FSM',583,'Federated States of Micronesia','Micronesia',1),(144,'MD','MDA',498,'Republic of Moldova','Moldava',1),(145,'MC','MCO',492,'Principality of Monaco','Monaco',1),(146,'MN','MNG',496,'Mongolia','Mongolia',1),(147,'ME','MNE',499,'Montenegro','Montenegro',1),(148,'MS','MSR',500,'Montserrat','Montserrat',0),(149,'MA','MAR',504,'Kingdom of Morocco','Morocco',1),(150,'MZ','MOZ',508,'Republic of Mozambique','Mozambique',1),(151,'MM','MMR',104,'Republic of the Union of Myanmar','Myanmar (Burma)',1),(152,'NA','NAM',516,'Republic of Namibia','Namibia',1),(153,'NR','NRU',520,'Republic of Nauru','Nauru',1),(154,'NP','NPL',524,'Federal Democratic Republic of Nepal','Nepal',1),(155,'NL','NLD',528,'Kingdom of the Netherlands','Netherlands',1),(156,'NC','NCL',540,'New Caledonia','New Caledonia',0),(157,'NZ','NZL',554,'New Zealand','New Zealand',1),(158,'NI','NIC',558,'Republic of Nicaragua','Nicaragua',1),(159,'NE','NER',562,'Republic of Niger','Niger',1),(160,'NG','NGA',566,'Federal Republic of Nigeria','Nigeria',1),(161,'NU','NIU',570,'Niue','Niue',0),(162,'NF','NFK',574,'Norfolk Island','Norfolk Island',0),(163,'KP','PRK',408,'Democratic People\'s Republic of Korea','North Korea',1),(164,'MP','MNP',580,'Northern Mariana Islands','Northern Mariana Islands',0),(165,'NO','NOR',578,'Kingdom of Norway','Norway',1),(166,'OM','OMN',512,'Sultanate of Oman','Oman',1),(167,'PK','PAK',586,'Islamic Republic of Pakistan','Pakistan',1),(168,'PW','PLW',585,'Republic of Palau','Palau',1),(169,'PS','PSE',275,'State of Palestine (or Occupied Palestinian Territory)','Palestine',0),(170,'PA','PAN',591,'Republic of Panama','Panama',1),(171,'PG','PNG',598,'Independent State of Papua New Guinea','Papua New Guinea',1),(172,'PY','PRY',600,'Republic of Paraguay','Paraguay',1),(173,'PE','PER',604,'Republic of Peru','Peru',1),(174,'PH','PHL',608,'Republic of the Philippines','Phillipines',1),(175,'PN','PCN',612,'Pitcairn','Pitcairn',0),(176,'PL','POL',616,'Republic of Poland','Poland',1),(177,'PT','PRT',620,'Portuguese Republic','Portugal',1),(178,'PR','PRI',630,'Commonwealth of Puerto Rico','Puerto Rico',0),(179,'QA','QAT',634,'State of Qatar','Qatar',1),(180,'RE','REU',638,'R&eacute;union','Reunion',0),(181,'RO','ROU',642,'Romania','Romania',1),(182,'RU','RUS',643,'Russian Federation','Russia',1),(183,'RW','RWA',646,'Republic of Rwanda','Rwanda',1),(184,'BL','BLM',652,'Saint Barth&eacute;lemy','Saint Barthelemy',0),(185,'SH','SHN',654,'Saint Helena, Ascension and Tristan da Cunha','Saint Helena',0),(186,'KN','KNA',659,'Federation of Saint Christopher and Nevis','Saint Kitts and Nevis',1),(187,'LC','LCA',662,'Saint Lucia','Saint Lucia',1),(188,'MF','MAF',663,'Saint Martin','Saint Martin',0),(189,'PM','SPM',666,'Saint Pierre and Miquelon','Saint Pierre and Miquelon',0),(190,'VC','VCT',670,'Saint Vincent and the Grenadines','Saint Vincent and the Grenadines',1),(191,'WS','WSM',882,'Independent State of Samoa','Samoa',1),(192,'SM','SMR',674,'Republic of San Marino','San Marino',1),(193,'ST','STP',678,'Democratic Republic of S&atilde;o Tom&eacute; and Pr&iacute;ncipe','Sao Tome and Principe',1),(194,'SA','SAU',682,'Kingdom of Saudi Arabia','Saudi Arabia',1),(195,'SN','SEN',686,'Republic of Senegal','Senegal',1),(196,'RS','SRB',688,'Republic of Serbia','Serbia',1),(197,'SC','SYC',690,'Republic of Seychelles','Seychelles',1),(198,'SL','SLE',694,'Republic of Sierra Leone','Sierra Leone',1),(199,'SG','SGP',702,'Republic of Singapore','Singapore',1),(200,'SX','SXM',534,'Sint Maarten','Sint Maarten',0),(201,'SK','SVK',703,'Slovak Republic','Slovakia',1),(202,'SI','SVN',705,'Republic of Slovenia','Slovenia',1),(203,'SB','SLB',90,'Solomon Islands','Solomon Islands',1),(204,'SO','SOM',706,'Somali Republic','Somalia',1),(205,'ZA','ZAF',710,'Republic of South Africa','South Africa',1),(206,'GS','SGS',239,'South Georgia and the South Sandwich Islands','South Georgia and the South Sandwich Islands',0),(207,'KR','KOR',410,'Republic of Korea','South Korea',1),(208,'SS','SSD',728,'Republic of South Sudan','South Sudan',1),(209,'ES','ESP',724,'Kingdom of Spain','Spain',1),(210,'LK','LKA',144,'Democratic Socialist Republic of Sri Lanka','Sri Lanka',1),(211,'SD','SDN',729,'Republic of the Sudan','Sudan',1),(212,'SR','SUR',740,'Republic of Suriname','Suriname',1),(213,'SJ','SJM',744,'Svalbard and Jan Mayen','Svalbard and Jan Mayen',0),(214,'SZ','SWZ',748,'Kingdom of Swaziland','Swaziland',1),(215,'SE','SWE',752,'Kingdom of Sweden','Sweden',1),(216,'CH','CHE',756,'Swiss Confederation','Switzerland',1),(217,'SY','SYR',760,'Syrian Arab Republic','Syria',1),(218,'TW','TWN',158,'Republic of China (Taiwan)','Taiwan',0),(219,'TJ','TJK',762,'Republic of Tajikistan','Tajikistan',1),(220,'TZ','TZA',834,'United Republic of Tanzania','Tanzania',1),(221,'TH','THA',764,'Kingdom of Thailand','Thailand',1),(222,'TL','TLS',626,'Democratic Republic of Timor-Leste','Timor-Leste (East Timor)',1),(223,'TG','TGO',768,'Togolese Republic','Togo',1),(224,'TK','TKL',772,'Tokelau','Tokelau',0),(225,'TO','TON',776,'Kingdom of Tonga','Tonga',1),(226,'TT','TTO',780,'Republic of Trinidad and Tobago','Trinidad and Tobago',1),(227,'TN','TUN',788,'Republic of Tunisia','Tunisia',1),(228,'TR','TUR',792,'Republic of Turkey','Turkey',1),(229,'TM','TKM',795,'Turkmenistan','Turkmenistan',1),(230,'TC','TCA',796,'Turks and Caicos Islands','Turks and Caicos Islands',0),(231,'TV','TUV',798,'Tuvalu','Tuvalu',1),(232,'UG','UGA',800,'Republic of Uganda','Uganda',1),(233,'UA','UKR',804,'Ukraine','Ukraine',1),(234,'AE','ARE',784,'United Arab Emirates','United Arab Emirates',1),(235,'GB','GBR',826,'United Kingdom of Great Britain and Nothern Ireland','United Kingdom',1),(236,'US','USA',840,'United States of America','United States',1),(237,'UM','UMI',581,'United States Minor Outlying Islands','United States Minor Outlying Islands',0),(238,'UY','URY',858,'Eastern Republic of Uruguay','Uruguay',1),(239,'UZ','UZB',860,'Republic of Uzbekistan','Uzbekistan',1),(240,'VU','VUT',548,'Republic of Vanuatu','Vanuatu',1),(241,'VA','VAT',336,'State of the Vatican City','Vatican City',0),(242,'VE','VEN',862,'Bolivarian Republic of Venezuela','Venezuela',1),(243,'VN','VNM',704,'Socialist Republic of Vietnam','Vietnam',1),(244,'VG','VGB',92,'British Virgin Islands','Virgin Islands, British',0),(245,'VI','VIR',850,'Virgin Islands of the United States','Virgin Islands, US',0),(246,'WF','WLF',876,'Wallis and Futuna','Wallis and Futuna',0),(247,'EH','ESH',732,'Western Sahara','Western Sahara',0),(248,'YE','YEM',887,'Republic of Yemen','Yemen',1),(249,'ZM','ZMB',894,'Republic of Zambia','Zambia',1),(250,'ZW','ZWE',716,'Republic of Zimbabwe','Zimbabwe',1);
/*!40000 ALTER TABLE `countries` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `testimonies`
--

DROP TABLE IF EXISTS `testimonies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `testimonies` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(200) DEFAULT NULL,
  `testimony` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `testimonies`
--

LOCK TABLES `testimonies` WRITE;
/*!40000 ALTER TABLE `testimonies` DISABLE KEYS */;
INSERT INTO `testimonies` VALUES (10,'asd','Good'),(11,'asd','asddd');
/*!40000 ALTER TABLE `testimonies` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-16 11:16:25