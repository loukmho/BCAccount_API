package model

type Statement struct {
	SaveFrom        int     `json:"save_from" db:"SaveFrom"`
	Source          int     `json:"source" db:"Source"`
	DocNo           string  `json:"doc_no" db:"DocNo"`
	DocDate         string  `json:"doc_date" db:"DocDate"`
	ApCode          string  `json:"ap_code" db:"ApCode"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	DepartCode      string  `json:"depart_code" db:"DepartCode"`
	SumOfInvoice    float64 `json:"sum_of_invoice" db:"SumOfInvoice"`
	SumOfDebitNote  float64 `json:"sum_of_debit_note" db:"SumOfDebitNote"`
	SumOfCreditNote float64 `json:"sum_of_credit_note" db:"SumOfCreditNote"`
	BeforeTaxAmount float64 `json:"before_tax_amount" db:"BeforeTaxAmount"`
	TaxAmount       float64 `json:"tax_amount" db:"TaxAmount"`
	TotalAmount     float64 `json:"total_amount" db:"TotalAmount"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	BillGroup       string  `json:"bill_group" db:"BillGroup"`
	BillStatus      int     `json:"bill_status" db:"BillStatus"`
	CreditDay       int     `json:"credit_day" db:"CreditDay"`
	IsCompleteSave  int     `json:"is_complete_save" db:"IsCompleteSave"`
	DueDate         string  `json:"due_date" db:"DueDate"`
	IsConfirm       int     `json:"is_confirm" db:"IsConfirm"`
	RecurName       string  `json:"recur_name" db:"RecurName"`
	ConfirmCode     string  `json:"confirm_code" db:"ConfirmCode"`
	ConfirmDateTime string  `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string  `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string  `json:"cancel_date_time" db:"CancelDateTime"`
	IsCancel        int     `json:"is_cancel" db:"IsCancel"`
}

type stmItem struct {
	InvoiceDate   string `json:"invoice_date" db:"InvoiceDate"`
	InvoiceNo     string
	InvBalance    float64
	PayAmount     float64
	PayBalance    float64
	MyDescription string
	IsCancel      int
	LineNumber    int
	BillType      int
	HomeAmount    float64
}

func (stm *Statement) SearchStatementByDocNo() error {
	//sql := `DocNo, DocDate, ApCode, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, DepartCode, SumOfInvoice, SumOfDebitNote, SumOfCreditNote, BeforeTaxAmount, TaxAmount,TotalAmount, MyDescription, BillGroup, BillStatus, CreditDay, IsCompleteSave, DueDate, IsConfirm, RecurName, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime, IsCancel`
	//sqlsub := `DocNo, DocDate, ApCode, InvoiceDate, InvoiceNo, InvBalance, PayAmount, PayBalance, MyDescription, IsCancel,LineNumber, BillType, HomeAmount`
	return nil
}
