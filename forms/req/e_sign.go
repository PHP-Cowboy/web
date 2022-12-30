package req

type GetTemplateReq struct {
	Paging
}

type GetDocTemplatesDetailReq struct {
	DocTemplateId string `json:"docTemplateId" form:"docTemplateId" binding:"required"`
}

type CreateByTemplateReq struct {
	TemplateId     string           `json:"templateId" binding:"required"`
	TemplateName   string           `json:"templateName" binding:"required"`
	FillComponents []FillComponents `json:"fill_components" binding:"required"`
}

type FillComponents struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

type CreateFlowOneStepReq struct {
	FileId     string `json:"file_id" binding:"required"`
	MerchantId int    `json:"merchant_id" binding:"required"`
}

type CreateUserFlowReq struct {
	MerchantId int `json:"merchant_id" binding:"required"`
	ContractId int `json:"contract_id" binding:"required"`
}

type InformLetterCreateFlowReq struct {
	MerchantId int `json:"merchant_id" binding:"required"`
	IlId       int `json:"contract_id" binding:"required"`
}

type SingleStoreContractReq struct {
	MerchantId              int      `json:"merchantId"`
	SalesmanId              int      `json:"salesmanId"`
	SalesmanName            string   `json:"salesmanName"`
	SalesmanTel             string   `json:"salesmanTel"`
	MerchantName            string   `json:"merchantName" binding:"required"`            //品牌方/乙方姓名
	IdentificationNumber    string   `json:"identificationNumber" binding:"required"`    //证件号码
	ContactAddress          string   `json:"contactAddress" binding:"required"`          //联系地址
	MerchantTel             string   `json:"merchantTel" binding:"required"`             //联系电话
	AccountName             string   `json:"accountName"`                                //户名
	BankAccount             string   `json:"bankAccount"`                                //银行账号
	BankOfDeposit           string   `json:"bankOfDeposit"`                              //开户行
	AddressRange            string   `json:"addressRange" binding:"required"`            //店铺地址范围
	DocumentType            string   `json:"documentType" binding:"required"`            //证件类型
	FranchiseFee            *float64 `json:"franchiseFee" binding:"required"`            //加盟费
	BrandFee                *float64 `json:"brandFee" binding:"required"`                //品牌管理费
	Bond                    *float64 `json:"bond" binding:"required"`                    //保证金
	InitialPaymentTerm      string   `json:"initialPaymentTerm" binding:"required"`      //首次付款期限
	InitialPaymentAmount    string   `json:"initialPaymentAmount" binding:"required"`    //首次付款金额
	BalanceSettlementPeriod string   `json:"balanceSettlementPeriod" binding:"required"` //余款结清期限
	AccountStatus           string   `json:"accountStatus" binding:"required"`           //账号状态
	//ContactEmail            string  `json:"contactEmail" binding:"required"`            //联系邮箱
	FranchiseFeeRemark string `json:"franchiseFeeRemark"`
	BrandFeeRemark     string `json:"brandFeeRemark"`
	BondRemark         string `json:"bondRemark"`
}

type GetDownloadUrlForm struct {
	FileId string `json:"file_id" form:"file_id"`
}

type DocumentDownloadForm struct {
	FlowId     string `json:"flow_id" form:"flow_id" binding:"required"`
	ContractId int    `json:"contract_id" form:"contract_id" binding:"required"`
}

type SiteSelectionContractForm struct {
	MerchantId           int    `json:"merchantId"`
	MerchantTel          string `json:"merchantTel"`
	SalesmanId           int    `json:"salesmanId"`
	SalesmanName         string `json:"salesmanName"`
	SalesmanTel          string `json:"salesmanTel"`
	MerchantName         string `json:"MerchantName" binding:"required"`         //品牌方/乙方姓名
	DocumentType         string `json:"documentType" binding:"required"`         //证件类型
	IdentificationNumber string `json:"identificationNumber" binding:"required"` //证件号码
	Address              string `json:"address" binding:"required"`              //门店地址
	Area                 string `json:"area" binding:"required"`                 //面积
	SpecialBusiness      string `json:"specialBusiness" binding:"required"`      //属于特殊商圈0:是1:否
	OpeningDate          string `json:"openingDate" binding:"required"`          //开业日期
	Year                 string `json:"year"`                                    //开业年
	Month                string `json:"month"`                                   //开业月
	Day                  string `json:"day"`                                     //开业日
	ContractType         int    `json:"contractType"  binding:"required"`        //合同类型  这里只允许传 2-加盟选址合同 5-电子迁址合同
	ShopId               int    `json:"shop_id"`                                 //门店ID
	CloseRemark          string `json:"close_remark"`                            //关闭备注
	CloseSrc             string `json:"close_src"`                               //关闭凭证
	LastContractId       int    `json:"last_contract_id" binding:"required"`     //上一份的合同ID
	PId                  int    `json:"p_id" binding:"required"`                 //上一份父级合同id
}

type SignCallbackForm struct {
	Action            string `json:"action"`
	FlowId            string `json:"flowId"`
	BusinessScence    string `json:"businessScence"`
	FlowStatus        string `json:"flowStatus"`
	CreateTime        string `json:"createTime"`
	EndTime           string `json:"endTime"`
	StatusDescription string `json:"statusDescription"`
	Timestamp         int64  `json:"timestamp"`
}

type DeleteESignUserByThirdIdForm struct {
	MerchantId *int `json:"merchantId" binding:"required"`
}

type GetPersonInfoForm struct {
	MerchantId *int `json:"merchantId" form:"merchantId" binding:"required"`
}

type RushSignForm struct {
	ContractId int `json:"contract_id" form:"contract_id" binding:"required"`
}

type RevokeForm struct {
	FlowId string `json:"flowId" form:"flowId" binding:"required"`
}

type SearchWordsPositionForm struct {
	FileId   string `json:"file_id" form:"file_id"`
	Keywords string `json:"keywords" form:"keywords"`
}

type AddComponentForm struct {
	TemplateId      string      `json:"templateId"`
	StructComponent []Component `json:"structComponent"`
}

type Component struct {
	Id      string           `json:"id"`
	Key     string           `json:"key"`
	Type    int              `json:"type"`
	Context ComponentContext `json:"context"`
}

type ComponentContext struct {
	Label    string       `json:"label"`
	Required bool         `json:"required"`
	Limit    string       `json:"limit"`
	Style    Style        `json:"style"`
	Pos      ComponentPos `json:"pos"`
}

type Style struct {
	Width     float64 `json:"width"`
	Height    float64 `json:"height"`
	Font      int     `json:"font"`
	FontSize  float64 `json:"fontSize"`
	TextColor string  `json:"textColor"`
}

type ComponentPos struct {
	Page int     `json:"page"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
}
