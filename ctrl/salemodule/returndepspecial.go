package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/loukmho/bcaccount_api/model/salemodule"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"fmt"
	"net/http"
)

func InsertAndEditReturnDepSpecial(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	rds := &model.ReturnDepSpecial{}
	err := c.BindJSON(rds)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = rds.InsertAndEditReturnDepSpecial(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = rds
		c.JSON(http.StatusOK, rs)
	}

}

func SearchReturnDepSpecialByDocNo(c *gin.Context){
	c.Keys = ct.HeaderKeys
	docno := c.Request.URL.Query().Get("docno")

	rds := new(model.ReturnDepSpecial)
	err := rds.SearchReturnDepSpecialByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = rds
		c.JSON(http.StatusOK, rs)
	}

}

func SearchReturnDepSpecialByKeyword(c *gin.Context){
	c.Keys = ct.HeaderKeys
	keyword := c.Request.URL.Query().Get("keyword")

	rds := new(model.ReturnDepSpecial)
	rdsList, err := rds.SearchReturnDepSpecialByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = rdsList
		c.JSON(http.StatusOK, rs)
	}

}
