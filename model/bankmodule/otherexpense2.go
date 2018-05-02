package model

type OtherExpense2 struct {
	DocNo           string  `json:"doc_no" db:"DocNo"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	SumofAmount     float64 `json:"sumof_amount" db:"SumofAmount"`
	AllocateCode    string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string  `json:"project_code" db:"ProjectCode"`
	ApCode          string  `json:"ap_code" db:"ApCode"`
	GLFormat        string  `json:"gl_format" db:"GLFormat"`
	PersonCode      string  `json:"person_code" db:"PersonCode"`
	BillStatus      int     `json:"bill_status" db:"BillStatus"`
	PettyCashAmount float64 `json:"petty_cash_amount" db:"PettyCashAmount"`
	SumCashAmount   float64 `json:"sum_cash_amount" db:"SumCashAmount"`
	SumChqAmount    float64 `json:"sum_chq_amount" db:"SumChqAmount"`
	SumBankAmount   float64 `json:"sum_bank_amount" db:"SumBankAmount"`
	OtherIncome     float64 `json:"other_income" db:"OtherIncome"`
	OtherExpense    float64 `json:"other_expense" db:"OtherExpense"`
	ExcessAmount1   float64 `json:"excess_amount_1" db:"ExcessAmount1"`
	ExcessAmount2   float64 `json:"excess_amount_2" db:"ExcessAmount2"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	RecurName       string  `json:"recur_name" db:"RecurName"`
	ConfirmCode     string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	CancelCode      string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string  `json:"cancel_date_time" db:"CancelDateTime"`
}

func (oxp2 *OtherExpense2) SearchOtherExpense2ByDocNo() error {
	//sql := `DocNo, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, DocDate, DepartCode, MyDescription, SumofAmount, AllocateCode, ProjectCode, ApCode, GLFormat,  PersonCode, BillStatus, PettyCashAmount, SumCashAmount, SumChqAmount, SumBankAmount, OtherIncome, OtherExpense, ExcessAmount1, ExcessAmount2, IsConfirm, RecurName, ConfirmCode, ConfirmDateTime, IsCancel, CancelCode, CancelDateTime from dbo.BCOTHEREXPENSE2 WITH (NOLOCK)`
	return nil
}
