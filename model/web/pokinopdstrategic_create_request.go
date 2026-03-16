package web

type PokinOpdStrategicCreateRequest struct {
	Parent       int    `json:"parent"`
	NamaPohon    string `json:"nama_pohon"`
	JenisPohon   string `json:"jenis_pohon"`
	LevelPohon   int    `json:"level_pohon"`
	KodeOpd      string `json:"kode_opd"`
	NamaOpd      string `json:"nama_opd"`
	Keterangan   string `json:"keterangan"`
	Tahun        int    `json:"tahun"`
	JumlahReview int    `json:"jumlah_review"`
	Status       string `json:"status"`
	Pelaksana    string `json:"pelaksana"`
	UpdatedBy    string `json:"updated_by"`
}
