package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type DebitNote2 struct {
	SaveFrom          int     `json:"save_from" db:"SaveFrom"`
	Source            int     `json:"source" db:"Source"`
	DocNo             string  `json:"doc_no" db:"DocNo"`
	DbnInPutTax
	DbnVendor
	DocDate           string  `json:"doc_date" db:"DocDate"`
	CreatorCode       string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime    string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode    string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT     string  `json:"last_edit_date_t" db:"LastEditDateT"`
	CreditDay         int     `json:"credit_day" db:"CreditDay"`
	DueDate           string  `json:"due_date" db:"DueDate"`
	TaxType           int     `json:"tax_type" db:"TaxType"`
	DepartCode        string  `json:"depart_code" db:"DepartCode"`
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
	AllocateCode      string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode       string  `json:"project_code" db:"ProjectCode"`
	BillGroup         string  `json:"bill_group" db:"BillGroup"`
	RecurName         string  `json:"recur_name" db:"RecurName"`
	ConfirmCode       string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime   string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode        string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime    string  `json:"cancel_date_time" db:"CancelDateTime"`
	PayBillAmount     float64 `json:"pay_bill_amount" db:"PayBillAmount"`
	UserCode          string     `json:"user_code" db:"UserCode"`
	Subs              []*DbnItem `json:"subs"`
}

type DbnVendor struct {
	ApCode string `json:"ap_code" db:"ApCode"`
	ApName string `json:"ap_name" db:"ApName"`
}

type DbnInPutTax struct {
	TaxNo    string `json:"tax_no" db:"TaxNo"`
	BookCode string `json:"book_code" db:"BookCode"`
}

type DbnItem struct {
	MyType         int     `json:"my_type" db:"MyType"`
	ItemCode       string  `json:"item_code" db:"ItemCode"`
	ItemType       int     `json:"item_type" db:"ItemType"`
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
	LotNumber      string  `json:"lot_number" db:"LotNumber"`
	PackingRate1   float64 `json:"packing_rate_1" db:"PackingRate1"`
	PackingRate2   float64 `json:"packing_rate_2" db:"PackingRate2"`
}

func InsertAndEditDebitNote2() {
	//sql := `insert into dbo.bcdebitnote2(DocNo, TaxNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, CreditDay, DueDate, TaxType, ApCode, DepartCode, TaxRate, IsConfirm, MyDescription, SumOfItemAmount, SumOldAmount, SumTrueAmount, SumofDiffAmount, SumofBeforeTax, SumOfTaxAmount, SumOfTotalTax, SumOfExceptTax, SumOfZeroTax, SumOfWTax, DiscountWord, DiscountAmount, NetDebtAmount, SumExchangeProfit, BillBalance, CurrencyCode, ExchangeRate, GLFormat, GLStartPosting, IsPostGL, IsCancel,  IsCompleteSave, ReturnStatus, CauseType, CauseCode, AllocateCode, ProjectCode, BillGroup, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime, PayBillAmount)`
	//sqlsub := `insert into dbo.BCDebitNoteSub2(MyType, DocNo, TaxNo,  TaxType, ItemCode, ItemType, DocDate, ApCode, DepartCode, MyDescription, ItemName, WHCode, ShelfCode, DiscQty, TempQty, BillQty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost,UnitCode,  InvoiceNo, ExceptTax,IsCancel, LineNumber, RefLineNumber,BarCode,AVERAGECOST, LotNumber, PackingRate1, PackingRate2) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
}

func (dbn *DebitNote2) SearchDebitNote2ByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select DocNo, isnull(a.TaxNo,'') as TaxNo, DocDate, isnull(a.CreatorCode,'') as CreatorCode, isnull(a.CreateDateTime,'') as CreateDateTime, isnull(a.LastEditorCode,'') as LastEditorCode, isnull(a.LastEditDateT,'') as LastEditDateT, a.CreditDay, isnull(DueDate,'') as DueDate, a.TaxType, isnull(ApCode,'') as ApCode, isnull(b.name1,'') as ApName, isnull(DepartCode,'') as DepartCode, TaxRate, IsConfirm, isnull(MyDescription,'') as MyDescription, SumOfItemAmount, SumOldAmount, SumTrueAmount, SumofDiffAmount, SumofBeforeTax, SumOfTaxAmount, SumOfTotalTax, SumOfExceptTax, SumOfZeroTax, SumOfWTax, isnull(DiscountWord,'') as DiscountWord, DiscountAmount, NetDebtAmount, SumExchangeProfit, BillBalance, isnull(CurrencyCode,'') as CurrencyCode, ExchangeRate, isnull(GLFormat,'') as GLFormat, GLStartPosting, IsPostGL, IsCancel, IsCompleteSave, ReturnStatus, isnull(CauseType,0) as CauseType, isnull(CauseCode,'') as CauseCode, isnull(AllocateCode,'') as AllocateCode, isnull(ProjectCode,'') as ProjectCode, isnull(BillGroup,'') as BillGroup, isnull(RecurName,'') as RecurName, isnull(a.ConfirmCode,'') as ConfirmCode, isnull(a.ConfirmDateTime,'') as ConfirmDateTime, isnull(a.CancelCode,'') as CancelCode, isnull(a.CancelDateTime,'') as CancelDateTime, PayBillAmount  from dbo.bcdebitnote2 a left join dbo.bcap b with (nolock) on a.apcode = b.code where a.docno = ? order by a.docno`
	err := db.Get(dbn, sql, docno)
	if err != nil {
		return nil
	}

	sqlsub := `set dateformat dmy     select MyType, DocNo, isnull(TaxNo,'') as TaxNo, isnull(TaxType,0) as TaxType, isnull(ItemCode,'') as ItemCode, isnull(ItemType,0) as ItemType, DocDate, ApCode, isnull(DepartCode,'') as DepartCode, isnull(SaleCode,'') as SaleCode, isnull(MyDescription,'') as MyDescription, isnull(ItemName,'') as ItemName, isnull(WHCode,'') as WHCode, isnull(ShelfCode,'') as ShelfCode, isnull(DiscQty,0) as DiscQty, isnull(TempQty,0) as TempQty, isnull(BillQty,0) as BillQty, isnull(Price,0) as Price, isnull(DiscountWord,'') as DiscountWord, isnull(DiscountAmount,0) as DiscountAmount, isnull(Amount,0) as Amount, isnull(NetAmount,0) as NetAmount, isnull(HomeAmount,0) as HomeAmount, isnull(SumOfCost,0) as SumOfCost, isnull(UnitCode,'') as UnitCode, isnull(InvoiceNo,'') as InvoiceNo, isnull(IsCancel,0) as IsCancel, isnull(LineNumber,0) as LineNumber, isnull(RefLineNumber,0) as RefLineNumber, isnull(BarCode,'') as BarCode, isnull(AVERAGECOST,0) as AverageCost, isnull(LotNumber,'') as LotNumber, isnull(PackingRate1,1) as PackingRate1, isnull(PackingRate2,1) as PackingRate2 from dbo.BCDebitNoteSub2 a With (Nolock) where a.docno = ?`
	err = db.Select(&dbn.Subs, sqlsub, docno)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//sqlsub := `MyType, DocNo, TaxNo,  TaxType, ItemCode, ItemType, DocDate, ApCode, DepartCode, MyDescription, ItemName, WHCode, ShelfCode, DiscQty, TempQty, BillQty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost,UnitCode,  InvoiceNo, ExceptTax,IsCancel, LineNumber, RefLineNumber,BarCode,AVERAGECOST, LotNumber, PackingRate1, PackingRate2`
	return nil
}
