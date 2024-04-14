package rsp

type TyphoidCatalogueRsp struct {
	Total int64              `json:"total"`
	List  []TyphoidCatalogue `json:"list"`
}

type TyphoidCatalogue struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Brief   string `json:"brief"`
	Content string `json:"content"`
}
