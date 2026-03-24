package web

type TargetPokinOpdOperationalResponse struct {
	Id                             int    `json:"id,omitempty" swaggerignore:"true"`
	IndikatorPokinOpdOperationalId int    `json:"indikator_pokin_opd_operational_id"`
	NilaiTarget                    int    `json:"target"`
	Satuan                         string `json:"satuan"`
}
