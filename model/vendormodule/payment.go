package model

type Payment struct {
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
	PettyCashAmount float64
	SumCashAmount   float64
	SumChqAmount    float64
	SumBankAmount   float64
	SumOfWTax       float64
	GLFormat        string
	IsCancel        int
	MyDescription   string
	AllocateCode    string
	ProjectCode     string
	BillGroup       string
	SumHomeAmount1  float64
	SumHomeAmount2  float64
	OtherIncome     float64
	OtherExpense    float64
	ExcessAmount1   float64
	ExcessAmount2   float64
	CurrencyCode    string
	ExchangeRate    float64
	IsCompleteSave  int
	IsConfirm       int
	RecurName       string
	ConfirmCode     string
	ConfirmDateTime string
	CancelCode      string
	CancelDateTime  string
	SumOfDeposit1   float64
	SumOfDeposit2   float64
	HomeAmountInv   float64
	HomeAmountDebt  float64
	HomeAmountCrd   float64
}

type pmtItem struct {
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

func (pmt *Payment) SearchPaymentByDocNo() error {
	//sql := `DocNo, TaxNo, DocDate, ApCode, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, DepartCode, SumOfInvoice, SumOfDebitNote, SumOfCreditNote, BeforeTaxAmount, TaxAmount,TotalAmount, DiscountAmount, NetAmount, PettyCashAmount, SumCashAmount, SumChqAmount, SumBankAmount, SumOfWTax, GLFormat, IsCancel, MyDescription, AllocateCode,ProjectCode, BillGroup, SumHomeAmount1, SumHomeAmount2, OtherIncome, OtherExpense, ExcessAmount1, ExcessAmount2, CurrencyCode, ExchangeRate, IsCompleteSave, IsConfirm,RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime, SumOfDeposit1, SumOfDeposit2, HomeAmountInv, HomeAmountDebt, HomeAmountCrd`
	//sqlsub:=`DocNo, DocDate, ApCode, InvoiceDate, InvoiceNo, InvBalance, PayAmount, MyDescription, BillType, PaybillNo, RefType,IsCancel, LineNumber, HomeAmount1, HomeAmount2`
	return nil
}

//paymentsub
