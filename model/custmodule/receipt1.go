package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"errors"
	m "github.com/loukmho/bcaccount_api/model"
)

type Receipt1 struct {
	Source          int                  `json:"source" db:"Source"`
	SaveFrom        int                  `json:"save_from" db:"SaveFrom"`
	DocNo           string               `json:"doc_no" db:"DocNo"`
	RcpOutPutTax
	DocDate         string               `json:"doc_date" db:"DocDate"`
	CreatorCode     string               `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string               `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string               `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string               `json:"last_edit_date_t" db:"LastEditDateT"`
	ArCode          string               `json:"ar_code" db:"ArCode"`
	SaleCode        string               `json:"sale_code" db:"SaleCode"`
	DepartCode      string               `json:"depart_code" db:"DepartCode"`
	SumOfInvoice    float64              `json:"sum_of_invoice" db:"SumOfInvoice"`
	SumOfDebitNote  float64              `json:"sum_of_debit_note" db:"SumOfDebitNote"`
	SumOfCreditNote float64              `json:"sum_of_credit_note" db:"SumOfCreditNote"`
	BeforeTaxAmount float64              `json:"before_tax_amount" db:"BeforeTaxAmount"`
	TaxAmount       float64              `json:"tax_amount" db:"TaxAmount"`
	TotalAmount     float64              `json:"total_amount" db:"TotalAmount"`
	DiscountAmount  float64              `json:"discount_amount" db:"DiscountAmount"`
	NetAmount       float64              `json:"net_amount" db:"NetAmount"`
	SumCashAmount   float64              `json:"sum_cash_amount" db:"SumCashAmount"`
	SumChqAmount    float64              `json:"sum_chq_amount" db:"SumChqAmount"`
	SumCreditAmount float64              `json:"sum_credit_amount" db:"SumCreditAmount"`
	ChargeAmount    float64              `json:"charge_amount" db:"ChargeAmount"`
	ChangeAmount    float64              `json:"change_amount" db:"ChangeAmount"`
	SumBankAmount   float64              `json:"sum_bank_amount" db:"SumBankAmount"`
	SumOfWTax       float64              `json:"sum_of_w_tax" db:"SumOfWTax"`
	ConfirmNo       string               `json:"confirm_no" db:"ConfirmNo"`
	CreditNo        string               `json:"credit_no" db:"CreditNo"`
	GLFormat        string               `json:"gl_format" db:"GLFormat"`
	IsCancel        int                  `json:"is_cancel" db:"IsCancel"`
	MyDescription   string               `json:"my_description" db:"MyDescription"`
	AllocateCode    string               `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string               `json:"project_code" db:"ProjectCode"`
	BillGroup       string               `json:"bill_group" db:"BillGroup"`
	IsCompleteSave  int                  `json:"is_complete_save" db:"IsCompleteSave"`
	SumHomeAmount1  float64              `json:"sum_home_amount_1" db:"SumHomeAmount1"`
	SumHomeAmount2  float64              `json:"sum_home_amount_2" db:"SumHomeAmount2"`
	DebtLimitReturn float64              `json:"debt_limit_return" db:"DebtLimitReturn"`
	CurrencyCode    string               `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate    float64              `json:"exchange_rate" db:"ExchangeRate"`
	OtherIncome     float64              `json:"other_income" db:"OtherIncome"`
	OtherExpense    float64              `json:"other_expense" db:"OtherExpense"`
	ExcessAmount1   float64              `json:"excess_amount_1" db:"ExcessAmount1"`
	ExcessAmount2   float64              `json:"excess_amount_2" db:"ExcessAmount2"`
	IsConfirm       int                  `json:"is_confirm" db:"IsConfirm"`
	RecurName       string               `json:"recur_name" db:"RecurName"`
	ConfirmCode     string               `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string               `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string               `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string               `json:"cancel_date_time" db:"CancelDateTime"`
	UserCode        string               `json:"user_code"`
	ListRcpRecMoney
	Subs            []*RcpItem           `json:"subs"`
	Cdcs            []*ListRcpCreditCard `json:"cdcs"`
	Chqs            []*ListRcpChqIn      `json:"chqs"`
}

type RcpOutPutTax struct {
	TaxNo    string  `json:"tax_no" db:"TaxNo"`
	TaxDate  string  `json:"tax_date" db:"TaxDate"`
	BookCode string  `json:"book_code" db:"BookCode"`
	TaxRate  float64 `json:"tax_rate" db:"TaxRate"`
}

type RcpItem struct {
	InvoiceDate     string  `json:"invoice_date" db:"InvoiceDate"`
	InvoiceNo       string  `json:"invoice_no" db:"InvoiceNo"`
	InvBalance      float64 `json:"inv_balance" db:"InvBalance"`
	PayAmount       float64 `json:"pay_amount" db:"PayAmount"`
	DebtLimitReturn float64 `json:"debt_limit_return" db:"DebtLimitReturn"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	LineNumber      int     `json:"line_number" db:"LineNumber"`
	BillType        int     `json:"bill_type" db:"BillType"`
	RefNo           string  `json:"ref_no" db:"RefNo"`
	RefType         int     `json:"ref_type" db:"RefType"`
	HomeAmount1     float64 `json:"home_amount_1" db:"HomeAmount1"`
	HomeAmount2     float64 `json:"home_amount_2" db:"HomeAmount2"`
}

type ListRcpRecMoney struct {
	CreditType     string `json:"credit_type" db:"CreditType"`
	CofirmNo       string `json:"cofirm_no" db:"CofirmNo"`
	CreditNo       string `json:"credit_no" db:"CreditNo"`
	CreditDueDate  string `json:"credit_due_date" db:"CreditDueDate"`
	BankCode       string `json:"bank_code" db:"BookCode"`
	BankBranchCode string `json:"bank_branch_code" db:"BankBranchCode"`
	BankRefNo      string `json:"bank_ref_no" db:"BankRefNo"`
	TransBankDate  string `json:"trans_bank_date" db:"TransBankDate"`
	RefDate        string `json:"ref_date" db:"RefDate"`
}

type ListRcpChqIn struct {
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

type ListRcpCreditCard struct {
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

func (rcp *Receipt1) InsertAndEditReceipt1(db *sqlx.DB) error {
	var check_exist int
	var sum_pay_amount float64
	var count_item int

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")
	if rcp.TaxRate == 0 {
		rcp.TaxRate = def.TaxRateDefault
	}

	fmt.Println("TaxRate = ", def.TaxRateDefault)

	sum_pay_amount = (rcp.SumCashAmount + rcp.SumCreditAmount + rcp.SumChqAmount + rcp.SumBankAmount + rcp.OtherExpense) - rcp.OtherIncome

	for _, Subs := range rcp.Subs {
		if (Subs.InvoiceNo != "") {
			count_item = count_item + 1
		}
	}

	sqlexist := `select count(docno) as check_exist from dbo.bcreceipt1 where docno = ?`
	err := db.Get(&check_exist, sqlexist, rcp.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil
	}

	switch {
	case rcp.DocNo == "":
		return errors.New("Docno is null")
	case rcp.ArCode == "":
		return errors.New("Arcode is null")
	case rcp.DocDate == "":
		return errors.New("Docdate is null")
	case rcp.IsCancel == 1:
		return errors.New("Docno is cancel")
	case rcp.IsConfirm == 1:
		return errors.New("Docno is confirm")
	case count_item == 0:
		return errors.New("Docno is not have item")
	case (rcp.SumCashAmount == 0 && rcp.SumCreditAmount == 0 && rcp.SumChqAmount == 0 && rcp.SumBankAmount == 0):
		return errors.New("Docno not set money to another type payment")
	case rcp.TotalAmount == 0:
		return errors.New("TotalAmount = 0")
	case sum_pay_amount > rcp.TotalAmount:
		return errors.New("Rec money is over totalamount")
	case rcp.SumChqAmount != 0 && len(rcp.Chqs) == 0:
		return errors.New("Docno not have chq use data")
	case rcp.SumCreditAmount != 0 && (rcp.CreditNo == "" || rcp.CofirmNo == "" || rcp.CreditType == "" || rcp.BankCode == "" || rcp.BankBranchCode == ""):
		return errors.New("credit card not have credittype,confirmno,creditno,bankcode,bankbranch data not complete")
	case rcp.SumBankAmount != 0 && rcp.BankRefNo == "":
		return errors.New("Bank transfer not have accountcode data not complete")
	}

	if rcp.ExchangeRate == 0 {
		rcp.ExchangeRate = def.ExchangeRateDefault
	}

	if rcp.SaveFrom == 0 {
		rcp.SaveFrom = def.ArInvoiceSaveFrom
	}

	rcp.TotalAmount = rcp.TotalAmount + rcp.OtherIncome - (rcp.SumCashAmount + rcp.SumCreditAmount + rcp.SumChqAmount + rcp.SumBankAmount + rcp.OtherExpense)

	if (((rcp.TotalAmount + rcp.OtherIncome - (rcp.SumCashAmount + rcp.SumCreditAmount + rcp.SumChqAmount + rcp.SumBankAmount + rcp.OtherExpense))) != 0) {
		return errors.New("Docno have money remain")
	}

	if (rcp.Source == 0) {
		rcp.Source = def.Receipt1Source
	}

	if (rcp.GLFormat == "") {
		rcp.GLFormat = def.Receipt1GLFormat
	}

	if (rcp.BookCode == "") {
		rcp.BookCode = def.Receipt1BookCode
	}

	if (rcp.TaxNo == "") {
		rcp.TaxNo = rcp.DocNo
	}

	rcp.TaxDate = rcp.DocDate

	if (rcp.SumCreditAmount != 0 && rcp.CreditDueDate == "") {
		rcp.CreditDueDate = rcp.DocDate
	}

	if (rcp.SumBankAmount != 0 && rcp.TransBankDate == "") {
		rcp.TransBankDate = rcp.DocDate
	}

	fmt.Println("UserCode = ", rcp.UserCode)

	if (rcp.IsCompleteSave == 0) {
		rcp.IsCompleteSave = 1
	}


	if (check_exist == 0) {
		//Insert//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		rcp.CreatorCode = rcp.UserCode
		rcp.NetAmount = rcp.TotalAmount

		sql := `Insert into dbo.BCReceipt1(DocNo,TaxNo,DocDate,CreatorCode,CreateDateTime,ArCode,SaleCode,DepartCode,SumOfInvoice,SumOfDebitNote,SumOfCreditNote,BeforeTaxAmount,TaxAmount,TotalAmount,DiscountAmount,NetAmount,SumCashAmount,SumChqAmount,SumCreditAmount,ChargeAmount,ChangeAmount,SumBankAmount,SumOfWTax,GLFormat,IsCancel,MyDescription,AllocateCode,ProjectCode,BillGroup,IsCompleteSave,SumHomeAmount1,SumHomeAmount2,DebtLimitReturn,CurrencyCode,ExchangeRate,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,IsConfirm,RecurName) values(?,?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sql, rcp.DocNo, rcp.TaxNo, rcp.DocDate, rcp.CreatorCode, rcp.ArCode, rcp.SaleCode, rcp.DepartCode, rcp.SumOfInvoice, rcp.SumOfDebitNote, rcp.SumOfCreditNote, rcp.BeforeTaxAmount, rcp.TaxAmount, rcp.TotalAmount, rcp.DiscountAmount, rcp.NetAmount, rcp.SumCashAmount, rcp.SumChqAmount, rcp.SumCreditAmount, rcp.ChargeAmount, rcp.ChangeAmount, rcp.SumBankAmount, rcp.SumOfWTax, rcp.GLFormat, rcp.IsCancel, rcp.MyDescription, rcp.AllocateCode, rcp.ProjectCode, rcp.BillGroup, rcp.IsCompleteSave, rcp.SumHomeAmount1, rcp.SumHomeAmount2, rcp.DebtLimitReturn, rcp.CurrencyCode, rcp.ExchangeRate, rcp.OtherIncome, rcp.OtherExpense, rcp.ExcessAmount1, rcp.ExcessAmount2, rcp.IsConfirm, rcp.RecurName)
		fmt.Println("sql =", sql, rcp.DocNo, rcp.DocDate, rcp.TaxNo, rcp.ArCode, rcp.SaleCode)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}

	} else {

		//Update/////////////////////////////////////////////////////////////////////////////////////////////////////////
		rcp.LastEditorCode = rcp.UserCode
		rcp.NetAmount = rcp.TotalAmount

		sql := `Update dbo.BCReceipt1 set TaxNo=?,DocDate=?,LastEditorCode=?,LastEditDateT=getdate(),ArCode=?,SaleCode=?,DepartCode=?,SumOfInvoice=?,SumOfDebitNote=?,SumOfCreditNote=?,BeforeTaxAmount=?,TaxAmount=?,TotalAmount=?,DiscountAmount=?,NetAmount=?,SumCashAmount=?,SumChqAmount=?,SumCreditAmount=?,ChargeAmount=?,ChangeAmount=?,SumBankAmount=?,SumOfWTax=?,GLFormat=?,IsCancel=?,MyDescription=?,AllocateCode=?,ProjectCode=?,BillGroup=?,IsCompleteSave=?,SumHomeAmount1=?,SumHomeAmount2=?,DebtLimitReturn=?,CurrencyCode=?,ExchangeRate=?,OtherIncome=?,OtherExpense=?,ExcessAmount1=?,ExcessAmount2=?,IsConfirm=?,RecurName=? where docno = ?`
		_, err := db.Exec(sql, rcp.TaxNo, rcp.DocDate, rcp.LastEditorCode, rcp.ArCode, rcp.SaleCode, rcp.DepartCode, rcp.SumOfInvoice, rcp.SumOfDebitNote, rcp.SumOfCreditNote, rcp.BeforeTaxAmount, rcp.TaxAmount, rcp.TotalAmount, rcp.DiscountAmount, rcp.NetAmount, rcp.SumCashAmount, rcp.SumChqAmount, rcp.SumCreditAmount, rcp.ChargeAmount, rcp.ChangeAmount, rcp.SumBankAmount, rcp.SumOfWTax, rcp.GLFormat, rcp.IsCancel, rcp.MyDescription, rcp.AllocateCode, rcp.ProjectCode, rcp.BillGroup, rcp.IsCompleteSave, rcp.SumHomeAmount1, rcp.SumHomeAmount2, rcp.DebtLimitReturn, rcp.CurrencyCode, rcp.ExchangeRate, rcp.OtherIncome, rcp.OtherExpense, rcp.ExcessAmount1, rcp.ExcessAmount2, rcp.IsConfirm, rcp.RecurName, rcp.DocNo)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}
	}

	sql_del_sub := `delete dbo.BCReceiptSub1 where docno = ?`
	_, err = db.Exec(sql_del_sub, rcp.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}

	var vLineNumber int

	for _, sub := range rcp.Subs {
		fmt.Println("ItemSub")
		sub.LineNumber = vLineNumber
		sub.IsCancel = 0
		sub.HomeAmount2 = sub.HomeAmount2

		sqlsub := ` insert into dbo.BCReceiptSub1(DocNo,DocDate,ArCode,SaleCode,InvoiceDate,InvoiceNo,InvBalance,PayAmount,DebtLimitReturn,MyDescription,IsCancel,LineNumber,BillType,RefNo,RefType,HomeAmount1,HomeAmount2) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlsub, rcp.DocNo,rcp.DocDate,rcp.ArCode,rcp.SaleCode,sub.InvoiceDate,sub.InvoiceNo,sub.InvBalance,sub.PayAmount,sub.DebtLimitReturn,sub.MyDescription,sub.IsCancel,sub.LineNumber,sub.BillType,sub.RefNo,sub.RefType,sub.HomeAmount1,sub.HomeAmount2)
		fmt.Println("sqltax = ", sqlsub)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}

		vLineNumber = vLineNumber + 1
	}

	sqldel := `delete dbo.BCOutputTax where docno = ?`
	_, err = db.Exec(sqldel, rcp.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}

	sqltax := "insert into dbo.BCOutputTax(SaveFrom,DocNo,BookCode,Source,DocDate,TaxDate,TaxNo,ArCode,ShortTaxDesc,TaxRate,Process,BeforeTaxAmount,TaxAmount,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?,?,'รับชำระ',?,1,?,?,?,getdate())"
	_, err = db.Exec(sqltax, rcp.SaveFrom, rcp.DocNo, rcp.BookCode, rcp.Source, rcp.DocDate, rcp.TaxDate, rcp.TaxNo, rcp.ArCode, rcp.TaxRate, rcp.BeforeTaxAmount, rcp.TaxAmount, rcp.UserCode)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}

	sqlrecdel := `delete dbo.BCRecMoney where docno = ?`
	_, err = db.Exec(sqlrecdel, rcp.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}

	var my_description_recmoney string

	my_description_recmoney = "รับชำระหนี้"

	fmt.Println("RecMoney")
	var linenumber int

	if (rcp.SumCashAmount != 0) { //subs.PaymentType == 0:
		fmt.Println("SumCashAmount")
		sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,PaymentType,SaveFrom,LineNumber,ProjectCode,DepartCode,SaleCode,MyDescription) values(?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlrec, rcp.DocNo, rcp.DocDate, rcp.ArCode, rcp.ExchangeRate, rcp.SumCashAmount, 0, rcp.SaveFrom, linenumber, rcp.ProjectCode, rcp.DepartCode, rcp.SaleCode, my_description_recmoney)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}
	}
	//case dp.SumCreditAmount != 0: //subs.PaymentType == 1:
	if (rcp.SumCreditAmount != 0) {
		fmt.Println("SumCreditAmount")
		if (rcp.SumCashAmount != 0) {
			linenumber = 1
		} else {
			linenumber = 0
		}
		sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,ChqTotalAmount,PaymentType,SaveFrom,CreditType,ConfirmNo,LineNumber,RefNo,BankCode,BankBranchCode,ProjectCode,DepartCode,SaleCode,MyDescription,RefDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlrec, rcp.DocNo, rcp.DocDate, rcp.ArCode, rcp.ExchangeRate, rcp.SumCreditAmount, rcp.SumCreditAmount, 1, rcp.SaveFrom, rcp.CreditType, rcp.ConfirmNo, linenumber, rcp.CreditNo, rcp.BankCode, rcp.BankBranchCode, rcp.ProjectCode, rcp.DepartCode, rcp.SaleCode, my_description_recmoney, rcp.DocDate)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}
	}

	//case dp.SumChqAmount != 0: //subs.PaymentType == 2:
	if (rcp.SumChqAmount != 0) {
		fmt.Println("SumChqAmount")
		if (rcp.SumCashAmount != 0 && rcp.SumCreditAmount != 0) {
			linenumber = 2
		} else if ((rcp.SumCashAmount != 0 && rcp.SumCreditAmount == 0)) {
			linenumber = 1
		} else if ((rcp.SumCashAmount == 0 && rcp.SumCreditAmount != 0)) {
			linenumber = 1
		} else if ((rcp.SumCashAmount == 0 && rcp.SumCreditAmount == 0)) {
			linenumber = 0
		}

		sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,PaymentType,SaveFrom,LineNumber,RefNo,BankCode,ProjectCode,DepartCode,SaleCode,BankBranchCode,MyDescription,RefDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlrec, rcp.DocNo, rcp.DocDate, rcp.ArCode, rcp.ExchangeRate, rcp.SumChqAmount, 2, rcp.SaveFrom, linenumber, rcp.CreditNo, rcp.BankCode, rcp.ProjectCode, rcp.DepartCode, rcp.SaleCode, rcp.BankBranchCode, my_description_recmoney, rcp.RefDate)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}
	}

	//case dp.SumBankAmount != 0: //subs.PaymentType == 3:
	if (rcp.SumBankAmount != 0) {
		fmt.Println("SumBankAmount")
		if (rcp.SumCashAmount != 0 && rcp.SumCreditAmount != 0 && rcp.SumChqAmount != 0) {
			linenumber = 3
		} else if (rcp.SumCashAmount != 0 && rcp.SumCreditAmount == 0 && rcp.SumChqAmount != 0) {
			linenumber = 2
		} else if (rcp.SumCashAmount == 0 && rcp.SumCreditAmount != 0 && rcp.SumChqAmount != 0) {
			linenumber = 2
		} else if (rcp.SumCashAmount != 0 && rcp.SumCreditAmount != 0 && rcp.SumChqAmount == 0) {
			linenumber = 2
		} else if (rcp.SumCashAmount != 0 && rcp.SumCreditAmount != 0 && rcp.SumChqAmount == 0) {
			linenumber = 2
		} else if (rcp.SumCashAmount != 0 && rcp.SumCreditAmount == 0 && rcp.SumChqAmount == 0) {
			linenumber = 1
		} else if (rcp.SumCashAmount == 0 && rcp.SumCreditAmount != 0 && rcp.SumChqAmount == 0) {
			linenumber = 1
		} else if (rcp.SumCashAmount == 0 && rcp.SumCreditAmount == 0 && rcp.SumChqAmount != 0) {
			linenumber = 1
		} else if (rcp.SumCashAmount == 0 && rcp.SumCreditAmount == 0 && rcp.SumChqAmount == 0) {
			linenumber = 0
		}

		sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,PaymentType,SaveFrom,LineNumber,RefNo,ProjectCode,DepartCode,SaleCode,MyDescription,RefDate,TransBankDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlrec, rcp.DocNo, rcp.DocDate, rcp.ArCode, rcp.ExchangeRate, rcp.SumBankAmount, 3, rcp.SaveFrom, linenumber, rcp.BankRefNo, rcp.ProjectCode, rcp.DepartCode, rcp.SaleCode, my_description_recmoney, rcp.DocDate, rcp.TransBankDate)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}
	}

	if (rcp.SumChqAmount != 0) {
		sqlchqdel := `delete dbo.BCChqIn where docno = ?`
		_, err = db.Exec(sqlchqdel, rcp.DocNo)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}

		for _, c := range rcp.Chqs {
			if ((c.ReceiveDate == "") || (c.DueDate == "")) {
				c.ReceiveDate = rcp.DocDate
				c.DueDate = rcp.DocDate
			}

			sqldep := `insert into dbo.bcchqin(BankCode,ChqNumber,DocNo,ArCode,SaleCode,ReceiveDate,DueDate,BookNo,Status,SaveFrom,StatusDate,StatusDocNo,DepartCode,BankBranchCode,Amount,Balance,MyDescription,ExchangeRate,RefChqRowOrder) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
			_, err = db.Exec(sqldep, c.BankCode, c.ChqNumber, rcp.DocNo, rcp.ArCode, rcp.SaleCode, c.ReceiveDate, c.DueDate, c.BookNo, c.Status, rcp.SaveFrom, c.StatusDate, c.StatusDocNo, rcp.DepartCode, c.BankBranchCode, c.Amount, c.Balance, my_description_recmoney, rcp.ExchangeRate, c.RefChqRowOrder)
			if err != nil {
				fmt.Println("Chq", err.Error())
				return err
			}
		}
	}

	if (rcp.SumCreditAmount != 0) {
		sqlcrddel := `delete dbo.BCCreditCard where docno = ?`
		_, err = db.Exec(sqlcrddel, rcp.DocNo)
		if err != nil {
			return err
		}

		fmt.Println("Cdcs ren =", len((rcp.Cdcs)))
		for _, d := range rcp.Cdcs {
			sqlcrd := `insert into dbo.bccreditcard(BankCode,CreditCardNo,DocNo,ArCode,SaleCode,ReceiveDate,DueDate,BookNo,Status,SaveFrom,StatusDate,StatusDocNo,DepartCode,BankBranchCode,Amount,MyDescription,ExchangeRate,CreditType,ConfirmNo,ChargeAmount,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,getdate())`
			_, err = db.Exec(sqlcrd, d.BankCode, d.CreditCardNo, rcp.DocNo, rcp.ArCode, rcp.SaleCode, d.ReceiveDate, d.DueDate, d.BookNo, d.Status, rcp.SaveFrom, d.StatusDate, d.StatusDocNo, rcp.DepartCode, d.BankBranchCode, d.Amount, my_description_recmoney, rcp.ExchangeRate, d.CreditType, d.ConfirmNo, d.ChargeAmount, rcp.UserCode)
			if err != nil {
				fmt.Println("Credit", err.Error())
				return err
			}
		}

	}

	return nil
}

func (rcp *Receipt1) SearchReceipt1ByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select DocNo,isnull(TaxNo,'') as TaxNo,DocDate,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(ArCode,'') as ArCode,isnull(SaleCode,'') as SaleCode,isnull(DepartCode,'') as DepartCode,SumOfInvoice,SumOfDebitNote,SumOfCreditNote,BeforeTaxAmount,TaxAmount,TotalAmount,DiscountAmount,NetAmount,SumCashAmount,SumChqAmount,SumCreditAmount,ChargeAmount,ChangeAmount,SumBankAmount,SumOfWTax,isnull(GLFormat,'') as GLFormat,IsCancel,isnull(MyDescription,'') as MyDescription,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(BillGroup,'') as BillGroup,IsCompleteSave,SumHomeAmount1,SumHomeAmount2,DebtLimitReturn,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,IsConfirm,RecurName,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime from dbo.bcreceipt1 with (nolock) with (nolock) where docno = ?`
	err := db.Get(rcp, sql,docno)
	if err != nil {
		fmt.Println("Search ", err.Error())
		return err
	}

	sqlsub := `set dateformat dmy     select isnull(InvoiceDate,'') as InvoiceDate, isnull(InvoiceNo,'') as InvoiceNo, InvBalance, PayAmount, DebtLimitReturn, isnull(MyDescription,'') as MyDescription, IsCancel, LineNumber, BillType, isnull(RefNo,'') as RefNo, RefType,  HomeAmount1, HomeAmount2 from dbo.BCReceiptSub1 with (nolock) where docno = ? order by linenumber`
	err = db.Select(rcp.Subs, sqlsub, docno)
	if err != nil {
		fmt.Println("Search Sub ", err.Error())
		return err
	}
	//sql := `DocNo, TaxNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, ArCode, SaleCode, DepartCode, SumOfInvoice, SumOfDebitNote, SumOfCreditNote, BeforeTaxAmount,TaxAmount, TotalAmount, DiscountAmount, NetAmount, SumCashAmount, SumChqAmount, SumCreditAmount, ChargeAmount, ChangeAmount, SumBankAmount, SumOfWTax, GLFormat, IsCancel,MyDescription, AllocateCode, ProjectCode, BillGroup, IsCompleteSave, SumHomeAmount1, SumHomeAmount2, DebtLimitReturn, CurrencyCode, ExchangeRate, OtherIncome,OtherExpense, ExcessAmount1, ExcessAmount2, IsConfirm, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime`
	//sqlsub:= `DocNo, DocDate, ArCode, SaleCode,  InvoiceDate, InvoiceNo, InvBalance, PayAmount, DebtLimitReturn,MyDescription,  IsCancel, LineNumber, BillType, RefNo, RefType,  HomeAmount1, HomeAmount2,`
	return nil
}


func (rcp *Receipt1) SearchReceipt1ByKeyword(db *sqlx.DB, keyword string) (rcps []*Receipt1, err error) {
	sql := `set dateformat dmy     select DocNo,isnull(TaxNo,'') as TaxNo,DocDate,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(ArCode,'') as ArCode,isnull(SaleCode,'') as SaleCode,isnull(DepartCode,'') as DepartCode,SumOfInvoice,SumOfDebitNote,SumOfCreditNote,BeforeTaxAmount,TaxAmount,TotalAmount,DiscountAmount,NetAmount,SumCashAmount,SumChqAmount,SumCreditAmount,ChargeAmount,ChangeAmount,SumBankAmount,SumOfWTax,isnull(GLFormat,'') as GLFormat,IsCancel,isnull(MyDescription,'') as MyDescription,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(BillGroup,'') as BillGroup,IsCompleteSave,SumHomeAmount1,SumHomeAmount2,DebtLimitReturn,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,IsConfirm,RecurName,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime from dbo.bcreceipt1 with (nolock) where (docno  like '%'+?+'%' or arcode like '%'+?+'%' or salecode like '%'+?+'%') order by docno`
	err = db.Select(&rcps, sql, keyword, keyword, keyword)
	if err != nil {
		fmt.Println("Search ", err.Error())
		return nil, err
	}

	for _, sub := range rcps {
		sqlsub := `set dateformat dmy     select isnull(InvoiceDate,'') as InvoiceDate, isnull(InvoiceNo,'') as InvoiceNo, InvBalance, PayAmount, DebtLimitReturn, isnull(MyDescription,'') as MyDescription, IsCancel, LineNumber, BillType, isnull(RefNo,'') as RefNo, RefType,  HomeAmount1, HomeAmount2 from dbo.BCReceiptSub1 with (nolock) where docno = ? order by linenumber`
		err = db.Select(&rcp.Subs, sqlsub, sub.DocNo)
		if err != nil {
			fmt.Println("Search Sub ", err.Error())
			return nil, err
		}
	}
	//sql := `DocNo, TaxNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, ArCode, SaleCode, DepartCode, SumOfInvoice, SumOfDebitNote, SumOfCreditNote, BeforeTaxAmount,TaxAmount, TotalAmount, DiscountAmount, NetAmount, SumCashAmount, SumChqAmount, SumCreditAmount, ChargeAmount, ChangeAmount, SumBankAmount, SumOfWTax, GLFormat, IsCancel,MyDescription, AllocateCode, ProjectCode, BillGroup, IsCompleteSave, SumHomeAmount1, SumHomeAmount2, DebtLimitReturn, CurrencyCode, ExchangeRate, OtherIncome,OtherExpense, ExcessAmount1, ExcessAmount2, IsConfirm, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime`
	//sqlsub:= `DocNo, DocDate, ArCode, SaleCode,  InvoiceDate, InvoiceNo, InvBalance, PayAmount, DebtLimitReturn,MyDescription,  IsCancel, LineNumber, BillType, RefNo, RefType,  HomeAmount1, HomeAmount2,`
	return rcps, nil
}
