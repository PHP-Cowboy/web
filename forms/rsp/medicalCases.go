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

type MindMethodRsp struct {
	Total int64        `json:"total"`
	List  []MindMethod `json:"list"`
}

type MindMethod struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Doctor  string `json:"doctor"`
	Content string `json:"content"`
}

type CommonlyPrescriptionCategoryRsp struct {
	Total int64                          `json:"total"`
	List  []CommonlyPrescriptionCategory `json:"list"`
}

type CommonlyPrescriptionCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CommonlyPrescriptionRsp struct {
	Total int64                  `json:"total"`
	List  []CommonlyPrescription `json:"list"`
}

type CommonlyPrescription struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Provenance  string `json:"provenance"`
	Constituent string `json:"constituent"` //组成成分
	Usage       string `json:"usage"`       //用法
	Efficacy    string `json:"efficacy"`    //功效
	Mainly      string `json:"mainly"`      //主治
	Song        string `json:"song"`        //方歌
	CategoryId  int    `json:"category_id"`
}

type CompleteCollectionPrescriptionRsp struct {
	Total int64                            `json:"total"`
	List  []CompleteCollectionPrescription `json:"list"`
}

type CompleteCollectionPrescription struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Provenance string `json:"provenance"`
}

type Group struct {
	Id       int     `json:"id"`
	ParentId int     `json:"parent_id"`
	Name     string  `json:"name"`
	Children []Group `json:"children"`
}

type QuestionCategoryRsp struct {
	Total int64              `json:"total"`
	List  []QuestionCategory `json:"list"`
}

type QuestionCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Question struct {
	Id         int    `json:"id"`
	CategoryId int    `json:"category_id"`
	Number     int    `json:"number"`
	Topic      string `json:"topic"`
	A          string `json:"a"`
	B          string `json:"b"`
	C          string `json:"c"`
	D          string `json:"d"`
	E          string `json:"e"`
	Answer     string `json:"answer"`
	Analysis   string `json:"analysis"`
}
