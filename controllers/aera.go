package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"iHome/models"
)

type AreaController struct {
	beego.Controller
}

func (this *AreaController)RetData(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *AreaController) GetArea() {
	beego.Info("connecct success!")

	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer c.RetData(resp)

	//从session模块中拿数据

	//从数据库拿到数据
	var areas []models.Area

	o := orm.NewOrm()
	num, err := o.QueryTable("area").All(&areas)
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	if num == 0{
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	}
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = areas

	beego.Info("query data success,resp=", resp, "num = ", num)

}
