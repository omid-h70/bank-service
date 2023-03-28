DROP DATABASE IF EXISTS webServiceDB;
CREATE DATABASE webServiceDB;
USE webServiceDB;

DROP TABLE IF EXISTS `customer`;
CREATE TABLE `customer`(
`customer_id` int(11) NOT NULL AUTO_INCREMENT,
`name` varchar(100) NOT NULL,
`phone_number` varchar(16) NOT NULL,
`status` tinyint(1) NOT NULL DEFAULT '1',
PRIMARY KEY (`customer_id`)
)ENGINE=InnoDB AUTO_INCREMENT = 1006 DEFAULT CHARSET=latin1;

INSERT INTO `customer` VALUES
(1001, 'John Doe',  'tehran',1),
(1002, 'Mammad',    'tehran',1),
(1003, 'Omid H',    'tehran',1),
(1004, 'Derpina',   'tehran',1),
(1005, 'Derp',      'tehran',1);

DROP TABLE IF EXISTS `account`;
CREATE TABLE `account`(
`account_id` int(11) NOT NULL AUTO_INCREMENT,
`customer_id` int(11) NOT NULL,
`opening_date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
`account_balance` double NOT NULL,
`pin` varchar(10) NOT NULL,
`status` tinyint(1) NOT NULL DEFAULT '1',
PRIMARY KEY (`account_id`),
KEY `account_FK` (`customer_id`),
CONSTRAINT `account_FK` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`customer_id`)
)ENGINE=InnoDB AUTO_INCREMENT = 2006 DEFAULT CHARSET=latin1;

INSERT INTO `account` VALUES
(2001, 1001, '2023-09-01 00:00:00', 'saving','1075',1),
(2002, 1002, '2023-09-02 00:00:01', 'saving','1111',1),
(2003, 1003, '2023-09-03 00:00:02', 'checking','1234',1),
(2004, 1004, '2023-09-04 00:00:03', 'saving','4567',1),
(2005, 1005, '2023-09-05 00:00:04', 'saving','9876',0);

DROP TABLE IF EXISTS `card`;
CREATE TABLE `card`(
`id` int(5) NOT NULL AUTO_INCREMENT,
`account_id` varchar(17) NOT NULL,
`card_number` date NOT NULL,
PRIMARY KEY (`customer_id`)
)ENGINE=InnoDB AUTO_INCREMENT = 3005 DEFAULT CHARSET=latin1;

INSERT INTO `card` VALUES
(3001, 2001, '6280231133106101'),
(3002, 2001, '6280231033106293'),
(3003, 2003, '5022291302421266'),
(3004, 2003, '6221061060903186'),
(3005, 2003, '6063731136152064');

DROP TABLE IF EXISTS `transaction`;
CREATE TABLE `transaction`(
`transaction_id` int(11) NOT NULL AUTO_INCREMENT,
`account_id_from` int(11) NOT NULL,
`account_id_to` int(11) NOT NULL,
`amount` int(11) NOT NULL,
`transaction_type` varchar(100) NOT NULL,
`transaction_date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`transaction_id`),
KEY `transaction_FK` (`account_id`),
CONSTRAINT `transaction_FK` FOREIGN KEY (`account_id`) REFERENCES `account` (`account_id`)
)ENGINE=InnoDB AUTO_INCREMENT = 1006 DEFAULT CHARSET=latin1;

SELECT "InitDB Done >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>"