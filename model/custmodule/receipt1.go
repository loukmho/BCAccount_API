package model

type Receipt1 struct {
	DocNo           string  `json:"doc_no" db:"DocNo"`
	TaxNo           string  `json:"tax_no" db:"TaxNo"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	ArCode          string  `json:"ar_code" db:"ArCode"`
	SaleCode        string  `json:"sale_code" db:"SaleCode"`
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	SumOfInvoice    float64 `json:"sum_of_invoice" db:"SumOfInvoice"`
	SumOfDebitNote  float64 `json:"sum_of_debit_note" db:"SumOfDebitNote"`
	SumOfCreditNote float64 `json:"sum_of_credit_note" db:"SumOfCreditNote"`
	BeforeTaxAmount float64 `json:"before_tax_amount" db:"BeforeTaxAmount"`
	TaxAmount       float64 `json:"tax_amount" db:"TaxAmount"`
	TotalAmount     float64 `json:"total_amount" db:"TotalAmount"`
	DiscountAmount  float64 `json:"discount_amount" db:"DiscountAmount"`
	NetAmount       float64 `json:"net_amount" db:"NetAmount"`
	SumCashAmount   float64 `json:"sum_cash_amount" db:"SumCashAmount"`
	SumChqAmount    float64 `json:"sum_chq_amount" db:"SumChqAmount"`
	SumCreditAmount float64 `json:"sum_credit_amount" db:"SumCreditAmount"`
	ChargeAmount    float64 `json:"charge_amount" db:"ChargeAmount"`
	ChangeAmount    float64 `json:"change_amount" db:"ChangeAmount"`
	SumBankAmount   float64 `json:"sum_bank_amount" db:"SumBankAmount"`
	SumOfWTax       float64 `json:"sum_of_w_tax" db:"SumOfWTax"`
	GLFormat        string  `json:"gl_format" db:"GLFormat"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	AllocateCode    string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string  `json:"project_code" db:"ProjectCode"`
	BillGroup       string  `json:"bill_group" db:"BillGroup"`
	IsCompleteSave  int     `json:"is_complete_save" db:"IsCompleteSave"`
	SumHomeAmount1  float64 `json:"sum_home_amount_1" db:"SumHomeAmount1"`
	SumHomeAmount2  float64 `json:"sum_home_amount_2" db:"SumHomeAmount2"`
	DebtLimitReturn float64 `json:"debt_limit_return" db:"DebtLimitReturn"`
	CurrencyCode    string  `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate    float64 `json:"exchange_rate" db:"ExchangeRate"`
	OtherIncome     float64 `json:"other_income" db:"OtherIncome"`
	OtherExpense    float64 `json:"other_expense" db:"OtherExpense"`
	ExcessAmount1   float64 `json:"excess_amount_1" db:"ExcessAmount1"`
	ExcessAmount2   float64 `json:"excess_amount_2" db:"ExcessAmount2"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	RecurName       string  `json:"recur_name" db:"RecurName"`
	ConfirmCode     string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string  `json:"cancel_date_time" db:"CancelDateTime"`
}

type rcpItem struct {
	InvoiceDate     string  `json:"invoice_date" db:"InvoiceDate"`
	InvoiceNo       string  `json:"invoice_no" db:"InvoiceNo"`
	InvBalance      float64 `json:"inv_balance" db:"InvBalance"`
	PayAmount       float64 `json:"pay_amount" db:"PayAmount"`
	DebtLimitReturn float64 `json:"debt_limit_return" db:"DebtLimitReturn"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	LineNumber      int     `json:"line_number" db:"LineNumber"`
	BillType        int     `json:"bill_type" db:"BillType"`
	RefNo           string  `json:"ref_no" db:"RefNo"`
	RefType         int     `json:"ref_type" db:"RefType"`
	HomeAmount1     float64 `json:"home_amount_1" db:"HomeAmount1"`
	HomeAmount2     float64 `json:"home_amount_2" db:"HomeAmount2"`
}

func (rcp *Receipt1) SearchReceiptByDocNo() error {
	//sql := `DocNo, TaxNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, ArCode, SaleCode, DepartCode, SumOfInvoice, SumOfDebitNote, SumOfCreditNote, BeforeTaxAmount,TaxAmount, TotalAmount, DiscountAmount, NetAmount, SumCashAmount, SumChqAmount, SumCreditAmount, ChargeAmount, ChangeAmount, SumBankAmount, SumOfWTax, GLFormat, IsCancel,MyDescription, AllocateCode, ProjectCode, BillGroup, IsCompleteSave, SumHomeAmount1, SumHomeAmount2, DebtLimitReturn, CurrencyCode, ExchangeRate, OtherIncome,OtherExpense, ExcessAmount1, ExcessAmount2, IsConfirm, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime`
	//sqlsub:= `DocNo, DocDate, ArCode, SaleCode,  InvoiceDate, InvoiceNo, InvBalance, PayAmount, DebtLimitReturn,MyDescription,  IsCancel, LineNumber, BillType, RefNo, RefType,  HomeAmount1, HomeAmount2,`
	return nil
}
