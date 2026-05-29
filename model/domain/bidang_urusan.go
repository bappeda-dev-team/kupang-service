package domain

import "time"

type BidangUrusan struct {
	Id                int
	KodeBidangUrusan  string
	NamaBidangUrusan  string
	Tahun             string
	CreatedDate       time.Time
	LastModifiedDate  time.Time
}
