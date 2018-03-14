package model

type ArOtherDebt struct {
	DocNo           string `json:"doc_no" db:""`
	DocDate         string `json:"doc_date" db:""`
	ArCode          string `json:"ar_code" db:""`
	GLBookCode      string `json:"gl_book_code" db:""`
	SumofDebit      float64 `json:"sumof_debit" db:""`
	SumofCredit     float64 `json:"sumof_credit" db:""`
	DepartCode      string `json:"depart_code" db:""`
	CreditDay       int `json:"credit_day" db:""`
	DueDate         string `json:"due_date" db:""`
	PayBillDate     string `json:"pay_bill_date" db:""`
	SaleCode        string `json:"sale_code" db:""`
	IsConfirm       int `json:"is_confirm" db:""`
	PayBillStatus   int `json:"pay_bill_status" db:""`
	MyDescription   string `json:"my_description" db:""`
	BillGroup       string `json:"bill_group" db:""`
	ContactCode     string `json:"contact_code" db:""`
	NetDebtAmount   float64 `json:"net_debt_amount" db:""`
	BillBalance     float64 `json:"bill_balance" db:""`
	CurrencyCode    string `json:"currency_code" db:""`
	ExchangeRate    float64 `json:"exchange_rate" db:""`
	IsCancel        int `json:"is_cancel" db:""`
	AllocateCode    string `json:"allocate_code" db:""`
	ProjectCode     string `json:"project_code" db: ""`
	RecurName       string `json:"recur_name" db:""`
	CreatorCode     string `json:"creator_code" db:""`
	CreateDateTime  string `json:"create_date_time" db:""`
	LastEditorCode  string `json:"last_editor_code" db:""`
	LastEditDateT   string `json:"last_edit_date_t" db:""`
	ConfirmCode     string `json:"confirm_code" db:""`
	ConfirmDateTime string `json:"confirm_date_time" db:""`
	CancelCode      string `json:"cancel_code" db:""`
	CancelDateTime  string `json:"cancel_date_time" db:""`
	PayBillAmount   float64 `json:"pay_bill_amount" db:""`
	BillTemporary   float64 `json:"bill_temporary" db:""`
}

func (ard *ArOtherDebt) SearchArOtherDebtByDocNo() error {
	//sql := `DocNo, DocDate, ArCode, GLBookCode, SumofDebit, SumofCredit, DepartCode, CreditDay, DueDate, PayBillDate, SaleCode, IsConfirm, PayBillStatus, MyDescription, BillGroup,ContactCode, NetDebtAmount, BillBalance, CurrencyCode, ExchangeRate, IsCancel, AllocateCode, ProjectCode, RecurName, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, ConfirmCode, ConfirmDateTime,CancelCode, CancelDateTime, PayBillAmount, BillTemporary`
	return nil
}
