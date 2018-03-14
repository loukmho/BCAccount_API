package model

type ArOtherDebt struct {
	DocNo           string
	DocDate         string
	ArCode          string
	GLBookCode      string
	SumofDebit      float64
	SumofCredit     float64
	DepartCode      string
	CreditDay       int
	DueDate         string
	PayBillDate     string
	SaleCode        string
	IsConfirm       int
	PayBillStatus   int
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
	CreatorCode     string
	CreateDateTime  string
	LastEditorCode  string
	LastEditDateT   string
	ConfirmCode     string
	ConfirmDateTime string
	CancelCode      string
	CancelDateTime  string
	PayBillAmount   float64
	BillTemporary   float64
}

func (ard *ArOtherDebt) SearchArOtherDebtByDocNo() error {
	//sql := `DocNo, DocDate, ArCode, GLBookCode, SumofDebit, SumofCredit, DepartCode, CreditDay, DueDate, PayBillDate, SaleCode, IsConfirm, PayBillStatus, MyDescription, BillGroup,ContactCode, NetDebtAmount, BillBalance, CurrencyCode, ExchangeRate, IsCancel, AllocateCode, ProjectCode, RecurName, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, ConfirmCode, ConfirmDateTime,CancelCode, CancelDateTime, PayBillAmount, BillTemporary`
	return nil
}
