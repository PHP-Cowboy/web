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
		medicalCasesList []model.MedicalCases
		total            int64
	)

	total, err = obj.GetLikeCount(db, form)

	if err != nil {
		return
	}

	medicalCasesList, err = obj.GetLikePageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.MedicalCases, 0, len(medicalCasesList))

	for _, m := range medicalCasesList {
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

	medicalCasesList := make([]model.MedicalCases, 0, form.Size)

	medicalCasesList, err = obj.GetRandList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.MedicalCases, 0, form.Size)

	for _, m := range medicalCasesList {
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

	medicalCases, err := obj.GetFirstById(db, form.Id)
	if err != nil {
		return
	}

	res.Id = medicalCases.Id
	res.Symptom = medicalCases.Symptom
	res.Prescription = medicalCases.Prescription
	res.Content = medicalCases.Content
	res.Provenance = medicalCases.Provenance

	return
}
