package ctrl

import (
	"github.com/loukmho/bcaccount_api/model/invenmodule"
	"fmt"
	"net/http"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"github.com/gin-gonic/gin"
)

func SearchFinishGoodsByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	fng := new(model.FinishGoods)
	err := fng.SearchFinishGoodsByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = fng
		c.JSON(http.StatusOK, rs)
	}
}

func SearchFinishGoodsByKeyword(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	fng := new(model.FinishGoods)
	fngs, err := fng.SearchFinishGoodsByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = fngs
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditFinishGoods(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	fng := &model.FinishGoods{}
	err := c.BindJSON(fng)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = fng.InsertAndEditFinishGoods(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = fng
		c.JSON(http.StatusOK, rs)
	}

}
