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

type ClassicsCategoryRsp struct {
	Total int64              `json:"total"`
	List  []ClassicsCategory `json:"list"`
}

type ClassicsCategory struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type BookListRsp struct {
	Total int64      `json:"total"`
	List  []BookList `json:"list"`
}

type BookList struct {
	Id         int    `json:"id"`
	CategoryId int    `json:"category_id"`
	Name       string `json:"name"`
	Author     string `json:"author"`
	Dynasty    string `json:"dynasty"`
}

type CatalogueRsp struct {
	Total int64       `json:"total"`
	List  []Catalogue `json:"list"`
}

type Catalogue struct {
	Id         int    `json:"id"`
	ClassicsId int    `json:"classics_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}
