package rsp

type PrescriptionRsp struct {
	Total int64          `json:"total"`
	List  []Prescription `json:"list"`
}

type Prescription struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Symptom       string `json:"symptom"`
	TongueQuality string `json:"tongue_quality"`
	CoatedTongue  string `json:"coated_tongue"`
	Pulse         string `json:"pulse"`
	ModernDisease string `json:"modern_disease"`
	SymptomPic    string `json:"symptom_pic"`
	PulsePic      string `json:"pulse_pic"`
	TonguePic     string `json:"tongue_pic"`
}

type CelebrityRsp struct {
	Total int64       `json:"total"`
	List  []Celebrity `json:"list"`
}

type Celebrity struct {
	Id             int    `json:"id"`
	NotableDoctor  string `json:"notable_doctor"`
	Provenance     string `json:"provenance"`
	Content        string `json:"content"`
	PrescriptionId int    `json:"prescription_id"`
}

type PrescriptionGraph struct {
	Id             int    `json:"id"`
	PrescriptionId int    `json:"prescription_id"`
	Type           int    `json:"type"`
	Name           string `json:"name"`
	Num            int    `json:"num"`
}
