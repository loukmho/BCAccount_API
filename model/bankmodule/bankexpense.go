package model

type BankExpense struct {
	DocNo           string  `json:"doc_no" db:""`
	DocDate         string  `json:"doc_date" db:""`
	BookNo          string  `json:"book_no" db:""`
	CreatorCode     string  `json:"creator_code" db:""`
	CreateDateTime  string  `json:"create_date_time" db:""`
	LastEditorCode  string  `json:"last_editor_code" db:""`
	LastEditDateT   string  `json:"last_edit_date_t" db:""`
	AccountCode     string  `json:"account_code" db:""`
	GLBookCode      string  `json:"gl_book_code" db:""`
	DepartCode      string  `json:"depart_code" db:""`
	MyDescription   string  `json:"my_description" db:""`
	Amount          float64 `json:"amount" db:""`
	AllocateCode    string  `json:"allocate_code" db:""`
	ProjectCode     string  `json:"project_code" db:""`
	IsCancel        int     `json:"is_cancel" db:""`
	IsConfirm       int     `json:"is_confirm" db:""`
	RecurName       string  `json:"recur_name" db:""`
	ConfirmCode     string  `json:"confirm_code" db:""`
	ConfirmDateTime string  `json:"confirm_date_time" db:""`
	CancelCode      string  `json:"cancel_code" db:""`
	CancelDateTime  string  `json:"cancel_date_time" db:""`
}

func (bep *BankExpense) SearchBankExpenseByDocNo() error {
	//sql := `DocNo, DocDate, BookNo, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, AccountCode, GLBookCode, DepartCode, MyDescription, Amount, AllocateCode, ProjectCode, IsCancel, IsConfirm, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime from dbo.BCBANKEXPENSE WITH (NOLOCK)`
	return nil
}
