package web

type TargetPokinOpdCreateRequest struct {
	IndikatorPokinOpdId int    `json:"indikator_pokin_opd_id"`
	NilaiTarget         int    `json:"target"`
	Satuan              string `json:"satuan"`
}
