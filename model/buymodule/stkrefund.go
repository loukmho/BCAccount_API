package model

type StkRefund struct {
	DocNo             string `json:"doc_no" db:""`
	TaxNo             string `json:"tax_no" db:""`
	DocDate           string `json:"doc_date" db:""`
	CreatorCode       string `json:"creator_code" db:""`
	CreateDateTime    string `json:"create_date_time" db:""`
	LastEditorCode    string `json:"last_editor_code" db:""`
	LastEditDateT     string `json:"last_edit_date_t" db:""`
	DueDate           string `json:"due_date" db:""`
	TaxType           int `json:"tax_type" db:""`
	ApCode            string `json:"ap_code" db:""`
	DepartCode        string `json:"depart_code" db:""`
	TaxRate           float64 `json:"tax_rate" db:""`
	IsConfirm         int `json:"is_confirm" db:""`
	MyDescription     string `json:"my_description" db:""`
	SumOfItemAmount   float64 `json:"sum_of_item_amount" db:""`
	SumOldAmount      float64 `json:"sum_old_amount" db:""`
	SumTrueAmount     float64 `json:"sum_true_amount" db:""`
	SumofDiffAmount   float64 `json:"sumof_diff_amount" db:""`
	DiscountWord      string `json:"discount_word" db:""`
	DiscountAmount    float64 `json:"discount_amount" db:""`
	SumofBeforeTax    float64 `json:"sumof_before_tax" db:""`
	SumOfTaxAmount    float64 `json:"sum_of_tax_amount" db:""`
	SumOfTotalTax     float64 `json:"sum_of_total_tax" db:""`
	SumOfExceptTax    float64 `json:"sum_of_except_tax" db:""`
	SumOfZeroTax      float64 `json:"sum_of_zero_tax" db:""`
	SumOfWTax         float64 `json:"sum_of_w_tax" db:""`
	NetDebtAmount     float64 `json:"net_debt_amount" db:""`
	SumExchangeProfit float64 `json:"sum_exchange_profit" db:""`
	BillBalance       float64 `json:"bill_balance" db:""`
	CurrencyCode      string `json:"currency_code" db:""`
	ExchangeRate      float64 `json:"exchange_rate" db:""`
	GLFormat          string `json:"gl_format" db:""`
	IsCancel          int `json:"is_cancel" db:""`
	IsCompleteSave    int `json:"is_complete_save" db:""`
	ReturnMoney       int `json:"return_money" db:""`
	BillType          int `json:"bill_type" db:""`
	CauseType         int `json:"cause_type" db:""`
	CauseCode         string
	AllocateCode      string
	ProjectCode       string
	BillGroup         string
	RecurName         string
	ConfirmCode       string
	ConfirmDateTime   string
	CancelCode        string
	CancelDateTime    string
	PayBillAmount     float64
}

type strItem struct {
	MyType         int
	ItemCode       string
	MyDescription  string
	WHCode         string
	ShelfCode      string
	DiscQty        float64
	TempQty        float64
	BillQty        float64
	Price          float64
	DiscountWord   string
	DiscountAmount float64
	Amount         float64
	NetAmount      float64
	HomeAmount     float64
	UnitCode       string
	IsCancel       int
	LineNumber     int
	RefLineNumber  int
	BarCode        string
	AVERAGECOST    float64
	SumOfCost      float64
	LotNumber      string
	ItemName       string
	PackingRate1   float64
	PackingRate2   float64
}

func (stf *StkRefund) SearchStkRefundByDocNo() error {
	//sql := `DocNo, TaxNo, DocDate,CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, DueDate, TaxType, ApCode, DepartCode, TaxRate, IsConfirm, MyDescription, SumOfItemAmount,SumOldAmount, SumTrueAmount, SumofDiffAmount, DiscountWord, DiscountAmount, SumofBeforeTax, SumOfTaxAmount, SumOfTotalTax, SumOfExceptTax, SumOfZeroTax, SumOfWTax, NetDebtAmount, SumExchangeProfit, BillBalance, CurrencyCode, ExchangeRate, GLFormat, IsCancel,  IsCompleteSave, ReturnMoney,   BillType, CauseType, CauseCode,AllocateCode, ProjectCode, BillGroup, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime, PayBillAmount`
	//sqlsub := `MyType, DocNo, TaxNo, TaxType, ItemCode, DocDate, ApCode, DepartCode, MyDescription, WHCode, ShelfCode,DiscQty, TempQty, BillQty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, UnitCode, IsCancel,  LineNumber, RefLineNumber, BarCode,AVERAGECOST, SumOfCost,LotNumber,  ItemName, TaxRate, PackingRate1, PackingRate2`
	return nil
}
