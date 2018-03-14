package model

type ApOtherDebt struct {
	DocNo           string
	DocDate         string
	ApCode          string
	GLBookCode      string
	SumofDebit      float64
	SumofCredit     float64
	DepartCode      string
	CreditDay       int
	DueDate         string
	PayBillDate     string
	IsConfirm       int
	MyDescription   string
	BillGroup       string
	ContactCode     string
	NetDebtAmount   float64
	BillBalance     float64
	CurrencyCode    string
	ExchangeRate    float64
	IsCancel        int
	AllocateCode    string
	ProjectCode     string
	RecurName       string
	ConfirmCode     string
	ConfirmDateTime string
	CancelCode      string
	CancelDateTime  string
	PayBillAmount   float64
	CreatorCode     string
	CreateDateTime  string
	LastEditorCode  string
	LastEditDateT   string
}

func (atd *ApOtherDebt) SearchApOtherDebtByDocNo() error {
	//sql := `DocNo, DocDate, ApCode, GLBookCode, SumofDebit, SumofCredit, DepartCode, CreditDay, DueDate, PayBillDate, IsConfirm, MyDescription, BillGroup, ContactCode,NetDebtAmount, BillBalance, CurrencyCode, ExchangeRate, IsCancel, AllocateCode, ProjectCode, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime, PayBillAmount,CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT`
	return nil
}
