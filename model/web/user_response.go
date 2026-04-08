package web

type UserResponse struct {
	Id          int     `json:"id,omitempty"`
	Nama        string  `json:"nama"`
	Nip         *string `json:"nip,omitempty"`
	Email       string  `json:"email"`
	Status      string  `json:"status"`
	Role        *string `json:"role,omitempty"`
	OpdId       *int    `json:"opd_id,omitempty"`
	PegawaiId   *int    `json:"pegawai_id,omitempty"`
	RoleId      *int    `json:"role_id,omitempty"`
	NamaOpd     *string `json:"nama_opd,omitempty"`
	NamaPegawai *string `json:"nama_pegawai,omitempty"`
	KodeOpd     *string `json:"kode_opd,omitempty"`
}
