package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	ct "github.com/loukmho/BCAccount_API/ctrl"
	"github.com/loukmho/BCAccount_API/model/bankmodule"
)

func SearchCashBankInByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	cbi := new(model.CashBankIn)
	err := cbi.SearchCashBankInByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cbi
		c.JSON(http.StatusOK, rs)
	}
}

func SearchCashBankInByKeyword(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	cbi := new(model.CashBankIn)
	cbis, err := cbi.SearchCashBankInByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cbis
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditCashBankIn(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	cbi := &model.CashBankIn{}
	err := c.BindJSON(cbi)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = cbi.InsertAndEditCashBankIn(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cbi
		c.JSON(http.StatusOK, rs)
	}

}
