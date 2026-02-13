package web

type TujuanPokinOpdResponse struct {
	Id                int    `json:"id,omitempty"`
	KodeOpd           string `json:"kode_opd"`
	NamaTujuan        string `json:"tujuan"`
	BidangUrusan      string `json:"bidang_urusan"`
	TahunAwalPeriode  int    `json:"tahun_awal_periode"`
	TahunAkhirPeriode int    `json:"tahun_akhir_periode"`
}
