package domain

import "database/sql"

type User struct {
	Id          int
	Nama        string
	Nip         sql.NullString
	Email       string
	Status      string
	Role        sql.NullString
	KodeOpd     sql.NullString
	OpdId       sql.NullInt64
	PegawaiId   sql.NullInt64
	RoleId      sql.NullInt64
	NamaOpd     sql.NullString
	NamaPegawai sql.NullString
}
