package web

type PohonKinerjaResponse struct {
	KodeOpd   string                          `json:"kode_opd"`
	NamaOpd   string                          `json:"nama_opd"`
	Tahun     int                             `json:"tahun"`
	TujuanOpd []PohonKinerjaTujuanResponse    `json:"tujuan_opd"`
	Childs    []PohonKinerjaStrategicResponse `json:"childs"`
}

type PohonKinerjaTujuanResponse struct {
	Id                int                             `json:"id"`
	KodeOpd           string                          `json:"kode_opd"`
	Tujuan            string                          `json:"tujuan"`
	BidangUrusan      string                          `json:"bidang_urusan"`
	TahunAwalPeriode  int                             `json:"tahun_awal_periode"`
	TahunAkhirPeriode int                             `json:"tahun_akhir_periode"`
	Indikator         []PohonKinerjaIndikatorResponse `json:"indikator"`
}

type PohonKinerjaIndikatorResponse struct {
	Id        int                          `json:"id"`
	Indikator string                       `json:"indikator"`
	Targets   []PohonKinerjaTargetResponse `json:"targets"`
}

type PohonKinerjaTargetResponse struct {
	Id     int    `json:"id"`
	Target int    `json:"target"`
	Satuan string `json:"satuan"`
}

type PohonKinerjaStrategicResponse struct {
	Id           int                                  `json:"id,omitempty"`
	Parent       int                                  `json:"parent"`
	NamaPohon    string                               `json:"nama_pohon"`
	JenisPohon   string                               `json:"jenis_pohon"`
	LevelPohon   int                                  `json:"level_pohon"`
	KodeOpd      string                               `json:"kode_opd"`
	NamaOpd      string                               `json:"nama_opd"`
	Keterangan   string                               `json:"keterangan"`
	Tahun        int                                  `json:"tahun"`
	JumlahReview int                                  `json:"jumlah_review"`
	Status       string                               `json:"status"`
	Pelaksana    string                               `json:"pelaksana"`
	UpdatedBy    string                               `json:"updated_by"`
	Indikator    []PokinOpdStrategicIndikatorResponse `json:"indikator"`
	Childs       []PokinOpdTacticalResponse           `json:"childs"`
}
