package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	m "github.com/loukmho/bcaccount_api/model"
	"errors"
	"time"
)

type StkIssue struct {
	DocNo           string  `json:"doc_no" db:"DocNo"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	IssueType       int     `json:"issue_type" db:"IssueType"`
	DayOfUse        int     `json:"day_of_use" db:"DayOfUse"`
	DueDate         string  `json:"due_date" db:"DueDate"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	PersonCode      string  `json:"person_code" db:"PersonCode"`
	ArCode          string  `json:"ar_code" db:"ArCode"`
	SumOfAmount     float64 `json:"sum_of_amount" db:"SumOfAmount"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
	GLFormat        string  `json:"gl_format" db:"GLFormat"`
	GLStartPosting  int     `json:"gl_start_posting" db:"GLStartPosting"`
	IsPostGL        int     `json:"is_post_gl" db:"IsPostGL"`
	IsCompleteSave  int     `json:"is_complete_save" db:"IsCompleteSave"`
	BillRetStatus   int     `json:"bill_ret_status" db:"BillRetStatus"`
	AllocateCode    string  `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string  `json:"project_code" db:"ProjectCode"`
	RecurName       string  `json:"recur_name" db:"RecurName"`
	ConfirmCode     string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string  `json:"cancel_date_time" db:"CancelDateTime"`
	UserCode        string     `json:"user_code" db:"UserCode"`
	Subs            []*IssItem `json:"subs"`
}

type IssItem struct {
	MyType        int     `json:"my_type" db:"MyType"`
	ItemCode      string  `json:"item_code" db:"ItemCode"`
	MyDescription string  `json:"my_description" db:"MyDescription"`
	ItemName      string  `json:"item_name" db:"ItemName"`
	WHCode        string  `json:"wh_code" db:"WHCode"`
	ShelfCode     string  `json:"shelf_code" db:"ShelfCode"`
	RefNo         string  `json:"ref_no" db:"RefNo"`
	Qty           float64 `json:"qty" db:"Qty"`
	RetQty        float64 `json:"ret_qty" db:"RetQty"`
	SumOfCost     float64 `json:"sum_of_cost" db:"SumOfCost"`
	Price         float64 `json:"price" db:"Price"`
	Amount        float64 `json:"amount" db:"Amount"`
	UnitCode      string  `json:"unit_code" db:"UnitCode"`
	BarCode       string  `json:"bar_code" db:"BarCode"`
	IsCancel      int     `json:"is_cancel" db:"IsCancel"`
	LineNumber    int     `json:"line_number" db:"LineNumber"`
	AverageCost   float64 `json:"average_cost" db:"AverageCost"`
	RefLineNumber int     `json:"ref_line_number" db:"RefLineNumber"`
	LotNumber     string  `json:"lot_number" db:"LotNumber"`
	PackingRate1  float64 `json:"packing_rate_1" db:"PackingRate1"`
	PackingRate2  float64 `json:"packing_rate_2" db:"PackingRate2"`
}

func (iss *StkIssue) InsertAndEditStkIssue(db *sqlx.DB) error {
	var check_exist int
	var count_item int

	now := time.Now()

	for _, sub_item := range iss.Subs {
		if (sub_item.Qty != 0) {
			count_item = count_item + 1
		}
	}

	switch {
	case iss.DocNo == "":
		return errors.New("Docno is null")
	case iss.DocDate == "":
		return errors.New("Docdate is null")
	case iss.IsCancel == 1:
		return errors.New("Docno is cancel")
	case iss.IsConfirm == 1:
		return errors.New("Docno is confirm")
	case count_item == 0:
		return errors.New("Docno is not have item")
	}


	sqlexist := `select count(docno) as check_exist from dbo.BCSTKIssue where docno = ?` //เช็คว่ามีเอกสารหรือยัง
	err := db.Get(&check_exist, sqlexist, iss.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")
	iss.GLFormat = def.StkIssueGLFormat

	iss.IsCompleteSave = 1

	if (iss.DayOfUse == 0 || iss.DueDate == "") {
		iss.DueDate = iss.DocDate
	} else {
		iss.DueDate = now.AddDate(0, 0, iss.DayOfUse).Format("2006-01-02")
	}

	if check_exist == 0 {
		iss.CreatorCode = iss.UserCode

		sql := `set dateformat dmy     Insert into dbo.BCSTKIssue(DocNo,DocDate,CreatorCode,CreateDateTime,DepartCode,IssueType,DayOfUse,DueDate,MyDescription,PersonCode,ArCode,SumOfAmount,IsConfirm,IsCancel,GLFormat,GLStartPosting,IsPostGL,IsCompleteSave,BillRetStatus,AllocateCode,ProjectCode,RecurName) values(?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sql, iss.DocNo,iss.DocDate,iss.CreatorCode,iss.DepartCode,iss.IssueType,iss.DayOfUse,iss.DueDate,iss.MyDescription,iss.PersonCode,iss.ArCode,iss.SumOfAmount,iss.IsConfirm,iss.IsCancel,iss.GLFormat,iss.GLStartPosting,iss.IsPostGL,iss.IsCompleteSave,iss.BillRetStatus,iss.AllocateCode,iss.ProjectCode,iss.RecurName)
		if err != nil {
			return err
		}

	} else {
		iss.LastEditorCode = iss.UserCode

		sql := `set dateformat dmy     update dbo.BCSTKIssue set DocDate=?,LastEditorCode=?,LastEditDateT=getdate(),DepartCode=?,IssueType=?,DayOfUse=?,DueDate=?,MyDescription=?,PersonCode=?,ArCode=?,SumOfAmount=?,IsConfirm=?,IsCancel=?,GLFormat=?,GLStartPosting=?,IsPostGL=?,IsCompleteSave=?,BillRetStatus=?,AllocateCode=?,ProjectCode=?,RecurName=? where docno = ?`
		_, err := db.Exec(sql, iss.DocDate,iss.LastEditorCode,iss.DepartCode,iss.IssueType,iss.DayOfUse,iss.DueDate,iss.MyDescription,iss.PersonCode,iss.ArCode,iss.SumOfAmount,iss.IsConfirm,iss.IsCancel,iss.GLFormat,iss.GLStartPosting,iss.IsPostGL,iss.IsCompleteSave,iss.BillRetStatus,iss.AllocateCode,iss.ProjectCode,iss.RecurName, iss.DocNo )
		if err != nil {
			return err
		}
	}

	sql_del_sub := `delete dbo.BCSTKIssueSub where docno = ?`
	_, err = db.Exec(sql_del_sub, iss.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}

	var vLineNumber int

	for _, sub := range iss.Subs {
		sub.MyType = def.StkIssueMyType
		sub.LineNumber = vLineNumber

		sub.IsCancel = 0
		if (sub.PackingRate1 == 0) {
			sub.PackingRate1 = 1
		}
		if (sub.PackingRate2 == 0) {
			sub.PackingRate2 = 1
		}
		if (sub.RetQty == 0){
			sub.RetQty = sub.Qty
		}

		sqlsub := `set dateformat dmy     Insert into dbo.BCSTKIssueSub(MyType,DocNo,IssueType,ItemCode,DocDate,DepartCode,Personcode,MyDescription,ItemName,WHCode,ShelfCode,RefNo,Qty,RetQty,SumOfCost,Price,Amount,UnitCode,BarCode,IsCancel,LineNumber,AVERAGECOST,RefLineNumber,LotNumber,PackingRate1,PackingRate2) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		db.Exec(sqlsub, sub.MyType, iss.DocNo, iss.IssueType, sub.ItemCode, iss.DocDate, iss.DepartCode, iss.PersonCode, sub.MyDescription, sub.ItemName, sub.WHCode, sub.ShelfCode, sub.RefNo, sub.Qty, sub.RetQty, sub.SumOfCost, sub.Price, sub.Amount, sub.UnitCode, sub.BarCode, sub.IsCancel, sub.LineNumber, sub.AverageCost, sub.RefLineNumber, sub.LotNumber, sub.PackingRate1, sub.PackingRate2)
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

//func (iss *StkIssue) SearchStkIssueByDocNo() error {
//	//sql := ` DocNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, DepartCode, IssueType, DayOfUse, DueDate, MyDescription, PersonCode, ArCode, SumOfAmount, IsConfirm,IsCancel, GLFormat, GLStartPosting, IsPostGL, IsCompleteSave,  BillRetStatus, AllocateCode, ProjectCode, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime `
//	//sqlsub := `MyType, DocNo,  IssueType, ItemCode, DocDate, DepartCode, Personcode, MyDescription, ItemName, WHCode,ShelfCode, RefNo, Qty, RetQty, SumOfCost, Price, Amount, UnitCode, BarCode,IsCancel, LineNumber,  AVERAGECOST, RefLineNumber, LotNumber, PackingRate1, PackingRate2`
//	return nil
//}


func (saj *StkAdjust) SearchStkIssueByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select DocNo, DocDate, isnull(CreatorCode,'') as CreatorCode, isnull(CreateDateTime,'') as CreateDateTime, isnull(LastEditorCode,'') as LastEditorCode, isnull(LastEditDateT,'') as LastEditDateT, isnull(DepartCode,'') as DepartCode, isnull(IssueType,'') as IssueType, DayOfUse, isnull(DueDate,'') as DueDate, isnull(MyDescription,'') as MyDescription, isnull(PersonCode,'') as PersonCode, isnull(ArCode,'') as ArCode, SumOfAmount, IsConfirm,IsCancel, isnull(GLFormat,'') as GLFormat, GLStartPosting, IsPostGL, IsCompleteSave,  BillRetStatus, isnull(AllocateCode,'') as AllocateCode, isnull(ProjectCode,'') as ProjectCode, isnull(RecurName,'') as RecurName, isnull(ConfirmCode,'') as ConfirmCode, isnull(ConfirmDateTime,'') as ConfirmDateTime, isnull(CancelCode,'') as CancelCode, isnull(CancelDateTime,'') as CancelDateTime from dbo.bcstkissue with (nolock) where docno = ?`
	err := db.Get(saj, sql, docno)
	if err != nil {
		return err
	}

	sqlsub := `set dateformat dmy     select MyType, DocNo, isnull(IssueType,'') as IssueType, isnull(ItemCode,'') as ItemCode, DocDate, isnull(DepartCode,'') as DepartCode, isnull(Personcode,'') as Personcode, isnull(MyDescription,'') as MyDescription, isnull(ItemName,'') as ItemName, isnull(WHCode,'') as WHCode, isnull(ShelfCode,'') as ShelfCode, isnull(RefNo,'') as RefNo, Qty, RetQty, SumOfCost, Price, Amount, isnull(UnitCode,'') as UnitCode, isnull(BarCode,'') as BarCode, IsCancel, LineNumber, AVERAGECOST, RefLineNumber, isnull(LotNumber,'') as LotNumber, PackingRate1, PackingRate2 from dbo.bcstkissuesub with (nolock) where docno = ? order by linenumber`
	err = db.Select(&saj.Subs, sqlsub, saj.DocNo)
	if err != nil {
		return err
	}

	return nil
}

func (saj *StkAdjust) SearchStkIssueByKeyword(db *sqlx.DB, keyword string) (sajs []*StkAdjust, err error) {
	sql := `set dateformat dmy     select DocNo, DocDate, isnull(CreatorCode,'') as CreatorCode, isnull(CreateDateTime,'') as CreateDateTime, isnull(LastEditorCode,'') as LastEditorCode, isnull(LastEditDateT,'') as LastEditDateT, isnull(DepartCode,'') as DepartCode, isnull(IssueType,'') as IssueType, DayOfUse, isnull(DueDate,'') as DueDate, isnull(MyDescription,'') as MyDescription, isnull(PersonCode,'') as PersonCode, isnull(ArCode,'') as ArCode, SumOfAmount, IsConfirm,IsCancel, isnull(GLFormat,'') as GLFormat, GLStartPosting, IsPostGL, IsCompleteSave,  BillRetStatus, isnull(AllocateCode,'') as AllocateCode, isnull(ProjectCode,'') as ProjectCode, isnull(RecurName,'') as RecurName, isnull(ConfirmCode,'') as ConfirmCode, isnull(ConfirmDateTime,'') as ConfirmDateTime, isnull(CancelCode,'') as CancelCode, isnull(CancelDateTime,'') as CancelDateTime from dbo.bcstkissue with (nolock) where (docno  like '%'+?+'%' or arcode like '%'+?+'%' or personcode like '%'+?+'%' ) order by docno`
	err = db.Select(saj, sql, keyword, keyword, keyword)
	if err != nil {
		return nil, err
	}

	for _, sub := range sajs {
		sqlsub := `set dateformat dmy     select MyType, DocNo, isnull(IssueType,'') as IssueType, isnull(ItemCode,'') as ItemCode, DocDate, isnull(DepartCode,'') as DepartCode, isnull(Personcode,'') as Personcode, isnull(MyDescription,'') as MyDescription, isnull(ItemName,'') as ItemName, isnull(WHCode,'') as WHCode, isnull(ShelfCode,'') as ShelfCode, isnull(RefNo,'') as RefNo, Qty, RetQty, SumOfCost, Price, Amount, isnull(UnitCode,'') as UnitCode, isnull(BarCode,'') as BarCode, IsCancel, LineNumber, AVERAGECOST, RefLineNumber, isnull(LotNumber,'') as LotNumber, PackingRate1, PackingRate2 from dbo.bcstkissuesub with (nolock) where docno = ? order by linenumber`
		err = db.Select(&sub.Subs, sqlsub, sub.DocNo)
		if err != nil {
			return nil, err
		}
	}

	return sajs, nil
}