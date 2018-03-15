package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type ArDepositSpecial struct {
	DocNo           string  `json:"doc_no" db:"DocNo"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	ArCode          string  `json:"ar_code" db:"ArCode"`
	DpsCustomer
	DpsSaleMan
	CreatorCode string `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	DepartCode      string  `json:"depart_code"  db:"DepartCode"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	CreditDay       int     `json:"credit_day" db:"CreditDay"`
	DueDate         string  `json:"due_date" db:"DueDate"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	TotalAmount     float64 `json:"total_amount" db:"TotalAmount"`
	SumOfWTax       float64 `json:"sum_of_w_tax" db:"SumOfWTax"`
	BeforeTaxAmount float64 `json:"before_tax_amount" db:"BeforeTaxAmount"`
	NetAmount       float64 `json:"net_amount" db:"NetAmount"`
	BillBalance     float64 `json:"bill_balance" db:"BillBalance"`
	ExcessAmount1   float64 `json:"excess_amount_1" db:"ExcessAmount1"`
	ExcessAmount2   float64 `json:"excess_amount_2" db:"ExcessAmount2"`
	ChargeAmount    float64 `json:"charge_amount" db:"ChargeAmount"`
	ChangeAmount    float64 `json:"change_amount" db:"ChangeAmount"`
	OtherIncome     float64 `json:"other_income" db:"OtherIncome"`
	OtherExpense    float64 `json:"other_expense" db:"OtherExpense"`
	RefNo           string  `json:"ref_no" db:"RefNo"`
	CurrencyCode    string  `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate    float64 `json:"exchange_rate" db:"ExchangeRate"`
	SumCashAmount   float64 `json:"sum_cash_amount" db:"SumCashAmount"`
	SumChqAmount    float64 `json:"sum_chq_amount" db:"SumChqAmount"`
	SumCreditAmount float64 `json:"sum_credit_amount" db:"SumCreditAmount"`
	SumBankAmount   float64 `json:"sum_bank_amount" db:"SumBankAmount"`
	PettyCashAmount float64 `json:"petty_cash_amount" db:"PettyCashAmount"`
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
}

type DpsCustomer struct {
	ArCode string `json:"ar_code" db:"ArCode"`
	ArName string `json:"ar_name" db:"ArName"`
}

type DpsSaleMan struct {
	SaleCode string `json:"sale_code" db:"SaleCode"`
	SaleName string `json:"sale_name" db:"SaleName"`
}

func (dps *ArDepositSpecial) SearchArDepositSpecialByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select a.DocNo,a.DocDate,a.ArCode,isnull(b.name1,'') as ArName,isnull(a.CreatorCode,'') as CreatorCode,isnull(a.CreateDateTime,'') as CreateDateTime,isnull(a.LastEditorCode,'') as LastEditorCode,isnull(a.LastEditDateT,'') as LastEditDateT,isnull(a.DepartCode,'') as DepartCode,isnull(a.SaleCode,'') as SaleCode,isnull(c.name,'') as SaleName,a.IsConfirm,a.CreditDay,isnull(a.DueDate,'') as DueDate,isnull(a.MyDescription,'') as MyDescription,a.TotalAmount,a.SumOfWTax,a.BeforeTaxAmount,a.NetAmount,a.BillBalance,a.ExcessAmount1,a.ExcessAmount2,a.ChargeAmount,a.ChangeAmount,a.OtherIncome,a.OtherExpense,isnull(a.RefNo,'') as RefNo,isnull(a.CurrencyCode,'') as CurrencyCode,a.ExchangeRate,a.SumCashAmount,a.SumChqAmount,a.SumCreditAmount,a.SumBankAmount,a.PettyCashAmount,isnull(a.GLFormat,'') as GLFormat,a.IsCancel,a.IsReturnMoney,isnull(a.AllocateCode,'') as AllocateCode,isnull(a.ProjectCode,'') as ProjectCode,isnull(a.BillGroup,'') as BillGroup,isnull(a.RecurName,'') as RecurName,isnull(a.ConfirmCode,'') as ConfirmCode,isnull(a.ConfirmDateTime,'') as ConfirmDateTime,isnull(a.CancelCode,'') as CancelCode,isnull(a.CancelDateTime,'') as CancelDateTime from dbo.BCArDepositSpecial a WITH (NOLOCK) left join dbo.BCAR b WITH (NOLOCK) on a.arcode = b.code left join dbo.BCSale c WITH (NOLOCK) on a.salecode = c.code where a.docno = ?`
	err := db.Get(dps, sql, docno)
	fmt.Println("sql =", sql)
	if err != nil {
		return err
	}
	return nil
}

func (dps *ArDepositSpecial) SearchArDepositSpecialByKeyword(db *sqlx.DB, keyword string) (dpsList []*ArDepositSpecial, err error) {
	sql := `set dateformat dmy     select a.DocNo,a.DocDate,a.ArCode,isnull(b.name1,'') as ArName,isnull(a.CreatorCode,'') as CreatorCode,isnull(a.CreateDateTime,'') as CreateDateTime,isnull(a.LastEditorCode,'') as LastEditorCode,isnull(a.LastEditDateT,'') as LastEditDateT,isnull(a.DepartCode,'') as DepartCode,isnull(a.SaleCode,'') as SaleCode,isnull(c.name,'') as SaleName,a.IsConfirm,a.CreditDay,isnull(a.DueDate,'') as DueDate,isnull(a.MyDescription,'') as MyDescription,a.TotalAmount,a.SumOfWTax,a.BeforeTaxAmount,a.NetAmount,a.BillBalance,a.ExcessAmount1,a.ExcessAmount2,a.ChargeAmount,a.ChangeAmount,a.OtherIncome,a.OtherExpense,isnull(a.RefNo,'') as RefNo,isnull(a.CurrencyCode,'') as CurrencyCode,a.ExchangeRate,a.SumCashAmount,a.SumChqAmount,a.SumCreditAmount,a.SumBankAmount,a.PettyCashAmount,isnull(a.GLFormat,'') as GLFormat,a.IsCancel,a.IsReturnMoney,isnull(a.AllocateCode,'') as AllocateCode,isnull(a.ProjectCode,'') as ProjectCode,isnull(a.BillGroup,'') as BillGroup,isnull(a.RecurName,'') as RecurName,isnull(a.ConfirmCode,'') as ConfirmCode,isnull(a.ConfirmDateTime,'') as ConfirmDateTime,isnull(a.CancelCode,'') as CancelCode,isnull(a.CancelDateTime,'') as CancelDateTime from dbo.BCArDepositSpecial a WITH (NOLOCK) left join dbo.BCAR b WITH (NOLOCK) on a.arcode = b.code left join dbo.BCSale c WITH (NOLOCK) on a.salecode = c.code where (a.docno  like '%'+?+'%' or a.arcode like '%'+?+'%' or a.salecode like '%'+?+'%' or b.name1 like '%'+?+'%')  order by a.docno`
	err = db.Select(&dpsList, sql, keyword, keyword, keyword, keyword )
	fmt.Println("sql =", sql)
	if err != nil {
		return nil, err
	}
	return dpsList, nil
}
