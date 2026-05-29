package web

type BidangUrusanResponse struct {
	Id               int    `json:"id,omitempty"`
	KodeBidangUrusan string `json:"kode_bidang_urusan"`
	NamaBidangUrusan string `json:"nama_bidang_urusan"`
	Tahun            string `json:"tahun"`
	CreatedDate      string `json:"created_date"`
	LastModifiedDate string `json:"last_modified_date"`
}
