package model

type CashBankIn struct {
	DocNo           string  `json:"doc_no" db:"DocNo"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	BookNo          string  `json:"book_no" db:"BookNo"`
	AccountCode     string  `json:"account_code" db:"AccountCode"`
	GLBookCode      string  `json:"gl_book_code" db:"GLBookCode"`
	IsBankBalance   float64 `json:"is_bank_balance" db:"IsBankBalance"`
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	Amount          float64 `json:"amount" db:"Amount"`
	AllocateCode    string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string  `json:"project_code" db:"ProjectCode"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	ConfirmCode     string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string  `json:"cancel_date_time" db:"CancelDateTime"`
}

func (cbi *CashBankIn) SearchCashBankInByDocNo() error {
	//sql := `DocNo, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, DocDate, BookNo, AccountCode, GLBookCode, IsBankBalance, DepartCode, MyDescription, Amount, AllocateCode, ProjectCode, IsCancel, IsConfirm, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime from dbo.BCCASHBANKIN WITH (NOLOCK)`
	return nil
}
