package model

type OtherExpense2 struct {
	DocNo           string
	CreatorCode     string
	CreateDateTime  string
	LastEditorCode  string
	LastEditDateT   string
	DocDate         string
	DepartCode      string
	MyDescription   string
	SumofAmount     float64
	AllocateCode    string
	ProjectCode     string
	ApCode          string
	GLFormat        string
	PersonCode      string
	BillStatus      int
	PettyCashAmount float64
	SumCashAmount   float64
	SumChqAmount    float64
	SumBankAmount   float64
	OtherIncome     float64
	OtherExpense    float64
	ExcessAmount1   float64
	ExcessAmount2   float64
	IsConfirm       int
	RecurName       string
	ConfirmCode     string
	ConfirmDateTime string
	IsCancel        int
	CancelCode      string
	CancelDateTime  string
}

func (oxp2 *OtherExpense2) SearchOtherExpense2ByDocNo() error {
	//sql := `DocNo, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, DocDate, DepartCode, MyDescription, SumofAmount, AllocateCode, ProjectCode, ApCode, GLFormat,  PersonCode, BillStatus, PettyCashAmount, SumCashAmount, SumChqAmount, SumBankAmount, OtherIncome, OtherExpense, ExcessAmount1, ExcessAmount2, IsConfirm, RecurName, ConfirmCode, ConfirmDateTime, IsCancel, CancelCode, CancelDateTime from dbo.BCOTHEREXPENSE2 WITH (NOLOCK)`
	return nil
}
