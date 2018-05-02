package model

type ApOtherDebt struct {
	DocNo           string  `json:"doc_no" db:"DocNo"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	ApCode          string  `json:"ap_code" db:"ApCode"`
	GLBookCode      string  `json:"gl_book_code" db:"GLBookCode"`
	SumofDebit      float64 `json:"sumof_debit" db:"SumofDebit"`
	SumofCredit     float64 `json:"sumof_credit" db:"SumofCredit"`
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	CreditDay       int     `json:"credit_day" db:"CreditDay"`
	DueDate         string  `json:"due_date" db:"DueDate"`
	PayBillDate     string  `json:"pay_bill_date" db:"PayBillDate"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	BillGroup       string  `json:"bill_group" db:"BillGroup"`
	ContactCode     string  `json:"contact_code" db:"ContactCode"`
	NetDebtAmount   float64 `json:"net_debt_amount" db:"NetDebtAmount"`
	BillBalance     float64 `json:"bill_balance" db:"BillBalance"`
	CurrencyCode    string  `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate    float64 `json:"exchange_rate" db:"ExchangeRate"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	AllocateCode    string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string  `json:"project_code" db:"ProjectCode"`
	RecurName       string  `json:"recur_name" db:"RecurName"`
	ConfirmCode     string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string  `json:"cancel_date_time" db:"CancelDateTime"`
	PayBillAmount   float64 `json:"pay_bill_amount" db:"PayBillAmount"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
}

func (atd *ApOtherDebt) SearchApOtherDebtByDocNo() error {
	//sql := `DocNo, DocDate, ApCode, GLBookCode, SumofDebit, SumofCredit, DepartCode, CreditDay, DueDate, PayBillDate, IsConfirm, MyDescription, BillGroup, ContactCode,NetDebtAmount, BillBalance, CurrencyCode, ExchangeRate, IsCancel, AllocateCode, ProjectCode, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime, PayBillAmount,CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT`
	return nil
}
