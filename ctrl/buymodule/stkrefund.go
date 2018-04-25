package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/loukmho/bcaccount_api/model/buymodule"
	ct "github.com/loukmho/bcaccount_api/ctrl"
	"net/http"
)

func SearchStkRefundByDocNo(c *gin.Context){
	c.Keys = ct.HeaderKeys

	docno := c.Request.URL.Query().Get("docno")
	srf := new(model.StkRefund)
	dbc := ct.Dbc
	err := srf.SearchStkRefundByDocNo(dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "success"
		rs.Data = srf
		c.JSON(http.StatusOK, rs)
	}

}

func SearchStkRefundByKeyword(c *gin.Context){
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")
	srf := new(model.StkRefund)
	dbc := ct.Dbc
	srfs, err := srf.SearchStkRefundByKeyword(dbc, keyword)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "success"
		rs.Data = srfs
		c.JSON(http.StatusOK, rs)
	}

}
