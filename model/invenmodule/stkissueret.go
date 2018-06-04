package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"errors"
	m "github.com/loukmho/bcaccount_api/model"
)

type StkIssueRet struct {
	DocNo           string  `json:"doc_no" db:"DocNo"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	IssueRefNo      string  `json:"issue_ref_no" db:"IssueRefNo"`
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	PersonCode      string  `json:"person_code" db:"PersonCode"`
	SumOfAmount     float64 `json:"sum_of_amount" db:"SumOfAmount"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	GLFormat        string  `json:"gl_format" db:"GLFormat"`
	IsCompleteSave  int     `json:"is_complete_save" db:"IsCompleteSave"`
	AllocateCode    string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string  `json:"project_code" db:"ProjectCode"`
	RecurName       string  `json:"recur_name" db:"RecurName"`
	ConfirmCode     string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string  `json:"cancel_date_time" db:"CancelDateTime"`
	UserCode        string     `json:"user_code" db:"UserCode"`
	Subs            []*IsrItem `json:"subs"`
}

type IsrItem struct {
	MyType        int     `json:"my_type" db:"MyType"`
	ItemCode      string  `json:"item_code" db:"ItemCode"`
	MyDescription string  `json:"my_description" db:"MyDescription"`
	ItemName      string  `json:"item_name" db:"ItemName"`
	WHCode        string  `json:"wh_code" db:"WHCode"`
	ShelfCode     string  `json:"shelf_code" db:"ShelfCode"`
	Qty           float64 `json:"qty" db:"Qty"`
	Price         float64 `json:"price" db:"Price"`
	Amount        float64 `json:"amount" db:"Amount"`
	HomeAmount    float64 `json:"home_amount" db:"HomeAmount"`
	SumOfCost     float64 `json:"sum_of_cost" db:"SumOfCost"`
	UnitCode      string  `json:"unit_code" db:"UnitCode"`
	BarCode       string  `json:"bar_code" db:"BarCode"`
	IsCancel      int     `json:"is_cancel" db:"IsCancel"`
	LineNumber    int     `json:"line_number" db:"LineNumber"`
	AverageCost   float64 `json:"average_cost" db:"AverageCost"`
	RefLinenumber int     `json:"ref_linenumber" db:"RefLinenumber"`
	LotNumber     string  `json:"lot_number" db:"LotNumber"`
	PackingRate1  float64 `json:"packing_rate_1" db:"PackingRate1"`
	PackingRate2  float64 `json:"packing_rate_2" db:"PackingRate2"`
}

//func (sir *StkIssueRet) SearchStkIssueRetByDocNo() error {
//	//sql := `DocNo, DocDate,CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, IssueRefNo, DepartCode, IsConfirm, IsCancel, PersonCode, SumOfAmount, MyDescription, GLFormat,IsCompleteSave, AllocateCode, ProjectCode, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime`
//	//sqlsub := `MyType, DocNo, IssueRefNo, ItemCode, DocDate, DepartCode, Personcode, MyDescription, ItemName, WHCode, ShelfCode, Qty, Price, Amount, HomeAmount, SumOfCost,UnitCode,  BarCode, IsCancel, LineNumber,  AVERAGECOST, RefLinenumber, LotNumber,PackingRate1, PackingRate2`
//	return nil
//
//}


func (sir *StkIssueRet) InsertAndEditStkIssueRet(db *sqlx.DB) error {
	var check_exist int
	var count_item int

	for _, sub_item := range sir.Subs {
		if (sub_item.Qty != 0) {
			count_item = count_item + 1
		}
	}

	switch {
	case sir.DocNo == "":
		return errors.New("Docno is null")
	case sir.DocDate == "":
		return errors.New("Docdate is null")
	case sir.IsCancel == 1:
		return errors.New("Docno is cancel")
	case sir.IsConfirm == 1:
		return errors.New("Docno is confirm")
	case count_item == 0:
		return errors.New("Docno is not have item")
	}


	sqlexist := `select count(docno) as check_exist from dbo.BCStkIssueRet where docno = ?` //เช็คว่ามีเอกสารหรือยัง
	err := db.Get(&check_exist, sqlexist, sir.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")
	sir.GLFormat = def.StkIssueGLFormat

	sir.IsCompleteSave = 1

	if check_exist == 0 {
		sir.CreatorCode = sir.UserCode

		sql := `set dateformat dmy     Insert into dbo.BCStkIssueRet(DocNo,DocDate,CreatorCode,CreateDateTime,IssueRefNo,DepartCode,IsConfirm,IsCancel,PersonCode,SumOfAmount,MyDescription,GLFormat,IsCompleteSave,AllocateCode,ProjectCode,RecurName) values(?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sql, sir.DocNo,sir.DocDate,sir.CreatorCode,sir.IssueRefNo,sir.DepartCode,sir.IsConfirm,sir.IsCancel,sir.PersonCode,sir.SumOfAmount,sir.MyDescription,sir.GLFormat,sir.IsCompleteSave,sir.AllocateCode,sir.ProjectCode,sir.RecurName)
		if err != nil {
			return err
		}

	} else {
		sir.LastEditorCode = sir.UserCode

		sql := `set dateformat dmy     update dbo.BCStkIssueRet set DocDate=?,LastEditorCode=?,LastEditDateT=getdate(),IssueRefNo=?,DepartCode=?,IsConfirm=?,IsCancel=?,PersonCode=?,SumOfAmount=?,MyDescription=?,GLFormat=?,IsCompleteSave=?,AllocateCode=?,ProjectCode=?,RecurName=? where docno = ?`
		_, err := db.Exec(sql, sir.DocDate,sir.LastEditorCode,sir.IssueRefNo,sir.DepartCode,sir.IsConfirm,sir.IsCancel,sir.PersonCode,sir.SumOfAmount,sir.MyDescription,sir.GLFormat,sir.IsCompleteSave,sir.AllocateCode,sir.ProjectCode,sir.RecurName, sir.DocNo )
		if err != nil {
			return err
		}
	}

	sql_del_sub := `delete dbo.BCStkIssRetSub where docno = ?`
	_, err = db.Exec(sql_del_sub, sir.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}

	var vLineNumber int

	for _, sub := range sir.Subs {
		sub.MyType = def.StkIssueMyType
		sub.LineNumber = vLineNumber

		sub.IsCancel = 0
		if (sub.PackingRate1 == 0) {
			sub.PackingRate1 = 1
		}
		if (sub.PackingRate2 == 0) {
			sub.PackingRate2 = 1
		}

		sqlsub := `set dateformat dmy     Insert into dbo.BCStkIssRetSub(MyType,DocNo,IssueRefNo,ItemCode,DocDate,DepartCode,Personcode,MyDescription,ItemName,WHCode,ShelfCode,Qty,Price,Amount,HomeAmount,SumOfCost,UnitCode,BarCode,IsCancel,LineNumber,AVERAGECOST,RefLinenumber,LotNumber,PackingRate1,PackingRate2) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		db.Exec(sqlsub, sub.MyType,sir.DocNo,sir.IssueRefNo,sub.ItemCode,sir.DocDate,sir.DepartCode,sir.PersonCode,sub.MyDescription,sub.ItemName,sub.WHCode,sub.ShelfCode,sub.Qty,sub.Price,sub.Amount,sub.HomeAmount,sub.SumOfCost,sub.UnitCode,sub.BarCode,sub.IsCancel,sub.LineNumber,sub.AverageCost,sub.RefLinenumber,sub.LotNumber,sub.PackingRate1,sub.PackingRate2)
		vLineNumber = vLineNumber + 1

		sqlprocess := ` insert into dbo.ProcessStock (ItemCode,ProcessFlag,FlowStatus) values(?, 1, 0)`
		_, err = db.Exec(sqlprocess, sub.ItemCode)
		fmt.Println("sqlprocess = ", sqlsub)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			fmt.Println(err.Error())
		}
	}

	return nil
}

func (sir *StkIssueRet) SearchStkIssueRetByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select DocNo, DocDate,isnull(CreatorCode,'') as CreatorCode, isnull(CreateDateTime,'') as CreateDateTime, isnull(LastEditorCode,'') as LastEditorCode, isnull(LastEditDateT,'') as LastEditDateT, isnull(IssueRefNo,'') as IssueRefNo, isnull(DepartCode,'') as DepartCode, IsConfirm, IsCancel, isnull(PersonCode,'') as PersonCode, SumOfAmount, isnull(MyDescription,'') as MyDescription, isnull(GLFormat,'') as GLFormat,IsCompleteSave, isnull(AllocateCode,'') as AllocateCode, isnull(ProjectCode,'') as ProjectCode, isnull(RecurName,'') as RecurName, isnull(ConfirmCode,'') as ConfirmCode, isnull(ConfirmDateTime,'') as ConfirmDateTime, isnull(CancelCode,'') as CancelCode, isnull(CancelDateTime,'') as CancelDateTime from dbo.BCStkIssueRet with (nolock) where docno = ?`
	err := db.Get(sir, sql, docno)
	if err != nil {
		return err
	}

	sqlsub := `set dateformat dmy     select MyType, DocNo, isnull(IssueRefNo,'') as IssueRefNo, ItemCode, DocDate, isnull(DepartCode,'') as DepartCode, isnull(Personcode,'') as Personcode, isnull(MyDescription,'') as MyDescription, isnull(ItemName,'') as ItemName, isnull(WHCode,'') as WHCode, isnull(ShelfCode,'') as ShelfCode, Qty, Price, Amount, HomeAmount, SumOfCost, isnull(UnitCode,'') as UnitCode, isnull(BarCode,'') as BarCode, IsCancel, LineNumber, AVERAGECOST, RefLinenumber, isnull(LotNumber,'') as LotNumber,isnull(PackingRate1,0) as PackingRate1, isnull(PackingRate2,0) as PackingRate2 from dbo.BCStkIssRetSub with (nolock) where docno = ? order by linenumber`
	err = db.Select(&sir.Subs, sqlsub, sir.DocNo)
	if err != nil {
		return err
	}

	return nil
}

func (sir *StkIssueRet) SearchStkIssueRetByKeyword(db *sqlx.DB, keyword string) (sirs []*StkIssueRet, err error) {
	sql := `set dateformat dmy     select DocNo, DocDate,isnull(CreatorCode,'') as CreatorCode, isnull(CreateDateTime,'') as CreateDateTime, isnull(LastEditorCode,'') as LastEditorCode, isnull(LastEditDateT,'') as LastEditDateT, isnull(IssueRefNo,'') as IssueRefNo, isnull(DepartCode,'') as DepartCode, IsConfirm, IsCancel, isnull(PersonCode,'') as PersonCode, SumOfAmount, isnull(MyDescription,'') as MyDescription, isnull(GLFormat,'') as GLFormat,IsCompleteSave, isnull(AllocateCode,'') as AllocateCode, isnull(ProjectCode,'') as ProjectCode, isnull(RecurName,'') as RecurName, isnull(ConfirmCode,'') as ConfirmCode, isnull(ConfirmDateTime,'') as ConfirmDateTime, isnull(CancelCode,'') as CancelCode, isnull(CancelDateTime,'') as CancelDateTime from dbo.BCStkIssueRet with (nolock) where (docno  like '%'+?+'%' or arcode like '%'+?+'%' or personcode like '%'+?+'%' ) order by docno`
	err = db.Select(&sirs, sql, keyword, keyword, keyword)
	if err != nil {
		return nil, err
	}

	for _, sub := range sirs {
		sqlsub := `set dateformat dmy     select MyType, DocNo, isnull(IssueRefNo,'') as IssueRefNo, ItemCode, DocDate, isnull(DepartCode,'') as DepartCode, isnull(Personcode,'') as Personcode, isnull(MyDescription,'') as MyDescription, isnull(ItemName,'') as ItemName, isnull(WHCode,'') as WHCode, isnull(ShelfCode,'') as ShelfCode, Qty, Price, Amount, HomeAmount, SumOfCost, isnull(UnitCode,'') as UnitCode, isnull(BarCode,'') as BarCode, IsCancel, LineNumber, AVERAGECOST, RefLinenumber, isnull(LotNumber,'') as LotNumber,isnull(PackingRate1,0) as PackingRate1, isnull(PackingRate2,0) as PackingRate2 from dbo.BCStkIssRetSub with (nolock) where docno = ? order by linenumber`
		err = db.Select(&sub.Subs, sqlsub, sub.DocNo)
		if err != nil {
			return nil, err
		}
	}

	return sirs, nil
}
