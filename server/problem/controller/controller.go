package controller

import (
	"fmt"
	"server/account/constant"
	"server/problem/model"
	"server/problem/serializer"
	"server/problem/validation"
	db "server/setting/model"
	"server/setting/request"
	"strconv"
)

type ProblemAPI struct{ request.Controller }
type ProblemCountAPI struct {
	request.Controller
}

func (c *ProblemCountAPI) Get() {
	if c.RequestUser().UserType != constant.Admin {
		c.Error("权限不足")
		return
	}
	var count = 0
	db.GetDB().Table("problem").Count(&count)
	c.Success(count)
}

func (c *ProblemAPI) Post() {

	data := validation.ProblemAddValid{}
	c.Check(&data, true, "admin")

	content := data.Content
	option := data.Option
	problemType := data.Type
	answer := data.Answer
	difficult := data.Difficult

	if c.RequestUser().UserType != constant.Admin {
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

	id, _ := c.GetInt("ID")
	content := c.GetString("Content")
	problemType := c.GetString("Type")
	difficult := c.GetString("Difficult")
	currentPage, _ := c.GetInt("CurrentPage")
	pageSize, _ := c.GetInt("PageSize")
	size,_ := c.GetInt("Size")




	var newIdStr = ""
	if id == 0 {
		newIdStr = ""
	} else {
		newIdStr = strconv.Itoa(id)
	}
	newId := "%" + newIdStr + "%"
	newContent := "%" + content + "%"
	newDifficult := "%" + difficult + "%"
	newproblemType := "%" + problemType + "%"

	type Problem struct {
		serializer.ProblemSerialize
	}
	var problems []Problem
	var problem = model.Problem{}

	if pageSize == 0 && currentPage == 0 && size == 0 {
		db.GetDB().Table("problem").Where("id like ? and content like ? and type like ? and difficult like ?", newId, newContent, newproblemType, newDifficult).Find(&problem)
		c.Success(problem)
	}

	if pageSize > 0 && currentPage > 0 {
		db.GetDB().Table("problem").Where("id like ? and content like ? and type like ? and difficult like ?", newId, newContent, newproblemType, newDifficult).Limit(pageSize).Offset((currentPage - 1) * pageSize).Order("problem.id").Find(&problems)
		c.Success(problems)
	}

	if size > 30 {
		size = 30
	}

	if size > 0 {
		//SELECT * FROM problem as t1 WHERE t1.id>=(RAND()*(SELECT MAX(id) FROM problem)) and id like '1%' LIMIT 3;
		//db.GetDB().Where("id >= ? and type like ? and difficult = ?", db.GetDB().Table("problem").Select("MAX(id)").SubQuery(), newproblemType, difficult).Limit(size).Find(&problem)
		db.GetDB().Table("problem").Where("type like ? and difficult = ?",newproblemType,difficult).Limit(size).Find(&problems)
		c.Success(problems)
	}
}
