package ctrl

import (
	"github.com/loukmho/bcaccount_api/model/custmodule"
	"fmt"
	"net/http"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"github.com/gin-gonic/gin"
)

func SearchReceipt1ByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	rcp := new(model.Receipt1)
	err := rcp.SearchReceipt1ByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = rcp
		c.JSON(http.StatusOK, rs)
	}
}

func SearchReceipt1ByKeyword(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	rcp := new(model.Receipt1)
	rcps, err := rcp.SearchReceipt1ByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = rcps
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditReceipt1(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	rcp := &model.Receipt1{}
	err := c.BindJSON(rcp)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = rcp.InsertAndEditReceipt1(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = rcp
		c.JSON(http.StatusOK, rs)
	}

}
