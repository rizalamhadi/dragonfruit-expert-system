
-- +migrate Up
CREATE TABLE `jenis_gejala` (
  `id_jenis_gejala` char(6) NOT NULL,
  `jenis_gejala` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id_jenis_gejala`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +migrate StatementBegin
INSERT INTO `jenis_gejala` VALUES
  ('JG0001','Sulur'),
  ('JG0002','Bercak Sulur'),
  ('JG0003','Bercak Buah'),
  ('JG0004','Bintik'),
  ('JG0005','Tekstur'),
  ('JG0006','Batang'),
  ('JG0007','Pembusukan'),
  ('JG0008','Pengelupasan'),
  ('JG0009','Bau'),
  ('JG0010','Ukuran Sulur');
-- +migrate StatementEnd

-- +migrate Down
DROP TABLE IF EXISTS `jenis_gejala`;
