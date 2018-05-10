package ctrl

import (
	"github.com/gin-gonic/gin"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"github.com/loukmho/bcaccount_api/model/bankmodule"
	"fmt"
	"net/http"
)

func InsertAndEditBankTransfer(c *gin.Context){
	c.Keys = ct.HeaderKeys
	btf := &model.BankTransfer{}
	err := c.BindJSON(btf)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = btf.InsertAndEditBankTransfer(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = btf
		c.JSON(http.StatusOK, rs)
	}
}

func SearchBankTransferByDocNo (c *gin.Context){
	c.Keys = ct.HeaderKeys

	docno := c.Request.URL.Query().Get("docno")

	btf := new(model.BankTransfer)
	err := btf.SearchBankTransferByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = btf
		c.JSON(http.StatusOK, rs)
	}
}


func SearchBankTransferByKeyword (c *gin.Context){
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	btf := new(model.BankTransfer)
	btfs, err := btf.SearchBankTransferByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = btfs
		c.JSON(http.StatusOK, rs)
	}
}

