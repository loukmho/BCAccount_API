package ctrl

import (
	"github.com/gin-gonic/gin"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"github.com/loukmho/bcaccount_api/model/salemodule"
	"net/http"
	"fmt"
)


func InsertAndEditArDepoitSpecial(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	dps := &model.ArDepositSpecial{}
	err := c.BindJSON(dps)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = dps.InsertAndEditArDepositSpecial(ct.Dbc)

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

func SearchArDepositSpecialByDocNo(c *gin.Context){
	c.Keys = ct.HeaderKeys
	docno := c.Request.URL.Query().Get("docno")

	dps := new(model.ArDepositSpecial)
	err := dps.SearchArDepositSpecialByDocNo(ct.Dbc, docno)

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

func SearchArDepositSpecialByKeyword(c *gin.Context){
	c.Keys = ct.HeaderKeys
	keyword := c.Request.URL.Query().Get("keyword")

	dps := new(model.ArDepositSpecial)
	dpsList, err := dps.SearchArDepositSpecialByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = dpsList
		c.JSON(http.StatusOK, rs)
	}

}
