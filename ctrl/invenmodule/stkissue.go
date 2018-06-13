package ctrl

import (
	"github.com/loukmho/bcaccount_api/model/invenmodule"
	"fmt"
	"net/http"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"github.com/gin-gonic/gin"
)

func SearchStkIssuetByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	iss := new(model.StkIssue)
	err := iss.SearchStkIssueByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = iss
		c.JSON(http.StatusOK, rs)
	}
}

func SearchStkIssueByKeyword(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	iss := new(model.StkIssue)
	isss, err := iss.SearchStkIssueByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = isss
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditStkIssue(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	iss := &model.StkIssue{}
	err := c.BindJSON(iss)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = iss.InsertAndEditStkIssue(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = iss
		c.JSON(http.StatusOK, rs)
	}

}
