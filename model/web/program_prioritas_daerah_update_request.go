package web

type ProgramPrioritasDaerahUpdateRequest struct {
	Id                         int    `json:"id" validate:"required"`
	KodeProgramPrioritasDaerah string `json:"kode_program_prioritas_daerah" validate:"required"`
	NamaProgramPrioritasDaerah string `json:"nama_program_prioritas_daerah" validate:"required"`
	RencanaImplementasi        string `json:"rencana_implementasi" validate:"required"`
	Keterangan                 string `json:"keterangan" validate:"required"`
	TahunAwal                  string `json:"tahun_awal" validate:"required"`
	TahunAkhir                 string `json:"tahun_akhir" validate:"required"`
	IsActive                   bool   `json:"is_active" validate:"required"`
}
