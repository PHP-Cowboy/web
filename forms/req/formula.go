package req

type FormulaList struct {
	Paging
	Name   string `json:"name" form:"name"`
	IsMy   bool
	UserId int
}

type SaveMyFormula struct {
	Name           string       `json:"name"`
	Abbreviation   string       `json:"abbreviation"`
	ProportionList []Proportion `json:"proportionList"`
	Content        string       `json:"content"`
	UserId         int
}

type Proportion struct {
	HerbId int `json:"herbId" binding:"required"`
	Weight int `json:"weight" binding:"required"`
}

type HerbList struct {
	Paging
	Name   string `json:"name" form:"name"`
	UserId int
}

type SaveHerb struct {
	Name         string `json:"name" form:"name"`
	Abbreviation string `json:"abbreviation" form:"abbreviation"`
	Nature       string `json:"nature" form:"nature"`
	Brief        string `json:"brief" form:"brief"`
	UserId       int
}
