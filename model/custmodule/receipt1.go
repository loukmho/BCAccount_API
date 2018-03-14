package model

type Receipt1 struct {
	DocNo           string
	TaxNo           string
	DocDate         string
	CreatorCode     string
	CreateDateTime  string
	LastEditorCode  string
	LastEditDateT   string
	ArCode          string
	SaleCode        string
	DepartCode      string
	SumOfInvoice    float64
	SumOfDebitNote  float64
	SumOfCreditNote float64
	BeforeTaxAmount float64
	TaxAmount       float64
	TotalAmount     float64
	DiscountAmount  float64
	NetAmount       float64
	SumCashAmount   float64
	SumChqAmount    float64
	SumCreditAmount float64
	ChargeAmount    float64
	ChangeAmount    float64
	SumBankAmount   float64
	SumOfWTax       float64
	GLFormat        string
	IsCancel        int
	MyDescription   string
	AllocateCode    string
	ProjectCode     string
	BillGroup       string
	IsCompleteSave  int
	SumHomeAmount1  float64
	SumHomeAmount2  float64
	DebtLimitReturn float64
	CurrencyCode    string
	ExchangeRate    float64
	OtherIncome     float64
	OtherExpense    float64
	ExcessAmount1   float64
	ExcessAmount2   float64
	IsConfirm       int
	RecurName       string
	ConfirmCode     string
	ConfirmDateTime string
	CancelCode      string
	CancelDateTime  string
}

type rcpItem struct {
	InvoiceDate     string
	InvoiceNo       string
	InvBalance      float64
	PayAmount       float64
	DebtLimitReturn float64
	MyDescription   string
	IsCancel        int
	LineNumber      int
	BillType        int
	RefNo           string
	RefType         int
	HomeAmount1     float64
	HomeAmount2     float64
}

func (rcp *Receipt1) SearchReceiptByDocNo() error {
	//sql := `DocNo, TaxNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, ArCode, SaleCode, DepartCode, SumOfInvoice, SumOfDebitNote, SumOfCreditNote, BeforeTaxAmount,TaxAmount, TotalAmount, DiscountAmount, NetAmount, SumCashAmount, SumChqAmount, SumCreditAmount, ChargeAmount, ChangeAmount, SumBankAmount, SumOfWTax, GLFormat, IsCancel,MyDescription, AllocateCode, ProjectCode, BillGroup, IsCompleteSave, SumHomeAmount1, SumHomeAmount2, DebtLimitReturn, CurrencyCode, ExchangeRate, OtherIncome,OtherExpense, ExcessAmount1, ExcessAmount2, IsConfirm, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime`
	//sqlsub:= `DocNo, DocDate, ArCode, SaleCode,  InvoiceDate, InvoiceNo, InvBalance, PayAmount, DebtLimitReturn,MyDescription,  IsCancel, LineNumber, BillType, RefNo, RefType,  HomeAmount1, HomeAmount2,`
	return nil
}
