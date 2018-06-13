package ctrl

import (
	"github.com/loukmho/bcaccount_api/model/invenmodule"
	"fmt"
	"net/http"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"github.com/gin-gonic/gin"
)

func SearchStkAdjustByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	sta := new(model.StkAdjust)
	err := sta.SearchStkAdjustByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = sta
		c.JSON(http.StatusOK, rs)
	}
}

func SearchStkAdjustByKeyword(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	sta := new(model.StkAdjust)
	stas, err := sta.SearchStkAdjustByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = stas
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditStkAdjust(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	sta := &model.StkAdjust{}
	err := c.BindJSON(sta)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = sta.InsertAndEditStkAdjust(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = sta
		c.JSON(http.StatusOK, rs)
	}

}
