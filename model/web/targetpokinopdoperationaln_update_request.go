package web

type TargetPokinOpdOperationalNUpdateRequest struct {
	Id                              int    `json:"id" validate:"required"`
	IndikatorPokinOpdOperationalNId int    `json:"indikator_pokin_opd_operationalN_id" validate:"required"`
	NilaiTarget                     int    `json:"target" validate:"required"`
	Satuan                          string `json:"satuan" validate:"required"`
}
