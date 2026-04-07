package domain

import "database/sql"

type Pegawai struct {
	Id           int
	Nama         string
	Nip          string
	JabatanId    sql.NullInt64
	NamaJabatan  sql.NullString
	TahunJabatan sql.NullString
	KodeOpd      string
	NamaOpd      string
	JenisPegawai sql.NullString
}
