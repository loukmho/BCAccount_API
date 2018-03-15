package model

type CashBankOut struct {
	DocNo           string  `json:"doc_no" db:""`
	DocDate         string  `json:"doc_date" db:""`
	CreatorCode     string  `json:"creator_code" db:""`
	CreateDateTime  string  `json:"create_date_time" db:""`
	LastEditorCode  string  `json:"last_editor_code" db:""`
	LastEditDateT   string  `json:"last_edit_date_t" db:""`
	BookNo          string  `json:"book_no" db:""`
	AccountCode     string  `json:"account_code" db:""`
	GLBookCode      string  `json:"gl_book_code" db:""`
	IsBankBalance   float64 `json:"is_bank_balance" db:""`
	DepartCode      string  `json:"depart_code" db:""`
	MyDescription   string  `json:"my_description" db:""`
	Amount          float64 `json:"amount" db:""`
	AllocateCode    string  `json:"allocate_code" db:""`
	ProjectCode     string  `json:"project_code" db:""`
	IsCancel        int     `json:"is_cancel" db:""`
	IsConfirm       int     `json:"is_confirm" db:""`
	ConfirmCode     string  `json:"confirm_code" db:""`
	ConfirmDateTime string  `json:"confirm_date_time" db:""`
	CancelCode      string  `json:"cancel_code" db:""`
	CancelDateTime  string  `json:"cancel_date_time" db:""`
}

func (cbo *CashBankOut) SearchCashBankOutByDocNo() error {
	//sql := `DocNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, BookNo, AccountCode, GLBookCode, IsBankBalance,DepartCode, MyDescription, Amount,  AllocateCode, ProjectCode, IsCancel, IsConfirm, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime from dbo.BCCashBankOut WITH (NOLOCK)`
	return nil
}
