package web

type PegawaiAddJabatanRequest struct {
	PegawaiId   int    `json:"pegawai_id" validate:"required"`
	NamaJabatan string `json:"nama_jabatan" validate:"required"`
}
