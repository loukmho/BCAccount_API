package model

type CreditNote struct {
	DocNo           string  `json:"doc_no" db:"DocNo"`
	TaxNo           string  `json:"tax_no" db:"TaxNo"`
	TaxDate         string  `json:"tax_date" db:"TaxDate"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	DueDate         string  `json:"due_date" db:"DueDate"`
	TaxType         int     `json:"tax_type" db:"TaxType"`
	ArCode          string  `json:"ar_code" db:"ArCode"`
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	SaleCode        string  `json:"sale_code" db:"SaleCode"`
	CashierCode     string  `json:"cashier_code" db:"CashierCode"`
	TaxRate         float64 `json:"tax_rate" db:"TaxRate"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	SumOfItemAmount float64 `json:"sum_of_item_amount" db:"SumOfItemAmount"`
	SumOldAmount    float64 `json:"sum_old_amount" db:"SumOldAmount"`
	SumTrueAmount   float64 `json:"sum_true_amount" db:"SumTrueAmount"`
	SumofDiffAmount float64 `json:"sumof_diff_amount" db:"SumofDiffAmount"`
	DiscountWord    string  `json:"discount_word" db:"DiscountWord"`
	DiscountAmount  float64 `json:"discount_amount" db:"DiscountAmount"`
	SumofBeforeTax  float64 `json:"sumof_before_tax" db:"SumofBeforeTax"`
	SumOfTaxAmount  float64 `json:"sum_of_tax_amount" db:"SumOfTaxAmount"`
	SumOfTotalTax   float64 `json:"sum_of_total_tax" db:"SumOfTotalTax"`
	SumOfZeroTax    float64 `json:"sum_of_zero_tax" db:"SumOfZeroTax"`
	SumOfExceptTax  float64 `json:"sum_of_except_tax" db:"SumOfExceptTax"`
	SumOfWTax       float64 `json:"sum_of_w_tax" db:"SumOfWTax"`
	NetDebtAmount   float64 `json:"net_debt_amount" db:"NetDebtAmount"`
	BillBalance     float64 `json:"bill_balance" db:"BillBalance"`
	CurrencyCode    string  `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate    float64 `json:"exchange_rate" db:"ExchangeRate"`
	GLFormat        string  `json:"gl_format" db:"GLFormat"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	IsCompleteSave  int     `json:"is_complete_save" db:"IsCompleteSave"`
	ReturnMoney     int     `json:"return_money" db:"ReturnMoney"`
	ReturnStatus    int     `json:"return_status" db:"ReturnStatus"`
	ReturnCash      int     `json:"return_cash" db:"ReturnCash"`
	OtherIncome     float64 `json:"other_income" db:"OtherIncome"`
	OtherExpense    float64 `json:"other_expense" db:"OtherExpense"`
	ExcessAmount1   float64 `json:"excess_amount_1" db:"ExcessAmount1"`
	ExcessAmount2   float64 `json:"excess_amount_2" db:"ExcessAmount2"`
	SumCashAmount   float64 `json:"sum_cash_amount" db:"SumCashAmount"`
	SumChqAmount    float64 `json:"sum_chq_amount" db:"SumChqAmount"`
	SumCreditAmount float64 `json:"sum_credit_amount" db:"SumCreditAmount"`
	SumBankAmount   float64 `json:"sum_bank_amount" db:"SumBankAmount"`
	ChargeAmount    float64 `json:"charge_amount" db:"ChargeAmount"`
	ChangeAmount    float64 `json:"change_amount" db:"ChangeAmount"`
	CauseType       int     `json:"cause_type" db:"CauseType"`
	PayBillStatus   int     `json:"pay_bill_status" db:"PayBillStatus"`
	IsCNDeposit     int     `json:"is_cn_deposit" db:"IsCNDeposit"`
	IsPos           int     `json:"is_pos" db:"IsPos"`
	PosDocNo        string  `json:"pos_doc_no" db:"PosDocNo"`
	CauseCode       string  `json:"cause_code" db:"CauseCode"`
	AllocateCode    string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string  `json:"project_code" db:"ProjectCode"`
	BillGroup       string  `json:"bill_group" db:"BillGroup"`
	RecurName       string  `json:"recur_name" db:"RecurName"`
	ConfirmCode     string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string  `json:"cancel_date_time" db:"CancelDateTime"`
	PayBillAmount   float64 `json:"pay_bill_amount" db:"PayBillAmount"`
	BillTemporary   float64 `json:"bill_temporary" db:"BillTemporary"`
}

var CrdItem struct {
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
	ExceptTax      int     `json:"except_tax" db:"ExceptTax"`
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

func (crd *CreditNote) SearchCreditNoteByDocNo() error {
	//sql := `DocNo,TaxNo,TaxDate ,CreatorCode,CreateDateTime,LastEditorCode,LastEditDateT,DocDate,DueDate,TaxType,ArCode,DepartCode,SaleCode,CashierCode,TaxRate,IsConfirm,MyDescription,SumOfItemAmount,SumOldAmount,SumTrueAmount,SumofDiffAmount,DiscountWord,DiscountAmount,SumofBeforeTax,SumOfTaxAmount,SumOfTotalTax,SumOfZeroTax,SumOfExceptTax,SumOfWTax,NetDebtAmount,BillBalance,CurrencyCode,ExchangeRate,GLFormat,GLStartPosting,IsPostGL,IsCancel,IsCompleteSave,GLTransData,ReturnMoney,ReturnStatus,ReturnCash,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,ChargeAmount,ChangeAmount,CauseType,PayBillStatus,IsCNDeposit,IsPos,PosDocNo,CauseCode,AllocateCode,ProjectCode,BillGroup,RecurName,ConfirmCode,ConfirmDateTime,CancelCode,CancelDateTime,PayBillAmount,BillTemporary`
	//sqlsub := `MyType, DocNo, TaxNo, TaxType, ItemCode, DocDate, ArCode, DepartCode, SaleCode, CashierCode, MyDescription, ItemName, WHCode, ShelfCode, DiscQty, TempQty, BillQty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost, UnitCode, InvoiceNo, ItemType, ExceptTax, IsPos, IsCancel, LineNumber, RefLineNumber, BarCode,AVERAGECOST, LotNumber, PackingRate1, PackingRate2`
	return nil
}
