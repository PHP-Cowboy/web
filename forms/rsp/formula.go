package rsp

type FormulaListRsp struct {
	Total int64     `json:"total"`
	List  []Formula `json:"list"`
}

type Formula struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	NameAbbreviation string `json:"nameAbbreviation"`
	Proportion       string `json:"proportion"`
	Dose             string `json:"dose"`
	Content          string `json:"content"`
}

type MedicineListRsp struct {
	Total int64      `json:"total"`
	List  []Medicine `json:"list"`
}

type Medicine struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
