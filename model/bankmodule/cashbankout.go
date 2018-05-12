package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type CashBankOut struct {
	DocNo           string  `json:"doc_no" db:"DocNo"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	BookNo          string  `json:"book_no" db:"BookNo"`
	AccountCode     string  `json:"account_code" db:"AccountCode"`
	GLBookCode      string  `json:"gl_book_code" db:"GLBookCode"`
	IsBankBalance   float64 `json:"is_bank_balance" db:"IsBankBalance"`
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	Amount          float64 `json:"amount" db:"Amount"`
	AllocateCode    string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string  `json:"project_code" db:"ProjectCode"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	ConfirmCode     string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string  `json:"cancel_date_time" db:"CancelDateTime"`
	UserCode        string  `json:"user_code"`
}

//func (cbo *CashBankOut) SearchCashBankOutByDocNo() error {
//	//sql := `DocNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, BookNo, AccountCode, GLBookCode, IsBankBalance,DepartCode, MyDescription, Amount,  AllocateCode, ProjectCode, IsCancel, IsConfirm, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime from dbo.BCCashBankOut WITH (NOLOCK)`
//	return nil
//}

func (cbo *CashBankOut) InsertAndEditCashBankOut(db *sqlx.DB) error {
	var check_exist int

	sqlexist := `select count(docno) as check_exist from dbo.bccashbankout where docno = ?` //เช็คว่ามีเอกสารหรือยัง
	err := db.Get(&check_exist, sqlexist, cbo.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	if (check_exist == 0) {
		//Insert//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		cbo.CreatorCode = cbo.UserCode

		sql := `Insert into dbo.BCCashBankOut(DocNo,DocDate,BookNo,CreatorCode,CreateDateTime,AccountCode,GLBookCode,DepartCode,MyDescription,Amount,AllocateCode,ProjectCode,IsCancel,IsConfirm,RecurName) values(?,?,?,?,getdate(),?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sql, cbo.DocNo, cbo.DocDate, cbo.BookNo, cbo.CreatorCode, cbo.AccountCode, cbo.GLBookCode, cbo.DepartCode, cbo.MyDescription, cbo.Amount, cbo.AllocateCode, cbo.ProjectCode, cbo.IsCancel, cbo.IsConfirm)
		if err != nil {
			return err
		}
	} else {
		cbo.LastEditorCode = cbo.UserCode

		sql := `Update dbo.BCCashBankOut set DocDate=?,BookNo=?,LastEditCode=?,LastEditDateT=getdate(),AccountCode=?,GLBookCode=?,DepartCode=?,MyDescription=?,Amount=?,AllocateCode=?,ProjectCode=?,IsCancel=?,IsConfirm=?,RecurName=? where docno = ?`
		_, err := db.Exec(sql, cbo.DocDate, cbo.BookNo, cbo.LastEditorCode, cbo.AccountCode, cbo.GLBookCode, cbo.DepartCode, cbo.MyDescription, cbo.Amount, cbo.AllocateCode, cbo.ProjectCode, cbo.IsCancel, cbo.IsConfirm, cbo.DocNo)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cbo *CashBankOut) SearchCashBankOutByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select DocNo,DocDate,isnull(BookNo,'') as BookNo,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(AccountCode,'') as AccountCode,isnull(GLBookCode,'') as GLBookCode,isnull(DepartCode,'') as DepartCode,isnull(MyDescription,'') as MyDescription,Amount,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,IsCancel,IsConfirm,isnull(RecurName,'') as RecurName from dbo.BCCashBankOut with (nolock) where docno = ?`
	err := db.Get(&cbo, sql, docno)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}
	return nil
}

func (cbo *CashBankOut) SearchCashBankOutByKeyword(db *sqlx.DB, docno string) (cbos *[]CashBankOut, err error) {
	sql := `set dateformat dmy     select DocNo,DocDate,isnull(BookNo,'') as BookNo,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(AccountCode,'') as AccountCode,isnull(GLBookCode,'') as GLBookCode,isnull(DepartCode,'') as DepartCode,isnull(MyDescription,'') as MyDescription,Amount,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,IsCancel,IsConfirm,isnull(RecurName,'') as RecurName from dbo.BCCashBankOut with (nolock) where (docno  like '%'+?+'%' ) order by docno`
	err = db.Select(&cbo, sql, docno)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}
	return cbos, nil
}
