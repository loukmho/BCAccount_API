package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type CashBankIn struct {
	DocNo           string  `json:"doc_no" db:"DocNo"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	BookNo          string  `json:"book_no" db:"BookNo"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
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

//func (cbi *CashBankIn) SearchCashBankInByDocNo() error {
//	//sql := `DocNo, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, DocDate, BookNo, AccountCode, GLBookCode, IsBankBalance, DepartCode, MyDescription, Amount, AllocateCode, ProjectCode, IsCancel, IsConfirm, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime from dbo.BCCASHBANKIN WITH (NOLOCK)`
//	return nil
//}


func (cbi *CashBankIn) InsertAndEditCashBankIn(db *sqlx.DB) error {
	var check_exist int

	sqlexist := `select count(docno) as check_exist from dbo.bccashbankin where docno = ?` //เช็คว่ามีเอกสารหรือยัง
	err := db.Get(&check_exist, sqlexist, cbi.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	if (check_exist == 0) {
		//Insert//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		cbi.CreatorCode = cbi.UserCode

		sql := `Insert into dbo.BCCashBankIn(DocNo,DocDate,BookNo,CreatorCode,CreateDateTime,AccountCode,GLBookCode,DepartCode,MyDescription,Amount,AllocateCode,ProjectCode,IsCancel,IsConfirm,RecurName) values(?,?,?,?,getdate(),?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sql, cbi.DocNo,cbi.DocDate,cbi.BookNo,cbi.CreatorCode,cbi.AccountCode,cbi.GLBookCode,cbi.DepartCode,cbi.MyDescription,cbi.Amount,cbi.AllocateCode,cbi.ProjectCode,cbi.IsCancel,cbi.IsConfirm)
		if err != nil {
			return err
		}
	} else {
		cbi.LastEditorCode = cbi.UserCode

		sql := `Update dbo.BCCashBankIn set DocDate=?,BookNo=?,LastEditCode=?,LastEditDateT=getdate(),AccountCode=?,GLBookCode=?,DepartCode=?,MyDescription=?,Amount=?,AllocateCode=?,ProjectCode=?,IsCancel=?,IsConfirm=?,RecurName=? where docno = ?`
		_, err := db.Exec(sql, cbi.DocDate,cbi.BookNo,cbi.LastEditorCode,cbi.AccountCode,cbi.GLBookCode,cbi.DepartCode,cbi.MyDescription,cbi.Amount,cbi.AllocateCode,cbi.ProjectCode,cbi.IsCancel,cbi.IsConfirm,cbi.DocNo)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cbi *CashBankIn) SearchCashBankInByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select DocNo,DocDate,isnull(BookNo,'') as BookNo,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(AccountCode,'') as AccountCode,isnull(GLBookCode,'') as GLBookCode,isnull(DepartCode,'') as DepartCode,isnull(MyDescription,'') as MyDescription,Amount,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,IsCancel,IsConfirm,isnull(RecurName,'') as RecurName from dbo.BCCashBankIn with (nolock) where docno = ?`
	err := db.Get(&cbi, sql, docno)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}
	return nil
}

func (cbi *CashBankIn) SearchCashBankInByKeyword(db *sqlx.DB, docno string) (cbis *[]CashBankIn, err error) {
	sql := `set dateformat dmy     select DocNo,DocDate,isnull(BookNo,'') as BookNo,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(AccountCode,'') as AccountCode,isnull(GLBookCode,'') as GLBookCode,isnull(DepartCode,'') as DepartCode,isnull(MyDescription,'') as MyDescription,Amount,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,IsCancel,IsConfirm,isnull(RecurName,'') as RecurName from dbo.BCCashBankIn with (nolock) where (docno  like '%'+?+'%' ) order by docno`
	err = db.Select(&cbis, sql, docno)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}
	return cbis, nil
}