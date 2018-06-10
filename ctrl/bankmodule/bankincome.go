package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	ct "github.com/loukmho/BCAccount_API/ctrl"
	"github.com/loukmho/BCAccount_API/model/bankmodule"
)

func SearchBankInComeByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	bic := new(model.BankInCome)
	err := bic.SearchBankInComeByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = bic
		c.JSON(http.StatusOK, rs)
	}
}

func SearchBankInComeByKeyword(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	bic := new(model.BankInCome)
	bics, err := bic.SearchBankInComeByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = bics
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditBankInCome(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	bic := &model.BankInCome{}
	err := c.BindJSON(bic)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = bic.InsertAndEditBankInCome(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = bic
		c.JSON(http.StatusOK, rs)
	}

}
