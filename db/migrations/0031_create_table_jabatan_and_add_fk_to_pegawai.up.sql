CREATE TABLE IF NOT EXISTS jabatan (
    id BIGSERIAL PRIMARY KEY,
    nama_jabatan VARCHAR(255) NOT NULL UNIQUE,
    created_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    last_modified_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.columns
        WHERE table_name = 'pegawai' AND column_name = 'jabatan_id'
    ) THEN
        ALTER TABLE pegawai ADD COLUMN jabatan_id BIGINT;
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.columns
        WHERE table_name = 'pegawai' AND column_name = 'nama_jabatan'
    ) THEN
        ALTER TABLE pegawai ADD COLUMN nama_jabatan VARCHAR(255);
    END IF;
END$$;

INSERT INTO jabatan (nama_jabatan)
SELECT DISTINCT jabatan FROM pegawai
WHERE jabatan IS NOT NULL
  AND jabatan <> ''
  AND NOT EXISTS (
      SELECT 1 FROM jabatan j WHERE j.nama_jabatan = pegawai.jabatan
  );

UPDATE pegawai p
SET jabatan_id = j.id,
    nama_jabatan = j.nama_jabatan
FROM jabatan j
WHERE p.jabatan = j.nama_jabatan;

ALTER TABLE pegawai
    ADD CONSTRAINT pegawai_jabatan_id_fkey
        FOREIGN KEY (jabatan_id)
        REFERENCES jabatan (id)
        ON UPDATE CASCADE
        ON DELETE SET NULL;

CREATE OR REPLACE FUNCTION set_pegawai_nama_jabatan_null()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE pegawai
    SET nama_jabatan = NULL
    WHERE jabatan_id = OLD.id;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_set_pegawai_nama_jabatan_null ON jabatan;
CREATE TRIGGER trg_set_pegawai_nama_jabatan_null
BEFORE DELETE ON jabatan
FOR EACH ROW
EXECUTE FUNCTION set_pegawai_nama_jabatan_null();

DO $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM information_schema.columns
        WHERE table_name = 'pegawai' AND column_name = 'jabatan'
    ) THEN
        ALTER TABLE pegawai DROP COLUMN jabatan;
    END IF;
END$$;
