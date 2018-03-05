package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/cors"
	s "github.com/loukmho/bcaccount_api/ctrl/salemodule"
	b "github.com/loukmho/bcaccount_api/ctrl/buymodule"
)

func main(){
	r := gin.New()
	r.Use(cors.Default())

	r.GET("/apinvoice", b.SearchApInvoiceByDocNo)

	r.GET("/arinvoice", s.SearchArInvoiceByDocNo)
	r.GET("/ardeposit", s.SearchArDepositByDocNo)
	r.GET("/ardeposits", s.SearchArDepositByKeyword)
	r.POST("/ardeposit", s.InsertArDespoit)
	r.PUT("/ardeposit", s.UpdateArDespoit)

	r.Run(":8002")
}
