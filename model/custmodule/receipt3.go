package model

type Receipt3 struct {
	DocNo           string
	TaxNo           string
	DocDate         string
	ArCode          string
	CreatorCode     string
	CreateDateTime  string
	LastEditorCode  string
	LastEditDateT   string
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
	GLFormat        string
	IsCancel        int
	MyDescription   string
	AllocateCode    string
	ProjectCode     string
	BillGroup       string
	IsCompleteSave  int
	SumHomeAmount1  float64
	SumHomeAmount2  float64
	CurrencyCode    string
	ExchangeRate    float64
	IsConfirm       int
	RecurName       string
	ConfirmCode     string
	ConfirmDateTime string
	CancelCode      string
	CancelDateTime  string
}

type rcp3Item struct {
	InvoiceNo     string
	InvBalance    float64
	PayAmount     float64
	MyDescription string
	IsCancel      int
	LineNumber    int
	BillType      int
	RefNo         string
	RefType       int
	HomeAmount1   float64
	HomeAmount2   float64
}

func (rct3 *Receipt3) SearchReceipt3ByDocNo() error {
	//sql:=`DocNo, TaxNo, DocDate, ArCode,  CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, SaleCode, DepartCode, SumOfInvoice, SumOfDebitNote, SumOfCreditNote, BeforeTaxAmount,TaxAmount, TotalAmount, DiscountAmount, NetAmount, GLFormat,  IsCancel, MyDescription, AllocateCode, ProjectCode, BillGroup,  IsCompleteSave, SumHomeAmount1,SumHomeAmount2, CurrencyCode, ExchangeRate, IsConfirm, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime`
	//sqlsub := `DocNo, DocDate,ArCode, SaleCode, InvoiceDate, InvoiceNo, InvBalance, PayAmount, MyDescription, IsCancel, LineNumber, BillType, RefNo, RefType,HomeAmount1, HomeAmount2`
	return nil
}
