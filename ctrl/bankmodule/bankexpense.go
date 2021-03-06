package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	ct "github.com/loukmho/BCAccount_API/ctrl"
	"github.com/loukmho/BCAccount_API/model/bankmodule"
)

func SearchBankExpenseByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	bep := new(model.BankExpense)
	err := bep.SearchBankExpenseByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = bep
		c.JSON(http.StatusOK, rs)
	}
}

func SearchBankExpenseByKeyword(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	bep := new(model.BankExpense)
	beps, err := bep.SearchBankExpenseByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = beps
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditBankExpense(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	bep := &model.BankExpense{}
	err := c.BindJSON(bep)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = bep.InsertAndEditBankExpense(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = bep
		c.JSON(http.StatusOK, rs)
	}

}

