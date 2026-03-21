ALTER TABLE tujuan_pokin_opd
    ADD COLUMN pokin_opd_id BIGINT;

UPDATE tujuan_pokin_opd t
SET pokin_opd_id = p.id
FROM pokin_opd p
WHERE p.kode_opd = t.kode_opd
  AND p.tahun = t.tahun_awal_periode
  AND t.pokin_opd_id IS NULL;

ALTER TABLE tujuan_pokin_opd
    ADD CONSTRAINT fk_tujuan_pokin_opd_pokin_opd
        FOREIGN KEY (pokin_opd_id) REFERENCES pokin_opd (id);

ALTER TABLE indikator_pokin_opd
    ADD COLUMN tujuan_pokin_opd_id BIGINT,
    ADD CONSTRAINT fk_indikator_pokin_opd_tujuan_pokin_opd
        FOREIGN KEY (tujuan_pokin_opd_id) REFERENCES tujuan_pokin_opd (id);

ALTER TABLE target_pokin_opd
    ADD COLUMN indikator_pokin_opd_id BIGINT,
    ADD CONSTRAINT fk_target_pokin_opd_indikator_pokin_opd
        FOREIGN KEY (indikator_pokin_opd_id) REFERENCES indikator_pokin_opd (id);
