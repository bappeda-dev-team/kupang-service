package web

type IndikatorPokinOpdOperationalCreateRequest struct {
	PokinOpdOperationalId int    `json:"pokin_opd_operational_id" validate:"required"`
	NamaIndikator         string `json:"indikator" validate:"required"`
}
