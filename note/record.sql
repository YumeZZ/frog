﻿CREATE TABLE `record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `userid` int(11) NOT NULL,
  `organismname` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `animal` tinyint(1) DEFAULT NULL,
  `categorytagchinese` varchar(30) COLLATE utf8_unicode_ci DEFAULT NULL,
  `categorytagenglish` varchar(30) COLLATE utf8_unicode_ci DEFAULT NULL,
  `locationname` varchar(30) COLLATE utf8_unicode_ci DEFAULT NULL,
  `kingdom` varchar(30) COLLATE utf8_unicode_ci DEFAULT NULL,
  `phylum` varchar(30) COLLATE utf8_unicode_ci DEFAULT NULL,
  `class` varchar(30) COLLATE utf8_unicode_ci DEFAULT NULL,
  `order` varchar(30) COLLATE utf8_unicode_ci DEFAULT NULL,
  `family` varchar(30) COLLATE utf8_unicode_ci DEFAULT NULL,
  `genus` varchar(30) COLLATE utf8_unicode_ci DEFAULT NULL,
  `species` varchar(30) COLLATE utf8_unicode_ci DEFAULT NULL,
  `food` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL,
  `stage` varchar(15) COLLATE utf8_unicode_ci DEFAULT NULL,
  `season` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `recorddate` DATE DEFAULT NULL,
  `time` datetime DEFAULT NULL,
  `note` TEXT DEFAULT NULL,
  `habitat` TEXT DEFAULT NULL,
  `createtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;