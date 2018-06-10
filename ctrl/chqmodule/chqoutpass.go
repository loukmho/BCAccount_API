package chqmodule

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	ct "github.com/loukmho/BCAccount_API/ctrl"
	"github.com/loukmho/BCAccount_API/model/chqmodule"
)

func SearchChqOutPassByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	cop := new(model.ChqOutPass)
	err := cop.SearchChqOutPassByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cop
		c.JSON(http.StatusOK, rs)
	}
}

func SearchChqOutPassByKeyword(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	cop := new(model.ChqOutPass)
	cops, err := cop.SearchChqOutPassByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cops
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditChqOutPass(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	cop := &model.ChqOutPass{}
	err := c.BindJSON(cop)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = cop.InsertAndEditChqOutPass(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cop
		c.JSON(http.StatusOK, rs)
	}

}
