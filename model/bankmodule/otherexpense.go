package model

type OtherExpense struct {
	DocNo           string  `json:"doc_no" db:""`
	DocDate         string  `json:"doc_date" db:""`
	CreatorCode     string  `json:"creator_code" db:""`
	CreateDateTime  string  `json:"create_date_time" db:""`
	LastEditorCode  string  `json:"last_editor_code" db:""`
	LastEditDateT   string  `json:"last_edit_date_t" db:""`
	RefDocNo        string  `json:"ref_doc_no" db:""`
	GLBookCode      string  `json:"gl_book_code" db:""`
	DepartCode      string  `json:"depart_code" db:""`
	MyDescription   string  `json:"my_description" db:""`
	SumofDebit      float64 `json:"sumof_debit" db:""`
	SumofCredit     float64 `json:"sumof_credit" db:""`
	NetAmount       float64 `json:"net_amount" db:""`
	AllocateCode    string  `json:"allocate_code" db:""`
	ProjectCode     string  `json:"project_code" db:""`
	ApCode          string  `json:"ap_code" db:""`
	PettyCashAmount float64 `json:"petty_cash_amount" db:""`
	SumOfWTax       float64 `json:"sum_of_w_tax" db:""`
	SumCashAmount   float64 `json:"sum_cash_amount" db:""`
	SumChqAmount    float64 `json:"sum_chq_amount" db:""`
	SumCreditAmount float64 `json:"sum_credit_amount" db:""`
	SumBankAmount   float64 `json:"sum_bank_amount" db:""`
	OtherIncome     float64 `json:"other_income" db:""`
	OtherExpense    float64 `json:"other_expense" db:""`
	ExcessAmount1   float64 `json:"excess_amount_1" db:""`
	ExcessAmount2   float64 `json:"excess_amount_2" db:""`
	BillGroup       string  `json:"bill_group" db:""`
	IsConfirm       int     `json:"is_confirm" db:""`
	IsCancel        int     `json:"is_cancel" db:""`
	RecurName       string  `json:"recur_name" db:""`
	ConfirmCode     string  `json:"confirm_code" db:""`
	ConfirmDateTime string  `json:"confirm_date_time" db:""`
	CancelCode      string  `json:"cancel_code" db:""`
	CancelDateTime  string  `json:"cancel_date_time" db:""`
}

func (ote *OtherExpense) SearchOtherExpenseByDocNo() error {
	//sql := `DocNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, RefDocNo, GLBookCode, DepartCode, MyDescription, SumofDebit, SumofCredit, NetAmount, AllocateCode, ProjectCode, ApCode, PettyCashAmount, SumOfWTax, SumCashAmount, SumChqAmount, SumCreditAmount, SumBankAmount, OtherIncome, OtherExpense, ExcessAmount1, ExcessAmount2, BillGroup, IsConfirm, IsCancel, RecurName, ConfirmCode, ConfirmDateTime, CancelCode,CancelDateTime from dbo.BCOtherExpense WITH (NOLOCK)`
	return nil
}
