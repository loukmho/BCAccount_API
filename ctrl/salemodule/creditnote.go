package ctrl

import (
	"github.com/gin-gonic/gin"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"github.com/loukmho/bcaccount_api/model/salemodule"
	"net/http"
	"fmt"
)

func InsertAndEditCreditNote(c *gin.Context){
	c.Keys = ct.HeaderKeys

	crd := &model.CreditNote{}
	err := c.BindJSON(crd)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = crd.InsertAndEditCreditNote(ct.Dbc)
	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "success"
		rs.Data = crd
		c.JSON(http.StatusNotFound, rs)
	}
}

func SearchCreditNoteByDocNo(c *gin.Context){
	c.Keys = ct.HeaderKeys

	docno := c.Request.URL.Query().Get("docno")

	crd := new(model.CreditNote)
	err := crd.SearchCreditNoteByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = crd
		c.JSON(http.StatusOK, rs)
	}

}

func SearchCreditNoteByKeyword(c *gin.Context){
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	crd := new(model.CreditNote)
	crds, err := crd.SearchCreditNoteByKeyword(ct.Dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = crds
		c.JSON(http.StatusOK, rs)
	}

}
