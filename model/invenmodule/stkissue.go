package model

type StkIssue struct {
	DocNo           string  `json:"doc_no" db:"DocNo"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	IssueType       int     `json:"issue_type" db:"IssueType"`
	DayOfUse        int     `json:"day_of_use" db:"DayOfUse"`
	DueDate         string  `json:"due_date" db:"DueDate"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	PersonCode      string  `json:"person_code" db:"PersonCode"`
	ArCode          string  `json:"ar_code" db:"ArCode"`
	SumOfAmount     float64 `json:"sum_of_amount" db:"SumOfAmount"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	GLFormat        string  `json:"gl_format" db:"GLFormat"`
	GLStartPosting  int     `json:"gl_start_posting" db:"GLStartPosting"`
	IsPostGL        int     `json:"is_post_gl" db:"IsPostGL"`
	IsCompleteSave  int     `json:"is_complete_save" db:"IsCompleteSave"`
	BillRetStatus   int     `json:"bill_ret_status" db:"BillRetStatus"`
	AllocateCode    string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string  `json:"project_code" db:"ProjectCode"`
	RecurName       string  `json:"recur_name" db:"RecurName"`
	ConfirmCode     string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string  `json:"cancel_date_time" db:"CancelDateTime"`
}

type IssItem struct {
	MyType        int     `json:"my_type" db:"MyType"`
	ItemCode      string  `json:"item_code" db:"ItemCode"`
	MyDescription string  `json:"my_description" db:"MyDescription"`
	ItemName      string  `json:"item_name" db:"ItemName"`
	WHCode        string  `json:"wh_code" db:"WHCode"`
	ShelfCode     string  `json:"shelf_code" db:"ShelfCode"`
	RefNo         string  `json:"ref_no" db:"RefNo"`
	Qty           float64 `json:"qty" db:"Qty"`
	RetQty        float64 `json:"ret_qty" db:"RetQty"`
	SumOfCost     float64 `json:"sum_of_cost" db:"SumOfCost"`
	Price         float64 `json:"price" db:"Price"`
	Amount        float64 `json:"amount" db:"Amount"`
	UnitCode      string  `json:"unit_code" db:"UnitCode"`
	BarCode       string  `json:"bar_code" db:"BarCode"`
	IsCancel      int     `json:"is_cancel" db:"IsCancel"`
	LineNumber    int     `json:"line_number" db:"LineNumber"`
	AVERAGECOST   float64 `json:"averagecost" db:"AVERAGECOST"`
	RefLineNumber int     `json:"ref_line_number" db:"RefLineNumber"`
	LotNumber     string  `json:"lot_number" db:"LotNumber"`
	PackingRate1  float64 `json:"packing_rate_1" db:"PackingRate1"`
	PackingRate2  float64 `json:"packing_rate_2" db:"PackingRate2"`
}

func (iss *StkIssue) SearchStkIssueByDocNo() error {
	//sql := ` DocNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, DepartCode, IssueType, DayOfUse, DueDate, MyDescription, PersonCode, ArCode, SumOfAmount, IsConfirm,IsCancel, GLFormat, GLStartPosting, IsPostGL, IsCompleteSave,  BillRetStatus, AllocateCode, ProjectCode, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime, `
	//sqlsub := `MyType, DocNo,  IssueType, ItemCode, DocDate, DepartCode, Personcode, MyDescription, ItemName, WHCode,ShelfCode, RefNo, Qty, RetQty, SumOfCost, Price, Amount, UnitCode, BarCode,IsCancel, LineNumber,  AVERAGECOST, RefLineNumber, LotNumber, PackingRate1, PackingRate2`
	return nil
}
