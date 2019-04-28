
-- +migrate Up
CREATE TABLE `rule_details` (
  `id_rule_detail` int(11) NOT NULL AUTO_INCREMENT,
  `id_rule` char(6) DEFAULT NULL,
  `id_gejala` char(6) DEFAULT NULL,
  PRIMARY KEY (`id_rule_detail`)
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +migrate StatementBegin
INSERT INTO `rule_details` VALUES
  (1,'R0001','G0001'),
  (2,'R0001','G0005'),
  (3,'R0002','G0003'),
  (4,'R0002','G0004'),
  (5,'R0002','G0010'),
  (6,'R0003','G0007'),
  (7,'R0003','G0031'),
  (8,'R0004','G0008'),
  (9,'R0004','G0032'),
  (10,'R0005','G0009'),
  (11,'R0005','G0002'),
  (12,'R0005','G0026'),
  (13,'R0006','G0009'),
  (14,'R0006','G0006'),
  (15,'R0006','G0028'),
  (16,'R0007','G0009'),
  (17,'R0007','G0012'),
  (18,'R0008','G0011'),
  (19,'R0008','G0019'),
  (20,'R0009','G0013'),
  (21,'R0009','G0014'),
  (22,'R0009','G0018'),
  (23,'R0010','G0015'),
  (24,'R0010','G0021'),
  (25,'R0010','G0029'),
  (26,'R0010','G0030'),
  (27,'R0011','G0016'),
  (28,'R0012','G0017'),
  (29,'R0012','G0018'),
  (30,'R0012','G0025'),
  (31,'R0013','G0018'),
  (32,'R0013','G0020'),
  (33,'R0014','G0018'),
  (34,'R0014','G0022'),
  (35,'R0015','G0023'),
  (36,'R0015','G0024'),
  (37,'R0015','G0027');
-- +migrate StatementEnd

-- +migrate Down
DROP TABLE IF EXISTS `rule_details`;
