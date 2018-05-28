package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/fatih/structs"
	"github.com/udistrital/api_financiera/models"
	"github.com/udistrital/utils_oas/formatdata"

	"github.com/astaxie/beego"
)

// AnulacionDisponibilidadController operations for AnulacionDisponibilidad
type AnulacionDisponibilidadController struct {
	beego.Controller
}

// URLMapping ...
func (c *AnulacionDisponibilidadController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create AnulacionDisponibilidad
// @Param	body		body 	models.AnulacionDisponibilidad	true		"body for AnulacionDisponibilidad content"
// @Success 201 {int} models.AnulacionDisponibilidad
// @Failure 403 body is empty
// @router / [post]
func (c *AnulacionDisponibilidadController) Post() {
	var v models.AnulacionDisponibilidad
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddAnulacionDisponibilidad(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get AnulacionDisponibilidad by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.AnulacionDisponibilidad
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AnulacionDisponibilidadController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetAnulacionDisponibilidadById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get AnulacionDisponibilidad
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.AnulacionDisponibilidad
// @Failure 403
// @router / [get]
func (c *AnulacionDisponibilidadController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllAnulacionDisponibilidad(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the AnulacionDisponibilidad
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.AnulacionDisponibilidad	true		"body for AnulacionDisponibilidad content"
// @Success 200 {object} models.AnulacionDisponibilidad
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AnulacionDisponibilidadController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	var fields []string
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	id, _ := strconv.Atoi(idStr)
	v := models.AnulacionDisponibilidad{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateAnulacionDisponibilidadById(&v, fields...); err == nil {
			c.Ctx.Output.SetStatus(201)
			alert := models.Alert{Type: "success", Code: "S_A01", Body: v}
			c.Data["json"] = alert
		} else {
			alertdb := structs.Map(err)
			var code string
			formatdata.FillStruct(alertdb["Code"], &code)
			alert := models.Alert{Type: "error", Code: "E_" + code, Body: err}
			c.Data["json"] = alert
		}
	} else {
		alert := models.Alert{Type: "success", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = alert
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the AnulacionDisponibilidad
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AnulacionDisponibilidadController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteAnulacionDisponibilidad(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}


// TotalAnulacionDisponibilidad ...
// @Title TotalAnulacionDisponibilidad
// @Description numero de dispònibilidades segun vigencia o rango de fechas
// @Param	vigencia		query 	string	true		"vigencia para la consulta del total de disponibilidades"
// @Param	UnidadEjecutora	query	string	false	"unidad ejecutora de las solicitudes a consultar"
// @Param	rangoinicio		query 	string	true		"opcional, fecha inicio de consulta de cdp"
// @Param	rangofin		query 	string	true		"opcional, fecha fin de consulta de cdp"
// @Success 201 {int} total
// @Failure 403 vigencia is empty
// @router /TotalAnulacionDisponibilidad/:vigencia [get]
func (c *AnulacionDisponibilidadController) TotalAnulacionDisponibilidad() {
	vigenciaStr := c.Ctx.Input.Param(":vigencia")
	vigencia, err := strconv.Atoi(vigenciaStr)
	//var startrange string
	//var endrange string
	/*if r := c.GetString("rangoinicio"); r != "" {
		startrange = r

	}
	if r := c.GetString("rangofin"); r != "" {
		endrange = r
	}*/
	UnidadEjecutora, err2 := c.GetInt("UnidadEjecutora")
	if err == nil && err2 == nil {
		total, err := models.GetTotalAnulacionDisponibilidades(vigencia, UnidadEjecutora)
		if err == nil {
			c.Data["json"] = total
		} else {
			c.Data["json"] = models.Alert{Code: "E_0458", Body: err.Error(), Type: "error"}
		}
	} else {
		c.Data["json"] = models.Alert{Code: "E_0458", Body: "Not enough parameter", Type: "error"}
	}

	c.ServeJSON()
}
