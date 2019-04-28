
-- +migrate Up
CREATE TABLE `gejala` (
  `id_gejala` char(6) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `id_jenis_gejala` char(6) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `gejala` varchar(200) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`id_gejala`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- +migrate StatementBegin
INSERT INTO `gejala` VALUES
  ('G0001','JG0001','Sulur melepuh'),
  ('G0002','JG0001','Sulur berkerut dan layu'),
  ('G0003','JG0001','Sulur layu dan kusam'),
  ('G0004','JG0001','Sulur berlendir'),
  ('G0005','JG0001','Muncul karat kecokelatan di permukaan sulur'),
  ('G0006','JG0001','Terdapat lapisan putih pada sulur'),
  ('G0007','JG0001','Menguning di bagian tengah dan berwarna hijau di tepi sulur'),
  ('G0008','JG0001','Menguning di seluruh bagian'),
  ('G0009','JG0001','Permukaan sulur berwarna cokelat'),
  ('G0010','JG0001','Sulur berwarna putih kekuningan'),
  ('G0011','JG0002','Terdapat bercak orange tidak beraturan pada sulur'),
  ('G0012','JG0002','Terdapat bercak putih membesar pada sulur'),
  ('G0013','JG0002','Terdapat bercak hitam dan cokelat pada sulur'),
  ('G0014','JG0002','Bercak menjadi berlendir'),
  ('G0015','JG0002','Terdapat bercak berair berwarna cokelat berukuran kecil pada sulur'),
  ('G0016','JG0003','Terdapat bercak berwarna orange seperti besi berkarat di kulit buah'),
  ('G0017','JG0003','Terdapat bercak hitam dan cokelat pada kulit buah'),
  ('G0018','JG0004','Terdapat bintik-bintik hitam tersusun beraturan'),
  ('G0019','JG0004','Terdapat bintik hitam atau cokelat'),
  ('G0020','JG0005','Tekstur seperti serbuk'),
  ('G0021','JG0005','Tekstur sulur yang terserang sangat berair dan mudah sobek'),
  ('G0022','JG0005','Permukaan sulur menonjol berwarna cokelat dan bagian pusat seperti berlubang'),
  ('G0023','JG0006','Terdapat lapisan putih pada batang'),
  ('G0024','JG0006','Batang berwarna kecokelatan'),
  ('G0025','JG0007','Buah menjadi busuk kering dan menghitam'),
  ('G0026','JG0007','Sulur membusuk'),
  ('G0027','JG0007','Busuk pada pangkal batang'),
  ('G0028','JG0008','Lapisan sulur mengelupas'),
  ('G0029','JG0008','Lapisan lilin dan daging sulur terkelupas'),
  ('G0030','JG0009','Bagian busuk lunak batang tercium bau tidak enak'),
  ('G0031','JG0010','Ukuran dan ketebalan sulur tidak berubah'),
  ('G0032','JG0010','Sulur menjadi sangat tipis');
-- +migrate StatementEnd

-- +migrate Down
DROP TABLE IF EXISTS `gejala`;
