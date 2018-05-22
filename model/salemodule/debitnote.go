package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"errors"
	m "github.com/loukmho/bcaccount_api/model"
)

type DebitNote struct {
	SaveFrom          int                  `json:"save_from" db"SaveFrom"`
	DocNo             string               `json:"doc_no" db:"DocNo"`
	TaxNo             string               `json:"tax_no" db:"TaxNo"`
	DocDate           string               `json:"doc_date" db:"DocDate"`
	CreatorCode       string               `json:"creator_code" db:"CreatorCode"`
	CreateDateTime    string               `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode    string               `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT     string               `json:"last_edit_date_t" db:"LastEditDateT"`
	CreditDay         int                  `json:"credit_day" db:"CreditDay"`
	DueDate           string               `json:"due_date" db:"DueDate"`
	TaxType           int                  `json:"tax_type" db:"TaxType"`
	DbtCustomer
	DbtSaleMan
	DepartCode        string               `json:"depart_code" db:"DepartCode"`
	TaxRate           float64              `json:"tax_rate" db:"TaxRate"`
	IsConfirm         int                  `json:"is_confirm" db:"IsConfirm"`
	MyDescription     string               `json:"my_description" db:"MyDescription"`
	SumOfItemAmount   float64              `json:"sum_of_item_amount" db:"SumOfItemAmount"`
	SumOldAmount      float64              `json:"sum_old_amount" db:"SumOldAmount"`
	SumTrueAmount     float64              `json:"sum_true_amount" db:"SumTrueAmount"`
	SumofDiffAmount   float64              `json:"sumof_diff_amount" db:"SumofDiffAmount"`
	SumofBeforeTax    float64              `json:"sumof_before_tax" db:"SumofBeforeTax"`
	SumOfTaxAmount    float64              `json:"sum_of_tax_amount" db:"SumOfTaxAmount"`
	SumOfTotalTax     float64              `json:"sum_of_total_tax" db:"SumOfTotalTax"`
	SumOfExceptTax    float64              `json:"sum_of_except_tax" db:"SumOfExceptTax"`
	SumOfZeroTax      float64              `json:"sum_of_zero_tax" db:"SumOfZeroTax"`
	SumOfWTax         float64              `json:"sum_of_w_tax" db:"SumOfWTax"`
	DiscountWord      string               `json:"discount_word" db:"DiscountWord"`
	DiscountAmount    float64              `json:"discount_amount" db:"DiscountAmount"`
	NetDebtAmount     float64              `json:"net_debt_amount" db:"NetDebtAmount"`
	BillBalance       float64              `json:"bill_balance" db:"BillBalance"`
	CurrencyCode      string               `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate      float64              `json:"exchange_rate" db:"ExchangeRate"`
	GLFormat          string               `json:"gl_format" db:"GLFormat"`
	IsCancel          int                  `json:"is_cancel" db:"IsCancel"`
	IsCompleteSave    int                  `json:"is_complete_save" db:"IsCompleteSave"`
	ReturnStatus      int                  `json:"return_status" db:"ReturnStatus"`
	CauseType         int                  `json:"cause_type" db:"CauseType"`
	CauseCode         string               `json:"cause_code" db:"CauseCode"`
	PayBillStatus     int                  `json:"pay_bill_status" db:"PayBillStatus"`
	AllocateCode      string               `json:"allocate_code" db:"AllocateCode"`
	ProjectCode       string               `json:"project_code" db:"ProjectCode"`
	BillGroup         string               `json:"bill_group" db:"BillGroup"`
	RecurName         string               `json:"recur_name" db:"RecurName"`
	ConfirmCode       string               `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime   string               `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode        string               `json:"cancel_code" db:"CancelCode"`
	CancelDateTime    string               `json:"cancel_date_time" db:"CancelDateTime"`
	PayBillAmount     float64              `json:"pay_bill_amount" db:"PayBillAmount"`
	BillTemporary     float64              `json:"bill_temporary" db:"BillTemporary"`
	UserCode          string               `json:"user_code" db:"UserCode"`
	ListCrdRecMoney
	Subs              []*DbtItem           `json:"subs"`
	Cdcs              []*ListDbtCreditCard `json:"cdcs"`
	Chqs              []*ListDbtChqIn      `json:"chqs"`
}

type DbtCustomer struct {
	ArCode string `json:"ar_code" db:"ArCode"`
	ArName string `json:"ar_name" db:"ArName"`
}

type DbtSaleMan struct {
	SaleCode string `json:"sale_code" db:"SaleCode"`
	SaleName string `json:"sale_name" db:"SaleName"`
}

type DbtItem struct {
	MyType         int     `json:"my_type" db:"MyType"`
	ItemCode       string  `json:"item_code" db:"ItemCode"`
	ItemType       int     `json:"item_type" db:"ItemType"`
	MyDescription  string  `json:"my_description" db:"MyDescription"`
	ItemName       string  `json:"item_name" db:"ItemName"`
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
	IsCancel       int     `json:"is_cancel" db:"IsCancel"`
	LineNumber     int     `json:"line_number" db:"LineNumber"`
	RefLineNumber  int     `json:"ref_line_number" db:"RefLineNumber"`
	BarCode        string  `json:"bar_code" db:"BarCode"`
	AverageCost    float64 `json:"average_cost" db:"AverageCost"`
	LotNumber      string  `json:"lot_number" db:"LotNumber"`
	PackingRate1   float64 `json:"packing_rate_1" db:"PackingRate1"`
	PackingRate2   float64 `json:"packing_rate_2" db:"PackingRate2"`
}

type ListDbtRecMoney struct {
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

type ListDbtChqIn struct {
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

type ListDbtCreditCard struct {
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

func (dbt *DebitNote) InsertAndEditDebitNote(db *sqlx.DB) error {
	var check_exist int
	var count_item int
	//var sum_pay_amount float64

	//	now := time.Now()

	//sum_pay_amount = (crd.SumCashAmount+crd.SumCreditAmount+crd.SumChqAmount+crd.SumBankAmount+crd.OtherExpense)-crd.OtherIncome

	for _, sub_item := range dbt.Subs {
		if (sub_item.DiscQty != 0) {
			count_item = count_item + 1
		}
	}

	sqlexist := `select count(docno) as check_exist from dbo.bcdebitnote1 where docno = ?`
	err := db.Get(&check_exist, sqlexist, dbt.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil
	}
	fmt.Println("Docno = ", dbt.DocNo)
	switch {
	case dbt.DocNo == "":
		return errors.New("Docno is null")
	case dbt.ArCode == "":
		return errors.New("Arcode is null")
	case dbt.DocDate == "":
		return errors.New("Docdate is null")
	case dbt.IsCancel == 1:
		return errors.New("Docno is cancel")
	case dbt.IsConfirm == 1:
		return errors.New("Docno is confirm")
	case dbt.SumOldAmount == 0:
		return errors.New("Docno not have old bill amount")
	case dbt.SumOfTotalTax == 0:
		return errors.New("Docno not have amount return")
	case dbt.SumOfItemAmount == 0 && count_item == 0:
		return errors.New("Docno not have SumOfItemAmount")
	}

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")

	fmt.Println("TaxRate = ", def.TaxRateDefault)

	if dbt.TaxRate == 0 {
		dbt.TaxRate = def.TaxRateDefault
	}
	if dbt.ExchangeRate == 0 {
		dbt.ExchangeRate = def.ExchangeRateDefault
	}

	dbt.SumofBeforeTax, dbt.SumOfTaxAmount = m.CalcTaxCredit(dbt.TaxType, dbt.TaxRate, dbt.SumOfTotalTax)
	fmt.Println("Befor,Tax", dbt.SumofBeforeTax, dbt.SumOfTaxAmount)

	dbt.BillBalance = dbt.SumOfTotalTax
	dbt.NetDebtAmount = dbt.SumOfTotalTax
	dbt.SumTrueAmount = dbt.SumOldAmount - dbt.SumOfTotalTax

	if dbt.TaxType == 1 {
		dbt.PayBillAmount = dbt.SumOfTotalTax
		dbt.BillTemporary = dbt.SumOfTotalTax
	}

	if (dbt.GLFormat == "") {
		dbt.GLFormat = def.CreditNoteGLFormat
	}
	if (dbt.TaxNo == "" ) {
		dbt.TaxNo = dbt.DocNo
	}

	fmt.Println("UserCode = ", dbt.UserCode)

	if (dbt.IsCompleteSave == 0) {
		dbt.IsCompleteSave = 1
	}

	if check_exist == 0 {
		sql := `insert into dbo.bcdebitnote1(DocNo,TaxNo,DocDate,CreatorCode,CreateDateTime,CreditDay,DueDate,TaxType,ArCode,DepartCode,SaleCode,TaxRate,IsConfirm,MyDescription,SumOfItemAmount,SumOldAmount,SumTrueAmount,SumofDiffAmount,SumofBeforeTax,SumOfTaxAmount,SumOfTotalTax,SumOfExceptTax,SumOfZeroTax,SumOfWTax,DiscountWord,DiscountAmount,NetDebtAmount,BillBalance,CurrencyCode,ExchangeRate,GLFormat,IsCancel,IsCompleteSave,ReturnStatus,CauseType,CauseCode,PayBillStatus,AllocateCode,ProjectCode,BillGroup,RecurName,PayBillAmount,BillTemporary) values(?,?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sql, dbt.DocNo,dbt.TaxNo,dbt.DocDate,dbt.UserCode,dbt.CreditDay,dbt.DueDate,dbt.TaxType,dbt.ArCode,dbt.DepartCode,dbt.SaleCode,dbt.TaxRate,dbt.IsConfirm,dbt.MyDescription,dbt.SumOfItemAmount,dbt.SumOldAmount,dbt.SumTrueAmount,dbt.SumofDiffAmount,dbt.SumofBeforeTax,dbt.SumOfTaxAmount,dbt.SumOfTotalTax,dbt.SumOfExceptTax,dbt.SumOfZeroTax,dbt.SumOfWTax,dbt.DiscountWord,dbt.DiscountAmount,dbt.NetDebtAmount,dbt.BillBalance,dbt.CurrencyCode,dbt.ExchangeRate,dbt.GLFormat,dbt.IsCancel,dbt.IsCompleteSave,dbt.ReturnStatus,dbt.CauseType,dbt.CauseCode,dbt.PayBillStatus,dbt.AllocateCode,dbt.ProjectCode,dbt.BillGroup,dbt.RecurName,dbt.PayBillAmount,dbt.BillTemporary)
		if err != nil {
			fmt.Println("Insert Debit =", err.Error())
			return err
		}
	} else {
		sql := `update dbo.bcdebitnote1 set TaxNo=?,DocDate=?,LastEditorCode=?,LastEditDateT=getdate(),CreditDay=?,DueDate=?,TaxType=?,ArCode=?,DepartCode=?,SaleCode=?,TaxRate=?,IsConfirm=?,MyDescription=?,SumOfItemAmount=?,SumOldAmount=?,SumTrueAmount=?,SumofDiffAmount=?,SumofBeforeTax=?,SumOfTaxAmount=?,SumOfTotalTax=?,SumOfExceptTax=?,SumOfZeroTax=?,SumOfWTax=?,DiscountWord=?,DiscountAmount=?,NetDebtAmount=?,BillBalance=?,CurrencyCode=?,ExchangeRate=?,GLFormat=?,IsCancel=?,IsCompleteSave=?,ReturnStatus=?,CauseType=?,CauseCode=?,PayBillStatus=?,AllocateCode=?,ProjectCode=?,BillGroup=?,RecurName=?,PayBillAmount=?,BillTemporary=? where docno = ? `
		_, err = db.Exec(sql, dbt.TaxNo,dbt.DocDate,dbt.UserCode,dbt.CreditDay,dbt.DueDate,dbt.TaxType,dbt.ArCode,dbt.DepartCode,dbt.SaleCode,dbt.TaxRate,dbt.IsConfirm,dbt.MyDescription,dbt.SumOfItemAmount,dbt.SumOldAmount,dbt.SumTrueAmount,dbt.SumofDiffAmount,dbt.SumofBeforeTax,dbt.SumOfTaxAmount,dbt.SumOfTotalTax,dbt.SumOfExceptTax,dbt.SumOfZeroTax,dbt.SumOfWTax,dbt.DiscountWord,dbt.DiscountAmount,dbt.NetDebtAmount,dbt.BillBalance,dbt.CurrencyCode,dbt.ExchangeRate,dbt.GLFormat,dbt.IsCancel,dbt.IsCompleteSave,dbt.ReturnStatus,dbt.CauseType,dbt.CauseCode,dbt.PayBillStatus,dbt.AllocateCode,dbt.ProjectCode,dbt.BillGroup,dbt.RecurName,dbt.PayBillAmount,dbt.BillTemporary,dbt.DocNo)
		if err != nil {
			fmt.Println("Update Debit =", err.Error())
			return err
		}
	}

	sql_del_sub := `delete dbo.bccreditnotesub where docno = ?`
	_, err = db.Exec(sql_del_sub, dbt.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}

	var vLineNumber int

	for _, item := range dbt.Subs {
		fmt.Println("ItemSub")
		item.MyType = def.DebitNoteMyType
		item.TempQty = item.DiscQty

		item.LineNumber = vLineNumber

		item.IsCancel = 0
		if (item.PackingRate1 == 0) {
			item.PackingRate1 = 1
		}
		if (item.PackingRate2 == 0) {
			item.PackingRate2 = 1
		}

		if item.Amount == 0 {
			item.Amount = (item.DiscQty * (item.Price - (item.DiscountAmount / item.DiscQty)))
		}

		switch {
		case dbt.TaxType == 0:
			item.HomeAmount = item.Amount
			item.NetAmount = item.Amount
		case dbt.TaxType == 1:
			taxamount := m.ToFixed(item.Amount-((item.Amount*100)/(100+float64(dbt.TaxRate))), 2)
			beforetaxamount := m.ToFixed(item.Amount-taxamount, 2)
			item.HomeAmount = beforetaxamount
			item.NetAmount = beforetaxamount
		case dbt.TaxType == 2:
			item.HomeAmount = item.Amount
			item.NetAmount = item.Amount
		}

		sqlsub := ` insert into dbo.BCDebitNoteSub1(MyType,DocNo,TaxNo,TaxType,ItemCode,ItemType,DocDate,ArCode,DepartCode,SaleCode,MyDescription,ItemName,WHCode,ShelfCode,DiscQty,TempQty,BillQty,Price,DiscountWord,DiscountAmount,Amount,NetAmount,HomeAmount,SumOfCost,UnitCode,InvoiceNo,IsCancel,LineNumber,RefLineNumber,BarCode,AverageCost,LotNumber,PackingRate1,PackingRate2) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlsub, item.MyType,dbt.DocNo,dbt.TaxNo,dbt.TaxType,item.ItemCode,item.ItemType,dbt.DocDate,dbt.ArCode,dbt.DepartCode,dbt.SaleCode,item.MyDescription,item.ItemName,item.WHCode,item.ShelfCode,item.DiscQty,item.TempQty,item.BillQty,item.Price,item.DiscountWord,item.DiscountAmount,item.Amount,item.NetAmount,item.HomeAmount,item.SumOfCost,item.UnitCode,item.InvoiceNo,item.IsCancel,item.LineNumber,item.RefLineNumber,item.BarCode,item.AverageCost,item.LotNumber,item.PackingRate1,item.PackingRate2)
		fmt.Println("sqltax = ", sqlsub)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}

		if (item.ItemType != 1){
			sqlprocess := ` insert into dbo.ProcessStock (ItemCode,ProcessFlag,FlowStatus) values(?, 1, 0)`
			_, err = db.Exec(sqlprocess, item.ItemCode )
			fmt.Println("sqlprocess = ", sqlsub)
			if err != nil {
				fmt.Println("Error = ", err.Error())
				fmt.Println(err.Error())
			}
		}

		vLineNumber = vLineNumber + 1
	}

	return nil
}

func (dbt *DebitNote) SearchDebitNoteByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select a.DocNo, isnull(a.TaxNo,'') as TaxNo, a.DocDate, a.CreatorCode, a.CreateDateTime, isnull(a.LastEditorCode,'') as LastEditorCode, isnull(a.LastEditDateT,'') as LastEditDateT, isnull(a.CreditDay,0) as CreditDay, isnull(a.DueDate,'') as DueDate, isnull(a.TaxType,0) as TaxType, isnull(a.ArCode,'') as ArCode, isnull(b.name1,'') as ArName,isnull(a.DepartCode,'') as DepartCode, isnull(a.SaleCode,'') as SaleCode, isnull(c.name,'') as SaleName, isnull(a.TaxRate,0) as TaxRate, isnull(a.IsConfirm,0) as IsConfirm, isnull(a.MyDescription,'') as MyDescription, isnull(a.SumOfItemAmount,0) as SumOfItemAmount, isnull(a.SumOldAmount,0) as SumOldAmount, isnull(a.SumTrueAmount,0) as SumTrueAmount, isnull(a.SumofDiffAmount,0) as SumofDiffAmount, isnull(SumofBeforeTax,0) as SumofBeforeTax, isnull(a.SumOfTaxAmount,0) as SumOfTaxAmount, isnull(a.SumOfTotalTax,0) as SumOfTotalTax, isnull(a.SumOfExceptTax,0) as SumOfExceptTax, isnull(a.SumOfZeroTax,0) as SumOfZeroTax, isnull(a.SumOfWTax,0) as SumOfWTax, isnull(a.DiscountWord,'') as DiscountWord, isnull(a.DiscountAmount,0) as DiscountAmount, isnull(a.NetDebtAmount,0) as NetDebtAmount, isnull(a.BillBalance,0) as BillBalance, isnull(a.CurrencyCode,'') as CurrencyCode, isnull(a.ExchangeRate,0) as ExchangeRate, isnull(a.GLFormat,'') as GLFormat, isnull(a.IsCancel,0) as IsCancel,  isnull(a.IsCompleteSave,0) as IsCompleteSave, isnull(a.ReturnStatus,0) as ReturnStatus, isnull(a.CauseType,0) as CauseType, isnull(a.CauseCode,'') as CauseCode, isnull(a.PayBillStatus,0) as PayBillStatus, isnull(a.AllocateCode,'') as AllocateCode, isnull(a.ProjectCode,'') as ProjectCode, isnull(a.BillGroup,'') as BillGroup, isnull(a.RecurName,'') as RecurName, isnull(a.ConfirmCode,'') as ConfirmCode, isnull(a.ConfirmDateTime,'') as ConfirmDateTime, isnull(a.CancelCode,'') as CancelCode, isnull(a.CancelDateTime,'') as CancelDateTime, isnull(a.PayBillAmount,0) as PayBillAmount, isnull(a.BillTemporary,0) as BillTemporary from dbo.BCDebitNote1  a With (Nolock) left join dbo.bcar b with (nolock) on a.arcode = b.code left join dbo.bcsale c with (nolock) on a.salecode = c.code where a.docno = ?`
	err := db.Get(dbt, sql, docno)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}
	//sqlsub := `MyType, DocNo, TaxNo, TaxType, ItemCode, DocDate, ArCode, DepartCode, SaleCode, CashierCode, MyDescription, ItemName, WHCode, ShelfCode, DiscQty, TempQty, BillQty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost, UnitCode, InvoiceNo, ItemType, ExceptTax, IsPos, IsCancel, LineNumber, RefLineNumber, BarCode,AVERAGECOST, LotNumber, PackingRate1, PackingRate2`
	sqlsub := `set dateformat dmy     select MyType, ItemCode, ItemType, isnull(MyDescription,'') as MyDescription, isnull(ItemName,'') as ItemName, isnull(WHCode,'') as WHCode, isnull(ShelfCode,'') as ShelfCode, isnull(DiscQty,0) as DiscQty, isnull(TempQty,0) as TempQty, isnull(BillQty,0) as BillQty, isnull(Price,0) as Price, isnull(DiscountWord,'') as DiscountWord, isnull(DiscountAmount,0) as DiscountAmount, isnull(Amount,0) as Amount, isnull(NetAmount,0) as NetAmount, isnull(HomeAmount,0) as HomeAmount, isnull(SumOfCost,0) as SumOfCost, isnull(UnitCode,'') as UnitCode, isnull(InvoiceNo,'') as InvoiceNo, isnull(IsCancel,0) as IsCancel, isnull(LineNumber,0) as LineNumber, isnull(RefLineNumber,0) as RefLineNumber, isnull(BarCode,'') as BarCode, isnull(AVERAGECOST,0) as AverageCost, isnull(LotNumber,'') as LotNumber, isnull(PackingRate1,1) as PackingRate1, isnull(PackingRate2,1) as PackingRate2 from dbo.BCDebitNoteSub1 a With (Nolock) where a.docno = ? order by linenumber`
	err = db.Select(&dbt.Subs, sqlsub, docno)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}
	return nil
}

func (dbt *DebitNote) SearchDebitNoteByKeyword(db *sqlx.DB, keyword string) (dbtList []*DebitNote, err error) {
	sql := `set dateformat dmy     select a.DocNo, isnull(a.TaxNo,'') as TaxNo, a.DocDate, a.CreatorCode, a.CreateDateTime, isnull(a.LastEditorCode,'') as LastEditorCode, isnull(a.LastEditDateT,'') as LastEditDateT, isnull(a.CreditDay,0) as CreditDay, isnull(a.DueDate,'') as DueDate, isnull(a.TaxType,0) as TaxType, isnull(a.ArCode,'') as ArCode, isnull(b.name1,'') as ArName,isnull(a.DepartCode,'') as DepartCode, isnull(a.SaleCode,'') as SaleCode, isnull(c.name,'') as SaleName, isnull(a.TaxRate,0) as TaxRate, isnull(a.IsConfirm,0) as IsConfirm, isnull(a.MyDescription,'') as MyDescription, isnull(a.SumOfItemAmount,0) as SumOfItemAmount, isnull(a.SumOldAmount,0) as SumOldAmount, isnull(a.SumTrueAmount,0) as SumTrueAmount, isnull(a.SumofDiffAmount,0) as SumofDiffAmount, isnull(SumofBeforeTax,0) as SumofBeforeTax, isnull(a.SumOfTaxAmount,0) as SumOfTaxAmount, isnull(a.SumOfTotalTax,0) as SumOfTotalTax, isnull(a.SumOfExceptTax,0) as SumOfExceptTax, isnull(a.SumOfZeroTax,0) as SumOfZeroTax, isnull(a.SumOfWTax,0) as SumOfWTax, isnull(a.DiscountWord,'') as DiscountWord, isnull(a.DiscountAmount,0) as DiscountAmount, isnull(a.NetDebtAmount,0) as NetDebtAmount,isnull(a.BillBalance,0) as BillBalance, isnull(a.CurrencyCode,'') as CurrencyCode, isnull(a.ExchangeRate,0) as ExchangeRate, isnull(a.GLFormat,'') as GLFormat, isnull(a.IsCancel,0) as IsCancel,  isnull(a.IsCompleteSave,0) as IsCompleteSave, isnull(a.ReturnStatus,0) as ReturnStatus, isnull(a.CauseType,0) as CauseType, isnull(a.CauseCode,'') as CauseCode, isnull(a.PayBillStatus,0) as PayBillStatus, isnull(a.AllocateCode,'') as AllocateCode, isnull(a.ProjectCode,'') as ProjectCode, isnull(a.BillGroup,'') as BillGroup, isnull(a.RecurName,'') as RecurName, isnull(a.ConfirmCode,'') as ConfirmCode, isnull(a.ConfirmDateTime,'') as ConfirmDateTime, isnull(a.CancelCode,'') as CancelCode, isnull(a.CancelDateTime,'') as CancelDateTime, isnull(a.PayBillAmount,0) as PayBillAmount, isnull(a.BillTemporary,0) as BillTemporary from dbo.BCDebitNote1  a With (Nolock) left join dbo.bcar b with (nolock) on a.arcode = b.code left join dbo.bcsale c with (nolock) on a.salecode = c.code where (a.docno  like '%'+?+'%' or a.arcode like '%'+?+'%' or a.salecode like '%'+?+'%' or b.name1 like '%'+?+'%' or c.name like '%'+?+'%' ) order by a.docdate, a.docno`
	err = db.Select(&dbtList, sql, keyword, keyword, keyword, keyword, keyword)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}
	for _, sub := range dbtList {
		//sqlsub := `MyType, DocNo, TaxNo, TaxType, ItemCode, DocDate, ArCode, DepartCode, SaleCode, CashierCode, MyDescription, ItemName, WHCode, ShelfCode, DiscQty, TempQty, BillQty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost, UnitCode, InvoiceNo, ItemType, ExceptTax, IsPos, IsCancel, LineNumber, RefLineNumber, BarCode,AVERAGECOST, LotNumber, PackingRate1, PackingRate2`
		sqlsub := `set dateformat dmy     select MyType, ItemCode, ItemType, isnull(MyDescription,'') as MyDescription, isnull(ItemName,'') as ItemName, isnull(WHCode,'') as WHCode, isnull(ShelfCode,'') as ShelfCode, isnull(DiscQty,0) as DiscQty, isnull(TempQty,0) as TempQty, isnull(BillQty,0) as BillQty, isnull(Price,0) as Price, isnull(DiscountWord,'') as DiscountWord, isnull(DiscountAmount,0) as DiscountAmount, isnull(Amount,0) as Amount, isnull(NetAmount,0) as NetAmount, isnull(HomeAmount,0) as HomeAmount, isnull(SumOfCost,0) as SumOfCost, isnull(UnitCode,'') as UnitCode, isnull(InvoiceNo,'') as InvoiceNo, isnull(IsCancel,0) as IsCancel, isnull(LineNumber,0) as LineNumber, isnull(RefLineNumber,0) as RefLineNumber, isnull(BarCode,'') as BarCode, isnull(AVERAGECOST,0) as AverageCost, isnull(LotNumber,'') as LotNumber, isnull(PackingRate1,1) as PackingRate1, isnull(PackingRate2,1) as PackingRate2 from dbo.BCDebitNoteSub1 a With (Nolock) where a.docno = ? order by linenumber`
		err = db.Select(&sub.Subs, sqlsub, sub.DocNo)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}
	}
	return dbtList, nil
}
