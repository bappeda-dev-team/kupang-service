package web

type PokinOpdOperationalNResponse struct {
	Id           int                                      `json:"id,omitempty"`
	Parent       int                                      `json:"parent"`
	NamaPohon    string                                   `json:"nama_pohon"`
	JenisPohon   string                                   `json:"jenis_pohon"`
	LevelPohon   int                                      `json:"level_pohon"`
	KodeOpd      string                                   `json:"kode_opd"`
	NamaOpd      string                                   `json:"nama_opd"`
	Keterangan   string                                   `json:"keterangan"`
	Tahun        int                                      `json:"tahun"`
	JumlahReview int                                      `json:"jumlah_review"`
	Status       string                                   `json:"status"`
	Pelaksana    string                                   `json:"pelaksana"`
	UpdatedBy    string                                   `json:"updated_by"`
	Indikator    []PokinOpdOperationalNIndikatorResponse   `json:"indikator"`
}

type PokinOpdOperationalNIndikatorResponse struct {
	IdIndikator   int                                  `json:"id_indikator"`
	NamaIndikator string                               `json:"nama_indikator"`
	Targets       []PokinOpdOperationalNTargetResponse  `json:"targets"`
}

type PokinOpdOperationalNTargetResponse struct {
	IdTarget    int    `json:"id_target"`
	IndikatorId int    `json:"indikator_id"`
	Target      int    `json:"target"`
	Satuan      string `json:"satuan"`
}
