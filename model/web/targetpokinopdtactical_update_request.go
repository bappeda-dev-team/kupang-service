package web

type TargetPokinOpdTacticalUpdateRequest struct {
	Id                          int    `json:"id" validate:"required"`
	IndikatorPokinOpdTacticalId int    `json:"indikator_pokin_opd_tactical_id" validate:"required"`
	NilaiTarget                 int    `json:"target" validate:"required"`
	Satuan                      string `json:"satuan" validate:"required"`
}
