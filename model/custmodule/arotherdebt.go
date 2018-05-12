package model

import "github.com/jmoiron/sqlx"

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
}

func (otd *ArOtherDebt)InsertAndEditArOtherDebt(db *sqlx.DB) error{
	sql := `set dateformat dmy     insert dbo.BCArOtherDebt(DocNo,DocDate,ArCode,GLBookCode,SumofDebit,SumofCredit,DepartCode,CreditDay,DueDate,PayBillDate,SaleCode,IsConfirm,PayBillStatus,MyDescription,BillGroup,ContactCode,NetDebtAmount,BillBalance,CurrencyCode,ExchangeRate,IsCancel,AllocateCode,ProjectCode,RecurName,CreatorCode,CreateDateTime,PayBillAmount,BillTemporary) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,getdate(),?,?)`
	db.Exec(sql, otd.DocNo,otd.DocDate,otd.ArCode,otd.GLBookCode,otd.SumofDebit,otd.SumofCredit,otd.DepartCode,otd.CreditDay,otd.DueDate,otd.PayBillDate,otd.SaleCode,otd.IsConfirm,otd.PayBillStatus,otd.MyDescription,otd.BillGroup,otd.ContactCode,otd.NetDebtAmount,otd.BillBalance,otd.CurrencyCode,otd.ExchangeRate,otd.IsCancel,otd.AllocateCode,otd.ProjectCode,otd.RecurName,otd.CreatorCode,otd.PayBillAmount,otd.BillTemporary)

	return nil
}

func (otd *ArOtherDebt) SearchArOtherDebtByDocNo() error {
	//sql := `DocNo, DocDate, ArCode, GLBookCode, SumofDebit, SumofCredit, DepartCode, CreditDay, DueDate, PayBillDate, SaleCode, IsConfirm, PayBillStatus, MyDescription, BillGroup,ContactCode, NetDebtAmount, BillBalance, CurrencyCode, ExchangeRate, IsCancel, AllocateCode, ProjectCode, RecurName, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, ConfirmCode, ConfirmDateTime,CancelCode, CancelDateTime, PayBillAmount, BillTemporary`
	return nil
}
