package web

type UrusanUpdateRequest struct {
	Id         int    `json:"id"`
	KodeUrusan string `json:"kode_urusan" validate:"required"`
	NamaUrusan string `json:"nama_urusan" validate:"required"`
}
