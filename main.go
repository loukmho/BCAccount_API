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
	r.GET("/apinvoices", b.SearchApInvoiceByKeyword)
	r.POST("/apinvoice", b.InsertAndEditApinvoice)

	r.GET("/stkrefund", b.SearchStkRefundByDocNo)
	r.GET("/stkrefunds", b.SearchStkRefundByKeyword)
	r.POST("/stkrefund", b.InsertAndEditStkRefund)

	r.GET("/debitnote2", b.SearchDebitNote2ByDocNo)
	r.GET("/debitnote2s", b.SearchDebitNote2ByKeyword)

	r.GET("/customer", c.SearchCustomerByCode)
	r.GET("/customers", c.SearchCustomerByKeyword)
	r.POST("/customer", c.InsertAndEditCustomer)

	r.GET("/arinvoice", s.SearchArInvoiceByDocNo)
	r.GET("/arinvoices", s.SearchArInvoiceByKeyword)
	r.POST("/arinvoice", s.InsertAndEditArInvoice)

	r.GET("/creditnote", s.SearchCreditNoteByDocNo)
	r.GET("/creditnotes", s.SearchCreditNoteByKeyword)
	r.POST("/creditnote", s.InsertAndEditCreditNote)

	r.GET("/debitnote", s.SearchDebitNoteByDocNo)
	r.GET("/debitnotes", s.SearchDebitNoteByKeyword)
	r.POST("/debitnote", s.InsertAndEditDebitNote)

	r.GET("/ardeposit", s.SearchArDepositByDocNo)
	r.GET("/ardeposits", s.SearchArDepositByKeyword)
	r.POST("/ardeposit", s.InsertAndEditArDespoit)

	r.GET("/ardepositspecial", s.SearchArDepositSpecialByDocNo)
	r.GET("/ardepositspecials", s.SearchArDepositSpecialByKeyword)
	r.POST("/ardepositspecial", s.InsertAndEditArDespoitSpecial)

	r.GET("/item", i.SearchItemByCode)
	r.GET("/items", i.SearchItemByKeyword)
	r.POST("/item", i.InsertAndEditItem)

	r.Run(":8002")
}
