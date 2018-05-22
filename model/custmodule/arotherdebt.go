package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"errors"
	"time"
	m "github.com/loukmho/bcaccount_api/model"
)

type ArOtherDebt struct {
	DocNo           string  `json:"doc_no" db:"DocNo"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	ArCode          string  `json:"ar_code" db:"ArCode"`
	GLBookCode      string  `json:"gl_book_code" db:"GLBookCode"`
	SumofDebit      float64 `json:"sumof_debit" db:"SumofDebit"`
	SumofCredit     float64 `json:"sumof_credit" db:"SumofCredit"`
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	CreditDay       int     `json:"credit_day" db:"CreditDay"`
	DueDate         string  `json:"due_date" db:"DueDate"`
	PayBillDate     string  `json:"pay_bill_date" db:"PayBillDate"`
	SaleCode        string  `json:"sale_code" db:"SaleCode"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	PayBillStatus   int     `json:"pay_bill_status" db:"PayBillStatus"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	BillGroup       string  `json:"bill_group" db:"BillGroup"`
	ContactCode     string  `json:"contact_code" db:"ContactCode"`
	NetDebtAmount   float64 `json:"net_debt_amount" db:"NetDebtAmount"`
	BillBalance     float64 `json:"bill_balance" db:"BillBalance"`
	CurrencyCode    string  `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate    float64 `json:"exchange_rate" db:"ExchangeRate"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	AllocateCode    string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string  `json:"project_code" db: "ProjectCode"`
	RecurName       string  `json:"recur_name" db:"RecurName"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	ConfirmCode     string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string  `json:"cancel_date_time" db:"CancelDateTime"`
	PayBillAmount   float64 `json:"pay_bill_amount" db:"PayBillAmount"`
	BillTemporary   float64 `json:"bill_temporary" db:"BillTemporary"`
	UserCode        string  `json:"user_code"`
}

func (otd *ArOtherDebt) InsertAndEditArOtherDebt(db *sqlx.DB) error {
	var check_exist int

	now := time.Now()

	sqlexist := `select count(docno) as check_exist from dbo.bcarotherdebt where docno = ?`
	err := db.Get(&check_exist, sqlexist, otd.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil
	}

	switch {
	case otd.DocNo == "":
		return errors.New("Docno is null")
	case otd.ArCode == "":
		return errors.New("Arcode is null")
	case otd.NetDebtAmount == 0:
		return errors.New("Totalamount is 0")
	case otd.DocDate == "":
		return errors.New("Docdate is null")
	case otd.SumofDebit == 0 :
		return errors.New("Docno not have debit")
	case otd.SumofCredit == 0 :
		return errors.New("Docno not have credit")
	case otd.IsCancel == 1:
		return errors.New("Docno is cancel can not edit data")
	case otd.IsConfirm == 1:
		return errors.New("Docno is confirm can not edit data")
	}

	if (otd.DueDate == "" && otd.CreditDay == 0) {
		otd.DueDate = otd.DocDate
	} else {
		otd.DueDate = now.AddDate(0, 0, otd.CreditDay).Format("2006-01-02")
	}

	fmt.Println("UserCode = ", otd.UserCode)

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")

	if otd.ExchangeRate == 0 {
		otd.ExchangeRate = def.ExchangeRateDefault
	}


	if (check_exist == 0){
		otd.CreatorCode = otd.UserCode

		sql := `set dateformat dmy     insert dbo.BCArOtherDebt(DocNo,DocDate,ArCode,GLBookCode,SumofDebit,SumofCredit,DepartCode,CreditDay,DueDate,PayBillDate,SaleCode,IsConfirm,PayBillStatus,MyDescription,BillGroup,ContactCode,NetDebtAmount,BillBalance,CurrencyCode,ExchangeRate,IsCancel,AllocateCode,ProjectCode,RecurName,CreatorCode,CreateDateTime,PayBillAmount,BillTemporary) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,getdate(),?,?)`
		_, err = db.Exec(sql, otd.DocNo, otd.DocDate, otd.ArCode, otd.GLBookCode, otd.SumofDebit, otd.SumofCredit, otd.DepartCode, otd.CreditDay, otd.DueDate, otd.PayBillDate, otd.SaleCode, otd.IsConfirm, otd.PayBillStatus, otd.MyDescription, otd.BillGroup, otd.ContactCode, otd.NetDebtAmount, otd.BillBalance, otd.CurrencyCode, otd.ExchangeRate, otd.IsCancel, otd.AllocateCode, otd.ProjectCode, otd.RecurName, otd.CreatorCode, otd.PayBillAmount, otd.BillTemporary)
		if err != nil {
			return err
		}

	}else{
		otd.LastEditorCode = otd.UserCode

		sql := `set dateformat dmy     update dbo.BCArOtherDebt set DocDate=?,ArCode=?,GLBookCode=?,SumofDebit=?,SumofCredit=?,DepartCode=?,CreditDay=?,DueDate=?,PayBillDate=?,SaleCode=?,IsConfirm=?,PayBillStatus=?,MyDescription=?,BillGroup=?,ContactCode=?,NetDebtAmount=?,BillBalance=?,CurrencyCode=?,ExchangeRate=?,IsCancel=?,AllocateCode=?,ProjectCode=?,RecurName=?,LastEditorCode=?,LastEditDateT=getdate(),PayBillAmount=?,BillTemporary=? where docno = ?`
		_, err = db.Exec(sql, otd.DocDate, otd.ArCode, otd.GLBookCode, otd.SumofDebit, otd.SumofCredit, otd.DepartCode, otd.CreditDay, otd.DueDate, otd.PayBillDate, otd.SaleCode, otd.IsConfirm, otd.PayBillStatus, otd.MyDescription, otd.BillGroup, otd.ContactCode, otd.NetDebtAmount, otd.BillBalance, otd.CurrencyCode, otd.ExchangeRate, otd.IsCancel, otd.AllocateCode, otd.ProjectCode, otd.RecurName, otd.CreatorCode, otd.PayBillAmount, otd.BillTemporary, otd.DocNo)
		if err != nil {
			return err
		}
	}

	return nil
}

func (otd *ArOtherDebt) SearchArOtherDebtByDocNo(db *sqlx.DB, docno string) error {
	//sql := `DocNo, DocDate, ArCode, GLBookCode, SumofDebit, SumofCredit, DepartCode, CreditDay, DueDate, PayBillDate, SaleCode, IsConfirm, PayBillStatus, MyDescription, BillGroup,ContactCode, NetDebtAmount, BillBalance, CurrencyCode, ExchangeRate, IsCancel, AllocateCode, ProjectCode, RecurName, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, ConfirmCode, ConfirmDateTime,CancelCode, CancelDateTime, PayBillAmount, BillTemporary`
	sql := `set dateformat dmy     select DocNo,DocDate,isnull(ArCode,'') as ArCode,isnull(GLBookCode,'') as GLBookCode,SumofDebit,SumofCredit,isnull(DepartCode,'') as DepartCode,CreditDay,isnull(DueDate,'') as DueDate,isnull(PayBillDate,'') as PayBillDate,isnull(SaleCode,'') as SaleCode,IsConfirm,PayBillStatus,isnull(MyDescription,'') as MyDescription,isnull(BillGroup,'') as BillGroup,isnull(ContactCode,'') as ContactCode,NetDebtAmount,BillBalance,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,IsCancel,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(RecurName,'') as RecurName,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime,PayBillAmount, BillTemporary from dbo.BCArOtherDebt with (nolock) where docno = ?`
	err := db.Get(otd, sql, docno)
	if err != nil {
		return err
	}

	return nil
}

func (otd *ArOtherDebt) SearchArOtherDebtByKeyword(db *sqlx.DB, keyword string) (otds []*ArOtherDebt, err error) {
	//sql := `DocNo, DocDate, ArCode, GLBookCode, SumofDebit, SumofCredit, DepartCode, CreditDay, DueDate, PayBillDate, SaleCode, IsConfirm, PayBillStatus, MyDescription, BillGroup,ContactCode, NetDebtAmount, BillBalance, CurrencyCode, ExchangeRate, IsCancel, AllocateCode, ProjectCode, RecurName, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, ConfirmCode, ConfirmDateTime,CancelCode, CancelDateTime, PayBillAmount, BillTemporary`
	sql := `set dateformat dmy     select DocNo,DocDate,isnull(ArCode,'') as ArCode,isnull(GLBookCode,'') as GLBookCode,SumofDebit,SumofCredit,isnull(DepartCode,'') as DepartCode,CreditDay,isnull(DueDate,'') as DueDate,isnull(PayBillDate,'') as PayBillDate,isnull(SaleCode,'') as SaleCode,IsConfirm,PayBillStatus,isnull(MyDescription,'') as MyDescription,isnull(BillGroup,'') as BillGroup,isnull(ContactCode,'') as ContactCode,NetDebtAmount,BillBalance,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,IsCancel,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(RecurName,'') as RecurName,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime,PayBillAmount, BillTemporary from dbo.BCArOtherDebt with (nolock) where (docno  like '%'+?+'%' or arcode like '%'+?+'%' or salecode like '%'+?+'%') order by docno`
	err = db.Select(otd, sql, keyword, keyword, keyword)
	if err != nil {
		return nil, err
	}

	return otds, nil
}
