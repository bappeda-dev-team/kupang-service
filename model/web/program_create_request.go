package web

type ProgramCreateRequest struct {
	KodeProgram string `json:"kode_program" validate:"required"`
	NamaProgram string `json:"nama_program" validate:"required"`
	Tahun       string `json:"tahun" validate:"required"`
}
