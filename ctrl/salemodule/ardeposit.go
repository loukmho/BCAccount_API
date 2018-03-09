package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/loukmho/bcaccount_api/model/salemodule"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"net/http"
	"fmt"
)

func SearchArDepositByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	dp := new(model.ArDeposit)
	err := dp.SearchArDepositByDocNo(ct.Dbc, docno)

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

	dp := new(model.ArDeposit)
	dps, err := dp.SearchArDepositByKeyword(ct.Dbc, keyword)

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

	dp := &model.ArDeposit{}
	err := c.BindJSON(dp)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = dp.InsertArDeposit(ct.Dbc)

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

	dp := &model.ArDeposit{}
	err := c.BindJSON(dp)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = dp.UpdateArDeposit(ct.Dbc)

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
