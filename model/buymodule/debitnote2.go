package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"errors"
	m "github.com/loukmho/bcaccount_api/model"
)

type DebitNote2 struct {
	SaveFrom          int        `json:"save_from" db:"SaveFrom"`
	Source            int        `json:"source" db:"Source"`
	DocNo             string     `json:"doc_no" db:"DocNo"`
	DbnInPutTax
	DbnVendor
	DocDate           string     `json:"doc_date" db:"DocDate"`
	CreatorCode       string     `json:"creator_code" db:"CreatorCode"`
	CreateDateTime    string     `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode    string     `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT     string     `json:"last_edit_date_t" db:"LastEditDateT"`
	CreditDay         int        `json:"credit_day" db:"CreditDay"`
	DueDate           string     `json:"due_date" db:"DueDate"`
	TaxType           int        `json:"tax_type" db:"TaxType"`
	DepartCode        string     `json:"depart_code" db:"DepartCode"`
	TaxRate           float64    `json:"tax_rate" db:"TaxRate"`
	IsConfirm         int        `json:"is_confirm" db:"IsConfirm"`
	MyDescription     string     `json:"my_description" db:"MyDescription"`
	SumOfItemAmount   float64    `json:"sum_of_item_amount" db:"SumOfItemAmount"`
	SumOldAmount      float64    `json:"sum_old_amount" db:"SumOldAmount"`
	SumTrueAmount     float64    `json:"sum_true_amount" db:"SumTrueAmount"`
	SumofDiffAmount   float64    `json:"sumof_diff_amount" db:"SumofDiffAmount"`
	SumofBeforeTax    float64    `json:"sumof_before_tax" db:"SumofBeforeTax"`
	SumOfTaxAmount    float64    `json:"sum_of_tax_amount" db:"SumOfTaxAmount"`
	SumOfTotalTax     float64    `json:"sum_of_total_tax" db:"SumOfTotalTax"`
	SumOfExceptTax    float64    `json:"sum_of_except_tax" db:"SumOfExceptTax"`
	SumOfZeroTax      float64    `json:"sum_of_zero_tax" db:"SumOfZeroTax"`
	SumOfWTax         float64    `json:"sum_of_w_tax" db:"SumOfWTax"`
	DiscountWord      string     `json:"discount_word" db:"DiscountWord"`
	DiscountAmount    float64    `json:"discount_amount" db:"DiscountAmount"`
	NetDebtAmount     float64    `json:"net_debt_amount" db:"NetDebtAmount"`
	SumExchangeProfit float64    `json:"sum_exchange_profit" db:"SumExchangeProfit"`
	BillBalance       float64    `json:"bill_balance" db:"BillBalance"`
	CurrencyCode      string     `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate      float64    `json:"exchange_rate" db:"ExchangeRate"`
	GLFormat          string     `json:"gl_format" db:"GLFormat"`
	GLStartPosting    int        `json:"gl_start_posting" db:"GLStartPosting"`
	IsPostGL          int        `json:"is_post_gl" db:"IsPostGL"`
	IsCancel          int        `json:"is_cancel" db:"IsCancel"`
	IsCompleteSave    int        `json:"is_complete_save" db:"IsCompleteSave"`
	ReturnStatus      int        `json:"return_status" db:"ReturnStatus"`
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
	Subs              []*DbnItem `json:"subs"`
}

type DbnVendor struct {
	ApCode string `json:"ap_code" db:"ApCode"`
	ApName string `json:"ap_name" db:"ApName"`
}

type DbnInPutTax struct {
	TaxNo    string `json:"tax_no" db:"TaxNo"`
	BookCode string `json:"book_code" db:"BookCode"`
}

type DbnItem struct {
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
	ExceptTax      int     `json:"except_tax" db:"ExceptTax"`
	IsCancel       int     `json:"is_cancel" db:"IsCancel"`
	LineNumber     int     `json:"line_number" db:"LineNumber"`
	RefLineNumber  int     `json:"ref_line_number" db:"RefLineNumber"`
	BarCode        string  `json:"bar_code" db:"BarCode"`
	AverageCost    float64 `json:"averagecost" db:"AverageCost"`
	LotNumber      string  `json:"lot_number" db:"LotNumber"`
	PackingRate1   float64 `json:"packing_rate_1" db:"PackingRate1"`
	PackingRate2   float64 `json:"packing_rate_2" db:"PackingRate2"`
}

func InsertAndEditDebitNote2() {
	//sql := `insert into dbo.bcdebitnote2(DocNo, TaxNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, CreditDay, DueDate, TaxType, ApCode, DepartCode, TaxRate, IsConfirm, MyDescription, SumOfItemAmount, SumOldAmount, SumTrueAmount, SumofDiffAmount, SumofBeforeTax, SumOfTaxAmount, SumOfTotalTax, SumOfExceptTax, SumOfZeroTax, SumOfWTax, DiscountWord, DiscountAmount, NetDebtAmount, SumExchangeProfit, BillBalance, CurrencyCode, ExchangeRate, GLFormat, GLStartPosting, IsPostGL, IsCancel,  IsCompleteSave, ReturnStatus, CauseType, CauseCode, AllocateCode, ProjectCode, BillGroup, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime, PayBillAmount)`
	//sqlsub := `insert into dbo.BCDebitNoteSub2(MyType, DocNo, TaxNo,  TaxType, ItemCode, ItemType, DocDate, ApCode, DepartCode, MyDescription, ItemName, WHCode, ShelfCode, DiscQty, TempQty, BillQty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost,UnitCode,  InvoiceNo, ExceptTax,IsCancel, LineNumber, RefLineNumber,BarCode,AVERAGECOST, LotNumber, PackingRate1, PackingRate2) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
}

func (dbn *DebitNote2) InsertAndEditDebitNote2(db *sqlx.DB) error {
	var check_exist int
	var count_item int

	for _, sub_item := range dbn.Subs {
		if (sub_item.DiscQty != 0) {
			count_item = count_item + 1
		}
	}

	sqlexist := `select count(docno) as check_exist from dbo.bcdebitnote2 where docno = ?`
	err := db.Get(&check_exist, sqlexist, dbn.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println("Docno = ", dbn.DocNo)
	switch {
	case dbn.DocNo == "":
		return errors.New("Docno is null")
	case dbn.ApCode == "":
		return errors.New("Apcode is null")
	case dbn.DocDate == "":
		return errors.New("Docdate is null")
	case dbn.IsCancel == 1:
		return errors.New("Docno is cancel")
	case dbn.IsConfirm == 1:
		return errors.New("Docno is confirm")
	case dbn.SumOldAmount == 0:
		return errors.New("Docno not have old bill amount")
	case dbn.SumOfTotalTax == 0:
		return errors.New("Docno not have amount return")
	case dbn.SumOfItemAmount == 0 && count_item == 0:
		return errors.New("Docno not have SumOfItemAmount")
	}

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")

	fmt.Println("TaxRate = ", def.TaxRateDefault)

	if dbn.TaxRate == 0 {
		dbn.TaxRate = def.TaxRateDefault
	}
	if dbn.ExchangeRate == 0 {
		dbn.ExchangeRate = def.ExchangeRateDefault
	}

	if dbn.SaveFrom == 0 {
		//dbn.SaveFrom = def.Debitnote2SaveFrom
	}

	dbn.SumofBeforeTax, dbn.SumOfTaxAmount = m.CalcTaxCredit(dbn.TaxType, dbn.TaxRate, dbn.SumOfTotalTax)
	fmt.Println("Befor,Tax", dbn.SumofBeforeTax, dbn.SumOfTaxAmount)

	dbn.BillBalance = dbn.SumOfTotalTax

	dbn.NetDebtAmount = dbn.SumOfTotalTax
	dbn.SumTrueAmount = dbn.SumOldAmount - dbn.SumOfTotalTax

	if dbn.TaxType == 1 {
		dbn.PayBillAmount = dbn.SumOfTotalTax
	}

	if (dbn.Source == 0) {
		//dbn.Source = def.Debitnote2Source
	}
	if (dbn.GLFormat == "") {
		//dbn.GLFormat = def.Debitnote2GLFormat
	}
	if (dbn.TaxNo == "") {
		dbn.TaxNo = dbn.DocNo
	}

	fmt.Println("UserCode = ", dbn.UserCode)

	if (dbn.IsCompleteSave == 0) {
		dbn.IsCompleteSave = 1
	}

	if check_exist == 0 {
		dbn.CreatorCode = dbn.UserCode
		sql := `insert into dbo.bcdebitnote2(DocNo,TaxNo,DocDate,CreatorCode,CreateDateTime,CreditDay,DueDate,TaxType,ApCode,DepartCode,TaxRate,IsConfirm,MyDescription,SumOfItemAmount,SumOldAmount,SumTrueAmount,SumofDiffAmount,SumofBeforeTax,SumOfTaxAmount,SumOfTotalTax,SumOfExceptTax,SumOfZeroTax,SumOfWTax,DiscountWord,DiscountAmount,NetDebtAmount,SumExchangeProfit,BillBalance,CurrencyCode,ExchangeRate,GLFormat,GLStartPosting,IsPostGL,IsCancel,IsCompleteSave,ReturnStatus,CauseType,CauseCode,AllocateCode,ProjectCode,BillGroup,RecurName,PayBillAmount) values(?,?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sql, dbn.DocNo,dbn.TaxNo,dbn.DocDate,dbn.CreatorCode,dbn.CreditDay,dbn.DueDate,dbn.TaxType,dbn.ApCode,dbn.DepartCode,dbn.TaxRate,dbn.IsConfirm,dbn.MyDescription,dbn.SumOfItemAmount,dbn.SumOldAmount,dbn.SumTrueAmount,dbn.SumofDiffAmount,dbn.SumofBeforeTax,dbn.SumOfTaxAmount,dbn.SumOfTotalTax,dbn.SumOfExceptTax,dbn.SumOfZeroTax,dbn.SumOfWTax,dbn.DiscountWord,dbn.DiscountAmount,dbn.NetDebtAmount,dbn.SumExchangeProfit,dbn.BillBalance,dbn.CurrencyCode,dbn.ExchangeRate,dbn.GLFormat,dbn.GLStartPosting,dbn.IsPostGL,dbn.IsCancel,dbn.IsCompleteSave,dbn.ReturnStatus,dbn.CauseType,dbn.CauseCode,dbn.AllocateCode,dbn.ProjectCode,dbn.BillGroup,dbn.RecurName,dbn.PayBillAmount)
		if err != nil {
			fmt.Println("Insert Credit =", err.Error())
			return err
		}
	} else {
		dbn.LastEditorCode = dbn.UserCode

		//sql := `update dbo.bcstkrefund set TaxNo=?,DocDate=?,CreatorCode=?,LastEditDateT=getdate(),DueDate=?,TaxType=?,ApCode=?,DepartCode=?,TaxRate=?,IsConfirm=?,MyDescription=?,SumOfItemAmount=?,SumOldAmount=?,SumTrueAmount=?,SumofDiffAmount=?,DiscountWord=?,DiscountAmount=?,SumofBeforeTax=?,SumOfTaxAmount=?,SumOfTotalTax=?,SumOfExceptTax=?,SumOfZeroTax=?,SumOfWTax=?,NetDebtAmount=?,SumExchangeProfit=?,BillBalance=?,CurrencyCode=?,ExchangeRate=?,GLFormat=?, GLStartPosting=?,IsCancel=?,IsCompleteSave=?,ReturnMoney=?,BillType=?,CauseType=?,CauseCode=?,AllocateCode=?,ProjectCode=?,BillGroup=?,RecurName=?,PayBillAmount=? where docno = ? `
		//_, err = db.Exec(sql, srf.TaxNo, srf.DocDate, srf.LastEditorCode, srf.DueDate, srf.TaxType, srf.ApCode, srf.DepartCode, srf.TaxRate, srf.IsConfirm, srf.MyDescription, srf.SumOfItemAmount, srf.SumOldAmount, srf.SumTrueAmount, srf.SumofDiffAmount, srf.DiscountWord, srf.DiscountAmount, srf.SumofBeforeTax, srf.SumOfTaxAmount, srf.SumOfTotalTax, srf.SumOfExceptTax, srf.SumOfZeroTax, srf.SumOfWTax, srf.NetDebtAmount, srf.SumExchangeProfit, srf.BillBalance, srf.CurrencyCode, srf.ExchangeRate, srf.GLFormat, srf.GLStartPosting, srf.IsCancel, srf.IsCompleteSave, srf.ReturnMoney, srf.BillType, srf.CauseType, srf.CauseCode, srf.AllocateCode, srf.ProjectCode, srf.BillGroup, srf.RecurName, srf.PayBillAmount, srf.DocNo)
		if err != nil {
			fmt.Println("Insert Credit =", err.Error())
			return err
		}
	}

	sql_del_sub := `delete dbo.bcdebitnotesub2 where docno = ?`
	_, err = db.Exec(sql_del_sub, dbn.DocNo)
	if err != nil {
		return err
	}

	var vLineNumber int

	for _, item := range dbn.Subs {
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
		case dbn.TaxType == 0:
			item.HomeAmount = item.Amount
			item.NetAmount = item.Amount
		case dbn.TaxType == 1:
//			taxamount := m.ToFixed(item.Amount-((item.Amount*100)/(100+float64(srf.TaxRate))), 2)
//			beforetaxamount := m.ToFixed(item.Amount-taxamount, 2)
			//item.HomeAmount = beforetaxamount
			//item.NetAmount = beforetaxamount
		case dbn.TaxType == 2:
			item.HomeAmount = item.Amount
			item.NetAmount = item.Amount
		}

		//sqlsub := ` insert into dbo.BCStkRefundSub(MyType,DocNo,TaxNo,TaxType,ItemType,ItemCode,DocDate,ApCode,DepartCode,MyDescription,WHCode,ShelfCode,DiscQty,TempQty,BillQty,Price,DiscountWord,DiscountAmount,Amount,NetAmount,HomeAmount,UnitCode,IsCancel,LineNumber,RefLineNumber,BarCode,AVERAGECOST,SumOfCost,LotNumber,ItemName,TaxRate,PackingRate1,PackingRate2) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		//_, err = db.Exec(sqlsub, item.MyType, srf.DocNo, srf.TaxNo, srf.TaxType, item.ItemType, item.ItemCode, srf.DocDate, srf.ApCode, srf.DepartCode, item.MyDescription, item.WHCode, item.ShelfCode, item.DiscQty, item.TempQty, item.BillQty, item.Price, item.DiscountWord, item.DiscountAmount, item.Amount, item.NetAmount, item.HomeAmount, item.UnitCode, item.IsCancel, item.LineNumber, item.RefLineNumber, item.BarCode, item.AverageCost, item.SumOfCost, item.LotNumber, item.ItemName, srf.TaxRate, item.PackingRate1, item.PackingRate2)
		//fmt.Println("sqltax = ", sqlsub)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		if (item.ItemType != 1) {
			sqlprocess := ` insert into dbo.ProcessStock (ItemCode,ProcessFlag,FlowStatus) values(?, 1, 0)`
			_, err = db.Exec(sqlprocess, item.ItemCode)
//			fmt.Println("sqlprocess = ", sqlsub)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		vLineNumber = vLineNumber + 1
	}

	return nil
}

func (dbn *DebitNote2) SearchDebitNote2ByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select DocNo, isnull(a.TaxNo,'') as TaxNo, DocDate, isnull(a.CreatorCode,'') as CreatorCode, isnull(a.CreateDateTime,'') as CreateDateTime, isnull(a.LastEditorCode,'') as LastEditorCode, isnull(a.LastEditDateT,'') as LastEditDateT, a.CreditDay, isnull(DueDate,'') as DueDate, a.TaxType, isnull(ApCode,'') as ApCode, isnull(b.name1,'') as ApName, isnull(DepartCode,'') as DepartCode, TaxRate, IsConfirm, isnull(MyDescription,'') as MyDescription, SumOfItemAmount, SumOldAmount, SumTrueAmount, SumofDiffAmount, SumofBeforeTax, SumOfTaxAmount, SumOfTotalTax, SumOfExceptTax, SumOfZeroTax, SumOfWTax, isnull(DiscountWord,'') as DiscountWord, DiscountAmount, NetDebtAmount, SumExchangeProfit, BillBalance, isnull(CurrencyCode,'') as CurrencyCode, ExchangeRate, isnull(GLFormat,'') as GLFormat, GLStartPosting, IsPostGL, IsCancel, IsCompleteSave, ReturnStatus, isnull(CauseType,0) as CauseType, isnull(CauseCode,'') as CauseCode, isnull(AllocateCode,'') as AllocateCode, isnull(ProjectCode,'') as ProjectCode, isnull(BillGroup,'') as BillGroup, isnull(RecurName,'') as RecurName, isnull(a.ConfirmCode,'') as ConfirmCode, isnull(a.ConfirmDateTime,'') as ConfirmDateTime, isnull(a.CancelCode,'') as CancelCode, isnull(a.CancelDateTime,'') as CancelDateTime, PayBillAmount  from dbo.bcdebitnote2 a with (nolock)  left join dbo.bcap b with (nolock) on a.apcode = b.code where a.docno = ? order by a.docno`
	err := db.Get(dbn, sql, docno)
	if err != nil {
		fmt.Println(err)
		return err
	}

	sqlsub := `set dateformat dmy     select MyType, isnull(ItemCode,'') as ItemCode, isnull(ItemType,0) as ItemType, isnull(MyDescription,'') as MyDescription, isnull(ItemName,'') as ItemName, isnull(WHCode,'') as WHCode, isnull(ShelfCode,'') as ShelfCode, isnull(DiscQty,0) as DiscQty, isnull(TempQty,0) as TempQty, isnull(BillQty,0) as BillQty, isnull(Price,0) as Price, isnull(DiscountWord,'') as DiscountWord, isnull(DiscountAmount,0) as DiscountAmount, isnull(Amount,0) as Amount, isnull(NetAmount,0) as NetAmount, isnull(HomeAmount,0) as HomeAmount, isnull(SumOfCost,0) as SumOfCost, isnull(UnitCode,'') as UnitCode, isnull(InvoiceNo,'') as InvoiceNo, isnull(IsCancel,0) as IsCancel, isnull(LineNumber,0) as LineNumber, isnull(RefLineNumber,0) as RefLineNumber, isnull(BarCode,'') as BarCode, isnull(AVERAGECOST,0) as AverageCost, isnull(LotNumber,'') as LotNumber, isnull(PackingRate1,1) as PackingRate1, isnull(PackingRate2,1) as PackingRate2 from dbo.BCDebitNoteSub2 a With (Nolock) where a.docno = ?`
	err = db.Select(&dbn.Subs, sqlsub, docno)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//sqlsub := `MyType, DocNo, TaxNo,  TaxType, ItemCode, ItemType, DocDate, ApCode, DepartCode, MyDescription, ItemName, WHCode, ShelfCode, DiscQty, TempQty, BillQty, Price, DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost,UnitCode,  InvoiceNo, ExceptTax,IsCancel, LineNumber, RefLineNumber,BarCode,AVERAGECOST, LotNumber, PackingRate1, PackingRate2`
	return nil
}

func (dbn *DebitNote2) SearchDebitnote2ByKeyword(db *sqlx.DB, keyword string) (dbns []*DebitNote2, err error) {
	fmt.Println("SearchByKeyword")
	sql := `set dateformat dmy     select DocNo, isnull(a.TaxNo,'') as TaxNo, DocDate, isnull(a.CreatorCode,'') as CreatorCode, isnull(a.CreateDateTime,'') as CreateDateTime, isnull(a.LastEditorCode,'') as LastEditorCode, isnull(a.LastEditDateT,'') as LastEditDateT, a.CreditDay, isnull(DueDate,'') as DueDate, a.TaxType, isnull(ApCode,'') as ApCode, isnull(b.name1,'') as ApName, isnull(DepartCode,'') as DepartCode, TaxRate, IsConfirm, isnull(MyDescription,'') as MyDescription, SumOfItemAmount, SumOldAmount, SumTrueAmount, SumofDiffAmount, SumofBeforeTax, SumOfTaxAmount, SumOfTotalTax, SumOfExceptTax, SumOfZeroTax, SumOfWTax, isnull(DiscountWord,'') as DiscountWord, DiscountAmount, NetDebtAmount, SumExchangeProfit, BillBalance, isnull(CurrencyCode,'') as CurrencyCode, ExchangeRate, isnull(GLFormat,'') as GLFormat, GLStartPosting, IsPostGL, IsCancel, IsCompleteSave, ReturnStatus, isnull(CauseType,0) as CauseType, isnull(CauseCode,'') as CauseCode, isnull(AllocateCode,'') as AllocateCode, isnull(ProjectCode,'') as ProjectCode, isnull(BillGroup,'') as BillGroup, isnull(RecurName,'') as RecurName, isnull(a.ConfirmCode,'') as ConfirmCode, isnull(a.ConfirmDateTime,'') as ConfirmDateTime, isnull(a.CancelCode,'') as CancelCode, isnull(a.CancelDateTime,'') as CancelDateTime, PayBillAmount  from dbo.bcdebitnote2 a with (nolock)  left join dbo.bcap b with (nolock) on a.apcode = b.code where (a.docno  like '%'+?+'%' or a.apcode like '%'+?+'%' or b.name1 like '%'+?+'%' ) order by docno`
	err = db.Select(&dbns, sql, keyword, keyword, keyword)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, sub := range dbns {
		fmt.Println("Docno = ", sub.DocNo)
		sqlsub := `set dateformat dmy     select MyType, isnull(ItemCode,'') as ItemCode, isnull(ItemType,0) as ItemType, isnull(MyDescription,'') as MyDescription, isnull(ItemName,'') as ItemName, isnull(WHCode,'') as WHCode, isnull(ShelfCode,'') as ShelfCode, isnull(DiscQty,0) as DiscQty, isnull(TempQty,0) as TempQty, isnull(BillQty,0) as BillQty, isnull(Price,0) as Price, isnull(DiscountWord,'') as DiscountWord, isnull(DiscountAmount,0) as DiscountAmount, isnull(Amount,0) as Amount, isnull(NetAmount,0) as NetAmount, isnull(HomeAmount,0) as HomeAmount, isnull(SumOfCost,0) as SumOfCost, isnull(UnitCode,'') as UnitCode, isnull(InvoiceNo,'') as InvoiceNo, isnull(IsCancel,0) as IsCancel, isnull(LineNumber,0) as LineNumber, isnull(RefLineNumber,0) as RefLineNumber, isnull(BarCode,'') as BarCode, isnull(AVERAGECOST,0) as AverageCost, isnull(LotNumber,'') as LotNumber, isnull(PackingRate1,1) as PackingRate1, isnull(PackingRate2,1) as PackingRate2 from dbo.BCDebitNoteSub2 a With (Nolock) where a.docno = ? order by linenumber`
		err = db.Select(&sub.Subs, sqlsub, sub.DocNo)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	return dbns, nil
}
