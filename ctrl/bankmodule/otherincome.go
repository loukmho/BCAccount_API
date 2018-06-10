package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	ct "github.com/loukmho/BCAccount_API/ctrl"
	"github.com/loukmho/BCAccount_API/model/bankmodule"
)

func SearchOtherInComeByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	ote := new(model.OtherInCome)
	err := ote.SearchOtherInComeByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = ote
		c.JSON(http.StatusOK, rs)
	}
}

func SearchOtherInComeByKeyword(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	ote := new(model.OtherInCome)
	otes, err := ote.SearchOtherInComeByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = otes
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditOtherInCome(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	ote := &model.OtherInCome{}
	err := c.BindJSON(ote)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = ote.InsertAndEditOtherInCome(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = ote
		c.JSON(http.StatusOK, rs)
	}

}