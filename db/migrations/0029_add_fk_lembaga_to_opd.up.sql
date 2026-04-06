-- Pastikan nilai kode_lembaga yang kosong set menjadi NULL
UPDATE opd SET kode_lembaga = NULL WHERE kode_lembaga = '';

-- Ijinkan kode_lembaga menjadi nullable dan hapus nilai default ''
ALTER TABLE opd
    ALTER COLUMN kode_lembaga DROP NOT NULL,
    ALTER COLUMN kode_lembaga DROP DEFAULT;

-- Pastikan nilai lembaga.kode_lembaga unique agar dapat dijadikan kolom FK
ALTER TABLE lembaga
    ADD CONSTRAINT lembaga_kode_lembaga_key UNIQUE (kode_lembaga);

-- Tambah FK dari tabel opd ke lembaga menggunakan kolom kode_lembaga
ALTER TABLE opd
    ADD CONSTRAINT opd_kode_lembaga_fkey
        FOREIGN KEY (kode_lembaga)
        REFERENCES lembaga (kode_lembaga)
        ON UPDATE CASCADE
        ON DELETE SET NULL;
