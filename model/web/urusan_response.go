package web

type UrusanResponse struct {
	Id               int    `json:"id,omitempty"`
	KodeUrusan       string `json:"kode_urusan"`
	NamaUrusan       string `json:"nama_urusan"`
	Tahun            string `json:"tahun"`
	CreatedDate      string `json:"created_date"`
	LastModifiedDate string `json:"last_modified_date"`
}
