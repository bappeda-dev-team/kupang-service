CREATE TABLE IF NOT EXISTS "program_prioritas_daerah" (
    id BIGSERIAL PRIMARY KEY,
    kode_program_prioritas_daerah VARCHAR(255) NOT NULL,
    nama_program_prioritas_daerah VARCHAR(255) NOT NULL,
    rencana_implementasi VARCHAR(255) NOT NULL,
    keterangan VARCHAR(255) NOT NULL,
    tahun_awal VARCHAR(255) NOT NULL,
    tahun_akhir VARCHAR(255) NOT NULL,
    is_active VARCHAR(255) NOT NULL,
    created_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    last_modified_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);
