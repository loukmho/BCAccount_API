package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"errors"
	m "github.com/loukmho/bcaccount_api/model"
	//"time"
)

type CreditNote struct {
	SaveFrom        int                  `json:"save_from" db"SaveFrom"`
	DocNo           string               `json:"doc_no" db:"DocNo"`
	TaxNo           string               `json:"tax_no" db:"TaxNo"`
	TaxDate         string               `json:"tax_date" db:"TaxDate"`
	CreatorCode     string               `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string               `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string               `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string               `json:"last_edit_date_t" db:"LastEditDateT"`
	DocDate         string               `json:"doc_date" db:"DocDate"`
	DueDate         string               `json:"due_date" db:"DueDate"`
	TaxType         int                  `json:"tax_type" db:"TaxType"`
	Source          int                  `json:"source" db:"Source"`
	CrdCustomer
	CrdSaleMan
	DepartCode      string               `json:"depart_code" db:"DepartCode"`
	CashierCode     string               `json:"cashier_code" db:"CashierCode"`
	TaxRate         float64              `json:"tax_rate" db:"TaxRate"`
	IsConfirm       int                  `json:"is_confirm" db:"IsConfirm"`
	MyDescription   string               `json:"my_description" db:"MyDescription"`
	SumOfItemAmount float64              `json:"sum_of_item_amount" db:"SumOfItemAmount"`
	SumOldAmount    float64              `json:"sum_old_amount" db:"SumOldAmount"`
	SumTrueAmount   float64              `json:"sum_true_amount" db:"SumTrueAmount"`
	SumofDiffAmount float64              `json:"sumof_diff_amount" db:"SumofDiffAmount"`
	DiscountWord    string               `json:"discount_word" db:"DiscountWord"`
	DiscountAmount  float64              `json:"discount_amount" db:"DiscountAmount"`
	SumofBeforeTax  float64              `json:"sumof_before_tax" db:"SumofBeforeTax"`
	SumOfTaxAmount  float64              `json:"sum_of_tax_amount" db:"SumOfTaxAmount"`
	SumOfTotalTax   float64              `json:"sum_of_total_tax" db:"SumOfTotalTax"`
	SumOfZeroTax    float64              `json:"sum_of_zero_tax" db:"SumOfZeroTax"`
	SumOfExceptTax  float64              `json:"sum_of_except_tax" db:"SumOfExceptTax"`
	SumOfWTax       float64              `json:"sum_of_w_tax" db:"SumOfWTax"`
	NetDebtAmount   float64              `json:"net_debt_amount" db:"NetDebtAmount"`
	BillBalance     float64              `json:"bill_balance" db:"BillBalance"`
	CurrencyCode    string               `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate    float64              `json:"exchange_rate" db:"ExchangeRate"`
	GLFormat        string               `json:"gl_format" db:"GLFormat"`
	IsCancel        int                  `json:"is_cancel" db:"IsCancel"`
	IsCompleteSave  int                  `json:"is_complete_save" db:"IsCompleteSave"`
	ReturnMoney     int                  `json:"return_money" db:"ReturnMoney"`
	ReturnStatus    int                  `json:"return_status" db:"ReturnStatus"`
	ReturnCash      int                  `json:"return_cash" db:"ReturnCash"`
	OtherIncome     float64              `json:"other_income" db:"OtherIncome"`
	OtherExpense    float64              `json:"other_expense" db:"OtherExpense"`
	ExcessAmount1   float64              `json:"excess_amount_1" db:"ExcessAmount1"`
	ExcessAmount2   float64              `json:"excess_amount_2" db:"ExcessAmount2"`
	SumCashAmount   float64              `json:"sum_cash_amount" db:"SumCashAmount"`
	SumChqAmount    float64              `json:"sum_chq_amount" db:"SumChqAmount"`
	SumCreditAmount float64              `json:"sum_credit_amount" db:"SumCreditAmount"`
	SumBankAmount   float64              `json:"sum_bank_amount" db:"SumBankAmount"`
	ChargeAmount    float64              `json:"charge_amount" db:"ChargeAmount"`
	ChangeAmount    float64              `json:"change_amount" db:"ChangeAmount"`
	CauseType       int                  `json:"cause_type" db:"CauseType"`
	PayBillStatus   int                  `json:"pay_bill_status" db:"PayBillStatus"`
	IsCNDeposit     int                  `json:"is_cn_deposit" db:"IsCNDeposit"`
	IsPos           int                  `json:"is_pos" db:"IsPos"`
	PosDocNo        string               `json:"pos_doc_no" db:"PosDocNo"`
	CauseCode       string               `json:"cause_code" db:"CauseCode"`
	AllocateCode    string               `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string               `json:"project_code" db:"ProjectCode"`
	BillGroup       string               `json:"bill_group" db:"BillGroup"`
	RecurName       string               `json:"recur_name" db:"RecurName"`
	ConfirmCode     string               `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string               `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string               `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string               `json:"cancel_date_time" db:"CancelDateTime"`
	PayBillAmount   float64              `json:"pay_bill_amount" db:"PayBillAmount"`
	BillTemporary   float64              `json:"bill_temporary" db:"BillTemporary"`
	UserCode        string               `json:"user_code" db:"UserCode"`
	ListCrdRecMoney
	Subs            []*CrdItem           `json:"subs"`
	Cdcs            []*ListCrdCreditCard `json:"cdcs"`
	Chqs            []*ListCrdChqIn      `json:"chqs"`
}

type CrdCustomer struct {
	ArCode string `json:"ar_code" db:"ArCode"`
	ArName string `json:"ar_name" db:"ArName"`
}

type CrdSaleMan struct {
	SaleCode string `json:"sale_code" db:"SaleCode"`
	SaleName string `json:"sale_name" db:"SaleName"`
}

type CrdItem struct {
	MyType         int     `json:"my_type" db:"MyType"`
	ItemCode       string  `json:"item_code" db:"ItemCode"`
	MyDescription  string  `json:"my_description" db:"MyDescription"`
	ItemName       string  `json:"item_name" db:"ItemName"`
	ItemType       int     `json:"item_type" db:"ItemType"`
	WHCode         string  `json:"wh_code" db:"WHCode"`
	ShelfCode      string  `json:"shelf_code" db:"ShelfCode"`
	DiscQty        float64 `json:"disc_qty" db:"DiscQty"`
	TempQty        float64 `json:"temp_qty" db:"TempQty"`
	BillQty        float64 `json:"bill_qty" db:"BillQty"`
	Price          float64 `json:"price" db:"Price"`
	DiscountWord   string  `json:"discount_word" db:"DiscountWord"`
	DiscountAmount float64 `json:"discount_amount" db:"DiscountAmount"`
	Amount         float64 `json:"amount" db:"Amount"`
	NetAmount      float64 `json:"net_amount" db:"NetAmount"`
	HomeAmount     float64 `json:"home_amount" db:"HomeAmount"`
	SumOfCost      float64 `json:"sum_of_cost" db:"SumOfCost"`
	UnitCode       string  `json:"unit_code" db:"UnitCode"`
	InvoiceNo      string  `json:"invoice_no" db:"InvoiceNo"`
	IsPos          int     `json:"is_pos" db:"IsPos"`
	IsCancel       int     `json:"is_cancel" db:"IsCancel"`
	LineNumber     int     `json:"line_number" db:"LineNumber"`
	RefLineNumber  int     `json:"ref_line_number" db:"RefLineNumber"`
	BarCode        string  `json:"bar_code" db:"BarCode"`
	AVERAGECOST    float64 `json:"averagecost" db:"AVERAGECOST"`
	TaxRate        float64 `json:"tax_rate" db:"TaxRate"`
	LotNumber      string  `json:"lot_number" db:"LotNumber"`
	PackingRate1   float64 `json:"packing_rate_1" db:"PackingRate1"`
	PackingRate2   float64 `json:"packing_rate_2" db:"PackingRate2"`
}

type ListCrdRecMoney struct {
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

type ListCrdChqIn struct {
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

type ListCrdCreditCard struct {
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

func (crd *CreditNote) InsertAndEditCreditNote(db *sqlx.DB) error {
	var check_exist int
	var count_item int
	//var sum_pay_amount float64

	//	now := time.Now()

	//sum_pay_amount = (crd.SumCashAmount+crd.SumCreditAmount+crd.SumChqAmount+crd.SumBankAmount+crd.OtherExpense)-crd.OtherIncome

	for _, sub_item := range crd.Subs {
		if (sub_item.DiscQty != 0) {
			count_item = count_item + 1
		}
	}

	sqlexist := `select count(docno) as check_exist from dbo.bccreditnote where docno = ?`
	err := db.Get(&check_exist, sqlexist, crd.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println("Docno = ", crd.DocNo)
	switch {
	case crd.DocNo == "":
		return errors.New("Docno is null")
	case crd.ArCode == "":
		return errors.New("Arcode is null")
	case crd.DocDate == "":
		return errors.New("Docdate is null")
	case crd.IsCancel == 1:
		return errors.New("Docno is cancel")
	case crd.IsConfirm == 1:
		return errors.New("Docno is confirm")
	case crd.SumOldAmount == 0:
		return errors.New("Docno not have old bill amount")
	case crd.SumOfTotalTax == 0:
		return errors.New("Docno not have amount return")
	case crd.SumChqAmount != 0 && len(crd.Chqs) == 0:
		return errors.New("Docno not have chq use data")
	case crd.SumCreditAmount != 0 && (crd.CreditRefNo == "" || crd.ConfirmNo == "" || crd.CreditType == "" || crd.BankCode == "" || crd.BankBranchCode == ""):
		return errors.New("credit card not have credittype,confirmno,creditno,bankcode,bankbranch data not complete")
	case crd.SumBankAmount != 0 && crd.BankRefNo == "":
		return errors.New("Bank transfer not have accountcode data not complete")
	case crd.SumOfItemAmount == 0 && count_item == 0:
		return errors.New("Docno not have SumOfItemAmount")
	}

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")

	fmt.Println("TaxRate = ", def.TaxRateDefault)

	if crd.TaxRate == 0 {
		crd.TaxRate = def.TaxRateDefault
	}
	if crd.ExchangeRate == 0 {
		crd.ExchangeRate = def.ExchangeRateDefault
	}

	if crd.SaveFrom == 0 {
		crd.SaveFrom = def.CreditNoteSaveFrom
	}

	crd.SumofBeforeTax, crd.SumOfTaxAmount = m.CalcTaxCredit(crd.TaxType, crd.TaxRate, crd.SumOfTotalTax)
	fmt.Println("Befor,Tax", crd.SumofBeforeTax, crd.SumOfTaxAmount)

	if crd.ReturnMoney == 0 {
		crd.BillBalance = crd.SumOfTotalTax
	} else {
		crd.BillBalance = 0
	}

	crd.NetDebtAmount = crd.SumOfTotalTax
	crd.SumTrueAmount = crd.SumOldAmount - crd.SumOfTotalTax

	if crd.TaxType == 1 {
		crd.PayBillAmount = crd.SumOfTotalTax
		crd.BillTemporary = crd.SumOfTotalTax
	}

	if (crd.Source == 0) {
		crd.Source = def.CreditNoteSource
	}
	if (crd.GLFormat == "") {
		crd.GLFormat = def.CreditNoteGLFormat
	}
	if (crd.TaxNo == "" ) {
		crd.TaxNo = crd.DocNo
	}
	if (crd.TaxDate == "" ) {
		crd.TaxDate = crd.DocDate
	}

	if (crd.SumCreditAmount != 0 && crd.CreditDueDate == "") {
		crd.CreditDueDate = crd.DocDate
	}

	if (crd.SumBankAmount != 0 && crd.TransBankDate == "") {
		crd.TransBankDate = crd.DocDate
	}

	fmt.Println("UserCode = ", crd.UserCode)

	if (crd.IsCompleteSave == 0) {
		crd.IsCompleteSave = 1
	}

	if check_exist == 0 {
		sql := `insert into dbo.bccreditnote(DocNo,TaxNo,TaxDate,CreatorCode,CreateDateTime,DocDate,DueDate,TaxType,ArCode,DepartCode,SaleCode,CashierCode,TaxRate,IsConfirm,MyDescription,SumOfItemAmount,SumOldAmount,SumTrueAmount,SumofDiffAmount,DiscountWord,DiscountAmount,SumofBeforeTax,SumOfTaxAmount,SumOfTotalTax,SumOfZeroTax,SumOfExceptTax,SumOfWTax,NetDebtAmount,BillBalance,CurrencyCode,ExchangeRate,GLFormat,IsCancel,IsCompleteSave,ReturnMoney,ReturnStatus,ReturnCash,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,ChargeAmount,ChangeAmount,CauseType,PayBillStatus,IsCNDeposit,IsPos,PosDocNo,CauseCode,AllocateCode,ProjectCode,BillGroup,RecurName,PayBillAmount,BillTemporary) values(?,?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sql, crd.DocNo, crd.TaxNo, crd.TaxDate, crd.UserCode, crd.DocDate, crd.DueDate, crd.TaxType, crd.ArCode, crd.DepartCode, crd.SaleCode, crd.CashierCode, crd.TaxRate, crd.IsConfirm, crd.MyDescription, crd.SumOfItemAmount, crd.SumOldAmount, crd.SumTrueAmount, crd.SumofDiffAmount, crd.DiscountWord, crd.DiscountAmount, crd.SumofBeforeTax, crd.SumOfTaxAmount, crd.SumOfTotalTax, crd.SumOfZeroTax, crd.SumOfExceptTax, crd.SumOfWTax, crd.NetDebtAmount, crd.BillBalance, crd.CurrencyCode, crd.ExchangeRate, crd.GLFormat, crd.IsCancel, crd.IsCompleteSave, crd.ReturnMoney, crd.ReturnStatus, crd.ReturnCash, crd.OtherIncome, crd.OtherExpense, crd.ExcessAmount1, crd.ExcessAmount2, crd.SumCashAmount, crd.SumChqAmount, crd.SumCreditAmount, crd.SumBankAmount, crd.ChargeAmount, crd.ChangeAmount, crd.CauseType, crd.PayBillStatus, crd.IsCNDeposit, crd.IsPos, crd.PosDocNo, crd.CauseCode, crd.AllocateCode, crd.ProjectCode, crd.BillGroup, crd.RecurName, crd.PayBillAmount, crd.BillTemporary)
		if err != nil {
			fmt.Println("Insert Credit =", err.Error())
			return err
		}
	} else {
		sql := `update dbo.bccreditnote set TaxNo=?,TaxDate=?,LastEditorCode=?,LastEditDateT=getdate(),DocDate=?,DueDate=?,TaxType=?,ArCode=?,DepartCode=?,SaleCode=?,CashierCode=?,TaxRate=?,IsConfirm=?,MyDescription=?,SumOfItemAmount=?,SumOldAmount=?,SumTrueAmount=?,SumofDiffAmount=?,DiscountWord=?,DiscountAmount=?,SumofBeforeTax=?,SumOfTaxAmount=?,SumOfTotalTax=?,SumOfZeroTax=?,SumOfExceptTax=?,SumOfWTax=?,NetDebtAmount=?,BillBalance=?,CurrencyCode=?,ExchangeRate=?,GLFormat=?,IsCancel=?,IsCompleteSave=?,ReturnMoney=?,ReturnStatus=?,ReturnCash=?,OtherIncome=?,OtherExpense=?,ExcessAmount1=?,ExcessAmount2=?,SumCashAmount=?,SumChqAmount=?,SumCreditAmount=?,SumBankAmount=?,ChargeAmount=?,ChangeAmount=?,CauseType=?,PayBillStatus=?,IsCNDeposit=?,IsPos=?,PosDocNo=?,CauseCode=?,AllocateCode=?,ProjectCode=?,BillGroup=?,RecurName=?,PayBillAmount=?,BillTemporary=? where docno = ? `
		_, err = db.Exec(sql, crd.TaxNo, crd.TaxDate, crd.UserCode, crd.DocDate, crd.DueDate, crd.TaxType, crd.ArCode, crd.DepartCode, crd.SaleCode, crd.CashierCode, crd.TaxRate, crd.IsConfirm, crd.MyDescription, crd.SumOfItemAmount, crd.SumOldAmount, crd.SumTrueAmount, crd.SumofDiffAmount, crd.DiscountWord, crd.DiscountAmount, crd.SumofBeforeTax, crd.SumOfTaxAmount, crd.SumOfTotalTax, crd.SumOfZeroTax, crd.SumOfExceptTax, crd.SumOfWTax, crd.NetDebtAmount, crd.BillBalance, crd.CurrencyCode, crd.ExchangeRate, crd.GLFormat, crd.IsCancel, crd.IsCompleteSave, crd.ReturnMoney, crd.ReturnStatus, crd.ReturnCash, crd.OtherIncome, crd.OtherExpense, crd.ExcessAmount1, crd.ExcessAmount2, crd.SumCashAmount, crd.SumChqAmount, crd.SumCreditAmount, crd.SumBankAmount, crd.ChargeAmount, crd.ChangeAmount, crd.CauseType, crd.PayBillStatus, crd.IsCNDeposit, crd.IsPos, crd.PosDocNo, crd.CauseCode, crd.AllocateCode, crd.ProjectCode, crd.BillGroup, crd.RecurName, crd.PayBillAmount, crd.BillTemporary, crd.DocNo)
		if err != nil {
			fmt.Println("Insert Credit =", err.Error())
			return err
		}
	}

	sql_del_sub := `delete dbo.bccreditnotesub where docno = ?`
	_, err = db.Exec(sql_del_sub, crd.DocNo)
	if err != nil {
		return err
	}

	var vLineNumber int

	for _, item := range crd.Subs {
		fmt.Println("ItemSub")
		item.MyType = def.CreditNoteMyType
		item.TempQty = item.DiscQty

		item.LineNumber = vLineNumber

		item.IsCancel = 0
		if (item.PackingRate1 == 0) {
			item.PackingRate1 = 1
		}
		if (item.PackingRate2 == 0) {
			item.PackingRate2 = 1
		}

		item.IsPos = crd.IsPos
		if item.Amount == 0 {
			item.Amount = (item.DiscQty * (item.Price - (item.DiscountAmount / item.DiscQty)))
		}

		switch {
		case crd.TaxType == 0:
			item.HomeAmount = item.Amount
			item.NetAmount = item.Amount
		case crd.TaxType == 1:
			taxamount := m.ToFixed(item.Amount-((item.Amount*100)/(100+float64(crd.TaxRate))), 2)
			beforetaxamount := m.ToFixed(item.Amount-taxamount, 2)
			item.HomeAmount = beforetaxamount
			item.NetAmount = beforetaxamount
		case crd.TaxType == 2:
			item.HomeAmount = item.Amount
			item.NetAmount = item.Amount
		}

		sqlsub := ` insert into dbo.BCCreditNoteSub(MyType, DocNo, TaxNo, TaxType, TaxRate, ItemCode, ItemType, DocDate, ArCode, DepartCode, SaleCode, CashierCode, MyDescription, ItemName, WHCode, ShelfCode, DiscQty, TempQty, BillQty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost, UnitCode, InvoiceNo, IsPos, IsCancel, LineNumber, RefLineNumber, BarCode,AVERAGECOST, LotNumber, PackingRate1, PackingRate2) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		_, err = db.Exec(sqlsub, item.MyType, crd.DocNo, crd.TaxNo, crd.TaxType, crd.TaxRate, item.ItemCode, item.ItemType, crd.DocDate, crd.ArCode, crd.DepartCode, crd.SaleCode, crd.CashierCode, item.MyDescription, item.ItemName, item.WHCode, item.ShelfCode, item.DiscQty, item.TempQty, item.BillQty, item.Price, item.DiscountWord, item.DiscountAmount, item.Amount, item.NetAmount, item.HomeAmount, item.SumOfCost, item.UnitCode, item.InvoiceNo, item.IsPos, item.IsCancel, item.LineNumber, item.RefLineNumber, item.BarCode, item.AVERAGECOST, item.LotNumber, item.PackingRate1, item.PackingRate2)
		fmt.Println("sqltax = ", sqlsub)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		if (item.ItemType  != 1){
			sqlprocess := ` insert into dbo.ProcessStock (ItemCode,ProcessFlag,FlowStatus) values(?, 1, 0)`
			_, err = db.Exec(sqlprocess, item.ItemCode )
			fmt.Println("sqlprocess = ", sqlsub)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		vLineNumber = vLineNumber + 1
	}

	return nil
}

func (crd *CreditNote) SearchCreditNoteByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select a.DocNo,isnull(a.TaxNo,'') as TaxNo,isnull(a.TaxDate,'') as TaxDate,isnull(a.CreatorCode,'') as CreatorCode,isnull(a.CreateDateTime,'') as CreateDateTime,isnull(a.LastEditorCode,'') as LastEditorCode,isnull(a.LastEditDateT,'') as LastEditDateT,a.DocDate,isnull(a.DueDate,'') as DueDate,a.TaxType,a.ArCode,isnull(b.name1,'') as ArName,isnull(a.DepartCode,'') as DepartCode,isnull(a.SaleCode,'') as SaleCode,isnull(c.name,'') as SaleName,isnull(a.CashierCode,'') as CashierCode,a.TaxRate,a.IsConfirm,isnull(a.MyDescription,'') as MyDescription,a.SumOfItemAmount,a.SumOldAmount,a.SumTrueAmount,a.SumofDiffAmount,isnull(a.DiscountWord,'') as DiscountWord,a.DiscountAmount,a.SumofBeforeTax,a.SumOfTaxAmount,a.SumOfTotalTax,a.SumOfZeroTax,a.SumOfExceptTax,a.SumOfWTax,a.NetDebtAmount,a.BillBalance,isnull(a.CurrencyCode,'') as CurrencyCode,a.ExchangeRate,isnull(a.GLFormat,'') as GLFormat,a.IsCancel,a.IsCompleteSave,a.ReturnMoney,a.ReturnStatus,a.ReturnCash,a.OtherIncome,a.OtherExpense,a.ExcessAmount1,a.ExcessAmount2,a.SumCashAmount,a.SumChqAmount,a.SumCreditAmount,a.SumBankAmount,a.ChargeAmount,a.ChangeAmount,a.CauseType,a.PayBillStatus,a.IsCNDeposit,a.IsPos,isnull(a.PosDocNo,'') as PosDocNo,isnull(a.CauseCode,'') as CauseCode,isnull(a.AllocateCode,'') as AllocateCode,isnull(a.ProjectCode,'') as ProjectCode,isnull(a.BillGroup,'') as BillGroup,isnull(a.RecurName,'') as RecurName,isnull(a.ConfirmCode,'') as ConfirmCode,isnull(a.ConfirmDateTime,'') as ConfirmDateTime,isnull(a.CancelCode,'') as CancelCode,isnull(a.CancelDateTime,'') as CancelDateTime,a.PayBillAmount,a.BillTemporary from dbo.BCCreditNote a With (Nolock) left join dbo.bcar b with (nolock) on a.arcode = b.code left join dbo.bcsale c with (nolock) on a.salecode = c.code where a.docno = ?`
	err := db.Get(crd, sql, docno)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//sqlsub := `MyType, DocNo, TaxNo, TaxType, ItemCode, DocDate, ArCode, DepartCode, SaleCode, CashierCode, MyDescription, ItemName, WHCode, ShelfCode, DiscQty, TempQty, BillQty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost, UnitCode, InvoiceNo, ItemType, ExceptTax, IsPos, IsCancel, LineNumber, RefLineNumber, BarCode,AVERAGECOST, LotNumber, PackingRate1, PackingRate2`
	sqlsub := `set dateformat dmy     select a.MyType,a.ItemCode,isnull(a.MyDescription,'') as MyDescription,isnull(a.ItemName,'') as ItemName,isnull(a.WHCode,'') as WHCode,isnull(a.ShelfCode,'') as ShelfCode,a.DiscQty,a.TempQty,a.BillQty,a.Price,isnull(a.DiscountWord,'') as DiscountWord,a.DiscountAmount,a.Amount,a.NetAmount,a.HomeAmount,a.SumOfCost,isnull(a.UnitCode,'') as UnitCode,isnull(a.InvoiceNo,'') as InvoiceNo,a.IsPos,a.IsCancel,a.LineNumber,a.RefLineNumber,isnull(a.BarCode,'') as BarCode,a.AVERAGECOST,isnull(a.LotNumber,'') as LotNumber,isnull(a.PackingRate1,1) as PackingRate1,isnull(a.PackingRate2,1) as PackingRate2 from dbo.BCCreditNoteSub a with (nolock) where a.docno = ?`
	err = db.Select(&crd.Subs, sqlsub, docno)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (crd *CreditNote) SearchCreditNoteByKeyword(db *sqlx.DB, keyword string) (crdList []*CreditNote, err error) {
	sql := `set dateformat dmy     select a.DocNo,isnull(a.TaxNo,'') as TaxNo,isnull(a.TaxDate,'') as TaxDate,isnull(a.CreatorCode,'') as CreatorCode,isnull(a.CreateDateTime,'') as CreateDateTime,isnull(a.LastEditorCode,'') as LastEditorCode,isnull(a.LastEditDateT,'') as LastEditDateT,a.DocDate,isnull(a.DueDate,'') as DueDate,a.TaxType,a.ArCode,isnull(b.name1,'') as ArName,isnull(a.DepartCode,'') as DepartCode,isnull(a.SaleCode,'') as SaleCode,isnull(c.name,'') as SaleName,isnull(a.CashierCode,'') as CashierCode,a.TaxRate,a.IsConfirm,isnull(a.MyDescription,'') as MyDescription,a.SumOfItemAmount,a.SumOldAmount,a.SumTrueAmount,a.SumofDiffAmount,isnull(a.DiscountWord,'') as DiscountWord,a.DiscountAmount,a.SumofBeforeTax,a.SumOfTaxAmount,a.SumOfTotalTax,a.SumOfZeroTax,a.SumOfExceptTax,a.SumOfWTax,a.NetDebtAmount,a.BillBalance,isnull(a.CurrencyCode,'') as CurrencyCode,a.ExchangeRate,isnull(a.GLFormat,'') as GLFormat,a.IsCancel,a.IsCompleteSave,a.ReturnMoney,a.ReturnStatus,a.ReturnCash,a.OtherIncome,a.OtherExpense,a.ExcessAmount1,a.ExcessAmount2,a.SumCashAmount,a.SumChqAmount,a.SumCreditAmount,a.SumBankAmount,a.ChargeAmount,a.ChangeAmount,a.CauseType,a.PayBillStatus,a.IsCNDeposit,a.IsPos,isnull(a.PosDocNo,'') as PosDocNo,isnull(a.CauseCode,'') as CauseCode,isnull(a.AllocateCode,'') as AllocateCode,isnull(a.ProjectCode,'') as ProjectCode,isnull(a.BillGroup,'') as BillGroup,isnull(a.RecurName,'') as RecurName,isnull(a.ConfirmCode,'') as ConfirmCode,isnull(a.ConfirmDateTime,'') as ConfirmDateTime,isnull(a.CancelCode,'') as CancelCode,isnull(a.CancelDateTime,'') as CancelDateTime,a.PayBillAmount,a.BillTemporary from dbo.BCCreditNote a With (Nolock) left join dbo.bcar b with (nolock) on a.arcode = b.code left join dbo.bcsale c with (nolock) on a.salecode = c.code where (a.docno  like '%'+?+'%' or a.arcode like '%'+?+'%' or a.salecode like '%'+?+'%' or b.name1 like '%'+?+'%' or c.name like '%'+?+'%' ) order by a.docdate, a.docno`
	err = db.Select(&crdList, sql, keyword, keyword, keyword, keyword, keyword)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, sub := range crdList {
		//sqlsub := `MyType, DocNo, TaxNo, TaxType, ItemCode, DocDate, ArCode, DepartCode, SaleCode, CashierCode, MyDescription, ItemName, WHCode, ShelfCode, DiscQty, TempQty, BillQty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost, UnitCode, InvoiceNo, ItemType, ExceptTax, IsPos, IsCancel, LineNumber, RefLineNumber, BarCode,AVERAGECOST, LotNumber, PackingRate1, PackingRate2`
		sqlsub := `set dateformat dmy     select a.MyType,a.ItemCode,isnull(a.MyDescription,'') as MyDescription,isnull(a.ItemName,'') as ItemName,isnull(a.WHCode,'') as WHCode,isnull(a.ShelfCode,'') as ShelfCode,a.DiscQty,a.TempQty,a.BillQty,a.Price,isnull(a.DiscountWord,'') as DiscountWord,a.DiscountAmount,a.Amount,a.NetAmount,a.HomeAmount,a.SumOfCost,isnull(a.UnitCode,'') as UnitCode,isnull(a.InvoiceNo,'') as InvoiceNo,a.IsPos,a.IsCancel,a.LineNumber,a.RefLineNumber,isnull(a.BarCode,'') as BarCode,a.AVERAGECOST,isnull(a.LotNumber,'') as LotNumber,isnull(a.PackingRate1,1) as PackingRate1,isnull(a.PackingRate2,1) as PackingRate2 from dbo.BCCreditNoteSub a with (nolock) where a.docno = ?`
		err = db.Select(&sub.Subs, sqlsub, sub.DocNo)
		fmt.Println("Docno =", sqlsub, sub.DocNo)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return crdList, nil
}
