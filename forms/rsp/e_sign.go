package rsp

type GetTemplateListRsp struct {
	Total int            `json:"total"`
	List  []TemplateList `json:"list"`
}

type TemplateList struct {
	FlowTemplateId   string `json:"flowTemplateId"`
	FlowTemplateName string `json:"flowTemplateName"`
	DocTemplateId    string `json:"docTemplateId"`
	DocTemplateName  string `json:"docTemplateName"`
}

type ESignTemplatesRsp struct {
	Code    int       `json:"code"`
	Data    Templates `json:"data"`
	Message string    `json:"message"`
}

type Templates struct {
	FlowTemplateBasicInfos []FlowTemplateBasicInfos `json:"flowTemplateBasicInfos"`
	Total                  int                      `json:"total"`
}

type FlowTemplateBasicInfos struct {
	FlowTemplateId   string         `json:"flowTemplateId"`
	FlowTemplateName string         `json:"flowTemplateName"`
	DocTemplates     []DocTemplates `json:"docTemplates"`
}

type DocTemplates struct {
	DocTemplateId   string `json:"docTemplateId"`
	DocTemplateName string `json:"docTemplateName"`
}

type ESignDocTemplatesDetailRsp struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    DocTemplatesDetail `json:"data"`
}

type DocTemplatesDetail struct {
	TemplateId       string       `json:"templateId"`
	TemplateName     string       `json:"templateName"`
	FileKey          string       `json:"fileKey"`
	DownloadUrl      string       `json:"downloadUrl"`
	StructComponents []Components `json:"structComponents"`
}

type Components struct {
	Id      string  `json:"id"`
	Key     string  `json:"key"`
	Context Context `json:"context"`
	RefId   string  `json:"refId"`
}

type Context struct {
	Label string `json:"label"`
	Pos   Pos    `json:"pos"`
}

type Pos struct {
	X    float32 `json:"x"`
	Y    float32 `json:"y"`
	Page int     `json:"page"`
}

type ESignCreateByTemplateRsp struct {
	Code    int                       `json:"code"`
	Data    ESignCreateByTemplateData `json:"data"`
	Message string                    `json:"message"`
}

type ESignCreateByTemplateData struct {
	ContractId  int    `json:"contractId"`
	DownloadUrl string `json:"downloadUrl"`
	FileId      string `json:"fileId"`
	FileName    string `json:"fileName"`
}

type SigningAddressRsp struct {
	Code    int     `json:"code"`
	Data    Address `json:"data"`
	Message string  `json:"message"`
}

type Address struct {
	ShortUrl string `json:"shortUrl"`
	Url      string `json:"url"`
}

type GetDownloadUrlRsp struct {
	Code    int         `json:"code"`
	Data    DownloadUrl `json:"data"`
	Message string      `json:"message"`
}

type DownloadUrl struct {
	DownloadUrl   string `json:"downloadUrl"`
	PdfTotalPages int    `json:"pdfTotalPages"`
}

type DocumentDownloadRsp struct {
	Code    int      `json:"code"`
	Data    Document `json:"data"`
	Message string   `json:"message"`
}

type Document struct {
	Docs []Docs `json:"docs"`
}

type Docs struct {
	FileUrl string `json:"fileUrl"`
}

type CreateUserFlowRsp struct {
	FlowId   string `json:"flow_id"`
	ShortUrl string `json:"short_url"`
}

type InformLetterCreateFlowReqRsp struct {
	FlowId   string `json:"flow_id"`
	ShortUrl string `json:"short_url"`
}

type GetPersonInfoRsp struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    PersonInfo `json:"data"`
}

type PersonInfo struct {
	Mobile             string `json:"mobile"`
	AccountId          string `json:"accountId"`
	Name               string `json:"name"`
	IdType             string `json:"idType"`
	IdNumber           string `json:"idNumber"`
	ThirdPartyUserType string `json:"thirdPartyUserType"`
}

type SearchWordsPositionRsp struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    []SearchWordsPosition `json:"data"`
}

type SearchWordsPosition struct {
	FileId       string         `json:"fileId"`
	Keyword      string         `json:"keyword"`
	PositionList []PositionList `json:"positionList"`
}

type PositionList struct {
	PageIndex      int              `json:"pageIndex"`
	CoordinateList []CoordinateList `json:"coordinateList"`
}

type CoordinateList struct {
	Posx float64 `json:"posx"`
	Posy float64 `json:"posy"`
}

type AddComponentsRsp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
