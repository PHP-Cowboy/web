package daos

import (
	"errors"
	"gorm.io/gorm"
	"web/forms/req"
	"web/forms/rsp"
	"web/global"
	"web/model"
)

func ToolList(form req.ToolList) (res rsp.ToolRsp, err error) {
	db := global.DB

	obj := model.Tool{}

	var (
		total    int64
		dataList []model.Tool
	)

	total, dataList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.Tool, 0, len(dataList))

	for _, cl := range dataList {
		list = append(list, rsp.Tool{
			Id:      cl.Id,
			Name:    cl.Name,
			Desc:    cl.Desc,
			Picture: cl.Picture,
			Router:  cl.Router,
		})
	}

	res.List = list
	res.Total = total

	return
}

// 思维导图列表
func MindMapList(form req.MindMapList) (res rsp.MindMapRsp, err error) {
	db := global.DB

	obj := model.MindMap{}

	var (
		total    int64
		dataList []model.MindMap
	)

	total, dataList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.MindMap, 0, len(dataList))

	for _, cl := range dataList {
		list = append(list, rsp.MindMap{
			Id:      cl.Id,
			Title:   cl.Title,
			Picture: cl.Picture,
		})
	}

	res.List = list
	res.Total = total

	return
}

// 疾病分类
func DiseaseCategory() (res []rsp.Group, err error) {
	db := global.DB

	obj := model.DiseaseCategory{}

	var (
		dataList []model.DiseaseCategory
	)

	dataList, err = obj.GetList(db)

	if err != nil {
		return
	}

	list := make([]rsp.Group, 0, len(dataList))

	for _, cl := range dataList {
		list = append(list, rsp.Group{
			Id:       cl.Id,
			ParentId: cl.ParentId,
			Name:     cl.Name,
			Children: []rsp.Group{},
		})
	}

	res = buildTree(list, 0)

	return
}

// 生成树形结构
func buildTree(groups []rsp.Group, parentId int) []rsp.Group {
	var tree []rsp.Group
	for _, group := range groups {
		if group.ParentId == parentId {
			children := buildTree(groups, group.Id)
			if children != nil {
				group.Children = children
			}
			tree = append(tree, group)
		}
	}
	return tree
}

// 疾病信息
func Disease(form req.DiseaseList) (res rsp.DiseaseRsp, err error) {
	db := global.DB

	obj := model.Disease{}

	var (
		total    int64
		dataList []model.Disease
	)

	dataList, err = obj.GetListByCategoryId(db, form.CategoryId)

	if err != nil {
		return
	}

	list := make([]rsp.Disease, 0, len(dataList))

	for _, cl := range dataList {
		list = append(list, rsp.Disease{
			Id:      cl.Id,
			Title:   cl.Title,
			Content: cl.Content,
		})
	}

	res.List = list
	res.Total = total

	return
}

// 名医心法
func MindMethodList(form req.MindMethodList) (res rsp.MindMethodRsp, err error) {
	db := global.DB

	obj := model.MindMethod{}

	var (
		total    int64
		dataList []model.MindMethod
	)

	total, dataList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.MindMethod, 0, len(dataList))

	for _, cl := range dataList {
		list = append(list, rsp.MindMethod{
			Id:     cl.Id,
			Title:  cl.Title,
			Doctor: cl.Doctor,
		})
	}

	res.List = list
	res.Total = total

	return
}

// 名医心法详情
func MindMethod(form req.Id) (res rsp.MindMethod, err error) {
	db := global.DB

	obj := model.MindMethod{}

	var data model.MindMethod

	data, err = obj.GetFirst(db, form)

	if err != nil {
		return
	}

	res.Id = data.Id
	res.Title = data.Title
	res.Doctor = data.Doctor
	res.Content = data.Content

	return
}

// 中医方剂分类列表
func CommonlyPrescriptionCategoryList(form req.CommonlyPrescriptionCategoryList) (res rsp.CommonlyPrescriptionCategoryRsp, err error) {
	db := global.DB

	obj := model.CommonlyPrescriptionCategory{}

	var (
		total    int64
		dataList []model.CommonlyPrescriptionCategory
	)

	total, dataList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.CommonlyPrescriptionCategory, 0, len(dataList))

	for _, cl := range dataList {
		list = append(list, rsp.CommonlyPrescriptionCategory{
			Id:   cl.Id,
			Name: cl.Name,
		})
	}

	res.List = list
	res.Total = total

	return
}

// 常用方剂列表
func CommonlyPrescriptionList(form req.CommonlyPrescriptionList) (res rsp.CommonlyPrescriptionRsp, err error) {
	db := global.DB

	obj := model.CommonlyPrescription{}

	var (
		total    int64
		dataList []model.CommonlyPrescription
	)

	total, dataList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.CommonlyPrescription, 0, len(dataList))

	for _, cl := range dataList {
		list = append(list, rsp.CommonlyPrescription{
			Id:          cl.Id,
			Title:       cl.Title,
			Provenance:  cl.Provenance,
			Constituent: cl.Constituent,
			Usage:       cl.Usage,
			Efficacy:    cl.Efficacy,
			Mainly:      cl.Mainly,
			Song:        cl.Song,
			CategoryId:  cl.CategoryId,
		})
	}

	res.List = list
	res.Total = total

	return
}

// 中医方剂详情
func CommonlyPrescription(form req.Id) (res rsp.CommonlyPrescription, err error) {
	db := global.DB

	obj := model.CommonlyPrescription{}

	var data model.CommonlyPrescription

	data, err = obj.GetFirst(db, form)

	if err != nil {
		return
	}

	res.Id = data.Id
	res.Title = data.Title
	res.Provenance = data.Provenance
	res.Constituent = data.Constituent
	res.Usage = data.Usage
	res.Efficacy = data.Efficacy
	res.Mainly = data.Mainly
	res.Song = data.Song
	res.CategoryId = data.CategoryId

	return
}

// 方剂大全列表
func CompleteCollectionPrescriptionList(form req.CompleteCollectionPrescriptionList) (res rsp.CompleteCollectionPrescriptionRsp, err error) {

	db := global.DB

	obj := model.CompleteCollectionPrescription{}

	var (
		total    int64
		dataList []model.CompleteCollectionPrescription
	)

	total, dataList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.CompleteCollectionPrescription, 0, len(dataList))

	for _, cl := range dataList {
		list = append(list, rsp.CompleteCollectionPrescription{
			Id:         cl.Id,
			Title:      cl.Title,
			Provenance: cl.Provenance,
		})
	}

	res.List = list
	res.Total = total

	return
}

// 方剂大全详情
func CompleteCollectionPrescription(form req.Id) (res rsp.CompleteCollectionPrescription, err error) {
	db := global.DB

	obj := model.CompleteCollectionPrescription{}

	var data model.CompleteCollectionPrescription

	data, err = obj.GetFirst(db, form)

	if err != nil {
		return
	}

	res.Id = data.Id
	res.Title = data.Title
	res.Provenance = data.Provenance
	res.Content = data.Content

	return
}

// 题库类别列表
func QuestionCategoryList(form req.QuestionCategoryList) (res rsp.QuestionCategoryRsp, err error) {

	db := global.DB

	obj := model.QuestionCategory{}

	var (
		total    int64
		dataList []model.QuestionCategory
	)

	total, dataList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.QuestionCategory, 0, len(dataList))

	for _, cl := range dataList {
		list = append(list, rsp.QuestionCategory{
			Id:   cl.Id,
			Name: cl.Name,
		})
	}

	res.List = list
	res.Total = total

	return
}

func QuestionList(form req.Id) (list []rsp.Question, err error) {

	db := global.DB

	obj := model.Question{}

	var (
		dataList []model.Question
	)

	dataList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list = make([]rsp.Question, 0, len(dataList))

	for _, cl := range dataList {
		list = append(list, rsp.Question{
			Id:         cl.Id,
			CategoryId: cl.CategoryId,
			Number:     cl.Number,
			Topic:      cl.Topic,
			A:          cl.A,
			B:          cl.B,
			C:          cl.C,
			D:          cl.D,
			E:          cl.E,
			Answer:     cl.Answer,
			Analysis:   cl.Analysis,
		})
	}

	return
}

// 题目的内容
func QuestionInfo(form req.Id) (res rsp.Question, err error) {
	db := global.DB

	obj := model.Question{}

	var (
		data model.Question
	)

	data, err = obj.GetOneById(db, form.Id)

	if err != nil {
		return
	}

	res.Id = data.Id
	res.CategoryId = data.CategoryId
	res.Number = data.Number
	res.Topic = data.Topic
	res.A = data.A
	res.B = data.B
	res.C = data.C
	res.D = data.D
	res.E = data.E
	res.Answer = data.Answer
	res.Analysis = data.Analysis

	return
}

// 上一题
func QuestionPrev(form req.Id) (res rsp.Question, err error) {
	db := global.DB

	obj := model.Question{}

	var (
		data model.Question
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
	res.CategoryId = data.CategoryId
	res.Number = data.Number
	res.Topic = data.Topic
	res.A = data.A
	res.B = data.B
	res.C = data.C
	res.D = data.D
	res.E = data.E
	res.Answer = data.Answer
	res.Analysis = data.Analysis

	return
}

// 下一题
func QuestionNext(form req.Id) (res rsp.Question, err error) {
	db := global.DB

	obj := model.Question{}

	var (
		data model.Question
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
	res.CategoryId = data.CategoryId
	res.Number = data.Number
	res.Topic = data.Topic
	res.A = data.A
	res.B = data.B
	res.C = data.C
	res.D = data.D
	res.E = data.E
	res.Answer = data.Answer
	res.Analysis = data.Analysis

	return
}
