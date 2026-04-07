DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.columns
        WHERE table_name = 'pegawai' AND column_name = 'jenis_pegawai'
    ) THEN
        ALTER TABLE pegawai ADD COLUMN jenis_pegawai VARCHAR(255);
    END IF;
END$$;
