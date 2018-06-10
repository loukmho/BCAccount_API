package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	ct "github.com/loukmho/BCAccount_API/ctrl"
	"github.com/loukmho/BCAccount_API/model/bankmodule"
)

func SearchCashBankOutByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	cbo := new(model.CashBankOut)
	err := cbo.SearchCashBankOutByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cbo
		c.JSON(http.StatusOK, rs)
	}
}

func SearchCashBankOutByKeyword(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	cbo := new(model.CashBankOut)
	cbos, err := cbo.SearchCashBankOutByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cbos
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditCashBankOut(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	cbo := &model.CashBankOut{}
	err := c.BindJSON(cbo)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = cbo.InsertAndEditCashBankOut(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cbo
		c.JSON(http.StatusOK, rs)
	}

}
