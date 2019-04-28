
-- +migrate Up
CREATE TABLE `rules` (
  `id_rules` char(6) NOT NULL,
  `id_penyakit` char(6) DEFAULT NULL,
  PRIMARY KEY (`id_rules`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +migrate StatementBegin
INSERT INTO `rules` VALUES
  ('R0001','PK0001'),
  ('R0002','PK0013'),
  ('R0003','PK0007'),
  ('R0004','PK0007'),
  ('R0005','PK0011'),
  ('R0006','PK0003'),
  ('R0007','PK0003'),
  ('R0008','PK0002'),
  ('R0009','PK0004'),
  ('R0010','PK0006'),
  ('R0011','PK0009'),
  ('R0012','PK0008'),
  ('R0013','PK0005'),
  ('R0014','PK0010'),
  ('R0015','PK0012');
-- +migrate StatementEnd

-- +migrate Down
DROP TABLE IF EXISTS `rules`;
