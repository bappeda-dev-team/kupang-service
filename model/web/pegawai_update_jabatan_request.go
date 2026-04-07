package web

type PegawaiUpdateJabatanRequest struct {
	PegawaiId   int    `json:"pegawai_id" validate:"required" example:"0"`
	JabatanId   int    `json:"jabatan_id" validate:"required" example:"0"`
	NamaJabatan string `json:"nama_jabatan" validate:"required" example:"string"`
}
