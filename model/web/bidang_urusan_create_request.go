package web

type BidangUrusanCreateRequest struct {
	KodeUrusan string `json:"kode_urusan" validate:"required"`
	NamaUrusan string `json:"nama_urusan" validate:"required"`
}
