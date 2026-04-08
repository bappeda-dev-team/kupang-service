package web

type PeriodeUpdateRequest struct {
	Id           int    `json:"id" validate:"required"`
	TahunAwal    string `json:"tahun_awal" validate:"required"`
	TahunAkhir   string `json:"tahun_akhir" validate:"required"`
	JenisPeriode string `json:"jenis_periode" validate:"required"`
}
