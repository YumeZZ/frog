CREATE TABLE `photo` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `userid` int(11) NOT NULL,
  `recordid` int(11) NOT NULL,
  `initorder` int(11) NOT NULL,
  `path` varchar(200) COLLATE utf8_unicode_ci NOT NULL,
  `name` varchar(200) COLLATE utf8_unicode_ci NOT NULL,
  `longitude` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `latitude` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `createtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

/*
  userid == 上傳該圖片的用戶ID
  recordid == 表示連同哪個生態記錄一起上傳
  initorder == 是該筆記錄的第幾張圖片
  path == 圖片存放的路徑
  name == 圖片未上傳前的檔名
  longitude == 圖片exif裡面的經度座標，尚未實作依據網頁上的輸入，修改照片本身的exif資料、mysql裡的資料
  latitude == 圖片exif裡面的緯度座標
  createtime == 上傳圖片的時間
*/