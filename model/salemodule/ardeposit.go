package model

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	m "github.com/loukmho/bcaccount_api/model"
	"time"
)

type ArDeposit struct {
	RowOrder        int                  `json:"row_order" db:"RowOrder"`
	DocNo           string               `json:"doc_no" db:"DocNo"`
	DocDate         string               `json:"doc_date" db:"DocDate"`
	TaxType         int                  `json:"tax_type" db:"TaxType"`
	Source          int                  `json:"source" db:"Source"`
	SaveFrom        int                  `json:"save_from" db:"SaveFrom"`
	OutPutTax
	Customer
	SaleMan
	DepartCode      string               `json:"depart_code" db:"DepartCode"`
	CreditDay       int                  `json:"credit_day" db:"CreditDay"`
	DueDate         string               `json:"due_date" db:"DueDate"`
	TaxRate         float64              `json:"tax_rate" db:"TaxRate"`
	IsConfirm       int                  `json:"is_confirm" db:"IsConfirm"`
	MyDescription   string               `json:"my_description" db:"MyDescription"`
	BeforeTaxAmount float64              `json:"before_tax_amount" db:"BeforeTaxAmount"`
	TaxAmount       float64              `json:"tax_amount" db:"TaxAmount"`
	TotalAmount     float64              `json:"total_amount" db:"TotalAmount"`
	SumOfWTax       float64              `json:"sum_of_w_tax" db:"SumOfWTax"`
	NetAmount       float64              `json:"net_amount" db:"NetAmount"`
	BillBalance     float64              `json:"bill_balance" db:"BillBalance"`
	OtherIncome     float64              `json:"other_income" db:"OtherIncome"`
	OtherExpense    float64              `json:"other_expense" db:"OtherExpense"`
	ExcessAmount1   float64              `json:"excess_amount_1" db:"ExcessAmount1"`
	ExcessAmount2   float64              `json:"excess_amount_2" db:"ExcessAmount2"`
	ChargeAmount    float64              `json:"charge_amount" db:"ChargeAmount"`
	ChangeAmount    float64              `json:"change_amount" db:"ChangeAmount"`
	RefNo           string               `json:"ref_no" db:"RefNo"`
	CurrencyCode    string               `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate    float64              `json:"exchange_rate" db:"ExchangeRate"`
	SumCashAmount   float64              `json:"sum_cash_amount" db:"SumCashAmount"`
	SumChqAmount    float64              `json:"sum_chq_amount" db:"SumChqAmount"`
	SumCreditAmount float64              `json:"sum_credit_amount" db:"SumCreditAmount"`
	SumBankAmount   float64              `json:"sum_bank_amount" db:"SumBankAmount"`
	GLFormat        string               `json:"gl_format" db:"GLFormat"`
	GLStartPosting  int                  `json:"gl_start_posting" db:"GLStartPosting"`
	IsPostGL        int                  `json:"is_post_gl" db:"IsPostGL"`
	IsCancel        int                  `json:"is_cancel" db:"IsCancel"`
	IsReturnMoney   int                  `json:"is_return_money" db:"IsReturnMoney"`
	AllocateCode    string               `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string               `json:"project_code" db:"ProjectCode"`
	BillGroup       string               `json:"bill_group" db:"BillGroup"`
	RecurName       string               `json:"recur_name" db:"RecurName"`
	ConfirmCode     string               `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string               `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string               `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string               `json:"cancel_date_time" db:"CancelCode"`
	CreatorCode     string               `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string               `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string               `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string               `json:"last_edit_date_t" db:"LastEditDateT"`
	UserCode        string               `json:"user_code" db:"UserCode"`
	ListArDepRecMoney
	Cdcs            []*ListDepCreditCard `json:"cdcs"`
	Chqs            []*ListDepChqIn      `json:"chqs"`
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
	BankRefNo      string `json:"bank_ref_no" db:"BankRefNo"`
	TransBankDate  string `json:"trans_bank_date" db:"TransBankDate"`
	RefDate        string `json:"ref_date" db:"RefDate"`
}

type ListDepChqIn struct {
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

type ListDepCreditCard struct {
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

func (dp *ArDeposit) InsertAndEditArDeposit(db *sqlx.DB) error {
	var check_exist int
	var sum_pay_amount float64

	now := time.Now()

	sqlexist := `select count(docno) as check_exist from dbo.bcardeposit where docno = ?` //เช็คว่ามีเอกสารหรือยัง
	err := db.Get(&check_exist, sqlexist, dp.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil
	}

	sum_pay_amount = (dp.SumCashAmount+dp.SumCreditAmount+dp.SumChqAmount+dp.SumBankAmount+dp.OtherExpense)-dp.OtherIncome

	switch {
	case dp.DocNo == "":
		return errors.New("Docno is null")
	case dp.ArCode == "":
		return errors.New("Arcode is null")
	case dp.TotalAmount == 0:
		return errors.New("Totalamount is 0")
	case dp.DocDate == "":
		return errors.New("Docdate is null")
	case dp.SumCashAmount == 0 && dp.SumCreditAmount == 0 && dp.SumChqAmount == 0 && dp.SumBankAmount == 0:
		return errors.New("Docno not set money to another type payment")
	case sum_pay_amount > dp.TotalAmount:
		return errors.New("Rec money is over totalamount")
	case dp.IsCancel == 1:
		return errors.New("Docno is cancel can not edit data")
	case dp.IsConfirm == 1:
		return errors.New("Docno is confirm can not edit data")
	case dp.SumCreditAmount != 0 && (dp.CreditType == "" || dp.ConfirmNo == "" || dp.CreditRefNo == "" || dp.BankCode == "" || dp.BankBranchCode == "" ):
		return errors.New("Creditcard not have credittype or confirmno or creditrefno")
	case dp.SumChqAmount != 0 && (dp.BankCode == "" || dp.BankBranchCode == "" || dp.CreditRefNo == ""):
		return errors.New("Chq not have BankCode or BankBranchCode or creditrefno")
	case dp.SumBankAmount != 0 && dp.BankRefNo == "":
		return errors.New("Bank transfer not have BankRefNo")
	}

	if (dp.TaxNo == "") {
		dp.TaxNo = dp.DocNo
	}
	if (dp.TaxDate == "" ) {
		dp.TaxDate = dp.DocDate
	}
	if (dp.DueDate == "" && dp.CreditDay == 0) {
		dp.DueDate = dp.DocDate
	} else {
		dp.DueDate = now.AddDate(0, 0, dp.CreditDay).Format("2006-01-02")
	}
	if (dp.ExchangeRate == 0) {
		dp.ExchangeRate = 1
	}

	fmt.Println("UserCode = ", dp.UserCode)

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")

	//fmt.Println("TaxRate = ", def.TaxRateDefault)

	if dp.TaxRate == 0 {
		dp.TaxRate = def.TaxRateDefault
	}
	if dp.ExchangeRate == 0 {
		dp.ExchangeRate = def.ExchangeRateDefault
	}

	if (dp.BookCode == "") {
		dp.BookCode = def.ArDepositBookCode
	}

	dp.Source = def.ArDepositSource

	if (dp.GLFormat == "") {
		dp.GLFormat = def.ArDepositGLFormat
	}

	if dp.SaveFrom == 0 {
		dp.SaveFrom = def.ArDepositSaveFrom
	}

	if (dp.SumBankAmount != 0 && dp.TransBankDate == "") {
		dp.TransBankDate = dp.DocDate
	}

	if (dp.BeforeTaxAmount == 0 && dp.TaxAmount == 0) {
		dp.BeforeTaxAmount, dp.TaxAmount = m.CalcTaxDoc(dp.TaxType, dp.TaxRate, dp.TotalAmount)
	}

	fmt.Println("BeforeTax,TaxAmount", dp.BeforeTaxAmount, dp.TaxAmount)

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
			fmt.Println("Error = ", err.Error())
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
			fmt.Println("Error = ", err.Error())
			return err
		}
	}

	sqldel := `delete dbo.BCOutputTax where docno = ?`
	_, err = db.Exec(sqldel, dp.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}

	sqltax := "insert into dbo.BCOutputTax(SaveFrom,DocNo,BookCode,Source,DocDate,TaxDate,TaxNo,ArCode,ShortTaxDesc,TaxRate,Process,BeforeTaxAmount,TaxAmount,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?,?,'ขายสินค้า',?,1,?,?,?,getdate())"
	_, err = db.Exec(sqltax, dp.SaveFrom, dp.DocNo, dp.BookCode, dp.Source, dp.DocDate, dp.TaxDate, dp.TaxNo, dp.ArCode, dp.TaxRate, dp.BeforeTaxAmount, dp.TaxAmount, dp.CreatorCode)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}

	sqlrecdel := `delete dbo.BCRecMoney where docno = ?`
	_, err = db.Exec(sqlrecdel, dp.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}

	var my_description_recmoney string

	my_description_recmoney = "รับเงินมัดจำ"

	fmt.Println("RecMoney")
	var linenumber int


	if (dp.SumCashAmount != 0) { //subs.PaymentType == 0:
		fmt.Println("SumCashAmount")
		sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,PaymentType,SaveFrom,LineNumber,ProjectCode,DepartCode,SaleCode,MyDescription) values(?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlrec, dp.DocNo, dp.DocDate, dp.ArCode, dp.ExchangeRate, dp.SumCashAmount, 0,dp.SaveFrom, linenumber, dp.ProjectCode, dp.DepartCode, dp.SaleCode, my_description_recmoney)
		if err != nil {
			fmt.Println("Error = ", err.Error())
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
		sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,ChqTotalAmount,PaymentType,SaveFrom,CreditType,ConfirmNo,LineNumber,RefNo,BankCode,BankBranchCode,ProjectCode,DepartCode,SaleCode,MyDescription,RefDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlrec, dp.DocNo, dp.DocDate, dp.ArCode, dp.ExchangeRate, dp.SumCreditAmount, dp.SumCreditAmount, 1,dp.SaveFrom, dp.CreditType, dp.ConfirmNo, linenumber, dp.CreditRefNo, dp.BankCode, dp.BankBranchCode, dp.ProjectCode, dp.DepartCode, dp.SaleCode, my_description_recmoney, dp.DocDate)
		if err != nil {
			fmt.Println("Error = ", err.Error())
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

		sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,PaymentType,SaveFrom,LineNumber,RefNo,BankCode,ProjectCode,DepartCode,SaleCode,BankBranchCode,MyDescription,RefDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlrec, dp.DocNo, dp.DocDate, dp.ArCode, dp.ExchangeRate, dp.SumChqAmount, 2,dp.SaveFrom, linenumber, dp.CreditRefNo, dp.BankCode, dp.ProjectCode, dp.DepartCode, dp.SaleCode, dp.BankBranchCode, my_description_recmoney, dp.RefDate)
		if err != nil {
			fmt.Println("Error = ", err.Error())
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

		sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,PaymentType,SaveFrom,LineNumber,RefNo,ProjectCode,DepartCode,SaleCode,MyDescription,RefDate,TransBankDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlrec, dp.DocNo, dp.DocDate, dp.ArCode, dp.ExchangeRate, dp.SumBankAmount, 3,dp.SaveFrom, linenumber, dp.BankRefNo, dp.ProjectCode, dp.DepartCode, dp.SaleCode, my_description_recmoney, dp.DocDate, dp.TransBankDate)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}
	}

	if (dp.SumChqAmount != 0) {
		sqlchqdel := `delete dbo.BCChqIn where docno = ?`
		_, err = db.Exec(sqlchqdel, dp.DocNo)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}

		for _, c := range dp.Chqs {
			if ((c.ReceiveDate == "") || (c.DueDate == "")) {
				c.ReceiveDate = dp.DocDate
				c.DueDate = dp.DocDate
			}

			sqldep := `insert into dbo.bcchqin(BankCode,ChqNumber,DocNo,ArCode,SaleCode,ReceiveDate,DueDate,BookNo,Status,SaveFrom,StatusDate,StatusDocNo,DepartCode,BankBranchCode,Amount,Balance,MyDescription,ExchangeRate,RefChqRowOrder) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
			_, err = db.Exec(sqldep, c.BankCode, c.ChqNumber, dp.DocNo, dp.ArCode, dp.SaleCode, c.ReceiveDate, c.DueDate, c.BookNo, c.Status, dp.SaveFrom, c.StatusDate, c.StatusDocNo, dp.DepartCode, c.BankBranchCode, c.Amount, c.Balance, my_description_recmoney, dp.ExchangeRate, c.RefChqRowOrder)
			if err != nil {
				fmt.Println("Chq", err.Error())
				return err
			}
		}
	}

	if (dp.SumCreditAmount != 0) {
		sqlcrddel := `delete dbo.BCCreditCard where docno = ?`
		_, err = db.Exec(sqlcrddel, dp.DocNo)
		if err != nil {
			return err
		}

		fmt.Println("Cdcs ren =", len((dp.Cdcs)))
		if len((dp.Cdcs)) != 0 {
			for _, d := range dp.Cdcs {
				sqlcrd := `insert into dbo.bccreditcard(BankCode,CreditCardNo,DocNo,ArCode,SaleCode,ReceiveDate,DueDate,BookNo,Status,SaveFrom,StatusDate,StatusDocNo,DepartCode,BankBranchCode,Amount,MyDescription,ExchangeRate,CreditType,ConfirmNo,ChargeAmount,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,getdate())`
				_, err = db.Exec(sqlcrd, d.BankCode, d.CreditCardNo, dp.DocNo, dp.ArCode, dp.SaleCode, d.ReceiveDate, d.DueDate, d.BookNo, d.Status, dp.SaveFrom, d.StatusDate, d.StatusDocNo, dp.DepartCode, d.BankBranchCode, d.Amount, my_description_recmoney, dp.ExchangeRate, d.CreditType, d.ConfirmNo, d.ChargeAmount, dp.UserCode)
				if err != nil {
					fmt.Println("Credit", err.Error())
					return err
				}
			}
		} else {
			BookNo := ""
			Status := 0

			sqlcrd := `insert into dbo.bccreditcard(BankCode,CreditCardNo,DocNo,ArCode,SaleCode,ReceiveDate,DueDate,BookNo,Status,SaveFrom,StatusDate,StatusDocNo,DepartCode,BankBranchCode,Amount,MyDescription,ExchangeRate,CreditType,ConfirmNo,ChargeAmount,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,getdate())`
			_, err = db.Exec(sqlcrd, dp.BankCode, dp.CreditRefNo, dp.DocNo, dp.ArCode, dp.SaleCode, dp.DocDate, dp.DueDate, BookNo, Status, dp.SaveFrom, dp.DocDate, dp.DocNo, dp.DepartCode, dp.BankBranchCode, dp.SumCreditAmount, my_description_recmoney, dp.ExchangeRate, dp.CreditType, dp.ConfirmNo, dp.ChargeAmount, dp.UserCode)
			if err != nil {
				fmt.Println("Credit", err.Error())
				return err
			}
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
	sql := `set dateformat dmy     select a.RowOrder,a.DocNo as DocNo,a.DocDate,isnull(a.TaxDate,'') as TaxDate,a.TaxType,isnull(a.TaxNo,'') as TaxNo,a.ArCode,isnull(b.name1,'') as ArName,isnull(a.DepartCode,'') as DepartCode,a.CreditDay,a.DueDate,isnull(a.SaleCode,'') as SaleCode,isnull(c.name,'') as SaleName,a.TaxRate,isnull(a.MyDescription,'') as MyDescription,a.BeforeTaxAmount,a.TaxAmount,a.TotalAmount,a.SumOfWTax,a.NetAmount,a.BillBalance,a.OtherIncome,a.OtherExpense,a.ExcessAmount1,a.ExcessAmount2,a.ChargeAmount,a.ChangeAmount,isnull(a.RefNo,'') as RefNo,isnull(a.CurrencyCode,'') as CurrencyCode,a.ExchangeRate,a.SumCashAmount,a.SumChqAmount,a.SumCreditAmount,a.SumBankAmount,isnull(a.GLFormat,'') as GLFormat,isnull(a.AllocateCode,'') as AllocateCode,isnull(a.ProjectCode,'') as ProjectCode,isnull(a.BillGroup,'') as BillGroup,isnull(a.RecurName,'') as RecurName,a.CreatorCode,a.CreateDateTime,isnull(d.bookcode,'') as BookCode from dbo.bcardeposit a WITH (NOLOCK) left join dbo.bcar b WITH (NOLOCK) on a.arcode = b.code left join dbo.bcsale c on a.salecode = c.code left join dbo.bcoutputtax d WITH (NOLOCK) on d.docno = a.docno and d.arcode = a.arcode where a.docno  = ?`
	err := db.Get(dp, sql, docno)
	fmt.Println("sql =", sql)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}
	return nil
}

func (dp *ArDeposit) SearchArDepositByKeyword(db *sqlx.DB, keyword string) (dps []*ArDeposit, err error) {
	sql := `set dateformat dmy     select a.RowOrder,a.DocNo as DocNo,a.DocDate,isnull(a.TaxDate,'') as TaxDate,a.TaxType,isnull(a.TaxNo,'') as TaxNo,a.ArCode,isnull(b.name1,'') as ArName,isnull(a.DepartCode,'') as DepartCode,a.CreditDay,a.DueDate,isnull(a.SaleCode,'') as SaleCode,isnull(c.name,'') as SaleName,a.TaxRate,isnull(a.MyDescription,'') as MyDescription,a.BeforeTaxAmount,a.TaxAmount,a.TotalAmount,a.SumOfWTax,a.NetAmount,a.BillBalance,a.OtherIncome,a.OtherExpense,a.ExcessAmount1,a.ExcessAmount2,a.ChargeAmount,a.ChangeAmount,isnull(a.RefNo,'') as RefNo,isnull(a.CurrencyCode,'') as CurrencyCode,a.ExchangeRate,a.SumCashAmount,a.SumChqAmount,a.SumCreditAmount,a.SumBankAmount,isnull(a.GLFormat,'') as GLFormat,isnull(a.AllocateCode,'') as AllocateCode,isnull(a.ProjectCode,'') as ProjectCode,isnull(a.BillGroup,'') as BillGroup,isnull(a.RecurName,'') as RecurName,a.CreatorCode,a.CreateDateTime,isnull(d.bookcode,'') as BookCode from dbo.bcardeposit a WITH (NOLOCK) left join dbo.bcar b WITH (NOLOCK) on a.arcode = b.code left join dbo.bcsale c on a.salecode = c.code left join dbo.bcoutputtax d WITH (NOLOCK) on d.docno = a.docno and d.arcode = a.arcode where (a.docno  like '%'+?+'%' or a.arcode like '%'+?+'%' or a.salecode like '%'+?+'%' or b.name1 like '%'+?+'%')  order by a.docno`
	err = db.Select(&dps, sql, keyword, keyword, keyword, keyword)
	fmt.Println("sql =", sql)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}

	return dps, nil
}
