package web

type JabatanUpdateRequest struct {
	Id          int    `json:"id" validate:"required" example:"0"`
	NamaJabatan string `json:"nama_jabatan" validate:"required" example:"string"`
}
