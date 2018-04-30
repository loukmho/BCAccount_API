package model

import (
	"github.com/jmoiron/sqlx"
	"time"
	"fmt"
	"errors"
	m "github.com/loukmho/bcaccount_api/model"
)


type ReturnDepSpecial struct {
	DocNo           string  `json:"doc_no" db:"DocNo"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	EarnestNo       string  `json:"earnest_no" db:"EarnestNo"`
	RdsCustomer
	RdsSaleMan
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	CreditDay       int     `json:"credit_day" db:"CreditDay"`
	DueDate         string  `json:"due_date" db:"DueDate"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	TotalAmount     float64 `json:"total_amount" db:"TotalAmount"`
	BeforeTaxAmount float64 `json:"before_tax_amount" db:"BeforeTaxAmount"`
	NetAmount       float64 `json:"net_amount" db:"NetAmount"`
	SumOfWTax       float64 `json:"sum_of_w_tax" db:"SumOfWTax"`
	ExcessAmount1   float64 `json:"excess_amount_1" db:"ExcessAmount1"`
	ExcessAmount2   float64 `json:"excess_amount_2" db:"ExcessAmount2"`
	BillBalance     float64 `json:"bill_balance" db:"BillBalance"`
	CurrencyCode    string  `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate    float64 `json:"exchange_rate" db:"ExchangeRate"`
	PettyCashAmount float64 `json:"petty_cash_amount" db:"PettyCashAmount"`
	SumCashAmount   float64 `json:"sum_cash_amount" db:"SumCashAmount"`
	SumChqAmount    float64 `json:"sum_chq_amount" db:"SumChqAmount"`
	SumCreditAmount float64 `json:"sum_credit_amount" db:"SumCreditAmount"`
	SumBankAmount   float64 `json:"sum_bank_amount" db:"SumBankAmount"`
	EarnestAmount   float64 `json:"earnest_amount" db:"EarnestAmount"`
	EarnestBalance  float64 `json:"earnest_balance" db:"EarnestBalance"`
	OtherIncome     float64 `json:"other_income" db:"OtherIncome"`
	OtherExpense    float64 `json:"other_expense" db:"OtherExpense"`
	GLFormat        string  `json:"gl_format" db:"GLFormat"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	IsReturnMoney   int     `json:"is_return_money" db:"IsReturnMoney"`
	AllocateCode    string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string  `json:"project_code" db:"ProjectCode"`
	BillGroup       string  `json:"bill_group" db:"BillGroup"`
	RecurName       string  `json:"recur_name" db:"RecurName"`
	ConfirmCode     string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string  `json:"cancel_date_time" db:"CancelDateTime"`
	UserCode        string  `json:"user_code" db:"UserCode"`
}

type RdsCustomer struct {
	ArCode string `json:"ar_code" db:"ArCode"`
	ArName string `json:"ar_name" db:"ArName"`
}

type RdsSaleMan struct {
	SaleCode string `json:"sale_code" db:"SaleCode"`
	SaleName string `json:"sale_name" db:"SaleName"`
}


func (rds *ReturnDepSpecial) SearchReturnDepSpecialByDocNo() error {
	//sql := `DocNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, EarnestNo, ArCode, DepartCode, IsConfirm, CreditDay, SaleCode, DueDate, MyDescription, TotalAmount,BeforeTaxAmount, NetAmount, SumOfWTax, ExcessAmount1, ExcessAmount2, BillBalance, CurrencyCode, ExchangeRate, PettyCashAmount, SumCashAmount, SumChqAmount, SumCreditAmount, SumBankAmount,EarnestAmount, EarnestBalance, OtherIncome, OtherExpense, GLFormat, IsCancel, IsReturnMoney, AllocateCode, ProjectCode, BillGroup, RecurName, ConfirmCode, ConfirmDateTime, CancelCode,CancelDateTime`
	return nil
}

func (rds *ReturnDepSpecial) InsertAndEditReturnDepSpecial(db *sqlx.DB) error{
	var check_exist int
	var sum_pay_amount float64

	now := time.Now()

	sqlexist := `select count(docno) as check_exist from dbo.BCReturnDepSpecial where docno = ?` //เช็คว่ามีเอกสารหรือยัง
	err := db.Get(&check_exist, sqlexist, rds.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	sum_pay_amount = (rds.SumCashAmount + rds.SumCreditAmount + rds.SumChqAmount + rds.SumBankAmount + rds.OtherExpense) - rds.OtherIncome

	switch {
	case rds.DocNo == "":
		return errors.New("Docno is null")
	case rds.ArCode == "":
		return errors.New("Arcode is null")
	case rds.TotalAmount == 0:
		return errors.New("Totalamount is 0")
	case rds.DocDate == "":
		return errors.New("Docdate is null")
	case rds.SumCashAmount == 0 && rds.SumCreditAmount == 0 && rds.SumChqAmount == 0 && rds.SumBankAmount == 0:
		return errors.New("Docno not set money to another type payment")
	case sum_pay_amount > rds.TotalAmount:
		return errors.New("Rec money is over totalamount")
	case rds.IsCancel == 1:
		return errors.New("Docno is cancel can not edit data")
	case rds.IsConfirm == 1:
		return errors.New("Docno is confirm can not edit data")
	}

	if (rds.DueDate == "" && rds.CreditDay == 0) {
		rds.DueDate = rds.DocDate
	} else {
		rds.DueDate = now.AddDate(0, 0, rds.CreditDay).Format("2006-01-02")
	}
	if (rds.ExchangeRate == 0) {
		rds.ExchangeRate = 1
	}

	fmt.Println("UserCode = ", rds.UserCode)

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")

	if rds.ExchangeRate == 0 {
		rds.ExchangeRate = def.ExchangeRateDefault
	}

	if (rds.GLFormat == "") {
		rds.GLFormat = def.ArDepositGLFormat
	}

	if (rds.BeforeTaxAmount == 0 ) {
		rds.BeforeTaxAmount = rds.TotalAmount
	}

	fmt.Println("check_exist = ", check_exist)

	rds.BillBalance = rds.TotalAmount
	rds.NetAmount = rds.TotalAmount

	if (check_exist == 0) {
		//Insert//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		rds.CreatorCode = rds.UserCode

		sql := `insert dbo.BCReturnDepSpecial(DocNo,DocDate,CreatorCode,CreateDateTime,EarnestNo,ArCode,DepartCode,IsConfirm,CreditDay,SaleCode,DueDate,MyDescription,TotalAmount,BeforeTaxAmount,NetAmount,SumOfWTax,ExcessAmount1,ExcessAmount2,BillBalance,CurrencyCode,ExchangeRate,PettyCashAmount,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,EarnestAmount,EarnestBalance,OtherIncome,OtherExpense,GLFormat,IsCancel,IsReturnMoney,AllocateCode,ProjectCode,BillGroup,RecurName) values(?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sql, rds.DocNo,rds.DocDate,rds.CreatorCode,rds.EarnestNo,rds.ArCode,rds.DepartCode,rds.IsConfirm,rds.CreditDay,rds.SaleCode,rds.DueDate,rds.MyDescription,rds.TotalAmount,rds.BeforeTaxAmount,rds.NetAmount,rds.SumOfWTax,rds.ExcessAmount1,rds.ExcessAmount2,rds.BillBalance,rds.CurrencyCode,rds.ExchangeRate,rds.PettyCashAmount,rds.SumCashAmount,rds.SumChqAmount,rds.SumCreditAmount,rds.SumBankAmount,rds.EarnestAmount,rds.EarnestBalance,rds.OtherIncome,rds.OtherExpense,rds.GLFormat,rds.IsCancel,rds.IsReturnMoney,rds.AllocateCode,rds.ProjectCode,rds.BillGroup,rds.RecurName)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

	} else {

		//Update/////////////////////////////////////////////////////////////////////////////////////////////////////////
		rds.LastEditorCode = rds.UserCode

		sql := `update dbo.BCReturnDepSpecial set DocDate=?,LastEditorCode=?,LastEditDateT=?,EarnestNo=?,ArCode=?,DepartCode=?,IsConfirm=?,CreditDay=?,SaleCode=?,DueDate=?,MyDescription=?,TotalAmount=?,BeforeTaxAmount=?,NetAmount=?,SumOfWTax=?,ExcessAmount1=?,ExcessAmount2=?,BillBalance=?,CurrencyCode=?,ExchangeRate=?,PettyCashAmount=?,SumCashAmount=?,SumChqAmount=?,SumCreditAmount=?,SumBankAmount=?,EarnestAmount=?,EarnestBalance=?,OtherIncome=?,OtherExpense=?,GLFormat=?,IsCancel=?,IsReturnMoney=?,AllocateCode=?,ProjectCode=?,BillGroup=?,RecurName) where docno = ?`
		_, err := db.Exec(sql, rds.DocDate,rds.LastEditorCode,rds.LastEditDateT,rds.EarnestNo,rds.ArCode,rds.DepartCode,rds.IsConfirm,rds.CreditDay,rds.SaleCode,rds.DueDate,rds.MyDescription,rds.TotalAmount,rds.BeforeTaxAmount,rds.NetAmount,rds.SumOfWTax,rds.ExcessAmount1,rds.ExcessAmount2,rds.BillBalance,rds.CurrencyCode,rds.ExchangeRate,rds.PettyCashAmount,rds.SumCashAmount,rds.SumChqAmount,rds.SumCreditAmount,rds.SumBankAmount,rds.EarnestAmount,rds.EarnestBalance,rds.OtherIncome,rds.OtherExpense,rds.GLFormat,rds.IsCancel,rds.IsReturnMoney,rds.AllocateCode,rds.ProjectCode,rds.BillGroup,rds.RecurName,rds.DocNo)
		if err != nil {
			return err
		}
	}

	return nil
}
