package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type OtherExpense struct {
	DocNo           string  `json:"doc_no" db:"DocNo"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	RefDocNo        string  `json:"ref_doc_no" db:"RefDocNo"`
	GLBookCode      string  `json:"gl_book_code" db:"GLBookCode"`
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	SumofDebit      float64 `json:"sumof_debit" db:"SumofDebit"`
	SumofCredit     float64 `json:"sumof_credit" db:"SumofCredit"`
	NetAmount       float64 `json:"net_amount" db:"NetAmount"`
	AllocateCode    string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string  `json:"project_code" db:"ProjectCode"`
	ApCode          string  `json:"ap_code" db:"ApCode"`
	PettyCashAmount float64 `json:"petty_cash_amount" db:"PettyCashAmount"`
	SumOfWTax       float64 `json:"sum_of_w_tax" db:"SumOfWTax"`
	SumCashAmount   float64 `json:"sum_cash_amount" db:"SumCashAmount"`
	SumChqAmount    float64 `json:"sum_chq_amount" db:"SumChqAmount"`
	SumCreditAmount float64 `json:"sum_credit_amount" db:"SumCreditAmount"`
	SumBankAmount   float64 `json:"sum_bank_amount" db:"SumBankAmount"`
	OtherIncome     float64 `json:"other_income" db:"OtherIncome"`
	OtherExpense    float64 `json:"other_expense" db:"OtherExpense"`
	ExcessAmount1   float64 `json:"excess_amount_1" db:"ExcessAmount1"`
	ExcessAmount2   float64 `json:"excess_amount_2" db:"ExcessAmount2"`
	BillGroup       string  `json:"bill_group" db:"BillGroup"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	RecurName       string  `json:"recur_name" db:"RecurName"`
	ConfirmCode     string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string  `json:"cancel_date_time" db:"CancelDateTime"`
	UserCode        string  `json:"user_code"`
}

//func (ote *OtherExpense) SearchOtherExpenseByDocNo() error {
//	//sql := `DocNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, RefDocNo, GLBookCode, DepartCode, MyDescription, SumofDebit, SumofCredit, NetAmount, AllocateCode, ProjectCode, ApCode, PettyCashAmount, SumOfWTax, SumCashAmount, SumChqAmount, SumCreditAmount, SumBankAmount, OtherIncome, OtherExpense, ExcessAmount1, ExcessAmount2, BillGroup, IsConfirm, IsCancel, RecurName, ConfirmCode, ConfirmDateTime, CancelCode,CancelDateTime from dbo.BCOtherExpense WITH (NOLOCK)`
//	return nil
//}

func (ote *OtherExpense) InsertAndEditOtherExpense(db *sqlx.DB) error {
	var check_exist int

	sqlexist := `select count(docno) as check_exist from dbo.bcoherexpense where docno = ?` //เช็คว่ามีเอกสารหรือยัง
	err := db.Get(&check_exist, sqlexist, ote.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	if (check_exist == 0) {
		//Insert//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		ote.CreatorCode = ote.UserCode

		sql := `Insert into dbo.BCOtherExpense(DocNo,DocDate,CreatorCode,CreateDateTime,RefDocNo,GLBookCode,DepartCode,MyDescription,SumofDebit,SumofCredit,NetAmount,AllocateCode,ProjectCode,ApCode,PettyCashAmount,SumOfWTax,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,BillGroup,IsConfirm,IsCancel,RecurName) values(?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sql, ote.DocNo,ote.DocDate,ote.CreatorCode,ote.RefDocNo,ote.GLBookCode,ote.DepartCode,ote.MyDescription,ote.SumofDebit,ote.SumofCredit,ote.NetAmount,ote.AllocateCode,ote.ProjectCode,ote.ApCode,ote.PettyCashAmount,ote.SumOfWTax,ote.SumCashAmount,ote.SumChqAmount,ote.SumCreditAmount,ote.SumBankAmount,ote.OtherIncome,ote.OtherExpense,ote.ExcessAmount1,ote.ExcessAmount2,ote.BillGroup,ote.IsConfirm,ote.IsCancel,ote.RecurName)
		if err != nil {
			return err
		}
	} else {
		ote.LastEditorCode = ote.UserCode

		sql := `Update dbo.BCOtherExpense set DocDate=?,LasteditorCode=?,LastEditDateT=getdate(),RefDocNo=?,GLBookCode=?,DepartCode=?,MyDescription=?,SumofDebit=?,SumofCredit=?,NetAmount=?,AllocateCode=?,ProjectCode=?,ApCode=?,PettyCashAmount=?,SumOfWTax=?,SumCashAmount=?,SumChqAmount=?,SumCreditAmount=?,SumBankAmount=?,OtherIncome=?,OtherExpense=?,ExcessAmount1=?,ExcessAmount2=?,BillGroup=?,IsConfirm=?,IsCancel=?,RecurName=? where docno = ?`
		_, err := db.Exec(sql, ote.DocDate,ote.LastEditorCode,ote.RefDocNo,ote.GLBookCode,ote.DepartCode,ote.MyDescription,ote.SumofDebit,ote.SumofCredit,ote.NetAmount,ote.AllocateCode,ote.ProjectCode,ote.ApCode,ote.PettyCashAmount,ote.SumOfWTax,ote.SumCashAmount,ote.SumChqAmount,ote.SumCreditAmount,ote.SumBankAmount,ote.OtherIncome,ote.OtherExpense,ote.ExcessAmount1,ote.ExcessAmount2,ote.BillGroup,ote.IsConfirm,ote.IsCancel,ote.RecurName, ote.DocNo)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ote *OtherExpense) SearchOtherExpenseByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select DocNo,DocDate,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(RefDocNo,'') as RefDocNo,isnull(GLBookCode,'') as GLBookCode,isnull(DepartCode,'') as DepartCode,isnull(MyDescription,'') as MyDescription,SumofDebit,SumofCredit,NetAmount,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(ApCode,'') as ApCode,PettyCashAmount,SumOfWTax,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,isnull(BillGroup,'') as BillGroup,IsConfirm,IsCancel,isnull(RecurName,'') as RecurName,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime from BCOtherExpense with (nolock) where docno = ?`
	err := db.Get(&ote, sql, docno)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}
	return nil
}

func (ote *OtherExpense) SearchOtherExpenseByKeyword(db *sqlx.DB, docno string) (otes *[]OtherExpense, err error) {
	sql := `set dateformat dmy     select DocNo,DocDate,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(RefDocNo,'') as RefDocNo,isnull(GLBookCode,'') as GLBookCode,isnull(DepartCode,'') as DepartCode,isnull(MyDescription,'') as MyDescription,SumofDebit,SumofCredit,NetAmount,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(ApCode,'') as ApCode,PettyCashAmount,SumOfWTax,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,isnull(BillGroup,'') as BillGroup,IsConfirm,IsCancel,isnull(RecurName,'') as RecurName,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime from BCOtherExpense with (nolock) where (docno  like '%'+?+'%' ) order by docno`
	err = db.Select(&otes, sql, docno)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}
	return otes, nil
}
