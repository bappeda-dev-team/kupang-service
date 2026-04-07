package domain

import "database/sql"

type Jabatan struct {
	Id          int
	NamaJabatan string
	Tahun       sql.NullString
}
