package model

type FinishGoods struct {
	DocNo           string
	DocDate         string
	CreatorCode     string
	CreateDateTime  string
	LastEditorCode  string
	LastEditDateT   string
	DepartCode      string
	PersonCode      string
	MyDescription   string
	SumOfAmount     float64
	IsConfirm       int
	GLFormat        string
	IsCancel        int
	IsCompleteSave  int
	SumOfCost       float64
	AllocateCode    string
	ProjectCode     string
	RecurName       string
	ConfirmCode     string
	ConfirmDateTime string
	CancelCode      string
	CancelDateTime  string
}

type fngItem struct {
	MyType         int
	ItemCode       string
	MyDescription  string
	ItemName       string
	WHCode         string
	ShelfCode      string
	Qty            float64
	Cost           float64
	SumOfCost      float64
	Amount         float64
	UnitCode       string
	BarCode        string
	IsCancel       int
	LineNumber     int
	AVERAGECOST    float64
	LotNumber      string
	PackingRate1   float64
	PackingRate2   float64
	GRDocNo        string
	GRRefLineNum   int
	IssueRemainQty float64
}

func (fng *FinishGoods) SearchFinishGoodsByDocNo() error {
	//sql := `DocNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, DepartCode, PersonCode, MyDescription, SumOfAmount, IsConfirm, GLFormat, IsCancel,IsCompleteSave, SumOfCost, AllocateCode, ProjectCode, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime`
	//sqlsub := `MyType, DocNo, ItemCode, DocDate,MyDescription, ItemName, WHCode, ShelfCode, Qty, Cost, SumOfCost,Amount, UnitCode, BarCode,  IsCancel, LineNumber, AVERAGECOST, LotNumber, PackingRate1, PackingRate2, GRDocNo, GRRefLineNum, IssueRemainQty`
	return nil
}
