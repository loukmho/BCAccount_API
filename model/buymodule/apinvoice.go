package model

import (
	"github.com/jmoiron/sqlx"
)

type ApInvoice struct {
	DocNo           string  `json:"doc_no" db:""`
	TaxNo           string  `json:"tax_no" db:""`
	DocDate         string  `json:"doc_date" db:""`
	TaxType         int     `json:"tax_type" db:""`
	ApCode          string  `json:"ap_code" db:""`
	DepartCode      string  `json:"depart_code" db:""`
	CreditDay       int     `json:"credit_day" db:""`
	DueDate         string  `json:"due_date" db:""`
	StatementDate   string  `json:"statement_date" db:""`
	TaxRate         float64 `json:"tax_rate" db:""`
	IsConfirm       int     `json:"is_confirm" db:""`
	MyDescription   string  `json:"my_description" db:""`
	BillType        int     `json:"bill_type" db:""`
	BillGroup       string  `json:"bill_group" db:""`
	ContactCode     string  `json:"contact_code" db:""`
	SumOfItemAmount float64 `json:"sum_of_item_amount" db:""`
	DiscountWord    string  `json:"discount_word" db:""`
	DiscountAmount  float64 `json:"discount_amount" db:""`
	AfterDiscount   float64 `json:"after_discount" db:""`
	BeforeTaxAmount float64 `json:"before_tax_amount" db:""`
	TaxAmount       float64 `json:"tax_amount" db:""`
	TotalAmount     float64 `json:"total_amount" db:""`
	DiscountCase    float64 `json:"discount_case" db:""`
	ExceptTaxAmount float64 `json:"except_tax_amount" db:""`
	ZeroTaxAmount   float64 `json:"zero_tax_amount" db:""`
	PettyCashAmount float64 `json:"petty_cash_amount" db:""`
	SumCashAmount   float64 `json:"sum_cash_amount" db:""`
	SumChqAmount    float64 `json:"sum_chq_amount" db:""`
	SumBankAmount   float64 `json:"sum_bank_amount" db:""`
	DepositIncTax   float64 `json:"deposit_inc_tax" db:""`
	SumOfDeposit1   float64 `json:"sum_of_deposit_1" db:""`
	SumOfDeposit2   float64 `json:"sum_of_deposit_2" db:""`
	SumOfWTax       float64 `json:"sum_of_w_tax" db:""`
	NetDebtAmount   float64 `json:"net_debt_amount" db:""`
	HomeAmount      float64 `json:"home_amount" db:""`
	OtherIncome     float64 `json:"other_income" db:""`
	OtherExpense    float64 `json:"other_expense" db:""`
	ExcessAmount1   float64 `json:"excess_amount_1" db:""`
	ExcessAmount2   float64 `json:"excess_amount_2" db:""`
	BillBalance     float64 `json:"bill_balance" db:""`
	CurrencyCode    string  `json:"currency_code" db:""`
	ExchangeRate    float64 `json:"exchange_rate" db:""`
	GLFormat        string  `json:"gl_format" db:""`
	GLStartPosting  int     `json:"gl_start_posting" db:""`
	GRBillStatus    int     `json:"gr_bill_status" db:""`
	GRIRBillStatus  int     `json:"grir_bill_status" db:""`
	IsCancel        int     `json:"is_cancel" db:""`
	IsCreditNote    int     `json:"is_credit_note" db:""`
	IsDebitNote     int     `json:"is_debit_note" db:""`
	IsCompleteSave  int     `json:"is_complete_save" db:""`
	AllocateCode    string  `json:"allocate_code" db:""`
	ProjectCode     string  `json:"project_code" db:""`
	RecurName       string  `json:"recur_name" db:""`
	ExchangeProfit  float64 `json:"exchange_profit" db:""`
	ConfirmCode     string  `json:"confirm_code" db:""`
	ConfirmDateTime string  `json:"confirm_date_time" db:""`
	CancelCode      string  `json:"cancel_code" db:""`
	CancelDateTime  string  `json:"cancel_date_time" db:""`
	PayBillAmount   float64 `json:"pay_bill_amount" db:""`
	RefDocNo        string  `json:"ref_doc_no" db:""`
	CreatorCode     string  `json:"creator_code" db:""`
	CreateDateTime  string  `json:"create_date_time" db:""`
	LastEditorCode  string  `json:"last_editor_code" db:""`
	LastEditDateT   string  `json:"last_edit_date_t" db:""`
}

type Apv struct {
	MyType         int     `json:"my_type" db:""`
	ItemCode       string  `json:"item_code" db:""`
	MyDescription  string  `json:"my_description" db:""`
	ItemName       string  `json:"item_name" db:""`
	WHCode         string  `json:"wh_code" db:""`
	ShelfCode      string  `json:"shelf_code" db:""`
	CNQty          float64 `json:"cn_qty" db:""`
	GRRemainQty    float64 `json:"gr_remain_qty" db:""`
	Qty            float64 `json:"qty" db:""`
	Price          float64 `json:"price" db:""`
	DiscountWord   string  `json:"discount_word" db:""`
	DiscountAmount float64 `json:"discount_amount" db:""`
	Amount         float64 `json:"amount" db:""`
	NetAmount      float64 `json:"net_amount" db:""`
	HomeAmount     float64 `json:"home_amount" db:""`
	BalanceAmount  float64 `json:"balance_amount" db:""`
	SumOfExpCost   float64 `json:"sum_of_exp_cost" db:""`
	UnitCode       string  `json:"unit_code" db:""`
	PORefNo        string  `json:"po_ref_no"  db:""`
	IRRefNo        string  `json:"ir_ref_no" db:""`
	IsCancel       int     `json:"is_cancel" db:""`
	LineNumber     int     `json:"line_number" db:""`
	BarCode        string  `json:"bar_code" db:""`
	PORefLinenum   int     `json:"po_ref_linenum" db:""`
	AVERAGECOST    float64 `json:"averagecost" db:""`
	LotNumber      string  `json:"lot_number" db:""`
	SumOfCost      float64 `json:"sum_of_cost" db:""`
	PackingRate1   float64 `json:"packing_rate_1" db:""`
	PackingRate2   float64 `json:"packing_rate_2" db:""`
}

type Personal struct {
	ApCode string `json:"ap_code" db:"ApCode"`
	//	CreatorCode string `json:"creator_code" db:"CreatorCode"`
}

func (ai *ApInvoice) SearchApInvoiceByDocNo(db *sqlx.DB) error {
	//sql := `DocNo, TaxNo, DocDate, TaxType, ApCode, DepartCode, CreditDay, DueDate, StatementDate, TaxRate, IsConfirm, MyDescription, BillType, BillGroup, ContactCode, SumOfItemAmount,DiscountWord, DiscountAmount, AfterDiscount, BeforeTaxAmount, TaxAmount, TotalAmount, DiscountCase, ExceptTaxAmount, ZeroTaxAmount, PettyCashAmount, SumCashAmount, SumChqAmount, SumBankAmount,DepositIncTax, SumOfDeposit1, SumOfDeposit2, SumOfWTax, NetDebtAmount, HomeAmount, OtherIncome, OtherExpense, ExcessAmount1, ExcessAmount2, BillBalance, CurrencyCode, ExchangeRate, GLFormat,GLStartPosting, GRBillStatus, GRIRBillStatus, IsCancel, IsCreditNote, IsDebitNote,  IsCompleteSave,  AllocateCode, ProjectCode, RecurName, ExchangeProfit, ConfirmCode,ConfirmDateTime, CancelCode, CancelDateTime, PayBillAmount, RefDocNo,CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT`
	//sqlsub := ` MyType, DocNo, TaxNo, TaxType, ItemCode, DocDate, ApCode, DepartCode, MyDescription, ItemName, WHCode, ShelfCode, CNQty, GRRemainQty, Qty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, BalanceAmount, SumOfExpCost, UnitCode, PORefNo, IRRefNo, IsCancel, LineNumber, BarCode, PORefLinenum,  AVERAGECOST,  LotNumber, SumOfCost, TaxRate, PackingRate1, PackingRate2`
	////err := db.Query(ai, sql)
	//fmt.Println("sql =", sql)
	//if err != nil {
	//	return err
	//}
	return nil
}
