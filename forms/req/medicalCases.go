package req

type MedicalCasesList struct {
	Paging
	KeyWords string `json:"key_words" form:"key_words"`
}
