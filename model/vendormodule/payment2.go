package model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	m "github.com/loukmho/bcaccount_api/model"
	"errors"
)

type Payment2 struct {
	SaveFrom        int         `json:"save_from" db:"SaveFrom"`
	Source          int         `json:"source" db:"Source"`
	DocNo           string      `json:"doc_no" db:"DocNo"`
	Pmt2InPutTax
	DocDate         string      `json:"doc_date" db:"DocDate"`
	ApCode          string      `json:"ap_code" db:"ApCode"`
	CreatorCode     string      `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string      `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string      `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string      `json:"last_edit_date_t" db:"LastEditDateT"`
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
	SumHomeAmount1  float64     `json:"sum_home_amount_1" db:"SumHomeAmount1"`
	SumHomeAmount2  float64     `json:"sum_home_amount_2" db:"SumHomeAmount1"`
	CurrencyCode    string      `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate    float64     `json:"exchange_rate" db:"ExchangeRate"`
	IsCompleteSave  int         `json:"is_complete_save" db:"IsCompleteSave"`
	IsConfirm       int         `json:"is_confirm" db:"IsConfirm"`
	RecurName       string      `json:"recur_name" db:"RecurName"`
	ConfirmCode     string      `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string      `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string      `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string      `json:"cancel_date_time" db:"CancelDateTime"`
	UserCode        string      `json:"user_code" db:"UserCode"`
	Subs            []*Pmt2Item `json:"subs"`
}

type Pmt2InPutTax struct {
	TaxNo    string `json:"tax_no" db:"TaxNo"`
	TaxDate  string `json:"tax_date" db:"TaxDate"`
	BookCode string `json:"book_code" db:"BookCode"`
}

type Pmt2Item struct {
	InvoiceDate   string  `json:"invoice_date" db:"InvoiceDate"`
	InvoiceNo     string  `json:"invoice_no" db:"InvoiceNo"`
	InvBalance    float64 `json:"inv_balance" db:"InvBalance"`
	PayAmount     float64 `json:"pay_amount" db:"PayAmount"`
	MyDescription string  `json:"my_description" db:"MyDescription"`
	BillType      int     `json:"bill_type" db:"BillType"`
	PaybillNo     string  `json:"paybill_no" db:"PaybillNo"`
	IsCancel      int     `json:"is_cancel" db:"IsCancel"`
	LineNumber    int     `json:"line_number" db:"LineNumber"`
	HomeAmount1   float64 `json:"home_amount_1" db:"HomeAmount1"`
	HomeAmount2   float64 `json:"home_amount_2" db:"HomeAmount2"`
}

func (pmt2 *Payment2) InsertAndEditPayment2(db *sqlx.DB) error {
	var check_exist int
	var count_item int

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")
	vTaxRate := def.TaxRateDefault

	fmt.Println("TaxRate = ", def.TaxRateDefault)

	for _, sub_item := range pmt2.Subs {
		if (sub_item.InvoiceNo != "") {
			count_item = count_item + 1
		}
	}

	sqlexist := `select count(docno) as check_exist from dbo.bcpayment2 where docno = ?`
	err := db.Get(&check_exist, sqlexist, pmt2.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	switch {
	case pmt2.DocNo == "":
		return errors.New("Docno is null")
	case pmt2.ApCode == "":
		return errors.New("Apcode is null")
	case pmt2.DocDate == "":
		return errors.New("Docdate is null")
	case pmt2.IsCancel == 1:
		return errors.New("Docno is cancel")
	case pmt2.IsConfirm == 1:
		return errors.New("Docno is confirm")
	case count_item == 0:
		return errors.New("Docno is not have item")
	}

	if pmt2.ExchangeRate == 0 {
		pmt2.ExchangeRate = def.ExchangeRateDefault
	}

	if pmt2.SaveFrom == 0 {
		pmt2.SaveFrom = def.ApInvoiceSaveFrom
	}

	if (pmt2.BookCode == "") {
		pmt2.BookCode = def.Payment2BookCode
	}

	if (pmt2.Source == 0) {
		pmt2.Source = def.Payment2Source
	}

	pmt2.GLFormat = def.Payment2GLFormat

	if (pmt2.TaxNo == "") {
		pmt2.TaxNo = pmt2.DocNo
	}


	fmt.Println("UserCode = ", pmt2.UserCode)

	if (pmt2.IsCompleteSave == 0) {
		pmt2.IsCompleteSave = 1
	}

	if (check_exist == 0) {
		//Insert//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		pmt2.CreatorCode = pmt2.UserCode
		pmt2.NetAmount = pmt2.TotalAmount

		sql := `Insert into dbo.BCPayment2(DocNo,TaxNo,DocDate,CreatorCode,CreateDateTime,ApCode,DepartCode,SumOfInvoice,SumOfDebitNote,SumOfCreditNote,BeforeTaxAmount,TaxAmount,TotalAmount,DiscountAmount,NetAmount,GLFormat,MyDescription,AllocateCode,ProjectCode,BillGroup,SumHomeAmount1,SumHomeAmount2, CurrencyCode, ExchangeRate,IsCompleteSave,IsConfirm,RecurName,IsCancel) values(?,?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sql, pmt2.DocNo,pmt2.TaxNo,pmt2.DocDate,pmt2.CreatorCode,pmt2.ApCode,pmt2.DepartCode,pmt2.SumOfInvoice,pmt2.SumOfDebitNote,pmt2.SumOfCreditNote,pmt2.BeforeTaxAmount,pmt2.TaxAmount,pmt2.TotalAmount,pmt2.DiscountAmount,pmt2.NetAmount,pmt2.GLFormat,pmt2.MyDescription,pmt2.AllocateCode,pmt2.ProjectCode,pmt2.BillGroup,pmt2.SumHomeAmount1,pmt2.SumHomeAmount2,pmt2.CurrencyCode,pmt2.ExchangeRate,pmt2.IsCompleteSave,pmt2.IsConfirm,pmt2.RecurName,pmt2.IsCancel)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}

	} else {

		//Update/////////////////////////////////////////////////////////////////////////////////////////////////////////
		pmt2.LastEditorCode = pmt2.UserCode
		pmt2.NetAmount = pmt2.TotalAmount

		sql := `update dbo.BCPayment2 set TaxNo=?,DocDate=?,LastEditorCode=?,LastEditDateT=getdate(),ApCode=?,DepartCode=?,SumOfInvoice=?,SumOfDebitNote=?,SumOfCreditNote=?,BeforeTaxAmount=?,TaxAmount=?,TotalAmount=?,DiscountAmount=?,NetAmount=?,GLFormat=?,MyDescription=?,AllocateCode=?,ProjectCode=?,BillGroup=?,SumHomeAmount1=?,SumHomeAmount2=?,CurrencyCode=?,ExchangeRate=?,IsCompleteSave=?,IsConfirm=?,RecurName=?,IsCancel=? where docno = ?`
		_, err := db.Exec(sql, pmt2.TaxNo,pmt2.DocDate,pmt2.CreatorCode,pmt2.ApCode,pmt2.DepartCode,pmt2.SumOfInvoice,pmt2.SumOfDebitNote,pmt2.SumOfCreditNote,pmt2.BeforeTaxAmount,pmt2.TaxAmount,pmt2.TotalAmount,pmt2.DiscountAmount,pmt2.NetAmount,pmt2.GLFormat,pmt2.MyDescription,pmt2.AllocateCode,pmt2.ProjectCode,pmt2.BillGroup,pmt2.SumHomeAmount1,pmt2.SumHomeAmount2,pmt2.CurrencyCode,pmt2.ExchangeRate,pmt2.IsCompleteSave,pmt2.IsConfirm,pmt2.RecurName,pmt2.IsCancel, pmt2.DocNo)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}
	}

	sql_del_sub := `delete dbo.bcpaymentsub2 where docno = ?`
	_, err = db.Exec(sql_del_sub, pmt2.DocNo)
	if err != nil {
		return err
	}

	var vLineNumber int

	for _, sub := range pmt2.Subs {
		fmt.Println("ItemSub")
		sub.LineNumber = vLineNumber
		sub.IsCancel = 0

		sqlsub := ` insert into dbo.bcpaymentsub2(DocNo,DocDate,ApCode,InvoiceDate,InvoiceNo,InvBalance,PayAmount,MyDescription,BillType,PaybillNo,IsCancel,LineNumber,HomeAmount1,HomeAmount2) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlsub, pmt2.DocNo, pmt2.DocDate, pmt2.ApCode, sub.InvoiceDate, sub.InvoiceNo, sub.InvBalance, sub.PayAmount, sub.MyDescription, sub.BillType, sub.PaybillNo, sub.IsCancel, sub.LineNumber, sub.HomeAmount1, sub.HomeAmount2)
		fmt.Println("sqlsub = ", sqlsub)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		vLineNumber = vLineNumber + 1
	}

	sqldel := `delete dbo.BCInputTax where docno = ?`
	_, err = db.Exec(sqldel, pmt2.DocNo)
	if err != nil {
		return err
	}

	sqltax := `insert into dbo.BCInputTax(SaveFrom,DocNo,BookCode,Source,DocDate,TaxDate,TaxNo,ApCode,ShortTaxDesc,TaxRate,Process,BeforeTaxAmount,TaxAmount,CreatorCode,CreateDateTime) values(?,?,?,?,?,?,?,?,'ซื้อสินค้า',?,1,?,?,?,getdate())`
	_, err = db.Exec(sqltax, pmt2.SaveFrom, pmt2.DocNo, pmt2.BookCode, pmt2.Source, pmt2.DocDate, pmt2.DocDate, pmt2.TaxNo, pmt2.ApCode, vTaxRate, pmt2.BeforeTaxAmount, pmt2.TaxAmount, pmt2.UserCode)
	fmt.Println("sqltax = ", sqltax)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (pmt2 *Payment2) SearchPayment2ByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select DocNo, isnull(TaxNo,'') as taxno, DocDate, isnull(ApCode,'') as ApCode, isnull(CreatorCode,'') as CreatorCode, isnull(CreateDateTime,'') as CreateDateTime, isnull(LastEditorCode,'') as LastEditorCode, isnull(LastEditDateT,'') as LastEditDateT, isnull(DepartCode,'') as DepartCode, SumOfInvoice, SumOfDebitNote, SumOfCreditNote, BeforeTaxAmount, TaxAmount,TotalAmount, DiscountAmount, NetAmount, isnull(GLFormat,'') as GLFormat, IsCancel, isnull(MyDescription,'') as MyDescription, isnull(AllocateCode,'') as AllocateCode, isnull(ProjectCode,'') as ProjectCode, isnull(BillGroup,'') as BillGroup, SumHomeAmount1, SumHomeAmount2, isnull(CurrencyCode,'') as CurrencyCode, ExchangeRate, IsCompleteSave, IsConfirm, isnull(RecurName,'') as RecurName, isnull(ConfirmCode,'') as ConfirmCode, isnull(ConfirmDateTime,'') as ConfirmDateTime, isnull(CancelCode,'') as CancelCode, isnull(CancelDateTime,'') as CancelDateTime from dbo.BCPayment2 a  with (nolock) where docno = ? `
	fmt.Println("sql=", sql)
	err := db.Get(pmt2, sql, docno)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	sqlsub := `set dateformat dmy    select isnull(InvoiceDate,'') as InvoiceDate, isnull(InvoiceNo,'') as InvoiceNo, InvBalance, PayAmount, isnull(MyDescription,'') as MyDescription, BillType, isnull(PaybillNo,'') as PaybillNo, IsCancel, LineNumber, HomeAmount1, HomeAmount2 from dbo.BCPaymentsub2  with (nolock)  where docno = ? order by LineNumber`
	fmt.Println("sqlsub=", sqlsub)
	err = db.Select(&pmt2.Subs, sqlsub, docno)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (pmt2 *Payment2) SearchPayment2ByKeyword(db *sqlx.DB, keyword string) (pmt2s []*Payment2, err error) {
	sql := `set dateformat dmy     select DocNo, isnull(TaxNo,'') as taxno, DocDate, isnull(ApCode,'') as ApCode, isnull(CreatorCode,'') as CreatorCode, isnull(CreateDateTime,'') as CreateDateTime, isnull(LastEditorCode,'') as LastEditorCode, isnull(LastEditDateT,'') as LastEditDateT, isnull(DepartCode,'') as DepartCode, SumOfInvoice, SumOfDebitNote, SumOfCreditNote, BeforeTaxAmount, TaxAmount,TotalAmount, DiscountAmount, NetAmount, isnull(GLFormat,'') as GLFormat, IsCancel, isnull(MyDescription,'') as MyDescription, isnull(AllocateCode,'') as AllocateCode, isnull(ProjectCode,'') as ProjectCode, isnull(BillGroup,'') as BillGroup, SumHomeAmount1, SumHomeAmount2, isnull(CurrencyCode,'') as CurrencyCode, ExchangeRate, IsCompleteSave, IsConfirm, isnull(RecurName,'') as RecurName, isnull(ConfirmCode,'') as ConfirmCode, isnull(ConfirmDateTime,'') as ConfirmDateTime, isnull(CancelCode,'') as CancelCode, isnull(CancelDateTime,'') as CancelDateTime from dbo.BCPayment2 where (docno  like '%'+?+'%' or apcode like '%'+?+'%') order by docno `
	fmt.Println("sql=", sql)
	err = db.Select(&pmt2s, sql, keyword, keyword, keyword)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	for _, sub := range pmt2s {
		sqlsub := `set dateformat dmy    select isnull(InvoiceDate,'') as InvoiceDate, isnull(InvoiceNo,'') as InvoiceNo, InvBalance, PayAmount, isnull(MyDescription,'') as MyDescription, BillType, isnull(PaybillNo,'') as PaybillNo, IsCancel, LineNumber, HomeAmount1, HomeAmount2 from dbo.BCPaymentsub2  with (nolock) where docno = ? order by LineNumber`
		fmt.Println("sqlsub=", sqlsub)
		err = db.Select(&sub.Subs, sqlsub, sub.DocNo)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}
	return pmt2s, nil
}
