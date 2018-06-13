package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	m "github.com/loukmho/bcaccount_api/model"
	"errors"
)

type Receipt3 struct {
	Source          int         `json:"source" db:"Source"`
	SaveFrom        int         `json:"save_from" db:"SaveFrom"`
	DocNo           string      `json:"doc_no" db:"DocNo"`
	TaxNo           string      `json:"tax_no" db:"TaxNo"`
	DocDate         string      `json:"doc_date" db:"DocDate"`
	ArCode          string      `json:"ar_code" db:"ArCode"`
	CreatorCode     string      `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string      `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string      `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string      `json:"last_edit_date_t" db:"LastEditDateT"`
	SaleCode        string      `json:"sale_code" db:"SaleCode"`
	DepartCode      string      `json:"depart_code" db:"DepartCode"`
	SumOfInvoice    float64     `json:"sum_of_invoice" db:"SumOfInvoice"`
	SumOfDebitNote  float64     `json:"sum_of_debit_note" db:"SumOfDebitNote"`
	SumOfCreditNote float64     `json:"sum_of_credit_note" db:"SumOfCreditNote"`
	BeforeTaxAmount float64     `json:"before_tax_amount" db:"BeforeTaxAmount"`
	TaxAmount       float64     `json:"tax_amount" db:"TaxAmount"`
	TotalAmount     float64     `json:"total_amount" db:"TotalAmount"`
	DiscountAmount  float64     `json:"discount_amount" db:"DiscountAmount"`
	NetAmount       float64     `json:"net_amount" db:"NetAmount"`
	GLFormat        string      `json:"gl_format" db:"GLFormat"`
	IsCancel        int         `json:"is_cancel" db:"IsCancel"`
	MyDescription   string      `json:"my_description" db:"MyDescription"`
	AllocateCode    string      `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string      `json:"project_code" db:"ProjectCode"`
	BillGroup       string      `json:"bill_group" db:"BillGroup"`
	IsCompleteSave  int         `json:"is_complete_save" db:"IsCompleteSave"`
	SumHomeAmount1  float64     `json:"sum_home_amount_1" db:"SumHomeAmount1"`
	SumHomeAmount2  float64     `json:"sum_home_amount_2" db:"SumHomeAmount2"`
	CurrencyCode    string      `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate    float64     `json:"exchange_rate" db:"ExchangeRate"`
	IsConfirm       int         `json:"is_confirm" db:"IsConfirm"`
	RecurName       string      `json:"recur_name" db:"RecurName"`
	ConfirmCode     string      `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string      `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string      `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string      `json:"cancel_date_time" db:"CancelDateTime"`
	UserCode        string      `json:"user_code"`
	Subs            []*rcp3Item `json:"subs"`
}

type rcp3Item struct {
	InvoiceDate   string  `json:"invoice_date" db:"InvoiceDate"`
	InvoiceNo     string  `json:"invoice_no" db:"InvoiceNo"`
	InvBalance    float64 `json:"inv_balance" db:"InvBalance"`
	PayAmount     float64 `json:"pay_amount" db:"PayAmount"`
	MyDescription string  `json:"my_description" db:"MyDescription"`
	IsCancel      int     `json:"is_cancel" db:"IsCancel"`
	LineNumber    int     `json:"line_number" db:"LineNumber"`
	BillType      int     `json:"bill_type" db:"BillType"`
	RefNo         string  `json:"ref_no" db:"RefNo"`
	RefType       int     `json:"ref_type" db:"RefType"`
	HomeAmount1   float64 `json:"home_amount_1" db:"HomeAmount1"`
	HomeAmount2   float64 `json:"home_amount_2" db:"HomeAmount2"`
}

func (rct3 *Receipt3) InsertAndEditReceipt3(db *sqlx.DB) error {
	var check_exist int
	var count_item int

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")

	fmt.Println("TaxRate = ", def.TaxRateDefault)

	for _, Subs := range rct3.Subs {
		if (Subs.InvoiceNo != "") {
			count_item = count_item + 1
		}
	}

	sqlexist := `select count(docno) as check_exist from dbo.bcreceipt3 where docno = ?`
	err := db.Get(&check_exist, sqlexist, rct3.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil
	}

	switch {
	case rct3.DocNo == "":
		return errors.New("Docno is null")
	case rct3.ArCode == "":
		return errors.New("Arcode is null")
	case rct3.DocDate == "":
		return errors.New("Docdate is null")
	case rct3.IsCancel == 1:
		return errors.New("Docno is cancel")
	case rct3.IsConfirm == 1:
		return errors.New("Docno is confirm")
	case count_item == 0:
		return errors.New("Docno is not have item")
	case rct3.TotalAmount == 0:
		return errors.New("TotalAmount = 0")
	}

	if rct3.ExchangeRate == 0 {
		rct3.ExchangeRate = def.ExchangeRateDefault
	}

	if rct3.SaveFrom == 0 {
		rct3.SaveFrom = def.Receipt3SaveFrom
	}

	if (rct3.Source == 0) {
		rct3.Source = def.Receipt1Source
	}

	if (rct3.GLFormat == "") {
		rct3.GLFormat = def.Receipt3GLFormat
	}

	if (rct3.TaxNo == "") {
		rct3.TaxNo = rct3.DocNo
	}

	fmt.Println("UserCode = ", rct3.UserCode)

	if (rct3.IsCompleteSave == 0) {
		rct3.IsCompleteSave = 1
	}

	if (check_exist == 0) {
		//Insert//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		rct3.CreatorCode = rct3.UserCode
		rct3.NetAmount = rct3.TotalAmount

		sql := `insert into dbo.BCReceipt3(DocNo,TaxNo,DocDate,ArCode,CreatorCode,CreateDateTime,SaleCode,DepartCode,SumOfInvoice,SumOfDebitNote,SumOfCreditNote,BeforeTaxAmount,TaxAmount,TotalAmount,DiscountAmount,NetAmount,GLFormat,IsCancel,MyDescription,AllocateCode,ProjectCode,BillGroup,IsCompleteSave,SumHomeAmount1,SumHomeAmount2,CurrencyCode,ExchangeRate,IsConfirm,RecurName) values(?,?,?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sql, rct3.DocNo, rct3.TaxNo, rct3.DocDate, rct3.ArCode, rct3.CreatorCode, rct3.SaleCode, rct3.DepartCode, rct3.SumOfInvoice, rct3.SumOfDebitNote, rct3.SumOfCreditNote, rct3.BeforeTaxAmount, rct3.TaxAmount, rct3.TotalAmount, rct3.DiscountAmount, rct3.NetAmount, rct3.GLFormat, rct3.IsCancel, rct3.MyDescription, rct3.AllocateCode, rct3.ProjectCode, rct3.BillGroup, rct3.IsCompleteSave, rct3.SumHomeAmount1, rct3.SumHomeAmount2, rct3.CurrencyCode, rct3.ExchangeRate, rct3.IsConfirm, rct3.RecurName)
		fmt.Println("sql =", sql, rct3.DocNo, rct3.DocDate, rct3.TaxNo, rct3.ArCode, rct3.SaleCode)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}

	} else {

		//Update/////////////////////////////////////////////////////////////////////////////////////////////////////////
		rct3.LastEditorCode = rct3.UserCode
		rct3.NetAmount = rct3.TotalAmount

		sql := `update dbo.BCReceipt3 set TaxNo=?,DocDate=?,ArCode=?,LastEditorCode=?,LastEditDateT=getdate(),SaleCode=?,DepartCode=?,SumOfInvoice=?,SumOfDebitNote=?,SumOfCreditNote=?,BeforeTaxAmount=?,TaxAmount=?,TotalAmount=?,DiscountAmount=?,NetAmount=?,GLFormat=?,IsCancel=?,MyDescription=?,AllocateCode=?,ProjectCode=?,BillGroup=?,IsCompleteSave=?,SumHomeAmount1=?,SumHomeAmount2=?,CurrencyCode=?,ExchangeRate=?,IsConfirm=?,RecurName=? where docno = ?`
		_, err := db.Exec(sql, rct3.TaxNo, rct3.DocDate, rct3.ArCode, rct3.LastEditorCode, rct3.SaleCode, rct3.DepartCode, rct3.SumOfInvoice, rct3.SumOfDebitNote, rct3.SumOfCreditNote, rct3.BeforeTaxAmount, rct3.TaxAmount, rct3.TotalAmount, rct3.DiscountAmount, rct3.NetAmount, rct3.GLFormat, rct3.IsCancel, rct3.MyDescription, rct3.AllocateCode, rct3.ProjectCode, rct3.BillGroup, rct3.IsCompleteSave, rct3.SumHomeAmount1, rct3.SumHomeAmount2, rct3.CurrencyCode, rct3.ExchangeRate, rct3.IsConfirm, rct3.RecurName, rct3.DocNo)
		fmt.Println("sql =", sql, rct3.DocNo, rct3.DocDate, rct3.TaxNo, rct3.ArCode, rct3.SaleCode)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}
	}

	sql_del_sub := `delete dbo.BCReceiptSub3 where docno = ?`
	_, err = db.Exec(sql_del_sub, rct3.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}

	var vLineNumber int

	for _, sub := range rct3.Subs {
		fmt.Println("ItemSub")
		sub.LineNumber = vLineNumber
		sub.IsCancel = 0
		sub.HomeAmount2 = sub.HomeAmount2

		sqlsub := ` insert into dbo.BCReceiptSub3(DocNo,DocDate,ArCode,SaleCode,InvoiceDate,InvoiceNo,InvBalance,PayAmount,MyDescription,IsCancel,LineNumber,BillType,RefNo,RefType,HomeAmount1,HomeAmount2) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlsub, rct3.DocNo, rct3.DocDate, rct3.ArCode, rct3.SaleCode, sub.InvoiceDate, sub.InvoiceNo, sub.InvBalance, sub.PayAmount, sub.MyDescription, sub.IsCancel, sub.LineNumber, sub.BillType, sub.RefNo, sub.RefType, sub.HomeAmount1, sub.HomeAmount2)
		fmt.Println("sqltax = ", sqlsub)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}

		vLineNumber = vLineNumber + 1
	}

	return nil

}

//func (rct3 *Receipt3) SearchReceipt3ByDocNo() error {
//	//sql:=`DocNo, TaxNo, DocDate, ArCode,  CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, SaleCode, DepartCode, SumOfInvoice, SumOfDebitNote, SumOfCreditNote, BeforeTaxAmount,TaxAmount, TotalAmount, DiscountAmount, NetAmount, GLFormat,  IsCancel, MyDescription, AllocateCode, ProjectCode, BillGroup,  IsCompleteSave, SumHomeAmount1,SumHomeAmount2, CurrencyCode, ExchangeRate, IsConfirm, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime`
//	//sqlsub := `DocNo, DocDate,ArCode, SaleCode, InvoiceDate, InvoiceNo, InvBalance, PayAmount, MyDescription, IsCancel, LineNumber, BillType, RefNo, RefType,HomeAmount1, HomeAmount2`
//	return nil
//}

func (rct3 *Receipt3) SearchReceipt3ByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     set dateformat dmy     select DocNo,isnull(TaxNo,'') as TaxNo,DocDate,isnull(ArCode,'') as ArCode,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(SaleCode,'') as SaleCode,isnull(DepartCode,'') as DepartCode,SumOfInvoice,SumOfDebitNote,SumOfCreditNote,BeforeTaxAmount,TaxAmount,TotalAmount,DiscountAmount,NetAmount,isnull(GLFormat,'') as GLFormat,IsCancel,isnull(MyDescription,'') as MyDescription,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(BillGroup,'') as BillGroup,IsCompleteSave,SumHomeAmount1,SumHomeAmount2,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,IsConfirm,isnull(RecurName,'') as RecurName,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime from dbo.bcreceipt3 with (nolock) with (nolock) where docno = ?`
	err := db.Get(rct3, sql,docno)
	if err != nil {
		fmt.Println("Search ", err.Error())
		return err
	}

	sqlsub := `set dateformat dmy     select isnull(InvoiceDate,'') as InvoiceDate, isnull(InvoiceNo,'') as InvoiceNo, InvBalance, PayAmount, isnull(MyDescription,'') as MyDescription, IsCancel, LineNumber, BillType, isnull(RefNo,'') as RefNo, RefType,  HomeAmount1, HomeAmount2 from dbo.BCReceiptSub1 with (nolock) where docno = ? order by linenumber`
	err = db.Select(rct3.Subs, sqlsub, docno)
	if err != nil {
		fmt.Println("Search Sub ", err.Error())
		return err
	}
	//sql := `DocNo, TaxNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, ArCode, SaleCode, DepartCode, SumOfInvoice, SumOfDebitNote, SumOfCreditNote, BeforeTaxAmount,TaxAmount, TotalAmount, DiscountAmount, NetAmount, SumCashAmount, SumChqAmount, SumCreditAmount, ChargeAmount, ChangeAmount, SumBankAmount, SumOfWTax, GLFormat, IsCancel,MyDescription, AllocateCode, ProjectCode, BillGroup, IsCompleteSave, SumHomeAmount1, SumHomeAmount2, DebtLimitReturn, CurrencyCode, ExchangeRate, OtherIncome,OtherExpense, ExcessAmount1, ExcessAmount2, IsConfirm, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime`
	//sqlsub:= `DocNo, DocDate, ArCode, SaleCode,  InvoiceDate, InvoiceNo, InvBalance, PayAmount, DebtLimitReturn,MyDescription,  IsCancel, LineNumber, BillType, RefNo, RefType,  HomeAmount1, HomeAmount2,`
	return nil
}


func (rct3 *Receipt3) SearchReceipt3ByKeyword(db *sqlx.DB, keyword string) (rct3s []*Receipt3, err error) {
	sql := `set dateformat dmy     set dateformat dmy     select DocNo,isnull(TaxNo,'') as TaxNo,DocDate,isnull(ArCode,'') as ArCode,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(SaleCode,'') as SaleCode,isnull(DepartCode,'') as DepartCode,SumOfInvoice,SumOfDebitNote,SumOfCreditNote,BeforeTaxAmount,TaxAmount,TotalAmount,DiscountAmount,NetAmount,isnull(GLFormat,'') as GLFormat,IsCancel,isnull(MyDescription,'') as MyDescription,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(BillGroup,'') as BillGroup,IsCompleteSave,SumHomeAmount1,SumHomeAmount2,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,IsConfirm,isnull(RecurName,'') as RecurName,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime from dbo.bcreceipt3 with (nolock) where (docno  like '%'+?+'%' or arcode like '%'+?+'%' or salecode like '%'+?+'%') order by docno`
	err = db.Select(&rct3s, sql, keyword, keyword, keyword)
	if err != nil {
		fmt.Println("Search ", err.Error())
		return nil, err
	}

	for _, sub := range rct3s {
		sqlsub := `set dateformat dmy     select isnull(InvoiceDate,'') as InvoiceDate, isnull(InvoiceNo,'') as InvoiceNo, InvBalance, PayAmount, isnull(MyDescription,'') as MyDescription, IsCancel, LineNumber, BillType, isnull(RefNo,'') as RefNo, RefType,  HomeAmount1, HomeAmount2 from dbo.BCReceiptSub1 with (nolock) where docno = ? order by linenumber`
		err = db.Select(&rct3.Subs, sqlsub, sub.DocNo)
		if err != nil {
			fmt.Println("Search Sub ", err.Error())
			return nil, err
		}
	}
	//sql := `DocNo, TaxNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, ArCode, SaleCode, DepartCode, SumOfInvoice, SumOfDebitNote, SumOfCreditNote, BeforeTaxAmount,TaxAmount, TotalAmount, DiscountAmount, NetAmount, SumCashAmount, SumChqAmount, SumCreditAmount, ChargeAmount, ChangeAmount, SumBankAmount, SumOfWTax, GLFormat, IsCancel,MyDescription, AllocateCode, ProjectCode, BillGroup, IsCompleteSave, SumHomeAmount1, SumHomeAmount2, DebtLimitReturn, CurrencyCode, ExchangeRate, OtherIncome,OtherExpense, ExcessAmount1, ExcessAmount2, IsConfirm, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime`
	//sqlsub:= `DocNo, DocDate, ArCode, SaleCode,  InvoiceDate, InvoiceNo, InvBalance, PayAmount, DebtLimitReturn,MyDescription,  IsCancel, LineNumber, BillType, RefNo, RefType,  HomeAmount1, HomeAmount2,`
	return rct3s, nil
}
