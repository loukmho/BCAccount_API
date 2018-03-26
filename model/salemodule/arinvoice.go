package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"errors"
	m "github.com/loukmho/bcaccount_api/model"
	"time"
)

type ArInvoice struct {
	SaveFrom         int                  `json:"save_from" db:"SaveFrom"`
	Source           int                  `json:"source" db:"Source"`
	DocNo            string               `json:"doc_no" db:"DocNo"`
	DocDate          string               `json:"doc_date" db:"DocDate"`
	InvOutPutTax
	InvCustomer
	InvSaleMan
	CreatorCode      string               `json:"creator_code" db:"CreatorCode"`
	CreateDateTime   string               `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode   string               `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT    string               `json:"last_edit_date_t" db:"LastEditDateT"`
	TaxType          int                  `json:"tax_type" db:"TaxType"`
	DepartCode       string               `json:"depart_code" db:"DepartCode"`
	CreditDay        int                  `json:"credit_day" db:"CreditDay"`
	DeliveryDay      int                  `json:"delivery_day" db:"DeliveryDay"`
	DeliveryDate     string               `json:"delivery_date" db:"DeliveryDate"`
	DueDate          string               `json:"due_date" db:"DueDate"`
	PayBillDate      string               `json:"pay_bill_date" db:"PayBillDate"`
	TaxRate          float64              `json:"tax_rate" db:"TaxRate"`
	IsConfirm        int                  `json:"is_confirm" db:"IsConfirm"`
	MyDescription    string               `json:"my_description" db:"MyDescription"`
	BillType         int                  `json:"bill_type" db:"BillType"`
	BillGroup        string               `json:"bill_group" db:"BillGroup"`
	RefDocNo         string               `json:"ref_doc_no" db:"RefDocNo"`
	DeliveryAddr     string               `json:"delivery_addr" db:"DeliveryAddr"`
	ContactCode      string               `json:"contact_code" db:"ContactCode"`
	SumOfItemAmount  float64              `json:"sum_of_item_amount" db:"SumOfItemAmount"`
	DiscountWord     string               `json:"discount_word" db:"DiscountWord"`
	DiscountAmount   float64              `json:"discount_amount" db:"DiscountAmount"`
	AfterDiscount    float64              `json:"after_discount" db:"AfterDiscount"`
	BeforeTaxAmount  float64              `json:"before_tax_amount" db:"BeforeTaxAmount"`
	TaxAmount        float64              `json:"tax_amount" db:"TaxAmount"`
	TotalAmount      float64              `json:"total_amount" db:"TotalAmount"`
	ZeroTaxAmount    float64              `json:"zero_tax_amount" db:"ZeroTaxAmount"`
	ExceptTaxAmount  float64              `json:"except_tax_amount" db:"ExceptTaxAmount"`
	SumCashAmount    float64              `json:"sum_cash_amount" db:"SumCashAmount"`
	SumChqAmount     float64              `json:"sum_chq_amount" db:"SumChqAmount"`
	SumCreditAmount  float64              `json:"sum_credit_amount" db:"SumCreditAmount"`
	SumBankAmount    float64              `json:"sum_bank_amount" db:"SumBankAmount"`
	DepositIncTax    int                  `json:"deposit_inc_tax" db:"DepositIncTax"`
	SumOfDeposit1    float64              `json:"sum_of_deposit_1" db:"SumOfDeposit1"`
	SumOfDeposit2    float64              `json:"sum_of_deposit_2" db:"SumOfDeposit2"`
	SumOfWTax        float64              `json:"sum_of_w_tax" db:"SumOfWTax"`
	NetDebtAmount    float64              `json:"net_debt_amount" db:"NetDebtAmount"`
	HomeAmount       float64              `json:"home_amount" db:"HomeAmount"`
	OtherIncome      float64              `json:"other_income" db:"OtherIncome"`
	OtherExpense     float64              `json:"other_expense" db:"OtherExpense"`
	ExcessAmount1    float64              `json:"excess_amount_1" db:"ExcessAmount1"`
	ExcessAmount2    float64              `json:"excess_amount_2" db:"ExcessAmount2"`
	BillBalance      float64              `json:"bill_balance" db:"BillBalance"`
	CurrencyCode     string               `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate     float64              `json:"exchange_rate" db:"ExchangeRate"`
	GLFormat         string               `json:"gl_format" db:"GLFormat"`
	IsCancel         int                  `json:"is_cancel" db:"IsCancel"`
	IsCompleteSave   int                  `json:"is_complete_save" db:"IsCompleteSave"`
	AllocateCode     string               `json:"allocate_code" db:"AllocateCode"`
	ProjectCode      string               `json:"project_code" db:"ProjectCode"`
	RecurName        string               `json:"recur_name" db:"RecurName"`
	ConfirmCode      string               `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime  string               `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode       string               `json:"cancel_code" db:"CancelCode"`
	CancelDateTime   string               `json:"cancel_date_time" db:"CancelDateTime"`
	IsConditionSend  int                  `json:"is_condition_send" db:"IsConditionSend"`
	PayBillAmount    float64              `json:"pay_bill_amount" db:"PayBillAmount"`
	SORefNo          string               `json:"so_ref_no" db:"SORefNo"`
	HoldingStatus    int                  `json:"holding_status" db:"HoldingStatus"`
	PosStatus        int                  `json:"pos_status" db:"PosStatus"`
	CreditType       string               `json:"credit_type" db:"CreditType"`
	CreditDueDate    string               `json:"credit_due_date"  db:"CreditDueDate"`
	CreditNo         string               `json:"credit_no" db:"CreditNo"`
	CofirmNo         string               `json:"cofirm_no" db:"CofirmNo"`
	CreditBaseAmount float64              `json:"credit_base_amount" db:"CreditBaseAmount"`
	UserCode         string               `json:"user_code" db:"UserCode"`
	Pos
	ListInvRecMoney
	Subs             []*InvItem           `json:"subs"`
	Deps             []*ListArDepositUsed `json:"deps"`
	Cdcs             []*ListCreditCard    `json:"cdcs"`
	Chqs             []*ListChqIn         `json:"chqs"`
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
	ShiftCode     string  `json:"shiftcode" db:"ShiftCode"`
	CashierCode   string  `json:"cashier_code" db:"CashierCode"`
	ShiftNo       string  `json:"shift_no" db:"ShiftNo"`
	MachineNo     string  `json:"machine_no" db:"MachineNo"`
	MachineCode   string  `json:"machine_code" db:"MachineCode"`
	BillTime      string  `json:"bill_time" db:"BillTime"`
	CoupongAmount float64 `json:"coupong_amount" db:"CoupongAmount"`
	ChangeAmount  float64 `json:"change_amount" db:"ChangeAmount"`
	ChargeAmount  float64 `json:"charge_amount" db:"ChargeAmount"`
	GrandTotal    float64 `json:"grand_total" db:"GrandTotal"`
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
	PackingRate1    float64 `json:"packing_rate_1" db:"PackingRate1"`
	PackingRate2    float64 `json:"packing_rate_2" db:"PackingRate2"`
}

type ListInvRecMoney struct {
	CreditTypeSub  string `json:"credit_type_sub" db:"CreditTypeSub"`
	ConfirmNo      string `json:"confirm_no" db:"ConfirmNo"`
	CreditRefNo    string `json:"credit_ref_no" db:"CreditRefNo"`
	BankCode       string `json:"bank_code" db:"BookCode"`
	BankBranchCode string `json:"bank_branch_code" db:"BankBranchCode"`
	TransBankDate  string `json:"trans_bank_date" db:"TransBankDate"`
	RefDate        string `json:"ref_date" db:"RefDate"`
}

type ListArDepositUsed struct {
	DepositNo      string `json:"deposit_no" db:"DepositNo"`
	Balance        float64 `json:"balance" db:"Balance"`
	Amount         float64 `json:"amount" db:"Amount"`
	NetAmount      float64 `json:"net_amount" db:"NetAmount"`
	LineNumber     int `json:"line_number" db:"LineNumber"`
}

type ListChqIn struct {
	ChqNumber      string
	BankCode       string
	BankBranchCode string
	BookNo         string
	ReceiveDate    string
	DueDate        string
	Status         int
	Amount         float64
	Balance        float64
	RefChqRowOrder int
	SaveFrom       int
	StatusDate     string
	StatusDocNo    string
}

type ListCreditCard struct {
}

//type ListInvRecMoney struct {
//	PayAmount      float64 `json:"pay_amount" db:"PayAmount"`
//	ChqTotalAmount float64 `json:"chq_total_amount" db:"ChqTotalAmount"`
//	PaymentType    int     `json:"payment_type" db:"PaymentType"`
//	CreditType     string  `json:"credit_type" db:"CreditType"`
//	ConfirmNo      string  `json:"confirm_no" db:"ConfirmNo"`
//	RefNo          string  `json:"ref_no" db:"RefNo"`
//	BankCode       string  `json:"bank_code" db:"BookCode"`
//	BankBranchCode string  `json:"bank_branch_code" db:"BankBranchCode"`
//	TransBankDate  string  `json:"trans_bank_date" db:"TransBankDate"`
//	RefDate        string  `json:"ref_date" db:"RefDate"`
//	LineNumber     int     `json:"line_number" db:"LineNumber"`
//}

func (inv *ArInvoice) InsertArInvoice(db *sqlx.DB) error {
	var check_exist int
	var count_item int
	var sum_pay_amount float64

	now := time.Now()

	for _, sub_item := range inv.Subs {
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
	case inv.DocDate == "":
		return errors.New("docdate is null")
	case inv.IsCancel == 1:
		return errors.New("docno is cancel")
	case inv.IsConfirm == 1:
		return errors.New("docno is confirm")
	case count_item == 0:
		return errors.New("docno is not have item")
	case (inv.SumCashAmount == 0 && inv.SumCreditAmount == 0 && inv.SumChqAmount == 0 && inv.SumBankAmount == 0 && inv.SumOfDeposit1 == 0 && inv.BillType == 0):
		return errors.New("docno not set money to another type payment")
	case inv.SumOfItemAmount == 0:
		return errors.New("sumofitemamount = 0")
	case sum_pay_amount > inv.TotalAmount:
		return errors.New("rec money is over totalamount")
	case inv.SumCreditAmount != 0 && inv.PosStatus != 0 && (inv.CreditType == "" || inv.CofirmNo == "" || inv.CreditNo == ""):
		return errors.New("credit card data not complete")
	case inv.PosStatus != 0 && inv.MachineCode == "" && inv.MachineNo == "" && inv.ShiftCode == "" && inv.ShiftCode == "" && inv.CashierCode == "":
		return errors.New("docno not have pos data")
	case inv.SumOfDeposit1 != 0 && len(inv.Deps) == 0:
		return errors.New("docno not have deposit use data")
	}

	inv.SaveFrom = 0
	inv.DepositIncTax = 1

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")

	fmt.Println("TaxRate = ", def.TaxRateDefault)

	if inv.TaxRate == 0 {
		inv.TaxRate = def.TaxRateDefault
	}
	if inv.ExchangeRate == 0 {
		inv.ExchangeRate = def.ExchangeRateDefault
	}

	inv.BeforeTaxAmount, inv.TaxAmount, inv.TotalAmount = m.CalcTaxItem(inv.TaxType, inv.TaxRate, inv.AfterDiscount)

	if (((inv.SumCreditAmount + inv.SumCashAmount + inv.SumBankAmount + inv.SumChqAmount) != inv.TotalAmount) && inv.BillType == 0) {
		return errors.New("receive not equal totalamount")
	}

	if (inv.BookCode == "") {
		inv.BookCode = def.ArInvoiceBookCode
	}
	if (inv.Source == 0) {
		inv.Source = def.ArInvoiceSource
	}
	if (inv.BillType == 0 && inv.GLFormat == "") {
		inv.GLFormat = def.ArInvoiceCashGLFormat
	} else {
		inv.GLFormat = def.ArInvoiceCreditGLFormat
	}
	if (inv.TaxNo == "" && inv.PosStatus == 0) {
		inv.TaxNo = inv.DocNo
	}
	if (inv.TaxDate == "" && inv.PosStatus == 0) {
		inv.TaxDate = inv.DocDate
	}
	if (inv.CreditDay == 0 || inv.DueDate == "") {
		inv.DueDate = inv.DocDate
	} else {
		inv.DueDate = now.AddDate(0, 0, inv.CreditDay).Format("2006-01-02")
	}
	if (inv.DeliveryDay == 0 && inv.DeliveryDate == "") {
		inv.DeliveryDate = inv.DocDate
	} else {
		inv.DeliveryDate = now.AddDate(0, 0, inv.DeliveryDay).Format("2006-01-02")
	}

	fmt.Println("UserCode = ", inv.UserCode)

	if (inv.IsCompleteSave == 0) {
		inv.IsCompleteSave = 1
	}

	if (check_exist == 0) {
		//Insert/////////////////////////////////////////////////////////////////////////////////////////////////////////
		inv.CreatorCode = inv.UserCode

		if (inv.PosStatus == 0) {
			sql := `insert into dbo.bcarinvoice(DocNo,DocDate,TaxNo,ArCode,SaleCode,TaxType,DepartCode,CreditDay,DeliveryDay,DeliveryDate,DueDate,PayBillDate,TaxRate,IsConfirm,MyDescription,BillType,BillGroup,RefDocNo,DeliveryAddr,ContactCode,SumOfItemAmount,DiscountWord,DiscountAmount,AfterDiscount,BeforeTaxAmount,TaxAmount,TotalAmount,ZeroTaxAmount,ExceptTaxAmount,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,DepositIncTax,SumOfDeposit1,SumOfDeposit2,SumOfWTax,NetDebtAmount,HomeAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,BillBalance,CurrencyCode,ExchangeRate,GLFormat,IsCancel,IsCompleteSave,AllocateCode,ProjectCode,RecurName,IsConditionSend,PayBillAmount,SORefNo,HoldingStatus,PosStatus,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?,?,?,?,?,?,?,0,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,0,1,?,?,?,?,?,?,?,0,?,getdate())`
			_, err = db.Exec(sql, inv.DocNo, inv.DocDate, inv.TaxNo, inv.ArCode, inv.SaleCode, inv.TaxType, inv.DepartCode, inv.CreditDay, inv.DeliveryDay, inv.DeliveryDate, inv.DueDate, inv.PayBillDate, inv.TaxRate, inv.MyDescription, inv.BillType, inv.BillGroup, inv.RefDocNo, inv.DeliveryAddr, inv.ContactCode, inv.SumOfItemAmount, inv.DiscountWord, inv.DiscountAmount, inv.AfterDiscount, inv.BeforeTaxAmount, inv.TaxAmount, inv.TotalAmount, inv.ZeroTaxAmount, inv.ExceptTaxAmount, inv.SumCashAmount, inv.SumChqAmount, inv.SumCreditAmount, inv.SumBankAmount, inv.DepositIncTax, inv.SumOfDeposit1, inv.SumOfDeposit2, inv.SumOfWTax, inv.NetDebtAmount, inv.HomeAmount, inv.OtherIncome, inv.OtherExpense, inv.ExcessAmount1, inv.ExcessAmount2, inv.BillBalance, inv.CurrencyCode, inv.ExchangeRate, inv.GLFormat, inv.AllocateCode, inv.ProjectCode, inv.RecurName, inv.IsConditionSend, inv.PayBillAmount, inv.SORefNo, inv.HoldingStatus, inv.CreatorCode)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
		} else {
			sql := `insert into dbo.bcarinvoice(DocNo,DocDate,TaxNo,ArCode,SaleCode,TaxType,DepartCode,CreditDay,DeliveryDay,DeliveryDate,DueDate,PayBillDate,TaxRate,IsConfirm,MyDescription,BillType,BillGroup,RefDocNo,DeliveryAddr,ContactCode,SumOfItemAmount,DiscountWord,DiscountAmount,AfterDiscount,BeforeTaxAmount,TaxAmount,TotalAmount,ZeroTaxAmount,ExceptTaxAmount,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,DepositIncTax,SumOfDeposit1,SumOfDeposit2,SumOfWTax,NetDebtAmount,HomeAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,BillBalance,CurrencyCode,ExchangeRate,GLFormat,IsCancel,IsCompleteSave,AllocateCode,ProjectCode,RecurName,IsConditionSend,PayBillAmount,SORefNo,HoldingStatus,PosStatus,CreatorCode,CreateDateTime,ShiftCode,CashierCode,ShiftNo,MachineNo,MachineCode,BillTime,CreditType,CreditDueDate,CreditNo,CofirmNo,CreditBaseAmount,CoupongAmount,ChangeAmount,ChargeAmount,GrandTotal) values(?,?,?,?,?,?,?,?,?,?,?,?,?,0,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,0,1,?,?,?,?,?,?,?,0,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
			_, err = db.Exec(sql, inv.DocNo, inv.DocDate, inv.TaxNo, inv.ArCode, inv.SaleCode, inv.TaxType, inv.DepartCode, inv.CreditDay, inv.DeliveryDay, inv.DeliveryDate, inv.DueDate, inv.PayBillDate, inv.TaxRate, inv.MyDescription, inv.BillType, inv.BillGroup, inv.RefDocNo, inv.DeliveryAddr, inv.ContactCode, inv.SumOfItemAmount, inv.DiscountWord, inv.DiscountAmount, inv.AfterDiscount, inv.BeforeTaxAmount, inv.TaxAmount, inv.TotalAmount, inv.ZeroTaxAmount, inv.ExceptTaxAmount, inv.SumCashAmount, inv.SumChqAmount, inv.SumCreditAmount, inv.SumBankAmount, inv.DepositIncTax, inv.SumOfDeposit1, inv.SumOfDeposit2, inv.SumOfWTax, inv.NetDebtAmount, inv.HomeAmount, inv.OtherIncome, inv.OtherExpense, inv.ExcessAmount1, inv.ExcessAmount2, inv.BillBalance, inv.CurrencyCode, inv.ExchangeRate, inv.GLFormat, inv.AllocateCode, inv.ProjectCode, inv.RecurName, inv.IsConditionSend, inv.PayBillAmount, inv.SORefNo, inv.HoldingStatus, inv.CreatorCode, inv.ShiftCode, inv.CashierCode, inv.ShiftNo, inv.MachineNo, inv.MachineCode, inv.BillTime, inv.CreditType, inv.CreditDueDate, inv.CreditNo, inv.CofirmNo, inv.CreditBaseAmount, inv.CoupongAmount, inv.ChangeAmount, inv.ChargeAmount, inv.GrandTotal)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
		}

		sqltax := `insert into dbo.BCOutputTax(SaveFrom,DocNo,BookCode,Source,DocDate,TaxDate,TaxNo,ArCode,ShortTaxDesc,TaxRate,Process,BeforeTaxAmount,TaxAmount,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?,?,'ขายสินค้า',?,1,?,?,?,getdate())`
		_, err = db.Exec(sqltax, inv.SaveFrom, inv.DocNo, inv.BookCode, inv.Source, inv.DocDate, inv.TaxDate, inv.TaxNo, inv.ArCode, inv.TaxRate, inv.BeforeTaxAmount, inv.TaxAmount, inv.CreatorCode)
		fmt.Println("sqltax = ", sqltax)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	} else {
		//Update/////////////////////////////////////////////////////////////////////////////////////////////////////////
		inv.LastEditorCode = inv.UserCode

		if (inv.PosStatus == 0) {
			sql := `update dbo.bcarinvoice set DocDate=?,TaxNo=?,ArCode=?,SaleCode=?,TaxType=?,DepartCode=?,CreditDay=?,DeliveryDay=?,DeliveryDate=?,DueDate=?,PayBillDate=?,TaxRate=?,IsConfirm=?,MyDescription=?,BillType=?,BillGroup=?,RefDocNo=?,DeliveryAddr=?,ContactCode=?,SumOfItemAmount=?,DiscountWord=?,DiscountAmount=?,AfterDiscount=?,BeforeTaxAmount=?,TaxAmount=?,TotalAmount=?,ZeroTaxAmount=?,ExceptTaxAmount=?,SumCashAmount=?,SumChqAmount=?,SumCreditAmount=?,SumBankAmount=?,DepositIncTax=?,SumOfDeposit1=?,SumOfDeposit2=?,SumOfWTax=?,NetDebtAmount=?,HomeAmount=?,OtherIncome=?,OtherExpense=?,ExcessAmount1=?,ExcessAmount2=?,BillBalance=?,CurrencyCode=?,ExchangeRate=?,GLFormat=?,IsCancel=?,AllocateCode=?,ProjectCode=?,RecurName=?,IsConditionSend=?,PayBillAmount=?,SORefNo=?,HoldingStatus=?,PosStatus=?,LastEditorCode=?,LastEditDateT=getdate() where docno = ?`
			_, err = db.Exec(sql, inv.DocDate, inv.TaxNo, inv.ArCode, inv.SaleCode, inv.TaxType, inv.DepartCode, inv.CreditDay, inv.DeliveryDay, inv.DeliveryDate, inv.DueDate, inv.PayBillDate, inv.TaxRate, inv.IsConfirm, inv.MyDescription, inv.BillType, inv.BillGroup, inv.RefDocNo, inv.DeliveryAddr, inv.ContactCode, inv.SumOfItemAmount, inv.DiscountWord, inv.DiscountAmount, inv.AfterDiscount, inv.BeforeTaxAmount, inv.TaxAmount, inv.TotalAmount, inv.ZeroTaxAmount, inv.ExceptTaxAmount, inv.SumCashAmount, inv.SumChqAmount, inv.SumCreditAmount, inv.SumBankAmount, inv.DepositIncTax, inv.SumOfDeposit1, inv.SumOfDeposit2, inv.SumOfWTax, inv.NetDebtAmount, inv.HomeAmount, inv.OtherIncome, inv.OtherExpense, inv.ExcessAmount1, inv.ExcessAmount2, inv.BillBalance, inv.CurrencyCode, inv.ExchangeRate, inv.GLFormat, inv.IsCancel, inv.AllocateCode, inv.ProjectCode, inv.RecurName, inv.IsConditionSend, inv.PayBillAmount, inv.SORefNo, inv.HoldingStatus, inv.PosStatus, inv.LastEditorCode, inv.DocNo)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
		} else {
			sql := `insert into dbo.bcarinvoice(DocNo,DocDate,TaxNo,ArCode,SaleCode,TaxType,DepartCode,CreditDay,DeliveryDay,DeliveryDate,DueDate,PayBillDate,TaxRate,IsConfirm,MyDescription,BillType,BillGroup,RefDocNo,DeliveryAddr,ContactCode,SumOfItemAmount,DiscountWord,DiscountAmount,AfterDiscount,BeforeTaxAmount,TaxAmount,TotalAmount,ZeroTaxAmount,ExceptTaxAmount,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,DepositIncTax,SumOfDeposit1,SumOfDeposit2,SumOfWTax,NetDebtAmount,HomeAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,BillBalance,CurrencyCode,ExchangeRate,GLFormat,IsCancel,IsCompleteSave,AllocateCode,ProjectCode,RecurName,IsConditionSend,PayBillAmount,SORefNo,HoldingStatus,PosStatus,CreatorCode,CreateDateTime,ShiftCode,CashierCode,ShiftNo,MachineNo,MachineCode,BillTime,CreditType,CreditDueDate,CreditNo,CofirmNo,CreditBaseAmount,CoupongAmount,ChangeAmount,ChargeAmount,GrandTotal) values(?,?,?,?,?,?,?,?,?,?,?,?,?,0,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,0,1,?,?,?,?,?,?,?,0,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
			_, err = db.Exec(sql, inv.DocNo, inv.DocDate, inv.TaxNo, inv.ArCode, inv.SaleCode, inv.TaxType, inv.DepartCode, inv.CreditDay, inv.DeliveryDay, inv.DeliveryDate, inv.DueDate, inv.PayBillDate, inv.TaxRate, inv.MyDescription, inv.BillType, inv.BillGroup, inv.RefDocNo, inv.DeliveryAddr, inv.ContactCode, inv.SumOfItemAmount, inv.DiscountWord, inv.DiscountAmount, inv.AfterDiscount, inv.BeforeTaxAmount, inv.TaxAmount, inv.TotalAmount, inv.ZeroTaxAmount, inv.ExceptTaxAmount, inv.SumCashAmount, inv.SumChqAmount, inv.SumCreditAmount, inv.SumBankAmount, inv.DepositIncTax, inv.SumOfDeposit1, inv.SumOfDeposit2, inv.SumOfWTax, inv.NetDebtAmount, inv.HomeAmount, inv.OtherIncome, inv.OtherExpense, inv.ExcessAmount1, inv.ExcessAmount2, inv.BillBalance, inv.CurrencyCode, inv.ExchangeRate, inv.GLFormat, inv.AllocateCode, inv.ProjectCode, inv.RecurName, inv.IsConditionSend, inv.PayBillAmount, inv.SORefNo, inv.HoldingStatus, inv.CreatorCode, inv.ShiftCode, inv.CashierCode, inv.ShiftNo, inv.MachineNo, inv.MachineCode, inv.BillTime, inv.CreditType, inv.CreditDueDate, inv.CreditNo, inv.CofirmNo, inv.CreditBaseAmount, inv.CoupongAmount, inv.ChangeAmount, inv.ChargeAmount, inv.GrandTotal)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
		}
	}

	sql_del_sub := `delete dbo.bcarinvoicesub where docno = ?`
	_, err = db.Exec(sql_del_sub, inv.DocNo)
	if err != nil {
		return err
	}

	var vLineNumber int

	for _, item := range inv.Subs {
		fmt.Println("ItemSub")
		item.MyType = 4
		item.CNQty = item.Qty
		item.BalanceAmount = item.Amount
		item.LineNumber = vLineNumber

		item.IsCancel = 0
		if (item.PackingRate1 == 0) {
			item.PackingRate1 = 1
		}
		if (item.PackingRate2 == 0) {
			item.PackingRate2 = 1
		}

		item.PosStatus = inv.PosStatus

		switch {
		case inv.TaxType == 0:
			item.HomeAmount = item.Amount
			item.NetAmount = item.Amount
		case inv.TaxType == 1:
			taxamount := m.ToFixed(item.Amount-((item.Amount*100)/(100+float64(inv.TaxRate))), 2)
			beforetaxamount := m.ToFixed(item.Amount-taxamount, 2)
			item.HomeAmount = beforetaxamount
			item.NetAmount = beforetaxamount
		case inv.TaxType == 2:
			item.HomeAmount = item.Amount
			item.NetAmount = item.Amount
		}

		sqlsub := ` insert into dbo.BCArInvoiceSub(MyType,DocNo, TaxNo, TaxType, ItemCode, DocDate, ArCode, DepartCode, SaleCode, MyDescription, ItemName, WHCode, ShelfCode, CNQty, Qty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost, BalanceAmount, UnitCode, SORefNo, PORefNo, LineNumber, RefLineNumber, IsCancel, BarCode, POSSTATUS,IsConditionSend, AVERAGECOST, LotNumber, PackingRate1, PackingRate2) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		_, err = db.Exec(sqlsub, item.MyType, inv.DocNo, inv.TaxNo, inv.TaxType, item.ItemCode, inv.DocDate, inv.ArCode, inv.DepartCode, inv.SaleCode, item.MyDescription, item.ItemName, item.WHCode, item.ShelfCode, item.CNQty, item.Qty, item.Price, item.DiscountWord, item.DiscountAmount, item.Amount, item.NetAmount, item.HomeAmount, item.SumOfCost, item.BalanceAmount, item.UnitCode, item.SORefNo, item.PORefNo, item.LineNumber, item.RefLineNumber, item.IsCancel, item.BarCode, item.PosStatus, item.IsConditionSend, item.AverageCost, item.LotNumber, item.PackingRate1, item.PackingRate2)
		fmt.Println("sqltax = ", sqlsub)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		vLineNumber = vLineNumber + 1
	}

	sqldel := `delete dbo.BCOutputTax where docno = ?`
	_, err = db.Exec(sqldel, inv.DocNo)
	if err != nil {
		return err
	}

	sqltax := `insert into dbo.BCOutputTax(SaveFrom,DocNo,BookCode,Source,DocDate,TaxDate,TaxNo,ArCode,ShortTaxDesc,TaxRate,Process,BeforeTaxAmount,TaxAmount,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?,?,'ขายสินค้า',?,1,?,?,?,getdate())`
	_, err = db.Exec(sqltax, inv.SaveFrom, inv.DocNo, inv.BookCode, inv.Source, inv.DocDate, inv.TaxDate, inv.TaxNo, inv.ArCode, inv.TaxRate, inv.BeforeTaxAmount, inv.TaxAmount, inv.CreatorCode)
	fmt.Println("sqltax = ", sqltax)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	sqlrecdel := `delete dbo.BCRecMoney where docno = ?`
	_, err = db.Exec(sqlrecdel, inv.DocNo)
	if err != nil {
		return err
	}

	var my_description_recmoney string

	if (inv.MyDescription == "") {
		my_description_recmoney = "ขายเงินสด"
	} else {
		my_description_recmoney = inv.MyDescription
	}

	fmt.Println("RecMoney")
	var linenumber int

	if (inv.BillType == 0) {
		if (inv.SumCashAmount != 0) { //subs.PaymentType == 0:
			fmt.Println("SumCashAmount")
			sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,PaymentType,LineNumber,ProjectCode,DepartCode,SaleCode,MyDescription) values(?,?,?,?,?,?,?,?,?,?,?)`
			_, err = db.Exec(sqlrec, inv.DocNo, inv.DocDate, inv.ArCode, inv.ExchangeRate, inv.SumCashAmount, 0, linenumber, inv.ProjectCode, inv.DepartCode, inv.SaleCode, my_description_recmoney)
			if err != nil {
				return err
			}
		}
		//case dp.SumCreditAmount != 0: //subs.PaymentType == 1:
		if (inv.SumCreditAmount != 0) {
			fmt.Println("SumCreditAmount")
			if (inv.SumCashAmount != 0) {
				linenumber = 1
			} else {
				linenumber = 0
			}
			sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,ChqTotalAmount,PaymentType,CreditType,ConfirmNo,LineNumber,RefNo,BankCode,ProjectCode,DepartCode,SaleCode,MyDescription,RefDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
			_, err = db.Exec(sqlrec, inv.DocNo, inv.DocDate, inv.ArCode, inv.ExchangeRate, inv.SumCreditAmount, inv.SumCreditAmount, 1, inv.CreditType, inv.ConfirmNo, linenumber, inv.CreditRefNo, inv.BankCode, inv.ProjectCode, inv.DepartCode, inv.SaleCode, my_description_recmoney, inv.DocDate)
			if err != nil {
				return err
			}
		}

		//case dp.SumChqAmount != 0: //subs.PaymentType == 2:
		if (inv.SumChqAmount != 0) {
			fmt.Println("SumChqAmount")
			if (inv.SumCashAmount != 0 && inv.SumCreditAmount != 0) {
				linenumber = 2
			} else if ((inv.SumCashAmount != 0 && inv.SumCreditAmount == 0)) {
				linenumber = 1
			} else if ((inv.SumCashAmount == 0 && inv.SumCreditAmount != 0)) {
				linenumber = 1
			} else if ((inv.SumCashAmount == 0 && inv.SumCreditAmount == 0)) {
				linenumber = 0
			}

			sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,PaymentType,LineNumber,RefNo,BankCode,ProjectCode,DepartCode,SaleCode,BankBranchCode,TransBankDate,MyDescription,RefDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
			_, err = db.Exec(sqlrec, inv.DocNo, inv.DocDate, inv.ArCode, inv.ExchangeRate, inv.SumChqAmount, 2, linenumber, inv.CreditRefNo, inv.BankCode, inv.ProjectCode, inv.DepartCode, inv.SaleCode, inv.BankBranchCode, inv.TransBankDate, my_description_recmoney, inv.RefDate)
			if err != nil {
				return err
			}
		}

		//case dp.SumBankAmount != 0: //subs.PaymentType == 3:
		if (inv.SumBankAmount != 0) {
			fmt.Println("SumBankAmount")
			if (inv.SumCashAmount != 0 && inv.SumCreditAmount != 0 && inv.SumChqAmount != 0) {
				linenumber = 3
			} else if (inv.SumCashAmount != 0 && inv.SumCreditAmount == 0 && inv.SumChqAmount != 0) {
				linenumber = 2
			} else if (inv.SumCashAmount == 0 && inv.SumCreditAmount != 0 && inv.SumChqAmount != 0) {
				linenumber = 2
			} else if (inv.SumCashAmount != 0 && inv.SumCreditAmount != 0 && inv.SumChqAmount == 0) {
				linenumber = 2
			} else if (inv.SumCashAmount != 0 && inv.SumCreditAmount != 0 && inv.SumChqAmount == 0) {
				linenumber = 2
			} else if (inv.SumCashAmount != 0 && inv.SumCreditAmount == 0 && inv.SumChqAmount == 0) {
				linenumber = 1
			} else if (inv.SumCashAmount == 0 && inv.SumCreditAmount != 0 && inv.SumChqAmount == 0) {
				linenumber = 1
			} else if (inv.SumCashAmount == 0 && inv.SumCreditAmount == 0 && inv.SumChqAmount != 0) {
				linenumber = 1
			} else if (inv.SumCashAmount == 0 && inv.SumCreditAmount == 0 && inv.SumChqAmount == 0) {
				linenumber = 0
			}

			sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,PaymentType,LineNumber,RefNo,ProjectCode,DepartCode,SaleCode,MyDescription,RefDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?)`
			_, err = db.Exec(sqlrec, inv.DocNo, inv.DocDate, inv.ArCode, inv.ExchangeRate, inv.SumBankAmount, 3, linenumber, inv.CreditRefNo, inv.ProjectCode, inv.DepartCode, inv.SaleCode, my_description_recmoney, inv.DocDate)
			if err != nil {
				return err
			}
		}
	}

	sqldepdel := `delete dbo.BCArDepositUse where docno = ?`
	_, err = db.Exec(sqldepdel, inv.DocNo)
	if err != nil {
		return err
	}

	var depLineNumber int

	if (inv.SumOfDeposit1 > 0) {
		for _, d := range inv.Deps {
			d.LineNumber = depLineNumber
			sqldep := `insert into dbo.bcardeposituse(DocNo,DepositNo,DocDate,MyDescription,Balance,Amount,DeposTaxType,NetAmount,LineNumber) values(?,?,?,?,?,?,?,?,?)`
			_, err = db.Exec(sqldep, inv.DocNo, d.DepositNo, inv.DocDate, my_description_recmoney, d.Balance, d.Amount, inv.TaxType, d.NetAmount, d.LineNumber)
			if err != nil {
				return err
			}
			depLineNumber = depLineNumber + 1
		}
	}

	sqlchqdel := `delete dbo.BCChqIn where docno = ?`
	_, err = db.Exec(sqlchqdel, inv.DocNo)
	if err != nil {
		return err
	}

	if (inv.SumChqAmount > 0) {
		for _, c := range inv.Chqs {
			sqldep := `set dateformat dmy     insert into dbo.bcchqin(BankCode,ChqNumber,DocNo,ArCode,SaleCode,ReceiveDate,DueDate,BookNo,Status,SaveFrom,StatusDate,StatusDocNo,DepartCode,BankBranchCode,Amount,Balance,MyDescription,ExchangeRate,RefChqRowOrder) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
			_, err = db.Exec(sqldep, c.BankCode, c.ChqNumber, inv.DocNo, inv.ArCode, inv.SaleCode, c.ReceiveDate, c.DueDate, c.BookNo, c.Status, c.SaveFrom, c.StatusDate, c.StatusDocNo, inv.DepartCode, c.BankBranchCode, c.Amount, c.Balance, my_description_recmoney, inv.ExchangeRate, c.RefChqRowOrder)
			if err != nil {
				return err
			}
		}
	}

	if (inv.SumCreditAmount > 0) {

	}

	return nil
}

func (inv *ArInvoice) SearchArInvoiceByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select DocNo,DocDate,isnull(TaxNo,'') as TaxNo,ArCode,isnull(SaleCode,'') as SaleCode,TaxType,isnull(DepartCode,'') as DepartCode,CreditDay,DeliveryDay,isnull(DeliveryDate,'') as DeliveryDate,isnull(DueDate,'') as DueDate,isnull(PayBillDate,'') as PayBillDate,TaxRate,IsConfirm,isnull(MyDescription,'') as MyDescription,BillType,isnull(BillGroup,'') as BillGroup,isnull(RefDocNo,'') as RefDocNo,isnull(DeliveryAddr,'') as DeliveryAddr,isnull(ContactCode,'') as ContactCode,SumOfItemAmount,isnull(DiscountWord,'') as DiscountWord,isnull(DiscountAmount,'') as DiscountAmount,AfterDiscount,BeforeTaxAmount,TaxAmount,TotalAmount,ZeroTaxAmount,ExceptTaxAmount,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,DepositIncTax,SumOfDeposit1,SumOfDeposit2,SumOfWTax,NetDebtAmount,HomeAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,BillBalance,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,isnull(GLFormat,'') as GLFormat,IsCancel,IsCompleteSave,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(RecurName,'') as RecurName,IsConditionSend,PayBillAmount,isnull(SORefNo,'') as SORefNo,HoldingStatus,PosStatus,CreatorCode,CreateDateTime,isnull(ShiftCode,'') as ShiftCode,isnull(CashierCode,'') as CashierCode,'' as ShiftNo,isnull(MachineNo,'') as MachineNo,isnull(MachineCode,'') as MachineCode,isnull(BillTime,'') as BillTime,isnull(CreditType,'') as CreditType,isnull(CreditDueDate,'') as CreditDueDate,isnull(CreditNo,'') as CreditNo,isnull(CofirmNo,'') as CofirmNo,CreditBaseAmount,CoupongAmount,ChangeAmount,ChargeAmount,GrandTotal from dbo.bcarinvoice where docno = ? `
	fmt.Println("sql=", sql)
	err := db.Get(inv, sql, docno)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	sqlsub := `set dateformat dmy    select MyType, ItemCode, isnull(MyDescription,'') as MyDescription, ItemName, WHCode, ShelfCode, CNQty, Qty, Price, isnull(DiscountWord,'') as DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost, BalanceAmount, UnitCode , isnull(SORefNo,'') as SORefNo, isnull(PORefNo,'') as PORefNo, LineNumber, RefLineNumber, IsCancel, isnull(BarCode,'') as BarCode, IsConditionSend, AverageCost , isnull(LotNumber,'') as LotNumber, isnull(PackingRate1,1) as PackingRate1, isnull(PackingRate2,1) as PackingRate2 from dbo.bcarinvoicesub  where docno = ? order by LineNumber`
	fmt.Println("sqlsub=", sqlsub)
	err = db.Select(&inv.Subs, sqlsub, docno)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (inv *ArInvoice) SearchArInvoiceByKeyword(db *sqlx.DB, keyword string) (invs []*ArInvoice, err error) {
	sql := `set dateformat dmy     select DocNo,DocDate,isnull(TaxNo,'') as TaxNo,ArCode,isnull(SaleCode,'') as SaleCode,TaxType,isnull(DepartCode,'') as DepartCode,CreditDay,DeliveryDay,isnull(DeliveryDate,'') as DeliveryDate,isnull(DueDate,'') as DueDate,isnull(PayBillDate,'') as PayBillDate,TaxRate,IsConfirm,isnull(MyDescription,'') as MyDescription,BillType,isnull(BillGroup,'') as BillGroup,isnull(RefDocNo,'') as RefDocNo,isnull(DeliveryAddr,'') as DeliveryAddr,isnull(ContactCode,'') as ContactCode,SumOfItemAmount,isnull(DiscountWord,'') as DiscountWord,isnull(DiscountAmount,'') as DiscountAmount,AfterDiscount,BeforeTaxAmount,TaxAmount,TotalAmount,ZeroTaxAmount,ExceptTaxAmount,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,DepositIncTax,SumOfDeposit1,SumOfDeposit2,SumOfWTax,NetDebtAmount,HomeAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,BillBalance,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,isnull(GLFormat,'') as GLFormat,IsCancel,IsCompleteSave,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(RecurName,'') as RecurName,IsConditionSend,PayBillAmount,isnull(SORefNo,'') as SORefNo,HoldingStatus,PosStatus,CreatorCode,CreateDateTime,isnull(ShiftCode,'') as ShiftCode,isnull(CashierCode,'') as CashierCode,'' as ShiftNo,isnull(MachineNo,'') as MachineNo,isnull(MachineCode,'') as MachineCode,isnull(BillTime,'') as BillTime,isnull(CreditType,'') as CreditType,isnull(CreditDueDate,'') as CreditDueDate,isnull(CreditNo,'') as CreditNo,isnull(CofirmNo,'') as CofirmNo,CreditBaseAmount,CoupongAmount,ChangeAmount,ChargeAmount,GrandTotal from dbo.bcarinvoice where (docno  like '%'+?+'%' or arcode like '%'+?+'%' ) order by docno`
	fmt.Println("sql=", sql)
	err = db.Select(&invs, sql, keyword, keyword)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	for _, sub := range invs {
		sqlsub := `set dateformat dmy    select MyType, ItemCode, isnull(MyDescription,'') as MyDescription, ItemName, isnull(WHCode,'') as WHCode, isnull(ShelfCode,'') as ShelfCode, CNQty, Qty, Price, isnull(DiscountWord,'') as DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost, BalanceAmount, isnull(UnitCode,'') as UnitCode, isnull(SORefNo,'') as SORefNo, isnull(PORefNo,'') as PORefNo, LineNumber, RefLineNumber, IsCancel, isnull(BarCode,'') as BarCode, IsConditionSend, AverageCost , isnull(LotNumber,'') as LotNumber, isnull(PackingRate1,1) as PackingRate1, isnull(PackingRate2,1) as PackingRate2 from dbo.bcarinvoicesub  where docno = ? order by LineNumber`
		fmt.Println("sqlsub=", sqlsub)
		err = db.Select(&sub.Subs, sqlsub, sub.DocNo)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}
	return invs, nil
}
