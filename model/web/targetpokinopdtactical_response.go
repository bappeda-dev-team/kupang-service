package web

type TargetPokinOpdTacticalResponse struct {
	Id                          int    `json:"id,omitempty" swaggerignore:"true"`
	IndikatorPokinOpdTacticalId int    `json:"indikator_pokin_opd_tactical_id"`
	NilaiTarget                 int    `json:"target"`
	Satuan                      string `json:"satuan"`
}
