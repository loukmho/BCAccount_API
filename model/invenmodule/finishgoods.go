package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"errors"
	m "github.com/loukmho/bcaccount_api/model"
)

type FinishGoods struct {
	DocNo           string     `json:"doc_no" db:"DocNo"`
	DocDate         string     `json:"doc_date" db:"DocDate"`
	CreatorCode     string     `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string     `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string     `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string     `json:"last_edit_date_t" db:"LastEditDateT"`
	DepartCode      string     `json:"depart_code" db:"DepartCode"`
	PersonCode      string     `json:"person_code" db:"PersonCode"`
	MyDescription   string     `json:"my_description" db:"MyDescription"`
	SumOfAmount     float64    `json:"sum_of_amount" db:"SumOfAmount"`
	IsConfirm       int        `json:"is_confirm" db:"IsConfirm"`
	GLFormat        string     `json:"gl_format" db:"GLFormat"`
	IsCancel        int        `json:"is_cancel" db:"IsCancel"`
	IsCompleteSave  int        `json:"is_complete_save" db:"IsCompleteSave"`
	SumOfCost       float64    `json:"sum_of_cost" db:"SumOfCost"`
	AllocateCode    string     `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string     `json:"project_code" db:"ProjectCode"`
	RecurName       string     `json:"recur_name" db:"RecurName"`
	ConfirmCode     string     `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string     `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string     `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string     `json:"cancel_date_time" db:"CancelDateTime"`
	UserCode        string     `json:"user_code" db:"UserCode"`
	Subs            []*FngItem `json:"subs"`
}

type FngItem struct {
	MyType         int     `json:"my_type" db:"MyType"`
	ItemCode       string  `json:"item_code" db:"ItemCode"`
	MyDescription  string  `json:"my_description" db:"MyDescription"`
	ItemName       string  `json:"item_name" db:"ItemName"`
	WHCode         string  `json:"wh_code" db:"WHCode"`
	ShelfCode      string  `json:"shelf_code" db:"ShelfCode"`
	Qty            float64 `json:"qty" db:"Qty"`
	Cost           float64 `json:"cost" db:"Cost"`
	SumOfCost      float64 `json:"sum_of_cost" db:"SumOfCost"`
	Amount         float64 `json:"amount" db:"Amount"`
	UnitCode       string  `json:"unit_code" db:"UnitCode"`
	BarCode        string  `json:"bar_code" db:"BarCode"`
	IsCancel       int     `json:"is_cancel" db:"IsCancel"`
	LineNumber     int     `json:"line_number" db:"LineNumber"`
	AverageCost    float64 `json:"average_cost" db:"AverageCost"`
	LotNumber      string  `json:"lot_number" db:"LotNumber"`
	PackingRate1   float64 `json:"packing_rate_1" db:"PackingRate1"`
	PackingRate2   float64 `json:"packing_rate_2" db:"PackingRate2"`
	GRDocNo        string  `json:"gr_doc_no" db:"GRDocNo"`
	GRRefLineNum   int     `json:"gr_ref_line_num" db:"GRRefLineNum"`
	IssueRemainQty float64 `json:"issue_remain_qty" db:"IssueRemainQty"`
}
//
//func (fng *FinishGoods) SearchFinishGoodsByDocNo() error {
//	//sql := `DocNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, DepartCode, PersonCode, MyDescription, SumOfAmount, IsConfirm, GLFormat, IsCancel,IsCompleteSave, SumOfCost, AllocateCode, ProjectCode, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime`
//	//sqlsub := `MyType, DocNo, ItemCode, DocDate,MyDescription, ItemName, WHCode, ShelfCode, Qty, Cost, SumOfCost,Amount, UnitCode, BarCode,  IsCancel, LineNumber, AVERAGECOST, LotNumber, PackingRate1, PackingRate2, GRDocNo, GRRefLineNum, IssueRemainQty`
//	return nil
//}

func (fng *FinishGoods) InsertAndEditFinishGoods(db *sqlx.DB) error {
	var check_exist int
	var count_item int

	for _, sub_item := range fng.Subs {
		if (sub_item.Qty != 0) {
			count_item = count_item + 1
		}
	}

	switch {
	case fng.DocNo == "":
		return errors.New("Docno is null")
	case fng.DocDate == "":
		return errors.New("Docdate is null")
	case fng.IsCancel == 1:
		return errors.New("Docno is cancel")
	case fng.IsConfirm == 1:
		return errors.New("Docno is confirm")
	case count_item == 0:
		return errors.New("Docno is not have item")
	}

	sqlexist := `select count(docno) as check_exist from dbo.BCFinishGoods where docno = ?` //เช็คว่ามีเอกสารหรือยัง
	err := db.Get(&check_exist, sqlexist, fng.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")
	fng.GLFormat = def.StkIssueGLFormat

	fng.IsCompleteSave = 1

	if check_exist == 0 {
		fng.CreatorCode = fng.UserCode

		sql := `set dateformat dmy     Insert into dbo.BCFinishGoods(DocNo,DocDate,CreatorCode,CreateDateTime,DepartCode,PersonCode,MyDescription,SumOfAmount,IsConfirm,GLFormat,IsCancel,IsCompleteSave,SumOfCost,AllocateCode,ProjectCode,RecurName) values(?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sql, fng.DocNo, fng.DocDate, fng.CreatorCode, fng.DepartCode, fng.PersonCode, fng.MyDescription, fng.SumOfAmount, fng.IsConfirm, fng.GLFormat, fng.IsCancel, fng.IsCompleteSave, fng.SumOfCost, fng.AllocateCode, fng.ProjectCode, fng.RecurName)
		if err != nil {
			return err
		}

	} else {
		fng.LastEditorCode = fng.UserCode

		sql := `set dateformat dmy     update dbo.BCFinishGoods set DocDate=?,LastEditorCode=?,LastEditDateT=?,DepartCode=?,PersonCode=?,MyDescription=?,SumOfAmount=?,IsConfirm=?,GLFormat=?,IsCancel=?,IsCompleteSave=?,SumOfCost=?,AllocateCode=?,ProjectCode=?,RecurName=? where docno = ?`
		_, err := db.Exec(sql, fng.DocDate, fng.LastEditorCode, fng.DepartCode, fng.PersonCode, fng.MyDescription, fng.SumOfAmount, fng.IsConfirm, fng.GLFormat, fng.IsCancel, fng.IsCompleteSave, fng.SumOfCost, fng.AllocateCode, fng.ProjectCode, fng.RecurName, fng.DocNo)
		if err != nil {
			return err
		}
	}

	sql_del_sub := `delete dbo.BCFinishGoodsSub where docno = ?`
	_, err = db.Exec(sql_del_sub, fng.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}

	var vLineNumber int

	for _, sub := range fng.Subs {
		sub.MyType = def.StkIssueMyType
		sub.LineNumber = vLineNumber

		sub.IsCancel = 0
		if (sub.PackingRate1 == 0) {
			sub.PackingRate1 = 1
		}
		if (sub.PackingRate2 == 0) {
			sub.PackingRate2 = 1
		}

		sqlsub := `set dateformat dmy     Insert into dbo.BCFinishGoodsSub(MyType,DocNo,ItemCode,DocDate,MyDescription,ItemName,WHCode,ShelfCode,Qty,Cost,SumOfCost,Amount,UnitCode,BarCode,IsCancel,LineNumber,AVERAGECOST,LotNumber,PackingRate1,PackingRate2,GRDocNo,GRRefLineNum,IssueRemainQty) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		db.Exec(sqlsub, sub.MyType, fng.DocNo, sub.ItemCode, fng.DocDate, sub.MyDescription, sub.ItemName, sub.WHCode, sub.ShelfCode, sub.Qty, sub.Cost, sub.SumOfCost, sub.Amount, sub.UnitCode, sub.BarCode, sub.IsCancel, sub.LineNumber, sub.AverageCost, sub.LotNumber, sub.PackingRate1, sub.PackingRate2, sub.GRDocNo, sub.GRRefLineNum, sub.IssueRemainQty)
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

func (fng *FinishGoods) SearchFinishGoodsByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select DocNo, DocDate, isnull(CreatorCode,'') as CreatorCode, isnull(CreateDateTime,'') as CreateDateTime, isnull(LastEditorCode,'') as LastEditorCode, isnull(LastEditDateT,'') as LastEditDateT, isnull(DepartCode,'') as DepartCode, isnull(PersonCode,'') as PersonCode, isnull(MyDescription,'') as MyDescription, SumOfAmount, IsConfirm, isnull(GLFormat,'') as GLFormat, IsCancel, IsCompleteSave, SumOfCost, isnull(AllocateCode,'') as AllocateCode, isnull(ProjectCode,'') as ProjectCode, isnull(RecurName,'') as RecurName, isnull(ConfirmCode,'') as ConfirmCode, isnull(ConfirmDateTime,'') as ConfirmDateTime, isnull(CancelCode,'') as CancelCode, isnull(CancelDateTime,'') as CancelDateTime from dbo.BCFinishGoods with (nolock) where docno = ?`
	err := db.Get(fng, sql, docno)
	if err != nil {
		return err
	}

	sqlsub := `set dateformat dmy     select MyType, DocNo, ItemCode, DocDate, isnull(MyDescription,'') as MyDescription, isnull(ItemName,'') as ItemName, isnull(WHCode,'') as WHCode, isnull(ShelfCode,'') as ShelfCode, Qty, Cost, SumOfCost, Amount, isnull(UnitCode,'') as UnitCode, isnull(BarCode,'') as BarCode, IsCancel, LineNumber, AVERAGECOST, isnull(LotNumber,'') as LotNumber, isnull(PackingRate1,0) as PackingRate1, isnull(PackingRate2,0) as PackingRate2, isnull(GRDocNo,'') as GRDocNo, GRRefLineNum, IssueRemainQty from dbo.BCFinishGoodsSub with (nolock) where docno = ? order by linenumber`
	err = db.Select(&fng.Subs, sqlsub, fng.DocNo)
	if err != nil {
		return err
	}

	return nil
}

func (fng *FinishGoods) SearchFinishGoodsByKeyword(db *sqlx.DB, keyword string) (fngs []*FinishGoods, err error) {
	sql := `set dateformat dmy     select DocNo, DocDate, isnull(CreatorCode,'') as CreatorCode, isnull(CreateDateTime,'') as CreateDateTime, isnull(LastEditorCode,'') as LastEditorCode, isnull(LastEditDateT,'') as LastEditDateT, isnull(DepartCode,'') as DepartCode, isnull(PersonCode,'') as PersonCode, isnull(MyDescription,'') as MyDescription, SumOfAmount, IsConfirm, isnull(GLFormat,'') as GLFormat, IsCancel, IsCompleteSave, SumOfCost, isnull(AllocateCode,'') as AllocateCode, isnull(ProjectCode,'') as ProjectCode, isnull(RecurName,'') as RecurName, isnull(ConfirmCode,'') as ConfirmCode, isnull(ConfirmDateTime,'') as ConfirmDateTime, isnull(CancelCode,'') as CancelCode, isnull(CancelDateTime,'') as CancelDateTime from dbo.BCFinishGoods with (nolock) where (docno  like '%'+?+'%' or arcode like '%'+?+'%' or personcode like '%'+?+'%' ) order by docno`
	err = db.Select(&fngs, sql, keyword, keyword, keyword)
	if err != nil {
		return nil, err
	}

	for _, sub := range fngs {
		sqlsub := `set dateformat dmy     select MyType, DocNo, ItemCode, DocDate, isnull(MyDescription,'') as MyDescription, isnull(ItemName,'') as ItemName, isnull(WHCode,'') as WHCode, isnull(ShelfCode,'') as ShelfCode, Qty, Cost, SumOfCost, Amount, isnull(UnitCode,'') as UnitCode, isnull(BarCode,'') as BarCode, IsCancel, LineNumber, AVERAGECOST, isnull(LotNumber,'') as LotNumber, isnull(PackingRate1,0) as PackingRate1, isnull(PackingRate2,0) as PackingRate2, isnull(GRDocNo,'') as GRDocNo, GRRefLineNum, IssueRemainQty from dbo.BCFinishGoodsSub with (nolock) where docno = ? order by linenumber`
		err = db.Select(&sub.Subs, sqlsub, sub.DocNo)
		if err != nil {
			return nil, err
		}
	}

	return sajs, nil
}
