ALTER TABLE indikator_pokin_opd_strategic
    ADD COLUMN pokin_opd_strategic_id BIGINT;

ALTER TABLE indikator_pokin_opd_strategic
    ADD CONSTRAINT fk_indikator_pokin_opd_strategic_pokin_opd_strategic
        FOREIGN KEY (pokin_opd_strategic_id) REFERENCES pokin_opd_strategic (id) ON DELETE CASCADE;
