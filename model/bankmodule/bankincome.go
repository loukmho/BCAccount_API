package model


type BankInCome struct {
	DocNo,
	DocDate,
	BookNo,
	CreatorCode,
	CreateDateTime,
	LastEditorCode,
	LastEditDateT,
	AccountCode,
	GLBookCode,
	DepartCode,
	MyDescription,
	Amount,
	AllocateCode,
	ProjectCode,
	IsCancel,
	IsConfirm,
	ConfirmCode,
	ConfirmDateTime,
	CancelCode,
	CancelDateTime,
	RecurName
}

func (bic *BankInCome) SearchBankInCome() error {
	sql := `DocNo, DocDate, BookNo, CreatorCode, CreateDateTime, LastEditorCode, LastEditDateT, AccountCode, GLBookCode,DepartCode, MyDescription, Amount, AllocateCode, ProjectCode,IsCancel, IsConfirm, ConfirmCode, ConfirmDateTime, CancelCode, CancelDateTime, RecurName from dbo.BCBankIncome WITH (NOLOCK)`
	return nil
}
