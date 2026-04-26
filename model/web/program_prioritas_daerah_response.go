package web

type ProgramPrioritasDaerahResponse struct {
	Id                         int    `json:"id,omitempty"`
	KodeProgramPrioritasDaerah string `json:"kode_program_prioritas_daerah"`
	NamaProgramPrioritasDaerah string `json:"nama_program_prioritas_daerah"`
	RencanaImplementasi        string `json:"rencana_implementasi"`
	Keterangan                 string `json:"keterangan"`
	TahunAwal                  string `json:"tahun_awal"`
	TahunAkhir                 string `json:"tahun_akhir"`
	IsActive                   bool   `json:"is_active"`
}
