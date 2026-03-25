package web

type TargetPokinOpdOperationalNResponse struct {
	Id                              int    `json:"id,omitempty" swaggerignore:"true"`
	IndikatorPokinOpdOperationalNId int    `json:"indikator_pokin_opd_operationalN_id"`
	NilaiTarget                     int    `json:"target"`
	Satuan                          string `json:"satuan"`
}
