package req

type FormulaList struct {
	Paging
	Name   string `json:"name" form:"name"`
	IsMy   bool
	UserId int
}

type SaveMyFormula struct {
	Name             string `json:"name"`
	NameAbbreviation string `json:"nameAbbreviation"`
	Proportion       string `json:"proportion"`
	Dose             string `json:"dose"`
	Content          string `json:"content"`
	UserId           int
}

type MedicineList struct {
	Paging
	Name   string `json:"name" form:"name"`
	UserId int
}

type SaveMedicine struct {
	Name   string `json:"name" form:"name"`
	UserId int
}
