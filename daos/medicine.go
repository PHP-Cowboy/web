package daos

import (
	"web/forms/req"
	"web/forms/rsp"
	"web/global"
	"web/model"
)

func MedicineList(form req.MedicineList) (res rsp.MedicineListRsp, err error) {
	db := global.DB

	obj := new(model.Medicine)

	var (
		total    int64
		dataList []model.Medicine
	)

	total, dataList, err = obj.GetPageList(db, form)
	if err != nil {
		return
	}

	list := make([]rsp.Medicine, 0, len(dataList))

	for _, d := range dataList {
		list = append(list, rsp.Medicine{
			Id:   d.Id,
			Name: d.Name,
		})
	}

	res.List = list

	res.Total = total

	return
}

// 用户保存方剂
func SaveMedicine(form req.SaveMedicine) (err error) {
	db := global.DB

	obj := new(model.Medicine)

	data := model.Medicine{
		Name:   form.Name,
		UserId: form.UserId,
	}

	err = obj.Create(db, data)
	if err != nil {
		return
	}

	return
}
