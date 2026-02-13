CREATE TABLE IF NOT EXISTS target_pokin_opd (
    id BIGSERIAL PRIMARY KEY,
    nilai_target INTEGER,
    satuan VARCHAR(50),
    created_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    last_modified_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);
