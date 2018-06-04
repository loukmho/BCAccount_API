package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"errors"
	m "github.com/loukmho/bcaccount_api/model"
)

type Payment struct {
	DocNo           string `json:"doc_no" db:"DocNo"`
	TaxNo           string `json:"tax_no" db:"TaxNo"`
	DocDate         string `json:"doc_date" db:"DocDate"`
	ApCode          string `json:"ap_code" db:"ApCode"`
	CreatorCode     string `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string `json:"last_edit_date_t" db:"LastEditDateT"`
	DepartCode      string `json:"depart_code" db:"DepartCode"`
	SumOfInvoice    float64 `json:"sum_of_invoice" db:"SumOfInvoice"`
	SumOfDebitNote  float64 `json:"sum_of_debit_note" db:"SumOfDebitNote"`
	SumOfCreditNote float64 `json:"sum_of_credit_note" db:"SumOfCreditNote"`
	BeforeTaxAmount float64 `json:"before_tax_amount" db:"BeforeTaxAmount"`
	TaxAmount       float64 `json:"tax_amount" db:"TaxAmount"`
	TotalAmount     float64 `json:"total_amount" db:"TotalAmount"`
	DiscountAmount  float64 `json:"discount_amount" db:"DiscountAmount"`
	NetAmount       float64 `json:"net_amount" db:"NetAmount"`
	PettyCashAmount float64 `json:"petty_cash_amount" db:"PettyCashAmount"`
	SumCashAmount   float64 `json:"sum_cash_amount" db:"SumCashAmount"`
	SumChqAmount    float64 `json:"sum_chq_amount" db:"SumChqAmount"`
	SumBankAmount   float64 `json:"sum_bank_amount" db:"SumBankAmount"`
	SumOfWTax       float64 `json:"sum_of_w_tax" db:"SumOfWTax"`
	GLFormat        string `json:"gl_format" db:"GLFormat"`
	IsCancel        int `json:"is_cancel" db:"IsCancel"`
	MyDescription   string `json:"my_description" db:"MyDescription"`
	AllocateCode    string `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string `json:"project_code" db:"ProjectCode"`
	BillGroup       string `json:"bill_group" db:"BillGroup"`
	SumHomeAmount1  float64 `json:"sum_home_amount_1" db:"SumHomeAmount1"`
	SumHomeAmount2  float64 `json:"sum_home_amount_2" db:"SumHomeAmount2"`
	OtherIncome     float64 `json:"other_income" db:"OtherIncome"`
	OtherExpense    float64 `json:"other_expense" db:"OtherExpense"`
	ExcessAmount1   float64 `json:"excess_amount_1" db:"ExcessAmount1"`
	ExcessAmount2   float64 `json:"excess_amount_2" db:"ExcessAmount2"`
	CurrencyCode    string `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate    float64 `json:"exchange_rate" db:"ExchangeRate"`
	IsCompleteSave  int `json:"is_complete_save" db:"IsCompleteSave"`
	IsConfirm       int `json:"is_confirm" db:"IsConfirm"`
	RecurName       string `json:"recur_name" db:"RecurName"`
	ConfirmCode     string `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string `json:"cancel_date_time" db:"CancelDateTime"`
	SumOfDeposit1   float64 `json:"sum_of_deposit_1" db:"SumOfDeposit1"`
	SumOfDeposit2   float64 `json:"sum_of_deposit_2" db:"SumOfDeposit2"`
	HomeAmountInv   float64 `json:"home_amount_inv" db:"HomeAmountInv"`
	HomeAmountDebt  float64 `json:"home_amount_debt" db:"HomeAmountDebt"`
	HomeAmountCrd   float64 `json:"home_amount_crd" db:"HomeAmountCrd"`
	UserCode string `json:"user_code" db:"UserCode"`
	ListPmtPayMoney
	Chqs            []*ListPmtChqOut      `json:"chqs"`
	Subs            []*PmtItem `json:"subs"`
}

type PmtItem struct {
	InvoiceDate   string `json:"invoice_date" db:"InvoiceDate"`
	InvoiceNo     string `json:"invoice_no" db:"InvoiceNo"`
	InvBalance    float64 `json:"inv_balance" db:"InvBalance"`
	PayAmount     float64 `json:"pay_amount" db:"PayAmount"`
	MyDescription string `json:"my_description" db:"MyDescription"`
	BillType      int `json:"bill_type" db:"BillType"`
	PaybillNo     string `json:"paybill_no" db:"PaybillNo"`
	RefType       int `json:"ref_type" db:"RefType"`
	IsCancel      int `json:"is_cancel" db:"IsCancel"`
	LineNumber    int `json:"line_number" db:"LineNumber"`
	HomeAmount1   float64 `json:"home_amount_1" db:"HomeAmount1"`
	HomeAmount2   float64 `json:"home_amount_2" db:"HomeAmount2"`
}

type ListPmtPayMoney struct {
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

type ListPmtChqOut struct {
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



func (pmt *Payment) InsertAndEditPayment(db *sqlx.DB) error {
	var check_exist int
	var sum_pay_amount float64
	var count_item int

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")

	sum_pay_amount = (pmt.SumCashAmount + pmt.SumChqAmount + pmt.SumBankAmount + pmt.OtherExpense) - pmt.OtherIncome

	for _, Subs := range pmt.Subs {
		if (Subs.InvoiceNo != "") {
			count_item = count_item + 1
		}
	}

	sqlexist := `select count(docno) as check_exist from dbo.bcpayment where docno = ?`
	err := db.Get(&check_exist, sqlexist, pmt.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil
	}

	switch {
	case pmt.DocNo == "":
		return errors.New("Docno is null")
	case pmt.ApCode == "":
		return errors.New("Arcode is null")
	case pmt.DocDate == "":
		return errors.New("Docdate is null")
	case pmt.IsCancel == 1:
		return errors.New("Docno is cancel")
	case pmt.IsConfirm == 1:
		return errors.New("Docno is confirm")
	case count_item == 0:
		return errors.New("Docno is not have item")
	case (pmt.SumCashAmount == 0 && pmt.SumChqAmount == 0 && pmt.SumBankAmount == 0):
		return errors.New("Docno not set money to another type payment")
	case pmt.TotalAmount == 0:
		return errors.New("TotalAmount = 0")
	case sum_pay_amount > pmt.TotalAmount:
		return errors.New("Rec money is over totalamount")
	case pmt.SumChqAmount != 0 && len(pmt.Chqs) == 0:
		return errors.New("Docno not have chq use data")
	}

	if pmt.ExchangeRate == 0 {
		pmt.ExchangeRate = def.ExchangeRateDefault
	}

	if pmt.SaveFrom == 0 {
		pmt.SaveFrom = def.ArInvoiceSaveFrom
	}

	pmt.TotalAmount = pmt.TotalAmount + pmt.OtherIncome - (pmt.SumCashAmount + pmt.SumCreditAmount + pmt.SumChqAmount + pmt.SumBankAmount + pmt.OtherExpense)

	if (((pmt.TotalAmount + pmt.OtherIncome - (pmt.SumCashAmount + pmt.SumCreditAmount + pmt.SumChqAmount + pmt.SumBankAmount + pmt.OtherExpense))) != 0) {
		return errors.New("Docno have money remain")
	}

	if (pmt.Source == 0) {
		pmt.Source = def.Receipt1Source
	}

	if (pmt.GLFormat == "") {
		pmt.GLFormat = def.Receipt1GLFormat
	} else {
		pmt.GLFormat = def.ArInvoiceCreditGLFormat
	}

	if (pmt.BookCode == "") {
		pmt.BookCode = def.Receipt1BookCode
	}

	if (pmt.TaxNo == "") {
		pmt.TaxNo = pmt.DocNo
	}

	pmt.TaxDate = pmt.DocDate

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

	sum_pay_amount = (rcp.SumCashAmount + rcp.SumCreditAmount + rcp.SumChqAmount + rcp.SumBankAmount + rcp.OtherExpense) - rcp.OtherIncome

	switch {
	case rcp.DocNo == "":
		return errors.New("Docno is null")
	case rcp.ArCode == "":
		return errors.New("Arcode is null")
	case rcp.TotalAmount == 0:
		return errors.New("Totalamount is 0")
	case rcp.DocDate == "":
		return errors.New("Docdate is null")
	case rcp.SumCashAmount == 0 && rcp.SumCreditAmount == 0 && rcp.SumChqAmount == 0 && rcp.SumBankAmount == 0:
		return errors.New("Docno not set money to another type payment")
	case sum_pay_amount > rcp.TotalAmount:
		return errors.New("Rec money is over totalamount")
	case rcp.IsCancel == 1:
		return errors.New("Docno is cancel can not edit data")
	case rcp.IsConfirm == 1:
		return errors.New("Docno is confirm can not edit data")
	case rcp.SumCreditAmount != 0 && len(rcp.Cdcs) == 0:
		return errors.New("Creditcard not have credittype or confirmno or creditrefno")
	case rcp.SumChqAmount != 0 && len(rcp.Chqs) == 0:
		return errors.New("Chq not have BankCode or BankBranchCode or creditrefno")
	case rcp.SumBankAmount != 0 && rcp.BankRefNo == "":
		return errors.New("Bank transfer not have BankRefNo")
	}

	if (rcp.TaxNo == "") {
		rcp.TaxNo = rcp.DocNo
	}

	if (check_exist == 0) {
		//Insert//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		rcp.CreatorCode = rcp.UserCode
		rcp.NetAmount = rcp.TotalAmount

		sql := `Insert into dbo.BCPayment(DocNo,TaxNo,DocDate,ApCode,CreatorCode,CreateDateTime,DepartCode,SumOfInvoice,SumOfDebitNote,SumOfCreditNote,BeforeTaxAmount,TaxAmount,TotalAmount,DiscountAmount,NetAmount,PettyCashAmount,SumCashAmount,SumChqAmount,SumBankAmount,SumOfWTax,GLFormat,IsCancel,MyDescription,AllocateCode,ProjectCode,BillGroup,SumHomeAmount1,SumHomeAmount2,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,CurrencyCode,ExchangeRate,IsCompleteSave,IsConfirm,RecurName,SumOfDeposit1,SumOfDeposit2,HomeAmountInv,HomeAmountDebt,HomeAmountCrd) values(?,?,?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sql, pmt.DocNo,pmt.TaxNo,pmt.DocDate,pmt.ApCode,pmt.CreatorCode,pmt.DepartCode,pmt.SumOfInvoice,pmt.SumOfDebitNote,pmt.SumOfCreditNote,pmt.BeforeTaxAmount,pmt.TaxAmount,pmt.TotalAmount,pmt.DiscountAmount,pmt.NetAmount,pmt.PettyCashAmount,pmt.SumCashAmount,pmt.SumChqAmount,pmt.SumBankAmount,pmt.SumOfWTax,pmt.GLFormat,pmt.IsCancel,pmt.MyDescription,pmt.AllocateCode,pmt.ProjectCode,pmt.BillGroup,pmt.SumHomeAmount1,pmt.SumHomeAmount2,pmt.OtherIncome,pmt.OtherExpense,pmt.ExcessAmount1,pmt.ExcessAmount2,pmt.CurrencyCode,pmt.ExchangeRate,pmt.IsCompleteSave,pmt.IsConfirm,pmt.RecurName,pmt.SumOfDeposit1,pmt.SumOfDeposit2,pmt.HomeAmountInv,pmt.HomeAmountDebt,pmt.HomeAmountCrd)
		if err != nil {
			return err
		}

		//sql := `Insert into dbo.BCReceipt1(DocNo,TaxNo,DocDate,CreatorCode,CreateDateTime,ArCode,SaleCode,DepartCode,SumOfInvoice,SumOfDebitNote,SumOfCreditNote,BeforeTaxAmount,TaxAmount,TotalAmount,DiscountAmount,NetAmount,SumCashAmount,SumChqAmount,SumCreditAmount,ChargeAmount,ChangeAmount,SumBankAmount,SumOfWTax,GLFormat,IsCancel,MyDescription,AllocateCode,ProjectCode,BillGroup,IsCompleteSave,SumHomeAmount1,SumHomeAmount2,DebtLimitReturn,CurrencyCode,ExchangeRate,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,IsConfirm,RecurName) values(?,?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		//_, err := db.Exec(sql, rcp.DocNo, rcp.TaxNo, rcp.DocDate, rcp.CreatorCode, rcp.ArCode, rcp.SaleCode, rcp.DepartCode, rcp.SumOfInvoice, rcp.SumOfDebitNote, rcp.SumOfCreditNote, rcp.BeforeTaxAmount, rcp.TaxAmount, rcp.TotalAmount, rcp.DiscountAmount, rcp.NetAmount, rcp.SumCashAmount, rcp.SumChqAmount, rcp.SumCreditAmount, rcp.ChargeAmount, rcp.ChangeAmount, rcp.SumBankAmount, rcp.SumOfWTax, rcp.GLFormat, rcp.IsCancel, rcp.MyDescription, rcp.AllocateCode, rcp.ProjectCode, rcp.BillGroup, rcp.IsCompleteSave, rcp.SumHomeAmount1, rcp.SumHomeAmount2, rcp.DebtLimitReturn, rcp.CurrencyCode, rcp.ExchangeRate, rcp.OtherIncome, rcp.OtherExpense, rcp.ExcessAmount1, rcp.ExcessAmount2, rcp.IsConfirm, rcp.RecurName)
		//fmt.Println("sql =", sql, rcp.DocNo, rcp.DocDate, rcp.TaxNo, rcp.ArCode, rcp.SaleCode)
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

	sqldel := `delete dbo.BCOutputTax where docno = ?`
	_, err = db.Exec(sqldel, rcp.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}

	sqltax := "insert into dbo.BCOutputTax(SaveFrom,DocNo,BookCode,Source,DocDate,TaxDate,TaxNo,ArCode,ShortTaxDesc,TaxRate,Process,BeforeTaxAmount,TaxAmount,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?,?,'รับชำระ',?,1,?,?,?,getdate())"
	_, err = db.Exec(sqltax, rcp.SaveFrom, rcp.DocNo, rcp.BookCode, rcp.Source, rcp.DocDate, rcp.TaxDate, rcp.TaxNo, rcp.ArCode, rcp.TaxRate, rcp.BeforeTaxAmount, rcp.TaxAmount, rcp.CreatorCode)
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

func (pmt *Payment) SearchPaymentByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select DocNo, isnull(TaxNo,'') as taxno, DocDate, isnull(ApCode,'') as ApCode, isnull(CreatorCode,'') as CreatorCode, isnull(CreateDateTime,'') as CreateDateTime, isnull(LastEditorCode,'') as LastEditorCode, isnull(LastEditDateT,'') as LastEditDateT, isnull(DepartCode,'') as DepartCode, SumOfInvoice, SumOfDebitNote, SumOfCreditNote, BeforeTaxAmount, TaxAmount,TotalAmount, DiscountAmount, NetAmount, PettyCashAmount, SumCashAmount, SumChqAmount, SumBankAmount, SumOfWTax, isnull(GLFormat,'') as GLFormat, IsCancel, isnull(MyDescription,'') as MyDescription, isnull(AllocateCode,'') as AllocateCode, isnull(ProjectCode,'') as ProjectCode, isnull(BillGroup,'') as BillGroup, SumHomeAmount1, SumHomeAmount2, OtherIncome, OtherExpense, ExcessAmount1, ExcessAmount2, isnull(CurrencyCode,'') as CurrencyCode, ExchangeRate, IsCompleteSave, IsConfirm, isnull(RecurName,'') as RecurName, isnull(ConfirmCode,'') as ConfirmCode, isnull(ConfirmDateTime,'') as ConfirmDateTime, isnull(CancelCode,'') as CancelCode, isnull(CancelDateTime,'') as CancelDateTime, SumOfDeposit1, SumOfDeposit2, HomeAmountInv, HomeAmountDebt, HomeAmountCrd from dbo.BCPayment a  with (nolock) where docno = ? `
	fmt.Println("sql=", sql)
	err := db.Get(pmt, sql, docno)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	sqlsub := `set dateformat dmy    select isnull(InvoiceDate,'') as InvoiceDate, isnull(InvoiceNo,'') as InvoiceNo, InvBalance, PayAmount, isnull(MyDescription,'') as MyDescription, BillType, isnull(PaybillNo,'') as PaybillNo, RefType, IsCancel, LineNumber, HomeAmount1, HomeAmount2 from dbo.BCPaymentsub  with (nolock)  where docno = ? order by LineNumber`
	fmt.Println("sqlsub=", sqlsub)
	err = db.Select(&pmt.Subs, sqlsub, docno)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (pmt *Payment) SearchPaymentByKeyword(db *sqlx.DB, keyword string) (pmts []*Payment, err error) {
	sql := `set dateformat dmy     select DocNo, isnull(TaxNo,'') as taxno, DocDate, isnull(ApCode,'') as ApCode, isnull(CreatorCode,'') as CreatorCode, isnull(CreateDateTime,'') as CreateDateTime, isnull(LastEditorCode,'') as LastEditorCode, isnull(LastEditDateT,'') as LastEditDateT, isnull(DepartCode,'') as DepartCode, SumOfInvoice, SumOfDebitNote, SumOfCreditNote, BeforeTaxAmount, TaxAmount,TotalAmount, DiscountAmount, NetAmount, PettyCashAmount, SumCashAmount, SumChqAmount, SumBankAmount, SumOfWTax, isnull(GLFormat,'') as GLFormat, IsCancel, isnull(MyDescription,'') as MyDescription, isnull(AllocateCode,'') as AllocateCode, isnull(ProjectCode,'') as ProjectCode, isnull(BillGroup,'') as BillGroup, SumHomeAmount1, SumHomeAmount2, OtherIncome, OtherExpense, ExcessAmount1, ExcessAmount2, isnull(CurrencyCode,'') as CurrencyCode, ExchangeRate, IsCompleteSave, IsConfirm, isnull(RecurName,'') as RecurName, isnull(ConfirmCode,'') as ConfirmCode, isnull(ConfirmDateTime,'') as ConfirmDateTime, isnull(CancelCode,'') as CancelCode, isnull(CancelDateTime,'') as CancelDateTime, SumOfDeposit1, SumOfDeposit2, HomeAmountInv, HomeAmountDebt, HomeAmountCrd from dbo.BCPayment where (docno  like '%'+?+'%' or apcode like '%'+?+'%') order by docno `
	fmt.Println("sql=", sql)
	err = db.Select(&pmts, sql, keyword, keyword, keyword)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	for _, sub := range pmts {
		sqlsub := `set dateformat dmy    select isnull(InvoiceDate,'') as InvoiceDate, isnull(InvoiceNo,'') as InvoiceNo, InvBalance, PayAmount, isnull(MyDescription,'') as MyDescription, BillType, isnull(PaybillNo,'') as PaybillNo, RefType, IsCancel, LineNumber, HomeAmount1, HomeAmount2 from dbo.BCPaymentsub  with (nolock) where docno = ? order by LineNumber`
		fmt.Println("sqlsub=", sqlsub)
		err = db.Select(&sub.Subs, sqlsub, sub.DocNo)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}
	return pmts, nil
}

//
//func (pmt *Payment) SearchPaymentByDocNo() error {
//	//sql := `DocNo, TaxNo, DocDate, ApCode, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, DepartCode, SumOfInvoice, SumOfDebitNote, SumOfCreditNote, BeforeTaxAmount, TaxAmount,TotalAmount, DiscountAmount, NetAmount, PettyCashAmount, SumCashAmount, SumChqAmount, SumBankAmount, SumOfWTax, GLFormat, IsCancel, MyDescription, AllocateCode,ProjectCode, BillGroup, SumHomeAmount1, SumHomeAmount2, OtherIncome, OtherExpense, ExcessAmount1, ExcessAmount2, CurrencyCode, ExchangeRate, IsCompleteSave, IsConfirm,RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime, SumOfDeposit1, SumOfDeposit2, HomeAmountInv, HomeAmountDebt, HomeAmountCrd`
//	//sqlsub:=`DocNo, DocDate, ApCode, InvoiceDate, InvoiceNo, InvBalance, PayAmount, MyDescription, BillType, PaybillNo, RefType,IsCancel, LineNumber, HomeAmount1, HomeAmount2`
//	return nil
//}

//paymentsub
