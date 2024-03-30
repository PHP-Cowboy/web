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
