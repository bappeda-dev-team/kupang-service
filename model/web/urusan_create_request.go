package web

type UrusanCreateRequest struct {
	KodeUrusan string `json:"kode_urusan" validate:"required"`
	NamaUrusan string `json:"nama_urusan" validate:"required"`
	Tahun      string `json:"tahun" validate:"required"`
}
