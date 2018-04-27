package model

import (
	"github.com/jmoiron/sqlx"
	m "github.com/loukmho/bcaccount_api/model"
	"fmt"
	"errors"
)

type StkRefund struct {
	SaveFrom          int        `json:"save_from" db:"SaveFrom"`
	Source            int        `json:"source" db:"Source"`
	DocNo             string     `json:"doc_no" db:"DocNo"`
	TaxNo             string     `json:"tax_no" db:"TaxNo"`
	DocDate           string     `json:"doc_date" db:"DocDate"`
	CreatorCode       string     `json:"creator_code" db:"CreatorCode"`
	CreateDateTime    string     `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode    string     `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT     string     `json:"last_edit_date_t" db:"LastEditDateT"`
	DueDate           string     `json:"due_date" db:"DueDate"`
	TaxType           int        `json:"tax_type" db:"TaxType"`
	SrfVendor
	SrfInPutTax
	DepartCode        string     `json:"depart_code" db:"DepartCode"`
	TaxRate           float64    `json:"tax_rate" db:"TaxRate"`
	IsConfirm         int        `json:"is_confirm" db:"IsConfirm"`
	MyDescription     string     `json:"my_description" db:"MyDescription"`
	SumOfItemAmount   float64    `json:"sum_of_item_amount" db:"SumOfItemAmount"`
	SumOldAmount      float64    `json:"sum_old_amount" db:"SumOldAmount"`
	SumTrueAmount     float64    `json:"sum_true_amount" db:"SumTrueAmount"`
	SumofDiffAmount   float64    `json:"sumof_diff_amount" db:"SumofDiffAmount"`
	DiscountWord      string     `json:"discount_word" db:"DiscountWord"`
	DiscountAmount    float64    `json:"discount_amount" db:"DiscountAmount"`
	SumofBeforeTax    float64    `json:"sumof_before_tax" db:"SumofBeforeTax"`
	SumOfTaxAmount    float64    `json:"sum_of_tax_amount" db:"SumOfTaxAmount"`
	SumOfTotalTax     float64    `json:"sum_of_total_tax" db:"SumOfTotalTax"`
	SumOfExceptTax    float64    `json:"sum_of_except_tax" db:"SumOfExceptTax"`
	SumOfZeroTax      float64    `json:"sum_of_zero_tax" db:"SumOfZeroTax"`
	SumOfWTax         float64    `json:"sum_of_w_tax" db:"SumOfWTax"`
	NetDebtAmount     float64    `json:"net_debt_amount" db:"NetDebtAmount"`
	SumExchangeProfit float64    `json:"sum_exchange_profit" db:"SumExchangeProfit"`
	BillBalance       float64    `json:"bill_balance" db:"BillBalance"`
	CurrencyCode      string     `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate      float64    `json:"exchange_rate" db:"ExchangeRate"`
	GLFormat          string     `json:"gl_format" db:"GLFormat"`
	GLStartPosting    int        `json:"gl_start_posting" db:"GLStartPosting"`
	IsCancel          int        `json:"is_cancel" db:"IsCancel"`
	IsCompleteSave    int        `json:"is_complete_save" db:"IsCompleteSave"`
	ReturnMoney       int        `json:"return_money" db:"ReturnMoney"`
	BillType          int        `json:"bill_type" db:"BillType"`
	CauseType         int        `json:"cause_type" db:"CauseType"`
	CauseCode         string     `json:"cause_code" db:"CauseCode"`
	AllocateCode      string     `json:"allocate_code" db:"AllocateCode"`
	ProjectCode       string     `json:"project_code" db:"ProjectCode"`
	BillGroup         string     `json:"bill_group" db:"BillGroup"`
	RecurName         string     `json:"recur_name" db:"RecurName"`
	ConfirmCode       string     `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime   string     `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode        string     `json:"cancel_code" db:"CancelCode"`
	CancelDateTime    string     `json:"cancel_date_time" db:"CancelDateTime"`
	PayBillAmount     float64    `json:"pay_bill_amount" db:"PayBillAmount"`
	UserCode          string     `json:"user_code" db:"UserCode"`
	Subs              []*SrfItem `json:"subs"`
}

type SrfVendor struct {
	ApCode string `json:"ap_code" db:"ApCode"`
	ApName string `json:"ap_name" db:"ApName"`
}

type SrfInPutTax struct {
	TaxNo    string `json:"tax_no" db:"TaxNo"`
	BookCode string `json:"book_code" db:"BookCode"`
}

type SrfItem struct {
	MyType         int     `json:"my_type" db:"MyType"`
	ItemCode       string  `json:"item_code" db:"ItemCode"`
	MyDescription  string  `json:"my_description" db:"MyDescription"`
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
	UnitCode       string  `json:"unit_code" db:"UnitCode"`
	IsCancel       int     `json:"is_cancel" db:"IsCancel"`
	LineNumber     int     `json:"line_number" db:"LineNumber"`
	RefLineNumber  int     `json:"ref_line_number" db:"RefLineNumber"`
	BarCode        string  `json:"bar_code" db:"BarCode"`
	AverageCost    float64 `json:"average_cost" db:"AverageCost"`
	SumOfCost      float64 `json:"sum_of_cost" db:"SumOfCost"`
	LotNumber      string  `json:"lot_number" db:"LotNumber"`
	ItemName       string  `json:"item_name" db:"ItemName"`
	PackingRate1   float64 `json:"packing_rate_1" db:"PackingRate1"`
	PackingRate2   float64 `json:"packing_rate_2" db:"PackingRate2"`
}

//func (stf *StkRefund) SearchStkRefundByDocNo() error {
//	//sql := `DocNo, TaxNo, DocDate,CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, DueDate, TaxType, ApCode, DepartCode, TaxRate, IsConfirm, MyDescription, SumOfItemAmount,SumOldAmount, SumTrueAmount, SumofDiffAmount, DiscountWord, DiscountAmount, SumofBeforeTax, SumOfTaxAmount, SumOfTotalTax, SumOfExceptTax, SumOfZeroTax, SumOfWTax, NetDebtAmount, SumExchangeProfit, BillBalance, CurrencyCode, ExchangeRate, GLFormat, IsCancel,  IsCompleteSave, ReturnMoney,   BillType, CauseType, CauseCode,AllocateCode, ProjectCode, BillGroup, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime, PayBillAmount`
//	//sqlsub := `MyType, DocNo, TaxNo, TaxType, ItemCode, DocDate, ApCode, DepartCode, MyDescription, WHCode, ShelfCode,DiscQty, TempQty, BillQty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, UnitCode, IsCancel,  LineNumber, RefLineNumber, BarCode,AVERAGECOST, SumOfCost,LotNumber,  ItemName, TaxRate, PackingRate1, PackingRate2`
//	return nil
//}

func (srf *StkRefund) InsertAndEditStkRefund(db *sqlx.DB) error {
	var check_exist int
	var count_item int
	//var sum_pay_amount float64

	//	now := time.Now()

	//sum_pay_amount = (crd.SumCashAmount+crd.SumCreditAmount+crd.SumChqAmount+crd.SumBankAmount+crd.OtherExpense)-crd.OtherIncome

	for _, sub_item := range srf.Subs {
		if (sub_item.DiscQty != 0) {
			count_item = count_item + 1
		}
	}

	sqlexist := `select count(docno) as check_exist from dbo.bcstkrefund where docno = ?`
	err := db.Get(&check_exist, sqlexist, srf.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println("Docno = ", srf.DocNo)
	switch {
	case srf.DocNo == "":
		return errors.New("Docno is null")
	case srf.ApCode == "":
		return errors.New("Apcode is null")
	case srf.DocDate == "":
		return errors.New("Docdate is null")
	case srf.IsCancel == 1:
		return errors.New("Docno is cancel")
	case srf.IsConfirm == 1:
		return errors.New("Docno is confirm")
	case srf.SumOldAmount == 0:
		return errors.New("Docno not have old bill amount")
	case srf.SumOfTotalTax == 0:
		return errors.New("Docno not have amount return")
	case srf.SumOfItemAmount == 0 && count_item == 0:
		return errors.New("Docno not have SumOfItemAmount")
	}

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")

	fmt.Println("TaxRate = ", def.TaxRateDefault)

	if srf.TaxRate == 0 {
		srf.TaxRate = def.TaxRateDefault
	}
	if srf.ExchangeRate == 0 {
		srf.ExchangeRate = def.ExchangeRateDefault
	}

	if srf.SaveFrom == 0 {
		srf.SaveFrom = def.StkRefundSaveFrom
	}

	srf.SumofBeforeTax, srf.SumOfTaxAmount = m.CalcTaxCredit(srf.TaxType, srf.TaxRate, srf.SumOfTotalTax)
	fmt.Println("Befor,Tax", srf.SumofBeforeTax, srf.SumOfTaxAmount)

	if srf.ReturnMoney == 0 {
		srf.BillBalance = srf.SumOfTotalTax
	} else {
		srf.BillBalance = 0
	}

	srf.NetDebtAmount = srf.SumOfTotalTax
	srf.SumTrueAmount = srf.SumOldAmount - srf.SumOfTotalTax

	if srf.TaxType == 1 {
		srf.PayBillAmount = srf.SumOfTotalTax
	}

	if (srf.Source == 0) {
		srf.Source = def.StkRefundSource
	}
	if (srf.GLFormat == "") {
		srf.GLFormat = def.StkRefundGLFormat
	}
	if (srf.TaxNo == "" ) {
		srf.TaxNo = srf.DocNo
	}

	fmt.Println("UserCode = ", srf.UserCode)

	if (srf.IsCompleteSave == 0) {
		srf.IsCompleteSave = 1
	}

	if check_exist == 0 {
		srf.CreatorCode = srf.UserCode

		sql := `insert into dbo.bcstkrefund(DocNo,TaxNo,DocDate,CreatorCode,CreateDateTime,DueDate,TaxType,ApCode,DepartCode,TaxRate,IsConfirm,MyDescription,SumOfItemAmount,SumOldAmount,SumTrueAmount,SumofDiffAmount,DiscountWord,DiscountAmount,SumofBeforeTax,SumOfTaxAmount,SumOfTotalTax,SumOfExceptTax,SumOfZeroTax,SumOfWTax,NetDebtAmount,SumExchangeProfit,BillBalance,CurrencyCode,ExchangeRate,GLFormat,GLStartPosting, IsCancel,IsCompleteSave,ReturnMoney,BillType,CauseType,CauseCode,AllocateCode,ProjectCode,BillGroup,RecurName,PayBillAmount) values(?,?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sql, srf.DocNo, srf.TaxNo, srf.DocDate, srf.CreatorCode, srf.DueDate, srf.TaxType, srf.ApCode, srf.DepartCode, srf.TaxRate, srf.IsConfirm, srf.MyDescription, srf.SumOfItemAmount, srf.SumOldAmount, srf.SumTrueAmount, srf.SumofDiffAmount, srf.DiscountWord, srf.DiscountAmount, srf.SumofBeforeTax, srf.SumOfTaxAmount, srf.SumOfTotalTax, srf.SumOfExceptTax, srf.SumOfZeroTax, srf.SumOfWTax, srf.NetDebtAmount, srf.SumExchangeProfit, srf.BillBalance, srf.CurrencyCode, srf.ExchangeRate, srf.GLFormat, srf.GLStartPosting, srf.IsCancel, srf.IsCompleteSave, srf.ReturnMoney, srf.BillType, srf.CauseType, srf.CauseCode, srf.AllocateCode, srf.ProjectCode, srf.BillGroup, srf.RecurName, srf.PayBillAmount)
		if err != nil {
			fmt.Println("Insert Credit =", err.Error())
			return err
		}
	} else {
		srf.LastEditorCode = srf.UserCode

		sql := `update dbo.bcstkrefund set TaxNo=?,DocDate=?,CreatorCode=?,LastEditDateT=getdate(),DueDate=?,TaxType=?,ApCode=?,DepartCode=?,TaxRate=?,IsConfirm=?,MyDescription=?,SumOfItemAmount=?,SumOldAmount=?,SumTrueAmount=?,SumofDiffAmount=?,DiscountWord=?,DiscountAmount=?,SumofBeforeTax=?,SumOfTaxAmount=?,SumOfTotalTax=?,SumOfExceptTax=?,SumOfZeroTax=?,SumOfWTax=?,NetDebtAmount=?,SumExchangeProfit=?,BillBalance=?,CurrencyCode=?,ExchangeRate=?,GLFormat=?, GLStartPosting=?,IsCancel=?,IsCompleteSave=?,ReturnMoney=?,BillType=?,CauseType=?,CauseCode=?,AllocateCode=?,ProjectCode=?,BillGroup=?,RecurName=?,PayBillAmount=? where docno = ? `
		_, err = db.Exec(sql, srf.TaxNo, srf.DocDate, srf.LastEditorCode, srf.DueDate, srf.TaxType, srf.ApCode, srf.DepartCode, srf.TaxRate, srf.IsConfirm, srf.MyDescription, srf.SumOfItemAmount, srf.SumOldAmount, srf.SumTrueAmount, srf.SumofDiffAmount, srf.DiscountWord, srf.DiscountAmount, srf.SumofBeforeTax, srf.SumOfTaxAmount, srf.SumOfTotalTax, srf.SumOfExceptTax, srf.SumOfZeroTax, srf.SumOfWTax, srf.NetDebtAmount, srf.SumExchangeProfit, srf.BillBalance, srf.CurrencyCode, srf.ExchangeRate, srf.GLFormat, srf.GLStartPosting, srf.IsCancel, srf.IsCompleteSave, srf.ReturnMoney, srf.BillType, srf.CauseType, srf.CauseCode, srf.AllocateCode, srf.ProjectCode, srf.BillGroup, srf.RecurName, srf.PayBillAmount, srf.DocNo)
		if err != nil {
			fmt.Println("Insert Credit =", err.Error())
			return err
		}
	}

	sql_del_sub := `delete dbo.bcstkrefundsub where docno = ?`
	_, err = db.Exec(sql_del_sub, srf.DocNo)
	if err != nil {
		return err
	}

	var vLineNumber int

	for _, item := range srf.Subs {
		fmt.Println("ItemSub")
		item.MyType = def.StkRefundMyType
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
		case srf.TaxType == 0:
			item.HomeAmount = item.Amount
			item.NetAmount = item.Amount
		case srf.TaxType == 1:
			taxamount := m.ToFixed(item.Amount-((item.Amount*100)/(100+float64(srf.TaxRate))), 2)
			beforetaxamount := m.ToFixed(item.Amount-taxamount, 2)
			item.HomeAmount = beforetaxamount
			item.NetAmount = beforetaxamount
		case srf.TaxType == 2:
			item.HomeAmount = item.Amount
			item.NetAmount = item.Amount
		}

		sqlsub := ` insert into dbo.BCStkRefundSub(MyType,DocNo,TaxNo,TaxType,ItemType,ItemCode,DocDate,ApCode,DepartCode,MyDescription,WHCode,ShelfCode,DiscQty,TempQty,BillQty,Price,DiscountWord,DiscountAmount,Amount,NetAmount,HomeAmount,UnitCode,IsCancel,LineNumber,RefLineNumber,BarCode,AVERAGECOST,SumOfCost,LotNumber,ItemName,TaxRate,PackingRate1,PackingRate2) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sqlsub, item.MyType, srf.DocNo, srf.TaxNo, srf.TaxType, item.ItemType, item.ItemCode, srf.DocDate, srf.ApCode, srf.DepartCode, item.MyDescription, item.WHCode, item.ShelfCode, item.DiscQty, item.TempQty, item.BillQty, item.Price, item.DiscountWord, item.DiscountAmount, item.Amount, item.NetAmount, item.HomeAmount, item.UnitCode, item.IsCancel, item.LineNumber, item.RefLineNumber, item.BarCode, item.AverageCost, item.SumOfCost, item.LotNumber, item.ItemName, srf.TaxRate, item.PackingRate1, item.PackingRate2)
		fmt.Println("sqltax = ", sqlsub)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		if (item.ItemType != 1) {
			sqlprocess := ` insert into dbo.ProcessStock (ItemCode,ProcessFlag,FlowStatus) values(?, 1, 0)`
			_, err = db.Exec(sqlprocess, item.ItemCode)
			fmt.Println("sqlprocess = ", sqlsub)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		vLineNumber = vLineNumber + 1
	}

	return nil
}

//func (srf *StkRefund) InsertAndEditStkRefund(db *sqlx.DB) error {
//	sql := `insert into dbo.BCSTKRefund (DocNo,TaxNo,DocDate,CreatorCode,CreateDateTime,LastEditorCode,LastEditDateT,DueDate,TaxType,ApCode,DepartCode,TaxRate,IsConfirm,MyDescription,SumOfItemAmount,SumOldAmount,SumTrueAmount,SumofDiffAmount,DiscountWord,DiscountAmount,SumofBeforeTax,SumOfTaxAmount,SumOfTotalTax,SumOfExceptTax,SumOfZeroTax,SumOfWTax,NetDebtAmount,SumExchangeProfit,BillBalance,CurrencyCode,ExchangeRate,GLFormat,IsCancel,IsCompleteSave,ReturnMoney,BillType,CauseType,CauseCode,AllocateCode,ProjectCode,BillGroup,RecurName,ConfirmCode,ConfirmDateTime,CancelCode,CancelDateTime,PayBillAmount) values()`
//	err := db.Get(srf, sql,)
//	if err != nil {
//		fmt.Println(err)
//		return err
//	}
//
//	//sqlsub := `insert into dbo.bcstkrefundsub(MyType,DocNo,TaxNo,TaxType,ItemCode,DocDate,ApCode,DepartCode,MyDescription,WHCode,ShelfCode,DiscQty,TempQty,BillQty,Price,DiscountWord,DiscountAmount,Amount,NetAmount,HomeAmount,UnitCode,IsCancel,LineNumber,RefLineNumber,BarCode,AVERAGECOST,SumOfCost,LotNumber,ItemName,TaxRate,PackingRate1,PackingRate2) values(MyType,DocNo,TaxNo,TaxType,ItemCode,DocDate,ApCode,DepartCode,MyDescription,WHCode,ShelfCode,DiscQty,TempQty,BillQty,Price,DiscountWord,DiscountAmount,Amount,NetAmount,HomeAmount,UnitCode,IsCancel,LineNumber,RefLineNumber,BarCode,AVERAGECOST,SumOfCost,LotNumber,ItemName,TaxRate,PackingRate1,PackingRate2)`
//	return nil
//}

func (srf *StkRefund) SearchStkRefundByDocNo(db *sqlx.DB, docno string) error {
	fmt.Println("StkRefund")
	sql := `set dateformat dmy     select DocNo,isnull(a.TaxNo,'') as TaxNo,DocDate,isnull(a.CreatorCode,'') as CreatorCode,isnull(a.CreateDateTime,'') as CreateDateTime,isnull(a.LastEditorCode,'') as LastEditorCode,isnull(a.LastEditDateT,'') as LastEditDateT,isnull(a.DueDate,'') as DueDate,isnull(a.TaxType,'') as TaxType,isnull(a.ApCode,'') as ApCode,isnull(b.name1,'') as ApName,isnull(a.DepartCode,'') as DepartCode,isnull(a.TaxRate,'') as TaxRate,isnull(IsConfirm,'') as IsConfirm,isnull(MyDescription,'') as MyDescription,SumOfItemAmount,SumOldAmount,SumTrueAmount,SumofDiffAmount,isnull(DiscountWord,'') as DiscountWord,DiscountAmount,SumofBeforeTax,SumOfTaxAmount,SumOfTotalTax,SumOfExceptTax,SumOfZeroTax,SumOfWTax,NetDebtAmount,SumExchangeProfit,BillBalance,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,isnull(GLFormat,'') as GLFormat,GLStartPosting,IsCancel,IsCompleteSave,ReturnMoney,BillType,isnull(CauseType,0) as CauseType,isnull(CauseCode,'') as CauseCode,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(BillGroup,'') as BillGroup,isnull(RecurName,'') as RecurName,isnull(a.ConfirmCode,'') as ConfirmCode,isnull(a.ConfirmDateTime,'') as ConfirmDateTime,isnull(a.CancelCode,'') as CancelCode,isnull(a.CancelDateTime,'') as CancelDateTime,PayBillAmount from dbo.BCStkRefund a left join dbo.bcap b with (nolock) on a.apcode = b.code where a.docno = ?`
	err := db.Get(srf, sql, docno)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//sqlsub := `MyType, DocNo, TaxNo, TaxType, ItemCode, DocDate, ArCode, DepartCode, SaleCode, CashierCode, MyDescription, ItemName, WHCode, ShelfCode, DiscQty, TempQty, BillQty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost, UnitCode, InvoiceNo, ItemType, ExceptTax, IsPos, IsCancel, LineNumber, RefLineNumber, BarCode,AVERAGECOST, LotNumber, PackingRate1, PackingRate2`
	sqlsub := `set dateformat dmy     select MyType,isnull(ItemCode,'') as ItemCode,isnull(MyDescription,'') as MyDescription,isnull(WHCode,'') as WHCode,isnull(ShelfCode,'') as ShelfCode,isnull(DiscQty,0) as DiscQty,isnull(TempQty,0) as TempQty,isnull(BillQty,0) as BillQty,isnull(Price,0) as Price,isnull(DiscountWord,'') as DiscountWord,DiscountAmount,Amount,NetAmount,HomeAmount,isnull(UnitCode,'') as UnitCode,IsCancel,LineNumber,RefLineNumber,isnull(BarCode,'') as BarCode,AVERAGECOST,SumOfCost,isnull(LotNumber,'') as LotNumber,isnull(ItemName,'') as ItemName,isnull(PackingRate1,1) as PackingRate1,isnull(PackingRate2,1) as PackingRate2 from dbo.bcstkrefundsub with (nolock) where docno = ?`
	err = db.Select(&srf.Subs, sqlsub, docno)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (srf *StkRefund) SearchStkRefundByKeyword(db *sqlx.DB, keyword string) (srfs []*StkRefund, err error) {
	fmt.Println("StkRefund")
	sql := `set dateformat dmy     select DocNo,isnull(a.TaxNo,'') as TaxNo,DocDate,isnull(a.CreatorCode,'') as CreatorCode,isnull(a.CreateDateTime,'') as CreateDateTime,isnull(a.LastEditorCode,'') as LastEditorCode,isnull(a.LastEditDateT,'') as LastEditDateT,isnull(a.DueDate,'') as DueDate,isnull(a.TaxType,'') as TaxType,isnull(a.ApCode,'') as ApCode,isnull(b.name1,'') as ApName,isnull(a.DepartCode,'') as DepartCode,isnull(a.TaxRate,'') as TaxRate,isnull(IsConfirm,'') as IsConfirm,isnull(MyDescription,'') as MyDescription,SumOfItemAmount,SumOldAmount,SumTrueAmount,SumofDiffAmount,isnull(DiscountWord,'') as DiscountWord,DiscountAmount,SumofBeforeTax,SumOfTaxAmount,SumOfTotalTax,SumOfExceptTax,SumOfZeroTax,SumOfWTax,NetDebtAmount,SumExchangeProfit,BillBalance,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,isnull(GLFormat,'') as GLFormat,GLStartPosting,IsCancel,IsCompleteSave,ReturnMoney,BillType,isnull(CauseType,0) as CauseType,isnull(CauseCode,'') as CauseCode,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(BillGroup,'') as BillGroup,isnull(RecurName,'') as RecurName,isnull(a.ConfirmCode,'') as ConfirmCode,isnull(a.ConfirmDateTime,'') as ConfirmDateTime,isnull(a.CancelCode,'') as CancelCode,isnull(a.CancelDateTime,'') as CancelDateTime,PayBillAmount from dbo.BCStkRefund a left join dbo.bcap b with (nolock) on a.apcode = b.code where (a.docno  like '%'+?+'%' or a.apcode like '%'+?+'%' or b.name1 like '%'+?+'%' ) order by docno`
	fmt.Println("StkRefund sql =", sql, keyword)
	err = db.Select(&srfs, sql, keyword, keyword, keyword)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//sqlsub := `MyType, DocNo, TaxNo, TaxType, ItemCode, DocDate, ArCode, DepartCode, SaleCode, CashierCode, MyDescription, ItemName, WHCode, ShelfCode, DiscQty, TempQty, BillQty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost, UnitCode, InvoiceNo, ItemType, ExceptTax, IsPos, IsCancel, LineNumber, RefLineNumber, BarCode,AVERAGECOST, LotNumber, PackingRate1, PackingRate2`
	for _, sub := range srfs {
		sqlsub := `set dateformat dmy     select MyType,isnull(ItemCode,'') as ItemCode,isnull(MyDescription,'') as MyDescription,isnull(WHCode,'') as WHCode,isnull(ShelfCode,'') as ShelfCode,isnull(DiscQty,0) as DiscQty,isnull(TempQty,0) as TempQty,isnull(BillQty,0) as BillQty,isnull(Price,0) as Price,isnull(DiscountWord,'') as DiscountWord,DiscountAmount,Amount,NetAmount,HomeAmount,isnull(UnitCode,'') as UnitCode,IsCancel,LineNumber,RefLineNumber,isnull(BarCode,'') as BarCode,AVERAGECOST,SumOfCost,isnull(LotNumber,'') as LotNumber,isnull(ItemName,'') as ItemName,isnull(PackingRate1,1) as PackingRate1,isnull(PackingRate2,1) as PackingRate2 from dbo.bcstkrefundsub with (nolock) where docno = ?`
		fmt.Println("StkRefund sqlsub =", sqlsub, sub.DocNo)
		err = db.Select(&sub.Subs, sqlsub, sub.DocNo)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	return srfs, nil
}
