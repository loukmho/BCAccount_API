package ctrl

import (
	"github.com/gin-gonic/gin"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"github.com/loukmho/bcaccount_api/model/custmodule"
	"net/http"
	"fmt"
)

func SearchCustomerByCode(c *gin.Context){
	c.Keys = ct.HeaderKeys

	code := c.Request.URL.Query().Get("code")

	ctm := new(model.Customer)
	err := ctm.SearchCustomerByCode(ct.Dbc, code)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		rs.Status = "success"
		rs.Data = ctm
		c.JSON(http.StatusOK ,rs)
	}
}

func SearchCustomerByKeyword(c *gin.Context){
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	ctm := new(model.Customer)
	ctms, err := ctm.SearchCustomerByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		rs.Status = "success"
		rs.Data = ctms
		c.JSON(http.StatusOK ,rs)
	}
}

func SaveUpdateCustomer(c *gin.Context){
	c.Keys = ct.HeaderKeys

	ctm := &model.Customer{}
	err := c.BindJSON(ctm)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = ctm.SaveCustomer(ct.Dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		rs.Status = "success"
		rs.Data = ctm
		c.JSON(http.StatusOK ,rs)
	}
}
