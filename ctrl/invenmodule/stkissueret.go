package ctrl

import (
	"github.com/loukmho/bcaccount_api/model/invenmodule"
	"fmt"
	"net/http"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"github.com/gin-gonic/gin"
)

func SearchStkIssueRettByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	sir := new(model.StkIssueRet)
	err := sir.SearchStkIssueRetByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = sir
		c.JSON(http.StatusOK, rs)
	}
}

func SearchStkIssueRetByKeyword(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	sir := new(model.StkIssueRet)
	sirs, err := sir.SearchStkIssueRetByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = sirs
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditStkIssueRet(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	sir := &model.StkIssueRet{}
	err := c.BindJSON(sir)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = sir.InsertAndEditStkIssueRet(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = sir
		c.JSON(http.StatusOK, rs)
	}

}
