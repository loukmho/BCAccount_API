package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	m "github.com/loukmho/BCAccount_API/model"
	"errors"
)

type ChqInSale struct {
	DocNo           string          `json:"doc_no" db:"DocNo"`
	DocDate         string          `json:"doc_date" db:"DocDate"`
	CreatorCode     string          `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string          `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string          `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string          `json:"last_edit_date_t" db:"LastEditDateT"`
	MyDescription   string          `json:"my_description" db:"MyDescription"`
	BookNo          string          `json:"book_no" db:"BookNo"`
	GLFormat        string          `json:"gl_format" db:"GLFormat"`
	GLStartPosting  int             `json:"gl_start_posting" db:"GLStartPosting"`
	IsPostGL        int             `json:"is_post_gl" db:"IsPostGL"`
	GLTransData     int             `json:"gl_trans_data" db:"GLTransData"`
	SumChqAmount    float64         `json:"sum_chq_amount" db:"SumChqAmount"`
	SumExpend       float64         `json:"sum_expend" db:"SumExpend"`
	NetAmount       float64         `json:"net_amount" db:"NetAmount"`
	IsConfirm       int             `json:"is_confirm" db:"IsConfirm"`
	RecurName       string          `json:"recur_name" db:"RecurName"`
	ConfirmCode     string          `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string          `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string          `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string          `json:"cancel_date_time" db:"CancelDateTime"`
	IsCancel        int             `json:"is_cancel" db:"IsCancel"`
	UserCode        string          `json:"user_code"`
	Subs            []*ChqInRetSub `json:"subs"`
}

type ChqInSaleSub struct {
	ChqRowOrder    int     `json:"chq_row_order" db:"ChqRowOrder"`
	LineNumber     int     `json:"line_number" db:"LineNumber"`
	TransState     int     `json:"trans_state" db:"TransState"`
	IsCancel       int     `json:"is_cancel" db:"IsCancel"`
	ChqNumber      string  `json:"chq_number" db:"ChqNumber"`
	ChqAmount      float64 `json:"chq_amount" db:"ChqAmount"`
	Expend         float64 `json:"expend" db:"Expend"`
	NetAmount      float64 `json:"net_amount" db:"NetAmount"`
	CurrencyCode   string  `json:"currency_code" db:"CurrencyCode"`
	ExchangeRate   float64 `json:"exchange_rate" db:"ExchangeRate"`
	HomeAmount     float64 `json:"home_amount" db:"HomeAmount"`
	Arcode         string  `json:"arcode" db:"Arcode"`
	Bankcode       string  `json:"bankcode" db:"Bankcode"`
	BankBranchCode string  `json:"bank_branch_code" db:"BankBranchCode"`
	RefChqRowOrder int     `json:"ref_chq_row_order" db:"RefChqRowOrder"`
}

func (cis *ChqInSale) InsertAndEditChqInSale(db *sqlx.DB) error {
	var check_exist int
	var count_item int

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")
	cis.GLFormat = def.ChqInSaleGLFormat

	sqlexist := `select count(docno) as check_exist from dbo.BCChqInSale where docno = ?` //เช็คว่ามีเอกสารหรือยัง
	err := db.Get(&check_exist, sqlexist, cis.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	for _, sub_item := range cis.Subs {
		if (sub_item.ChqNumber != "") {
			count_item = count_item + 1
		}
	}

	switch {
	case cis.DocNo == "":
		return errors.New("Docno is null")
	case cis.DocDate == "":
		return errors.New("Docdate is null")
	case cis.IsCancel == 1:
		return errors.New("Docno is cancel")
	case cis.BookNo == "":
		return errors.New("BookNo is null")
	case cis.IsConfirm == 1:
		return errors.New("Docno is confirm")
	case count_item == 0:
		return errors.New("Docno is not have item")
	case cis.SumChqAmount == 0:
		return errors.New("SumChqAmount = 0")
	}

	if (check_exist == 0) {
		//Insert//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		cis.CreatorCode = cis.UserCode

		sql := `set dateformat dmy     insert into dbo.BCChqInSale(DocNo,DocDate,CreatorCode,CreateDateTime,MyDescription,BookNo,GLFormat,GLStartPosting,IsPostGL,GLTransData,SumChqAmount,SumExpend,NetAmount,IsConfirm,RecurName,IsCancel) values(?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sql, cis.DocNo, cis.DocDate, cis.CreatorCode, cis.MyDescription, cis.BookNo, cis.GLFormat, cis.GLStartPosting, cis.IsPostGL, cis.GLTransData, cis.SumChqAmount, cis.SumExpend, cis.NetAmount, cis.IsConfirm, cis.RecurName, cis.IsCancel)
		if err != nil {
			fmt.Println("sql insert= ", err.Error())
			return err
		}

	} else {
		cis.LastEditorCode = cis.UserCode

		sql := `set dateformat dmy     update dbo.BCChqInSale set DocDate=?,LastEditorCode=?,LastEditDateT=getdate(),MyDescription=?,BookNo=?,GLFormat=?,GLStartPosting=?,IsPostGL=?,GLTransData=?,SumChqAmount=?,SumExpend=?,NetAmount=?,IsConfirm=?,RecurName=?,IsCancel=? where docno = ?`
		_, err := db.Exec(sql, cis.DocDate, cis.LastEditorCode, cis.MyDescription, cis.BookNo, cis.GLFormat, cis.GLStartPosting, cis.IsPostGL, cis.GLTransData, cis.SumChqAmount, cis.SumExpend, cis.NetAmount, cis.IsConfirm, cis.RecurName, cis.IsCancel, cis.DocNo)
		if err != nil {
			fmt.Println("sql insert= ", err.Error())
			return err
		}
	}

	sql_del_sub := `delete dbo.BCChqInSaleSub where docno = ?`
	_, err = db.Exec(sql_del_sub, cis.DocNo)
	if err != nil {
		return err
	}

	var vLineNumber int

	for _, sub := range cis.Subs {
		fmt.Println("ItemSub")
		sub.LineNumber = vLineNumber
		sub.IsCancel = 0
		sub.HomeAmount = sub.NetAmount
		sub.ExchangeRate = def.ExchangeRateDefault

		sqlsub := `set dateformat dmy     insert into dbo.BCChqInSaleSub(DocNo,DocDate,BookNo,ChqRowOrder,LineNumber,TransState,IsCancel,ChqNumber,ChqAmount,Expend,NetAmount,CurrencyCode,ExchangeRate,HomeAmount,Arcode,Bankcode,BankBranchCode,RefChqRowOrder) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sqlsub, cis.DocNo, cis.DocDate, cis.BookNo, sub.ChqRowOrder, sub.LineNumber, sub.TransState, sub.IsCancel, sub.ChqNumber, sub.ChqAmount, sub.Expend, sub.NetAmount, sub.CurrencyCode, sub.ExchangeRate, sub.HomeAmount, sub.Arcode, sub.Bankcode, sub.BankBranchCode, sub.RefChqRowOrder)
		if err != nil {
			fmt.Println("sql insert sub = ", err.Error())
			return err
		}

		vLineNumber = vLineNumber + 1
	}

	return nil
}

func (cis *ChqInSale) SearchChqInSaleByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     SELECT DocNo,DocDate,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(MyDescription,'') as MyDescription,isnull(BookNo,'') as BookNo,isnull(GLFormat,'') as GLFormat,GLStartPosting,IsPostGL,GLTransData,SumChqAmount,SumExpend,NetAmount,IsConfirm,isnull(RecurName,'') as RecurName,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime,IsCancel  FROM dbo.BCChqInSale with (nolock) where docno = ?`
	err := db.Get(cis, sql, docno)
	if err != nil {
		fmt.Println("sql search = ", err.Error())
		return err
	}

	sqlsub := `set dateformat dmy     SELECT DocNo,DocDate,isnull(BookNo,'') as BookNo,ChqRowOrder,LineNumber,TransState,IsCancel,isnull(ChqNumber,'') as ChqNumber,ChqAmount,Expend,NetAmount,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,HomeAmount,isnull(Arcode,'') as Arcode,isnull(Bankcode,'') as Bankcode,isnull(BankBranchCode,'') as BankBranchCode,RefChqRowOrder  FROM dbo.BCChqInSaleSub with (nolock) where  Docno = ? order by LineNumber`
	err = db.Select(&cis.Subs, sqlsub, docno)
	if err != nil {
		fmt.Println("sql search sub = ", err.Error())
		return err
	}
	return nil
}


func (cis *ChqInSale) SearchChqInSaleByKeyword(db *sqlx.DB, keyword string) (ciss []*ChqInSale, err error) {
	sql := `set dateformat dmy     SELECT DocNo,DocDate,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(MyDescription,'') as MyDescription,isnull(BookNo,'') as BookNo,isnull(GLFormat,'') as GLFormat,GLStartPosting,IsPostGL,GLTransData,SumChqAmount,SumExpend,NetAmount,IsConfirm,isnull(RecurName,'') as RecurName,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime,IsCancel  FROM dbo.BCChqInSale with (nolock) where (docno  like '%'+?+'%' ) order by docno`
	err = db.Get(&ciss, sql, keyword)
	if err != nil {
		fmt.Println("sql search = ", err.Error())
		return nil, err
	}

	for _, sub := range ciss{
		sqlsub := `set dateformat dmy     SELECT DocNo,DocDate,isnull(BookNo,'') as BookNo,ChqRowOrder,LineNumber,TransState,IsCancel,isnull(ChqNumber,'') as ChqNumber,ChqAmount,Expend,NetAmount,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,HomeAmount,isnull(Arcode,'') as Arcode,isnull(Bankcode,'') as Bankcode,isnull(BankBranchCode,'') as BankBranchCode,RefChqRowOrder  FROM dbo.BCChqInSaleSub with (nolock) where  Docno = ? order by LineNumber`
		err = db.Select(cis.Subs, sqlsub, sub.DocNo)
		if err != nil {
			fmt.Println("sql search sub = ", err.Error())
			return nil, err
		}
	}

	return ciss, nil
}

//
//func (cis *ChqInSale) InsertAndEditChqInSale(db *sqlx.DB) error {
//	sql := `set dateformat dmy     insert into dbo.BCChqInSale(DocNo,DocDate,CreatorCode,CreateDateTime,LastEditorCode,LastEditDateT,MyDescription,BookNo,GLFormat,GLStartPosting,IsPostGL,GLTransData,SumChqAmount,SumExpend,NetAmount,IsConfirm,RecurName,ConfirmCode,ConfirmDateTime,CancelCode,CancelDateTime,IsCancel) values(DocNo,DocDate,CreatorCode,CreateDateTime,LastEditorCode,LastEditDateT,MyDescription,BookNo,GLFormat,GLStartPosting,IsPostGL,GLTransData,SumChqAmount,SumExpend,NetAmount,IsConfirm,RecurName,ConfirmCode,ConfirmDateTime,CancelCode,CancelDateTime,IsCancel)`
//	db.Exec(sql, cis.DocNo)
//
//
//	sqlsub := `set dateformat dmy     insert into dbo.BCChqInSaleSub(DocNo,DocDate,BookNo,ChqRowOrder,LineNumber,TransState,IsCancel,ChqNumber,ChqAmount,Expend,NetAmount,CurrencyCode,ExchangeRate,HomeAmount,OrgChqStatus,Arcode,Bankcode,BankBranchCode,RefChqRowOrder) values(DocNo,DocDate,BookNo,ChqRowOrder,LineNumber,TransState,IsCancel,ChqNumber,ChqAmount,Expend,NetAmount,CurrencyCode,ExchangeRate,HomeAmount,OrgChqStatus,Arcode,Bankcode,BankBranchCode,RefChqRowOrder)`
//	db.Exec(sqlsub)
//	return nil
//}
//
//func (cis *ChqInSale) SearchChqInSaleByDocNo(db *sqlx.DB, docno string) error {
//	sql := `set dateformat dmy     SELECT DocNo,DocDate,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(MyDescription,'') as MyDescription,isnull(BookNo,'') as BookNo,isnull(GLFormat,'') as GLFormat,GLStartPosting,IsPostGL,GLTransData,SumChqAmount,SumExpend,NetAmount,IsConfirm,isnull(RecurName,'') as RecurName,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime,IsCancel  FROM dbo.BCChqInSale where docno = ?`
//	db.Get(sql, docno)
//
//	sqlsub := `set dateformat dmy     SELECT DocNo,DocDate,isnull(BookNo,'') as BookNo,ChqRowOrder,LineNumber,TransState,IsCancel,isnull(ChqNumber,'') as ChqNumber,ChqAmount,Expend,NetAmount,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,HomeAmount,OrgChqStatus,isnull(Arcode,'') as Arcode,isnull(Bankcode,'') as Bankcode,isnull(BankBranchCode,'') as BankBranchCode,RefChqRowOrder  FROM dbo.BCChqInSaleSub`
//	db.Select(cis.Subs, sqlsub)
//	return nil
//}
