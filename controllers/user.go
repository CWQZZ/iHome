package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"iHome/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController)RetData(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *UserController) Reg() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)

	//获取从前端传过来的json数据
	json.Unmarshal(this.Ctx.Input.RequestBody, &resp)
	//resp["errno"] = 4001
	//resp["errmsg"] = "查询失败"
	//this.RetData(resp)
	//beego.Info(`resp["mobile"] = `, resp["mobile"])
	//beego.Info(`resp["password"] = `, resp["password"])
	//beego.Info(`resp["sms_code"] = `, resp["sms_code"])

	//插入数据
	o := orm.NewOrm()
	user := models.User{}
	user.Password_hash = resp["password"].(string)
	user.Name = resp["mobile"].(string)
	user.Mobile = resp["mobile"].(string)

	id, err := o.Insert(&user)
	if err != nil {
		resp["errno"] = 4002
		resp["errmsg"] = "注册失败"
		return
	}
	beego.Info("reg success, id = ", id)
	resp["errno"] = 0
	resp["errmsg"] = "注册成功"

	this.SetSession("name",user.Name)
}