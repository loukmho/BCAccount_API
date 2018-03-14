package model

type OtherExpense struct {
	DocNo           string
	DocDate         string
	CreatorCode     string
	CreateDateTime  string
	LastEditorCode  string
	LastEditDateT   string
	RefDocNo        string
	GLBookCode      string
	DepartCode      string
	MyDescription   string
	SumofDebit      float64
	SumofCredit     float64
	NetAmount       float64
	AllocateCode    string
	ProjectCode     string
	ApCode          string
	PettyCashAmount float64
	SumOfWTax       float64
	SumCashAmount   float64
	SumChqAmount    float64
	SumCreditAmount float64
	SumBankAmount   float64
	OtherIncome     float64
	OtherExpense    float64
	ExcessAmount1   float64
	ExcessAmount2   float64
	BillGroup       string
	IsConfirm       int
	IsCancel        int
	RecurName       string
	ConfirmCode     string
	ConfirmDateTime string
	CancelCode      string
	CancelDateTime  string
}

func (ote *OtherExpense) SearchOtherExpenseByDocNo() error {
	//sql := `DocNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, RefDocNo, GLBookCode, DepartCode, MyDescription, SumofDebit, SumofCredit, NetAmount, AllocateCode, ProjectCode, ApCode, PettyCashAmount, SumOfWTax, SumCashAmount, SumChqAmount, SumCreditAmount, SumBankAmount, OtherIncome, OtherExpense, ExcessAmount1, ExcessAmount2, BillGroup, IsConfirm, IsCancel, RecurName, ConfirmCode, ConfirmDateTime, CancelCode,CancelDateTime from dbo.BCOtherExpense WITH (NOLOCK)`
	return nil
}
