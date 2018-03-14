package model

type BankExpense struct {
	DocNo           string
	DocDate         string
	BookNo          string
	CreatorCode     string
	CreateDateTime  string
	LastEditorCode  string
	LastEditDateT   string
	AccountCode     string
	GLBookCode      string
	DepartCode      string
	MyDescription   string
	Amount          float64
	AllocateCode    string
	ProjectCode     string
	IsCancel        int
	IsConfirm       int
	RecurName       string
	ConfirmCode     string
	ConfirmDateTime string
	CancelCode      string
	CancelDateTime  string
}

func (bep *BankExpense) SearchBankExpenseByDocNo() error {
	//sql := `DocNo, DocDate, BookNo, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, AccountCode, GLBookCode, DepartCode, MyDescription, Amount, AllocateCode, ProjectCode, IsCancel, IsConfirm, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime from dbo.BCBANKEXPENSE WITH (NOLOCK)`
	return nil
}
