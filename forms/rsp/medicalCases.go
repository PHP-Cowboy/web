package rsp

type MedicalCasesRsp struct {
	Total int64          `json:"total"`
	List  []MedicalCases `json:"list"`
}

type MedicalCases struct {
	Id           int    `json:"id"`
	Symptom      string `json:"symptom"`      //症状
	Prescription string `json:"prescription"` //方药
	Content      string `json:"content"`      //内容
	Provenance   string `json:"provenance"`   //出处
}
