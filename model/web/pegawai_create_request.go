package web

type PegawaiCreateRequest struct {
	Nama    string `json:"nama" validate:"required"`
	Nip     string `json:"nip" validate:"required"`
	Jabatan string `json:"jabatan" validate:"required"`
	KodeOpd string `json:"kode_opd" validate:"required"`
	NamaOpd string `json:"nama_opd" validate:"required"`
}
