package model

type StkAdjust struct {
	DocNo           string  `json:"doc_no" db:""`
	DocDate         string  `json:"doc_date" db:""`
	CreatorCode     string  `json:"creator_code" db:""`
	CreateDateTime  string  `json:"create_date_time" db:""`
	LastEditorCode  string  `json:"last_editor_code" db:""`
	LastEditDateT   string  `json:"last_edit_date_t" db:""`
	cancelcode      string  `json:"cancelcode" db:""`
	canceldatetime  string  `json:"canceldatetime" db:""`
	InspectNo       string  `json:"inspect_no" db:""`
	MyDescription   string  `json:"my_description" db:""`
	GLFormat        string  `json:"gl_format" db:""`
	IsConfirm       int     `json:"is_confirm" db:""`
	IsCancel        int     `json:"is_cancel" db:""`
	SumAmount       float64 `json:"sum_amount" db:""`
	SumAmount2      float64 `json:"sum_amount_2" db:""`
	IsCompleteSave  int     `json:"is_complete_save" db:""`
	TaxAmount       float64 `json:"tax_amount" db:""`
	TotalAmount     float64 `json:"total_amount" db:""`
	SumOfExceptTax  float64 `json:"sum_of_except_tax" db:""`
	OutputTaxStatus int     `json:"output_tax_status" db:""`
	AdjustType      int     `json:"adjust_type" db:""`
	ConfirmCode     string  `json:"confirm_code" db:""`
	ConfirmDateTime string  `json:"confirm_date_time" db:""`
	SumOfCost       float64 `json:"sum_of_cost" db:""`
	AllocateCode    string  `json:"allocate_code" db:""`
	ProjectCode     string  `json:"project_code" db:""`
	DepartCode      string  `json:"depart_code" db:""`
}

type AdjItem struct {
	MyType        int     `json:"my_type" db:""`
	ItemCode      string  `json:"item_code" db:""`
	UnitCode      string  `json:"unit_code" db:""`
	MyDescription string  `json:"my_description" db:""`
	WHCode        string  `json:"wh_code" db:""`
	ShelfCode     string  `json:"shelf_code" db:""`
	IsCancel      int     `json:"is_cancel" db:""`
	AdjMark       int     `json:"adj_mark" db:""`
	Qty           float64 `json:"qty" db:""`
	AdjCost       float64 `json:"adj_cost" db:""`
	Amount        float64 `json:"amount" db:""`
	SumOfCost     float64 `json:"sum_of_cost" db:""`
	ExceptTax     float64 `json:"except_tax" db:""`
	Price         float64 `json:"price" db:""`
	LineNumber    int     `json:"line_number" db:""`
	NetAmount     float64 `json:"net_amount" db:""`
	AVERAGECOST   float64 `json:"averagecost" db:""`
	LotNumber     string  `json:"lot_number" db:""`
	BarCode       string  `json:"bar_code" db:""`
	PackingRate1  float64 `json:"packing_rate_1" db:""`
	PackingRate2  float64 `json:"packing_rate_2" db:""`
}

func (sta *StkAdjust) SearchStkAdjustByDocNo() error {
	//sql:=`cancelcode, canceldatetime, DocNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, InspectNo, MyDescription, GLFormat, IsConfirm, IsCancel, SumAmount, SumAmount2, IsCompleteSave, TaxAmount, TotalAmount, SumOfExceptTax, OutputTaxStatus, AdjustType, ConfirmCode, ConfirmDateTime, SumOfCost, AllocateCode, ProjectCode, DepartCode`
	//sqlsub := `MyType, DocNo, DocDate, InspectNo, ItemCode, UnitCode, MyDescription, WHCode, ShelfCode, IsCancel, AdjMark, Qty, AdjCost, Amount, SumOfCost, ExceptTax, Price, LineNumber, NetAmount, AVERAGECOST, LotNumber,BarCode,  PackingRate1, PackingRate2`
	return nil
}
