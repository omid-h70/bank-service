DROP DATABASE IF EXISTS webServiceDB;
CREATE DATABASE webServiceDB;
USE webServiceDB;

DROP TABLE IF EXISTS `customer`;
CREATE TABLE `customer`(
`customer_id` int(11) NOT NULL AUTO_INCREMENT,
`name` varchar(100) NOT NULL,
`phone_number` varchar(32) NOT NULL,
`status` tinyint(1) NOT NULL DEFAULT '1',
PRIMARY KEY (`customer_id`)
)ENGINE=InnoDB AUTO_INCREMENT = 1006 DEFAULT CHARSET=latin1;

INSERT INTO `customer` VALUES
(1001, 'John Doe',  '+989123993699',1),
(1002, 'Mammad',    '+989033934262',1),
(1003, 'Omid H',    '+989123993699',1),
(1004, 'Derpina',   '+989033934262',1),
(1005, 'Derp',      '+989123993699',1);

DROP TABLE IF EXISTS `account`;
CREATE TABLE `account`(
`account_id` int(11) NOT NULL AUTO_INCREMENT,
`customer_id` int(11) NOT NULL,
`account_rule_id` int(11) NOT NULL,
`opening_date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
`account_balance` bigint NOT NULL,
`pin` varchar(10) NOT NULL,
`status` tinyint(1) NOT NULL DEFAULT '1',
PRIMARY KEY (`account_id`),
KEY `account_FK` (`customer_id`),
CONSTRAINT `account_FK` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`customer_id`)
)ENGINE=InnoDB AUTO_INCREMENT = 2007 DEFAULT CHARSET=latin1;

INSERT INTO `account` VALUES
(2001, 1001, 3001, '2023-09-01 00:00:00', 500000,'1075',1),
(2002, 1002, 3001, '2023-09-02 00:00:01', 100000,'1111',1),
(2003, 1003, 3001, '2023-09-03 00:00:02', 150000,'1234',1),
(2004, 1004, 3001, '2023-09-04 00:00:03', 700000,'4567',1),
(2005, 1005, 3001, '2023-09-05 00:00:04', 800000,'9876',1),


DROP TABLE IF EXISTS `account_rule`;
CREATE TABLE `account_rule`(
`account_rule_id` int(5) NOT NULL AUTO_INCREMENT,
`min_amount` bigint NOT NULL,
`max_amount` bigint NOT NULL,
PRIMARY KEY (`account_rule_id`)
)ENGINE=InnoDB AUTO_INCREMENT = 3002 DEFAULT CHARSET=latin1;

INSERT INTO `account_rule` VALUES
(3001, 10000, 500000000);


DROP TABLE IF EXISTS `card`;
CREATE TABLE `card`(
`card_id` int(5) NOT NULL AUTO_INCREMENT,
`account_id` int(5) NOT NULL,
`card_number` varchar(17) NOT NULL,
PRIMARY KEY (`card_id`)
)ENGINE=InnoDB AUTO_INCREMENT = 4011 DEFAULT CHARSET=latin1;

INSERT INTO `card` VALUES
(4001, 2001, '6280231133106101'),
(4002, 2001, '6280231033106293'),
(4003, 2003, '5022291302421266'),
(4004, 2003, '6221061060903186'),
(4005, 2003, '6063731136152064'),
(4006, 2004, '1957202120304154'),
(4007, 2004, '5041721005782710'),
(4008, 2005, '6037994501250318'),
(4009, 2005, '6219861058787200'),
(4010, 2002, '6037991760382659');

DROP TABLE IF EXISTS `transaction`;
CREATE TABLE `transaction`(
`transaction_id` int(11) NOT NULL AUTO_INCREMENT,
`card_id_from` int(11) NOT NULL,
`card_id_to` int(11) NOT NULL,
`amount` bigint NOT NULL,
`transaction_type` int(2) NOT NULL,
`transaction_date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`transaction_id`),
KEY `transaction_FK_1` (`card_id_to`),
CONSTRAINT `transaction_FK_1` FOREIGN KEY (`card_id_to`) REFERENCES `card` (`card_id`)
)ENGINE=InnoDB AUTO_INCREMENT = 1006 DEFAULT CHARSET=latin1;

SELECT "InitDB Done >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>"