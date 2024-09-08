package daos

import (
	"web/forms/req"
	"web/forms/rsp"
	"web/global"
	"web/model"
)

func FormulaList(form req.FormulaList) (res rsp.FormulaListRsp, err error) {
	db := global.DB

	obj := new(model.Formula)

	var (
		total    int64
		dataList []model.Formula
	)

	total, dataList, err = obj.GetPageList(db, form)
	if err != nil {
		return
	}

	list := make([]rsp.Formula, 0, len(dataList))

	for _, d := range dataList {
		list = append(list, rsp.Formula{
			Id:               d.Id,
			Name:             d.Name,
			NameAbbreviation: d.NameAbbreviation,
			Proportion:       d.Proportion,
			Dose:             d.Dose,
			Content:          d.Content,
		})
	}

	res.List = list

	res.Total = total

	return
}

// 用户保存方剂
func SaveMyFormula(form req.SaveMyFormula) (err error) {
	db := global.DB

	obj := new(model.Formula)

	data := model.Formula{
		Name:             form.Name,
		NameAbbreviation: form.NameAbbreviation,
		Proportion:       form.Proportion,
		Dose:             form.Dose,
		Content:          form.Content,
		UserId:           form.UserId,
	}

	err = obj.Create(db, data)
	if err != nil {
		return
	}

	return
}
