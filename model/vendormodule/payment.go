package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"errors"
	m "github.com/loukmho/bcaccount_api/model"
)

type Payment struct {
	SaveFrom        int                   `json:"save_from" db:"SaveFrom"`
	Source          int                   `json:"source" db:"Source"`
	DocNo           string                `json:"doc_no" db:"DocNo"`
	PmtInPutTax
	DocDate         string                `json:"doc_date" db:"DocDate"`
	ApCode          string                `json:"ap_code" db:"ApCode"`
	CreatorCode     string                `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string                `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string                `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string                `json:"last_edit_date_t" db:"LastEditDateT"`
	DepartCode      string                `json:"depart_code" db:"DepartCode"`
	SumOfInvoice    float64               `json:"sum_of_invoice" db:"SumOfInvoice"`
	SumOfDebitNote  float64               `json:"sum_of_debit_note" db:"SumOfDebitNote"`
	SumOfCreditNote float64               `json:"sum_of_credit_note" db:"SumOfCreditNote"`
	BeforeTaxAmount float64               `json:"before_tax_amount" db:"BeforeTaxAmount"`
	TaxAmount       float64               `json:"tax_amount" db:"TaxAmount"`
	TotalAmount     float64               `json:"total_amount" db:"TotalAmount"`
	DiscountAmount  float64               `json:"discount_amount" db:"DiscountAmount"`
	NetAmount       float64               `json:"net_amount" db:"NetAmount"`
	PettyCashAmount float64               `json:"petty_cash_amount" db:"PettyCashAmount"`
	SumCashAmount   float64               `json:"sum_cash_amount" db:"SumCashAmount"`
	SumChqAmount    float64               `json:"sum_chq_amount" db:"SumChqAmount"`
	SumBankAmount   float64               `json:"sum_bank_amount" db:"SumBankAmount"`
	SumOfWTax       float64               `json:"sum_of_w_tax" db:"SumOfWTax"`
	GLFormat        string                `json:"gl_format" db:"GLFormat"`
	IsCancel        int                   `json:"is_cancel" db:"IsCancel"`
	MyDescription   string                `json:"my_description" db:"MyDescription"`
	AllocateCode    string                `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string                `json:"project_code" db:"ProjectCode"`
	BillGroup       string                `json:"bill_group" db:"BillGroup"`
	SumHomeAmount1  float64               `json:"sum_home_amount_1" db:"SumHomeAmount1"`
	SumHomeAmount2  float64               `json:"sum_home_amount_2" db:"SumHomeAmount2"`
	OtherIncome     float64               `json:"other_income" db:"OtherIncome"`
	OtherExpense    float64               `json:"other_expense" db:"OtherExpense"`
	ExcessAmount1   float64               `json:"excess_amount_1" db:"ExcessAmount1"`
	ExcessAmount2   float64               `json:"excess_amount_2" db:"ExcessAmount2"`
	CurrencyCode    string                `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate    float64               `json:"exchange_rate" db:"ExchangeRate"`
	IsCompleteSave  int                   `json:"is_complete_save" db:"IsCompleteSave"`
	IsConfirm       int                   `json:"is_confirm" db:"IsConfirm"`
	RecurName       string                `json:"recur_name" db:"RecurName"`
	ConfirmCode     string                `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string                `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string                `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string                `json:"cancel_date_time" db:"CancelDateTime"`
	SumOfDeposit1   float64               `json:"sum_of_deposit_1" db:"SumOfDeposit1"`
	SumOfDeposit2   float64               `json:"sum_of_deposit_2" db:"SumOfDeposit2"`
	HomeAmountInv   float64               `json:"home_amount_inv" db:"HomeAmountInv"`
	HomeAmountDebt  float64               `json:"home_amount_debt" db:"HomeAmountDebt"`
	HomeAmountCrd   float64               `json:"home_amount_crd" db:"HomeAmountCrd"`
	UserCode        string                `json:"user_code" db:"UserCode"`
	ListPmtPayMoney
	Deps            []*ListPmtDepositUsed `json:"deps"`
	Chqs            []*ListPmtChqOut      `json:"chqs"`
	Subs            []*PmtItem            `json:"subs"`
}

type PmtInPutTax struct {
	TaxNo    string `json:"tax_no" db:"TaxNo"`
	TaxDate  string `json:"tax_date" db:"TaxDate"`
	BookCode string `json:"book_code" db:"BookCode"`
}

type PmtItem struct {
	InvoiceDate   string  `json:"invoice_date" db:"InvoiceDate"`
	InvoiceNo     string  `json:"invoice_no" db:"InvoiceNo"`
	InvBalance    float64 `json:"inv_balance" db:"InvBalance"`
	PayAmount     float64 `json:"pay_amount" db:"PayAmount"`
	MyDescription string  `json:"my_description" db:"MyDescription"`
	BillType      int     `json:"bill_type" db:"BillType"`
	PaybillNo     string  `json:"paybill_no" db:"PaybillNo"`
	RefType       int     `json:"ref_type" db:"RefType"`
	IsCancel      int     `json:"is_cancel" db:"IsCancel"`
	LineNumber    int     `json:"line_number" db:"LineNumber"`
	HomeAmount1   float64 `json:"home_amount_1" db:"HomeAmount1"`
	HomeAmount2   float64 `json:"home_amount_2" db:"HomeAmount2"`
}

type ListPmtDepositUsed struct {
	DepositNo  string  `json:"deposit_no" db:"DepositNo"`
	Balance    float64 `json:"balance" db:"Balance"`
	Amount     float64 `json:"amount" db:"Amount"`
	NetAmount  float64 `json:"net_amount" db:"NetAmount"`
	LineNumber int     `json:"line_number" db:"LineNumber"`
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
	var count_item int
	var sum_pay_amount float64

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")

	fmt.Println("TaxRate = ", def.TaxRateDefault)

	sum_pay_amount = (pmt.SumCashAmount + pmt.SumChqAmount + pmt.SumBankAmount + pmt.OtherExpense) - pmt.OtherIncome

	for _, sub_item := range pmt.Subs {
		if (sub_item.InvoiceNo != "") {
			count_item = count_item + 1
		}
	}

	sqlexist := `select count(docno) as check_exist from dbo.bcapinvoice where docno = ?`
	err := db.Get(&check_exist, sqlexist, pmt.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	switch {
	case pmt.DocNo == "":
		return errors.New("Docno is null")
	case pmt.ApCode == "":
		return errors.New("Apcode is null")
	case pmt.DocDate == "":
		return errors.New("Docdate is null")
	case pmt.IsCancel == 1:
		return errors.New("Docno is cancel")
	case pmt.IsConfirm == 1:
		return errors.New("Docno is confirm")
	case count_item == 0:
		return errors.New("Docno is not have item")
	case (pmt.SumCashAmount == 0 && pmt.SumChqAmount == 0 && pmt.SumBankAmount == 0 && pmt.SumOfDeposit1 == 0):
		return errors.New("Docno not set money to another type payment")
	case sum_pay_amount > pmt.TotalAmount:
		return errors.New("Rec money is over totalamount")
	case pmt.SumChqAmount != 0 && len(pmt.Chqs) == 0:
		return errors.New("Docno not have chq use data")
	case pmt.SumOfDeposit1 != 0 && len(pmt.Deps) == 0:
		return errors.New("Docno not have deposit use data")
	case pmt.SumBankAmount != 0 && pmt.BankRefNo == "":
		return errors.New("Bank transfer not have accountcode data not complete")
	}

	var sumDepNetAmount float64

	for _, dep := range pmt.Deps {
		sumDepNetAmount = sumDepNetAmount + dep.Amount
	}

	if pmt.SumOfDeposit1 != sumDepNetAmount {
		return errors.New("Docno not have deposit not equa receive data")
	}

	if pmt.ExchangeRate == 0 {
		pmt.ExchangeRate = def.ExchangeRateDefault
	}

	if pmt.SaveFrom == 0 {
		pmt.SaveFrom = def.ApInvoiceSaveFrom
	}

	if (((pmt.TotalAmount + pmt.OtherIncome - (pmt.SumCashAmount + pmt.SumChqAmount + pmt.SumBankAmount + pmt.OtherExpense))) != 0) {
		return errors.New("Docno have money remain")
	}

	if (pmt.BookCode == "") {
		pmt.BookCode = def.ArInvoiceBookCode
	}
	if (pmt.Source == 0) {
		pmt.Source = def.ArInvoiceSource
	}

	pmt.GLFormat = def.ApInvoiceCreditGLFormat

	if (pmt.TaxNo == "") {
		pmt.TaxNo = pmt.DocNo
	}

	if (pmt.SumBankAmount != 0 && pmt.TransBankDate == "") {
		pmt.TransBankDate = pmt.DocDate
	}

	if (pmt.SumOfDeposit1 != 0) {
		pmt.SumOfDeposit2 = pmt.SumOfDeposit1 //+ m.ToFixed(inv.SumOfDeposit1-((inv.SumOfDeposit1*100)/(100+float64(inv.TaxRate))), 2)
	}

	fmt.Println("UserCode = ", pmt.UserCode)

	if (pmt.IsCompleteSave == 0) {
		pmt.IsCompleteSave = 1
	}

	if (check_exist == 0) {
		//Insert//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		pmt.CreatorCode = pmt.UserCode
		pmt.NetAmount = pmt.TotalAmount

		sql := `Insert into dbo.BCPayment(DocNo,TaxNo,DocDate,ApCode,CreatorCode,CreateDateTime,DepartCode,SumOfInvoice,SumOfDebitNote,SumOfCreditNote,BeforeTaxAmount,TaxAmount,TotalAmount,DiscountAmount,NetAmount,PettyCashAmount,SumCashAmount,SumChqAmount,SumBankAmount,SumOfWTax,GLFormat,IsCancel,MyDescription,AllocateCode,ProjectCode,BillGroup,SumHomeAmount1,SumHomeAmount2,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,CurrencyCode,ExchangeRate,IsCompleteSave,IsConfirm,RecurName,SumOfDeposit1,SumOfDeposit2,HomeAmountInv,HomeAmountDebt,HomeAmountCrd) values(?,?,?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sql, pmt.DocNo, pmt.TaxNo, pmt.DocDate, pmt.ApCode, pmt.CreatorCode, pmt.DepartCode, pmt.SumOfInvoice, pmt.SumOfDebitNote, pmt.SumOfCreditNote, pmt.BeforeTaxAmount, pmt.TaxAmount, pmt.TotalAmount, pmt.DiscountAmount, pmt.NetAmount, pmt.PettyCashAmount, pmt.SumCashAmount, pmt.SumChqAmount, pmt.SumBankAmount, pmt.SumOfWTax, pmt.GLFormat, pmt.IsCancel, pmt.MyDescription, pmt.AllocateCode, pmt.ProjectCode, pmt.BillGroup, pmt.SumHomeAmount1, pmt.SumHomeAmount2, pmt.OtherIncome, pmt.OtherExpense, pmt.ExcessAmount1, pmt.ExcessAmount2, pmt.CurrencyCode, pmt.ExchangeRate, pmt.IsCompleteSave, pmt.IsConfirm, pmt.RecurName, pmt.SumOfDeposit1, pmt.SumOfDeposit2, pmt.HomeAmountInv, pmt.HomeAmountDebt, pmt.HomeAmountCrd)
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
		pmt.LastEditorCode = pmt.UserCode
		pmt.NetAmount = pmt.TotalAmount

		sql := `Update dbo.BCReceipt1 set TaxNo=?,DocDate=?,LastEditorCode=?,LastEditDateT=getdate(),ArCode=?,SaleCode=?,DepartCode=?,SumOfInvoice=?,SumOfDebitNote=?,SumOfCreditNote=?,BeforeTaxAmount=?,TaxAmount=?,TotalAmount=?,DiscountAmount=?,NetAmount=?,SumCashAmount=?,SumChqAmount=?,SumCreditAmount=?,ChargeAmount=?,ChangeAmount=?,SumBankAmount=?,SumOfWTax=?,GLFormat=?,IsCancel=?,MyDescription=?,AllocateCode=?,ProjectCode=?,BillGroup=?,IsCompleteSave=?,SumHomeAmount1=?,SumHomeAmount2=?,DebtLimitReturn=?,CurrencyCode=?,ExchangeRate=?,OtherIncome=?,OtherExpense=?,ExcessAmount1=?,ExcessAmount2=?,IsConfirm=?,RecurName=? where docno = ?`
		_, err := db.Exec(sql, rcp.TaxNo, rcp.DocDate, rcp.LastEditorCode, rcp.ArCode, rcp.SaleCode, rcp.DepartCode, rcp.SumOfInvoice, rcp.SumOfDebitNote, rcp.SumOfCreditNote, rcp.BeforeTaxAmount, rcp.TaxAmount, rcp.TotalAmount, rcp.DiscountAmount, rcp.NetAmount, rcp.SumCashAmount, rcp.SumChqAmount, rcp.SumCreditAmount, rcp.ChargeAmount, rcp.ChangeAmount, rcp.SumBankAmount, rcp.SumOfWTax, rcp.GLFormat, rcp.IsCancel, rcp.MyDescription, rcp.AllocateCode, rcp.ProjectCode, rcp.BillGroup, rcp.IsCompleteSave, rcp.SumHomeAmount1, rcp.SumHomeAmount2, rcp.DebtLimitReturn, rcp.CurrencyCode, rcp.ExchangeRate, rcp.OtherIncome, rcp.OtherExpense, rcp.ExcessAmount1, rcp.ExcessAmount2, rcp.IsConfirm, rcp.RecurName, rcp.DocNo)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}
	}

	sql_del_sub := `delete dbo.bcapinvoicesub where docno = ?`
	_, err = db.Exec(sql_del_sub, pmt.DocNo)
	if err != nil {
		return err
	}

	var vLineNumber int

	for _, item := range pmt.Subs {
		fmt.Println("ItemSub")

		sqlsub := ` insert into dbo.BCApInvoiceSub(MyType,DocNo,TaxNo,TaxType,ItemCode,DocDate,ApCode,DepartCode,MyDescription,ItemName,WHCode,ShelfCode,CNQty,GRRemainQty,Qty,Price,DiscountWord,DiscountAmount,Amount,NetAmount,HomeAmount,BalanceAmount,SumOfExpCost,UnitCode,PORefNo,IRRefNo,IsCancel,LineNumber,BarCode,PORefLinenum,AVERAGECOST,LotNumber,SumOfCost,TaxRate,PackingRate1,PackingRate2) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		_, err = db.Exec(sqlsub, item.MyType, apv.DocNo, apv.TaxNo, apv.TaxType, item.ItemCode, apv.DocDate, apv.ApCode, apv.DepartCode, item.MyDescription, item.ItemName, item.WHCode, item.ShelfCode, item.CNQty, item.GRRemainQty, item.Qty, item.Price, item.DiscountWord, item.DiscountAmount, item.Amount, item.NetAmount, item.HomeAmount, item.BalanceAmount, item.SumOfExpCost, item.UnitCode, item.PORefNo, item.IRRefNo, item.IsCancel, item.LineNumber, item.BarCode, item.PORefLinenum, item.AVERAGECOST, item.LotNumber, item.SumOfCost, apv.TaxRate, item.PackingRate1, item.PackingRate2)
		fmt.Println("sqlsub = ", sqlsub)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		vLineNumber = vLineNumber + 1
	}

	sqldel := `delete dbo.BCInputTax where docno = ?`
	_, err = db.Exec(sqldel, pmt.DocNo)
	if err != nil {
		return err
	}

	sqltax := `insert into dbo.BCInputTax(SaveFrom,DocNo,BookCode,Source,DocDate,TaxDate,TaxNo,ApCode,ShortTaxDesc,TaxRate,Process,BeforeTaxAmount,TaxAmount,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?,?,'ซื้อสินค้า',?,1,?,?,?,getdate())`
	_, err = db.Exec(sqltax, pmt.SaveFrom, pmt.DocNo, pmt.BookCode, pmt.Source, pmt.DocDate, pmt.DocDate, pmt.TaxNo, pmt.ApCode, pmt.TaxRate, pmt.BeforeTaxAmount, pmt.TaxAmount, pmt.CreatorCode)
	fmt.Println("sqltax = ", sqltax)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	sqlrecdel := `delete dbo.BCPayMoney where docno = ?`
	_, err = db.Exec(sqlrecdel, pmt.DocNo)
	if err != nil {
		return err
	}

	var my_description_recmoney string

	my_description_recmoney = "ขายเงินสด"

	fmt.Println("RecMoney")
	var linenumber int

	if (pmt.SumCashAmount != 0) { //subs.PaymentType == 0:
		fmt.Println("SumCashAmount")
		sqlrec := `insert	into dbo.BCPayMoney(DocNo,DocDate,ApCode,ExchangeRate,PayAmount,PaymentType,SaveFrom,LineNumber,ProjectCode,DepartCode,MyDescription) values(?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlrec, pmt.DocNo, pmt.DocDate, pmt.ApCode, pmt.ExchangeRate, pmt.SumCashAmount, 0, pmt.SaveFrom, linenumber, pmt.ProjectCode, pmt.DepartCode, my_description_recmoney)
		if err != nil {
			return err
		}
	}
	//case dp.SumCreditAmount != 0: //subs.PaymentType == 1:
	if (pmt.SumChqAmount != 0) {
		fmt.Println("SumChqAmount")
		if (pmt.SumCashAmount != 0) {
			linenumber = 1
		} else {
			linenumber = 0
		}
		sqlrec := `insert	into dbo.BCPayMoney(DocNo,DocDate,ApCode,ExchangeRate,PayAmount,PaymentType,SaveFrom,LineNumber,RefNo,BankCode,ProjectCode,DepartCode,BankBranchCode,MyDescription,RefDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlrec, pmt.DocNo, pmt.DocDate, pmt.ApCode, pmt.ExchangeRate, pmt.SumChqAmount, 2, pmt.SaveFrom, linenumber, pmt.CreditRefNo, pmt.BankCode, pmt.ProjectCode, pmt.DepartCode, pmt.BankBranchCode, my_description_recmoney, pmt.RefDate)
		if err != nil {
			return err
		}
	}

	//case dp.SumChqAmount != 0: //subs.PaymentType == 2:
	if (pmt.SumBankAmount != 0) {
		fmt.Println("SumChqAmount")
		if (pmt.SumCashAmount != 0 && pmt.SumChqAmount != 0) {
			linenumber = 2
		} else if ((pmt.SumCashAmount != 0 && pmt.SumChqAmount == 0)) {
			linenumber = 1
		} else if ((pmt.SumCashAmount == 0 && pmt.SumChqAmount != 0)) {
			linenumber = 1
		} else if ((pmt.SumCashAmount == 0 && pmt.SumChqAmount == 0)) {
			linenumber = 0
		}

		sqlrec := `insert	into dbo.BCPayMoney(DocNo,DocDate,ApCode,ExchangeRate,PayAmount,PaymentType,SaveFrom,LineNumber,RefNo,ProjectCode,DepartCode,MyDescription,RefDate,TransBankDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlrec, pmt.DocNo, pmt.DocDate, pmt.ApCode, pmt.ExchangeRate, pmt.SumBankAmount, 3, pmt.SaveFrom, linenumber, pmt.BankRefNo, pmt.ProjectCode, pmt.DepartCode, my_description_recmoney, pmt.DocDate, pmt.TransBankDate)
		if err != nil {
			return err
		}
	}

	var depLineNumber int

	if (pmt.SumOfDeposit1 != 0) {
		sqldepdel := `delete dbo.BCApDepositUse where docno = ?`
		_, err = db.Exec(sqldepdel, pmt.DocNo)
		if err != nil {
			return err
		}

		for _, d := range pmt.Deps {
			var depNetAmount float64
			d.LineNumber = depLineNumber

			depNetAmount = d.Amount - (m.ToFixed(d.Amount-((d.Amount*100)/(100+float64(def.TaxRateDefault))), 2))

			sqldep := `insert into dbo.bcapdeposituse(DocNo,DepositNo,DocDate,MyDescription,Balance,Amount,DeposTaxType,NetAmount,LineNumber) values(?,?,?,?,?,?,?,?,?)`
			_, err = db.Exec(sqldep, pmt.DocNo, d.DepositNo, pmt.DocDate, my_description_recmoney, d.Balance, d.Amount, 0, depNetAmount, d.LineNumber)
			if err != nil {
				return err
			}
			depLineNumber = depLineNumber + 1
		}
	}

	if (pmt.SumChqAmount != 0) {
		sqlchqdel := `delete dbo.BCChqOut where docno = ?`
		_, err = db.Exec(sqlchqdel, pmt.DocNo)
		if err != nil {
			return err
		}

		for _, c := range pmt.Chqs {
			if ((c.ReceiveDate == "") || (c.DueDate == "")) {
				c.ReceiveDate = pmt.DocDate
				c.DueDate = pmt.DocDate
			}

			sqldep := `insert into dbo.bcchqout(BankCode,ChqNumber,DocNo,ApCode,DueDate,BookNo,Status,SaveFrom,StatusDate,StatusDocNo,DepartCode,BankBranchCode,Amount,Balance,MyDescription,ExchangeRate,RefChqRowOrder) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
			_, err = db.Exec(sqldep, c.BankCode, c.ChqNumber, pmt.DocNo, pmt.ApCode, c.DueDate, c.BookNo, c.Status, pmt.SaveFrom, c.StatusDate, c.StatusDocNo, pmt.DepartCode, c.BankBranchCode, c.Amount, c.Balance, my_description_recmoney, pmt.ExchangeRate, c.RefChqRowOrder)
			if err != nil {
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
