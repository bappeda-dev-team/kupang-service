CREATE TABLE IF NOT EXISTS indikator_pokin_opd (
    id BIGSERIAL PRIMARY KEY,
    nama_indikator VARCHAR(255),
    created_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    last_modified_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);
