package daos

import (
	"errors"
	"gorm.io/gorm"
	"web/forms/req"
	"web/forms/rsp"
	"web/global"
	"web/model"
)

// 伤寒论目录列表
func TyphoidCatalogueList(form req.TyphoidCatalogueList) (res rsp.TyphoidCatalogueRsp, err error) {
	db := global.DB

	obj := model.Typhoid{}

	var (
		total    int64
		dataList []model.Typhoid
	)

	total, dataList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.TyphoidCatalogue, 0, len(dataList))

	for _, cl := range dataList {
		list = append(list, rsp.TyphoidCatalogue{
			Id:    cl.Id,
			Title: cl.Title,
			Brief: cl.Brief,
		})
	}

	res.List = list
	res.Total = total

	return
}

// 伤寒论的内容
func TyphoidInfo(form req.Id) (res rsp.TyphoidCatalogue, err error) {
	db := global.DB

	obj := model.Typhoid{}

	var (
		data model.Typhoid
	)

	data, err = obj.GetOneById(db, form.Id)

	if err != nil {
		return
	}

	res.Id = data.Id
	res.Title = data.Title
	res.Brief = data.Brief
	res.Content = data.Content

	return
}

// 伤寒论上一章
func TyphoidPrev(form req.Id) (res rsp.TyphoidCatalogue, err error) {
	db := global.DB

	obj := model.Typhoid{}

	var (
		data model.Typhoid
	)

	data, err = obj.GetPrevById(db, form.Id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("已经是第一章了")
			return
		}
		return
	}

	res.Id = data.Id
	res.Title = data.Title
	res.Brief = data.Brief
	res.Content = data.Content

	return
}

// 伤寒论下一章
func TyphoidNext(form req.Id) (res rsp.TyphoidCatalogue, err error) {
	db := global.DB

	obj := model.Typhoid{}

	var (
		data model.Typhoid
	)

	data, err = obj.GetNextById(db, form.Id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("已经是最后一章了")
			return
		}
		return
	}

	res.Id = data.Id
	res.Title = data.Title
	res.Brief = data.Brief
	res.Content = data.Content

	return
}
