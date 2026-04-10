package domain

import "time"

type Musrenbang struct {
	Id               int
	Usulan           string
	Alamat           string
	Uraian           string
	Tahun            string
	KodeOpd          string
	NamaOpd          string
	Status           string
	CreatedDate      time.Time
	LastModifiedDate time.Time
}
