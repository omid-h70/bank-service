DROP DATABASE IF EXISTS webServiceDB;
CREATE DATABASE webServiceDB;
USE webServiceDB;

DROP TABLE IF EXISTS `customer`;
CREATE TABLE `customer`(
`customer_id` int(11) NOT NULL AUTO_INCREMENT,
`name` varchar(100) NOT NULL,
`date_of_birth` date NOT NULL,
`phone_number` varchar(16) NOT NULL,
#`zipcode` varchar(100) NOT NULL,
`status` tinyint(1) NOT NULL DEFAULT '1',
PRIMARY KEY (`customer_id`)
)ENGINE=InnoDB AUTO_INCREMENT = 1006 DEFAULT CHARSET=latin1;

INSERT INTO `customer` VALUES
(1001, 'John Doe', '2023-09-01', 'tehran',1),
(1002, 'John Nash', '2023-09-02', 'tehran',1),
(1003, 'Omid H', '2023-09-03', 'tehran',0),
(1004, 'Foo Bar', '2023-09-04', 'tehran',1),
(1005, 'Derp', '2023-09-05', 'tehran',0);

DROP TABLE IF EXISTS `account`;
CREATE TABLE `account`(
`account_id` int(11) NOT NULL AUTO_INCREMENT,
`customer_id` int(11) NOT NULL,
`opening_date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
`account_type` varchar(10) NOT NULL,
`pin` varchar(10) NOT NULL,
`status` tinyint(1) NOT NULL DEFAULT '1',
PRIMARY KEY (`account_id`),
KEY `account_FK` (`customer_id`),
CONSTRAINT `account_FK` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`customer_id`)
)ENGINE=InnoDB AUTO_INCREMENT = 95475 DEFAULT CHARSET=latin1;

INSERT INTO `account` VALUES
(95470, 1001, '2023-09-01 00:00:00', 'saving','1075',1),
(95471, 1002, '2023-09-02 00:00:01', 'saving','1111',1),
(95472, 1003, '2023-09-03 00:00:02', 'checking','1234',1),
(95473, 1004, '2023-09-04 00:00:03', 'saving','4567',1),
(95474, 1005, '2023-09-05 00:00:04', 'saving','9876',0);

DROP TABLE IF EXISTS `transaction`;
CREATE TABLE `transaction`(
`transaction_id` int(11) NOT NULL AUTO_INCREMENT,
`account_id` int(11) NOT NULL,
`amount` int(11) NOT NULL,
`transaction_type` varchar(100) NOT NULL,
`transaction_date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`transaction_id`),
KEY `transaction_FK` (`account_id`),
CONSTRAINT `transaction_FK` FOREIGN KEY (`account_id`) REFERENCES `account` (`account_id`)
)ENGINE=InnoDB AUTO_INCREMENT = 1006 DEFAULT CHARSET=latin1;

\! echo "InitDB Done >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>"