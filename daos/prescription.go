package daos

import (
	"web/forms/req"
	"web/forms/rsp"
	"web/global"
	"web/model"
)

// 大数据随机获取列表
func PrescriptionRandList(form req.PrescriptionRandList) (res rsp.PrescriptionRsp, err error) {
	db := global.DB
	obj := new(model.Prescription)

	var (
		total    int64
		dataList []model.Prescription
	)

	dataList, err = obj.GetRandList(db, form)

	if err != nil {
		return
	}

	total = int64(len(dataList))

	list := make([]rsp.Prescription, 0, total)

	for _, d := range dataList {
		list = append(list, rsp.Prescription{
			Id:            d.Id,
			Name:          d.Name,
			Symptom:       d.Symptom,
			TongueQuality: d.TongueQuality,
			CoatedTongue:  d.CoatedTongue,
			Pulse:         d.Pulse,
			ModernDisease: d.ModernDisease,
		})
	}

	res.Total = total
	res.List = list
	return
}

// 大数据方剂列表
func PrescriptionList(form req.PrescriptionList) (res rsp.PrescriptionRsp, err error) {
	db := global.DB
	obj := new(model.Prescription)

	var (
		total    int64
		dataList []model.Prescription
	)

	total, dataList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.Prescription, 0, len(dataList))

	for _, d := range dataList {
		list = append(list, rsp.Prescription{
			Id:            d.Id,
			Name:          d.Name,
			Symptom:       d.Symptom,
			TongueQuality: d.TongueQuality,
			CoatedTongue:  d.CoatedTongue,
			Pulse:         d.Pulse,
			ModernDisease: d.ModernDisease,
		})
	}

	res.Total = total
	res.List = list
	return
}

// 大数据方剂详情
func PrescriptionInfo(form req.Id) (res rsp.Prescription, err error) {
	db := global.DB
	obj := new(model.Prescription)

	var data model.Prescription

	data, err = obj.GetFirst(db, form)

	if err != nil {
		return
	}

	res.Id = data.Id
	res.Name = data.Name
	res.Symptom = data.Symptom
	res.TongueQuality = data.TongueQuality
	res.CoatedTongue = data.CoatedTongue
	res.Pulse = data.Pulse
	res.ModernDisease = data.ModernDisease
	res.SymptomPic = data.SymptomPic
	res.PulsePic = data.PulsePic
	res.TonguePic = data.TonguePic

	return
}

// 大数据方剂名家列表
func PrescriptionCelebrityList(form req.PrescriptionCelebrityList) (res rsp.CelebrityRsp, err error) {
	db := global.DB
	obj := new(model.Celebrity)

	var (
		total    int64
		dataList []model.Celebrity
	)

	total, dataList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.Celebrity, 0, len(dataList))

	for _, d := range dataList {
		list = append(list, rsp.Celebrity{
			Id:             d.Id,
			NotableDoctor:  d.NotableDoctor,
			Provenance:     d.Provenance,
			Content:        d.Content,
			PrescriptionId: d.PrescriptionId,
		})
	}

	res.Total = total
	res.List = list
	return
}

// 大数据方剂名家列表
func PrescriptionCelebrityInfo(form req.Id) (res rsp.Celebrity, err error) {
	db := global.DB
	obj := new(model.Celebrity)

	var data model.Celebrity

	data, err = obj.GetFirst(db, form)

	if err != nil {
		return
	}

	res.Id = data.Id
	res.NotableDoctor = data.NotableDoctor
	res.Provenance = data.Provenance
	res.Content = data.Content
	res.PrescriptionId = data.PrescriptionId

	return
}

// 大数据方剂名家列表
func PrescriptionCelebrityInfoByPrescriptionId(form req.Id) (res rsp.Celebrity, err error) {
	db := global.DB
	obj := new(model.Celebrity)

	var data model.Celebrity

	data, err = obj.GetFirstByPrescriptionId(db, form.Id)

	if err != nil {
		return
	}

	res.Id = data.Id
	res.NotableDoctor = data.NotableDoctor
	res.Provenance = data.Provenance
	res.Content = data.Content
	res.PrescriptionId = data.PrescriptionId

	return
}

// 大数据方剂名家列表
func PrescriptionGraph(form req.PrescriptionGraph) (res []rsp.PrescriptionGraph, err error) {
	db := global.DB
	obj := new(model.PrescriptionGraph)

	var dataList []model.PrescriptionGraph

	dataList, err = obj.GetPageListByPrescriptionIdType(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.PrescriptionGraph, 0, len(dataList))

	for _, l := range dataList {
		list = append(list, rsp.PrescriptionGraph{
			Id:             l.Id,
			PrescriptionId: l.PrescriptionId,
			Type:           l.Type,
			Name:           l.Name,
			Num:            l.Num,
		})

	}

	res = list

	return
}
