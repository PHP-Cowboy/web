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
		total        int64
		categoryList []model.ClassicsCategory
	)

	total, categoryList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.ClassicsCategory, 0, len(categoryList))

	for _, cl := range categoryList {
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
		total        int64
		classicsList []model.Classics
	)

	total, classicsList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.BookList, 0, len(classicsList))

	for _, cl := range classicsList {
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
		total       int64
		contentList []model.ClassicsContent
	)

	total, contentList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.Catalogue, 0, len(contentList))

	for _, cl := range contentList {
		list = append(list, rsp.Catalogue{
			Id:         cl.ClassicsId,
			ClassicsId: cl.ClassicsId,
			Title:      cl.Title,
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
		content model.ClassicsContent
	)

	content, err = obj.GetOneById(db, form.Id)

	if err != nil {
		return
	}

	res.Id = content.Id
	res.ClassicsId = content.ClassicsId
	res.Title = content.Title
	res.Content = content.Content

	return
}

// 上一章
func PrevBookContent(form req.PrevNext) (res rsp.Catalogue, err error) {
	db := global.DB

	obj := model.ClassicsContent{}

	var (
		content model.ClassicsContent
	)

	content, err = obj.GetPrevByIdClassicsId(db, form.Id, form.ClassicsId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("已经是第一章了")
			return
		}
		return
	}

	res.Id = content.Id
	res.ClassicsId = content.ClassicsId
	res.Title = content.Title
	res.Content = content.Content

	return
}

// 下一章
func NextBookContent(form req.PrevNext) (res rsp.Catalogue, err error) {
	db := global.DB

	obj := model.ClassicsContent{}

	var (
		content model.ClassicsContent
	)

	content, err = obj.GetNextByIdClassicsId(db, form.Id, form.ClassicsId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("已经是最后一章了")
			return
		}
		return
	}

	res.Id = content.Id
	res.ClassicsId = content.ClassicsId
	res.Title = content.Title
	res.Content = content.Content

	return
}
