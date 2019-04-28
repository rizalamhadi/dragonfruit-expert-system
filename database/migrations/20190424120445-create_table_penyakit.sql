
-- +migrate Up
CREATE TABLE `penyakit` (
  `id_penyakit` char(6) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `nama_penyakit` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `solusi_penyakit` varchar(500) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`id_penyakit`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- +migrate StatementBegin
INSERT INTO `penyakit` VALUES
  ('PK0001','Karat Merah Alga','Menunas sulur secara teratur dan melakukan penyemprotan dengan fungisida tembaga dengan dosis 2.5-5 gram/2 liter air dengan interval penyemprotan 1 minggu.'),
  ('PK0002','Bercak Orange Sulur','Lakukan eradikasi (pemusnahan total sampai ke akar) pada tanaman yang terserang penyakit. Menggunakan agens hayati Gliocladium sp dalam kompos, yang diberikan dalam lubang tanam pada saat penanaman. Pengapuran atau pemberian abu dapur untuk menaikkan atau menjaga kestabilan pH tanah, dan penggunaan media ampas tebu yang ditambah urea dapat mengurangi perkembangan organism pathogen.  '),
  ('PK0003','Putih Sulur','Lakukan penyemprotan menggunakan nutrisi organik secara rutin dengan interval 7 hari sekali atau melakukan penyemprotan pestisida nabati seperti nimba, tagetes, eceng gondok atau rumput laut dilakukan secara rutin dengan interval 3-4 hari sekali.'),
  ('PK0004','Hawar Sulur','Menyemprotkan benlate dengan dosis 2 gram/liter air, atau menggunakan redomile 2 gram/liter air sebulan sekali. Bila muncul gejala kekuningan pada pangkal batang maka segera dilakukan penyemprotan pada seluruh batang dan diutamakan pada pangkal yang terserang.'),
  ('PK0005','Kusam Putih Sulur','Penanganan dengan menyemprotkan, pungisida dethane manzante atau pigune 3 kali seminggu. Memangkas tangkai bunga dan bagian tanaman yang terserang.'),
  ('PK0006','Busuk Lunak Batang','Penanganannya dengan memotong buah, yang terserang. penyemprotan kanon dengan dosis 1-2 cc atau liter air seminggu sekali disekitar buah yang terserang'),
  ('PK0007','Kuning Sulur','Penanganan dengan menyemprotkan, pungisida dethane manzante atau pigune 3 kali seminggu. memangkas tangkai bunga dan bagian tanaman yang terserang.'),
  ('PK0008','Antraknosa Buah','Penanganannya dengan memotong buah, yang terserang. penyemprotan kanon dengan dosis 1-2 cc atau liter air seminggu sekali disekitar buah yang terserang'),
  ('PK0009','Bercak Orange Buah','Meyemprotkan benlate dengan dosis 2 gram/liter air dalam seminggu 1-2 kali, penyemprotan pada bagian batang dan cabang.'),
  ('PK0010','Bintik Hitam pada Sulur','Untuk pengendaliannya dengan pemberian pupuk kandang secara rutin dengan interval 2 minggu sekali dan memperbaiki irigasi. Lakukan penyemprotan fungisida pada tanaman dan diutamakan pada bagian yang terserang penyakit.'),
  ('PK0011','Fusarium','Meyemprotkan benlate dengan dosis 2 gram/liter air dalam seminggu 1-2 kalli, penyemprotan pada bagian batang dan cabang.'),
  ('PK0012','Busuk Pangkal Batang','Untuk pengendalian pada penyakit ini dapat dilakukan dengan penyemprotan Benlate dosis 2 g/liter air atau dengan menggunakan Ridomil 2 g/liter air sebulan sekali. Bila muncul gejala kekuningan pada pangkal batang maka segera dilakukan penyemprotan pada seluruh batang dan diutamakan pada pangkal batang yang terserang.'),
  ('PK0013','Busuk Bakteri','Untuk pengendalian penyakit ini dapat dilakukan dengan mencabut tanaman yang sakit kemudian pada lubang tanam diberi Basamid dengan dosis 0,5-1 gram dalam bentuk serbuk kemudian pada lubang tanam tersebut ditanam bibit baru.');
-- +migrate StatementEnd

-- +migrate Down
DROP TABLE IF EXISTS `penyakit`;
