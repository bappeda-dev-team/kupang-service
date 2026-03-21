CREATE TABLE IF NOT EXISTS tujuan_pokin_opd (
    id BIGSERIAL PRIMARY KEY,
    kode_opd VARCHAR(255),
    nama_tujuan VARCHAR(255),
    bidang_urusan VARCHAR(255),
    tahun_awal_periode INTEGER,
    tahun_akhir_periode INTEGER,
    created_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    last_modified_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);
