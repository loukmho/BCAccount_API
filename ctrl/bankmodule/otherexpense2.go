package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	ct "github.com/loukmho/BCAccount_API/ctrl"
	"github.com/loukmho/BCAccount_API/model/bankmodule"
)

func SearchOtherExpense2ByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	ote2 := new(model.OtherExpense2)
	err := ote2.SearchOtherExpense2ByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = ote2
		c.JSON(http.StatusOK, rs)
	}
}

func SearchOtherExpense2ByKeyword(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	ote2 := new(model.OtherExpense2)
	ote2s, err := ote2.SearchOtherExpense2ByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = ote2s
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditOtherExpense2(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	ote2 := &model.OtherExpense2{}
	err := c.BindJSON(ote2)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = ote2.InsertAndEditOtherExpense2(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = ote2
		c.JSON(http.StatusOK, rs)
	}

}
