package model

type CreditNote struct {
	DocNo           string  `json:"doc_no" db:""`
	TaxNo           string  `json:"tax_no" db:""`
	TaxDate         string  `json:"tax_date" db:""`
	CreatorCode     string  `json:"creator_code" db:""`
	CreateDateTime  string  `json:"create_date_time"`
	LastEditorCode  string  `json:"last_editor_code" db:""`
	LastEditDateT   string  `json:"last_edit_date_t" db:""`
	DocDate         string  `json:"doc_date" db:""`
	DueDate         string  `json:"due_date" db:""`
	TaxType         int     `json:"tax_type" db:""`
	ArCode          string  `json:"ar_code" db:""`
	DepartCode      string  `json:"depart_code" db:""`
	SaleCode        string  `json:"sale_code" db:""`
	CashierCode     string  `json:"cashier_code" db:""`
	TaxRate         float64 `json:"tax_rate" db:""`
	IsConfirm       int     `json:"is_confirm" db:""`
	MyDescription   string  `json:"my_description" db:""`
	SumOfItemAmount float64 `json:"sum_of_item_amount" db:""`
	SumOldAmount    float64 `json:"sum_old_amount" db:""`
	SumTrueAmount   float64 `json:"sum_true_amount" db:""`
	SumofDiffAmount float64 `json:"sumof_diff_amount" db:""`
	DiscountWord    string  `json:"discount_word" db:""`
	DiscountAmount  float64 `json:"discount_amount" db:""`
	SumofBeforeTax  float64 `json:"sumof_before_tax" db:""`
	SumOfTaxAmount  float64 `json:"sum_of_tax_amount" db:""`
	SumOfTotalTax   float64 `json:"sum_of_total_tax" db:""`
	SumOfZeroTax    float64 `json:"sum_of_zero_tax" db:""`
	SumOfExceptTax  float64 `json:"sum_of_except_tax" db:""`
	SumOfWTax       float64 `json:"sum_of_w_tax" db:""`
	NetDebtAmount   float64 `json:"net_debt_amount" db:""`
	BillBalance     float64 `json:"bill_balance" db:""`
	CurrencyCode    string  `json:"currency_code" db:""`
	ExchangeRate    float64 `json:"exchange_rate" db:""`
	GLFormat        string  `json:"gl_format" db:""`
	IsCancel        int     `json:"is_cancel" db:""`
	IsCompleteSave  int     `json:"is_complete_save"`
	ReturnMoney     int     `json:"return_money" db:""`
	ReturnStatus    int     `json:"return_status" db:""`
	ReturnCash      int     `json:"return_cash" db:""`
	OtherIncome     float64 `json:"other_income" db:""`
	OtherExpense    float64 `json:"other_expense" db:""`
	ExcessAmount1   float64 `json:"excess_amount_1" db:""`
	ExcessAmount2   float64 `json:"excess_amount_2" db:""`
	SumCashAmount   float64 `json:"sum_cash_amount" db:""`
	SumChqAmount    float64 `json:"sum_chq_amount" db:""`
	SumCreditAmount float64 `json:"sum_credit_amount" db:""`
	SumBankAmount   float64 `json:"sum_bank_amount" db:""`
	ChargeAmount    float64 `json:"charge_amount" db:""`
	ChangeAmount    float64 `json:"change_amount" db:""`
	CauseType       int     `json:"cause_type" db:""`
	PayBillStatus   int     `json:"pay_bill_status" db:""`
	IsCNDeposit     int     `json:"is_cn_deposit" db:""`
	IsPos           int     `json:"is_pos" db:""`
	PosDocNo        string  `json:"pos_doc_no" db:""`
	CauseCode       string  `json:"cause_code" db:""`
	AllocateCode    string  `json:"allocate_code" db:""`
	ProjectCode     string  `json:"project_code" db:""`
	BillGroup       string  `json:"bill_group" db:""`
	RecurName       string  `json:"recur_name" db:""`
	ConfirmCode     string  `json:"confirm_code" db:""`
	ConfirmDateTime string  `json:"confirm_date_time" db:""`
	CancelCode      string  `json:"cancel_code" db:""`
	CancelDateTime  string  `json:"cancel_date_time" db:""`
	PayBillAmount   float64 `json:"pay_bill_amount" db:""`
	BillTemporary   float64 `json:"bill_temporary" db:""`
}

var CrdItem struct {
	MyType         int     `json:"my_type" db:""`
	ItemCode       string  `json:"item_code" db:""`
	MyDescription  string  `json:"my_description" db:""`
	ItemName       string  `json:"item_name" db:""`
	WHCode         string  `json:"wh_code" db:""`
	ShelfCode      string  `json:"shelf_code" db:""`
	DiscQty        float64 `json:"disc_qty" db:""`
	TempQty        float64 `json:"temp_qty" db:""`
	BillQty        float64 `json:"bill_qty" db:""`
	Price          float64 `json:"price" db:""`
	DiscountWord   string  `json:"discount_word" db:""`
	DiscountAmount float64 `json:"discount_amount" db:""`
	Amount         float64 `json:"amount" db:""`
	NetAmount      float64 `json:"net_amount" db:""`
	HomeAmount     float64 `json:"home_amount" db:""`
	SumOfCost      float64 `json:"sum_of_cost" db:""`
	UnitCode       string  `json:"unit_code" db:""`
	InvoiceNo      string  `json:"invoice_no" db:""`
	ExceptTax      int     `json:"except_tax" db:""`
	IsPos          int     `json:"is_pos" db:""`
	IsCancel       int     `json:"is_cancel" db:""`
	LineNumber     int     `json:"line_number" db:""`
	RefLineNumber  int     `json:"ref_line_number" db:""`
	BarCode        string  `json:"bar_code" db:""`
	AVERAGECOST    float64 `json:"averagecost" db:""`
	LotNumber      string  `json:"lot_number" db:""`
	PackingRate1   float64 `json:"packing_rate_1" db:""`
	PackingRate2   float64 `json:"packing_rate_2" db:""`
}

func (crd *CreditNote) SearchCreditNoteByDocNo() error {
	//sql := `DocNo,TaxNo,TaxDate ,CreatorCode,CreateDateTime,LastEditorCode,LastEditDateT,DocDate,DueDate,TaxType,ArCode,DepartCode,SaleCode,CashierCode,TaxRate,IsConfirm,MyDescription,SumOfItemAmount,SumOldAmount,SumTrueAmount,SumofDiffAmount,DiscountWord,DiscountAmount,SumofBeforeTax,SumOfTaxAmount,SumOfTotalTax,SumOfZeroTax,SumOfExceptTax,SumOfWTax,NetDebtAmount,BillBalance,CurrencyCode,ExchangeRate,GLFormat,GLStartPosting,IsPostGL,IsCancel,IsCompleteSave,GLTransData,ReturnMoney,ReturnStatus,ReturnCash,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,ChargeAmount,ChangeAmount,CauseType,PayBillStatus,IsCNDeposit,IsPos,PosDocNo,CauseCode,AllocateCode,ProjectCode,BillGroup,RecurName,ConfirmCode,ConfirmDateTime,CancelCode,CancelDateTime,PayBillAmount,BillTemporary`
	//sqlsub := `MyType, DocNo, TaxNo, TaxType, ItemCode, DocDate, ArCode, DepartCode, SaleCode, CashierCode, MyDescription, ItemName, WHCode, ShelfCode, DiscQty, TempQty, BillQty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost, UnitCode, InvoiceNo, ItemType, ExceptTax, IsPos, IsCancel, LineNumber, RefLineNumber, BarCode,AVERAGECOST, LotNumber, PackingRate1, PackingRate2`
	return nil
}
