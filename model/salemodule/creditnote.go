package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type CreditNote struct {
	DocNo           string    `json:"doc_no" db:"DocNo"`
	TaxNo           string    `json:"tax_no" db:"TaxNo"`
	TaxDate         string    `json:"tax_date" db:"TaxDate"`
	CreatorCode     string    `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string    `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string    `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string    `json:"last_edit_date_t" db:"LastEditDateT"`
	DocDate         string    `json:"doc_date" db:"DocDate"`
	DueDate         string    `json:"due_date" db:"DueDate"`
	TaxType         int       `json:"tax_type" db:"TaxType"`
	CrdCustomer
	CrdSaleMan
	DepartCode      string    `json:"depart_code" db:"DepartCode"`
	CashierCode     string    `json:"cashier_code" db:"CashierCode"`
	TaxRate         float64   `json:"tax_rate" db:"TaxRate"`
	IsConfirm       int       `json:"is_confirm" db:"IsConfirm"`
	MyDescription   string    `json:"my_description" db:"MyDescription"`
	SumOfItemAmount float64   `json:"sum_of_item_amount" db:"SumOfItemAmount"`
	SumOldAmount    float64   `json:"sum_old_amount" db:"SumOldAmount"`
	SumTrueAmount   float64   `json:"sum_true_amount" db:"SumTrueAmount"`
	SumofDiffAmount float64   `json:"sumof_diff_amount" db:"SumofDiffAmount"`
	DiscountWord    string    `json:"discount_word" db:"DiscountWord"`
	DiscountAmount  float64   `json:"discount_amount" db:"DiscountAmount"`
	SumofBeforeTax  float64   `json:"sumof_before_tax" db:"SumofBeforeTax"`
	SumOfTaxAmount  float64   `json:"sum_of_tax_amount" db:"SumOfTaxAmount"`
	SumOfTotalTax   float64   `json:"sum_of_total_tax" db:"SumOfTotalTax"`
	SumOfZeroTax    float64   `json:"sum_of_zero_tax" db:"SumOfZeroTax"`
	SumOfExceptTax  float64   `json:"sum_of_except_tax" db:"SumOfExceptTax"`
	SumOfWTax       float64   `json:"sum_of_w_tax" db:"SumOfWTax"`
	NetDebtAmount   float64   `json:"net_debt_amount" db:"NetDebtAmount"`
	BillBalance     float64   `json:"bill_balance" db:"BillBalance"`
	CurrencyCode    string    `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate    float64   `json:"exchange_rate" db:"ExchangeRate"`
	GLFormat        string    `json:"gl_format" db:"GLFormat"`
	IsCancel        int       `json:"is_cancel" db:"IsCancel"`
	IsCompleteSave  int       `json:"is_complete_save" db:"IsCompleteSave"`
	ReturnMoney     int       `json:"return_money" db:"ReturnMoney"`
	ReturnStatus    int       `json:"return_status" db:"ReturnStatus"`
	ReturnCash      int       `json:"return_cash" db:"ReturnCash"`
	OtherIncome     float64   `json:"other_income" db:"OtherIncome"`
	OtherExpense    float64   `json:"other_expense" db:"OtherExpense"`
	ExcessAmount1   float64   `json:"excess_amount_1" db:"ExcessAmount1"`
	ExcessAmount2   float64   `json:"excess_amount_2" db:"ExcessAmount2"`
	SumCashAmount   float64   `json:"sum_cash_amount" db:"SumCashAmount"`
	SumChqAmount    float64   `json:"sum_chq_amount" db:"SumChqAmount"`
	SumCreditAmount float64   `json:"sum_credit_amount" db:"SumCreditAmount"`
	SumBankAmount   float64   `json:"sum_bank_amount" db:"SumBankAmount"`
	ChargeAmount    float64   `json:"charge_amount" db:"ChargeAmount"`
	ChangeAmount    float64   `json:"change_amount" db:"ChangeAmount"`
	CauseType       int       `json:"cause_type" db:"CauseType"`
	PayBillStatus   int       `json:"pay_bill_status" db:"PayBillStatus"`
	IsCNDeposit     int       `json:"is_cn_deposit" db:"IsCNDeposit"`
	IsPos           int       `json:"is_pos" db:"IsPos"`
	PosDocNo        string    `json:"pos_doc_no" db:"PosDocNo"`
	CauseCode       string    `json:"cause_code" db:"CauseCode"`
	AllocateCode    string    `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string    `json:"project_code" db:"ProjectCode"`
	BillGroup       string    `json:"bill_group" db:"BillGroup"`
	RecurName       string    `json:"recur_name" db:"RecurName"`
	ConfirmCode     string    `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string    `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string    `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string    `json:"cancel_date_time" db:"CancelDateTime"`
	PayBillAmount   float64   `json:"pay_bill_amount" db:"PayBillAmount"`
	BillTemporary   float64   `json:"bill_temporary" db:"BillTemporary"`
	Item            []CrdItem `json:"item"`
}

type CrdCustomer struct {
	ArCode string `json:"ar_code" db:"ArCode"`
	ArName string `json:"ar_name" db:"ArName"`
}

type CrdSaleMan struct {
	SaleCode string `json:"sale_code" db:"SaleCode"`
	SaleName string `json:"sale_name" db:"SaleName"`
}

type CrdItem struct {
	MyType         int     `json:"my_type" db:"MyType"`
	ItemCode       string  `json:"item_code" db:"ItemCode"`
	MyDescription  string  `json:"my_description" db:"MyDescription"`
	ItemName       string  `json:"item_name" db:"ItemName"`
	WHCode         string  `json:"wh_code" db:"WHCode"`
	ShelfCode      string  `json:"shelf_code" db:"ShelfCode"`
	DiscQty        float64 `json:"disc_qty" db:"DiscQty"`
	TempQty        float64 `json:"temp_qty" db:"TempQty"`
	BillQty        float64 `json:"bill_qty" db:"BillQty"`
	Price          float64 `json:"price" db:"Price"`
	DiscountWord   string  `json:"discount_word" db:"DiscountWord"`
	DiscountAmount float64 `json:"discount_amount" db:"DiscountAmount"`
	Amount         float64 `json:"amount" db:"Amount"`
	NetAmount      float64 `json:"net_amount" db:"NetAmount"`
	HomeAmount     float64 `json:"home_amount" db:"HomeAmount"`
	SumOfCost      float64 `json:"sum_of_cost" db:"SumOfCost"`
	UnitCode       string  `json:"unit_code" db:"UnitCode"`
	InvoiceNo      string  `json:"invoice_no" db:"InvoiceNo"`
	IsPos          int     `json:"is_pos" db:"IsPos"`
	IsCancel       int     `json:"is_cancel" db:"IsCancel"`
	LineNumber     int     `json:"line_number" db:"LineNumber"`
	RefLineNumber  int     `json:"ref_line_number" db:"RefLineNumber"`
	BarCode        string  `json:"bar_code" db:"BarCode"`
	AVERAGECOST    float64 `json:"averagecost" db:"AVERAGECOST"`
	LotNumber      string  `json:"lot_number" db:"LotNumber"`
	PackingRate1   float64 `json:"packing_rate_1" db:"PackingRate1"`
	PackingRate2   float64 `json:"packing_rate_2" db:"PackingRate2"`
}


func (crd *CreditNote) SearchCreditNoteByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select a.DocNo,isnull(a.TaxNo,'') as TaxNo,isnull(a.TaxDate,'') as TaxDate,isnull(a.CreatorCode,'') as CreatorCode,isnull(a.CreateDateTime,'') as CreateDateTime,isnull(a.LastEditorCode,'') as LastEditorCode,isnull(a.LastEditDateT,'') as LastEditDateT,a.DocDate,isnull(a.DueDate,'') as DueDate,a.TaxType,a.ArCode,isnull(b.name1,'') as ArName,isnull(a.DepartCode,'') as DepartCode,isnull(a.SaleCode,'') as SaleCode,isnull(c.name,'') as SaleName,isnull(a.CashierCode,'') as CashierCode,a.TaxRate,a.IsConfirm,isnull(a.MyDescription,'') as MyDescription,a.SumOfItemAmount,a.SumOldAmount,a.SumTrueAmount,a.SumofDiffAmount,isnull(a.DiscountWord,'') as DiscountWord,a.DiscountAmount,a.SumofBeforeTax,a.SumOfTaxAmount,a.SumOfTotalTax,a.SumOfZeroTax,a.SumOfExceptTax,a.SumOfWTax,a.NetDebtAmount,a.BillBalance,isnull(a.CurrencyCode,'') as CurrencyCode,a.ExchangeRate,isnull(a.GLFormat,'') as GLFormat,a.IsCancel,a.IsCompleteSave,a.ReturnMoney,a.ReturnStatus,a.ReturnCash,a.OtherIncome,a.OtherExpense,a.ExcessAmount1,a.ExcessAmount2,a.SumCashAmount,a.SumChqAmount,a.SumCreditAmount,a.SumBankAmount,a.ChargeAmount,a.ChangeAmount,a.CauseType,a.PayBillStatus,a.IsCNDeposit,a.IsPos,isnull(a.PosDocNo,'') as PosDocNo,isnull(a.CauseCode,'') as CauseCode,isnull(a.AllocateCode,'') as AllocateCode,isnull(a.ProjectCode,'') as ProjectCode,isnull(a.BillGroup,'') as BillGroup,isnull(a.RecurName,'') as RecurName,isnull(a.ConfirmCode,'') as ConfirmCode,isnull(a.ConfirmDateTime,'') as ConfirmDateTime,isnull(a.CancelCode,'') as CancelCode,isnull(a.CancelDateTime,'') as CancelDateTime,a.PayBillAmount,a.BillTemporary from dbo.BCCreditNote a With (Nolock) left join dbo.bcar b with (nolock) on a.arcode = b.code left join dbo.bcsale c with (nolock) on a.salecode = c.code where a.docno = ?`
	err := db.Get(crd, sql, docno)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//sqlsub := `MyType, DocNo, TaxNo, TaxType, ItemCode, DocDate, ArCode, DepartCode, SaleCode, CashierCode, MyDescription, ItemName, WHCode, ShelfCode, DiscQty, TempQty, BillQty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost, UnitCode, InvoiceNo, ItemType, ExceptTax, IsPos, IsCancel, LineNumber, RefLineNumber, BarCode,AVERAGECOST, LotNumber, PackingRate1, PackingRate2`
	sqlsub := `set dateformat dmy     select a.MyType,a.ItemCode,isnull(a.MyDescription,'') as MyDescription,isnull(a.ItemName,'') as ItemName,isnull(a.WHCode,'') as WHCode,isnull(a.ShelfCode,'') as ShelfCode,a.DiscQty,a.TempQty,a.BillQty,a.Price,isnull(a.DiscountWord,'') as DiscountWord,a.DiscountAmount,a.Amount,a.NetAmount,a.HomeAmount,a.SumOfCost,isnull(a.UnitCode,'') as UnitCode,isnull(a.InvoiceNo,'') as InvoiceNo,a.IsPos,a.IsCancel,a.LineNumber,a.RefLineNumber,isnull(a.BarCode,'') as BarCode,a.AVERAGECOST,isnull(a.LotNumber,'') as LotNumber,a.PackingRate1,a.PackingRate2 from dbo.BCCreditNoteSub a with (nolock) where a.docno = ?`
	err = db.Select(&crd.Item, sqlsub, docno)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return nil
}
