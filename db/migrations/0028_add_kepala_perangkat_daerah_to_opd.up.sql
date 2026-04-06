ALTER TABLE opd
    ADD COLUMN nama_kepala_perangkat_daerah VARCHAR(255) NOT NULL DEFAULT '',
    ADD COLUMN nip_kepala_perangkat_daerah VARCHAR(255) NOT NULL DEFAULT '',
    ADD COLUMN pangkat_kepala_perangkat_daerah VARCHAR(255) NOT NULL DEFAULT '';
