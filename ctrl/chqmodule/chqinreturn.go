package chqmodule

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	ct "github.com/loukmho/BCAccount_API/ctrl"
	"github.com/loukmho/BCAccount_API/model/chqmodule"
)

func SearchChqInReturnByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	cir := new(model.ChqInReturn)
	err := cir.SearchChqInReturnByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cir
		c.JSON(http.StatusOK, rs)
	}
}

func SearchChqInReturnByKeyword(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	cir := new(model.ChqInReturn)
	cirs, err := cir.SearchChqInReturnByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cirs
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditChqInReturn(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	cir := &model.ChqInReturn{}
	err := c.BindJSON(cir)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = cir.InsertAndEditChqInReturn(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cir
		c.JSON(http.StatusOK, rs)
	}

}