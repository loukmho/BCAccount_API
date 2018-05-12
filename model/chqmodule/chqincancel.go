package model

import "github.com/jmoiron/sqlx"

type ChqInCancel struct {
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
	sql := `set dateformat dmy     insert into dbo.BCChqInCancel(DocNo,DocDate,CreatorCode,CreateDateTime,LastEditorCode,LastEditDateT,MyDescription,BookNo,GLFormat,GLStartPosting,IsPostGL,GLTransData,SumChqAmount,SumExpend,NetAmount,IsConfirm,RecurName,ConfirmCode,ConfirmDateTime,CancelCode,CancelDateTime,IsCancel) values(DocNo,DocDate,CreatorCode,CreateDateTime,LastEditorCode,LastEditDateT,MyDescription,BookNo,GLFormat,GLStartPosting,IsPostGL,GLTransData,SumChqAmount,SumExpend,NetAmount,IsConfirm,RecurName,ConfirmCode,ConfirmDateTime,CancelCode,CancelDateTime,IsCancel)`

	db.Exec(sql, cic.DocNo)


	sqlsub := `set dateformat dmy     insert into dbo.BCChqInCancelSub(DocNo,DocDate,BookNo,ChqRowOrder,LineNumber,TransState,IsCancel,ChqNumber,ChqAmount,Expend,NetAmount,CurrencyCode,ExchangeRate,HomeAmount,Arcode,Bankcode,BankBranchCode,RefChqRowOrder) values(DocNo,DocDate,BookNo,ChqRowOrder,LineNumber,TransState,IsCancel,ChqNumber,ChqAmount,Expend,NetAmount,CurrencyCode,ExchangeRate,HomeAmount,Arcode,Bankcode,BankBranchCode,RefChqRowOrder)`
	db.Exec(sqlsub)
	return nil
}

func (cic *ChqInCancel) SearchChqInCancelByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     SELECT DocNo,DocDate,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(MyDescription,'') as MyDescription,isnull(BookNo,'') as BookNo,isnull(GLFormat,'') as GLFormat,GLStartPosting,IsPostGL,GLTransData,SumChqAmount,SumExpend,NetAmount,IsConfirm,isnull(RecurName,'') as RecurName,isnull(ConfirmCode,'') as ConfirmCode,isnull(ConfirmDateTime,'') as ConfirmDateTime,isnull(CancelCode,'') as CancelCode,isnull(CancelDateTime,'') as CancelDateTime,IsCancel  FROM dbo.BCChqInCancel where docno = ?`
	db.Get(sql, docno)

	sqlsub := `set dateformat dmy     SELECT DocNo,DocDate,isnull(BookNo,'') as BookNo,ChqRowOrder,LineNumber,TransState,IsCancel,isnull(ChqNumber,'') as ChqNumber,ChqAmount,Expend,NetAmount,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,HomeAmount,isnull(Arcode,'') as Arcode,isnull(Bankcode,'') as Bankcode,isnull(BankBranchCode,'') as BankBranchCode,RefChqRowOrder  FROM dbo.BCChqInCancelSub`
	db.Select(cic.Subs, sqlsub)
	return nil
}