package web

type PegawaiAddJabatanRequest struct {
	Id          int    `json:"id" validate:"required"`
	NamaJabatan string `json:"nama_jabatan" validate:"required"`
}
