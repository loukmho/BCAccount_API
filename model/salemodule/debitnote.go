package model

type DebitNote struct {
	DocNo             string  `json:"doc_no" db:"DocNo"`
	TaxNo             string  `json:"tax_no" db:"TaxNo"`
	DocDate           string  `json:"doc_date" db:"DocDate"`
	CreatorCode       string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime    string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode    string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT     string  `json:"last_edit_date_t" db:"LastEditDateT"`
	CreditDay         int     `json:"credit_day" db:"CreditDay"`
	DueDate           string  `json:"due_date" db:"DueDate"`
	TaxType           int     `json:"tax_type" db:"TaxType"`
	ArCode            string  `json:"ar_code" db:"ArCode"`
	DepartCode        string  `json:"depart_code" db:"DepartCode"`
	SaleCode          string  `json:"sale_code" db:"SaleCode"`
	TaxRate           float64 `json:"tax_rate" db:"TaxRate"`
	IsConfirm         int     `json:"is_confirm" db:"IsConfirm"`
	MyDescription     string  `json:"my_description" db:"MyDescription"`
	SumOfItemAmount   float64 `json:"sum_of_item_amount" db:"SumOfItemAmount"`
	SumOldAmount      float64 `json:"sum_old_amount" db:"SumOldAmount"`
	SumTrueAmount     float64 `json:"sum_true_amount" db:"SumTrueAmount"`
	SumofDiffAmount   float64 `json:"sumof_diff_amount" db:"SumofDiffAmount"`
	SumofBeforeTax    float64 `json:"sumof_before_tax" db:"SumofBeforeTax"`
	SumOfTaxAmount    float64 `json:"sum_of_tax_amount" db:"SumOfTaxAmount"`
	SumOfTotalTax     float64 `json:"sum_of_total_tax" db:"SumOfTotalTax"`
	SumOfExceptTax    float64 `json:"sum_of_except_tax" db:"SumOfExceptTax"`
	SumOfZeroTax      float64 `json:"sum_of_zero_tax" db:"SumOfZeroTax"`
	SumOfWTax         float64 `json:"sum_of_w_tax" db:"SumOfWTax"`
	DiscountWord      string  `json:"discount_word" db:"DiscountWord"`
	DiscountAmount    float64 `json:"discount_amount" db:"DiscountAmount"`
	NetDebtAmount     float64 `json:"net_debt_amount" db:"NetDebtAmount"`
	SumExchangeProfit float64 `json:"sum_exchange_profit" db:"SumExchangeProfit"`
	BillBalance       float64 `json:"bill_balance" db:"BillBalance"`
	CurrencyCode      string  `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate      float64 `json:"exchange_rate" db:"ExchangeRate"`
	GLFormat          string  `json:"gl_format" db:"GLFormat"`
	IsCancel          int     `json:"is_cancel" db:"IsCancel"`
	IsCompleteSave    int     `json:"is_complete_save" db:"IsCompleteSave"`
	ReturnStatus      int     `json:"return_status" db:"ReturnStatus"`
	CauseType         int     `json:"cause_type" db:"CauseType"`
	CauseCode         string  `json:"cause_code" db:"CauseCode"`
	PayBillStatus     int     `json:"pay_bill_status" db:""PayBillStatus`
	AllocateCode      string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode       string  `json:"project_code" db:"ProjectCode"`
	BillGroup         string  `json:"bill_group" db:"BillGroup"`
	RecurName         string  `json:"recur_name" db:"RecurName"`
	ConfirmCode       string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime   string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode        string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime    string  `json:"cancel_date_time" db:"CancelDateTime"`
	PayBillAmount     float64 `json:"pay_bill_amount" db:"PayBillAmount"`
	BillTemporary     float64 `json:"bill_temporary" db:"BillTemporary"`
}

type DbtItem struct {
	MyType         int     `json:"my_type" db:"MyType"`
	ItemCode       string  `json:"item_code" db:"ItemCode"`
	MyDescription  string  `json:"my_description" db:"MyDescription"`
	ItemName       string  `json:"item_name" db:"ItemName"`
	WHCode         string  `json:"wh_code" db:"WHCode"`
	ShelfCode      string  `json:"shelf_code" db:"ShelfCode"`
	DiscQty        float64 `json:"disc_qty" db:"DiscQty"`
	TempQty        float64 `json:"temp_qty" db:"TempQty"`
	BillQty        float64 `json:"bill_qty" db:"BillQty"`
	Price          float64 `json:"price" db:"Price"`
	DiscountWord   string  `json:"discount_word" db:"DiscountWord"`
	DiscountAmount float64 `json:"discount_amount" db:"DiscountAmount"`
	Amount         float64 `json:"amount" db:"Amount"`
	NetAmount      float64 `json:"net_amount" db:"NetAmount"`
	HomeAmount     float64 `json:"home_amount" db:"HomeAmount"`
	SumOfCost      float64 `json:"sum_of_cost" db:"SumOfCost"`
	UnitCode       string  `json:"unit_code" db:"UnitCode"`
	InvoiceNo      string  `json:"invoice_no" db:"InvoiceNo"`
	ExceptTax      int     `json:"except_tax" db:"ExceptTax"`
	IsCancel       int     `json:"is_cancel" db:"IsCancel"`
	LineNumber     int     `json:"line_number" db:"LineNumber"`
	RefLineNumber  int     `json:"ref_line_number" db:"RefLineNumber"`
	BarCode        string  `json:"bar_code" db:"BarCode"`
	AVERAGECOST    float64 `json:"averagecost" db:"AVERAGECOST"`
	LotNumber      string  `json:"lot_number" db:""`
	PackingRate1   float64 `json:"packing_rate_1" db:"PackingRate1"`
	PackingRate2   float64 `json:"packing_rate_2" db:"PackingRate2"`
}

func (dbn *DebitNote) SearchDebitNoteByDocNo() error {
	//sql := ` DocNo, TaxNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, CreditDay, DueDate, TaxType, ArCode, DepartCode, SaleCode, TaxRate, IsConfirm, MyDescription, SumOfItemAmount, SumOldAmount, SumTrueAmount, SumofDiffAmount, SumofBeforeTax, SumOfTaxAmount, SumOfTotalTax, SumOfExceptTax, SumOfZeroTax, SumOfWTax, DiscountWord, DiscountAmount, NetDebtAmount, SumExchangeProfit, BillBalance, CurrencyCode, ExchangeRate, GLFormat, GLStartPosting, IsPostGL, IsCancel,  IsCompleteSave, ReturnStatus, CauseType, CauseCode, PayBillStatus, AllocateCode, ProjectCode, BillGroup, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime, PayBillAmount, BillTemporary`

	//sqlsub := `MyType, DocNo, TaxNo,  TaxType, ItemCode, DocDate, ArCode, DepartCode, SaleCode, MyDescription, ItemName, WHCode, ShelfCode, DiscQty, TempQty, BillQty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost,UnitCode,  InvoiceNo, ExceptTax,IsCancel, LineNumber, RefLineNumber,BarCode,AVERAGECOST, LotNumber, PackingRate1, PackingRate2`
	return nil
}
