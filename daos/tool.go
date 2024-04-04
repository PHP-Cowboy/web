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
