package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type OtherInCome struct {
	DocNo           string  `json:"doc_no" db:"DocNo"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	GLBookCode      string  `json:"gl_book_code" db:"GLBookCode"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	SumofDebit      float64 `json:"sumof_debit" db:"SumofDebit"`
	SumofCredit     float64 `json:"sumof_credit" db:"SumofCredit"`
	NetAmount       float64 `json:"net_amount" db:"NetAmount"`
	AllocateCode    string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string  `json:"project_code" db:"ProjectCode"`
	ArCode          string  `json:"ar_code" db:"ArCode"`
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

func (oti *OtherInCome) InsertAndEditOtherInCome(db *sqlx.DB) error {
	var check_exist int

	sqlexist := `select count(docno) as check_exist from dbo.bcoherincome where docno = ?` //เช็คว่ามีเอกสารหรือยัง
	err := db.Get(&check_exist, sqlexist, oti.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	if (check_exist == 0) {
		//Insert//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		oti.CreatorCode = oti.UserCode

		sql := `set dateformat dmy     insert into dbo.BCOTHERINCOME(DocNo,DocDate,GLBookCode,CreatorCode,CreateDateTime,DepartCode,MyDescription,SumofDebit,SumofCredit,NetAmount,AllocateCode,ProjectCode,ArCode,PettyCashAmount,SumOfWTax,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,BillGroup,IsConfirm,IsCancel,RecurName) values(?,?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sql, oti.DocNo, oti.DocDate, oti.GLBookCode, oti.CreatorCode, oti.DepartCode, oti.MyDescription, oti.SumofDebit, oti.SumofCredit, oti.NetAmount, oti.AllocateCode, oti.ProjectCode, oti.ArCode, oti.PettyCashAmount, oti.SumOfWTax, oti.SumCashAmount, oti.SumChqAmount, oti.SumCreditAmount, oti.SumBankAmount, oti.OtherIncome, oti.OtherExpense, oti.ExcessAmount1, oti.ExcessAmount2, oti.BillGroup, oti.IsConfirm, oti.IsCancel, oti.RecurName)
		if err != nil {
			return err
		}
	} else {
		//Update//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		oti.LastEditorCode = oti.UserCode

		sql := `set dateformat dmy     update dbo.BCOTHERINCOME set DocDate=?,GLBookCode=?,LastEditorCode=?,LastEditDateT=getdate(),DepartCode=?,MyDescription=?,SumofDebit=?,SumofCredit=?,NetAmount=?,AllocateCode=?,ProjectCode=?,ArCode=?,PettyCashAmount=?,SumOfWTax=?,SumCashAmount=?,SumChqAmount=?,SumCreditAmount=?,SumBankAmount=?,OtherIncome=?,OtherExpense=?,ExcessAmount1=?,ExcessAmount2=?,BillGroup=?,IsConfirm=?,IsCancel=?,RecurName=? where docno = ?`
		_, err := db.Exec(sql, oti.DocDate,oti.GLBookCode,oti.LastEditorCode,oti.DepartCode,oti.MyDescription,oti.SumofDebit,oti.SumofCredit,oti.NetAmount,oti.AllocateCode,oti.ProjectCode,oti.ArCode,oti.PettyCashAmount,oti.SumOfWTax,oti.SumCashAmount,oti.SumChqAmount,oti.SumCreditAmount,oti.SumBankAmount,oti.OtherIncome,oti.OtherExpense,oti.ExcessAmount1,oti.ExcessAmount2,oti.BillGroup,oti.IsConfirm,oti.IsCancel,oti.RecurName, oti.DocNo)
		if err != nil {
			return err
		}
	}

	return nil
}

func (oti *OtherInCome) SearchOtherInComeByDocNo(db *sqlx.DB, docno string) error {
	//sql := `DocNo, DocDate, GLBookCode,CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, DepartCode, MyDescription, SumofDebit, SumofCredit, NetAmount, AllocateCode,  ProjectCode, ArCode, PettyCashAmount, SumOfWTax, SumCashAmount, SumChqAmount, SumCreditAmount, SumBankAmount, OtherIncome, OtherExpense, ExcessAmount1, ExcessAmount2, BillGroup, IsConfirm, IsCancel, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime from dbo.BCOTHERINCOME WITH (NOLOCK)`
	sql := `set dateformat dmy     select DocNo,DocDate,isnull(GLBookCode,'') as GLBookCode,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(DepartCode,'') as DepartCode,isnull(MyDescription,'') as MyDescription,SumofDebit,SumofCredit,NetAmount,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(ArCode,'') as ArCode,PettyCashAmount,SumOfWTax,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,isnull(BillGroup,'') as BillGroup,IsConfirm,IsCancel,isnull(RecurName,'') as RecurName,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT from dbo.BCOtherIncome with (nolock) where docno = ?`
	err := db.Get(&oti, sql, docno)
	if err != nil {
		return err
	}
	return nil
}



func (oti *OtherInCome) SearchOtherInComeByKeyword(db *sqlx.DB, keyword string) (otis *[]OtherInCome, err error) {
	//sql := `DocNo, DocDate, GLBookCode,CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, DepartCode, MyDescription, SumofDebit, SumofCredit, NetAmount, AllocateCode,  ProjectCode, ArCode, PettyCashAmount, SumOfWTax, SumCashAmount, SumChqAmount, SumCreditAmount, SumBankAmount, OtherIncome, OtherExpense, ExcessAmount1, ExcessAmount2, BillGroup, IsConfirm, IsCancel, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime from dbo.BCOTHERINCOME WITH (NOLOCK)`
	sql := `set dateformat dmy     select DocNo,DocDate,isnull(GLBookCode,'') as GLBookCode,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(DepartCode,'') as DepartCode,isnull(MyDescription,'') as MyDescription,SumofDebit,SumofCredit,NetAmount,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(ArCode,'') as ArCode,PettyCashAmount,SumOfWTax,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,isnull(BillGroup,'') as BillGroup,IsConfirm,IsCancel,isnull(RecurName,'') as RecurName,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT from dbo.BCOtherIncome with (nolock) where (docno  like '%'+?+'%' ) order by docno`
	err = db.Select(&otis, sql, keyword)
	if err != nil {
		return nil, err
	}
	return otis, nil
}



