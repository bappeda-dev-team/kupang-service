CREATE TABLE IF NOT EXISTS "subkegiatan" (
    id BIGSERIAL PRIMARY KEY,
    kode_subkegiatan VARCHAR(255) NOT NULL,
    nama_subkegiatan VARCHAR(255) NOT NULL,
    tahun VARCHAR(255) NOT NULL,
    created_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    last_modified_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);
