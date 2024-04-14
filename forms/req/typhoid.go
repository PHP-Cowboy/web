package req

type TyphoidCatalogueList struct {
	Paging
	KeyWords string `json:"key_words"`
}
