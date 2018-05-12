package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type OtherExpense2 struct {
	DocNo           string  `json:"doc_no" db:"DocNo"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	SumofAmount     float64 `json:"sumof_amount" db:"SumofAmount"`
	AllocateCode    string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string  `json:"project_code" db:"ProjectCode"`
	ApCode          string  `json:"ap_code" db:"ApCode"`
	GLFormat        string  `json:"gl_format" db:"GLFormat"`
	PersonCode      string  `json:"person_code" db:"PersonCode"`
	BillStatus      int     `json:"bill_status" db:"BillStatus"`
	PettyCashAmount float64 `json:"petty_cash_amount" db:"PettyCashAmount"`
	SumCashAmount   float64 `json:"sum_cash_amount" db:"SumCashAmount"`
	SumChqAmount    float64 `json:"sum_chq_amount" db:"SumChqAmount"`
	SumBankAmount   float64 `json:"sum_bank_amount" db:"SumBankAmount"`
	OtherIncome     float64 `json:"other_income" db:"OtherIncome"`
	OtherExpense    float64 `json:"other_expense" db:"OtherExpense"`
	ExcessAmount1   float64 `json:"excess_amount_1" db:"ExcessAmount1"`
	ExcessAmount2   float64 `json:"excess_amount_2" db:"ExcessAmount2"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	RecurName       string  `json:"recur_name" db:"RecurName"`
	ConfirmCode     string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	CancelCode      string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string  `json:"cancel_date_time" db:"CancelDateTime"`
	UserCode        string  `json:"user_code"`
}

func (oxp2 *OtherExpense2) InsertAndEditOtherExpense2(db *sqlx.DB) error {
	var check_exist int

	sqlexist := `select count(docno) as check_exist from dbo.BCOTHEREXPENSE2 where docno = ?` //เช็คว่ามีเอกสารหรือยัง
	err := db.Get(&check_exist, sqlexist, oxp2.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	if (check_exist == 0) {
		//Insert//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		oxp2.CreatorCode = oxp2.UserCode

		sql := `set dateformat dmy     insert into dbo.BCOTHEREXPENSE2(DocNo,CreatorCode,CreateDateTime,DocDate,DepartCode,MyDescription,SumofAmount,AllocateCode,ProjectCode,ApCode,GLFormat,PersonCode,BillStatus,PettyCashAmount,SumCashAmount,SumChqAmount,SumBankAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,IsConfirm,RecurName) values(?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sql, oxp2.DocNo, oxp2.CreatorCode, oxp2.DocDate, oxp2.DepartCode, oxp2.MyDescription, oxp2.SumofAmount, oxp2.AllocateCode, oxp2.ProjectCode, oxp2.ApCode, oxp2.GLFormat, oxp2.PersonCode, oxp2.BillStatus, oxp2.PettyCashAmount, oxp2.SumCashAmount, oxp2.SumChqAmount, oxp2.SumBankAmount, oxp2.OtherIncome, oxp2.OtherExpense, oxp2.ExcessAmount1, oxp2.ExcessAmount2, oxp2.IsConfirm, oxp2.RecurName)
		if err != nil {
			return err
		}

	} else {
		oxp2.LastEditorCode = oxp2.UserCode

		sql := `set dateformat dmy     update dbo.BCOTHEREXPENSE2 set LastEditorCode=?,LastEditDateT=getdate(),DocDate=?,DepartCode=?,MyDescription=?,SumofAmount=?,AllocateCode=?,ProjectCode=?,ApCode=?,GLFormat=?,PersonCode=?,BillStatus=?,PettyCashAmount=?,SumCashAmount=?,SumChqAmount=?,SumBankAmount=?,OtherIncome=?,OtherExpense=?,ExcessAmount1=?,ExcessAmount2=?,IsConfirm=?,RecurName=? where docno = ?`
		_, err := db.Exec(sql, oxp2.LastEditorCode, oxp2.DocDate, oxp2.DepartCode, oxp2.MyDescription, oxp2.SumofAmount, oxp2.AllocateCode, oxp2.ProjectCode, oxp2.ApCode, oxp2.GLFormat, oxp2.PersonCode, oxp2.BillStatus, oxp2.PettyCashAmount, oxp2.SumCashAmount, oxp2.SumChqAmount, oxp2.SumBankAmount, oxp2.OtherIncome, oxp2.OtherExpense, oxp2.ExcessAmount1, oxp2.ExcessAmount2, oxp2.IsConfirm, oxp2.RecurName, oxp2.DocNo)
		if err != nil {
			return err
		}
	}

	return nil
}

func (oxp2 *OtherExpense2) SearchOtherExpense2ByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select DocNo,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,DocDate,isnull(DepartCode,'') as DepartCode,isnull(MyDescription,'') as MyDescription,SumofAmount,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(ApCode,'') as ApCode,isnull(GLFormat,'') as GLFormat,isnull(PersonCode,'') as PersonCode,BillStatus,PettyCashAmount,SumCashAmount,SumChqAmount,SumBankAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,IsConfirm,isnull(RecurName,'') as RecurName,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,IsCancel,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime from dbo.BCOTHEREXPENSE2 WITH (NOLOCK) where docno = ?`
	err := db.Get(&oxp2, sql, docno)
	if err != nil {
		return err
	}
	return nil
}

func (oxp2 *OtherExpense2) SearchOtherExpense2ByKeyword(db *sqlx.DB, keyword string) (oxps2 *[]OtherExpense2, err error) {
	sql := `set dateformat dmy     select DocNo,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,DocDate,isnull(DepartCode,'') as DepartCode,isnull(MyDescription,'') as MyDescription,SumofAmount,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(ApCode,'') as ApCode,isnull(GLFormat,'') as GLFormat,isnull(PersonCode,'') as PersonCode,BillStatus,PettyCashAmount,SumCashAmount,SumChqAmount,SumBankAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,IsConfirm,isnull(RecurName,'') as RecurName,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,IsCancel,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime from dbo.BCOTHEREXPENSE2 WITH (NOLOCK) where (docno  like '%'+?+'%' ) order by docno`
	err = db.Select(&oxps2, sql, keyword)
	if err != nil {
		return nil, err
	}
	return oxps2, nil
}
