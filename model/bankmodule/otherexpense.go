package model

type OtherExpense struct {
	DocNo           string  `json:"doc_no" db:"DocNo"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	RefDocNo        string  `json:"ref_doc_no" db:"RefDocNo"`
	GLBookCode      string  `json:"gl_book_code" db:"GLBookCode"`
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	SumofDebit      float64 `json:"sumof_debit" db:"SumofDebit"`
	SumofCredit     float64 `json:"sumof_credit" db:"SumofCredit"`
	NetAmount       float64 `json:"net_amount" db:"NetAmount"`
	AllocateCode    string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string  `json:"project_code" db:"ProjectCode"`
	ApCode          string  `json:"ap_code" db:"ApCode"`
	PettyCashAmount float64 `json:"petty_cash_amount" db:"PettyCashAmount"`
	SumOfWTax       float64 `json:"sum_of_w_tax" db:"SumOfWTax"`
	SumCashAmount   float64 `json:"sum_cash_amount" db:"SumCashAmount"`
	SumChqAmount    float64 `json:"sum_chq_amount" db:"SumChqAmount"`
	SumCreditAmount float64 `json:"sum_credit_amount" db:"SumCreditAmount"`
	SumBankAmount   float64 `json:"sum_bank_amount" db:"SumBankAmount"`
	OtherIncome     float64 `json:"other_income" db:"OtherIncome"`
	OtherExpense    float64 `json:"other_expense" db:"OtherExpense"`
	ExcessAmount1   float64 `json:"excess_amount_1" db:"ExcessAmount1"`
	ExcessAmount2   float64 `json:"excess_amount_2" db:"ExcessAmount2"`
	BillGroup       string  `json:"bill_group" db:"BillGroup"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	RecurName       string  `json:"recur_name" db:"RecurName"`
	ConfirmCode     string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string  `json:"cancel_date_time" db:"CancelDateTime"`
}

func (ote *OtherExpense) SearchOtherExpenseByDocNo() error {
	//sql := `DocNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, RefDocNo, GLBookCode, DepartCode, MyDescription, SumofDebit, SumofCredit, NetAmount, AllocateCode, ProjectCode, ApCode, PettyCashAmount, SumOfWTax, SumCashAmount, SumChqAmount, SumCreditAmount, SumBankAmount, OtherIncome, OtherExpense, ExcessAmount1, ExcessAmount2, BillGroup, IsConfirm, IsCancel, RecurName, ConfirmCode, ConfirmDateTime, CancelCode,CancelDateTime from dbo.BCOtherExpense WITH (NOLOCK)`
	return nil
}
