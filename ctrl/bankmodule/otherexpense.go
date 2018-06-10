package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	ct "github.com/loukmho/BCAccount_API/ctrl"
	"github.com/loukmho/BCAccount_API/model/bankmodule"
)

func SearchOtherExpenseByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	ote := new(model.OtherExpense)
	err := ote.SearchOtherExpenseByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = ote
		c.JSON(http.StatusOK, rs)
	}
}

func SearchOtherExpenseByKeyword(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	ote := new(model.OtherExpense)
	otes, err := ote.SearchOtherExpenseByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = otes
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditOtherExpense(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	ote := &model.OtherExpense{}
	err := c.BindJSON(ote)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = ote.InsertAndEditOtherExpense(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = ote
		c.JSON(http.StatusOK, rs)
	}

}
