package model

type Receipt1 struct {
	DocNo           string `json:"doc_no" db:""`
	TaxNo           string `json:"tax_no" db:""`
	DocDate         string `json:"doc_date" db:""`
	CreatorCode     string `json:"creator_code" db:""`
	CreateDateTime  string `json:"create_date_time" db:""`
	LastEditorCode  string `json:"last_editor_code" db:""`
	LastEditDateT   string `json:"last_edit_date_t" db:""`
	ArCode          string `json:"ar_code" db:""`
	SaleCode        string `json:"sale_code" db:""`
	DepartCode      string `json:"depart_code" db:""`
	SumOfInvoice    float64 `json:"sum_of_invoice" db:""`
	SumOfDebitNote  float64 `json:"sum_of_debit_note" db:""`
	SumOfCreditNote float64 `json:"sum_of_credit_note" db:""`
	BeforeTaxAmount float64 `json:"before_tax_amount" db:""`
	TaxAmount       float64 `json:"tax_amount" db:""`
	TotalAmount     float64 `json:"total_amount" db:""`
	DiscountAmount  float64 `json:"discount_amount" db:""`
	NetAmount       float64 `json:"net_amount" db:""`
	SumCashAmount   float64 `json:"sum_cash_amount" db:""`
	SumChqAmount    float64 `json:"sum_chq_amount" db:""`
	SumCreditAmount float64 `json:"sum_credit_amount" db:""`
	ChargeAmount    float64 `json:"charge_amount" db:""`
	ChangeAmount    float64 `json:"change_amount" db:""`
	SumBankAmount   float64 `json:"sum_bank_amount" db:""`
	SumOfWTax       float64 `json:"sum_of_w_tax" db:""`
	GLFormat        string `json:"gl_format" db:""`
	IsCancel        int `json:"is_cancel" db:""`
	MyDescription   string `json:"my_description" db:""`
	AllocateCode    string `json:"allocate_code" db:""`
	ProjectCode     string `json:"project_code" db:""`
	BillGroup       string `json:"bill_group" db:""`
	IsCompleteSave  int `json:"is_complete_save" db:""`
	SumHomeAmount1  float64 `json:"sum_home_amount_1" db:""`
	SumHomeAmount2  float64 `json:"sum_home_amount_2" db:""`
	DebtLimitReturn float64 `json:"debt_limit_return" db:""`
	CurrencyCode    string `json:"currency_code" db:""`
	ExchangeRate    float64 `json:"exchange_rate" db:""`
	OtherIncome     float64 `json:"other_income" db:""`
	OtherExpense    float64 `json:"other_expense" db:""`
	ExcessAmount1   float64 `json:"excess_amount_1" db:""`
	ExcessAmount2   float64 `json:"excess_amount_2" db:""`
	IsConfirm       int `json:"is_confirm" db:""`
	RecurName       string `json:"recur_name" db:""`
	ConfirmCode     string `json:"confirm_code" db:""`
	ConfirmDateTime string `json:"confirm_date_time" db:""`
	CancelCode      string `json:"cancel_code" db:""`
	CancelDateTime  string `json:"cancel_date_time" db:""`
}

type rcpItem struct {
	InvoiceDate     string `json:"invoice_date" db:""`
	InvoiceNo       string `json:"invoice_no" db:""`
	InvBalance      float64 `json:"inv_balance" db:""`
	PayAmount       float64 `json:"pay_amount" db:""`
	DebtLimitReturn float64 `json:"debt_limit_return" db:""`
	MyDescription   string `json:"my_description" db:""`
	IsCancel        int `json:"is_cancel" db:""`
	LineNumber      int `json:"line_number" db:""`
	BillType        int `json:"bill_type" db:""`
	RefNo           string `json:"ref_no" db:""`
	RefType         int `json:"ref_type" db:""`
	HomeAmount1     float64 `json:"home_amount_1" db:""`
	HomeAmount2     float64 `json:"home_amount_2" db:""`
}

func (rcp *Receipt1) SearchReceiptByDocNo() error {
	//sql := `DocNo, TaxNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, ArCode, SaleCode, DepartCode, SumOfInvoice, SumOfDebitNote, SumOfCreditNote, BeforeTaxAmount,TaxAmount, TotalAmount, DiscountAmount, NetAmount, SumCashAmount, SumChqAmount, SumCreditAmount, ChargeAmount, ChangeAmount, SumBankAmount, SumOfWTax, GLFormat, IsCancel,MyDescription, AllocateCode, ProjectCode, BillGroup, IsCompleteSave, SumHomeAmount1, SumHomeAmount2, DebtLimitReturn, CurrencyCode, ExchangeRate, OtherIncome,OtherExpense, ExcessAmount1, ExcessAmount2, IsConfirm, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime`
	//sqlsub:= `DocNo, DocDate, ArCode, SaleCode,  InvoiceDate, InvoiceNo, InvBalance, PayAmount, DebtLimitReturn,MyDescription,  IsCancel, LineNumber, BillType, RefNo, RefType,  HomeAmount1, HomeAmount2,`
	return nil
}
