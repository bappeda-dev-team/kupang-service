package web

type BidangUrusanResponse struct {
	Id         int    `json:"id,omitempty"`
	KodeUrusan string `json:"kode_urusan"`
	NamaUrusan string `json:"nama_urusan"`
}
