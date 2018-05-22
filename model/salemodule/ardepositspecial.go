package model

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	m "github.com/loukmho/bcaccount_api/model"
	"time"
)

type ArDepositSpecial struct {
	SaveFrom        int                    `json:"save_from" db"SaveFrom"`
	DocNo           string                 `json:"doc_no" db:"DocNo"`
	DocDate         string                 `json:"doc_date" db:"DocDate"`
	ArCode          string                 `json:"ar_code" db:"ArCode"`
	DpsCustomer
	DpsSaleMan
	Source          int                    `json:"source" db:"Source"`
	CreatorCode     string                 `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string                 `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string                 `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string                 `json:"last_edit_date_t" db:"LastEditDateT"`
	DepartCode      string                 `json:"depart_code"  db:"DepartCode"`
	IsConfirm       int                    `json:"is_confirm" db:"IsConfirm"`
	CreditDay       int                    `json:"credit_day" db:"CreditDay"`
	DueDate         string                 `json:"due_date" db:"DueDate"`
	MyDescription   string                 `json:"my_description" db:"MyDescription"`
	TotalAmount     float64                `json:"total_amount" db:"TotalAmount"`
	SumOfWTax       float64                `json:"sum_of_w_tax" db:"SumOfWTax"`
	BeforeTaxAmount float64                `json:"before_tax_amount" db:"BeforeTaxAmount"`
	NetAmount       float64                `json:"net_amount" db:"NetAmount"`
	BillBalance     float64                `json:"bill_balance" db:"BillBalance"`
	ExcessAmount1   float64                `json:"excess_amount_1" db:"ExcessAmount1"`
	ExcessAmount2   float64                `json:"excess_amount_2" db:"ExcessAmount2"`
	ChargeAmount    float64                `json:"charge_amount" db:"ChargeAmount"`
	ChangeAmount    float64                `json:"change_amount" db:"ChangeAmount"`
	OtherIncome     float64                `json:"other_income" db:"OtherIncome"`
	OtherExpense    float64                `json:"other_expense" db:"OtherExpense"`
	RefNo           string                 `json:"ref_no" db:"RefNo"`
	CurrencyCode    string                 `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate    float64                `json:"exchange_rate" db:"ExchangeRate"`
	SumCashAmount   float64                `json:"sum_cash_amount" db:"SumCashAmount"`
	SumChqAmount    float64                `json:"sum_chq_amount" db:"SumChqAmount"`
	SumCreditAmount float64                `json:"sum_credit_amount" db:"SumCreditAmount"`
	SumBankAmount   float64                `json:"sum_bank_amount" db:"SumBankAmount"`
	PettyCashAmount float64                `json:"petty_cash_amount" db:"PettyCashAmount"`
	GLFormat        string                 `json:"gl_format" db:"GLFormat"`
	IsCancel        int                    `json:"is_cancel" db:"IsCancel"`
	IsReturnMoney   int                    `json:"is_return_money" db:"IsReturnMoney"`
	AllocateCode    string                 `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string                 `json:"project_code" db:"ProjectCode"`
	BillGroup       string                 `json:"bill_group" db:"BillGroup"`
	RecurName       string                 `json:"recur_name" db:"RecurName"`
	ConfirmCode     string                 `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string                 `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string                 `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string                 `json:"cancel_date_time" db:"CancelDateTime"`
	UserCode        string                 `json:"user_code" db:"UserCode"`
	ArDpsOutPutTax
	ListArDpsRecMoney
	Cdcs            []*ListArDpsCreditCard `json:"cdcs"`
	Chqs            []*ListArDpsChqIn      `json:"chqs"`
}

type DpsCustomer struct {
	ArCode string `json:"ar_code" db:"ArCode"`
	ArName string `json:"ar_name" db:"ArName"`
}

type DpsSaleMan struct {
	SaleCode string `json:"sale_code" db:"SaleCode"`
	SaleName string `json:"sale_name" db:"SaleName"`
}

type ArDpsOutPutTax struct {
	TaxNo    string `json:"tax_no" db:"TaxNo"`
	TaxDate  string `json:"tax_date" db:"TaxDate"`
	BookCode string `json:"book_code" db:"BookCode"`
}

type ListArDpsRecMoney struct {
	CreditType     string `json:"credit_type" db:"CreditType"`
	ConfirmNo      string `json:"confirm_no" db:"ConfirmNo"`
	CreditRefNo    string `json:"credit_ref_no" db:"CreditRefNo"`
	BankCode       string `json:"bank_code" db:"BookCode"`
	BankBranchCode string `json:"bank_branch_code" db:"BankBranchCode"`
	BankRefNo      string `json:"bank_ref_no" db:"BankRefNo"`
	TransBankDate  string `json:"trans_bank_date" db:"TransBankDate"`
	RefDate        string `json:"ref_date" db:"RefDate"`
}

type ListArDpsChqIn struct {
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

type ListArDpsCreditCard struct {
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

func (dps *ArDepositSpecial) InsertAndEditArDepositSpecial(db *sqlx.DB) error {
	var check_exist int
	var sum_pay_amount float64

	now := time.Now()

	sqlexist := `select count(docno) as check_exist from dbo.bcardepositspecial where docno = ?` //เช็คว่ามีเอกสารหรือยัง
	err := db.Get(&check_exist, sqlexist, dps.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil
	}

	sum_pay_amount = (dps.SumCashAmount + dps.SumCreditAmount + dps.SumChqAmount + dps.SumBankAmount + dps.OtherExpense) - dps.OtherIncome

	switch {
	case dps.DocNo == "":
		return errors.New("Docno is null")
	case dps.ArCode == "":
		return errors.New("Arcode is null")
	case dps.TotalAmount == 0:
		return errors.New("Totalamount is 0")
	case dps.DocDate == "":
		return errors.New("Docdate is null")
	case dps.SumCashAmount == 0 && dps.SumCreditAmount == 0 && dps.SumChqAmount == 0 && dps.SumBankAmount == 0:
		return errors.New("Docno not set money to another type payment")
	case sum_pay_amount > dps.TotalAmount:
		return errors.New("Rec money is over totalamount")
	case dps.IsCancel == 1:
		return errors.New("Docno is cancel can not edit data")
	case dps.IsConfirm == 1:
		return errors.New("Docno is confirm can not edit data")
	case dps.SumCreditAmount != 0 && (dps.CreditType == "" || dps.ConfirmNo == "" || dps.CreditRefNo == "" || dps.BankCode == "" || dps.BankBranchCode == "" ):
		return errors.New("Creditcard not have credittype or confirmno or creditrefno")
	case dps.SumChqAmount != 0 && (dps.BankCode == "" || dps.BankBranchCode == "" || dps.CreditRefNo == ""):
		return errors.New("Chq not have BankCode or BankBranchCode or creditrefno")
	case dps.SumBankAmount != 0 && dps.BankRefNo == "":
		return errors.New("Bank transfer not have BankRefNo")
	}

	if (dps.DueDate == "" && dps.CreditDay == 0) {
		dps.DueDate = dps.DocDate
	} else {
		dps.DueDate = now.AddDate(0, 0, dps.CreditDay).Format("2006-01-02")
	}

	fmt.Println("UserCode = ", dps.UserCode)

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")

	if dps.ExchangeRate == 0 {
		dps.ExchangeRate = def.ExchangeRateDefault
	}

	if (dps.BookCode == "") {
		dps.BookCode = def.ArDepositBookCode
	}

	dps.Source = def.ArDepositSource

	if (dps.GLFormat == "") {
		dps.GLFormat = def.ArDepositGLFormat
	}

	if dps.SaveFrom == 0 {
		dps.SaveFrom = def.ArDepositSaveFrom
	}

	if (dps.SumBankAmount != 0 && dps.TransBankDate == "") {
		dps.TransBankDate = dps.DocDate
	}

	if (dps.BeforeTaxAmount == 0 ) {
		dps.BeforeTaxAmount = dps.TotalAmount
	}

	fmt.Println("check_exist = ", check_exist)

	if (check_exist == 0) {
		//Insert//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		dps.CreatorCode = dps.UserCode
		dps.BillBalance = dps.TotalAmount
		dps.NetAmount = dps.TotalAmount

		sql := `Insert into dbo.BCArDepositSpecial(DocNo,DocDate,ArCode,CreatorCode,CreateDateTime,DepartCode,SaleCode,IsConfirm,CreditDay,DueDate,MyDescription,TotalAmount,SumOfWTax,BeforeTaxAmount,NetAmount,BillBalance,ExcessAmount1,ExcessAmount2,ChargeAmount,ChangeAmount,OtherIncome,OtherExpense,RefNo,CurrencyCode,ExchangeRate,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,PettyCashAmount,GLFormat,IsCancel,IsReturnMoney,AllocateCode,ProjectCode,BillGroup,RecurName) values(?,?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sql, dps.DocNo,dps.DocDate,dps.ArCode,dps.UserCode,dps.DepartCode,dps.SaleCode,dps.IsConfirm,dps.CreditDay,dps.DueDate,dps.MyDescription,dps.TotalAmount,dps.SumOfWTax,dps.BeforeTaxAmount,dps.NetAmount,dps.BillBalance,dps.ExcessAmount1,dps.ExcessAmount2,dps.ChargeAmount,dps.ChangeAmount,dps.OtherIncome,dps.OtherExpense,dps.RefNo,dps.CurrencyCode,dps.ExchangeRate,dps.SumCashAmount,dps.SumChqAmount,dps.SumCreditAmount,dps.SumBankAmount,dps.PettyCashAmount,dps.GLFormat,dps.IsCancel,dps.IsReturnMoney,dps.AllocateCode,dps.ProjectCode,dps.BillGroup,dps.RecurName)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}

	} else {

		//Update/////////////////////////////////////////////////////////////////////////////////////////////////////////
		dps.LastEditorCode = dps.UserCode
		dps.BillBalance = dps.TotalAmount
		dps.NetAmount = dps.TotalAmount

		sql := `Update dbo.BCArDepositSpecial set DocDate=?,ArCode=?,DepartCode=?,SaleCode=?,IsConfirm=?,CreditDay=?,DueDate=?,MyDescription=?,TotalAmount=?,SumOfWTax=?,BeforeTaxAmount=?,NetAmount=?,BillBalance=?,ExcessAmount1=?,ExcessAmount2=?,ChargeAmount=?,ChangeAmount=?,OtherIncome=?,OtherExpense=?,RefNo=?,CurrencyCode=?,ExchangeRate=?,SumCashAmount=?,SumChqAmount=?,SumCreditAmount=?,SumBankAmount=?,PettyCashAmount=?,GLFormat=?,IsCancel=?,IsReturnMoney=?,AllocateCode=?,ProjectCode=?,BillGroup=?,RecurName=?,LastEditorCode=?,LastEditDateT=getdate() where docno = ?`
		_, err := db.Exec(sql, dps.DocDate,dps.ArCode,dps.DepartCode,dps.SaleCode,dps.IsConfirm,dps.CreditDay,dps.DueDate,dps.MyDescription,dps.TotalAmount,dps.SumOfWTax,dps.BeforeTaxAmount,dps.NetAmount,dps.BillBalance,dps.ExcessAmount1,dps.ExcessAmount2,dps.ChargeAmount,dps.ChangeAmount,dps.OtherIncome,dps.OtherExpense,dps.RefNo,dps.CurrencyCode,dps.ExchangeRate,dps.SumCashAmount,dps.SumChqAmount,dps.SumCreditAmount,dps.SumBankAmount,dps.PettyCashAmount,dps.GLFormat,dps.IsCancel,dps.IsReturnMoney,dps.AllocateCode,dps.ProjectCode,dps.BillGroup,dps.RecurName,dps.UserCode, dps.DocNo)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}
	}

	sqlrecdel := `delete dbo.BCRecMoney where docno = ?`
	_, err = db.Exec(sqlrecdel, dps.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}

	var my_description_recmoney string

	my_description_recmoney = "รับเงินมัดจำ"

	fmt.Println("RecMoney")
	var linenumber int

	if (dps.SumCashAmount != 0) { //subs.PaymentType == 0:
		fmt.Println("SumCashAmount")
		sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,PaymentType,SaveFrom,LineNumber,ProjectCode,DepartCode,SaleCode,MyDescription) values(?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlrec, dps.DocNo, dps.DocDate, dps.ArCode, dps.ExchangeRate, dps.SumCashAmount, 0, dps.SaveFrom, linenumber, dps.ProjectCode, dps.DepartCode, dps.SaleCode, my_description_recmoney)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}
	}
	//case dp.SumCreditAmount != 0: //subs.PaymentType == 1:
	if (dps.SumCreditAmount != 0) {
		fmt.Println("SumCreditAmount")
		if (dps.SumCashAmount != 0) {
			linenumber = 1
		} else {
			linenumber = 0
		}
		sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,ChqTotalAmount,PaymentType,SaveFrom,CreditType,ConfirmNo,LineNumber,RefNo,BankCode,BankBranchCode,ProjectCode,DepartCode,SaleCode,MyDescription,RefDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlrec, dps.DocNo, dps.DocDate, dps.ArCode, dps.ExchangeRate, dps.SumCreditAmount, dps.SumCreditAmount, 1, dps.SaveFrom, dps.CreditType, dps.ConfirmNo, linenumber, dps.CreditRefNo, dps.BankCode, dps.BankBranchCode, dps.ProjectCode, dps.DepartCode, dps.SaleCode, my_description_recmoney, dps.DocDate)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}
	}

	//case dp.SumChqAmount != 0: //subs.PaymentType == 2:
	if (dps.SumChqAmount != 0) {
		fmt.Println("SumChqAmount")
		if (dps.SumCashAmount != 0 && dps.SumCreditAmount != 0) {
			linenumber = 2
		} else if ((dps.SumCashAmount != 0 && dps.SumCreditAmount == 0)) {
			linenumber = 1
		} else if ((dps.SumCashAmount == 0 && dps.SumCreditAmount != 0)) {
			linenumber = 1
		} else if ((dps.SumCashAmount == 0 && dps.SumCreditAmount == 0)) {
			linenumber = 0
		}

		sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,PaymentType,SaveFrom,LineNumber,RefNo,BankCode,ProjectCode,DepartCode,SaleCode,BankBranchCode,MyDescription,RefDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlrec, dps.DocNo, dps.DocDate, dps.ArCode, dps.ExchangeRate, dps.SumChqAmount, 2, dps.SaveFrom, linenumber, dps.CreditRefNo, dps.BankCode, dps.ProjectCode, dps.DepartCode, dps.SaleCode, dps.BankBranchCode, my_description_recmoney, dps.RefDate)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}
	}

	//case dp.SumBankAmount != 0: //subs.PaymentType == 3:
	if (dps.SumBankAmount != 0) {
		fmt.Println("SumBankAmount")
		if (dps.SumCashAmount != 0 && dps.SumCreditAmount != 0 && dps.SumChqAmount != 0) {
			linenumber = 3
		} else if (dps.SumCashAmount != 0 && dps.SumCreditAmount == 0 && dps.SumChqAmount != 0) {
			linenumber = 2
		} else if (dps.SumCashAmount == 0 && dps.SumCreditAmount != 0 && dps.SumChqAmount != 0) {
			linenumber = 2
		} else if (dps.SumCashAmount != 0 && dps.SumCreditAmount != 0 && dps.SumChqAmount == 0) {
			linenumber = 2
		} else if (dps.SumCashAmount != 0 && dps.SumCreditAmount != 0 && dps.SumChqAmount == 0) {
			linenumber = 2
		} else if (dps.SumCashAmount != 0 && dps.SumCreditAmount == 0 && dps.SumChqAmount == 0) {
			linenumber = 1
		} else if (dps.SumCashAmount == 0 && dps.SumCreditAmount != 0 && dps.SumChqAmount == 0) {
			linenumber = 1
		} else if (dps.SumCashAmount == 0 && dps.SumCreditAmount == 0 && dps.SumChqAmount != 0) {
			linenumber = 1
		} else if (dps.SumCashAmount == 0 && dps.SumCreditAmount == 0 && dps.SumChqAmount == 0) {
			linenumber = 0
		}

		sqlrec := `insert	into dbo.BCRecMoney(DocNo,DocDate,ArCode,ExchangeRate,PayAmount,PaymentType,SaveFrom,LineNumber,RefNo,ProjectCode,DepartCode,SaleCode,MyDescription,RefDate,TransBankDate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlrec, dps.DocNo, dps.DocDate, dps.ArCode, dps.ExchangeRate, dps.SumBankAmount, 3, dps.SaveFrom, linenumber, dps.BankRefNo, dps.ProjectCode, dps.DepartCode, dps.SaleCode, my_description_recmoney, dps.DocDate, dps.TransBankDate)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}
	}

	if (dps.SumChqAmount != 0) {
		sqlchqdel := `delete dbo.BCChqIn where docno = ?`
		_, err = db.Exec(sqlchqdel, dps.DocNo)
		if err != nil {
			return err
		}

		for _, c := range dps.Chqs {
			if ((c.ReceiveDate == "") || (c.DueDate == "")) {
				c.ReceiveDate = dps.DocDate
				c.DueDate = dps.DocDate
			}

			sqldep := `insert into dbo.bcchqin(BankCode,ChqNumber,DocNo,ArCode,SaleCode,ReceiveDate,DueDate,BookNo,Status,SaveFrom,StatusDate,StatusDocNo,DepartCode,BankBranchCode,Amount,Balance,MyDescription,ExchangeRate,RefChqRowOrder) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
			_, err = db.Exec(sqldep, c.BankCode, c.ChqNumber, dps.DocNo, dps.ArCode, dps.SaleCode, c.ReceiveDate, c.DueDate, c.BookNo, c.Status, dps.SaveFrom, c.StatusDate, c.StatusDocNo, dps.DepartCode, c.BankBranchCode, c.Amount, c.Balance, my_description_recmoney, dps.ExchangeRate, c.RefChqRowOrder)
			if err != nil {
				fmt.Println("Chq", err.Error())
				return err
			}
		}
	}

	if (dps.SumCreditAmount != 0) {
		sqlcrddel := `delete dbo.BCCreditCard where docno = ?`
		_, err = db.Exec(sqlcrddel, dps.DocNo)
		if err != nil {
			return err
		}

		fmt.Println("Cdcs ren =", len((dps.Cdcs)))
		if len((dps.Cdcs)) != 0 {
			for _, d := range dps.Cdcs {
				sqlcrd := `insert into dbo.bccreditcard(BankCode,CreditCardNo,DocNo,ArCode,SaleCode,ReceiveDate,DueDate,BookNo,Status,SaveFrom,StatusDate,StatusDocNo,DepartCode,BankBranchCode,Amount,MyDescription,ExchangeRate,CreditType,ConfirmNo,ChargeAmount,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,getdate())`
				_, err = db.Exec(sqlcrd, d.BankCode, d.CreditCardNo, dps.DocNo, dps.ArCode, dps.SaleCode, d.ReceiveDate, d.DueDate, d.BookNo, d.Status, dps.SaveFrom, d.StatusDate, d.StatusDocNo, dps.DepartCode, d.BankBranchCode, d.Amount, my_description_recmoney, dps.ExchangeRate, d.CreditType, d.ConfirmNo, d.ChargeAmount, dps.UserCode)
				if err != nil {
					fmt.Println("Credit", err.Error())
					return err
				}
			}
		} else {
			BookNo := ""
			Status := 0

			sqlcrd := `insert into dbo.bccreditcard(BankCode,CreditCardNo,DocNo,ArCode,SaleCode,ReceiveDate,DueDate,BookNo,Status,SaveFrom,StatusDate,StatusDocNo,DepartCode,BankBranchCode,Amount,MyDescription,ExchangeRate,CreditType,ConfirmNo,ChargeAmount,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,getdate())`
			_, err = db.Exec(sqlcrd, dps.BankCode, dps.CreditRefNo, dps.DocNo, dps.ArCode, dps.SaleCode, dps.DocDate, dps.DueDate, BookNo, Status, dps.SaveFrom, dps.DocDate, dps.DocNo, dps.DepartCode, dps.BankBranchCode, dps.SumCreditAmount, my_description_recmoney, dps.ExchangeRate, dps.CreditType, dps.ConfirmNo, dps.ChargeAmount, dps.UserCode)
			if err != nil {
				fmt.Println("Credit", err.Error())
				return err
			}
		}

	}

	return nil
}

func (dps *ArDepositSpecial) SearchArDepositSpecialByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select a.DocNo,a.DocDate,a.ArCode,isnull(b.name1,'') as ArName,isnull(a.CreatorCode,'') as CreatorCode,isnull(a.CreateDateTime,'') as CreateDateTime,isnull(a.LastEditorCode,'') as LastEditorCode,isnull(a.LastEditDateT,'') as LastEditDateT,isnull(a.DepartCode,'') as DepartCode,isnull(a.SaleCode,'') as SaleCode,isnull(c.name,'') as SaleName,a.IsConfirm,a.CreditDay,isnull(a.DueDate,'') as DueDate,isnull(a.MyDescription,'') as MyDescription,a.TotalAmount,a.SumOfWTax,a.BeforeTaxAmount,a.NetAmount,a.BillBalance,a.ExcessAmount1,a.ExcessAmount2,a.ChargeAmount,a.ChangeAmount,a.OtherIncome,a.OtherExpense,isnull(a.RefNo,'') as RefNo,isnull(a.CurrencyCode,'') as CurrencyCode,a.ExchangeRate,a.SumCashAmount,a.SumChqAmount,a.SumCreditAmount,a.SumBankAmount,a.PettyCashAmount,isnull(a.GLFormat,'') as GLFormat,a.IsCancel,a.IsReturnMoney,isnull(a.AllocateCode,'') as AllocateCode,isnull(a.ProjectCode,'') as ProjectCode,isnull(a.BillGroup,'') as BillGroup,isnull(a.RecurName,'') as RecurName,isnull(a.ConfirmCode,'') as ConfirmCode,isnull(a.ConfirmDateTime,'') as ConfirmDateTime,isnull(a.CancelCode,'') as CancelCode,isnull(a.CancelDateTime,'') as CancelDateTime from dbo.BCArDepositSpecial a WITH (NOLOCK) left join dbo.BCAR b WITH (NOLOCK) on a.arcode = b.code left join dbo.BCSale c WITH (NOLOCK) on a.salecode = c.code where a.docno = ?`
	err := db.Get(dps, sql, docno)
	fmt.Println("sql =", sql)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}
	return nil
}

func (dps *ArDepositSpecial) SearchArDepositSpecialByKeyword(db *sqlx.DB, keyword string) (dpsList []*ArDepositSpecial, err error) {
	sql := `set dateformat dmy     select a.DocNo,a.DocDate,a.ArCode,isnull(b.name1,'') as ArName,isnull(a.CreatorCode,'') as CreatorCode,isnull(a.CreateDateTime,'') as CreateDateTime,isnull(a.LastEditorCode,'') as LastEditorCode,isnull(a.LastEditDateT,'') as LastEditDateT,isnull(a.DepartCode,'') as DepartCode,isnull(a.SaleCode,'') as SaleCode,isnull(c.name,'') as SaleName,a.IsConfirm,a.CreditDay,isnull(a.DueDate,'') as DueDate,isnull(a.MyDescription,'') as MyDescription,a.TotalAmount,a.SumOfWTax,a.BeforeTaxAmount,a.NetAmount,a.BillBalance,a.ExcessAmount1,a.ExcessAmount2,a.ChargeAmount,a.ChangeAmount,a.OtherIncome,a.OtherExpense,isnull(a.RefNo,'') as RefNo,isnull(a.CurrencyCode,'') as CurrencyCode,a.ExchangeRate,a.SumCashAmount,a.SumChqAmount,a.SumCreditAmount,a.SumBankAmount,a.PettyCashAmount,isnull(a.GLFormat,'') as GLFormat,a.IsCancel,a.IsReturnMoney,isnull(a.AllocateCode,'') as AllocateCode,isnull(a.ProjectCode,'') as ProjectCode,isnull(a.BillGroup,'') as BillGroup,isnull(a.RecurName,'') as RecurName,isnull(a.ConfirmCode,'') as ConfirmCode,isnull(a.ConfirmDateTime,'') as ConfirmDateTime,isnull(a.CancelCode,'') as CancelCode,isnull(a.CancelDateTime,'') as CancelDateTime from dbo.BCArDepositSpecial a WITH (NOLOCK) left join dbo.BCAR b WITH (NOLOCK) on a.arcode = b.code left join dbo.BCSale c WITH (NOLOCK) on a.salecode = c.code where (a.docno  like '%'+?+'%' or a.arcode like '%'+?+'%' or a.salecode like '%'+?+'%' or b.name1 like '%'+?+'%')  order by a.docno`
	err = db.Select(&dpsList, sql, keyword, keyword, keyword, keyword)
	fmt.Println("sql =", sql)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}
	return dpsList, nil
}
