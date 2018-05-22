package ctrl

import (
	"github.com/loukmho/bcaccount_api/model/custmodule"
	"fmt"
	"net/http"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"github.com/gin-gonic/gin"
)

func SearchArOtherDebtByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	otd := new(model.ArOtherDebt)
	err := otd.SearchArOtherDebtByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = otd
		c.JSON(http.StatusOK, rs)
	}
}

func SearchArDepositByKeyword(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	otd := new(model.ArOtherDebt)
	otds, err := otd.SearchArOtherDebtByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = otds
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditArOtherDebt(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	otd := &model.ArOtherDebt{}
	err := c.BindJSON(otd)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = otd.InsertAndEditArOtherDebt(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = otd
		c.JSON(http.StatusOK, rs)
	}

}
