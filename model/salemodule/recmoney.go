package model

import (
)

type RecMoney struct {
	Data []ListData
}

type ListData struct {
	DocNo string `json:"doc_no" db:"DocNo"`
	DocDate string `json:"doc_date" db:"DocDate"`
	ArCode string `json:"ar_code" db:"ArCode"`
	ExchangeRate int `json:"exchange_rate" db:"ExchangeRate"`
	PayAmount float64 `json:"pay_amount" db:"PayAmount"`
	ChqTotalAmount float64 `json:"chq_total_amount" db:"ChqTotalAmount"`
	PaymentType int `json:"payment_type" db:"PaymentType"`
	CreditType int `json:"credit_type" db:"CreditType"`
	ConfirmNo string `json:"confirm_no" db:"ConfirmNo"`
	RefNo string `json:"ref_no" db:"RefNo"`
	BankCode string `json:"bank_code" db:"BookCode"`
	ProjectCode string `json:"project_code" db:"ProjectCode"`
	DepartCode string `json:"depart_code" db:"DepartCode"`
	SaleCode string `json:"sale_code" db:"SaleCode"`
	BankBranchCode string `json:"bank_branch_code" db:"BankBranchCode"`
	TransBankDate string `json:"trans_bank_date" db:"TransBankDate"`
	RefDate string `json:"ref_date" db:"RefDate"`
}

