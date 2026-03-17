ALTER TABLE indikator_pokin_opd_tactical
    ADD COLUMN pokin_opd_tactical_id BIGINT;

ALTER TABLE indikator_pokin_opd_tactical
    ADD CONSTRAINT fk_indikator_pokin_opd_tactical_pokin_opd_tactical
        FOREIGN KEY (pokin_opd_tactical_id) REFERENCES pokin_opd_tactical (id) ON DELETE CASCADE;
