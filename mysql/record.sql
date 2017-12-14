CREATE TABLE `record` (
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

/*
  userid 上傳者的ID
  organismname 生物名稱
  animal 是不是動物
  categorytagchinese 中文類別標籤：如青蛙、蝴蝶、植物
  categorytagenglish 英文類別標籤：如frog、butterfly、plant
  locationname 地名
  kingdom 界
  phylum  門
  class   綱
  order   目
  family  科
  genus   屬
  species 種
  stage   幼年、中年、老年
  season  出沒季節
  recorddate 上傳該記錄時可指定記錄的日期 這樣要搜是搜這個 還是mysql記錄的 createtime 是個問題
  time    上傳該記錄的日期跟時間
  note 額外的描述
  habitat 棲息地
*/