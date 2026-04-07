package web

type PegawaiCreateRequest struct {
	Nama      string `json:"nama" validate:"required" example:"string"`
	Nip       string `json:"nip" validate:"required" example:"string"`
	JabatanId *int   `json:"jabatan_id" validate:"omitempty" example:"0"`
	KodeOpd   string `json:"kode_opd" validate:"required" example:"string"`
	NamaOpd   string `json:"nama_opd" validate:"required" example:"string"`
}
