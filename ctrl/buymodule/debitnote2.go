package ctrl

import (
	"github.com/gin-gonic/gin"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"github.com/loukmho/bcaccount_api/model/buymodule"
	"net/http"
)


func SearchDebitNote2ByDocNo(c *gin.Context){
	c.Keys = ct.HeaderKeys
	docno := c.Request.URL.Query().Get("docno")

	dbc := ct.Dbc
	dbn := new(model.DebitNote2)

	err := dbn.SearchDebitNote2ByDocNo(dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message= ""
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "true"
		rs.Data = dbn
		c.JSON(http.StatusOK, dbn)
	}
}


func SearchDebitNote2ByKeyword(c *gin.Context){
	c.Keys = ct.HeaderKeys
	keyword := c.Request.URL.Query().Get("keyword")

	dbc := ct.Dbc
	dbn := new(model.DebitNote2)

	dbns, err := dbn.SearchDebitnote2ByKeyword(dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message= ""
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "true"
		rs.Data = dbns
		c.JSON(http.StatusOK, dbns)
	}
}
