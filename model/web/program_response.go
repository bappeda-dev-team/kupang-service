package web

type ProgramResponse struct {
	Id          int    `json:"id,omitempty"`
	KodeProgram string `json:"kode_program"`
	NamaProgram string `json:"nama_program"`
	Tahun       string `json:"tahun"`
}
