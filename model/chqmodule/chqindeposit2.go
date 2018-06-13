package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	m "github.com/loukmho/bcaccount_api/model"
	"errors"
)

type ChqInDeposit2 struct {
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
	Subs            []*ChqInDepo2Sub `json:"subs"`
}

type ChqInDepo2Sub struct {
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


func (cid *ChqInDeposit2) InsertAndEditChqInDeposit2(db *sqlx.DB) error {
	var check_exist int
	var count_item int

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")
	cid.GLFormat = def.ChqInDeposit2GLFormat

	sqlexist := `select count(docno) as check_exist from dbo.BCChqInDeposit2 where docno = ?` //เช็คว่ามีเอกสารหรือยัง
	err := db.Get(&check_exist, sqlexist, cid.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	for _, sub_item := range cid.Subs {
		if (sub_item.ChqNumber != "") {
			count_item = count_item + 1
		}
	}

	switch {
	case cid.DocNo == "":
		return errors.New("Docno is null")
	case cid.DocDate == "":
		return errors.New("Docdate is null")
	case cid.IsCancel == 1:
		return errors.New("Docno is cancel")
	case cid.BookNo == "":
		return errors.New("BookNo is null")
	case cid.IsConfirm == 1:
		return errors.New("Docno is confirm")
	case count_item == 0:
		return errors.New("Docno is not have item")
	case cid.SumChqAmount == 0:
		return errors.New("SumChqAmount = 0")
	}

	if (check_exist == 0) {
		//Insert//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		cid.CreatorCode = cid.UserCode

		sql := `set dateformat dmy     insert into dbo.BCChqInDeposit2(DocNo,DocDate,CreatorCode,CreateDateTime,MyDescription,BookNo,GLFormat,GLStartPosting,IsPostGL,GLTransData,SumChqAmount,SumExpend,NetAmount,IsConfirm,RecurName,IsCancel) values(?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sql, cid.DocNo, cid.DocDate, cid.CreatorCode, cid.MyDescription, cid.BookNo, cid.GLFormat, cid.GLStartPosting, cid.IsPostGL, cid.GLTransData, cid.SumChqAmount, cid.SumExpend, cid.NetAmount, cid.IsConfirm, cid.RecurName, cid.IsCancel)
		if err != nil {
			fmt.Println("sql insert= ", err.Error())
			return err
		}

	} else {
		cid.LastEditorCode = cid.UserCode

		sql := `set dateformat dmy     update dbo.BCChqInDeposit2 set DocDate=?,LastEditorCode=?,LastEditDateT=getdate(),MyDescription=?,BookNo=?,GLFormat=?,GLStartPosting=?,IsPostGL=?,GLTransData=?,SumChqAmount=?,SumExpend=?,NetAmount=?,IsConfirm=?,RecurName=?,IsCancel=? where docno = ?`
		_, err := db.Exec(sql, cid.DocDate, cid.LastEditorCode, cid.MyDescription, cid.BookNo, cid.GLFormat, cid.GLStartPosting, cid.IsPostGL, cid.GLTransData, cid.SumChqAmount, cid.SumExpend, cid.NetAmount, cid.IsConfirm, cid.RecurName, cid.IsCancel, cid.DocNo)
		if err != nil {
			fmt.Println("sql insert= ", err.Error())
			return err
		}
	}

	sql_del_sub := `delete dbo.BCChqInDepo2Sub where docno = ?`
	_, err = db.Exec(sql_del_sub, cid.DocNo)
	if err != nil {
		return err
	}

	var vLineNumber int

	for _, sub := range cid.Subs {
		fmt.Println("ItemSub")
		sub.LineNumber = vLineNumber
		sub.IsCancel = 0
		sub.HomeAmount = sub.NetAmount
		sub.ExchangeRate = def.ExchangeRateDefault

		sqlsub := `set dateformat dmy     insert into dbo.BCChqInDepo2Sub(DocNo,DocDate,BookNo,ChqRowOrder,LineNumber,TransState,IsCancel,ChqNumber,ChqAmount,Expend,NetAmount,CurrencyCode,ExchangeRate,HomeAmount,Arcode,Bankcode,BankBranchCode,RefChqRowOrder) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sqlsub, cid.DocNo, cid.DocDate, cid.BookNo, sub.ChqRowOrder, sub.LineNumber, sub.TransState, sub.IsCancel, sub.ChqNumber, sub.ChqAmount, sub.Expend, sub.NetAmount, sub.CurrencyCode, sub.ExchangeRate, sub.HomeAmount, sub.Arcode, sub.Bankcode, sub.BankBranchCode, sub.RefChqRowOrder)
		if err != nil {
			fmt.Println("sql insert sub = ", err.Error())
			return err
		}

		vLineNumber = vLineNumber + 1
	}

	return nil
	//sql := `set dateformat dmy     insert into dbo.BCChqInDeposit2(DocNo,DocDate,CreatorCode,CreateDateTime,LastEditorCode,LastEditDateT,MyDescription,BookNo,GLFormat,GLStartPosting,IsPostGL,GLTransData,SumChqAmount,SumExpend,NetAmount,IsConfirm,RecurName,ConfirmCode,ConfirmDateTime,CancelCode,CancelDateTime,IsCancel) values(DocNo,DocDate,CreatorCode,CreateDateTime,LastEditorCode,LastEditDateT,MyDescription,BookNo,GLFormat,GLStartPosting,IsPostGL,GLTransData,SumChqAmount,SumExpend,NetAmount,IsConfirm,RecurName,ConfirmCode,ConfirmDateTime,CancelCode,CancelDateTime,IsCancel)`
	//
	//db.Exec(sql, cid.DocNo)
	//
	//
	//sqlsub := `set dateformat dmy     insert into dbo.BCChqInDepo2Sub(DocNo,DocDate,BookNo,ChqRowOrder,LineNumber,TransState,IsCancel,ChqNumber,ChqAmount,Expend,NetAmount,CurrencyCode,ExchangeRate,HomeAmount,Arcode,Bankcode,BankBranchCode,RefChqRowOrder) values(DocNo,DocDate,BookNo,ChqRowOrder,LineNumber,TransState,IsCancel,ChqNumber,ChqAmount,Expend,NetAmount,CurrencyCode,ExchangeRate,HomeAmount,Arcode,Bankcode,BankBranchCode,RefChqRowOrder)`
	//db.Exec(sqlsub)
	//return nil
}

//func (cid *ChqInDeposit2) SearchChqInDeposit2ByDocNo(db *sqlx.DB, docno string) error {
//	sql := `set dateformat dmy     SELECT DocNo,DocDate,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(MyDescription,'') as MyDescription,isnull(BookNo,'') as BookNo,isnull(GLFormat,'') as GLFormat,GLStartPosting,IsPostGL,GLTransData,SumChqAmount,SumExpend,NetAmount,IsConfirm,isnull(RecurName,'') as RecurName,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime,IsCancel  FROM dbo.BCChqInDeposit2 where docno = ?`
//	db.Get(sql, docno)
//
//	sqlsub := `set dateformat dmy     SELECT DocNo,DocDate,isnull(BookNo,'') as BookNo,ChqRowOrder,LineNumber,TransState,IsCancel,isnull(ChqNumber,'') as ChqNumber,ChqAmount,Expend,NetAmount,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,HomeAmount,isnull(Arcode,'') as Arcode,isnull(Bankcode,'') as Bankcode,isnull(BankBranchCode,'') as BankBranchCode,RefChqRowOrder  FROM dbo.BCChqInDepo2Sub`
//	db.Select(cid.Subs, sqlsub)
//	return nil
//}

func (cid *ChqInDeposit2) SearchChqInDeposit2ByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     SELECT DocNo,DocDate,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(MyDescription,'') as MyDescription,isnull(BookNo,'') as BookNo,isnull(GLFormat,'') as GLFormat,GLStartPosting,IsPostGL,GLTransData,SumChqAmount,SumExpend,NetAmount,IsConfirm,isnull(RecurName,'') as RecurName,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime,IsCancel  FROM dbo.BCChqInDeposit2 with (nolock) where docno = ?`
	err := db.Get(cid, sql, docno)
	if err != nil {
		fmt.Println("sql search = ", err.Error())
		return err
	}

	sqlsub := `set dateformat dmy     SELECT DocNo,DocDate,isnull(BookNo,'') as BookNo,ChqRowOrder,LineNumber,TransState,IsCancel,isnull(ChqNumber,'') as ChqNumber,ChqAmount,Expend,NetAmount,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,HomeAmount,isnull(Arcode,'') as Arcode,isnull(Bankcode,'') as Bankcode,isnull(BankBranchCode,'') as BankBranchCode,RefChqRowOrder  FROM dbo.BCChqInDepo2Sub with (nolock) where  Docno = ? order by LineNumber`
	err = db.Select(&cid.Subs, sqlsub, docno)
	if err != nil {
		fmt.Println("sql search sub = ", err.Error())
		return err
	}
	return nil
}


func (cid *ChqInDeposit2) SearchChqInDeposit2ByKeyword(db *sqlx.DB, keyword string) (cids []*ChqInDeposit2, err error) {
	sql := `set dateformat dmy     SELECT DocNo,DocDate,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(MyDescription,'') as MyDescription,isnull(BookNo,'') as BookNo,isnull(GLFormat,'') as GLFormat,GLStartPosting,IsPostGL,GLTransData,SumChqAmount,SumExpend,NetAmount,IsConfirm,isnull(RecurName,'') as RecurName,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime,IsCancel  FROM dbo.BCChqInDeposit2 with (nolock) where (docno  like '%'+?+'%' ) order by docno`
	err = db.Get(&cids, sql, keyword)
	if err != nil {
		fmt.Println("sql search = ", err.Error())
		return nil, err
	}

	for _, sub := range cids{
		sqlsub := `set dateformat dmy     SELECT DocNo,DocDate,isnull(BookNo,'') as BookNo,ChqRowOrder,LineNumber,TransState,IsCancel,isnull(ChqNumber,'') as ChqNumber,ChqAmount,Expend,NetAmount,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,HomeAmount,isnull(Arcode,'') as Arcode,isnull(Bankcode,'') as Bankcode,isnull(BankBranchCode,'') as BankBranchCode,RefChqRowOrder  FROM dbo.BCChqInDepo2Sub with (nolock) where  Docno = ? order by LineNumber`
		err = db.Select(cid.Subs, sqlsub, sub.DocNo)
		if err != nil {
			fmt.Println("sql search sub = ", err.Error())
			return nil, err
		}
	}

	return cids, nil
}