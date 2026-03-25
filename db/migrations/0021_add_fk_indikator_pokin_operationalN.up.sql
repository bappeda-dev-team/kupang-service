ALTER TABLE indikator_pokin_opd_operationalN
    ADD COLUMN pokin_opd_operationalN_id BIGINT;

ALTER TABLE indikator_pokin_opd_operationalN
    ADD CONSTRAINT fk_indikator_pokin_opd_operationalN_pokin_opd_operationalN
        FOREIGN KEY (pokin_opd_operationalN_id) REFERENCES pokin_opd_operationalN (id) ON DELETE CASCADE;
