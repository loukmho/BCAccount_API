package ctrl

import (
	"github.com/gin-gonic/gin"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"github.com/loukmho/bcaccount_api/model/salemodule"
	"net/http"
	"fmt"
)


func InsertAndEditDebitNote(c *gin.Context){
	c.Keys = ct.HeaderKeys

	dbt := &model.DebitNote{}
	err := c.BindJSON(dbt)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = dbt.InsertAndEditDebitNote(ct.Dbc)
	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "success"
		rs.Data = dbt
		c.JSON(http.StatusNotFound, rs)
	}
}

func SearchDebitNoteByDocNo(c *gin.Context){
	c.Keys = ct.HeaderKeys

	docno := c.Request.URL.Query().Get("docno")

	dbt := new(model.DebitNote)
	err := dbt.SearchDebitNoteByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = dbt
		c.JSON(http.StatusOK, rs)
	}

}

func SearchDebitNoteByKeyword(c *gin.Context){
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	dbt := new(model.DebitNote)
	dbts, err := dbt.SearchDebitNoteByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = dbts
		c.JSON(http.StatusOK, rs)
	}

}
