package model

type ReturnDepSpecial struct {
	DocNo           string  `json:"doc_no" db:""`
	DocDate         string  `json:"doc_date" db:""`
	CreatorCode     string  `json:"creator_code" db:""`
	CreateDateTime  string  `json:"create_date_time" db:""`
	LastEditorCode  string  `json:"last_editor_code" db:""`
	LastEditDateT   string  `json:"last_edit_date_t" db:""`
	EarnestNo       string  `json:"earnest_no" db:""`
	ArCode          string  `json:"ar_code" db:""`
	DepartCode      string  `json:"depart_code" db:""`
	IsConfirm       int     `json:"is_confirm" db:""`
	CreditDay       int     `json:"credit_day" db:""`
	SaleCode        string  `json:"sale_code" db:""`
	DueDate         string  `json:"due_date" db:""`
	MyDescription   string  `json:"my_description" db:""`
	TotalAmount     float64 `json:"total_amount" db:""`
	BeforeTaxAmount float64 `json:"before_tax_amount" db:""`
	NetAmount       float64 `json:"net_amount" db:""`
	SumOfWTax       float64 `json:"sum_of_w_tax" db:""`
	ExcessAmount1   float64 `json:"excess_amount_1" db:""`
	ExcessAmount2   float64 `json:"excess_amount_2" db:""`
	BillBalance     float64 `json:"bill_balance" db:""`
	CurrencyCode    string  `json:"currency_code" db:""`
	ExchangeRate    float64 `json:"exchange_rate" db:""`
	PettyCashAmount float64 `json:"petty_cash_amount" db:""`
	SumCashAmount   float64 `json:"sum_cash_amount" db:""`
	SumChqAmount    float64 `json:"sum_chq_amount" db:""`
	SumCreditAmount float64 `json:"sum_credit_amount" db:""`
	SumBankAmount   float64 `json:"sum_bank_amount" db:""`
	EarnestAmount   float64 `json:"earnest_amount" db:""`
	EarnestBalance  float64 `json:"earnest_balance" db:""`
	OtherIncome     float64 `json:"other_income" db:""`
	OtherExpense    float64 `json:"other_expense" db:""`
	GLFormat        string  `json:"gl_format" db:""`
	IsCancel        int     `json:"is_cancel" db:""`
	IsReturnMoney   int     `json:"is_return_money" db:""`
	AllocateCode    string  `json:"allocate_code" db:""`
	ProjectCode     string  `json:"project_code" db:""`
	BillGroup       string  `json:"bill_group" db:""`
	RecurName       string  `json:"recur_name" db:""`
	ConfirmCode     string  `json:"confirm_code" db:""`
	ConfirmDateTime string  `json:"confirm_date_time" db:""`
	CancelCode      string  `json:"cancel_code" db:""`
	CancelDateTime  string  `json:"cancel_date_time" db:""`
}

func (rds *ReturnDepSpecial) SearchReturnDepSpecialByDocNo() error {
	//sql := `DocNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, EarnestNo, ArCode, DepartCode, IsConfirm, CreditDay, SaleCode, DueDate, MyDescription, TotalAmount,BeforeTaxAmount, NetAmount, SumOfWTax, ExcessAmount1, ExcessAmount2, BillBalance, CurrencyCode, ExchangeRate, PettyCashAmount, SumCashAmount, SumChqAmount, SumCreditAmount, SumBankAmount,EarnestAmount, EarnestBalance, OtherIncome, OtherExpense, GLFormat, IsCancel, IsReturnMoney, AllocateCode, ProjectCode, BillGroup, RecurName, ConfirmCode, ConfirmDateTime, CancelCode,CancelDateTime`
	return nil
}
