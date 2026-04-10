package domain

import (
	"database/sql"
	"time"
)

type PokokPikiran struct {
	Id               int
	Usulan           string
	Alamat           string
	Uraian           string
	Tahun            sql.NullString
	KodeOpd          string
	NamaOpd          string
	Status           sql.NullString
	CreatedDate      time.Time
	LastModifiedDate time.Time
}
