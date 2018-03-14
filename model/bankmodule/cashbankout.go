package model

type CashBankOut struct {
	DocNo           string
	DocDate         string
	CreatorCode     string
	CreateDateTime  string
	LastEditorCode  string
	LastEditDateT   string
	BookNo          string
	AccountCode     string
	GLBookCode      string
	IsBankBalance   float64
	DepartCode      string
	MyDescription   string
	Amount          float64
	AllocateCode    string
	ProjectCode     string
	IsCancel        int
	IsConfirm       int
	ConfirmCode     string
	ConfirmDateTime string
	CancelCode      string
	CancelDateTime  string
}

func (cbo *CashBankOut) SearchCashBankOutByDocNo() error {
	//sql := `DocNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, BookNo, AccountCode, GLBookCode, IsBankBalance,DepartCode, MyDescription, Amount,  AllocateCode, ProjectCode, IsCancel, IsConfirm, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime from dbo.BCCashBankOut WITH (NOLOCK)`
	return nil
}
