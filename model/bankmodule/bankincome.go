package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type BankInCome struct {
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

//func (bic *BankInCome) SearchBankInCome() error {
//	//sql := `DocNo, DocDate, BookNo, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, AccountCode, GLBookCode,DepartCode, MyDescription, Amount, AllocateCode, ProjectCode,IsCancel, IsConfirm, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime, RecurName from dbo.BCBankIncome WITH (NOLOCK)`
//	return nil
//}

func (bic *BankInCome) InsertAndEditBankInCome(db *sqlx.DB) error {
	var check_exist int

	sqlexist := `select count(docno) as check_exist from dbo.bcbankincome where docno = ?` //เช็คว่ามีเอกสารหรือยัง
	err := db.Get(&check_exist, sqlexist, bic.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	if (check_exist == 0) {
		//Insert//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		bic.CreatorCode = bic.UserCode

		sql := `Insert into dbo.BCBankInCome(DocNo,DocDate,BookNo,CreatorCode,CreateDateTime,AccountCode,GLBookCode,DepartCode,MyDescription,Amount,AllocateCode,ProjectCode,IsCancel,IsConfirm,RecurName) values(?,?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sql, bic.DocNo, bic.DocDate, bic.BookNo, bic.CreatorCode, bic.AccountCode, bic.GLBookCode, bic.DepartCode, bic.MyDescription, bic.Amount, bic.AllocateCode, bic.ProjectCode, bic.IsCancel, bic.IsConfirm, bic.RecurName)
		if err != nil {
			return err
		}
	} else {
		bic.LastEditorCode = bic.UserCode

		sql := `Update dbo.BCBankInCome set DocDate=?,BookNo=?,LastEditCode=?,LastEditDateT=getdate(),AccountCode=?,GLBookCode=?,DepartCode=?,MyDescription=?,Amount=?,AllocateCode=?,ProjectCode=?,IsCancel=?,IsConfirm=?,RecurName=? where docno = ?`
		_, err := db.Exec(sql, bic.DocDate, bic.BookNo, bic.LastEditorCode, bic.AccountCode, bic.GLBookCode, bic.DepartCode, bic.MyDescription, bic.Amount, bic.AllocateCode, bic.ProjectCode, bic.IsCancel, bic.IsConfirm, bic.RecurName, bic.DocNo)
		if err != nil {
			return err
		}
	}

	return nil
}

func (bic *BankInCome) SearchBankInComeByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select DocNo,DocDate,isnull(BookNo,'') as BookNo,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(AccountCode,'') as AccountCode,isnull(GLBookCode,'') as GLBookCode,isnull(DepartCode,'') as DepartCode,isnull(MyDescription,'') as MyDescription,Amount,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,IsCancel,IsConfirm,isnull(RecurName,'') as RecurName from dbo.BCBankInCome with (nolock) where docno = ?`
	err := db.Get(&bic, sql, docno)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}
	return nil
}

func (bic *BankInCome) SearchBankInComeByKeyword(db *sqlx.DB, docno string) (bics *[]BankInCome, err error) {
	sql := `set dateformat dmy     select DocNo,DocDate,isnull(BookNo,'') as BookNo,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(AccountCode,'') as AccountCode,isnull(GLBookCode,'') as GLBookCode,isnull(DepartCode,'') as DepartCode,isnull(MyDescription,'') as MyDescription,Amount,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,IsCancel,IsConfirm,isnull(RecurName,'') as RecurName from dbo.BCBankInCome with (nolock) where (docno  like '%'+?+'%' ) order by docno`
	err = db.Select(&bic, sql, docno)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}
	return bics, nil
}
