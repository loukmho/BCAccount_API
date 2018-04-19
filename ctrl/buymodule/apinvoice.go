package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/loukmho/bcaccount_api/model/buymodule"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"github.com/jmoiron/sqlx"
	"fmt"
)


func InsertAndEditApinvoice(c *gin.Context){
	c.Keys = ct.HeaderKeys

	apv := &model.ApInvoice{}
	err := c.BindJSON(apv)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = apv.InsertAndEditApInvoice(ct.Dbc)
	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "success"
		rs.Data = apv
		c.JSON(http.StatusNotFound, rs)
	}
}

var dbc *sqlx.DB
func SearchApInvoiceByDocNo(c *gin.Context){
	c.Keys = ct.HeaderKeys

	docno := c.Request.URL.Query().Get("docno")
	dbc = ct.Dbc
	apv := new(model.ApInvoice)
	err := apv.SearchApInvoiceByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "success"
		rs.Data = apv
		c.JSON(http.StatusOK, rs)
	}
}

func SearchApInvoiceByKeyword(c *gin.Context){
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")
	dbc = ct.Dbc
	apv := new(model.ApInvoice)
	apvs, err := apv.SearchApInvoiceByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "success"
		rs.Data = apvs
		c.JSON(http.StatusOK, rs)
	}
}
