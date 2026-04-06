package web

type PegawaiResponse struct {
	Id      int    `json:"id,omitempty"`
	Nama    string `json:"nama"`
	Nip     string `json:"nip"`
	Jabatan string `json:"jabatan"`
	KodeOpd string `json:"kode_opd"`
	NamaOpd string `json:"nama_opd"`
}
