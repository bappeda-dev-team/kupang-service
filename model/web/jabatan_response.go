package web

type JabatanResponse struct {
	Id          int     `json:"id,omitempty" example:"0"`
	NamaJabatan string  `json:"nama_jabatan" example:"string"`
	Tahun       *string `json:"tahun,omitempty" example:"2024"`
}
