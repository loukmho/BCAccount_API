package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
)

type BankTransfer struct {
	DocNo           string  `json:"doc_no" db:"DocNo"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	FromBook        string  `json:"from_book" db:"FromBook"`
	ToBook          string  `json:"to_book" db:"ToBook"`
	GLBookCode      string  `json:"gl_book_code" db:"GLBookCode"`
	IsPostGL        int     `json:"is_post_gl" db:"IsPostGL"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	AllocateCode    string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string  `json:"project_code" db:"ProjectCode"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	Amount          float64 `json:"amount" db:"Amount"`
	ChargeAmount    float64 `json:"charge_amount" db:"ChargeAmount"`
	TotalAmount     float64 `json:"total_amount" db:"TotalAmount"`
	RecurName       string  `json:"recur_name" db:"RecurName"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	ConfirmCode     string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string  `json:"cancel_date_time" db:"CancelDateTime"`
	UserCode        string  `json:"user_code" db:"UserCode"`
}

func (btf *BankTransfer) InsertAndEditBankTransfer(db *sqlx.DB) error {
	var check_exist int

	sqlexist := `select count(docno) as check_exist from dbo.bcbanktransfer where docno = ?` //เช็คว่ามีเอกสารหรือยัง
	err := db.Get(&check_exist, sqlexist, btf.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	if (check_exist == 0) {
		//Insert//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		btf.CreatorCode = btf.UserCode

		sql := `Insert into dbo.BCBankTransfer(DocNo,DocDate,CreatorCode,CreateDateTime,FromBook,ToBook,GLBookCode,IsPostGL,IsCancel,DepartCode,AllocateCode,ProjectCode,MyDescription,Amount,ChargeAmount,TotalAmount,RecurName,IsConfirm) values(?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sql, btf.DocNo, btf.DocDate, btf.CreatorCode, btf.FromBook, btf.ToBook, btf.GLBookCode, btf.IsPostGL, btf.IsCancel, btf.DepartCode, btf.AllocateCode, btf.ProjectCode, btf.MyDescription, btf.Amount, btf.ChargeAmount, btf.TotalAmount, btf.RecurName, btf.IsConfirm)
		if err != nil {
			return err
		}
	} else {
		sql := `Update dbo.BCBankTransfer set DocDate=?,LastEditorCode=?,LastEditDateT=getdate(),FromBook=?,ToBook=?,GLBookCode=?,IsPostGL=?,IsCancel=?,DepartCode=?,AllocateCode=?,ProjectCode=?,MyDescription=?,Amount=?,ChargeAmount=?,TotalAmount=?,RecurName=?,IsConfirm=? where docno = ?`
		_, err := db.Exec(sql, btf.DocDate, btf.LastEditorCode, btf.FromBook, btf.ToBook, btf.GLBookCode, btf.IsPostGL, btf.IsCancel, btf.DepartCode, btf.AllocateCode, btf.ProjectCode, btf.MyDescription, btf.Amount, btf.ChargeAmount, btf.TotalAmount, btf.RecurName, btf.IsConfirm, btf.DocNo)
		if err != nil {
			return err
		}
	}

	return nil
}

func (btf *BankTransfer) SearchBankTransferByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select DocNo,DocDate,CreatorCode,CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(FromBook,'') as FromBook,isnull(ToBook,'') as ToBook,isnull(GLBookCode,'') as GLBookCode,IsPostGL,IsCancel,isnull(DepartCode,'') as DepartCode,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(MyDescription,'') as MyDescription,Amount,ChargeAmount,TotalAmount,isnull(RecurName,'') as RecurName, IsConfirm,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime from dbo.bcbanktransfer where docno = ?`
	err := db.Get(&btf, sql, docno)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}
	return nil
}

func (btf *BankTransfer) SearchBankTransferByKeyword(db *sqlx.DB, docno string) (btfs []*BankTransfer, err error) {
	sql := `set dateformat dmy     select DocNo,DocDate,CreatorCode,CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(FromBook,'') as FromBook,isnull(ToBook,'') as ToBook,isnull(GLBookCode,'') as GLBookCode,IsPostGL,IsCancel,isnull(DepartCode,'') as DepartCode,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(MyDescription,'') as MyDescription,Amount,ChargeAmount,TotalAmount,isnull(RecurName,'') as RecurName, IsConfirm,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime from dbo.bcbanktransfer where (docno  like '%'+?+'%' ) order by docno`
	err = db.Select(&btfs, sql, docno)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}
	return btfs,nil
}
