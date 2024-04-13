package daos

import (
	"web/forms/req"
	"web/forms/rsp"
	"web/global"
	"web/model"
)

// 医疗案例列表
func MedicalCasesList(form req.MedicalCasesList) (res rsp.MedicalCasesRsp, err error) {
	db := global.DB

	obj := &model.MedicalCases{}

	var (
		dataList []model.MedicalCases
		total    int64
	)

	total, err = obj.GetLikeCount(db, form)

	if err != nil {
		return
	}

	dataList, err = obj.GetLikePageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.MedicalCases, 0, len(dataList))

	for _, m := range dataList {
		list = append(list, rsp.MedicalCases{
			Id:           m.Id,
			Symptom:      m.Symptom,
			Prescription: m.Prescription,
			Content:      m.Content,
			Provenance:   m.Provenance,
			Picture:      m.Picture,
		})
	}

	res.Total = total
	res.List = list

	return
}

// 随机获取医案列表
func RandMedicalCasesList(form req.MedicalCasesList) (res rsp.MedicalCasesRsp, err error) {
	db := global.DB

	obj := &model.MedicalCases{}

	dataList := make([]model.MedicalCases, 0, form.Size)

	dataList, err = obj.GetRandList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.MedicalCases, 0, form.Size)

	for _, m := range dataList {
		list = append(list, rsp.MedicalCases{
			Id:           m.Id,
			Symptom:      m.Symptom,
			Prescription: m.Prescription,
			Content:      m.Content,
			Provenance:   m.Provenance,
			Picture:      m.Picture,
		})
	}

	res.List = list

	return
}

// 医疗案例详情
func MedicalCasesDetail(form req.Id) (res rsp.MedicalCases, err error) {
	db := global.DB

	obj := &model.MedicalCases{}

	data, err := obj.GetFirstById(db, form.Id)
	if err != nil {
		return
	}

	res.Id = data.Id
	res.Symptom = data.Symptom
	res.Prescription = data.Prescription
	res.Content = data.Content
	res.Provenance = data.Provenance

	return
}
