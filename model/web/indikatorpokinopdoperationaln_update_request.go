package web

type IndikatorPokinOpdOperationalNUpdateRequest struct {
	Id                     int    `json:"id" validate:"required"`
	PokinOpdOperationalNId int    `json:"pokin_opd_operationalN_id" validate:"required"`
	NamaIndikator          string `json:"indikator"`
}
