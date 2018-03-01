package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/loukmho/bcaccount_api/model/salemodule"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"net/http"
	"fmt"
	"github.com/jmoiron/sqlx"
)

var dbc *sqlx.DB
func SearchArDepositByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	dbc = ct.Dbc
	docno := c.Request.URL.Query().Get("docno")
	dp := new(model.ArDeposit)
	err := dp.SearchArDepositByDocNo(dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = dp
		c.JSON(http.StatusOK, rs)
	}
}

func SearchArDepositByKeyword(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	dbc = ct.Dbc
	dp := new(model.ArDeposit)
	dps, err := dp.SearchArDepositByKeyword(dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = dps
		c.JSON(http.StatusOK, rs)
	}

}

func InsertArDespoit(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	dbc = ct.Dbc
	dp := &model.ArDeposit{}
	err := c.BindJSON(dp)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = dp.InsertArDeposit(dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = dp
		c.JSON(http.StatusOK, rs)
	}

}

func UpdateArDespoit(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	dbc = ct.Dbc
	dp := &model.ArDeposit{}
	err := c.BindJSON(dp)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = dp.UpdateArDeposit(dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = dp
		c.JSON(http.StatusOK, rs)
	}

}
