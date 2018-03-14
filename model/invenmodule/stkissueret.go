package model

type StkIssueRet struct {
	DocNo           string `json:"doc_no" db:""`
	DocDate         string `json:"doc_date" db:""`
	CreatorCode     string `json:"creator_code" db:""`
	CreateDateTime  string `json:"create_date_time" db:""`
	LastEditorCode  string `json:"last_editor_code" db:""`
	LastEditDateT   string `json:"last_edit_date_t" db:""`
	IssueRefNo      string `json:"issue_ref_no" db:""`
	DepartCode      string `json:"depart_code" db:""`
	IsConfirm       int `json:"is_confirm" db:""`
	IsCancel        int `json:"is_cancel" db:""`
	PersonCode      string `json:"person_code" db:""`
	SumOfAmount     float64 `json:"sum_of_amount" db:""`
	MyDescription   string `json:"my_description" db:""`
	GLFormat        string `json:"gl_format" db:""`
	IsCompleteSave  int `json:"is_complete_save" db:""`
	AllocateCode    string `json:"allocate_code" db:""`
	ProjectCode     string `json:"project_code" db:""`
	RecurName       string `json:"recur_name" db:""`
	ConfirmCode     string `json:"confirm_code" db:""`
	ConfirmDateTime string `json:"confirm_date_time" db:""`
	CancelCode      string `json:"cancel_code" db:""`
	CancelDateTime  string `json:"cancel_date_time" db:""`
}

type IsrItem struct {
	MyType        int `json:"my_type" db:""`
	ItemCode      string `json:"item_code" db:""`
	MyDescription string `json:"my_description" db:""`
	ItemName      string `json:"item_name" db:""`
	WHCode        string `json:"wh_code" db:""`
	ShelfCode     string `json:"shelf_code" db:""`
	Qty           float64 `json:"qty" db:""`
	Price         float64 `json:"price" db:""`
	Amount        float64 `json:"amount" db:""`
	HomeAmount    float64 `json:"home_amount" db:""`
	SumOfCost     float64 `json:"sum_of_cost" db:""`
	UnitCode      string `json:"unit_code" db:""`
	BarCode       string `json:"bar_code" db:""`
	IsCancel      int `json:"is_cancel" db:""`
	LineNumber    int `json:"line_number" db:""`
	AVERAGECOST   float64 `json:"averagecost" db:""`
	RefLinenumber int `json:"ref_linenumber" db:""`
	LotNumber     string `json:"lot_number" db:""`
	PackingRate1  float64 `json:"packing_rate_1" db:""`
	PackingRate2  float64 `json:"packing_rate_2" db:""`
}

func (str *StkIssueRet) SearchStkIssueRetByDocNo() error {
	//sql := `DocNo, DocDate,CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, IssueRefNo, DepartCode, IsConfirm, IsCancel, PersonCode, SumOfAmount, MyDescription, GLFormat,IsCompleteSave, AllocateCode, ProjectCode, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime`
	//sqlsub := `MyType, DocNo, IssueRefNo, ItemCode, DocDate, DepartCode, Personcode, MyDescription, ItemName, WHCode, ShelfCode, Qty, Price, Amount, HomeAmount, SumOfCost,UnitCode,  BarCode, IsCancel, LineNumber,  AVERAGECOST, RefLinenumber, LotNumber,PackingRate1, PackingRate2`
	return nil

}
