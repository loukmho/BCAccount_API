package ctrl

import(
	"github.com/gin-gonic/gin"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"github.com/loukmho/bcaccount_api/model/salemodule"
	"net/http"
)


func SearchArInvoiceByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	docno := c.Request.URL.Query().Get("docno")
	inv := new(model.ArInvoice)
	err := inv.SearchArInvoiceByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		rs.Status = "success"
		rs.Data = inv
		c.JSON(http.StatusOK, rs)
	}
}
