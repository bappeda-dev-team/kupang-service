ALTER TABLE indikator_pokin_opd_operational
    ADD COLUMN pokin_opd_operational_id BIGINT;

ALTER TABLE indikator_pokin_opd_operational
    ADD CONSTRAINT fk_indikator_pokin_opd_operational_pokin_opd_operational
        FOREIGN KEY (pokin_opd_operational_id) REFERENCES pokin_opd_operational (id) ON DELETE CASCADE;
