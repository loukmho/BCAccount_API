package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"errors"
)

type ArInvoice struct {
	SaveFrom        int               `json:"save_from" db:"SaveFrom"`
	Source          int               `json:"source" db:"Source"`
	DocNo           string            `json:"doc_no" db:"DocNo"`
	DocDate         string            `json:"doc_date" db:"DocDate"`
	InvOutPutTax
	InvCustomer
	InvSaleMan
	CreatorCode     string            `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string            `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string            `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string            `json:"last_edit_date_t" db:"LastEditDateT"`
	TaxType         int               `json:"tax_type" db:"TaxType"`
	DepartCode      string            `json:"depart_code" db:"DepartCode"`
	CreditDay       int               `json:"credit_day" db:"CreditDay"`
	DeliveryDay     int               `json:"delivery_day" db:"DeliveryDay"`
	DeliveryDate    string            `json:"delivery_date" db:"DeliveryDate"`
	DueDate         string            `json:"due_date" db:"DueDate"`
	PayBillDate     string            `json:"pay_bill_date" db:"PayBillDate"`
	TaxRate         int               `json:"tax_rate" db:"TaxRate"`
	IsConfirm       int               `json:"is_confirm" db:"IsConfirm"`
	MyDescription   string            `json:"my_description" db:"MyDescription"`
	BillType        int               `json:"bill_type" db:"BillType"`
	BillGroup       string            `json:"bill_group" db:"BillGroup"`
	RefDocNo        string            `json:"ref_doc_no" db:"RefDocNo"`
	DeliveryAddr    string            `json:"delivery_addr" db:"DeliveryAddr"`
	ContactCode     string            `json:"contact_code" db:""ContactCode`
	SumOfItemAmount float64           `json:"sum_of_item_amount" db:"SumOfItemAmount"`
	DiscountWord    string            `json:"discount_word" db:"DiscountWord"`
	DiscountAmount  float64           `json:"discount_amount" db:"DiscountAmount"`
	AfterDiscount   float64           `json:"after_discount" db:"AfterDiscount"`
	BeforeTaxAmount float64           `json:"before_tax_amount" db:"BeforeTaxAmount"`
	TaxAmount       float64           `json:"tax_amount" db:"TaxAmount"`
	TotalAmount     float64           `json:"total_amount" db:"TotalAmount"`
	ZeroTaxAmount   float64           `json:"zero_tax_amount" db:"ZeroTaxAmount"`
	ExceptTaxAmount float64           `json:"except_tax_amount" db:"ExceptTax"`
	SumCashAmount   float64           `json:"sum_cash_amount" db:""SumCashAmount`
	SumChqAmount    float64           `json:"sum_chq_amount" db:"SumChqAmount"`
	SumCreditAmount float64           `json:"sum_credit_amount" db:"SumCreditAmount"`
	SumBankAmount   float64           `json:"sum_bank_amount" db:"SumBankAmount"`
	DepositIncTax   float64           `json:"deposit_inc_tax" db:"DepositIncTax"`
	SumOfDeposit1   float64           `json:"sum_of_deposit_1" db:"SumOfDeposit1"`
	SumOfDeposit2   float64           `json:"sum_of_deposit_2" db:"SumOfDeposit2"`
	SumOfWTax       float64           `json:"sum_of_w_tax" db:"SumOfWTax"`
	NetDebtAmount   float64           `json:"net_debt_amount" db:"NetDebtAmount"`
	HomeAmount      float64           `json:"home_amount" db:"HomeAmount"`
	OtherIncome     float64           `json:"other_income" db:"OtherIncome"`
	OtherExpense    float64           `json:"other_expense" db:"OtherExpense"`
	ExcessAmount1   float64           `json:"excess_amount_1" db:"ExcessAmount1"`
	ExcessAmount2   float64           `json:"excess_amount_2" db:"ExcessAmount2"`
	BillBalance     float64           `json:"bill_balance" db:"BillBalance"`
	CurrencyCode    string            `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate    float64           `json:"exchange_rate" db:"ExchangeRate"`
	GLFormat        string            `json:"gl_format" db:"GLFormat"`
	IsCancel        int               `json:"is_cancel" db:"IsCancel"`
	IsCompleteSave  int               `json:"is_complete_save" db:"IsCompleteSave"`
	AllocateCode    string            `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string            `json:"project_code" db:"ProjectCode"`
	RecurName       string            `json:"recur_name" db:"RecurName"`
	ConfirmCode     string            `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string            `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string            `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string            `json:"cancel_date_time" db:"CancelDateTime"`
	IsConditionSend int               `json:"is_condition_send" db:"IsConditionSend"`
	PayBillAmount   float64           `json:"pay_bill_amount" db:"PayBillAmount"`
	SORefNo         string            `json:"so_ref_no" db:"SORe"`
	HoldingStatus   int               `json:"holding_status" db:"HoldingStatus"`
	PosStatus       int               `json:"pos_status" db:"PosStatus"`
	Pos
	RecMoney        []ListInvRecMoney `json:"rec_money"`
	Item            []InvItem         `json:"item"`
}

type InvOutPutTax struct {
	TaxNo    string `json:"tax_no" db:"TaxNo"`
	TaxDate  string `json:"tax_date" db:"TaxDate"`
	BookCode string `json:"book_code" db:"BookCode"`
}

type InvCustomer struct {
	ArCode string `json:"ar_code" db:"ArCode"`
	ArName string `json:"ar_name" db:"ArName"`
}

type InvSaleMan struct {
	SaleCode string `json:"sale_code" db:"SaleCode"`
	SaleName string `json:"sale_name" db:"SaleName"`
}

type Pos struct {
	ShiftCode        string  `json:"shiftcode" db:"ShiftCode"`
	CashierCode      string  `json:"cashier_code" db:"CashierCode"`
	ShiftNo          string  `json:"shift_no" db:"ShiftNo"`
	MachineNo        string  `json:"machine_no" db:"MachineNo"`
	MachineCode      string  `json:"machine_code" db:"MachineCode"`
	BillTime         string  `json:"bill_time" db:"BillTime"`
	CreditType       string  `json:"credit_type" db:"CreditType"`
	CreditDueDate    string  `json:"credit_due_date"  db:"CreditDueDate"`
	CreditNo         string  `json:"credit_no" db:"CreditNo"`
	CofirmNo         string  `json:"cofirm_no" db:"CofirmNo"`
	CreditBaseAmount float64 `json:"credit_base_amount" db:"CreditBaseAmount"`
	CoupongAmount    float64 `json:"coupong_amount" db:"CoupongAmount"`
	ChangeAmount     float64 `json:"change_amount" db:"ChangeAmount"`
	ChargeAmount     float64 `json:"charge_amount" db:"ChargeAmount"`
	GrandTotal       float64 `json:"grand_total" db:"GrandTotal"`
}

type InvItem struct {
	MyType          int     `json:"my_type" db:"MyType"`
	ItemCode        string  `json:"item_code" db:"ItemCode"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	ItemName        string  `json:"item_name" db:"ItemName"`
	WHCode          string  `json:"wh_code" db:"WHCode"`
	ShelfCode       string  `json:"shelf_code" db:"ShelfCode"`
	CNQty           float64 `json:"cn_qty" db:"CNQty"`
	Qty             float64 `json:"qty" db:"Qty"`
	Price           float64 `json:"price" db:"Price"`
	DiscountWord    string  `json:"discount_word" db:"DiscountWord"`
	DiscountAmount  float64 `json:"discount_amount" db:"DiscountAmount"`
	Amount          float64 `json:"amount" db:"Amount"`
	NetAmount       float64 `json:"net_amount" db:"NetAmount"`
	HomeAmount      float64 `json:"home_amount" db:"HomeAmount"`
	SumOfCost       float64 `json:"sum_of_cost" db:"SumOfCost"`
	BalanceAmount   float64 `json:"balance_amount" db:"BalanceAmount"`
	UnitCode        string  `json:"unit_code" db:"UnitCode"`
	SORefNo         string  `json:"so_ref_no" db:"SORefNo"`
	PORefNo         string  `json:"po_ref_no" db:"PORefNo"`
	LineNumber      int     `json:"line_number" db:"LineNumber"`
	RefLineNumber   int     `json:"ref_line_number" db:"RefLineNumber"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	BarCode         string  `json:"bar_code" db:"BarCode"`
	PosStatus       int     `json:"posstatus" db:"PosStatus"`
	IsConditionSend int     `json:"is_condition_send" db:"IsConditionSend"`
	AverageCost     float64 `json:"averagecost" db:"AverageCost"`
	LotNumber       string  `json:"lot_number" db:"LotNumber"`
	PackingRate1    int     `json:"packing_rate_1" db:"PackingRate1"`
	PackingRate2    int     `json:"packing_rate_2" db:"PackingRate2"`
}

type ListInvRecMoney struct {
	PayAmount      float64 `json:"pay_amount" db:"PayAmount"`
	ChqTotalAmount float64 `json:"chq_total_amount" db:"ChqTotalAmount"`
	PaymentType    int     `json:"payment_type" db:"PaymentType"`
	CreditType     string  `json:"credit_type" db:"CreditType"`
	ConfirmNo      string  `json:"confirm_no" db:"ConfirmNo"`
	RefNo          string  `json:"ref_no" db:"RefNo"`
	BankCode       string  `json:"bank_code" db:"BookCode"`
	BankBranchCode string  `json:"bank_branch_code" db:"BankBranchCode"`
	TransBankDate  string  `json:"trans_bank_date" db:"TransBankDate"`
	RefDate        string  `json:"ref_date" db:"RefDate"`
	LineNumber     int     `json:"line_number" db:"LineNumber"`
}

func (inv *ArInvoice) InsertArInvoice(db *sqlx.DB) error {
	var check_exist int
	var count_item int
	var sum_pay_amount float64

	for _, sub_recmoney := range inv.RecMoney {
		sum_pay_amount = sum_pay_amount + sub_recmoney.PayAmount
	}
	for _, sub_item := range inv.Item {
		if (sub_item.Qty != 0) {
			count_item = count_item + 1
		}
	}

	sqlexist := `select count(docno) as check_exist from dbo.bcarinvoice where docno = ? and arcode = ?`
	err := db.Get(&check_exist, sqlexist, inv.DocNo, inv.ArCode)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	switch {
	case inv.DocNo == "":
		return errors.New("docno is null")
	case inv.ArCode == "":
		return errors.New("arcode is null")
	case inv.TotalAmount == 0:
		return errors.New("totalamount is 0")
	case inv.DocDate == "":
		return errors.New("docdate is null")
	case count_item == 0:
		return errors.New("docno is not have item")
	case inv.TaxNo == "" && inv.PosStatus == 0:
		return errors.New("TaxNo is null")
	case inv.SumCashAmount == 0 && inv.SumCreditAmount == 0 && inv.SumChqAmount == 0 && inv.SumBankAmount == 0:
		return errors.New("docno not set money to another type payment")
	case sum_pay_amount > inv.TotalAmount:
		return errors.New("rec money is over totalamount")
	case check_exist > 0:
		return errors.New("docno is exist")
	case inv.SumCreditAmount != 0 && inv.PosStatus != 0 && (inv.CreditType == "" || inv.CofirmNo==""|| inv.CreditNo==""):
		return errors.New("credit card data not complete")
	case inv.PosStatus != 0 && inv.MachineCode == "" && inv.MachineNo=="" && inv.ShiftCode=="" && inv.ShiftCode=="" && inv.CashierCode == "":
		return errors.New("docno not have pos data")
	}

	if (inv.PosStatus == 0) {
		sql := `insert into dbo.bcarinvoice(DocNo,DocDate,TaxNo,ArCode,SaleCode,TaxType,DepartCode,CreditDay,DeliveryDay,DeliveryDate,DueDate,PayBillDate,TaxRate,IsConfirm,MyDescription,BillType,BillGroup,RefDocNo,DeliveryAddr,ContactCode,SumOfItemAmount,DiscountWord,DiscountAmount,AfterDiscount,BeforeTaxAmount,TaxAmount,TotalAmount,ZeroTaxAmount,ExceptTaxAmount,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,DepositIncTax,SumOfDeposit1,SumOfDeposit2,SumOfWTax,NetDebtAmount,HomeAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,BillBalance,CurrencyCode,ExchangeRate,GLFormat,IsCancel,IsCompleteSave,AllocateCode,ProjectCode,RecurName,IsConditionSend,PayBillAmount,SORefNo,HoldingStatus,PosStatus,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?,?,?,?,?,?,?,0,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,0,1,?,?,?,?,?,?,?,0,?,getdate())`
		_, err = db.Exec(sql, inv.DocNo, inv.DocDate, inv.TaxNo, inv.ArCode, inv.SaleCode, inv.TaxType, inv.DepartCode, inv.CreditDay, inv.DeliveryDay, inv.DeliveryDate, inv.DueDate, inv.PayBillDate, inv.TaxRate, inv.MyDescription, inv.BillType, inv.BillGroup, inv.RefDocNo, inv.DeliveryAddr, inv.ContactCode, inv.SumOfItemAmount, inv.DiscountWord, inv.DiscountAmount, inv.AfterDiscount, inv.BeforeTaxAmount, inv.TaxAmount, inv.TotalAmount, inv.ZeroTaxAmount, inv.ExceptTaxAmount, inv.SumCashAmount, inv.SumChqAmount, inv.SumCreditAmount, inv.SumBankAmount, inv.DepositIncTax, inv.SumOfDeposit1, inv.SumOfDeposit2, inv.SumOfWTax, inv.NetDebtAmount, inv.HomeAmount, inv.OtherIncome, inv.OtherExpense, inv.ExcessAmount1, inv.ExcessAmount2, inv.BillBalance, inv.CurrencyCode, inv.ExchangeRate, inv.GLFormat, inv.AllocateCode, inv.ProjectCode, inv.RecurName, inv.IsConditionSend, inv.PayBillAmount, inv.SORefNo, inv.HoldingStatus, inv.CreatorCode)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}else{
		sql := `insert into dbo.bcarinvoice(DocNo,DocDate,TaxNo,ArCode,SaleCode,TaxType,DepartCode,CreditDay,DeliveryDay,DeliveryDate,DueDate,PayBillDate,TaxRate,IsConfirm,MyDescription,BillType,BillGroup,RefDocNo,DeliveryAddr,ContactCode,SumOfItemAmount,DiscountWord,DiscountAmount,AfterDiscount,BeforeTaxAmount,TaxAmount,TotalAmount,ZeroTaxAmount,ExceptTaxAmount,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,DepositIncTax,SumOfDeposit1,SumOfDeposit2,SumOfWTax,NetDebtAmount,HomeAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,BillBalance,CurrencyCode,ExchangeRate,GLFormat,IsCancel,IsCompleteSave,AllocateCode,ProjectCode,RecurName,IsConditionSend,PayBillAmount,SORefNo,HoldingStatus,PosStatus,CreatorCode,CreateDateTime,ShiftCode,CashierCode,ShiftNo,MachineNo,MachineCode,BillTime,CreditType,CreditDueDate,CreditNo,CofirmNo,CreditBaseAmount,CoupongAmount,ChangeAmount,ChargeAmount,GrandTotal) values(?,?,?,?,?,?,?,?,?,?,?,?,?,0,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,0,1,?,?,?,?,?,?,?,0,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sql, inv.DocNo, inv.DocDate, inv.TaxNo, inv.ArCode, inv.SaleCode, inv.TaxType, inv.DepartCode, inv.CreditDay, inv.DeliveryDay, inv.DeliveryDate, inv.DueDate, inv.PayBillDate, inv.TaxRate, inv.MyDescription, inv.BillType, inv.BillGroup, inv.RefDocNo, inv.DeliveryAddr, inv.ContactCode, inv.SumOfItemAmount, inv.DiscountWord, inv.DiscountAmount, inv.AfterDiscount, inv.BeforeTaxAmount, inv.TaxAmount, inv.TotalAmount, inv.ZeroTaxAmount, inv.ExceptTaxAmount, inv.SumCashAmount, inv.SumChqAmount, inv.SumCreditAmount, inv.SumBankAmount, inv.DepositIncTax, inv.SumOfDeposit1, inv.SumOfDeposit2, inv.SumOfWTax, inv.NetDebtAmount, inv.HomeAmount, inv.OtherIncome, inv.OtherExpense, inv.ExcessAmount1, inv.ExcessAmount2, inv.BillBalance, inv.CurrencyCode, inv.ExchangeRate, inv.GLFormat, inv.AllocateCode, inv.ProjectCode, inv.RecurName, inv.IsConditionSend, inv.PayBillAmount, inv.SORefNo, inv.HoldingStatus, inv.CreatorCode,inv.ShiftCode,inv.CashierCode,inv.ShiftNo,inv.MachineNo,inv.MachineCode,inv.BillTime,inv.CreditType,inv.CreditDueDate,inv.CreditNo,inv.CofirmNo,inv.CreditBaseAmount,inv.CoupongAmount,inv.ChangeAmount,inv.ChargeAmount,inv.GrandTotal)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	inv.SaveFrom = 0
	inv.Source = 6
	inv.BookCode = "20"

	sqltax := `insert into dbo.BCOutputTax(SaveFrom,DocNo,BookCode,Source,DocDate,TaxDate,TaxNo,ArCode,ShortTaxDesc,TaxRate,Process,BeforeTaxAmount,TaxAmount,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?,?,'ขายสินค้า',?,1,?,?,?,getdate())`
	_, err = db.Exec(sqltax, inv.SaveFrom, inv.DocNo, inv.BookCode, inv.Source, inv.DocDate, inv.TaxDate, inv.TaxNo, inv.ArCode, inv.TaxRate, inv.BeforeTaxAmount, inv.TaxAmount, inv.CreatorCode)
	fmt.Println("sqltax = ", sqltax)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	for _, item := range inv.Item {
		item.MyType = 4
		item.CNQty = item.Qty
		item.HomeAmount = item.NetAmount
		item.IsCancel = 0
		item.PackingRate2 = 1
		item.PosStatus = 0

		sqlsub := `insert into dbo.BCArInvoiceSub(MyType,DocNo, TaxNo, TaxType, ItemCode, DocDate, ArCode, DepartCode, SaleCode, MyDescription, ItemName, WHCode, ShelfCode, CNQty, Qty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost, BalanceAmount, UnitCode, SORefNo, PORefNo, LineNumber, RefLineNumber, IsCancel, BarCode, POSSTATUS,IsConditionSend, AVERAGECOST, LotNumber, PackingRate1, PackingRate2) values(MyType,DocNo, TaxNo, TaxType, ItemCode, DocDate, ArCode, DepartCode, SaleCode, MyDescription, ItemName, WHCode, ShelfCode, CNQty, Qty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost, BalanceAmount, UnitCode, SORefNo, PORefNo, LineNumber, RefLineNumber, IsCancel, BarCode, POSSTATUS,IsConditionSend, AVERAGECOST, LotNumber, PackingRate1, PackingRate2)`
		db.Exec(sqlsub, item.MyType, inv.DocNo, inv.TaxNo, inv.TaxType, item.ItemCode, inv.DocDate, inv.ArCode, inv.DepartCode, inv.SaleCode, item.MyDescription, item.ItemName, item.WHCode, item.ShelfCode, item.CNQty, item.Qty, item.Price, item.DiscountWord, item.DiscountAmount, item.Amount, item.NetAmount, item.HomeAmount, item.SumOfCost, item.BalanceAmount, item.UnitCode, item.SORefNo, item.PORefNo, item.LineNumber, item.RefLineNumber, item.IsCancel, item.BarCode, item.PosStatus, item.IsConditionSend, item.AverageCost, item.LotNumber, item.PackingRate1, item.PackingRate2)
		fmt.Println("sqltax = ", sqltax)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}

	return nil
}

func (inv *ArInvoice)SearchArInvoiceByDocNo(db *sqlx.DB, docno string)error {
	sql := `select DocNo,DocDate,isnull(TaxNo,'') as TaxNo,ArCode,SaleCode,TaxType,isnull(DepartCode,'') as DepartCode,CreditDay,DeliveryDay,isnull(DeliveryDate,'') as DeliveryDate,isnull(DueDate,'') as DueDate,isnull(PayBillDate,'') as PayBillDate,TaxRate,IsConfirm,isnull(MyDescription,'') as MyDescription,BillType,isnull(BillGroup,'') as BillGroup,isnull(RefDocNo,'') as RefDocNo,isnull(DeliveryAddr,'') as DeliveryAddr,isnull(ContactCode,'') as ContactCode,SumOfItemAmount,DiscountWord,isnull(DiscountAmount,'') as DiscountAmount,AfterDiscount,BeforeTaxAmount,TaxAmount,TotalAmount,ZeroTaxAmount,ExceptTaxAmount,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,DepositIncTax,SumOfDeposit1,SumOfDeposit2,SumOfWTax,NetDebtAmount,HomeAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,BillBalance,CurrencyCode,ExchangeRate,isnull(GLFormat,'') as GLFormat,IsCancel,IsCompleteSave,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(RecurName,'') as RecurName,IsConditionSend,PayBillAmount,isnull(SORefNo,'') as SORefNo,HoldingStatus,PosStatus,CreatorCode,CreateDateTime,isnull(ShiftCode,'') as ShiftCode,isnull(CashierCode,'') as CashierCode,isnull(ShiftNo,'') as ShiftNo,isnull(MachineNo,'') as MachineNo,isnull(MachineCode,'') as MachineCode,isnull(BillTime,'') as BillTime,isnull(CreditType,'') as CreditType,isnull(CreditDueDate,'') as CreditDueDate,isnull(CreditNo,'') as CreditNo,isnull(CofirmNo,'') as CofirmNo,CreditBaseAmount,CoupongAmount,ChangeAmount,ChargeAmount,GrandTotal from dbo.bcarinvoice where docno = ? `
	err := db.Get(inv, sql, docno)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	sqlsub := `select select MyType, ItemCode, isnull(MyDescription,'') as MyDescription, ItemName, WHCode, ShelfCode, CNQty, Qty, Price, isnull(DiscountWord,'') as DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost, BalanceAmount, UnitCode, isnull(SORefNo,'') as SORefNo, isnull(PORefNo,'') as PORefNo, LineNumber, RefLineNumber, IsCancel, isnull(BarCode,'') as BarCode, IsConditionSend, AVERAGECOST, isnull(LotNumber,'') as LotNumber, PackingRate1, PackingRate2 from dbo.bcarinvoicesub from dbo.bcarinvoicesub where docno = ? order by linenumber`
	err = db.Select(&inv.Item, sqlsub, docno)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
