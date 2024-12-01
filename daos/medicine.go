package daos

import (
	"errors"
	"web/forms/req"
	"web/forms/rsp"
	"web/global"
	"web/model"
)

func HerbList(form req.HerbList) (res rsp.HerbListRsp, err error) {
	db := global.DB

	obj := new(model.Herb)

	var (
		total    int64
		dataList []model.Herb
	)

	total, dataList, err = obj.GetPageList(db, form)
	if err != nil {
		return
	}

	list := make([]rsp.Herb, 0, len(dataList))

	for _, d := range dataList {
		list = append(list, rsp.Herb{
			Id:           d.Id,
			Name:         d.Name,
			Abbreviation: d.Abbreviation,
			Nature:       d.Nature,
			Brief:        d.Brief,
			UserId:       d.UserId,
		})
	}

	res.List = list

	res.Total = total

	return
}

// 用户保存中药
func SaveHerb(form req.SaveHerb) (err error) {
	db := global.DB

	obj := new(model.Herb)

	list, err := obj.GetListByNameAndUserId(db, form.Name, form.UserId)

	if err != nil {
		return err
	}

	if len(list) > 0 {
		return errors.New("中药名称已存在")
	}

	data := model.Herb{
		Name:         form.Name,
		Abbreviation: form.Abbreviation,
		Nature:       form.Nature,
		Brief:        form.Brief,
		UserId:       form.UserId,
	}

	err = obj.Create(db, data)
	if err != nil {
		return
	}

	return
}
