package chqmodule


import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"github.com/loukmho/bcaccount_api/model/chqmodule"
)

func SearchChqInDeposit2ByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	cid := new(model.ChqInDeposit2)
	err := cid.SearchChqInDeposit2ByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cid
		c.JSON(http.StatusOK, rs)
	}
}

func SearchChqInDeposit2ByKeyword(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	cid := new(model.ChqInDeposit2)
	cids, err := cid.SearchChqInDeposit2ByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cids
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditChqInDeposit2(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	cid := &model.ChqInDeposit2{}
	err := c.BindJSON(cid)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = cid.InsertAndEditChqInDeposit2(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = cid
		c.JSON(http.StatusOK, rs)
	}

}
