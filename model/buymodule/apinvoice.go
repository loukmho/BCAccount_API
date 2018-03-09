package model

import (
	"github.com/jmoiron/sqlx"
)

type ApInvoice struct {
	DocNo string `json:"doc_no" db:"DocNo"`
	//ApCode string `json:"ap_code" db:"ApCode"`
	Personal
	CreatorCode string `json:"creator_code" db:"CreatorCode"`
}

type Personal struct {
	ApCode string `json:"ap_code" db:"ApCode"`
//	CreatorCode string `json:"creator_code" db:"CreatorCode"`
}

func (ai *ApInvoice)SearchApInvoiceByDocNo(db *sqlx.DB) error{
	//sql := `select DocNo,ApCode as ApCode,CreatorCode from dbo.bcapinvoice where docno = 'CGP6001-0001'`
	////err := db.Query(ai, sql)
	//fmt.Println("sql =", sql)
	//if err != nil {
	//	return err
	//}
	return nil
}
