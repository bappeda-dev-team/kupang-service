package web

type UserCreateRequest struct {
	Nama        string `json:"nama" validate:"required"`
	Nip         string `json:"nip" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Status      string `json:"status" validate:"required"`
	Role        string `json:"role" validate:"required"`
	KodeOpd     string `json:"kode_opd" validate:"required"`
	OpdId       int    `json:"opd_id" validate:"required"`
	PegawaiId   int    `json:"pegawai_id" validate:"required"`
	RoleId      int    `json:"role_id" validate:"required"`
	NamaOpd     string `json:"nama_opd,omitempty"`
	NamaPegawai string `json:"nama_pegawai,omitempty"`
}
