CREATE TABLE IF NOT EXISTS pokin_opd_tactical (
    id BIGSERIAL PRIMARY KEY,
    parent INTEGER,
    nama_pohon VARCHAR(255),
    jenis_pohon VARCHAR(255),
    level_pohon INTEGER,
    kode_opd VARCHAR(255),
    nama_opd VARCHAR(255),
    keterangan VARCHAR(255),
    tahun INTEGER,
    jumlah_review INTEGER,
    status VARCHAR(255),
    pelaksana VARCHAR(255),
    updated_by VARCHAR(255),
    created_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    last_modified_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);
