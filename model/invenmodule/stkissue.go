package model

type StkIssue struct {
	DocNo           string
	DocDate         string
	CreatorCode     string
	CreateDateTime  string
	LastEditorCode  string
	LastEditDateT   string
	DepartCode      string
	IssueType       int
	DayOfUse        int
	DueDate         string
	MyDescription   string
	PersonCode      string
	ArCode          string
	SumOfAmount     float64
	IsConfirm       int
	IsCancel        int
	GLFormat        string
	GLStartPosting  int
	IsPostGL        int
	IsCompleteSave  int
	BillRetStatus   int
	AllocateCode    string
	ProjectCode     string
	RecurName       string
	ConfirmCode     string
	ConfirmDateTime string
	CancelCode      string
	CancelDateTime  string
}

type IssItem struct {
	MyType        int
	ItemCode      string
	MyDescription string
	ItemName      string
	WHCode        string
	ShelfCode     string
	RefNo         string
	Qty           float64
	RetQty        float64
	SumOfCost     float64
	Price         float64
	Amount        float64
	UnitCode      string
	BarCode       string
	IsCancel      int
	LineNumber    int
	AVERAGECOST   float64
	RefLineNumber int
	LotNumber     string
	PackingRate1  float64
	PackingRate2  float64
}

func (iss *StkIssue) SearchStkIssueByDocNo() error {
	//sql := ` DocNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, DepartCode, IssueType, DayOfUse, DueDate, MyDescription, PersonCode, ArCode, SumOfAmount, IsConfirm,IsCancel, GLFormat, GLStartPosting, IsPostGL, IsCompleteSave,  BillRetStatus, AllocateCode, ProjectCode, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime, `
	//sqlsub := `MyType, DocNo,  IssueType, ItemCode, DocDate, DepartCode, Personcode, MyDescription, ItemName, WHCode,ShelfCode, RefNo, Qty, RetQty, SumOfCost, Price, Amount, UnitCode, BarCode,IsCancel, LineNumber,  AVERAGECOST, RefLineNumber, LotNumber, PackingRate1, PackingRate2`
	return nil
}
