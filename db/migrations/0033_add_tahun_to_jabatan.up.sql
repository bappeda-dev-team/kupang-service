DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM information_schema.columns
        WHERE table_name = 'jabatan'
          AND column_name = 'tahun'
    ) THEN
        ALTER TABLE jabatan ADD COLUMN tahun VARCHAR(255);
    END IF;
END$$;
