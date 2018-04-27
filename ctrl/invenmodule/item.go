package ctrl

import (
	"github.com/gin-gonic/gin"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"github.com/loukmho/bcaccount_api/model/invenmodule"
	"net/http"
	"fmt"
)

func SearchItemByCode(c *gin.Context){
	c.Keys = ct.HeaderKeys

	code := c.Request.URL.Query().Get("code")

	itm := new(model.Item)
	err := itm.SearchItemByCode(ct.Dbc, code)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "success"
		rs.Data = itm
		c.JSON(http.StatusOK, rs)
	}
}

func SearchItemByKeyword(c *gin.Context){
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	itm := new(model.Item)
	itms, err := itm.SearchItemByKeycode(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "success"
		rs.Data = itms
		c.JSON(http.StatusOK, rs)
	}
}

func InsertAndEditItem(c *gin.Context){
	c.Keys = ct.HeaderKeys

	itm := &model.Item{}
	err := c.BindJSON(itm)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = itm.InsertAndEditItem(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		rs.Status = "success"
		rs.Data = itm
		c.JSON(http.StatusOK ,rs)
	}
}