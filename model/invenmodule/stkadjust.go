package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	m "github.com/loukmho/bcaccount_api/model"
	"errors"
)

type StkAdjust struct {
	DocNo           string     `json:"doc_no" db:"DocNo"`
	DocDate         string     `json:"doc_date" db:"DocDate"`
	CreatorCode     string     `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string     `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string     `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string     `json:"last_edit_date_t" db:"LastEditDateT"`
	cancelcode      string     `json:"cancelcode" db:"cancelcode"`
	canceldatetime  string     `json:"canceldatetime" db:"canceldatetime"`
	InspectNo       string     `json:"inspect_no" db:"InspectNo"`
	MyDescription   string     `json:"my_description" db:"MyDescription"`
	GLFormat        string     `json:"gl_format" db:"GLFormat"`
	IsConfirm       int        `json:"is_confirm" db:"IsConfirm"`
	IsCancel        int        `json:"is_cancel" db:"IsCancel"`
	SumAmount       float64    `json:"sum_amount" db:"SumAmount"`
	SumAmount2      float64    `json:"sum_amount_2" db:"SumAmount2"`
	IsCompleteSave  int        `json:"is_complete_save" db:"IsCompleteSave"`
	TaxAmount       float64    `json:"tax_amount" db:"TaxAmount"`
	TotalAmount     float64    `json:"total_amount" db:"TotalAmount"`
	SumOfExceptTax  float64    `json:"sum_of_except_tax" db:"SumOfExceptTax"`
	OutputTaxStatus int        `json:"output_tax_status" db:"OutputTaxStatus"`
	AdjustType      int        `json:"adjust_type" db:"AdjustType"`
	ConfirmCode     string     `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string     `json:"confirm_date_time" db:"ConfirmDateTime"`
	SumOfCost       float64    `json:"sum_of_cost" db:"SumOfCost"`
	AllocateCode    string     `json:"allocate_code" db:"AllocateCode"`
	ProjectCode     string     `json:"project_code" db:"ProjectCode"`
	DepartCode      string     `json:"depart_code" db:"DepartCode"`
	UserCode        string     `json:"user_code" db:"UserCode"`
	Subs            []*AdjItem `json:"subs"`
}

type AdjItem struct {
	MyType        int     `json:"my_type" db:"MyType"`
	ItemCode      string  `json:"item_code" db:"ItemCode"`
	ItemName      string  `json:"item_name" db:"ItemName"`
	UnitCode      string  `json:"unit_code" db:"UnitCode"`
	MyDescription string  `json:"my_description" db:"MyDescription"`
	WHCode        string  `json:"wh_code" db:"WHCode"`
	ShelfCode     string  `json:"shelf_code" db:"ShelfCode"`
	IsCancel      int     `json:"is_cancel" db:"IsCancel"`
	AdjMark       int     `json:"adj_mark" db:"AdjMark"`
	Qty           float64 `json:"qty" db:"Qty"`
	AdjCost       float64 `json:"adj_cost" db:"AdjCost"`
	Amount        float64 `json:"amount" db:"Amount"`
	SumOfCost     float64 `json:"sum_of_cost" db:"SumOfCost"`
	ExceptTax     float64 `json:"except_tax" db:"ExceptTax"`
	Price         float64 `json:"price" db:"Price"`
	LineNumber    int     `json:"line_number" db:"LineNumber"`
	NetAmount     float64 `json:"net_amount" db:"NetAmount"`
	AverageCost   float64 `json:"average_cost" db:"AverageCost"`
	LotNumber     string  `json:"lot_number" db:"LotNumber"`
	BarCode       string  `json:"bar_code" db:"BarCode"`
	PackingRate1  float64 `json:"packing_rate_1" db:"PackingRate1"`
	PackingRate2  float64 `json:"packing_rate_2" db:"PackingRate2"`
}

func (sta *StkAdjust) InsertAndEditStkAdjust(db *sqlx.DB) error {
	var check_exist int
	var count_item int

	for _, sub_item := range sta.Subs {
		if (sub_item.Qty != 0) {
			count_item = count_item + 1
		}
	}

	switch {
	case sta.DocNo == "":
		return errors.New("Docno is null")
	case sta.DocDate == "":
		return errors.New("Docdate is null")
	case sta.IsCancel == 1:
		return errors.New("Docno is cancel")
	case sta.IsConfirm == 1:
		return errors.New("Docno is confirm")
	case count_item == 0:
		return errors.New("Docno is not have item")
	}

	sqlexist := `select count(docno) as check_exist from dbo.BCSTKAdjust where docno = ?` //เช็คว่ามีเอกสารหรือยัง
	err := db.Get(&check_exist, sqlexist, sta.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")
	sta.GLFormat = def.StkAdjustGLFormat

	sta.IsCompleteSave = 1

	if check_exist == 0 {
		sta.CreatorCode = sta.UserCode

		sql := `set dateformat dmy     Insert into dbo.BCSTKAdjust(DocNo,DocDate,CreatorCode,CreateDateTime,InspectNo,MyDescription,GLFormat,IsConfirm,IsCancel,SumAmount,SumAmount2,IsCompleteSave,TaxAmount,TotalAmount,SumOfExceptTax,OutputTaxStatus,AdjustType,SumOfCost,AllocateCode,ProjectCode,DepartCode) values(?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sql, sta.DocNo, sta.DocDate, sta.CreatorCode, sta.InspectNo, sta.MyDescription, sta.GLFormat, sta.IsConfirm, sta.IsCancel, sta.SumAmount, sta.SumAmount2, sta.IsCompleteSave, sta.TaxAmount, sta.TotalAmount, sta.SumOfExceptTax, sta.OutputTaxStatus, sta.AdjustType, sta.SumOfCost, sta.AllocateCode, sta.ProjectCode, sta.DepartCode)
		if err != nil {
			return err
		}

	} else {
		sta.LastEditorCode = sta.UserCode

		sql := `set dateformat dmy     update dbo.BCSTKAdjust set DocDate=?,LastEditorCode=?,LastEditDateT=getdate(),InspectNo=?,MyDescription=?,GLFormat=?,IsConfirm=?,IsCancel=?,SumAmount=?,SumAmount2=?,IsCompleteSave=?,TaxAmount=?,TotalAmount=?,SumOfExceptTax=?,OutputTaxStatus=?,AdjustType=?,SumOfCost=?,AllocateCode=?,ProjectCode=?,DepartCode=? where docno = ?`
		_, err := db.Exec(sql, sta.DocDate, sta.LastEditorCode, sta.InspectNo, sta.MyDescription, sta.GLFormat, sta.IsConfirm, sta.IsCancel, sta.SumAmount, sta.SumAmount2, sta.IsCompleteSave, sta.TaxAmount, sta.TotalAmount, sta.SumOfExceptTax, sta.OutputTaxStatus, sta.AdjustType, sta.SumOfCost, sta.AllocateCode, sta.ProjectCode, sta.DepartCode, sta.DocNo)
		if err != nil {
			return err
		}
	}

	sql_del_sub := `delete dbo.BCSTKAdjustSub where docno = ?`
	_, err = db.Exec(sql_del_sub, sta.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}

	var vLineNumber int

	for _, sub := range sta.Subs {
		sub.MyType = def.StkAdjustMyType
		sub.LineNumber = vLineNumber

		sub.IsCancel = 0
		if (sub.PackingRate1 == 0) {
			sub.PackingRate1 = 1
		}
		if (sub.PackingRate2 == 0) {
			sub.PackingRate2 = 1
		}

		sqlsub := `set dateformat dmy     Insert into dbo.BCSTKAdjustSub(MyType,DocNo,DocDate,InspectNo,ItemCode,UnitCode,MyDescription,WHCode,ShelfCode,IsCancel,AdjMark,Qty,AdjCost,Amount,SumOfCost,ExceptTax,Price,LineNumber,NetAmount,AVERAGECOST,LotNumber,BarCode,PackingRate1,PackingRate2) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		db.Exec(sqlsub, sub.MyType, sta.DocNo, sta.DocDate, sta.InspectNo, sub.ItemCode, sub.UnitCode, sub.MyDescription, sub.WHCode, sub.ShelfCode, sub.IsCancel, sub.AdjMark, sub.Qty, sub.AdjCost, sub.Amount, sub.SumOfCost, sub.ExceptTax, sub.Price, sub.LineNumber, sub.NetAmount, sub.AverageCost, sub.LotNumber, sub.BarCode, sub.PackingRate1, sub.PackingRate2)
		vLineNumber = vLineNumber + 1

		sqlprocess := ` insert into dbo.ProcessStock (ItemCode,ProcessFlag,FlowStatus) values(?, 1, 0)`
		_, err = db.Exec(sqlprocess, sub.ItemCode)
		fmt.Println("sqlprocess = ", sqlsub)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			fmt.Println(err.Error())
		}

	}

	//sql:=`cancelcode, canceldatetime, DocNo, DocDate, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, InspectNo, MyDescription, GLFormat, IsConfirm, IsCancel, SumAmount, SumAmount2, IsCompleteSave, TaxAmount, TotalAmount, SumOfExceptTax, OutputTaxStatus, AdjustType, ConfirmCode, ConfirmDateTime, SumOfCost, AllocateCode, ProjectCode, DepartCode`
	//sqlsub := `MyType, DocNo, DocDate, InspectNo, ItemCode, UnitCode, MyDescription, WHCode, ShelfCode, IsCancel, AdjMark, Qty, AdjCost, Amount, SumOfCost, ExceptTax, Price, LineNumber, NetAmount, AVERAGECOST, LotNumber,BarCode,  PackingRate1, PackingRate2`
	return nil
}

func (saj *StkAdjust) SearchStkAdjustByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select isnull(cancelcode,'') as cancelcode, isnull(canceldatetime,'') as canceldatetime, DocNo, DocDate, isnull(CreatorCode,'') as CreatorCode, isnull(CreateDateTime,'') as CreateDateTime, isnull(LastEditorCode,'') as LastEditorCode, isnull(LastEditDateT,'') as LastEditDateT, isnull(InspectNo,'') as InspectNo, isnull(MyDescription,'') as MyDescription, isnull(GLFormat,'') as GLFormat, IsConfirm, IsCancel, SumAmount, SumAmount2, IsCompleteSave, TaxAmount, TotalAmount, SumOfExceptTax, OutputTaxStatus, AdjustType, isnull(ConfirmCode,'') as ConfirmCode, isnull(ConfirmDateTime,'') as ConfirmDateTime, SumOfCost, isnull(AllocateCode,'') as AllocateCode, isnull(ProjectCode,'') as ProjectCode, isnull(DepartCode,'') as DepartCode from dbo.BCSTKAdjust with (nolock) where docno = ?`
	err := db.Get(saj, sql, docno)
	if err != nil {
		return err
	}

	sqlsub := `set dateformat dmy     select MyType, DocNo, DocDate, isnull(InspectNo,'') as InspectNo, ItemCode, b.name1 as ItemName, isnull(UnitCode,'') as UnitCode, isnull(a.MyDescription,'') as MyDescription, isnull(WHCode,'') as WHCode, isnull(ShelfCode,'') as ShelfCode, IsCancel, AdjMark, Qty, AdjCost, a.Amount, SumOfCost, a.ExceptTax, Price, LineNumber, NetAmount, a.AVERAGECOST, isnull(LotNumber,'') as LotNumber, isnull(BarCode,'') as BarCode, isnull(PackingRate1,0) as PackingRate1, isnull(PackingRate2,0) as PackingRate2 from dbo.BCSTKAdjustSub a with (nolock) left join dbo.bcitem b with (nolock) on a.itemcode = b.code where docno = ? order by linenumber`
	err = db.Select(&saj.Subs, sqlsub, saj.DocNo)
	if err != nil {
		return err
	}

	return nil
}

func (saj *StkAdjust) SearchStkAdjustByKeyword(db *sqlx.DB, keyword string) (sajs []*StkAdjust, err error) {
	sql := `set dateformat dmy     select isnull(cancelcode,'') as cancelcode, isnull(canceldatetime,'') as canceldatetime, DocNo, DocDate, isnull(CreatorCode,'') as CreatorCode, isnull(CreateDateTime,'') as CreateDateTime, isnull(LastEditorCode,'') as LastEditorCode, isnull(LastEditDateT,'') as LastEditDateT, isnull(InspectNo,'') as InspectNo, isnull(MyDescription,'') as MyDescription, isnull(GLFormat,'') as GLFormat, IsConfirm, IsCancel, SumAmount, SumAmount2, IsCompleteSave, TaxAmount, TotalAmount, SumOfExceptTax, OutputTaxStatus, AdjustType, isnull(ConfirmCode,'') as ConfirmCode, isnull(ConfirmDateTime,'') as ConfirmDateTime, SumOfCost, isnull(AllocateCode,'') as AllocateCode, isnull(ProjectCode,'') as ProjectCode, isnull(DepartCode,'') as DepartCode from dbo.BCSTKAdjust with (nolock) where (docno  like '%'+?+'%' or InspectNo like '%'+?+'%' or mydescription like '%'+?+'%' ) order by docno`
	err = db.Select(&sajs, sql, keyword, keyword, keyword)
	if err != nil {
		return nil, err
	}

	for _, sub := range sajs {
		sqlsub := `set dateformat dmy     select MyType, DocNo, DocDate, isnull(InspectNo,'') as InspectNo, ItemCode, isnull(b.name1,'') as ItemName, isnull(UnitCode,'') as UnitCode, isnull(a.MyDescription,'') as MyDescription, isnull(WHCode,'') as WHCode, isnull(ShelfCode,'') as ShelfCode, IsCancel, AdjMark, Qty, AdjCost, a.Amount, SumOfCost, a.ExceptTax, Price, LineNumber, NetAmount, a.AVERAGECOST, isnull(LotNumber,'') as LotNumber, isnull(BarCode,'') as BarCode, isnull(PackingRate1,0) as PackingRate1, isnull(PackingRate2,0) as PackingRate2 from dbo.BCSTKAdjustSub a with (nolock) left join dbo.bcitem b with (nolock) on a.itemcode = b.code where docno = ? order by linenumber`
		err = db.Select(&sub.Subs, sqlsub, sub.DocNo)
		if err != nil {
			return nil, err
		}
	}

	return sajs, nil
}
