package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/cors"
	c "github.com/loukmho/bcaccount_api/ctrl/custmodule"
	s "github.com/loukmho/bcaccount_api/ctrl/salemodule"
	b "github.com/loukmho/bcaccount_api/ctrl/buymodule"
	i "github.com/loukmho/bcaccount_api/ctrl/invenmodule"
	//y "github.com/betacraft/yaag/gin/v1"
	  "github.com/betacraft/yaag/yaag"
)

//yaag.Init(&yaag.Config{On: true, DocTitle: "Core", DocPath: "apidoc.html", BaseUrls: map[string]string{"Production": "", "Staging": ""}})

func main() {
	r := gin.New()
	r.Use(cors.Default())

	yaag.Init(&yaag.Config{On: true, DocTitle: "Gin", DocPath: "apidoc.html", BaseUrls: map[string]string{"Production": "", "Staging": ""}})
	//r.Use(y.Document())

	r.GET("/apinvoice", b.SearchApInvoiceByDocNo)

	r.GET("/customer", c.SearchCustomerByCode)
	r.GET("/customers", c.SearchCustomerByKeyword)
	r.POST("/customer", c.SaveUpdateCustomer)

	r.GET("/arinvoice", s.SearchArInvoiceByDocNo)
	r.GET("/arinvoices", s.SearchArInvoiceByKeyword)
	r.POST("/arinvoice", s.InsertArinvoice)
	r.PUT("/arinvoice", s.InsertArinvoice)

	r.GET("/creditnote", s.SearchCreditNoteByDocNo)
	r.GET("/creditnotes", s.SearchCreditNoteByKeyword)

	r.GET("/ardeposit", s.SearchArDepositByDocNo)
	r.GET("/ardeposits", s.SearchArDepositByKeyword)
	r.POST("/ardeposit", s.InsertArDespoit)
	//r.PUT("/ardeposit", s.UpdateArDespoit)

	r.GET("/ardepositspecial", s.SearchArDepositSpecialByDocNo)
	r.GET("/ardepositspecials", s.SearchArDepositSpecialByKeyword)
	r.POST("/ardepositspecial", s.InsertArDespoit)
	//r.PUT("/ardepositspecial", s.UpdateArDespoit)

	r.GET("/item", i.SearchItemByCode)
	r.GET("/items", i.SearchItemByKeyword)
	r.POST("/item", i.SaveUpdateItem)

	r.Run(":8002")
}
