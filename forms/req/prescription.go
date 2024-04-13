package req

type PrescriptionRandList struct {
	Paging
}

type PrescriptionList struct {
	Paging
	KeyWords string `json:"key_words"`
}

type PrescriptionCelebrityList struct {
	Paging
	KeyWords       string `json:"key_words"`
	PrescriptionId int    `json:"prescription_id"`
}

type PrescriptionGraph struct {
	PrescriptionId int `json:"prescription_id"`
	Type           int `json:"type"`
}
