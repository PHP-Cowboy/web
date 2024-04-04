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
	Picture      string `json:"picture"`      //图片地址
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

type ToolRsp struct {
	Total int64  `json:"total"`
	List  []Tool `json:"list"`
}

type Tool struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Picture string `json:"picture"`
	Router  string `json:"router"`
}

type MindMapRsp struct {
	Total int64     `json:"total"`
	List  []MindMap `json:"list"`
}

type MindMap struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Picture string `json:"picture"`
}

type DiseaseRsp struct {
	Total int64     `json:"total"`
	List  []Disease `json:"list"`
}

type Disease struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ClinicalRsp struct {
	Total int64      `json:"total"`
	List  []Clinical `json:"list"`
}

type Clinical struct {
	Id      int    `json:"id"`
	Symptom string `json:"symptom"`
}

type BigDataPieChartRsp struct {
	Category Category  `json:"category"`
	List     []BigData `json:"big_data"`
}

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type BigData struct {
	Name string `json:"name"`
	Num  int    `json:"num"`
}

type Group struct {
	Id       int     `json:"id"`
	ParentId int     `json:"parent_id"`
	Name     string  `json:"name"`
	Children []Group `json:"children"`
}
