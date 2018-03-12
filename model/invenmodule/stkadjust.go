package model

type StkAdjust struct {
	DocNo           string
	DocDate         string
	CreatorCode     string
	CreateDateTime  string
	LastEditorCode  string
	LastEditDateT   string
	cancelcode      string
	canceldatetime  string
	InspectNo       string
	MyDescription   string
	GLFormat        string
	IsConfirm       int
	IsCancel        int
	SumAmount       float64
	SumAmount2      float64
	IsCompleteSave  int
	TaxAmount       float64
	TotalAmount     float64
	SumOfExceptTax  float64
	OutputTaxStatus int
	AdjustType      int
	ConfirmCode     string
	ConfirmDateTime string
	SumOfCost       float64
	AllocateCode    string
	ProjectCode     string
	DepartCode      string
}

type AdjItem struct {
	MyType        int
	ItemCode      string
	UnitCode      string
	MyDescription string
	WHCode        string
	ShelfCode     string
	IsCancel      int
	AdjMark       int
	Qty           float64
	AdjCost       float64
	Amount        float64
	SumOfCost     float64
	ExceptTax     float64
	Price         float64
	LineNumber    int
	NetAmount     float64
	AVERAGECOST   float64
	LotNumber     string
	BarCode       string
	PackingRate1  float64
	PackingRate2  float64
}

func (sta *StkAdjust) SearchStkAdjustByDocNo() error {
	//sql:=`cancelcode, canceldatetime, DocNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, InspectNo, MyDescription, GLFormat, IsConfirm, IsCancel, SumAmount, SumAmount2, IsCompleteSave, TaxAmount, TotalAmount, SumOfExceptTax, OutputTaxStatus, AdjustType, ConfirmCode, ConfirmDateTime, SumOfCost, AllocateCode, ProjectCode, DepartCode`
	//sqlsub := `MyType, DocNo, DocDate, InspectNo, ItemCode, UnitCode, MyDescription, WHCode, ShelfCode, IsCancel, AdjMark, Qty, AdjCost, Amount, SumOfCost, ExceptTax, Price, LineNumber, NetAmount, AVERAGECOST, LotNumber,BarCode,  PackingRate1, PackingRate2`
	return nil
}
