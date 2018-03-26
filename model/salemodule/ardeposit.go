package model

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	m "github.com/loukmho/bcaccount_api/model"
	"time"
)

type ArDeposit struct {
	RowOrder        int     `json:"row_order" db:"RowOrder"`
	DocNo           string  `json:"doc_no" db:"DocNo"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	TaxType         int     `json:"tax_type" db:"TaxType"`
	OutPutTax
	Customer
	SaleMan
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	CreditDay       int     `json:"credit_day" db:"CreditDay"`
	DueDate         string  `json:"due_date" db:"DueDate"`
	TaxRate         float64 `json:"tax_rate" db:"TaxRate"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	BeforeTaxAmount float64 `json:"before_tax_amount" db:"BeforeTaxAmount"`
	TaxAmount       float64 `json:"tax_amount" db:"TaxAmount"`
	TotalAmount     float64 `json:"total_amount" db:"TotalAmount"`
	SumOfWTax       float64 `json:"sum_of_w_tax" db:"SumOfWTax"`
	NetAmount       float64 `json:"net_amount" db:"NetAmount"`
	BillBalance     float64 `json:"bill_balance" db:"BillBalance"`
	OtherIncome     float64 `json:"other_income" db:"OtherIncome"`
	OtherExpense    float64 `json:"other_expense" db:"OtherExpense"`
	ExcessAmount1   float64 `json:"excess_amount_1" db:"ExcessAmount1"`
	ExcessAmount2   float64 `json:"excess_amount_2" db:"ExcessAmount2"`
	ChargeAmount    float64 `json:"charge_amount" db:"ChargeAmount"`
	ChangeAmount    float64 `json:"change_amount" db:"ChangeAmount"`
	RefNo           string  `json:"ref_no" db:"RefNo"`
	CurrencyCode    string  `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate    float64 `json:"exchange_rate" db:"ExchangeRate"`
	SumCashAmount   float64 `json:"sum_cash_amount" db:"SumCashAmount"`
	SumChqAmount    float64 `json:"sum_chq_amount" db:"SumChqAmount"`
	SumCreditAmount float64 `json:"sum_credit_amount" db:"SumCreditAmount"`
	SumBankAmount   float64 `json:"sum_bank_amount" db:"SumBankAmount"`
	GLFormat        string  `json:"gl_format" db:"GLFormat"`
	GLStartPosting  int     `json:"gl_start_posting" db:"GLStartPosting"`
	IsPostGL        int     `json:"is_post_gl" db:"IsPostGL"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	IsReturnMoney   int     `json:"is_return_money" db:"IsReturnMoney"`
	AllocateCode    string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string  `json:"project_code" db:"ProjectCode"`
	BillGroup       string  `json:"bill_group" db:"BillGroup"`
	RecurName       string  `json:"recur_name" db:"RecurName"`
	ConfirmCode     string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string  `json:"cancel_date_time" db:"CancelCode"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	UserCode        string  `json:"user_code" db:"UserCode"`
	ListArDepRecMoney
}

type Customer struct {
	ArCode string `json:"ar_code" db:"ArCode"`
	ArName string `json:"ar_name" db:"ArName"`
}

type OutPutTax struct {
	TaxNo    string `json:"tax_no" db:"TaxNo"`
	TaxDate  string `json:"tax_date" db:"TaxDate"`
	BookCode string `json:"book_code" db:"BookCode"`
}

type SaleMan struct {
	SaleCode string `json:"sale_code" db:"SaleCode"`
	SaleName string `json:"sale_name" db:"SaleName"`
}

type ListArDepRecMoney struct {
	CreditType     string `json:"credit_type" db:"CreditType"`
	ConfirmNo      string `json:"confirm_no" db:"ConfirmNo"`
	CreditRefNo    string `json:"credit_ref_no" db:"CreditRefNo"`
	BankCode       string `json:"bank_code" db:"BookCode"`
	BankBranchCode string `json:"bank_branch_code" db:"BankBranchCode"`
	TransBankDate  string `json:"trans_bank_date" db:"TransBankDate"`
	RefDate        string `json:"ref_date" db:"RefDate"`
}

//type ListArDepRecMoney struct {
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

func (dp *ArDeposit) SaveArDeposit(db *sqlx.DB) error {
	var check_exist int
	var sum_pay_amount float64
	var source int

	now := time.Now()

	sqlexist := `select count(docno) as check_exist from dbo.bcardeposit where docno = ? and arcode = ?`//เช็คว่ามีเอกสารหรือยัง
	err := db.Get(&check_exist, sqlexist, dp.DocNo, dp.ArCode)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	switch {
	case dp.DocNo == "":
		return errors.New("docno is null")
	case dp.ArCode == "":
		return errors.New("arcode is null")
	case dp.TotalAmount == 0:
		return errors.New("totalamount is 0")
	case dp.DocDate == "":
		return errors.New("docdate is null")
	case dp.SumCashAmount == 0 && dp.SumCreditAmount == 0 && dp.SumChqAmount == 0 && dp.SumBankAmount == 0:
		return errors.New("docno not set money to another type payment")
	case sum_pay_amount > dp.TotalAmount:
		return errors.New("rec money is over totalamount")
	case dp.IsCancel == 1:
		return errors.New("docno is cancel can not edit data")
	case dp.IsConfirm == 1:
		return errors.New("docno is confirm can not edit data")
	case dp.SumCreditAmount != 0 && (dp.CreditType == "" || dp.ConfirmNo == "" || dp.CreditRefNo == ""):
		return errors.New("creditcard not have credittype or confirmno or creditrefno")
	}

	if (dp.TaxNo == "") {
		dp.TaxNo = dp.DocNo
	}
	if (dp.TaxDate == "") {
		dp.TaxDate = dp.DocDate
	}
	if (dp.DueDate == "" && dp.CreditDay == 0) {
		dp.DueDate = dp.DocDate
	}else{
		dp.DueDate = now.AddDate(0, 0, dp.CreditDay).Format("2006-01-02")
	}
	if (dp.ExchangeRate == 0) {
		dp.ExchangeRate = 1
	}

	fmt.Println("UserCode = ",dp.UserCode)

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")

	//fmt.Println("TaxRate = ", def.TaxRateDefault)

	if dp.TaxRate == 0 {
		dp.TaxRate = def.TaxRateDefault
	}
	if dp.ExchangeRate == 0 {
		dp.ExchangeRate = def.ExchangeRateDefault
	}

	if (dp.BookCode==""){
		dp.BookCode = def.ArDepositBookCode
	}
	source = def.ArDepositSource
	if (dp.GLFormat==""){
		dp.GLFormat = def.ArDepositGLFormat
	}

	dp.BeforeTaxAmount, dp.TaxAmount = m.CalcTaxDoc(dp.TaxType, dp.TaxRate, dp.TotalAmount)

	//fmt.Println("BeforeTax,TaxAmount", dp.BeforeTaxAmount, dp.TaxAmount)

	fmt.Println("check_exist = ", check_exist)

	if (check_exist == 0) {
		//Insert//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		dp.CreatorCode = dp.UserCode
		dp.BillBalance = dp.TotalAmount
		dp.NetAmount = dp.TotalAmount

		sql := `Insert into dbo.BCArDeposit(DocNo,DocDate,TaxDate,TaxType,TaxNo,ArCode,DepartCode,CreditDay,DueDate,SaleCode,TaxRate,MyDescription,BeforeTaxAmount,TaxAmount,TotalAmount,SumOfWTax,NetAmount,BillBalance,OtherIncome,OtherExpense, ExcessAmount1,ExcessAmount2,ChargeAmount,ChangeAmount,RefNo,CurrencyCode,ExchangeRate,SumCashAmount,SumChqAmount,SumCreditAmount, SumBankAmount,GLFormat,AllocateCode,ProjectCode,BillGroup,RecurName,CreatorCode,CreateDateTime)
			values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,getdate())`
		_, err = db.Exec(sql, dp.DocNo, dp.DocDate, dp.TaxDate, dp.TaxType, dp.TaxNo, dp.ArCode, dp.DepartCode, dp.CreditDay, dp.DueDate, dp.SaleCode, dp.TaxRate, dp.MyDescription, dp.BeforeTaxAmount, dp.TaxAmount, dp.TotalAmount, dp.SumOfWTax, dp.NetAmount, dp.BillBalance, dp.OtherIncome, dp.OtherExpense, dp.ExcessAmount1, dp.ExcessAmount2, dp.ChargeAmount, dp.ChangeAmount, dp.RefNo, dp.CurrencyCode, dp.ExchangeRate, dp.SumCashAmount, dp.SumChqAmount, dp.SumCreditAmount, dp.SumBankAmount, dp.GLFormat, dp.AllocateCode, dp.ProjectCode, dp.BillGroup, dp.RecurName, dp.CreatorCode)
		fmt.Println("sql =", sql, dp.DocNo, dp.DocDate, dp.DueDate, dp.TaxDate, dp.TaxNo, dp.ArCode, dp.SaleCode)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

	} else {

		//Update/////////////////////////////////////////////////////////////////////////////////////////////////////////
		dp.LastEditorCode = dp.UserCode
		dp.BillBalance = dp.TotalAmount
		dp.NetAmount = dp.TotalAmount

		sql := `Update dbo.BCArDeposit set DocDate=?,TaxDate=?,TaxType=?,TaxNo=?,ArCode=?,DepartCode=?,CreditDay=?,DueDate=?,SaleCode=?,TaxRate=?,MyDescription=?,BeforeTaxAmount=?,TaxAmount=?,TotalAmount=?,SumOfWTax=?,NetAmount=?,BillBalance=?,OtherIncome=?,OtherExpense=?,ExcessAmount1=?,ExcessAmount2=?,ChargeAmount=?,ChangeAmount=?,RefNo=?,CurrencyCode=?,ExchangeRate=?,SumCashAmount=?,SumChqAmount=?,SumCreditAmount=?,SumBankAmount=?,GLFormat=?,AllocateCode=?,ProjectCode=?,BillGroup=?,RecurName=?,LastEditorCode=?,LastEditDateT=getdate() where docno = ?`
		_, err := db.Exec(sql, dp.DocDate, dp.TaxDate, dp.TaxType, dp.TaxNo, dp.ArCode, dp.DepartCode, dp.CreditDay, dp.DueDate, dp.SaleCode, dp.TaxRate, dp.MyDescription, dp.BeforeTaxAmount, dp.TaxAmount, dp.TotalAmount, dp.SumOfWTax, dp.NetAmount, dp.BillBalance, dp.OtherIncome, dp.OtherExpense, dp.ExcessAmount1, dp.ExcessAmount2, dp.ChargeAmount, dp.ChangeAmount, dp.RefNo, dp.CurrencyCode, dp.ExchangeRate, dp.SumCashAmount, dp.SumChqAmount, dp.SumCreditAmount, dp.SumBankAmount, dp.GLFormat, dp.AllocateCode, dp.ProjectCode, dp.BillGroup, dp.RecurName, dp.LastEditorCode, dp.DocNo)
		if err != nil {
			return err
		}
	}

	sqldel := `delete dbo.BCOutputTax where docno = ?`
	_, err = db.Exec(sqldel, dp.DocNo)
	if err != nil {
		return err
	}

	sqltax := "insert into dbo.BCOutputTax(SaveFrom,DocNo,BookCode,Source,DocDate,TaxDate,TaxNo,ArCode,ShortTaxDesc,TaxRate,Process,BeforeTaxAmount,TaxAmount,CreatorCode,CreateDateTime) values(1,?,?,?,?,?,?,?,'ขายสินค้า',?,1,?,?,?,getdate())"
	_, err = db.Exec(sqltax, dp.DocNo, dp.BookCode, source, dp.DocDate, dp.TaxDate, dp.TaxNo, dp.ArCode, dp.TaxRate, dp.BeforeTaxAmount, dp.TaxAmount, dp.CreatorCode)
	if err != nil {
		return err
	}

	sqlrecdel := `delete dbo.BCRecMoney where docno = ?`
	_, err = db.Exec(sqlrecdel, dp.DocNo)
	if err != nil {
		return err
	}

	var my_description_recmoney string

	if (dp.MyDescription == "") {
		my_description_recmoney = "รับเงินมัดจำ"
	} else {
		my_description_recmoney = dp.MyDescription
	}

	fmt.Println("RecMoney")
	var linenumber int

	if (dp.SumCashAmount != 0) { //subs.PaymentType == 0:
		fmt.Println("SumCashAmount")
		sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,PaymentType,LineNumber,ProjectCode,DepartCode,SaleCode,MyDescription) values(?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlrec, dp.DocNo, dp.DocDate, dp.ArCode, dp.ExchangeRate, dp.SumCashAmount, 0, linenumber, dp.ProjectCode, dp.DepartCode, dp.SaleCode, my_description_recmoney)
		if err != nil {
			return err
		}
	}
	//case dp.SumCreditAmount != 0: //subs.PaymentType == 1:
	if (dp.SumCreditAmount != 0) {
		fmt.Println("SumCreditAmount")
		if (dp.SumCashAmount != 0) {
			linenumber = 1
		} else {
			linenumber = 0
		}
		sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,ChqTotalAmount,PaymentType,CreditType,ConfirmNo,LineNumber,RefNo,BankCode,ProjectCode,DepartCode,SaleCode,MyDescription,RefDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlrec, dp.DocNo, dp.DocDate, dp.ArCode, dp.ExchangeRate, dp.SumCreditAmount, dp.SumCreditAmount, 1, dp.CreditType, dp.ConfirmNo, linenumber, dp.CreditRefNo, dp.BankCode, dp.ProjectCode, dp.DepartCode, dp.SaleCode, my_description_recmoney, dp.DocDate)
		if err != nil {
			return err
		}
	}

	//case dp.SumChqAmount != 0: //subs.PaymentType == 2:
	if (dp.SumChqAmount != 0) {
		fmt.Println("SumChqAmount")
		if (dp.SumCashAmount != 0 && dp.SumCreditAmount != 0) {
			linenumber = 2
		} else if ((dp.SumCashAmount != 0 && dp.SumCreditAmount == 0)) {
			linenumber = 1
		} else if ((dp.SumCashAmount == 0 && dp.SumCreditAmount != 0)) {
			linenumber = 1
		} else if ((dp.SumCashAmount == 0 && dp.SumCreditAmount == 0)) {
			linenumber = 0
		}

		sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,PaymentType,LineNumber,RefNo,BankCode,ProjectCode,DepartCode,SaleCode,BankBranchCode,TransBankDate,MyDescription,RefDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlrec, dp.DocNo, dp.DocDate, dp.ArCode, dp.ExchangeRate, dp.SumChqAmount, 2, linenumber, dp.CreditRefNo, dp.BankCode, dp.ProjectCode, dp.DepartCode, dp.SaleCode, dp.BankBranchCode, dp.TransBankDate, my_description_recmoney, dp.RefDate)
		if err != nil {
			return err
		}
	}

	//case dp.SumBankAmount != 0: //subs.PaymentType == 3:
	if (dp.SumBankAmount != 0) {
		fmt.Println("SumBankAmount")
		if (dp.SumCashAmount != 0 && dp.SumCreditAmount != 0 && dp.SumChqAmount != 0) {
			linenumber = 3
		} else if (dp.SumCashAmount != 0 && dp.SumCreditAmount == 0 && dp.SumChqAmount != 0) {
			linenumber = 2
		} else if (dp.SumCashAmount == 0 && dp.SumCreditAmount != 0 && dp.SumChqAmount != 0) {
			linenumber = 2
		} else if (dp.SumCashAmount != 0 && dp.SumCreditAmount != 0 && dp.SumChqAmount == 0) {
			linenumber = 2
		} else if (dp.SumCashAmount != 0 && dp.SumCreditAmount != 0 && dp.SumChqAmount == 0) {
			linenumber = 2
		} else if (dp.SumCashAmount != 0 && dp.SumCreditAmount == 0 && dp.SumChqAmount == 0) {
			linenumber = 1
		} else if (dp.SumCashAmount == 0 && dp.SumCreditAmount != 0 && dp.SumChqAmount == 0) {
			linenumber = 1
		} else if (dp.SumCashAmount == 0 && dp.SumCreditAmount == 0 && dp.SumChqAmount != 0) {
			linenumber = 1
		} else if (dp.SumCashAmount == 0 && dp.SumCreditAmount == 0 && dp.SumChqAmount == 0) {
			linenumber = 0
		}

		sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,PaymentType,LineNumber,RefNo,ProjectCode,DepartCode,SaleCode,MyDescription,RefDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlrec, dp.DocNo, dp.DocDate, dp.ArCode, dp.ExchangeRate, dp.SumBankAmount, 3, linenumber, dp.CreditRefNo, dp.ProjectCode, dp.DepartCode, dp.SaleCode, my_description_recmoney, dp.DocDate)
		if err != nil {
			return err
		}
	}

	return nil
}

//func (dp *ArDeposit) UpdateArDeposit(db *sqlx.DB) error {
//	var check_exist int
//	var sum_pay_amount float64
//
//	switch {
//	case dp.DocNo == "":
//		return errors.New("docno is null")
//	case dp.ArCode == "":
//		return errors.New("arcode is null")
//	case dp.TotalAmount == 0:
//		return errors.New("totalamount is 0")
//	case dp.DocDate == "":
//		return errors.New("docdate is null")
//	case dp.TaxNo == "":
//		return errors.New("taxno is null")
//	case dp.TaxDate == "":
//		return errors.New("taxdate is null")
//	case dp.SumCashAmount == 0 && dp.SumCreditAmount == 0 && dp.SumChqAmount == 0 && dp.SumBankAmount == 0:
//		return errors.New("docno not set money to another type payment")
//	case sum_pay_amount > dp.TotalAmount:
//		return errors.New("rec money is over totalamount")
//	}
//
//	sqlexist := `select count(docno) as check_exist from dbo.bcardeposit where docno = ? and arcode = ?`
//	err := db.Get(&check_exist, sqlexist, dp.DocNo, dp.ArCode)
//	if err != nil {
//		fmt.Println(err.Error())
//		return nil
//	}
//
//	fmt.Println("check_exist = ", check_exist)
//	if (check_exist != 0) {
//		sql := `Update dbo.BCArDeposit set DocDate=?,TaxDate=?,TaxType=?,TaxNo=?,ArCode=?,DepartCode=?,CreditDay=?,DueDate=?,SaleCode=?,TaxRate=?,MyDescription=?,BeforeTaxAmount=?,TaxAmount=?,TotalAmount=?,SumOfWTax=?,NetAmount=?,BillBalance=?,OtherIncome=?,OtherExpense=?,ExcessAmount1=?,ExcessAmount2=?,ChargeAmount=?,ChangeAmount=?,RefNo=?,CurrencyCode=?,ExchangeRate=?,SumCashAmount=?,SumChqAmount=?,SumCreditAmount=?,SumBankAmount=?,GLFormat=?,AllocateCode=?,ProjectCode=?,BillGroup=?,RecurName=?,LastEditorCode=?,LastEditDateT=getdate() where docno = ?`
//		_, err := db.Exec(sql, dp.DocDate, dp.TaxDate, dp.TaxType, dp.TaxNo, dp.ArCode, dp.DepartCode, dp.CreditDay, dp.DueDate, dp.SaleCode, dp.TaxRate, dp.MyDescription, dp.BeforeTaxAmount, dp.TaxAmount, dp.TotalAmount, dp.SumOfWTax, dp.NetAmount, dp.BillBalance, dp.OtherIncome, dp.OtherExpense, dp.ExcessAmount1, dp.ExcessAmount2, dp.ChargeAmount, dp.ChangeAmount, dp.RefNo, dp.CurrencyCode, dp.ExchangeRate, dp.SumCashAmount, dp.SumChqAmount, dp.SumCreditAmount, dp.SumBankAmount, dp.GLFormat, dp.AllocateCode, dp.ProjectCode, dp.BillGroup, dp.RecurName, dp.LastEditorCode, dp.DocNo)
//		if err != nil {
//			return err
//		}
//		sqldel := `delete dbo.BCOutputTax where docno = ?`
//		_, err = db.Exec(sqldel, dp.DocNo)
//		if err != nil {
//			return err
//		}
//
//		sqltax := "insert into dbo.BCOutputTax(SaveFrom,DocNo,BookCode,Source,DocDate,TaxDate,TaxNo,ArCode,ShortTaxDesc,TaxRate,Process,BeforeTaxAmount,TaxAmount,CreatorCode,CreateDateTime) values(1,?,?,6,?,?,?,?,'ขายสินค้า',?,1,?,?,?,getdate())"
//		_, err = db.Exec(sqltax, dp.DocNo, dp.BookCode, dp.DocDate, dp.TaxDate, dp.TaxNo, dp.ArCode, dp.TaxRate, dp.BeforeTaxAmount, dp.TaxAmount, dp.CreatorCode)
//		if err != nil {
//			return err
//		}
//
//		sqlrecdel := `delete dbo.BCRecMoney where docno = ?`
//		_, err = db.Exec(sqlrecdel, dp.DocNo)
//		if err != nil {
//			return err
//		}
//
//		//var my_description_recmoney string
//
//		//if (dp.MyDescription == "") {
//		//	my_description_recmoney = "รับเงินมัดจำ"
//		//} else {
//		//	my_description_recmoney = dp.MyDescription
//		//}
//
//		//for _, subs := range dp.RecMoney {
//		//	switch {
//		//	case subs.PaymentType == 0:
//		//		sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,PaymentType,LineNumber,ProjectCode,DepartCode,SaleCode,MyDescription) values(?,?,?,?,?,?,?,?,?,?,?)`
//		//		_, err = db.Exec(sqlrec, dp.DocNo, dp.DocDate, dp.ArCode, dp.ExchangeRate, subs.PayAmount, subs.PaymentType, subs.LineNumber, dp.ProjectCode, dp.DepartCode, dp.SaleCode, my_description_recmoney)
//		//		if err != nil {
//		//			return err
//		//		}
//		//	case subs.PaymentType == 1:
//		//		sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,ChqTotalAmount,PaymentType,CreditType,ConfirmNo,LineNumber,RefNo,BankCode,ProjectCode,DepartCode,SaleCode,MyDescription,RefDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
//		//		_, err = db.Exec(sqlrec, dp.DocNo, dp.DocDate, dp.ArCode, dp.ExchangeRate, subs.PayAmount, subs.ChqTotalAmount, subs.PaymentType, subs.CreditType, subs.ConfirmNo, subs.LineNumber, subs.RefNo, subs.BankCode, dp.ProjectCode, dp.DepartCode, dp.SaleCode, my_description_recmoney, subs.RefDate)
//		//		if err != nil {
//		//			return err
//		//		}
//		//
//		//	case subs.PaymentType == 2:
//		//		sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,ChqTotalAmount,PaymentType,CreditType,ConfirmNo,LineNumber,RefNo,BankCode,ProjectCode,DepartCode,SaleCode,BankBranchCode,TransBankDate,MyDescription,RefDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
//		//		_, err = db.Exec(sqlrec, dp.DocNo, dp.DocDate, dp.ArCode, dp.ExchangeRate, subs.PayAmount, subs.ChqTotalAmount, subs.PaymentType, subs.CreditType, subs.ConfirmNo, subs.LineNumber, subs.RefNo, subs.BankCode, dp.ProjectCode, dp.DepartCode, dp.SaleCode, subs.BankBranchCode, subs.TransBankDate, my_description_recmoney, subs.RefDate)
//		//		if err != nil {
//		//			return err
//		//		}
//		//
//		//	case subs.PaymentType == 3:
//		//		sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,PaymentType,LineNumber,RefNo,ProjectCode,DepartCode,SaleCode,MyDescription,RefDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?)`
//		//		_, err = db.Exec(sqlrec, dp.DocNo, dp.DocDate, dp.ArCode, dp.ExchangeRate, subs.PayAmount, subs.PaymentType, subs.LineNumber, subs.RefNo, dp.ProjectCode, dp.DepartCode, dp.SaleCode, my_description_recmoney, subs.RefDate)
//		//		if err != nil {
//		//			return err
//		//		}
//		//	}
//		//}
//	}
//	return nil
//}

func (dp *ArDeposit) SearchArDepositByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select a.RowOrder,a.DocNo as DocNo,a.DocDate,a.TaxDate,a.TaxType,a.TaxNo,a.ArCode,isnull(b.name1,'') as ArName,isnull(a.DepartCode,'') as DepartCode,a.CreditDay,a.DueDate,a.SaleCode,isnull(c.name,'') as SaleName,a.TaxRate,isnull(a.MyDescription,'') as MyDescription,a.BeforeTaxAmount,a.TaxAmount,a.TotalAmount,a.SumOfWTax,a.NetAmount,a.BillBalance,a.OtherIncome,a.OtherExpense,a.ExcessAmount1,a.ExcessAmount2,a.ChargeAmount,a.ChangeAmount,isnull(a.RefNo,'') as RefNo,isnull(a.CurrencyCode,'') as CurrencyCode,a.ExchangeRate,a.SumCashAmount,a.SumChqAmount,a.SumCreditAmount,a.SumBankAmount,isnull(a.GLFormat,'') as GLFormat,isnull(a.AllocateCode,'') as AllocateCode,isnull(a.ProjectCode,'') as ProjectCode,isnull(a.BillGroup,'') as BillGroup,isnull(a.RecurName,'') as RecurName,a.CreatorCode,a.CreateDateTime,isnull(d.bookcode,'') as BookCode from dbo.bcardeposit a WITH (NOLOCK) left join dbo.bcar b WITH (NOLOCK) on a.arcode = b.code left join dbo.bcsale c on a.salecode = c.code left join dbo.bcoutputtax d WITH (NOLOCK) on d.docno = a.docno and d.arcode = a.arcode where a.docno  = ?`
	err := db.Get(dp, sql, docno)
	fmt.Println("sql =", sql)
	if err != nil {
		return err
	}
	return nil
}

func (dp *ArDeposit) SearchArDepositByKeyword(db *sqlx.DB, keyword string) (dps []*ArDeposit, err error) {
	sql := `set dateformat dmy     select a.RowOrder,a.DocNo as DocNo,a.DocDate,a.TaxDate,a.TaxType,a.TaxNo,a.ArCode,isnull(b.name1,'') as ArName,isnull(a.DepartCode,'') as DepartCode,a.CreditDay,a.DueDate,a.SaleCode,isnull(c.name,'') as SaleName,a.TaxRate,isnull(a.MyDescription,'') as MyDescription,a.BeforeTaxAmount,a.TaxAmount,a.TotalAmount,a.SumOfWTax,a.NetAmount,a.BillBalance,a.OtherIncome,a.OtherExpense,a.ExcessAmount1,a.ExcessAmount2,a.ChargeAmount,a.ChangeAmount,isnull(a.RefNo,'') as RefNo,isnull(a.CurrencyCode,'') as CurrencyCode,a.ExchangeRate,a.SumCashAmount,a.SumChqAmount,a.SumCreditAmount,a.SumBankAmount,isnull(a.GLFormat,'') as GLFormat,isnull(a.AllocateCode,'') as AllocateCode,isnull(a.ProjectCode,'') as ProjectCode,isnull(a.BillGroup,'') as BillGroup,isnull(a.RecurName,'') as RecurName,a.CreatorCode,a.CreateDateTime,isnull(d.bookcode,'') as BookCode from dbo.bcardeposit a WITH (NOLOCK) left join dbo.bcar b WITH (NOLOCK) on a.arcode = b.code left join dbo.bcsale c on a.salecode = c.code left join dbo.bcoutputtax d WITH (NOLOCK) on d.docno = a.docno and d.arcode = a.arcode where (a.docno  like '%'+?+'%' or a.arcode like '%'+?+'%' or a.salecode like '%'+?+'%' or b.name1 like '%'+?+'%')  order by a.docno`
	err = db.Select(&dps, sql, keyword, keyword, keyword, keyword)
	fmt.Println("sql =", sql)
	if err != nil {
		return nil, err
	}

	return dps, nil
}
