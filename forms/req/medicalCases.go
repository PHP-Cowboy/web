package req

type MedicalCasesList struct {
	Paging
	KeyWords string `json:"key_words" form:"key_words"`
}

type ClassicsCategoryList struct {
	Paging
	Name string `json:"name"`
}

type BookListByCategory struct {
	Paging
	CategoryId int `json:"category_id"`
}

type CatalogueList struct {
	Paging
	ClassicsId int `json:"classics_id"`
}

type PrevNext struct {
	Id         int `json:"id"`
	ClassicsId int `json:"classics_id"`
}

type ToolList struct {
	Paging
}

type MindMapList struct {
	Paging
	KeyWords string `json:"key_words"`
}

type DiseaseList struct {
	CategoryId int `json:"category_id" binding:"required"`
}

type MindMethodList struct {
	Paging
	KeyWords string `json:"key_words"`
}

type CommonlyPrescriptionCategoryList struct {
	Paging
	KeyWords string `json:"key_words"`
}

type CommonlyPrescriptionList struct {
	Paging
	KeyWords   string `json:"key_words"`
	CategoryId int    `json:"category_id"`
}

type CompleteCollectionPrescriptionList struct {
	Paging
	KeyWords   string `json:"key_words"`
	CategoryId int    `json:"category_id"`
}
