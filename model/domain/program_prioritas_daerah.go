package domain

import "time"

type ProgramPrioritasDaerah struct {
	Id                         int
	KodeProgramPrioritasDaerah string
	NamaProgramPrioritasDaerah string
	RencanaImplementasi        string
	Keterangan                 string
	TahunAwal                  string
	TahunAkhir                 string
	IsActive                   string
	CreatedDate                time.Time
	LastModifiedDate           time.Time
}
