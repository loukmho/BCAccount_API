package chqmodule

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	ct "github.com/loukmho/BCAccount_API/ctrl"
	"github.com/loukmho/BCAccount_API/model/chqmodule"
)

func SearchChqInPassByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	cip := new(model.ChqInPass)
	err := cip.SearchChqInPassByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cip
		c.JSON(http.StatusOK, rs)
	}
}

func SearchChqInPassByKeyword(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	cip := new(model.ChqInPass)
	cips, err := cip.SearchChqInPassByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cips
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditChqInPass(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	cip := &model.ChqInPass{}
	err := c.BindJSON(cip)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = cip.InsertAndEditChqInPass(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cip
		c.JSON(http.StatusOK, rs)
	}

}
