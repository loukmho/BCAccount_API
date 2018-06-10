package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	m "github.com/loukmho/bcaccount_api/model"
	"errors"
)

type ChqInCancel struct {
	DocNo           string         `json:"doc_no" db:"DocNo"`
	DocDate         string         `json:"doc_date" db:"DocDate"`
	CreatorCode     string         `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string         `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string         `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string         `json:"last_edit_date_t" db:"LastEditDateT"`
	MyDescription   string         `json:"my_description" db:"MyDescription"`
	BookNo          string         `json:"book_no" db:"BookNo"`
	GLFormat        string         `json:"gl_format" db:"GLFormat"`
	GLStartPosting  int            `json:"gl_start_posting" db:"GLStartPosting"`
	IsPostGL        int            `json:"is_post_gl" db:"IsPostGL"`
	GLTransData     int            `json:"gl_trans_data" db:"GLTransData"`
	SumChqAmount    float64        `json:"sum_chq_amount" db:"SumChqAmount"`
	SumExpend       float64        `json:"sum_expend" db:"SumExpend"`
	NetAmount       float64        `json:"net_amount" db:"NetAmount"`
	IsConfirm       int            `json:"is_confirm" db:"IsConfirm"`
	RecurName       string         `json:"recur_name" db:"RecurName"`
	ConfirmCode     string         `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string         `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string         `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string         `json:"cancel_date_time" db:"CancelDateTime"`
	IsCancel        int            `json:"is_cancel" db:"IsCancel"`
	UserCode        string         `json:"user_code"`
	Subs            []*ChqInRetSub `json:"subs"`
}

type ChqInCancelSub struct {
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

func (cic *ChqInCancel) InsertAndEditChqInCancel(db *sqlx.DB) error {
	var check_exist int
	var count_item int

	def := m.Default{}
	def = m.LoadDefaultData("bcdata.json")
	cic.GLFormat = def.ChqInCancelGLFormat

	sqlexist := `select count(docno) as check_exist from dbo.BCChqInCancel where docno = ?` //เช็คว่ามีเอกสารหรือยัง
	err := db.Get(&check_exist, sqlexist, cic.DocNo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	for _, sub_item := range cic.Subs {
		if (sub_item.ChqNumber != "") {
			count_item = count_item + 1
		}
	}

	switch {
	case cic.DocNo == "":
		return errors.New("Docno is null")
	case cic.DocDate == "":
		return errors.New("Docdate is null")
	case cic.IsCancel == 1:
		return errors.New("Docno is cancel")
	case cic.BookNo == "":
		return errors.New("BookNo is null")
	case cic.IsConfirm == 1:
		return errors.New("Docno is confirm")
	case count_item == 0:
		return errors.New("Docno is not have item")
	case cic.SumChqAmount == 0:
		return errors.New("SumChqAmount = 0")
	}

	if (check_exist == 0) {
		//Insert//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		cic.CreatorCode = cic.UserCode

		sql := `set dateformat dmy     insert into dbo.BCChqInCancel(DocNo,DocDate,CreatorCode,CreateDateTime,MyDescription,BookNo,GLFormat,GLStartPosting,IsPostGL,GLTransData,SumChqAmount,SumExpend,NetAmount,IsConfirm,RecurName,IsCancel) values(?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sql, cic.DocNo, cic.DocDate, cic.CreatorCode, cic.MyDescription, cic.BookNo, cic.GLFormat, cic.GLStartPosting, cic.IsPostGL, cic.GLTransData, cic.SumChqAmount, cic.SumExpend, cic.NetAmount, cic.IsConfirm, cic.RecurName, cic.IsCancel)
		if err != nil {
			fmt.Println("sql insert= ", err.Error())
			return err
		}

	} else {
		cic.LastEditorCode = cic.UserCode

		sql := `set dateformat dmy     update dbo.BCChqInCancel set DocDate=?,LastEditorCode=?,LastEditDateT=getdate(),MyDescription=?,BookNo=?,GLFormat=?,GLStartPosting=?,IsPostGL=?,GLTransData=?,SumChqAmount=?,SumExpend=?,NetAmount=?,IsConfirm=?,RecurName=?,IsCancel=? where docno = ?`
		_, err := db.Exec(sql, cic.DocDate, cic.LastEditorCode, cic.MyDescription, cic.BookNo, cic.GLFormat, cic.GLStartPosting, cic.IsPostGL, cic.GLTransData, cic.SumChqAmount, cic.SumExpend, cic.NetAmount, cic.IsConfirm, cic.RecurName, cic.IsCancel, cic.DocNo)
		if err != nil {
			fmt.Println("sql insert= ", err.Error())
			return err
		}
	}

	sql_del_sub := `delete dbo.BCChqInCancelSub where docno = ?`
	_, err = db.Exec(sql_del_sub, cic.DocNo)
	if err != nil {
		return err
	}

	var vLineNumber int

	for _, sub := range cic.Subs {
		fmt.Println("ItemSub")
		sub.LineNumber = vLineNumber
		sub.IsCancel = 0
		sub.HomeAmount = sub.NetAmount
		sub.ExchangeRate = def.ExchangeRateDefault

		sqlsub := `set dateformat dmy     insert into dbo.BCChqInCancelSub(DocNo,DocDate,BookNo,ChqRowOrder,LineNumber,TransState,IsCancel,ChqNumber,ChqAmount,Expend,NetAmount,CurrencyCode,ExchangeRate,HomeAmount,Arcode,Bankcode,BankBranchCode,RefChqRowOrder) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err := db.Exec(sqlsub, cic.DocNo, cic.DocDate, cic.BookNo, sub.ChqRowOrder, sub.LineNumber, sub.TransState, sub.IsCancel, sub.ChqNumber, sub.ChqAmount, sub.Expend, sub.NetAmount, sub.CurrencyCode, sub.ExchangeRate, sub.HomeAmount, sub.Arcode, sub.Bankcode, sub.BankBranchCode, sub.RefChqRowOrder)
		if err != nil {
			fmt.Println("sql insert sub = ", err.Error())
			return err
		}

		vLineNumber = vLineNumber + 1
	}

	return nil
}

func (cic *ChqInCancel) SearchChqInCancelByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     SELECT DocNo,DocDate,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(MyDescription,'') as MyDescription,isnull(BookNo,'') as BookNo,isnull(GLFormat,'') as GLFormat,GLStartPosting,IsPostGL,GLTransData,SumChqAmount,SumExpend,NetAmount,IsConfirm,isnull(RecurName,'') as RecurName,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime,IsCancel  FROM dbo.BCChqInCancel with (nolock) where docno = ?`
	err := db.Get(cic, sql, docno)
	if err != nil {
		fmt.Println("sql search = ", err.Error())
		return err
	}

	sqlsub := `set dateformat dmy     SELECT DocNo,DocDate,isnull(BookNo,'') as BookNo,ChqRowOrder,LineNumber,TransState,IsCancel,isnull(ChqNumber,'') as ChqNumber,ChqAmount,Expend,NetAmount,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,HomeAmount,isnull(Arcode,'') as Arcode,isnull(Bankcode,'') as Bankcode,isnull(BankBranchCode,'') as BankBranchCode,RefChqRowOrder  FROM dbo.BCChqInCancelSub with (nolock) where  Docno = ? order by LineNumber`
	err = db.Select(&cic.Subs, sqlsub, docno)
	if err != nil {
		fmt.Println("sql search sub = ", err.Error())
		return err
	}
	return nil
}


func (cic *ChqInCancel) SearchChqInCancelByKeyword(db *sqlx.DB, keyword string) (cics []*ChqInCancel, err error) {
	sql := `set dateformat dmy     SELECT DocNo,DocDate,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(MyDescription,'') as MyDescription,isnull(BookNo,'') as BookNo,isnull(GLFormat,'') as GLFormat,GLStartPosting,IsPostGL,GLTransData,SumChqAmount,SumExpend,NetAmount,IsConfirm,isnull(RecurName,'') as RecurName,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime,IsCancel  FROM dbo.BCChqInCancel with (nolock) where (docno  like '%'+?+'%' ) order by docno`
	err = db.Get(&cics, sql, keyword)
	if err != nil {
		fmt.Println("sql search = ", err.Error())
		return nil, err
	}

	for _, sub := range cics{
		sqlsub := `set dateformat dmy     SELECT DocNo,DocDate,isnull(BookNo,'') as BookNo,ChqRowOrder,LineNumber,TransState,IsCancel,isnull(ChqNumber,'') as ChqNumber,ChqAmount,Expend,NetAmount,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,HomeAmount,isnull(Arcode,'') as Arcode,isnull(Bankcode,'') as Bankcode,isnull(BankBranchCode,'') as BankBranchCode,RefChqRowOrder  FROM dbo.BCChqInCancelSub with (nolock) where  Docno = ? order by LineNumber`
		err = db.Select(cic.Subs, sqlsub, sub.DocNo)
		if err != nil {
			fmt.Println("sql search sub = ", err.Error())
			return nil, err
		}
	}

	return cics, nil
}


