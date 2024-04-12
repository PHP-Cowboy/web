package daos

import (
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
		toolList []model.Tool
	)

	total, toolList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.Tool, 0, len(toolList))

	for _, cl := range toolList {
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
		mindList []model.MindMap
	)

	total, mindList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.MindMap, 0, len(mindList))

	for _, cl := range mindList {
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
		cateList []model.DiseaseCategory
	)

	cateList, err = obj.GetList(db)

	if err != nil {
		return
	}

	list := make([]rsp.Group, 0, len(cateList))

	for _, cl := range cateList {
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
		total       int64
		diseaseList []model.Disease
	)

	diseaseList, err = obj.GetListByCategoryId(db, form.CategoryId)

	if err != nil {
		return
	}

	list := make([]rsp.Disease, 0, len(diseaseList))

	for _, cl := range diseaseList {
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
		total          int64
		mindMethodList []model.MindMethod
	)

	total, mindMethodList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.MindMethod, 0, len(mindMethodList))

	for _, cl := range mindMethodList {
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

	var mindMethod model.MindMethod

	mindMethod, err = obj.GetFirst(db, form)

	if err != nil {
		return
	}

	res.Id = mindMethod.Id
	res.Title = mindMethod.Title
	res.Doctor = mindMethod.Doctor
	res.Content = mindMethod.Content

	return
}

// 中医方剂分类列表
func PrescriptionCategoryList(form req.PrescriptionCategoryList) (res rsp.PrescriptionCategoryRsp, err error) {
	db := global.DB

	obj := model.PrescriptionCategory{}

	var (
		total          int64
		mindMethodList []model.PrescriptionCategory
	)

	total, mindMethodList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.PrescriptionCategory, 0, len(mindMethodList))

	for _, cl := range mindMethodList {
		list = append(list, rsp.PrescriptionCategory{
			Id:   cl.Id,
			Name: cl.Name,
		})
	}

	res.List = list
	res.Total = total

	return
}

// 中医方剂列表
func PrescriptionList(form req.PrescriptionList) (res rsp.PrescriptionRsp, err error) {
	db := global.DB

	obj := model.Prescription{}

	var (
		total          int64
		mindMethodList []model.Prescription
	)

	total, mindMethodList, err = obj.GetPageList(db, form)

	if err != nil {
		return
	}

	list := make([]rsp.Prescription, 0, len(mindMethodList))

	for _, cl := range mindMethodList {
		list = append(list, rsp.Prescription{
			Id:         cl.Id,
			Title:      cl.Title,
			Content:    cl.Content,
			Provenance: cl.Provenance,
			CategoryId: cl.CategoryId,
		})
	}

	res.List = list
	res.Total = total

	return
}

// 中医方剂详情
func Prescription(form req.Id) (res rsp.Prescription, err error) {
	db := global.DB

	obj := model.Prescription{}

	var mindMethod model.Prescription

	mindMethod, err = obj.GetFirst(db, form)

	if err != nil {
		return
	}

	res.Id = mindMethod.Id
	res.Title = mindMethod.Title
	res.Content = mindMethod.Content
	res.Provenance = mindMethod.Provenance
	res.CategoryId = mindMethod.CategoryId

	return
}
