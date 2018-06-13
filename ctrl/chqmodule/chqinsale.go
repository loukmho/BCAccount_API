package chqmodule

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"github.com/loukmho/bcaccount_api/model/chqmodule"
)

func SearchChqInSaleByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	cis := new(model.ChqInSale)
	err := cis.SearchChqInSaleByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cis
		c.JSON(http.StatusOK, rs)
	}
}

func SearchChqInSaleByKeyword(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	cis := new(model.ChqInSale)
	ciss, err := cis.SearchChqInSaleByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = ciss
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditChqInSale(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	cis := &model.ChqInSale{}
	err := c.BindJSON(cis)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = cis.InsertAndEditChqInSale(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cis
		c.JSON(http.StatusOK, rs)
	}

}
