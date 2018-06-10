package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	m "github.com/loukmho/BCAccount_API/model"
)

type Receipt3 struct {
	DocNo           string      `json:"doc_no" db:"DocNo"`
	TaxNo           string      `json:"tax_no" db:"TaxNo"`
	DocDate         string      `json:"doc_date" db:"DocDate"`
	ArCode          string      `json:"ar_code" db:"ArCode"`
	CreatorCode     string      `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string      `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string      `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string      `json:"last_edit_date_t" db:"LastEditDateT"`
	SaleCode        string      `json:"sale_code" db:"SaleCode"`
	DepartCode      string      `json:"depart_code" db:"DepartCode"`
	SumOfInvoice    float64     `json:"sum_of_invoice" db:"SumOfInvoice"`
	SumOfDebitNote  float64     `json:"sum_of_debit_note" db:"SumOfDebitNote"`
	SumOfCreditNote float64     `json:"sum_of_credit_note" db:"SumOfCreditNote"`
	BeforeTaxAmount float64     `json:"before_tax_amount" db:"BeforeTaxAmount"`
	TaxAmount       float64     `json:"tax_amount" db:"TaxAmount"`
	TotalAmount     float64     `json:"total_amount" db:"TotalAmount"`
	DiscountAmount  float64     `json:"discount_amount" db:"DiscountAmount"`
	NetAmount       float64     `json:"net_amount" db:"NetAmount"`
	GLFormat        string      `json:"gl_format" db:"GLFormat"`
	IsCancel        int         `json:"is_cancel" db:"IsCancel"`
	MyDescription   string      `json:"my_description" db:"MyDescription"`
	AllocateCode    string      `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string      `json:"project_code" db:"ProjectCode"`
	BillGroup       string      `json:"bill_group" db:"BillGroup"`
	IsCompleteSave  int         `json:"is_complete_save" db:"IsCompleteSave"`
	SumHomeAmount1  float64     `json:"sum_home_amount_1" db:"SumHomeAmount1"`
	SumHomeAmount2  float64     `json:"sum_home_amount_2" db:"SumHomeAmount2"`
	CurrencyCode    string      `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate    float64     `json:"exchange_rate" db:"ExchangeRate"`
	IsConfirm       int         `json:"is_confirm" db:"IsConfirm"`
	RecurName       string      `json:"recur_name" db:"RecurName"`
	ConfirmCode     string      `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string      `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string      `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string      `json:"cancel_date_time" db:"CancelDateTime"`
	UserCode        string      `json:"user_code"`
	Subs            []*rcp3Item `json:"subs"`
}

type rcp3Item struct {
	InvoiceNo     string  `json:"invoice_no" db:"InvoiceNo"`
	InvBalance    float64 `json:"inv_balance" db:"InvBalance"`
	PayAmount     float64 `json:"pay_amount" db:"PayAmount"`
	MyDescription string  `json:"my_description" db:"MyDescription"`
	IsCancel      int     `json:"is_cancel" db:"IsCancel"`
	LineNumber    int     `json:"line_number" db:"LineNumber"`
	BillType      int     `json:"bill_type" db:"BillType"`
	RefNo         string  `json:"ref_no" db:"RefNo"`
	RefType       int     `json:"ref_type" db:"RefType"`
	HomeAmount1   float64 `json:"home_amount_1" db:"HomeAmount1"`
	HomeAmount2   float64 `json:"home_amount_2" db:"HomeAmount2"`
}

func (rct3 *Receipt3) InsertAndEditReceipt3(db *sqlx.DB) error{
	var check_exist int
	var count_item int

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")
	if rct3.TaxRate == 0 {
		rct3.TaxRate = def.TaxRateDefault
	}

	for _, Subs := range rcp.Subs {
		if (Subs.InvoiceNo != "") {
			count_item = count_item + 1
		}
	}

	sqlexist := `select count(docno) as check_exist from dbo.bcreceipt1 where docno = ?`
	err := db.Get(&check_exist, sqlexist, rcp.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil
	}
	sql := `DocNo, TaxNo, DocDate, ArCode,  CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, SaleCode, DepartCode, SumOfInvoice, SumOfDebitNote, SumOfCreditNote, BeforeTaxAmount,TaxAmount, TotalAmount, DiscountAmount, NetAmount, GLFormat,  IsCancel, MyDescription, AllocateCode, ProjectCode, BillGroup,  IsCompleteSave, SumHomeAmount1,SumHomeAmount2, CurrencyCode, ExchangeRate, IsConfirm, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime`
}

func (rct3 *Receipt3) SearchReceipt3ByDocNo() error {
	//sql:=`DocNo, TaxNo, DocDate, ArCode,  CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, SaleCode, DepartCode, SumOfInvoice, SumOfDebitNote, SumOfCreditNote, BeforeTaxAmount,TaxAmount, TotalAmount, DiscountAmount, NetAmount, GLFormat,  IsCancel, MyDescription, AllocateCode, ProjectCode, BillGroup,  IsCompleteSave, SumHomeAmount1,SumHomeAmount2, CurrencyCode, ExchangeRate, IsConfirm, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime`
	//sqlsub := `DocNo, DocDate,ArCode, SaleCode, InvoiceDate, InvoiceNo, InvBalance, PayAmount, MyDescription, IsCancel, LineNumber, BillType, RefNo, RefType,HomeAmount1, HomeAmount2`
	return nil
}
