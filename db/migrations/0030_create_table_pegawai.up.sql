DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE table_name = 'opd' AND constraint_name = 'opd_kode_opd_key'
    ) THEN
        ALTER TABLE opd ADD CONSTRAINT opd_kode_opd_key UNIQUE (kode_opd);
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS pegawai (
    id BIGSERIAL PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    nip VARCHAR(255) NOT NULL,
    jabatan VARCHAR(255) NOT NULL,
    kode_opd VARCHAR(255) NOT NULL,
    nama_opd VARCHAR(255) NOT NULL,
    created_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    last_modified_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    CONSTRAINT pegawai_nip_key UNIQUE (nip)
);

ALTER TABLE pegawai
    ADD CONSTRAINT pegawai_kode_opd_fkey
        FOREIGN KEY (kode_opd)
        REFERENCES opd (kode_opd)
        ON UPDATE CASCADE
        ON DELETE RESTRICT;
