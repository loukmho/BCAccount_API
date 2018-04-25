package model

import (
	"math"
	"os"
	"fmt"
	"encoding/json"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func CalcTaxDoc(taxtype int, taxrate float64, totalamount float64) (beforetaxamount float64, taxamount float64) {
	fmt.Println("taxtype,taxrate,total", taxtype, taxrate, totalamount)

	switch taxtype {
	case 0:
		taxamount = ToFixed(((totalamount*(100+float64(taxrate)))/(100))-totalamount, 2)
		beforetaxamount = ToFixed(totalamount-taxamount, 2)
	case 1:
		taxamount = ToFixed(totalamount-((totalamount*100)/(100+float64(taxrate))), 2)
		beforetaxamount = ToFixed(totalamount-taxamount, 2)
	case 2:
		beforetaxamount = ToFixed(totalamount, 2)
		taxamount = 0
	}

	fmt.Println("Before,Tax", beforetaxamount, taxamount)

	return beforetaxamount, taxamount
}

func CalcTaxItem(taxtype int, taxrate float64, afterdiscountamount float64) (beforetaxamount float64, taxamount float64, totalamount float64) {
	switch taxtype {
	case 0:
		beforetaxamount = ToFixed(afterdiscountamount, 2)
		taxamount = ToFixed(((afterdiscountamount*(100+float64(taxrate)))/(100))-afterdiscountamount, 2)
		totalamount = ToFixed(beforetaxamount+taxamount, 2)
	case 1:
		taxamount = ToFixed(afterdiscountamount-((afterdiscountamount*100)/(100+float64(taxrate))), 2)
		beforetaxamount = ToFixed(afterdiscountamount-taxamount, 2)
		totalamount = ToFixed(afterdiscountamount, 2)
	case 2:
		beforetaxamount = ToFixed(afterdiscountamount, 2)
		taxamount = 0
		totalamount = ToFixed(afterdiscountamount, 2)
	}

	fmt.Println("taxtype,taxrate,beforetaxamount,taxamount,totalamount", taxtype, taxrate, beforetaxamount, taxamount, totalamount)

	return beforetaxamount, taxamount, totalamount
}

func CalcTaxCredit(taxtype int, taxrate float64, totalamount float64) (beforetaxamount float64, taxamount float64) {
	fmt.Println("taxtype,taxrate,total", taxtype, taxrate, totalamount)

	switch taxtype {
	case 0:
		beforetaxamount = ToFixed(((totalamount * (100)) / (100 + float64(taxrate))), 2)
		taxamount = ToFixed(totalamount-beforetaxamount, 2)
	case 1:
		taxamount = ToFixed(totalamount-((totalamount*100)/(100+float64(taxrate))), 2)
		beforetaxamount = ToFixed(totalamount-taxamount, 2)
	case 2:
		beforetaxamount = ToFixed(totalamount, 2)
		taxamount = 0
	}

	fmt.Println("Before,Tax", beforetaxamount, taxamount)

	return beforetaxamount, taxamount
}

type Default struct {
	TaxRateDefault      float64 `json:"tax_rate_default"`
	ExchangeRateDefault float64 `json:"exchange_rate_default"`

	ApInvoiceCashGLFormat   string `json:"ap_invoice_cash_gl_format"`
	ApInvoiceCreditGLFormat string `json:"ap_invoice_credit_gl_format"`
	ApInvoiceBookCode       string `json:"ap_invoice_book_code"`
	ApInvoiceSource         int    `json:"ap_invoice_source"`
	ApInvoiceSaveFrom       int    `json:"ap_invoice_save_from"`
	ApInvoiceMyType         int    `json:"ap_invoice_my_type"`

	StkRefundGLFormat   string `json:"stk_refund_gl_format"`
	StkRefundBookCode       string `json:"StkRefund_book_code"`
	StkRefundSource         int    `json:"StkRefund_source"`
	StkRefundSaveFrom       int    `json:"StkRefund_save_from"`
	StkRefundMyType         int    `json:"StkRefund_my_type"`

	ArDepositGLFormat string `json:"ar_deposit_gl_format"`
	ArDepositBookCode string `json:"ar_deposit_book_code"`
	ArDepositSource   int    `json:"ar_deposit_source"`
	ArDepositSaveFrom int    `json:"ar_deposit_save_from"`

	ArInvoiceCashGLFormat   string `json:"ar_invoice_cash_gl_format"`
	ArInvoiceCreditGLFormat string `json:"ar_invoice_credit_gl_format"`
	ArInvoiceBookCode       string `json:"ar_invoice_book_code"`
	ArInvoiceSource         int    `json:"ar_invoice_source"`
	ArInvoiceSaveFrom       int    `json:"ar_invoice_save_from"`
	ArInvoiceMyType         int    `json:"ar_invoice_my_type"`

	CreditNoteGLFormat string `json:"credit_note_gl_format"`
	CreditNoteBookCode string `json:"credit_note_book_code"`
	CreditNoteSource   int    `json:"credit_note_source"`
	CreditNoteSaveFrom int    `json:"credit_note_save_from"`
	CreditNoteMyType   int    `json:"credit_note_my_type"`

	DebitNoteGLFormat string `json:"debit_note_gl_format"`
	DebitNoteBookCode string `json:"debit_note_book_code"`
	DebitNoteSource   int    `json:"debit_note_source"`
	DebitNoteSaveFrom int    `json:"debit_note_save_from"`
	DebitNoteMyType   int    `json:"debit_note_my_type"`

	ArDepositSpecialGLFormat string `json:"ar_deposit_special_gl_format"`
}

func LoadDefaultData(fileName string) (df Default) {

	// dbType 0 = mySql , 1 = MsSql
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Err Open file %v: Error is:", file, err)
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	//c := new(Default)
	err = decoder.Decode(&df)
	if err != nil {
		fmt.Println("error Decode Json:", err)
	}
	//dsn := ""
	////if dbType == 0 {
	////	dsn = c.DBUser + ":" + c.DBPass + "@" + c.DBHost + "/" + c.DBName + "?parseTime=true"
	////}
	////if dbType == 1 {
	////	dsn = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable", c.DBHost, c.DBUser, c.DBPass, c.DBPort, c.DBName)
	////	fmt.Println("DSN =", dsn)
	////}

	return df
}
