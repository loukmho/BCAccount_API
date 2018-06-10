package chqmodule

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	ct "github.com/loukmho/BCAccount_API/ctrl"
	"github.com/loukmho/BCAccount_API/model/chqmodule"
)

func SearchChqOutCancelByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	coc := new(model.ChqOutCancel)
	err := coc.SearchChqOutCancelByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = coc
		c.JSON(http.StatusOK, rs)
	}
}

func SearchChqOutCancelByKeyword(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	coc := new(model.ChqOutCancel)
	cocs, err := coc.SearchChqOutCancelByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cocs
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditChqOutCancel(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	coc := &model.ChqOutCancel{}
	err := c.BindJSON(coc)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = coc.InsertAndEditChqOutCancel(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = coc
		c.JSON(http.StatusOK, rs)
	}

}
