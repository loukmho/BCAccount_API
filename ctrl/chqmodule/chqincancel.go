package chqmodule

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"github.com/loukmho/bcaccount_api/model/chqmodule"
)

func SearchChqInCancelByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	cic := new(model.ChqInCancel)
	err := cic.SearchChqInCancelByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cic
		c.JSON(http.StatusOK, rs)
	}
}

func SearchChqInCancelByKeyword(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	cic := new(model.ChqInCancel)
	cics, err := cic.SearchChqInCancelByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cics
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditChqInCancel(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	cic := &model.ChqInCancel{}
	err := c.BindJSON(cic)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = cic.InsertAndEditChqInCancel(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cic
		c.JSON(http.StatusOK, rs)
	}

}
