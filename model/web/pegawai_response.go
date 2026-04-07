package web

type PegawaiResponse struct {
	Id           int     `json:"id,omitempty" example:"0"`
	Nama         string  `json:"nama" example:"string"`
	Nip          string  `json:"nip" example:"string"`
	JabatanId    *int    `json:"jabatan_id,omitempty" example:"0"`
	NamaJabatan  *string `json:"nama_jabatan,omitempty" example:"string"`
	TahunJabatan *string `json:"tahun_jabatan,omitempty" example:"2024"`
	KodeOpd      string  `json:"kode_opd" example:"string"`
	NamaOpd      string  `json:"nama_opd" example:"string"`
	JenisPegawai *string `json:"jenis_pegawai,omitempty" example:"string"`
}
