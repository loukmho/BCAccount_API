package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"errors"
)

type Customer struct {
	Code           string  `json:"code" db:"Code"`
	Name1          string  `json:"name_1" db:"Name1"`
	Name2          string  `json:"name_2" db:"Name2"`
	BillAddress    string  `json:"bill_address" db:"BillAddress"`
	WorkAddress    string  `json:"work_address" db:"WorkAddress"`
	Telephone      string  `json:"telephone" db:"Telephone"`
	Fax            string  `json:"fax" db:"Fax"`
	AccountCode    string  `json:"account_code" db:"AccountCode"`
	IDCardNo       string  `json:"id_card_no" db:"IDCardNo"`
	SaleCode       string  `json:"sale_code" db:"SaleCode"`
	CreditStatus   int     `json:"credit_status" db:"CreditStatus"`
	DebtLimit1     float64 `json:"debt_limit_1" db:"DebtLimit1"`
	DebtLimit2     float64 `json:"debt_limit_2" db:"DebtLimit2"`
	DebtLimitBal   float64 `json:"debt_limit_bal" db:"DebtLimitBal"`
	TaxType        int     `json:"tax_type" db:"TaxType"`
	TaxNo          string  `json:"tax_no" db:"TaxNo"`
	PicFileName    string  `json:"pic_file_name" db:"PicFileName"`
	TypeCode       string  `json:"type_code" db:"TypeCode"`
	GroupCode      string  `json:"group_code" db:"GroupCode"`
	EmailAddress   string  `json:"email_address" db:"EmailAddress"`
	GroupOfDebt    string  `json:"group_of_debt" db:"GroupOfDebt"`
	PersonType     int     `json:"person_type" db:"PersonType"`
	BirthDay       string  `json:"birth_day" db:"BirthDay"`
	CustAge        int     `json:"cust_age" db:"CustAge"`
	CustStatus     int     `json:"cust_status" db:"CustStatus"`
	CustCreditType int     `json:"cust_credit_type" db:"CustCreditType"`
	DepartCode     string  `json:"depart_code" db:"DepartCode"`
	MEMBERID       string  `json:"memberid" db:"MEMBERID"`
	ActiveStatus   int     `json:"active_status" db:"ActiveStatus"`
	CreatorCode    string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT  string  `json:"last_edit_date_t" db:"LastEditDateT"`
	UserCode       string  `json:"user_code" db:"UserCode"`
}
//
//func (ctm *Customer) InsertAndEditCustomer(db *sqlx.DB) error {
//	var check_exist int
//
//	sqlexist := `select count(code) as check_exist from dbo.bcar where code = ?` //เช็คว่ามีเอกสารหรือยัง
//	err := db.Get(&check_exist, sqlexist, ctm.Code)
//	if err != nil {
//		fmt.Println(err.Error())
//		return nil
//	}
//
//	switch {
//	case ctm.Code == "":
//		return errors.New("Code is null")
//	case ctm.Name1 == "":
//		return errors.New("Name1 is null")
//	case ctm.BillAddress == "":
//		return errors.New("BillAddress is null")
//	}
//
//	if (check_exist == 0) {
//		//Insert//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//		ctm.CreatorCode = ctm.UserCode
//
//		sql := `insert dbo.BCAr(Code,Name1,Name2,BillAddress,WorkAddress,Telephone,Fax,AccountCode,IDCardNo,SaleCode,CreditStatus,DebtLimit1,DebtLimit2,DebtLimitBal,TaxType,TaxNo,PicFileName,TypeCode,GroupCode,EmailAddress,GroupOfDebt,PersonType,BirthDay,CustAge,CustStatus,CustCreditType,DepartCode,MEMBERID,CreatorCode,CreateDateTime,ActiveStatus) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,getdate(),1)`
//		db.Exec(sql, ctm.Code, ctm.Name1, ctm.Name2, ctm.BillAddress, ctm.WorkAddress, ctm.Telephone, ctm.Fax, ctm.AccountCode, ctm.IDCardNo, ctm.SaleCode, ctm.CreditStatus, ctm.DebtLimit1, ctm.DebtLimit2, ctm.DebtLimitBal, ctm.TaxType, ctm.TaxNo, ctm.PicFileName, ctm.TypeCode, ctm.GroupCode, ctm.EmailAddress, ctm.GroupOfDebt, ctm.PersonType, ctm.BirthDay, ctm.CustAge, ctm.CustStatus, ctm.CustCreditType, ctm.DepartCode, ctm.MEMBERID, ctm.CreatorCode)
//		if err != nil {
//			fmt.Println(err.Error())
//			return err
//		}
//	} else {
//		ctm.LastEditorCode = ctm.UserCode
//
//		sql := `update dbo.BCAr set Name1=?,Name2=?,BillAddress=?,WorkAddress=?,Telephone=?,Fax=?,AccountCode=?,IDCardNo=?,SaleCode=?,CreditStatus=?,DebtLimit1=?,DebtLimit2=?,DebtLimitBal=?,TaxType=?,TaxNo=?,PicFileName=?,TypeCode=?,GroupCode=?,EmailAddress=?,GroupOfDebt=?,PersonType=?,BirthDay=?,CustAge=?,CustStatus=?,CustCreditType=?,DepartCode=?,MEMBERID=?,LastEditorCode=?,LastEditDateT=getdate(),ActiveStatus=? where code = ?`
//		db.Exec(sql, ctm.Name1, ctm.Name2, ctm.BillAddress, ctm.WorkAddress, ctm.Telephone, ctm.Fax, ctm.AccountCode, ctm.IDCardNo, ctm.SaleCode, ctm.CreditStatus, ctm.DebtLimit1, ctm.DebtLimit2, ctm.DebtLimitBal, ctm.TaxType, ctm.TaxNo, ctm.PicFileName, ctm.TypeCode, ctm.GroupCode, ctm.EmailAddress, ctm.GroupOfDebt, ctm.PersonType, ctm.BirthDay, ctm.CustAge, ctm.CustStatus, ctm.CustCreditType, ctm.DepartCode, ctm.MEMBERID, ctm.LastEditorCode, ctm.ActiveStatus, ctm.Code)
//		if err != nil {
//			fmt.Println(err.Error())
//			return err
//		}
//	}
//	return nil
//}

func (ctm *Customer) SearchCustomerByCode(db *sqlx.DB, code string) error {

	sql := `set dateformat dmy     select Code,Name1,isnull(Name2,'') as Name2,isnull(BillAddress,'') as BillAddress,isnull(WorkAddress,'') as WorkAddress,isnull(Telephone,'') as Telephone,isnull(Fax,'') as Fax,isnull(AccountCode,'') as AccountCode,isnull(IDCardNo,'') as IDCardNo,isnull(SaleCode,'') as SaleCode,CreditStatus,DebtLimit1,DebtLimit2,DebtLimitBal,TaxType,isnull(TaxNo,'') as TaxNo,isnull(PicFileName,'') as PicFileName,isnull(TypeCode,'') as TypeCode,isnull(GroupCode,'') as GroupCode,isnull(EmailAddress,'') as EmailAddress,isnull(GroupOfDebt,'') as GroupOfDebt,PersonType,isnull(BirthDay,'') as BirthDay,CustAge,CustStatus,CustCreditType,isnull(DepartCode,'') as DepartCode,isnull(MEMBERID,'') as MEMBERID,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT, isnull(ActiveStatus,0) as ActiveStatus from dbo.bcar where code = ?`
	err := db.Get(ctm, sql, code)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (ctm *Customer) SearchCustomerByKeyword(db *sqlx.DB, keyword string) (ctmList []*Customer, err error) {
	sql := `set dateformat dmy     select Code,Name1,isnull(Name2,'') as Name2,isnull(BillAddress,'') as BillAddress,isnull(WorkAddress,'') as WorkAddress,isnull(Telephone,'') as Telephone,isnull(Fax,'') as Fax,isnull(AccountCode,'') as AccountCode,isnull(IDCardNo,'') as IDCardNo,isnull(SaleCode,'') as SaleCode,CreditStatus,DebtLimit1,DebtLimit2,DebtLimitBal,TaxType,isnull(TaxNo,'') as TaxNo,isnull(PicFileName,'') as PicFileName,isnull(TypeCode,'') as TypeCode,isnull(GroupCode,'') as GroupCode,isnull(EmailAddress,'') as EmailAddress,isnull(GroupOfDebt,'') as GroupOfDebt,PersonType,isnull(BirthDay,'') as BirthDay,CustAge,CustStatus,CustCreditType,isnull(DepartCode,'') as DepartCode,isnull(MEMBERID,'') as MEMBERID,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT, isnull(ActiveStatus,0) as ActiveStatus from dbo.bcar where (code  like '%'+?+'%' or name1 like '%'+?+'%' or BillAddress like '%'+?+'%' or memberid like '%'+?+'%')`
	err = db.Select(&ctmList, sql, keyword, keyword, keyword, keyword)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return ctmList, nil
}

func (ctm *Customer) InsertAndEditCustomer(db *sqlx.DB) error {
	var check_exist int

	sqlexist := `select count(code) as check_exist from dbo.bcar where code = ?`
	err := db.Get(&check_exist, sqlexist, ctm.Code)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	switch {
	case ctm.Code == "":
		return errors.New("code is null")
	case ctm.Name1 == "":
		return errors.New("Name1 is null")
	case ctm.BillAddress == "":
		return errors.New("address is null")
	}

	if (ctm.WorkAddress==""){
		ctm.WorkAddress = ctm.BillAddress
	}


	if (check_exist == 0) {
		ctm.ActiveStatus = 1

		sql := `Insert into dbo.bcar(Code,Name1,Name2,BillAddress,WorkAddress,Telephone,Fax,AccountCode,IDCardNo,SaleCode,CreditStatus,DebtLimit1,DebtLimit2,DebtLimitBal,TaxType,TaxNo,PicFileName,TypeCode,GroupCode,EmailAddress,GroupOfDebt,PersonType,BirthDay,CustAge,CustStatus,CustCreditType,DepartCode,MEMBERID,CreatorCode,CreateDateTime,ActiveStatus) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,getdate(),?)`
		_, err = db.Exec(sql, ctm.Code, ctm.Name1, ctm.Name2, ctm.BillAddress, ctm.WorkAddress, ctm.Telephone, ctm.Fax, ctm.AccountCode, ctm.IDCardNo, ctm.SaleCode, ctm.CreditStatus, ctm.DebtLimit1, ctm.DebtLimit2, ctm.DebtLimitBal, ctm.TaxType, ctm.TaxNo, ctm.PicFileName, ctm.TypeCode, ctm.GroupCode, ctm.EmailAddress, ctm.GroupOfDebt, ctm.PersonType, ctm.BirthDay, ctm.CustAge, ctm.CustStatus, ctm.CustCreditType, ctm.DepartCode, ctm.MEMBERID, ctm.CreatorCode, ctm.ActiveStatus)
	} else {
		sql := `Update dbo.bcar set Name1=?,Name2=?,BillAddress=?,WorkAddress=?,Telephone=?,Fax=?,AccountCode=?,IDCardNo=?,SaleCode=?,CreditStatus=?,DebtLimit1=?,DebtLimit2=?,DebtLimitBal=?,TaxType=?,TaxNo=?,PicFileName=?,TypeCode=?,GroupCode=?,EmailAddress=?,GroupOfDebt=?,PersonType=?,BirthDay=?,CustAge=?,CustStatus=?,CustCreditType=?,DepartCode=?,MEMBERID=?,LastEditorCode=?,LastEditDateT=getdate(),ActiveStatus=? where code = ?`
		_, err = db.Exec(sql, ctm.Name1, ctm.Name2, ctm.BillAddress, ctm.WorkAddress, ctm.Telephone, ctm.Fax, ctm.AccountCode, ctm.IDCardNo, ctm.SaleCode, ctm.CreditStatus, ctm.DebtLimit1, ctm.DebtLimit2, ctm.DebtLimitBal, ctm.TaxType, ctm.TaxNo, ctm.PicFileName, ctm.TypeCode, ctm.GroupCode, ctm.EmailAddress, ctm.GroupOfDebt, ctm.PersonType, ctm.BirthDay, ctm.CustAge, ctm.CustStatus, ctm.CustCreditType, ctm.DepartCode, ctm.MEMBERID, ctm.CreatorCode, ctm.ActiveStatus, ctm.Code)
	}
	if err != nil {
		fmt.Println("error = ", err.Error())
		return err
	}

	return nil
}
