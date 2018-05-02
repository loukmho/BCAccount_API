package model

type StkIssueRet struct {
	DocNo           string  `json:"doc_no" db:"DocNo"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	IssueRefNo      string  `json:"issue_ref_no" db:"IssueRefNo"`
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	PersonCode      string  `json:"person_code" db:"PersonCode"`
	SumOfAmount     float64 `json:"sum_of_amount" db:"SumOfAmount"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	GLFormat        string  `json:"gl_format" db:"GLFormat"`
	IsCompleteSave  int     `json:"is_complete_save" db:"IsCompleteSave"`
	AllocateCode    string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string  `json:"project_code" db:"ProjectCode"`
	RecurName       string  `json:"recur_name" db:"RecurName"`
	ConfirmCode     string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string  `json:"cancel_date_time" db:"CancelDateTime"`
}

type IsrItem struct {
	MyType        int     `json:"my_type" db:"MyType"`
	ItemCode      string  `json:"item_code" db:"ItemCode"`
	MyDescription string  `json:"my_description" db:"MyDescription"`
	ItemName      string  `json:"item_name" db:"ItemName"`
	WHCode        string  `json:"wh_code" db:"WHCode"`
	ShelfCode     string  `json:"shelf_code" db:"ShelfCode"`
	Qty           float64 `json:"qty" db:"Qty"`
	Price         float64 `json:"price" db:"Price"`
	Amount        float64 `json:"amount" db:"Amount"`
	HomeAmount    float64 `json:"home_amount" db:"HomeAmount"`
	SumOfCost     float64 `json:"sum_of_cost" db:"SumOfCost"`
	UnitCode      string  `json:"unit_code" db:"UnitCode"`
	BarCode       string  `json:"bar_code" db:"BarCode"`
	IsCancel      int     `json:"is_cancel" db:"IsCancel"`
	LineNumber    int     `json:"line_number" db:"LineNumber"`
	AVERAGECOST   float64 `json:"averagecost" db:"AVERAGECOST"`
	RefLinenumber int     `json:"ref_linenumber" db:"RefLinenumber"`
	LotNumber     string  `json:"lot_number" db:"LotNumber"`
	PackingRate1  float64 `json:"packing_rate_1" db:"PackingRate1"`
	PackingRate2  float64 `json:"packing_rate_2" db:"PackingRate2"`
}

func (str *StkIssueRet) SearchStkIssueRetByDocNo() error {
	//sql := `DocNo, DocDate,CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, IssueRefNo, DepartCode, IsConfirm, IsCancel, PersonCode, SumOfAmount, MyDescription, GLFormat,IsCompleteSave, AllocateCode, ProjectCode, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime`
	//sqlsub := `MyType, DocNo, IssueRefNo, ItemCode, DocDate, DepartCode, Personcode, MyDescription, ItemName, WHCode, ShelfCode, Qty, Price, Amount, HomeAmount, SumOfCost,UnitCode,  BarCode, IsCancel, LineNumber,  AVERAGECOST, RefLinenumber, LotNumber,PackingRate1, PackingRate2`
	return nil

}
