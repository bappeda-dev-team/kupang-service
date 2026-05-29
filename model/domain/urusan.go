package domain

import "time"

type Urusan struct {
	Id               int
	KodeUrusan       string
	NamaUrusan       string
	Tahun            string
	CreatedDate      time.Time
	LastModifiedDate time.Time
}
