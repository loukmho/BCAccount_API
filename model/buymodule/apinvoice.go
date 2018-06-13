package model

import (
	"github.com/jmoiron/sqlx"
	m "github.com/loukmho/bcaccount_api/model"
	"errors"
	"fmt"
	"time"
)

type ApInvoice struct {
	SaveFrom        int                     `json:"save_from" db:"SaveFrom"`
	Source          int                     `json:"source" db:"Source"`
	DocNo           string                  `json:"doc_no" db:"DocNo"`
	ApvInPutTax
	DocDate         string                  `json:"doc_date" db:"DocDate"`
	TaxType         int                     `json:"tax_type" db:"TaxType"`
	ApvVendor
	DepartCode      string                  `json:"depart_code" db:"DepartCode"`
	CreditDay       int                     `json:"credit_day" db:"CreditDay"`
	DueDate         string                  `json:"due_date" db:"DueDate"`
	StatementDate   string                  `json:"statement_date" db:"StatementDate"`
	TaxRate         float64                 `json:"tax_rate" db:"TaxRate"`
	IsConfirm       int                     `json:"is_confirm" db:"IsConfirm"`
	MyDescription   string                  `json:"my_description" db:"MyDescription"`
	BillType        int                     `json:"bill_type" db:"BillType"`
	BillGroup       string                  `json:"bill_group" db:"BillGroup"`
	ContactCode     string                  `json:"contact_code" db:"ContactCode"`
	SumOfItemAmount float64                 `json:"sum_of_item_amount" db:"SumOfItemAmount"`
	DiscountWord    string                  `json:"discount_word" db:"DiscountWord"`
	DiscountAmount  float64                 `json:"discount_amount" db:"DiscountAmount"`
	AfterDiscount   float64                 `json:"after_discount" db:"AfterDiscount"`
	BeforeTaxAmount float64                 `json:"before_tax_amount" db:"BeforeTaxAmount"`
	TaxAmount       float64                 `json:"tax_amount" db:"TaxAmount"`
	TotalAmount     float64                 `json:"total_amount" db:"TotalAmount"`
	ExceptTaxAmount float64                 `json:"except_tax_amount" db:"ExceptTaxAmount"`
	ZeroTaxAmount   float64                 `json:"zero_tax_amount" db:"ZeroTaxAmount"`
	PettyCashAmount float64                 `json:"petty_cash_amount" db:"PettyCashAmount"`
	SumCashAmount   float64                 `json:"sum_cash_amount" db:"SumCashAmount"`
	SumChqAmount    float64                 `json:"sum_chq_amount" db:"SumChqAmount"`
	SumBankAmount   float64                 `json:"sum_bank_amount" db:"SumBankAmount"`
	DepositIncTax   float64                 `json:"deposit_inc_tax" db:"DepositIncTax"`
	SumOfDeposit1   float64                 `json:"sum_of_deposit_1" db:"SumOfDeposit1"`
	SumOfDeposit2   float64                 `json:"sum_of_deposit_2" db:"SumOfDeposit2"`
	SumOfWTax       float64                 `json:"sum_of_w_tax" db:"SumOfWTax"`
	NetDebtAmount   float64                 `json:"net_debt_amount" db:"NetDebtAmount"`
	HomeAmount      float64                 `json:"home_amount" db:"HomeAmount"`
	OtherIncome     float64                 `json:"other_income" db:"OtherIncome"`
	OtherExpense    float64                 `json:"other_expense" db:"OtherExpense"`
	ExcessAmount1   float64                 `json:"excess_amount_1" db:"ExcessAmount1"`
	ExcessAmount2   float64                 `json:"excess_amount_2" db:"ExcessAmount2"`
	BillBalance     float64                 `json:"bill_balance" db:"BillBalance"`
	CurrencyCode    string                  `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate    float64                 `json:"exchange_rate" db:"ExchangeRate"`
	GLFormat        string                  `json:"gl_format" db:"GLFormat"`
	GLStartPosting  int                     `json:"gl_start_posting" db:"GLStartPosting"`
	GRBillStatus    int                     `json:"gr_bill_status" db:"GRBillStatus"`
	GRIRBillStatus  int                     `json:"grir_bill_status" db:"GRIRBillStatus"`
	IsCancel        int                     `json:"is_cancel" db:"IsCancel"`
	IsCreditNote    int                     `json:"is_credit_note" db:"IsCreditNote"`
	IsDebitNote     int                     `json:"is_debit_note" db:"IsDebitNote"`
	IsCompleteSave  int                     `json:"is_complete_save" db:"IsCompleteSave"`
	AllocateCode    string                  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string                  `json:"project_code" db:"ProjectCode"`
	RecurName       string                  `json:"recur_name" db:"RecurName"`
	ExchangeProfit  float64                 `json:"exchange_profit" db:"ExchangeProfit"`
	ConfirmCode     string                  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string                  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string                  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string                  `json:"cancel_date_time" db:"CancelDateTime"`
	PayBillAmount   float64                 `json:"pay_bill_amount" db:"PayBillAmount"`
	RefDocNo        string                  `json:"ref_doc_no" db:"RefDocNo"`
	CreatorCode     string                  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string                  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string                  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string                  `json:"last_edit_date_t" db:"LastEditDateT"`
	UserCode        string                  `json:"user_code" db:"UserCode"`
	ListApvPayMoney
	Subs            []*ApvItem              `json:"subs"`
	Deps            []*ListApvApDepositUsed `json:"deps"`
	Cdcs            []*ListApvCreditCard    `json:"cdcs"`
	Chqs            []*ListApvChqOut        `json:"chqs"`
}

type ApvInPutTax struct {
	TaxNo    string `json:"tax_no" db:"TaxNo"`
	BookCode string `json:"book_code" db:"BookCode"`
}

type ApvVendor struct {
	ApCode string `json:"ap_code" db:"ApCode"`
	ApName string `json:"ap_name" db:"ApName"`
}

type ApvItem struct {
	MyType         int     `json:"my_type" db:"MyType"`
	StockType      int     `json:"stock_type" db:"StockType"`
	ItemCode       string  `json:"item_code" db:"ItemCode"`
	MyDescription  string  `json:"my_description" db:"MyDescription"`
	ItemName       string  `json:"item_name" db:"ItemName"`
	WHCode         string  `json:"wh_code" db:"WHCode"`
	ShelfCode      string  `json:"shelf_code" db:"ShelfCode"`
	CNQty          float64 `json:"cn_qty" db:"CNQty"`
	GRRemainQty    float64 `json:"gr_remain_qty" db:"GRRemainQty"`
	Qty            float64 `json:"qty" db:"Qty"`
	Price          float64 `json:"price" db:"Price"`
	DiscountWord   string  `json:"discount_word" db:"DiscountWord"`
	DiscountAmount float64 `json:"discount_amount" db:"DiscountAmount"`
	Amount         float64 `json:"amount" db:"Amount"`
	NetAmount      float64 `json:"net_amount" db:"NetAmount"`
	HomeAmount     float64 `json:"home_amount" db:"HomeAmount"`
	BalanceAmount  float64 `json:"balance_amount" db:"BalanceAmount"`
	SumOfExpCost   float64 `json:"sum_of_exp_cost" db:"SumOfExpCost"`
	UnitCode       string  `json:"unit_code" db:"UnitCode"`
	PORefNo        string  `json:"po_ref_no"  db:"PORefNo"`
	IRRefNo        string  `json:"ir_ref_no" db:"IRRefNo"`
	IsCancel       int     `json:"is_cancel" db:"IsCancel"`
	LineNumber     int     `json:"line_number" db:"LineNumber"`
	BarCode        string  `json:"bar_code" db:"BarCode"`
	PORefLinenum   int     `json:"po_ref_linenum" db:"PORefLinenum"`
	AVERAGECOST    float64 `json:"averagecost" db:"AVERAGECOST"`
	LotNumber      string  `json:"lot_number" db:"LotNumber"`
	SumOfCost      float64 `json:"sum_of_cost" db:"SumOfCost"`
	PackingRate1   float64 `json:"packing_rate_1" db:"PackingRate1"`
	PackingRate2   float64 `json:"packing_rate_2" db:"PackingRate2"`
}

type ListApvPayMoney struct {
	CreditType     string `json:"credit_type" db:"CreditType"`
	ConfirmNo      string `json:"confirm_no" db:"ConfirmNo"`
	CreditRefNo    string `json:"credit_ref_no" db:"CreditRefNo"`
	CreditDueDate  string `json:"credit_due_date" db:"CreditDueDate"`
	BankCode       string `json:"bank_code" db:"BookCode"`
	BankBranchCode string `json:"bank_branch_code" db:"BankBranchCode"`
	BankRefNo      string `json:"bank_ref_no" db:"BankRefNo"`
	TransBankDate  string `json:"trans_bank_date" db:"TransBankDate"`
	RefDate        string `json:"ref_date" db:"RefDate"`
}

type ListApvApDepositUsed struct {
	DepositNo  string  `json:"deposit_no" db:"DepositNo"`
	Balance    float64 `json:"balance" db:"Balance"`
	Amount     float64 `json:"amount" db:"Amount"`
	NetAmount  float64 `json:"net_amount" db:"NetAmount"`
	LineNumber int     `json:"line_number" db:"LineNumber"`
}

type ListApvChqOut struct {
	ChqNumber      string  `json:"chq_number" db:"ChqNumber"`
	BankCode       string  `json:"bank_code" db:"BankCode"`
	BankBranchCode string  `json:"bank_branch_code" db:"BankBranchCode"`
	BookNo         string  `json:"book_no" db:"BookNo"`
	ReceiveDate    string  `json:"receive_date" db:"ReceiveDate"`
	DueDate        string  `json:"due_date" db:"DueDate"`
	Status         int     `json:"status" db:"Status"`
	Amount         float64 `json:"amount" db:"Amount"`
	Balance        float64 `json:"balance" db:"Balance"`
	RefChqRowOrder int     `json:"ref_chq_row_order" db:"RefChqRowOrder"`
	StatusDate     string  `json:"status_date" db:"StatusDate"`
	StatusDocNo    string  `json:"status_doc_no" db:"StatusDocNo"`
}

type ListApvCreditCard struct {
	BankCode       string  `json:"bank_code" db:"BankCode"`
	CreditCardNo   string  `json:"credit_card_no" db:"CreditCardNo"`
	ReceiveDate    string  `json:"receive_date" db:"ReceiveDate"`
	DueDate        string  `json:"due_date" db:"DueDate"`
	BookNo         string  `json:"book_no" db:"BookNo"`
	Status         int     `json:"status" db:"Status"`
	StatusDate     string  `json:"status_date" db:"StatusDate"`
	StatusDocNo    string  `json:"status_doc_no" db:"StatusDocNo"`
	BankBranchCode string  `json:"bank_branch_code" db:"BankBranchCode"`
	Amount         float64 `json:"amount" db:"Amount"`
	MyDescription  string  `json:"my_description" db:"MyDescription"`
	CreditType     string  `json:"credit_type" db:"CreditType"`
	ConfirmNo      string  `json:"confirm_no" db:"ConfirmNo"`
	ChargeAmount   float64 `json:"charge_amount" db:"ChargeAmount"`
}

func (apv *ApInvoice) InsertAndEditApInvoice(db *sqlx.DB) error {
	var check_exist int
	var count_item int
	var sum_pay_amount float64

	now := time.Now()

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")

	fmt.Println("TaxRate = ", def.TaxRateDefault)

	if apv.TaxRate == 0 {
		apv.TaxRate = def.TaxRateDefault
	}

	sum_pay_amount = (apv.SumCashAmount + apv.SumChqAmount + apv.SumBankAmount + apv.OtherExpense) - apv.OtherIncome

	var AfterDepositAmount float64

	AfterDepositAmount = apv.AfterDiscount - apv.SumOfDeposit1

	if (apv.BeforeTaxAmount == 0 && apv.TaxAmount == 0) {
		apv.BeforeTaxAmount, apv.TaxAmount, apv.TotalAmount = m.CalcTaxItem(apv.TaxType, apv.TaxRate, AfterDepositAmount)
	}
	for _, sub_item := range apv.Subs {
		if (sub_item.Qty != 0) {
			count_item = count_item + 1
		}
	}

	sqlexist := `select count(docno) as check_exist from dbo.bcapinvoice where docno = ?`
	err := db.Get(&check_exist, sqlexist, apv.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	switch {
	case apv.DocNo == "":
		return errors.New("Docno is null")
	case apv.ApCode == "":
		return errors.New("Apcode is null")
	case apv.DocDate == "":
		return errors.New("Docdate is null")
	case apv.IsCancel == 1:
		return errors.New("Docno is cancel")
	case apv.IsConfirm == 1:
		return errors.New("Docno is confirm")
	case count_item == 0:
		return errors.New("Docno is not have item")
	case (apv.SumCashAmount == 0 && apv.SumChqAmount == 0 && apv.SumBankAmount == 0 && apv.SumOfDeposit1 == 0 && apv.BillType == 0):
		return errors.New("Docno not set money to another type payment")
	case apv.SumOfItemAmount == 0:
		return errors.New("Sumofitemamount = 0")
	case sum_pay_amount > apv.TotalAmount && apv.BillType == 0:
		return errors.New("Rec money is over totalamount")
	case apv.SumChqAmount != 0 && len(apv.Chqs) == 0:
		return errors.New("Docno not have chq use data")
	case apv.SumOfDeposit1 != 0 && len(apv.Deps) == 0:
		return errors.New("Docno not have deposit use data")
	case apv.SumBankAmount != 0 && apv.BankRefNo == "":
		return errors.New("Bank transfer not have accountcode data not complete")
	}

	var sumDepNetAmount float64

	if apv.TaxType == 1 {
		for _, dep := range apv.Deps {
			sumDepNetAmount = sumDepNetAmount + dep.Amount
		}

		if apv.SumOfDeposit1 != sumDepNetAmount {
			return errors.New("Docno not have deposit not equa receive data")
		}
	}

	//if (apv.SumOfDeposit1 != 0) {
	apv.DepositIncTax = 1
	//} else {
	//	apv.DepositIncTax = 0
	//}

	if apv.ExchangeRate == 0 {
		apv.ExchangeRate = def.ExchangeRateDefault
	}

	if apv.SaveFrom == 0 {
		apv.SaveFrom = def.ApInvoiceSaveFrom
	}

	apv.NetDebtAmount = apv.TotalAmount
	apv.HomeAmount = apv.TotalAmount

	apv.BillBalance = apv.TotalAmount + apv.OtherIncome - (apv.SumCashAmount + apv.SumChqAmount + apv.SumBankAmount + apv.OtherExpense)

	if (apv.BillType == 0 && ((apv.TotalAmount+apv.OtherIncome-(apv.SumCashAmount+apv.SumChqAmount+apv.SumBankAmount+apv.OtherExpense))) != 0) {
		return errors.New("Docno have money remain")
	}

	if (apv.BookCode == "") {
		apv.BookCode = def.ArInvoiceBookCode
	}
	if (apv.Source == 0) {
		apv.Source = def.ArInvoiceSource
	}
	if (apv.BillType == 0 && apv.GLFormat == "") {
		apv.GLFormat = def.ApInvoiceCashGLFormat
	} else {
		apv.GLFormat = def.ApInvoiceCreditGLFormat
	}
	if (apv.TaxNo == "") {
		apv.TaxNo = apv.DocNo
	}
	if (apv.CreditDay == 0 || apv.DueDate == "") {
		apv.DueDate = apv.DocDate
		apv.StatementDate = apv.DocDate
	} else {
		apv.DueDate = now.AddDate(0, 0, apv.CreditDay).Format("2006-01-02")
		apv.StatementDate = now.AddDate(0, 0, apv.CreditDay).Format("2006-01-02")
	}

	if (apv.SumBankAmount != 0 && apv.TransBankDate == "") {
		apv.TransBankDate = apv.DocDate
	}

	if (apv.SumOfDeposit1 != 0 && apv.TaxRate != 0) {
		if apv.TaxType == 1 {
			apv.SumOfDeposit2 = apv.SumOfDeposit1 - m.ToFixed(apv.SumOfDeposit1-((apv.SumOfDeposit1*100)/(100+float64(apv.TaxRate))), 2)
		} else {
			apv.SumOfDeposit2 = apv.SumOfDeposit1 //+ m.ToFixed(inv.SumOfDeposit1-((inv.SumOfDeposit1*100)/(100+float64(inv.TaxRate))), 2)
		}
	}

	fmt.Println("UserCode = ", apv.UserCode)

	if (apv.IsCompleteSave == 0) {
		apv.IsCompleteSave = 1
	}

	fmt.Println("check_exist = ", check_exist)
	if (check_exist == 0) {
		//Insert/////////////////////////////////////////////////////////////////////////////////////////////////////////
		apv.CreatorCode = apv.UserCode

		sql := `insert into dbo.bcapinvoice(DocNo,TaxNo,DocDate,TaxType,ApCode,DepartCode,CreditDay,DueDate,StatementDate,TaxRate,IsConfirm,MyDescription,BillType,BillGroup,ContactCode,SumOfItemAmount,DiscountWord,DiscountAmount,AfterDiscount,BeforeTaxAmount,TaxAmount,TotalAmount,ExceptTaxAmount,ZeroTaxAmount,PettyCashAmount,SumCashAmount,SumChqAmount,SumBankAmount,DepositIncTax,SumOfDeposit1,SumOfDeposit2,SumOfWTax,NetDebtAmount,HomeAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,BillBalance,CurrencyCode,ExchangeRate,GLFormat,GLStartPosting,GRBillStatus,GRIRBillStatus,IsCancel,IsCreditNote,IsDebitNote, IsCompleteSave, AllocateCode,ProjectCode,RecurName,ExchangeProfit,PayBillAmount,RefDocNo,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,getdate())`
		_, err = db.Exec(sql, apv.DocNo, apv.TaxNo, apv.DocDate, apv.TaxType, apv.ApCode, apv.DepartCode, apv.CreditDay, apv.DueDate, apv.StatementDate, apv.TaxRate, apv.IsConfirm, apv.MyDescription, apv.BillType, apv.BillGroup, apv.ContactCode, apv.SumOfItemAmount, apv.DiscountWord, apv.DiscountAmount, apv.AfterDiscount, apv.BeforeTaxAmount, apv.TaxAmount, apv.TotalAmount, apv.ExceptTaxAmount, apv.ZeroTaxAmount, apv.PettyCashAmount, apv.SumCashAmount, apv.SumChqAmount, apv.SumBankAmount, apv.DepositIncTax, apv.SumOfDeposit1, apv.SumOfDeposit2, apv.SumOfWTax, apv.NetDebtAmount, apv.HomeAmount, apv.OtherIncome, apv.OtherExpense, apv.ExcessAmount1, apv.ExcessAmount2, apv.BillBalance, apv.CurrencyCode, apv.ExchangeRate, apv.GLFormat, apv.GLStartPosting, apv.GRBillStatus, apv.GRIRBillStatus, apv.IsCancel, apv.IsCreditNote, apv.IsDebitNote, apv.IsCompleteSave, apv.AllocateCode, apv.ProjectCode, apv.RecurName, apv.ExchangeProfit, apv.PayBillAmount, apv.RefDocNo, apv.CreatorCode)
		if err != nil {
			fmt.Println("sql insert= ", err.Error())
			return err
		}

	} else {
		//Update/////////////////////////////////////////////////////////////////////////////////////////////////////////
		apv.LastEditorCode = apv.UserCode
		sql := `update dbo.bcapinvoice set TaxNo=?,DocDate=?,TaxType=?,ApCode=?,DepartCode=?,CreditDay=?,DueDate=?,StatementDate=?,TaxRate=?,IsConfirm=?,MyDescription=?,BillType=?,BillGroup=?,ContactCode=?,SumOfItemAmount=?,DiscountWord=?,DiscountAmount=?,AfterDiscount=?,BeforeTaxAmount=?,TaxAmount=?,TotalAmount=?,ExceptTaxAmount=?,ZeroTaxAmount=?,PettyCashAmount=?,SumCashAmount=?,SumChqAmount=?,SumBankAmount=?,DepositIncTax=?,SumOfDeposit1=?,SumOfDeposit2=?,SumOfWTax=?,NetDebtAmount=?,HomeAmount=?,OtherIncome=?,OtherExpense=?,ExcessAmount1=?,ExcessAmount2=?,BillBalance=?,CurrencyCode=?,ExchangeRate=?,GLFormat=?,GLStartPosting=?,GRBillStatus=?,GRIRBillStatus=?,IsCancel=?,IsCreditNote=?,IsDebitNote=?,AllocateCode=?,ProjectCode=?,RecurName=?,ExchangeProfit=?,PayBillAmount=?,RefDocNo=?,LastEditorCode=?,LastEditDateT=getdate() where docno = ?`
		_, err = db.Exec(sql, apv.TaxNo, apv.DocDate, apv.TaxType, apv.ApCode, apv.DepartCode, apv.CreditDay, apv.DueDate, apv.StatementDate, apv.TaxRate, apv.IsConfirm, apv.MyDescription, apv.BillType, apv.BillGroup, apv.ContactCode, apv.SumOfItemAmount, apv.DiscountWord, apv.DiscountAmount, apv.AfterDiscount, apv.BeforeTaxAmount, apv.TaxAmount, apv.TotalAmount, apv.ExceptTaxAmount, apv.ZeroTaxAmount, apv.PettyCashAmount, apv.SumCashAmount, apv.SumChqAmount, apv.SumBankAmount, apv.DepositIncTax, apv.SumOfDeposit1, apv.SumOfDeposit2, apv.SumOfWTax, apv.NetDebtAmount, apv.HomeAmount, apv.OtherIncome, apv.OtherExpense, apv.ExcessAmount1, apv.ExcessAmount2, apv.BillBalance, apv.CurrencyCode, apv.ExchangeRate, apv.GLFormat, apv.GLStartPosting, apv.GRBillStatus, apv.GRIRBillStatus, apv.IsCancel, apv.IsCreditNote, apv.IsDebitNote, apv.AllocateCode, apv.ProjectCode, apv.RecurName, apv.ExchangeProfit, apv.PayBillAmount, apv.RefDocNo, apv.LastEditorCode, apv.DocNo)
		if err != nil {
			fmt.Println("sql update= ", err.Error())
			return err
		}
	}

	sql_del_sub := `delete dbo.bcapinvoicesub where docno = ?`
	_, err = db.Exec(sql_del_sub, apv.DocNo)
	if err != nil {
		return err
	}

	var vLineNumber int

	for _, item := range apv.Subs {
		fmt.Println("ItemSub")
		item.MyType = def.ArInvoiceMyType
		item.CNQty = item.Qty
		item.GRRemainQty = item.Qty
		item.BalanceAmount = item.Amount
		item.LineNumber = vLineNumber

		item.IsCancel = 0
		if (item.PackingRate1 == 0) {
			item.PackingRate1 = 1
		}
		if (item.PackingRate2 == 0) {
			item.PackingRate2 = 1
		}

		switch {
		case apv.TaxType == 0:
			item.HomeAmount = item.Amount
			item.NetAmount = item.Amount
		case apv.TaxType == 1:
			taxamount := m.ToFixed(item.Amount-((item.Amount*100)/(100+float64(apv.TaxRate))), 2)
			beforetaxamount := m.ToFixed(item.Amount-taxamount, 2)
			item.HomeAmount = beforetaxamount
			item.NetAmount = beforetaxamount
		case apv.TaxType == 2:
			item.HomeAmount = item.Amount
			item.NetAmount = item.Amount
		}

		sqlsub := ` insert into dbo.BCApInvoiceSub(MyType,DocNo,TaxNo,TaxType,ItemCode,DocDate,ApCode,DepartCode,MyDescription,ItemName,WHCode,ShelfCode,CNQty,GRRemainQty,Qty,Price,DiscountWord,DiscountAmount,Amount,NetAmount,HomeAmount,BalanceAmount,SumOfExpCost,UnitCode,PORefNo,IRRefNo,IsCancel,LineNumber,BarCode,PORefLinenum,AVERAGECOST,LotNumber,SumOfCost,TaxRate,PackingRate1,PackingRate2) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		_, err = db.Exec(sqlsub, item.MyType, apv.DocNo, apv.TaxNo, apv.TaxType, item.ItemCode, apv.DocDate, apv.ApCode, apv.DepartCode, item.MyDescription, item.ItemName, item.WHCode, item.ShelfCode, item.CNQty, item.GRRemainQty, item.Qty, item.Price, item.DiscountWord, item.DiscountAmount, item.Amount, item.NetAmount, item.HomeAmount, item.BalanceAmount, item.SumOfExpCost, item.UnitCode, item.PORefNo, item.IRRefNo, item.IsCancel, item.LineNumber, item.BarCode, item.PORefLinenum, item.AVERAGECOST, item.LotNumber, item.SumOfCost, apv.TaxRate, item.PackingRate1, item.PackingRate2)
		fmt.Println("sqlsub = ", sqlsub)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		if (item.StockType != 1) {
			sqlprocess := ` insert into dbo.ProcessStock (ItemCode,ProcessFlag,FlowStatus) values(?, 1, 0)`
			_, err = db.Exec(sqlprocess, item.ItemCode)
			fmt.Println("sqlprocess = ", sqlsub)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		vLineNumber = vLineNumber + 1
	}

	sqldel := `delete dbo.BCInputTax where docno = ?`
	_, err = db.Exec(sqldel, apv.DocNo)
	if err != nil {
		return err
	}

	sqltax := `insert into dbo.BCInputTax(SaveFrom,DocNo,BookCode,Source,DocDate,TaxDate,TaxNo,ApCode,ShortTaxDesc,TaxRate,Process,BeforeTaxAmount,TaxAmount,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?,?,'ซื้อสินค้า',?,1,?,?,?,getdate())`
	_, err = db.Exec(sqltax, apv.SaveFrom, apv.DocNo, apv.BookCode, apv.Source, apv.DocDate, apv.DocDate, apv.TaxNo, apv.ApCode, apv.TaxRate, apv.BeforeTaxAmount, apv.TaxAmount, apv.UserCode)
	fmt.Println("sqltax = ", sqltax)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	sqlrecdel := `delete dbo.BCPayMoney where docno = ?`
	_, err = db.Exec(sqlrecdel, apv.DocNo)
	if err != nil {
		return err
	}

	var my_description_recmoney string

	my_description_recmoney = "ขายเงินสด"

	fmt.Println("RecMoney")
	var linenumber int

	if (apv.BillType == 0) {
		if (apv.SumCashAmount != 0) { //subs.PaymentType == 0:
			fmt.Println("SumCashAmount")
			sqlrec := `insert	into dbo.BCPayMoney(DocNo,DocDate,ApCode,ExchangeRate,PayAmount,PaymentType,SaveFrom,LineNumber,ProjectCode,DepartCode,MyDescription) values(?,?,?,?,?,?,?,?,?,?,?)`
			_, err = db.Exec(sqlrec, apv.DocNo, apv.DocDate, apv.ApCode, apv.ExchangeRate, apv.SumCashAmount, 0, apv.SaveFrom, linenumber, apv.ProjectCode, apv.DepartCode, my_description_recmoney)
			if err != nil {
				return err
			}
		}
		//case dp.SumCreditAmount != 0: //subs.PaymentType == 1:
		if (apv.SumChqAmount != 0) {
			fmt.Println("SumChqAmount")
			if (apv.SumCashAmount != 0) {
				linenumber = 1
			} else {
				linenumber = 0
			}
			sqlrec := `insert	into dbo.BCPayMoney(DocNo,DocDate,ApCode,ExchangeRate,PayAmount,PaymentType,SaveFrom,LineNumber,RefNo,BankCode,ProjectCode,DepartCode,BankBranchCode,MyDescription,RefDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
			_, err = db.Exec(sqlrec, apv.DocNo, apv.DocDate, apv.ApCode, apv.ExchangeRate, apv.SumChqAmount, 2, apv.SaveFrom, linenumber, apv.CreditRefNo, apv.BankCode, apv.ProjectCode, apv.DepartCode, apv.BankBranchCode, my_description_recmoney, apv.RefDate)
			if err != nil {
				return err
			}
		}

		//case dp.SumChqAmount != 0: //subs.PaymentType == 2:
		if (apv.SumBankAmount != 0) {
			fmt.Println("SumChqAmount")
			if (apv.SumCashAmount != 0 && apv.SumChqAmount != 0) {
				linenumber = 2
			} else if ((apv.SumCashAmount != 0 && apv.SumChqAmount == 0)) {
				linenumber = 1
			} else if ((apv.SumCashAmount == 0 && apv.SumChqAmount != 0)) {
				linenumber = 1
			} else if ((apv.SumCashAmount == 0 && apv.SumChqAmount == 0)) {
				linenumber = 0
			}

			sqlrec := `insert	into dbo.BCPayMoney(DocNo,DocDate,ApCode,ExchangeRate,PayAmount,PaymentType,SaveFrom,LineNumber,RefNo,ProjectCode,DepartCode,MyDescription,RefDate,TransBankDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
			_, err = db.Exec(sqlrec, apv.DocNo, apv.DocDate, apv.ApCode, apv.ExchangeRate, apv.SumBankAmount, 3, apv.SaveFrom, linenumber, apv.BankRefNo, apv.ProjectCode, apv.DepartCode, my_description_recmoney, apv.DocDate, apv.TransBankDate)
			if err != nil {
				return err
			}
		}

		//case dp.SumBankAmount != 0: //subs.PaymentType == 3:
		//if (inv.SumBankAmount != 0) {
		//	fmt.Println("SumBankAmount")
		//	if (inv.SumCashAmount != 0 && inv.SumCreditAmount != 0 && inv.SumChqAmount != 0) {
		//		linenumber = 3
		//	} else if (inv.SumCashAmount != 0 && inv.SumCreditAmount == 0 && inv.SumChqAmount != 0) {
		//		linenumber = 2
		//	} else if (inv.SumCashAmount == 0 && inv.SumCreditAmount != 0 && inv.SumChqAmount != 0) {
		//		linenumber = 2
		//	} else if (inv.SumCashAmount != 0 && inv.SumCreditAmount != 0 && inv.SumChqAmount == 0) {
		//		linenumber = 2
		//	} else if (inv.SumCashAmount != 0 && inv.SumCreditAmount != 0 && inv.SumChqAmount == 0) {
		//		linenumber = 2
		//	} else if (inv.SumCashAmount != 0 && inv.SumCreditAmount == 0 && inv.SumChqAmount == 0) {
		//		linenumber = 1
		//	} else if (inv.SumCashAmount == 0 && inv.SumCreditAmount != 0 && inv.SumChqAmount == 0) {
		//		linenumber = 1
		//	} else if (inv.SumCashAmount == 0 && inv.SumCreditAmount == 0 && inv.SumChqAmount != 0) {
		//		linenumber = 1
		//	} else if (inv.SumCashAmount == 0 && inv.SumCreditAmount == 0 && inv.SumChqAmount == 0) {
		//		linenumber = 0
		//	}
		//
		//	sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,PaymentType,SaveFrom,LineNumber,RefNo,ProjectCode,DepartCode,SaleCode,MyDescription,RefDate,TransBankDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		//	_, err = db.Exec(sqlrec, inv.DocNo, inv.DocDate, inv.ArCode, inv.ExchangeRate, inv.SumBankAmount, 3, inv.SaveFrom, linenumber, inv.BankRefNo, inv.ProjectCode, inv.DepartCode, inv.SaleCode, my_description_recmoney, inv.DocDate, inv.TransBankDate)
		//	if err != nil {
		//		return err
		//	}
		//}
	}

	var depLineNumber int

	if (apv.SumOfDeposit1 != 0) {
		sqldepdel := `delete dbo.BCApDepositUse where docno = ?`
		_, err = db.Exec(sqldepdel, apv.DocNo)
		if err != nil {
			return err
		}

		for _, d := range apv.Deps {
			var depNetAmount float64
			d.LineNumber = depLineNumber

			if apv.TaxType == 0 {
				depNetAmount = d.Amount - (m.ToFixed(d.Amount-((d.Amount*100)/(100+float64(apv.TaxRate))), 2))

			} else {
				depNetAmount = d.Amount - (m.ToFixed(d.Amount-((d.Amount*100)/(100+float64(apv.TaxRate))), 2))
			}

			sqldep := `insert into dbo.bcapdeposituse(DocNo,DepositNo,DocDate,MyDescription,Balance,Amount,DeposTaxType,NetAmount,LineNumber) values(?,?,?,?,?,?,?,?,?)`
			_, err = db.Exec(sqldep, apv.DocNo, d.DepositNo, apv.DocDate, my_description_recmoney, d.Balance, d.Amount, apv.TaxType, depNetAmount, d.LineNumber)
			if err != nil {
				return err
			}
			depLineNumber = depLineNumber + 1
		}
	}

	if (apv.SumChqAmount != 0) {
		sqlchqdel := `delete dbo.BCChqOut where docno = ?`
		_, err = db.Exec(sqlchqdel, apv.DocNo)
		if err != nil {
			return err
		}

		for _, c := range apv.Chqs {
			if ((c.ReceiveDate == "") || (c.DueDate == "")) {
				c.ReceiveDate = apv.DocDate
				c.DueDate = apv.DocDate
			}

			sqldep := `insert into dbo.bcchqout(BankCode,ChqNumber,DocNo,ApCode,DueDate,BookNo,Status,SaveFrom,StatusDate,StatusDocNo,DepartCode,BankBranchCode,Amount,Balance,MyDescription,ExchangeRate,RefChqRowOrder) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
			_, err = db.Exec(sqldep, c.BankCode, c.ChqNumber, apv.DocNo, apv.ApCode, c.DueDate, c.BookNo, c.Status, apv.SaveFrom, c.StatusDate, c.StatusDocNo, apv.DepartCode, c.BankBranchCode, c.Amount, c.Balance, my_description_recmoney, apv.ExchangeRate, c.RefChqRowOrder)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

//
//func (ai *ApInvoice) SearchApInvoiceByDocNo(db *sqlx.DB) error {
//	sql := `select a.DocNo,isnull(a.TaxNo,'') as TaxNo,DocDate,a.TaxType,isnull(ApCode,'') as ApCode,isnull(DepartCode,'') as DepartCode,a.CreditDay,isnull(DueDate,'') as DueDate,isnull(StatementDate,'') as StatementDate,TaxRate,IsConfirm,isnull(MyDescription,'') as MyDescription,BillType,isnull(BillGroup,'') as BillGroup,isnull(ContactCode,'') as ContactCode,SumOfItemAmount,isnull(DiscountWord,'') as DiscountWord,DiscountAmount,AfterDiscount,BeforeTaxAmount,TaxAmount,TotalAmount,ExceptTaxAmount,ZeroTaxAmount,PettyCashAmount,SumCashAmount,SumChqAmount,SumBankAmount,DepositIncTax,SumOfDeposit1,SumOfDeposit2,SumOfWTax,NetDebtAmount,HomeAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,BillBalance,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,isnull(GLFormat,'') as GLFormat,GLStartPosting,GRBillStatus,GRIRBillStatus,IsCancel,IsCreditNote,IsDebitNote,IsCompleteSave,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(RecurName,'') as RecurName,ExchangeProfit,isnull(a.ConfirmCode,'') as ConfirmCode,isnull(a.ConfirmDateTime,'') as ConfirmDateTime,isnull(a.CancelCode,'') as CancelCode,isnull(a.CancelDateTime,'') as CancelDateTime,PayBillAmount,isnull(RefDocNo,'') as RefDocNo,isnull(a.CreatorCode,'') as CreatorCode,isnull(a.CreateDateTime,'') as CreateDateTime,isnull(a.LastEditorCode,'') as LastEditorCode,isnull(a.LastEditDateT,'') as LastEditDateT from dbo.bcapinvoice a left join dbo.bcap b with (nolock) on a.apcode = b.code`
//	//sql := `DocNo, TaxNo, DocDate, TaxType, ApCode, DepartCode, CreditDay, DueDate, StatementDate, TaxRate, IsConfirm, MyDescription, BillType, BillGroup, ContactCode, SumOfItemAmount,DiscountWord, DiscountAmount, AfterDiscount, BeforeTaxAmount, TaxAmount, TotalAmount, DiscountCase, ExceptTaxAmount, ZeroTaxAmount, PettyCashAmount, SumCashAmount, SumChqAmount, SumBankAmount,DepositIncTax, SumOfDeposit1, SumOfDeposit2, SumOfWTax, NetDebtAmount, HomeAmount, OtherIncome, OtherExpense, ExcessAmount1, ExcessAmount2, BillBalance, CurrencyCode, ExchangeRate, GLFormat,GLStartPosting, GRBillStatus, GRIRBillStatus, IsCancel, IsCreditNote, IsDebitNote,  IsCompleteSave,  AllocateCode, ProjectCode, RecurName, ExchangeProfit, ConfirmCode,ConfirmDateTime, CancelCode, CancelDateTime, PayBillAmount, RefDocNo,CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT`
//	//sqlsub := ` MyType, DocNo, TaxNo, TaxType, ItemCode, DocDate, ApCode, DepartCode, MyDescription, ItemName, WHCode, ShelfCode, CNQty, GRRemainQty, Qty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, BalanceAmount, SumOfExpCost, UnitCode, PORefNo, IRRefNo, IsCancel, LineNumber, BarCode, PORefLinenum,  AVERAGECOST,  LotNumber, SumOfCost, TaxRate, PackingRate1, PackingRate2`
//	////err := db.Query(ai, sql)
//	//fmt.Println("sql =", sql)
//	//if err != nil {
//	//	return err
//	//}
//	return nil
//}

func (inv *ApInvoice) SearchApInvoiceByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select DocNo,isnull(a.TaxNo,'') as TaxNo,DocDate,a.TaxType,isnull(ApCode,'') as ApCode,isnull(DepartCode,'') as DepartCode,a.CreditDay,isnull(DueDate,'') as DueDate,isnull(StatementDate,'') as StatementDate,TaxRate,IsConfirm,isnull(MyDescription,'') as MyDescription,BillType,isnull(BillGroup,'') as BillGroup,isnull(ContactCode,'') as ContactCode,SumOfItemAmount,isnull(DiscountWord,'') as DiscountWord,DiscountAmount,AfterDiscount,BeforeTaxAmount,TaxAmount,TotalAmount,ExceptTaxAmount,ZeroTaxAmount,PettyCashAmount,SumCashAmount,SumChqAmount,SumBankAmount,DepositIncTax,SumOfDeposit1,SumOfDeposit2,SumOfWTax,NetDebtAmount,HomeAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,BillBalance,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,isnull(GLFormat,'') as GLFormat,GLStartPosting,GRBillStatus,GRIRBillStatus,IsCancel,IsCreditNote,IsDebitNote,IsCompleteSave,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(RecurName,'') as RecurName,ExchangeProfit,isnull(a.ConfirmCode,'') as ConfirmCode,isnull(a.ConfirmDateTime,'') as ConfirmDateTime,isnull(a.CancelCode,'') as CancelCode,isnull(a.CancelDateTime,'') as CancelDateTime,PayBillAmount,isnull(RefDocNo,'') as RefDocNo,isnull(a.CreatorCode,'') as CreatorCode,isnull(a.CreateDateTime,'') as CreateDateTime,isnull(a.LastEditorCode,'') as LastEditorCode,isnull(a.LastEditDateT,'') as LastEditDateT from dbo.bcapinvoice a  with (nolock) left join dbo.bcap b with (nolock) on a.apcode = b.code where docno = ? `
	fmt.Println("sql=", sql)
	err := db.Get(inv, sql, docno)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	sqlsub := `set dateformat dmy    select MyType,StockType,isnull(ItemCode,'') as ItemCode,isnull(MyDescription,'') as ItemName,isnull(ItemName,'') as ItemName,isnull(WHCode,'') as WHCode,isnull(ShelfCode,'') as ShelfCode,CNQty,GRRemainQty,Qty,Price,isnull(DiscountWord,'') as DiscountWord,DiscountAmount,Amount,NetAmount,HomeAmount,BalanceAmount,SumOfExpCost,isnull(UnitCode,'') as UnitCode,isnull(PORefNo,'') as PORefNo,isnull(IRRefNo,'') as IRRefNo,IsCancel,LineNumber,isnull(BarCode,'') as BarCode,PORefLinenum,AVERAGECOST,isnull(LotNumber,'') as LotNumber,SumOfCost,isnull(PackingRate1,1) as PackingRate1,isnull(PackingRate2,1) as PackingRate2 from dbo.bcapinvoicesub  with (nolock)  where docno = ? order by LineNumber`
	fmt.Println("sqlsub=", sqlsub)
	err = db.Select(&inv.Subs, sqlsub, docno)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (inv *ApInvoice) SearchApInvoiceByKeyword(db *sqlx.DB, keyword string) (apvs []*ApInvoice, err error) {
	sql := `set dateformat dmy     select DocNo,isnull(a.TaxNo,'') as TaxNo,DocDate,a.TaxType,isnull(ApCode,'') as ApCode,isnull(DepartCode,'') as DepartCode,a.CreditDay,isnull(DueDate,'') as DueDate,isnull(StatementDate,'') as StatementDate,TaxRate,IsConfirm,isnull(MyDescription,'') as MyDescription,BillType,isnull(BillGroup,'') as BillGroup,isnull(ContactCode,'') as ContactCode,SumOfItemAmount,isnull(DiscountWord,'') as DiscountWord,DiscountAmount,AfterDiscount,BeforeTaxAmount,TaxAmount,TotalAmount,ExceptTaxAmount,ZeroTaxAmount,PettyCashAmount,SumCashAmount,SumChqAmount,SumBankAmount,DepositIncTax,SumOfDeposit1,SumOfDeposit2,SumOfWTax,NetDebtAmount,HomeAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,BillBalance,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,isnull(GLFormat,'') as GLFormat,GLStartPosting,GRBillStatus,GRIRBillStatus,IsCancel,IsCreditNote,IsDebitNote,IsCompleteSave,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(RecurName,'') as RecurName,ExchangeProfit,isnull(a.ConfirmCode,'') as ConfirmCode,isnull(a.ConfirmDateTime,'') as ConfirmDateTime,isnull(a.CancelCode,'') as CancelCode,isnull(a.CancelDateTime,'') as CancelDateTime,PayBillAmount,isnull(RefDocNo,'') as RefDocNo,isnull(a.CreatorCode,'') as CreatorCode,isnull(a.CreateDateTime,'') as CreateDateTime,isnull(a.LastEditorCode,'') as LastEditorCode,isnull(a.LastEditDateT,'') as LastEditDateT from dbo.bcapinvoice a   with (nolock) left join dbo.bcap b with (nolock) on a.apcode = b.code where (docno  like '%'+?+'%' or apcode like '%'+?+'%' or b.name1 like '%'+?+'%' ) order by docno `
	fmt.Println("sql=", sql)
	err = db.Select(&apvs, sql, keyword, keyword, keyword)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	for _, sub := range apvs {
		sqlsub := `set dateformat dmy    select MyType,StockType,isnull(ItemCode,'') as ItemCode,isnull(MyDescription,'') as ItemName,isnull(ItemName,'') as ItemName,isnull(WHCode,'') as WHCode,isnull(ShelfCode,'') as ShelfCode,CNQty,GRRemainQty,Qty,Price,isnull(DiscountWord,'') as DiscountWord,DiscountAmount,Amount,NetAmount,HomeAmount,BalanceAmount,SumOfExpCost,isnull(UnitCode,'') as UnitCode,isnull(PORefNo,'') as PORefNo,isnull(IRRefNo,'') as IRRefNo,IsCancel,LineNumber,isnull(BarCode,'') as BarCode,PORefLinenum,AVERAGECOST,isnull(LotNumber,'') as LotNumber,SumOfCost,isnull(PackingRate1,1) as PackingRate1,isnull(PackingRate2,1) as PackingRate2 from dbo.bcapinvoicesub  with (nolock) where docno = ? order by LineNumber`
		fmt.Println("sqlsub=", sqlsub)
		err = db.Select(&sub.Subs, sqlsub, sub.DocNo)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}
	return apvs, nil
}
