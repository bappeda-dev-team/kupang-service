ALTER TABLE lembaga
    ADD COLUMN jabatan_kepala_lembaga VARCHAR(255) NOT NULL DEFAULT '',
    ADD COLUMN nama_kepala_lembaga VARCHAR(255) NOT NULL DEFAULT '',
    ADD COLUMN nip_kepala_lembaga VARCHAR(255) NOT NULL DEFAULT '';