package model

var DebitNote struct {
	DocNo             string
	TaxNo             string
	DocDate           string
	CreatorCode       string
	CreateDateTime    string
	LastEditorCode    string
	LastEditDateT     string
	CreditDay         int
	DueDate           string
	TaxType           int
	ArCode            string
	DepartCode        string
	SaleCode          string
	TaxRate           float64
	IsConfirm         int
	MyDescription     string
	SumOfItemAmount   float64
	SumOldAmount      float64
	SumTrueAmount     float64
	SumofDiffAmount   float64
	SumofBeforeTax    float64
	SumOfTaxAmount    float64
	SumOfTotalTax     float64
	SumOfExceptTax    float64
	SumOfZeroTax      float64
	SumOfWTax         float64
	DiscountWord      string
	DiscountAmount    float64
	NetDebtAmount     float64
	SumExchangeProfit float64
	BillBalance       float64
	CurrencyCode      string
	ExchangeRate      float64
	GLFormat          string
	GLStartPosting    int
	IsPostGL          int
	IsCancel          int
	IsCompleteSave    int
	ReturnStatus      int
	CauseType         int
	CauseCode         string
	PayBillStatus     int
	AllocateCode      string
	ProjectCode       string
	BillGroup         string
	RecurName         string
	ConfirmCode       string
	ConfirmDateTime   string
	CancelCode        string
	CancelDateTime    string
	PayBillAmount     float64
	BillTemporary     float64
}

type DbtItem struct {
	MyType         int
	ItemCode       string
	MyDescription  string
	ItemName       string
	WHCode         string
	ShelfCode      string
	DiscQty        float64
	TempQty        float64
	BillQty        float64
	Price          float64
	DiscountWord   string
	DiscountAmount float64
	Amount         float64
	NetAmount      float64
	HomeAmount     float64
	SumOfCost      float64
	UnitCode       string
	InvoiceNo      string
	ExceptTax      int
	IsCancel       int
	LineNumber     int
	RefLineNumber  int
	BarCode        string
	AVERAGECOST    float64
	LotNumber      string
	PackingRate1   float64
	PackingRate2   float64
}

func (dbn *DebitNote) SearchDebitNoteByDocNo() error {
	//sql := ` DocNo, TaxNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, CreditDay, DueDate, TaxType, ArCode, DepartCode, SaleCode, TaxRate, IsConfirm, MyDescription, SumOfItemAmount, SumOldAmount, SumTrueAmount, SumofDiffAmount, SumofBeforeTax, SumOfTaxAmount, SumOfTotalTax, SumOfExceptTax, SumOfZeroTax, SumOfWTax, DiscountWord, DiscountAmount, NetDebtAmount, SumExchangeProfit, BillBalance, CurrencyCode, ExchangeRate, GLFormat, GLStartPosting, IsPostGL, IsCancel,  IsCompleteSave, ReturnStatus, CauseType, CauseCode, PayBillStatus, AllocateCode, ProjectCode, BillGroup, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime, PayBillAmount, BillTemporary`

	//sqlsub := `MyType, DocNo, TaxNo,  TaxType, ItemCode, DocDate, ArCode, DepartCode, SaleCode, MyDescription, ItemName, WHCode, ShelfCode, DiscQty, TempQty, BillQty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost,UnitCode,  InvoiceNo, ItemType, ExceptTax, ExchangeRate,IsCancel, LineNumber, RefLineNumber,BarCode,AVERAGECOST, LotNumber,TaxRate, PackingRate1, PackingRate2`
	return nil
}
