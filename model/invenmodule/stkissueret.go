package model

type StkIssueRet struct {
	DocNo           string
	DocDate         string
	CreatorCode     string
	CreateDateTime  string
	LastEditorCode  string
	LastEditDateT   string
	IssueRefNo      string
	DepartCode      string
	IsConfirm       int
	IsCancel        int
	PersonCode      string
	SumOfAmount     float64
	MyDescription   string
	GLFormat        string
	IsCompleteSave  int
	AllocateCode    string
	ProjectCode     string
	RecurName       string
	ConfirmCode     string
	ConfirmDateTime string
	CancelCode      string
	CancelDateTime  string
}

type IsrItem struct {
	MyType        int
	ItemCode      string
	MyDescription string
	ItemName      string
	WHCode        string
	ShelfCode     string
	Qty           float64
	Price         float64
	Amount        float64
	HomeAmount    float64
	SumOfCost     float64
	UnitCode      string
	BarCode       string
	IsCancel      int
	LineNumber    int
	AVERAGECOST   float64
	RefLinenumber int
	LotNumber     string
	PackingRate1  float64
	PackingRate2  float64
}

func (str *StkIssueRet) SearchStkIssueRetByDocNo() error {
	//sql := `DocNo, DocDate,CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, IssueRefNo, DepartCode, IsConfirm, IsCancel, PersonCode, SumOfAmount, MyDescription, GLFormat,IsCompleteSave, AllocateCode, ProjectCode, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime`
	//sqlsub := `MyType, DocNo, IssueRefNo, ItemCode, DocDate, DepartCode, Personcode, MyDescription, ItemName, WHCode, ShelfCode, Qty, Price, Amount, HomeAmount, SumOfCost,UnitCode,  BarCode, IsCancel, LineNumber,  AVERAGECOST, RefLinenumber, LotNumber,PackingRate1, PackingRate2`
	return nil

}
