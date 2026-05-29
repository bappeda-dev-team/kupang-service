package web

type ProgramUpdateRequest struct {
	Id          int    `json:"id"`
	KodeProgram string `json:"kode_program" validate:"required"`
	NamaProgram string `json:"nama_program" validate:"required"`
	Tahun       string `json:"tahun" validate:"required"`
}
