package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	m "github.com/loukmho/bcaccount_api/model"
	"errors"
	"time"
)

type ApOtherDebt struct {
	SaveFrom        int     `json:"save_from" db:"SaveFrom"`
	Source          int     `json:"source" db:"Source"`
	DocNo           string  `json:"doc_no" db:"DocNo"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	ApCode          string  `json:"ap_code" db:"ApCode"`
	GLBookCode      string  `json:"gl_book_code" db:"GLBookCode"`
	SumofDebit      float64 `json:"sumof_debit" db:"SumofDebit"`
	SumofCredit     float64 `json:"sumof_credit" db:"SumofCredit"`
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	CreditDay       int     `json:"credit_day" db:"CreditDay"`
	DueDate         string  `json:"due_date" db:"DueDate"`
	PayBillDate     string  `json:"pay_bill_date" db:"PayBillDate"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	BillGroup       string  `json:"bill_group" db:"BillGroup"`
	ContactCode     string  `json:"contact_code" db:"ContactCode"`
	NetDebtAmount   float64 `json:"net_debt_amount" db:"NetDebtAmount"`
	BillBalance     float64 `json:"bill_balance" db:"BillBalance"`
	CurrencyCode    string  `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate    float64 `json:"exchange_rate" db:"ExchangeRate"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	AllocateCode    string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string  `json:"project_code" db:"ProjectCode"`
	RecurName       string  `json:"recur_name" db:"RecurName"`
	ConfirmCode     string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string  `json:"cancel_date_time" db:"CancelDateTime"`
	PayBillAmount   float64 `json:"pay_bill_amount" db:"PayBillAmount"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	UserCode        string  `json:"user_code" db:"UserCode"`
	AtdInPutTax
}

type AtdInPutTax struct {
	TaxNo    string `json:"tax_no" db:"TaxNo"`
	TaxDate  string `json:"tax_date" db:"TaxDate"`
	BookCode string `json:"book_code" db:"BookCode"`
	BeforeTaxAmount float64 `json:"before_tax_amount" db:"BeforeTaxAmount"`
	TaxAmount float64 `json:"tax_amount" db:"TaxAmount"`
}

//func (atd *ApOtherDebt) SearchApOtherDebtByDocNo() error {
//	//sql := `DocNo, DocDate, ApCode, GLBookCode, SumofDebit, SumofCredit, DepartCode, CreditDay, DueDate, PayBillDate, IsConfirm, MyDescription, BillGroup, ContactCode,NetDebtAmount, BillBalance, CurrencyCode, ExchangeRate, IsCancel, AllocateCode, ProjectCode, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime, PayBillAmount,CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT`
//	return nil
//}

func (atd *ApOtherDebt) InsertAndEditApOtherDebt(db *sqlx.DB) error {
	var check_exist int

	now := time.Now()

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")
	vTaxRate := def.TaxRateDefault

	fmt.Println("TaxRate = ", def.TaxRateDefault)

	sqlexist := `select count(docno) as check_exist from dbo.BCApOtherDebt where docno = ?`
	err := db.Get(&check_exist, sqlexist, atd.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	switch {
	case atd.DocNo == "":
		return errors.New("Docno is null")
	case atd.ApCode == "":
		return errors.New("Apcode is null")
	case atd.DocDate == "":
		return errors.New("Docdate is null")
	case atd.IsCancel == 1:
		return errors.New("Docno is cancel")
	case atd.IsConfirm == 1:
		return errors.New("Docno is confirm")
	case atd.NetDebtAmount == 0:
		return errors.New("NetDebtAmount = 0")
	case atd.SumofDebit != atd.SumofCredit:
		return errors.New("CreditAmount <> DebitAmount")
	}

	if atd.ExchangeRate == 0 {
		atd.ExchangeRate = def.ExchangeRateDefault
	}

	if atd.SaveFrom == 0 {
		atd.SaveFrom = def.ApInvoiceSaveFrom
	}


	if (atd.BookCode == "") {
		atd.BookCode = def.ApOtherDebtBookCode
	}
	if (atd.Source == 0) {
		atd.Source = def.ApOtherDebtSource
	}

	if (atd.TaxNo == "") {
		atd.TaxNo = atd.DocNo
	}

	if (atd.CreditDay == 0 || atd.DueDate == "") {
		atd.DueDate = atd.DocDate
		atd.PayBillDate = atd.DocDate
	} else {
		atd.DueDate = now.AddDate(0, 0, atd.CreditDay).Format("2006-01-02")
		atd.PayBillDate = now.AddDate(0,0,atd.CreditDay).Format("2006-01-02")
	}

	fmt.Println("UserCode = ", atd.UserCode)

	if (check_exist == 0) {
		atd.CreatorCode = atd.UserCode

		sql := `insert into dbo.BCApOtherDebt(DocNo,DocDate,ApCode,GLBookCode,SumofDebit,SumofCredit,DepartCode,CreditDay,DueDate,PayBillDate,IsConfirm,MyDescription,BillGroup,ContactCode,NetDebtAmount,BillBalance,CurrencyCode,ExchangeRate,IsCancel,AllocateCode,ProjectCode,RecurName,PayBillAmount,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,getdate())`
		db.Exec(sql, atd.DocNo, atd.DocDate, atd.ApCode, atd.GLBookCode, atd.SumofDebit, atd.SumofCredit, atd.DepartCode, atd.CreditDay, atd.DueDate, atd.PayBillDate, atd.IsConfirm, atd.MyDescription, atd.BillGroup, atd.ContactCode, atd.NetDebtAmount, atd.BillBalance, atd.CurrencyCode, atd.ExchangeRate, atd.IsCancel, atd.AllocateCode, atd.ProjectCode, atd.RecurName, atd.PayBillAmount, atd.CreatorCode)

	}else{
		atd.LastEditorCode = atd.UserCode

		sql := `update dbo.BCApOtherDebt set DocDate=?,ApCode=?,GLBookCode=?,SumofDebit=?,SumofCredit=?,DepartCode=?,CreditDay=?,DueDate=?,PayBillDate=?,IsConfirm=?,MyDescription=?,BillGroup=?,ContactCode=?,NetDebtAmount=?,BillBalance=?,CurrencyCode=?,ExchangeRate=?,IsCancel=?,AllocateCode=?,ProjectCode=?,RecurName=?,PayBillAmount=?,LastEditorCode=?,LastEditDateT=getdate() where docno = ?`
		db.Exec(sql, atd.DocDate, atd.ApCode, atd.GLBookCode, atd.SumofDebit, atd.SumofCredit, atd.DepartCode, atd.CreditDay, atd.DueDate, atd.PayBillDate, atd.IsConfirm, atd.MyDescription, atd.BillGroup, atd.ContactCode, atd.NetDebtAmount, atd.BillBalance, atd.CurrencyCode, atd.ExchangeRate, atd.IsCancel, atd.AllocateCode, atd.ProjectCode, atd.RecurName, atd.PayBillAmount, atd.LastEditorCode, atd.DocNo)
	}


	sqldel := `delete dbo.BCInputTax where docno = ?`
	_, err = db.Exec(sqldel, atd.DocNo)
	if err != nil {
		return err
	}

	atd.BeforeTaxAmount, atd.TaxAmount = m.CalcTaxCredit(1,vTaxRate,atd.SumofCredit)

	sqltax := `insert into dbo.BCInputTax(SaveFrom,DocNo,BookCode,Source,DocDate,TaxDate,TaxNo,ApCode,ShortTaxDesc,TaxRate,Process,BeforeTaxAmount,TaxAmount,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?,?,'ซื้อสินค้า',?,1,?,?,?,getdate())`
	_, err = db.Exec(sqltax, atd.SaveFrom, atd.DocNo, atd.BookCode, atd.Source, atd.DocDate, atd.DocDate, atd.TaxNo, atd.ApCode, vTaxRate, atd.BeforeTaxAmount, atd.TaxAmount, atd.UserCode)
	fmt.Println("sqltax = ", sqltax)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (atd *ApOtherDebt) SearchApOtherDebtByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select DocNo, DocDate, isnull(ApCode,'') as ApCode, isnull(GLBookCode,'') as GLBookCode, SumofDebit, SumofCredit, isnull(DepartCode,'') as DepartCode, CreditDay, isnull(DueDate,'') as DueDate, isnull(PayBillDate,'') as PayBillDate, IsConfirm, isnull(MyDescription,'') as MyDescription, isnull(BillGroup,'') as BillGroup, isnull(ContactCode,'') as ContactCode, isnull(NetDebtAmount,'') as NetDebtAmount, BillBalance, isnull(CurrencyCode,'') as CurrencyCode, ExchangeRate, IsCancel, isnull(AllocateCode,'') as AllocateCode, isnull(ProjectCode,'') as ProjectCode, isnull(RecurName,'') as RecurName, isnull(ConfirmCode,'') as ConfirmCode, isnull(ConfirmDateTime,'') as ConfirmDateTime, isnull(CancelCode,'') as CancelCode, isnull(CancelDateTime,'') as CancelDateTime, PayBillAmount,isnull(CreatorCode,'') as CreatorCode, isnull(CreateDateTime,'') as CreateDateTime, isnull(LastEditorCode,'') as LastEditorCode, isnull(LastEditDateT,'') as LastEditDateT from dbo.BCApOtherDebt with (nolock) where docno = ?`
	err := db.Get(&atd, sql, docno)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}

	return nil
}


func (atd *ApOtherDebt) SearchApOtherDebtByKeyword(db *sqlx.DB, keycode string) (atds *[]ApOtherDebt, err error) {
	sql := `set dateformat dmy     select DocNo, DocDate, isnull(ApCode,'') as ApCode, isnull(GLBookCode,'') as GLBookCode, SumofDebit, SumofCredit, isnull(DepartCode,'') as DepartCode, CreditDay, isnull(DueDate,'') as DueDate, isnull(PayBillDate,'') as PayBillDate, IsConfirm, isnull(MyDescription,'') as MyDescription, isnull(BillGroup,'') as BillGroup, isnull(ContactCode,'') as ContactCode, isnull(NetDebtAmount,'') as NetDebtAmount, BillBalance, isnull(CurrencyCode,'') as CurrencyCode, ExchangeRate, IsCancel, isnull(AllocateCode,'') as AllocateCode, isnull(ProjectCode,'') as ProjectCode, isnull(RecurName,'') as RecurName, isnull(ConfirmCode,'') as ConfirmCode, isnull(ConfirmDateTime,'') as ConfirmDateTime, isnull(CancelCode,'') as CancelCode, isnull(CancelDateTime,'') as CancelDateTime, PayBillAmount,isnull(CreatorCode,'') as CreatorCode, isnull(CreateDateTime,'') as CreateDateTime, isnull(LastEditorCode,'') as LastEditorCode, isnull(LastEditDateT,'') as LastEditDateT from dbo.BCApOtherDebt with (nolock) where (docno  like '%'+?+'%' or apcode like '%'+?+'%') order by docno`
	err = db.Select(&atds, sql, keycode, keycode)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}
	return atds, nil
}
