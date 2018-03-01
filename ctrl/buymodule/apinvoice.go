package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/loukmho/bcaccount_api/model/buymodule"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"github.com/jmoiron/sqlx"
)

var dbc *sqlx.DB
func SearchApInvoiceByDocNo(c *gin.Context){
	c.Keys = ct.HeaderKeys

	//docno := c.Request.URL.Query().Get("docno")
	dbc = ct.Dbc
	ai := new(model.ApInvoice)
	err := ai.SearchApInvoiceByDocNo(dbc)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "success"
		rs.Data = ai
		c.JSON(http.StatusOK, rs)
	}
}
