package model

type FinishGoods struct {
	DocNo           string  `json:"doc_no" db:"DocNo"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	PersonCode      string  `json:"person_code" db:"PersonCode"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	SumOfAmount     float64 `json:"sum_of_amount" db:"SumOfAmount"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	GLFormat        string  `json:"gl_format" db:"GLFormat"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	IsCompleteSave  int     `json:"is_complete_save" db:"IsCompleteSave"`
	SumOfCost       float64 `json:"sum_of_cost" db:"SumOfCost"`
	AllocateCode    string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string  `json:"project_code" db:"ProjectCode"`
	RecurName       string  `json:"recur_name" db:"RecurName"`
	ConfirmCode     string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string  `json:"cancel_date_time" db:"CancelDateTime"`
}

type fngItem struct {
	MyType         int     `json:"my_type" db:"MyType"`
	ItemCode       string  `json:"item_code" db:"ItemCode"`
	MyDescription  string  `json:"my_description" db:"MyDescription"`
	ItemName       string  `json:"item_name" db:"ItemName"`
	WHCode         string  `json:"wh_code" db:"WHCode"`
	ShelfCode      string  `json:"shelf_code" db:"ShelfCode"`
	Qty            float64 `json:"qty" db:"Qty"`
	Cost           float64 `json:"cost" db:"Cost"`
	SumOfCost      float64 `json:"sum_of_cost" db:"SumOfCost"`
	Amount         float64 `json:"amount" db:"Amount"`
	UnitCode       string  `json:"unit_code" db:"UnitCode"`
	BarCode        string  `json:"bar_code" db:"BarCode"`
	IsCancel       int     `json:"is_cancel" db:"IsCancel"`
	LineNumber     int     `json:"line_number" db:"LineNumber"`
	AVERAGECOST    float64 `json:"averagecost" db:"AVERAGECOST"`
	LotNumber      string  `json:"lot_number" db:"LotNumber"`
	PackingRate1   float64 `json:"packing_rate_1" db:"PackingRate1"`
	PackingRate2   float64 `json:"packing_rate_2" db:"PackingRate2"`
	GRDocNo        string  `json:"gr_doc_no" db:"GRDocNo"`
	GRRefLineNum   int     `json:"gr_ref_line_num" db:"GRRefLineNum"`
	IssueRemainQty float64 `json:"issue_remain_qty" db:"IssueRemainQty"`
}

func (fng *FinishGoods) SearchFinishGoodsByDocNo() error {
	//sql := `DocNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, DepartCode, PersonCode, MyDescription, SumOfAmount, IsConfirm, GLFormat, IsCancel,IsCompleteSave, SumOfCost, AllocateCode, ProjectCode, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime`
	//sqlsub := `MyType, DocNo, ItemCode, DocDate,MyDescription, ItemName, WHCode, ShelfCode, Qty, Cost, SumOfCost,Amount, UnitCode, BarCode,  IsCancel, LineNumber, AVERAGECOST, LotNumber, PackingRate1, PackingRate2, GRDocNo, GRRefLineNum, IssueRemainQty`
	return nil
}
