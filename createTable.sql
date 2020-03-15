CREATE TABLE `user` (
 `user_id` int(11) NOT NULL AUTO_INCREMENT,
 `username` varchar(25) NOT NULL,
 `password` varchar(25) NOT NULL,
 `usertype` enum('staff','admin') NOT NULL DEFAULT 'staff',
 `email` varchar(25) NOT NULL,
 PRIMARY KEY (`user_id`),
 UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1