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
			Id:           d.Id,
			Name:         d.Name,
			Abbreviation: d.Abbreviation,
			Content:      d.Content,
		})
	}

	res.List = list

	res.Total = total

	return
}

func Formula(form req.Id) (res rsp.FormulaDetail, err error) {
	db := global.DB

	obj := new(model.Formula)

	var formula model.Formula

	formula, err = obj.GetOne(db, form)

	if err != nil {
		return
	}

	var (
		dataList []model.Dose
	)

	doseObj := new(model.Dose)

	dataList, err = doseObj.GetListByFormulaId(db, formula.Id)
	if err != nil {
		return
	}

	herbIds := make([]int, 0, len(dataList))

	for _, d := range dataList {
		herbIds = append(herbIds, d.HerbId)
	}

	herbObj := new(model.Herb)

	var herbList []model.Herb

	herbList, err = herbObj.GetListByIds(db, herbIds)
	if err != nil {
		return
	}

	herbMp := make(map[int]string)

	for _, herb := range herbList {
		herbMp[herb.Id] = herb.Name
	}

	p := make([]rsp.Proportion, 0, len(herbList))
	for _, d := range dataList {
		name, ok := herbMp[d.HerbId]

		if !ok {
			name = ""
		}

		p = append(p, rsp.Proportion{
			HerbName: name,
			Weight:   d.Weight,
		})
	}

	res.Name = formula.Name
	res.Proportion = p
	res.Content = formula.Content

	return
}

// 用户保存方剂
func SaveMyFormula(form req.SaveMyFormula) (err error) {
	tx := global.DB.Begin()

	data := model.Formula{
		Name:         form.Name,
		Abbreviation: form.Abbreviation,
		Content:      form.Content,
		UserId:       form.UserId,
	}

	err = tx.Model(&model.Formula{}).Create(&data).Error
	if err != nil {
		tx.Rollback()
		return
	}

	doseList := make([]model.Dose, 0, len(form.ProportionList))

	for _, proportion := range form.ProportionList {
		doseList = append(doseList, model.Dose{
			FormulaId: data.Id,
			HerbId:    proportion.HerbId,
			Weight:    proportion.Weight,
		})
	}

	obj := new(model.Dose)

	err = obj.CreateInBatches(tx, doseList)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return
}
