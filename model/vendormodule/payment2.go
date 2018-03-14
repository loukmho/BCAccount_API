package model

type PaymentSub2 struct {
	DocNo           string
	TaxNo           string
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
	DiscountAmount  float64
	NetAmount       float64
	GLFormat        string
	IsCancel        int
	MyDescription   string
	AllocateCode    string
	ProjectCode     string
	BillGroup       string
	SumHomeAmount1  float64
	SumHomeAmount2  float64
	CurrencyCode    string
	ExchangeRate    float64
	IsCompleteSave  int
	IsConfirm       int
	RecurName       string
	ConfirmCode     string
	ConfirmDateTime string
	CancelCode      string
	CancelDateTime  string
}

type pmt2Item struct {
	InvoiceDate   string
	InvoiceNo     string
	InvBalance    float64
	PayAmount     float64
	MyDescription string
	BillType      int
	PaybillNo     string
	RefType       int
	IsCancel      int
	LineNumber    int
	HomeAmount1   float64
	HomeAmount2   float64
}
