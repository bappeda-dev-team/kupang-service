CREATE TABLE IF NOT EXISTS target_pokin_opd_operationalN (
    id BIGSERIAL PRIMARY KEY,
    indikator_pokin_opd_operationalN_id BIGINT,
    nilai_target INTEGER,
    satuan VARCHAR(50),
    created_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    last_modified_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    CONSTRAINT fk_target_pokin_opd_operationalN_indikator
        FOREIGN KEY (indikator_pokin_opd_operationalN_id) REFERENCES indikator_pokin_opd_operationalN (id) ON DELETE CASCADE
);
