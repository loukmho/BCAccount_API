package model

type Statement struct {
	DocNo           string
	DocDate         string
	ApCode          string
	CreatorCode     string
	CreateDateTime  string
	LastEditorCode  string
	LastEditDateT   string
	DepartCode      string
	SumOfInvoice    float64
	SumOfDebitNote  float64
	SumOfCreditNote float64
	BeforeTaxAmount float64
	TaxAmount       float64
	TotalAmount     float64
	MyDescription   string
	BillGroup       string
	BillStatus      int
	CreditDay       int
	IsCompleteSave  int
	DueDate         string
	IsConfirm       int
	RecurName       string
	ConfirmCode     string
	ConfirmDateTime string
	CancelCode      string
	CancelDateTime  string
	IsCancel        int
}

type stmItem struct {
	InvoiceDate   string
	InvoiceNo     string
	InvBalance    float64
	PayAmount     float64
	PayBalance    float64
	MyDescription string
	IsCancel      int
	LineNumber    int
	BillType      int
	HomeAmount    float64
}

func (stm *Statement) SearchStatementByDocNo() error {
	//sql := `DocNo, DocDate, ApCode, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, DepartCode, SumOfInvoice, SumOfDebitNote, SumOfCreditNote, BeforeTaxAmount, TaxAmount,TotalAmount, MyDescription, BillGroup, BillStatus, CreditDay, IsCompleteSave, DueDate, IsConfirm, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime, IsCancel`
	//sqlsub := `DocNo, DocDate, ApCode, InvoiceDate, InvoiceNo, InvBalance, PayAmount, PayBalance, MyDescription, IsCancel,LineNumber, BillType, HomeAmount`
	return nil
}
