package rsp

type FormulaListRsp struct {
	Total int64     `json:"total"`
	List  []Formula `json:"list"`
}

type Formula struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Content      string `json:"content"`
}
type FormulaDetail struct {
	Name       string       `json:"name"`
	Proportion []Proportion `json:"proportion"`
	Content    string       `json:"content"`
}

type Proportion struct {
	HerbName string `json:"herbName"`
	Weight   int    `json:"weight"`
}

type HerbListRsp struct {
	Total int64  `json:"total"`
	List  []Herb `json:"list"`
}

type Herb struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Nature       string `json:"nature"`
	Brief        string `json:"brief"`
	UserId       int    `json:"user_id"`
}
