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
		total    int64
		dataList []model.Clinical
	)

	total, dataList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.Clinical, 0, len(dataList))

	for _, cl := range dataList {
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
		data        model.BigDataCategory
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

	data, err = categoryObj.GetOneById(db, form.CategoryId)

	if err != nil {
		return
	}

	res.List = list
	res.Category = rsp.Category{
		Id:   data.Id,
		Name: data.Name,
	}

	return
}
