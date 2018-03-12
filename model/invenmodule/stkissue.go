package model

type StkIssue struct {
	DocNo           string  `json:"doc_no" db:""`
	DocDate         string  `json:"doc_date" db:""`
	CreatorCode     string  `json:"creator_code" db:""`
	CreateDateTime  string  `json:"create_date_time" db:""`
	LastEditorCode  string  `json:"last_editor_code" db:""`
	LastEditDateT   string  `json:"last_edit_date_t" db:""`
	DepartCode      string  `json:"depart_code" db:""`
	IssueType       int     `json:"issue_type" db:""`
	DayOfUse        int     `json:"day_of_use" db:""`
	DueDate         string  `json:"due_date" db:""`
	MyDescription   string  `json:"my_description" db:""`
	PersonCode      string  `json:"person_code" db:""`
	ArCode          string  `json:"ar_code" db:""`
	SumOfAmount     float64 `json:"sum_of_amount" db:""`
	IsConfirm       int     `json:"is_confirm" db:""`
	IsCancel        int     `json:"is_cancel" db:""`
	GLFormat        string  `json:"gl_format" db:""`
	GLStartPosting  int     `json:"gl_start_posting" db:""`
	IsPostGL        int     `json:"is_post_gl" db:""`
	IsCompleteSave  int     `json:"is_complete_save" db:""`
	BillRetStatus   int     `json:"bill_ret_status" db:""`
	AllocateCode    string  `json:"allocate_code" db:""`
	ProjectCode     string  `json:"project_code" db:""`
	RecurName       string  `json:"recur_name" db:""`
	ConfirmCode     string  `json:"confirm_code" db:""`
	ConfirmDateTime string  `json:"confirm_date_time" db:""`
	CancelCode      string  `json:"cancel_code" db:""`
	CancelDateTime  string  `json:"cancel_date_time" db:""`
}

type IssItem struct {
	MyType        int     `json:"my_type" db:""`
	ItemCode      string  `json:"item_code" db:""`
	MyDescription string  `json:"my_description" db:""`
	ItemName      string  `json:"item_name" db:""`
	WHCode        string  `json:"wh_code" db:""`
	ShelfCode     string  `json:"shelf_code" db:""`
	RefNo         string  `json:"ref_no" db:""`
	Qty           float64 `json:"qty" db:""`
	RetQty        float64 `json:"ret_qty" db:""`
	SumOfCost     float64 `json:"sum_of_cost" db:""`
	Price         float64 `json:"price" db:""`
	Amount        float64 `json:"amount" db:""`
	UnitCode      string  `json:"unit_code" db:""`
	BarCode       string  `json:"bar_code" db:""`
	IsCancel      int     `json:"is_cancel" db:""`
	LineNumber    int     `json:"line_number" db:""`
	AVERAGECOST   float64 `json:"averagecost" db:""`
	RefLineNumber int     `json:"ref_line_number" db:""`
	LotNumber     string  `json:"lot_number" db:""`
	PackingRate1  float64 `json:"packing_rate_1" db:""`
	PackingRate2  float64 `json:"packing_rate_2" db:""`
}

func (iss *StkIssue) SearchStkIssueByDocNo() error {
	//sql := ` DocNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, DepartCode, IssueType, DayOfUse, DueDate, MyDescription, PersonCode, ArCode, SumOfAmount, IsConfirm,IsCancel, GLFormat, GLStartPosting, IsPostGL, IsCompleteSave,  BillRetStatus, AllocateCode, ProjectCode, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime, `
	//sqlsub := `MyType, DocNo,  IssueType, ItemCode, DocDate, DepartCode, Personcode, MyDescription, ItemName, WHCode,ShelfCode, RefNo, Qty, RetQty, SumOfCost, Price, Amount, UnitCode, BarCode,IsCancel, LineNumber,  AVERAGECOST, RefLineNumber, LotNumber, PackingRate1, PackingRate2`
	return nil
}
