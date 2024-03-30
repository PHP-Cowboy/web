package daos

import (
	"web/forms/req"
	"web/forms/rsp"
	"web/global"
	"web/model"
)

func ClinicalList(form req.ClinicalList) (res rsp.ClinicalRsp, err error) {
	db := global.DB

	obj := model.Clinical{}

	var (
		total        int64
		clinicalList []model.Clinical
	)

	total, clinicalList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.Clinical, 0, len(clinicalList))

	for _, cl := range clinicalList {
		list = append(list, rsp.Clinical{
			Id:      cl.Id,
			Symptom: cl.Symptom,
		})
	}

	res.List = list
	res.Total = total

	return
}

// 大数据饼图
func BigDataPieChart(form req.BigDataPieChart) (res rsp.BigDataPieChartRsp, err error) {
	db := global.DB

	bigDataObj := model.BigData{}
	categoryObj := model.BigDataCategory{}

	var (
		bigDataList []model.BigData
		category    model.BigDataCategory
	)

	bigDataList, err = bigDataObj.GetListByCategoryId(db, form.CategoryId)

	if err != nil {
		return
	}

	list := make([]rsp.BigData, 0, len(bigDataList))

	for _, l := range bigDataList {
		list = append(list, rsp.BigData{
			Name: l.Name,
			Num:  l.Num,
		})
	}

	category, err = categoryObj.GetOneById(db, form.CategoryId)

	if err != nil {
		return
	}

	res.List = list
	res.Category = rsp.Category{
		Id:   category.Id,
		Name: category.Name,
	}

	return
}
