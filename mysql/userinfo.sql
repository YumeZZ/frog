CREATE TABLE `userinfo` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL,
  `username` varchar(30) COLLATE utf8_unicode_ci NOT NULL,
  `password` varchar(200) COLLATE utf8_unicode_ci NOT NULL,
  `surname` varchar(15) COLLATE utf8_unicode_ci DEFAULT NULL,
  `givenname` varchar(15) COLLATE utf8_unicode_ci DEFAULT NULL,
  `nickname` varchar(15) COLLATE utf8_unicode_ci DEFAULT NULL,
  `phone` varchar(15) COLLATE utf8_unicode_ci DEFAULT NULL,
  `address` text DEFAULT NULL COLLATE utf8_unicode_ci,
  `birthday` date DEFAULT NULL,
  `nationality` varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL,
  `gender` varchar(15) COLLATE utf8_unicode_ci DEFAULT NULL,
  `religion` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `age` int(11) DEFAULT NULL,
  `vegetarian` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `createtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(id),
  UNIQUE KEY(username),
  UNIQUE KEY(email),
  UNIQUE KEY(phone)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

insert into userinfo(username,password) values('testun','68656c6c6f70773230313708a2f8c9e2fce8e8eb5f181ee569ed836b9d4dfb4e0fbbed24d816b8e0549f2de004ebba08ed4f5fbd463fe390bd827e9b496c109f26440fb942cd20f3a3579a');

/*
hellopw2017
*/