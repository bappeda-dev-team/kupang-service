package web

type BidangUrusanCreateRequest struct {
	KodeBidangUrusan string `json:"kode_bidang_urusan" validate:"required"`
	NamaBidangUrusan string `json:"nama_bidang_urusan" validate:"required"`
}
