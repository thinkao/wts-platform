package controller

import (
	"server/account/constant"
	"server/problem/model"
	"server/problem/serializer"
	"server/problem/validation"
	db "server/setting/model"
	"server/setting/request"
)

type ProblemAPI struct{ request.Controller }

func (c *ProblemAPI) Post() {

	data := validation.ProblemAddValid{}
	c.Check(&data, true, "admin")

	content := data.Content
	option := data.Option
	problemType := data.Type
	answer := data.Answer
	difficult := data.Difficult

	if c.RquestUser().UserType != constant.Admin {
		c.Error("权限不足")
		return
	}

	problem := model.Problem{Content: content, Option: option, Answer: answer, Type: problemType, Difficult: difficult}
	db.GetDB().Create(&problem)
	c.Success(nil)

}

func (c *ProblemAPI) Delete() {

	data := validation.DeleteByIdValid{}
	c.Check(&data, true, "admin")
	id := data.Id
	if db.GetDB().Where("id = ?", id).First(&model.Problem{}).Error == nil {
		db.GetDB().Delete(model.Problem{}, "id = ?", id)
	} else {
		c.Error("系统没有该题信息")
		return
	}
	c.Success(nil)
}

func (c *ProblemAPI) Put() {

	data := validation.ProblemUpdateValid{}
	c.Check(&data, true, "admin")
	id := data.Id
	content := data.Content
	option := data.Option
	problemType := data.Type
	answer := data.Answer
	difficult := data.Difficult

	problem := model.Problem{}

	if db.GetDB().Where("id = ?", id).First(&model.Problem{}).Error != nil {

		c.Error("系统没有该题信息")
		return

	}

	db.GetDB().Where("id = ?", id).Model(&problem).Updates(model.Problem{Content: content, Option: option, Type: problemType, Answer: answer, Difficult: difficult})

	c.Success(nil)
}

func (c *ProblemAPI) Get() {

	data := validation.ProblemGetValid{}
	c.Check(&data, true, "all")

	problemType := data.Type
	difficult := data.Difficult
	size := data.Size

	type Problem struct {
		serializer.ProblemSerialize
	}

	var problem []Problem
	//SELECT * FROM problem as t1 WHERE t1.id>=(RAND()*(SELECT MAX(id) FROM problem)) and id like '1%' LIMIT 3;
	db.GetDB().Where("id >= ? and type = ? and difficult = ?", db.GetDB().Table("problem").Select("MAX(id)").SubQuery(), problemType, difficult).Limit(size).Find(&problem)

	c.Success(nil)
}
