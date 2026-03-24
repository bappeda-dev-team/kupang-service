package web

type IndikatorPokinOpdOperationalUpdateRequest struct {
	Id                     int    `json:"id" validate:"required"`
	PokinOpdOperationalId  int    `json:"pokin_opd_operational_id" validate:"required"`
	NamaIndikator          string `json:"indikator"`
}
