CREATE TABLE IF NOT EXISTS "user" (
    id BIGSERIAL PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    kode_opd VARCHAR(255),
    opd_id BIGINT,
    nama_opd VARCHAR(255),
    nip VARCHAR(255),
    pegawai_id BIGINT,
    nama_pegawai VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    role VARCHAR(255),
    role_id BIGINT,
    created_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    last_modified_date TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    CONSTRAINT user_opd_id_fkey FOREIGN KEY (opd_id) REFERENCES opd (id) ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT user_pegawai_id_fkey FOREIGN KEY (pegawai_id) REFERENCES pegawai (id) ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT user_role_id_fkey FOREIGN KEY (role_id) REFERENCES "role" (id) ON UPDATE CASCADE ON DELETE SET NULL
);
