package model

type CashBankIn struct {
	DocNo           string
	CreatorCode     string
	CreateDateTime  string
	LastEditorCode  string
	LastEditDateT   string
	DocDate         string
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

func (cbi *CashBankIn) SearchCashBankInByDocNo() error {
	//sql := `DocNo, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, DocDate, BookNo, AccountCode, GLBookCode, IsBankBalance, DepartCode, MyDescription, Amount, AllocateCode, ProjectCode, IsCancel, IsConfirm, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime from dbo.BCCASHBANKIN WITH (NOLOCK)`
	return nil
}
