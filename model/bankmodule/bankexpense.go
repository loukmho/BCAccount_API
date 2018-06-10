package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type BankExpense struct {
	DocNo           string  `json:"doc_no" db:"DocNo"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	BookNo          string  `json:"book_no" db:"BookNo"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	AccountCode     string  `json:"account_code" db:"AccountCode"`
	GLBookCode      string  `json:"gl_book_code" db:"GLBookCode"`
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	Amount          float64 `json:"amount" db:"Amount"`
	AllocateCode    string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string  `json:"project_code" db:"ProjectCode"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	RecurName       string  `json:"recur_name" db:"RecurName"`
	ConfirmCode     string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string  `json:"cancel_date_time" db:"CancelDateTime"`
	UserCode        string  `json:"user_code"`
}

//func (bep *BankExpense) SearchBankExpenseByDocNo() error {
//	//sql := `DocNo, DocDate, BookNo, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, AccountCode, GLBookCode, DepartCode, MyDescription, Amount, AllocateCode, ProjectCode, IsCancel, IsConfirm, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime from dbo.BCBANKEXPENSE WITH (NOLOCK)`
//	return nil
//}

func (bep *BankExpense) InsertAndEditBankExpense(db *sqlx.DB) error {
	var check_exist int

	sqlexist := `select count(docno) as check_exist from dbo.bcbankexpense where docno = ?` //เช็คว่ามีเอกสารหรือยัง
	err := db.Get(&check_exist, sqlexist, bep.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	if (check_exist == 0) {
		//Insert//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		bep.CreatorCode = bep.UserCode

		sql := `Insert into dbo.BCBankExpense(DocNo,DocDate,BookNo,CreatorCode,CreateDateTime,AccountCode,GLBookCode,DepartCode,MyDescription,Amount,AllocateCode,ProjectCode,IsCancel,IsConfirm,RecurName) values(?,?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sql, bep.DocNo,bep.DocDate,bep.BookNo,bep.CreatorCode,bep.AccountCode,bep.GLBookCode,bep.DepartCode,bep.MyDescription,bep.Amount,bep.AllocateCode,bep.ProjectCode,bep.IsCancel,bep.IsConfirm,bep.RecurName)
		if err != nil {
			return err
		}
	} else {
		bep.LastEditorCode = bep.UserCode

		sql := `Update dbo.BCBankExpense set DocDate=?,BookNo=?,LastEditCode=?,LastEditDateT=getdate(),AccountCode=?,GLBookCode=?,DepartCode=?,MyDescription=?,Amount=?,AllocateCode=?,ProjectCode=?,IsCancel=?,IsConfirm=?,RecurName=? where docno = ?`
		_, err := db.Exec(sql, bep.DocDate,bep.BookNo,bep.LastEditorCode,bep.AccountCode,bep.GLBookCode,bep.DepartCode,bep.MyDescription,bep.Amount,bep.AllocateCode,bep.ProjectCode,bep.IsCancel,bep.IsConfirm,bep.RecurName, bep.DocNo)
		if err != nil {
			return err
		}
	}

	return nil
}

func (bep *BankExpense) SearchBankExpenseByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select DocNo,DocDate,isnull(BookNo,'') as BookNo,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(AccountCode,'') as AccountCode,isnull(GLBookCode,'') as GLBookCode,isnull(DepartCode,'') as DepartCode,isnull(MyDescription,'') as MyDescription,Amount,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,IsCancel,IsConfirm,isnull(RecurName,'') as RecurName from dbo.BCBankExpense with (nolock) where docno = ?`
	err := db.Get(&bep, sql, docno)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}
	return nil
}

func (bep *BankExpense) SearchBankExpenseByKeyword(db *sqlx.DB, docno string) (beps *[]BankExpense, err error) {
	sql := `set dateformat dmy     select DocNo,DocDate,isnull(BookNo,'') as BookNo,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(AccountCode,'') as AccountCode,isnull(GLBookCode,'') as GLBookCode,isnull(DepartCode,'') as DepartCode,isnull(MyDescription,'') as MyDescription,Amount,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,IsCancel,IsConfirm,isnull(RecurName,'') as RecurName from dbo.BCBankExpense with (nolock) where (docno  like '%'+?+'%' ) order by docno`
	err = db.Select(&beps, sql, docno)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}
	return beps, nil
}
