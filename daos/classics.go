package daos

import (
	"errors"
	"gorm.io/gorm"
	"web/forms/req"
	"web/forms/rsp"
	"web/global"
	"web/model"
)

// 典籍分类列表
func ClassicsCategoryList(form req.ClassicsCategoryList) (res rsp.ClassicsCategoryRsp, err error) {
	db := global.DB

	obj := model.ClassicsCategory{}

	var (
		total    int64
		dataList []model.ClassicsCategory
	)

	total, dataList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.ClassicsCategory, 0, len(dataList))

	for _, cl := range dataList {
		list = append(list, rsp.ClassicsCategory{
			Id:      cl.Id,
			Name:    cl.Name,
			Picture: cl.Picture,
		})
	}

	res.List = list
	res.Total = total

	return
}

// 某个分类的典籍列表
func BookListByCategory(form req.BookListByCategory) (res rsp.BookListRsp, err error) {
	db := global.DB

	obj := model.Classics{}

	var (
		total    int64
		dataList []model.Classics
	)

	total, dataList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.BookList, 0, len(dataList))

	for _, cl := range dataList {
		list = append(list, rsp.BookList{
			Id:         cl.Id,
			CategoryId: cl.CategoryId,
			Name:       cl.Name,
			Author:     cl.Author,
			Dynasty:    cl.Dynasty,
		})
	}

	res.List = list
	res.Total = total

	return
}

// 典籍目录列表
func CatalogueList(form req.CatalogueList) (res rsp.CatalogueRsp, err error) {
	db := global.DB

	obj := model.ClassicsContent{}

	var (
		total    int64
		dataList []model.ClassicsContent
	)

	total, dataList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.Catalogue, 0, len(dataList))

	for _, cl := range dataList {
		list = append(list, rsp.Catalogue{
			Id:         cl.ClassicsId,
			ClassicsId: cl.ClassicsId,
			Title:      cl.Title,
			Chapter:    cl.Chapter,
		})
	}

	res.List = list
	res.Total = total

	return
}

// 典籍的内容
func BookContent(form req.Id) (res rsp.Catalogue, err error) {
	db := global.DB

	obj := model.ClassicsContent{}

	var (
		data model.ClassicsContent
	)

	data, err = obj.GetOneById(db, form.Id)

	if err != nil {
		return
	}

	res.Id = data.Id
	res.ClassicsId = data.ClassicsId
	res.Title = data.Title
	res.Content = data.Content

	return
}

// 上一章
func PrevBookContent(form req.PrevNext) (res rsp.Catalogue, err error) {
	db := global.DB

	obj := model.ClassicsContent{}

	var (
		data model.ClassicsContent
	)

	data, err = obj.GetPrevByIdClassicsId(db, form.Id, form.ClassicsId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("已经是第一章了")
			return
		}
		return
	}

	res.Id = data.Id
	res.ClassicsId = data.ClassicsId
	res.Title = data.Title
	res.Content = data.Content

	return
}

// 下一章
func NextBookContent(form req.PrevNext) (res rsp.Catalogue, err error) {
	db := global.DB

	obj := model.ClassicsContent{}

	var (
		data model.ClassicsContent
	)

	data, err = obj.GetNextByIdClassicsId(db, form.Id, form.ClassicsId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("已经是最后一章了")
			return
		}
		return
	}

	res.Id = data.Id
	res.ClassicsId = data.ClassicsId
	res.Title = data.Title
	res.Content = data.Content

	return
}
