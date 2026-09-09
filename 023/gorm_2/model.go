package main

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int            `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Bank struct {
	BaseModel `gorm:"embedded"`

	Code    string `gorm:"size:10;unique;not null" json:"code"`
	Name    string `gorm:"size:30;not null" json:"name"`
	Address string `gorm:"type:text;not null" json:"address"`

	Branches []Branch `gorm:"foreignKey:BranchBankCode;references:Code;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"branches"`
}

type Branch struct {
	BaseModel `gorm:"embedded"`

	BranchBankCode string `gorm:"size:10;not null;uniqueIndex:bank_branch" json:"branch_bank_code"`
	BranchNumber   string `gorm:"size:10;not null;uniqueIndex:bank_branch" json:"branch_number"`
	Address        string `gorm:"type:text;not null" json:"address"`

	Bank Bank `gorm:"foreignKey:BranchBankCode;references:Code;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"bank"`

	Accounts []Account `gorm:"foreignKey:AccBankCode,AccBranchNumber;references:BranchBankCode,BranchNumber;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"accounts"`
	Loans    []Loan    `gorm:"foreignKey:LoanBankCode,LoanBranchNumber;references:BranchBankCode,BranchNumber;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"loans"`
}

type AccountType string

const (
	ProMax  AccountType = "pro-max"
	Pro     AccountType = "pro"
	Trial   AccountType = "trial"
	Limited AccountType = "limited"
)

type Account struct {
	BaseModel `gorm:"embedded"`

	AccBankCode     string `gorm:"size:10;not null;uniqueIndex:branch_account" json:"acc_bank_code"`
	AccBranchNumber string `gorm:"size:10;not null;uniqueIndex:branch_account" json:"acc_branch_number"`
	AccountNumber   string `gorm:"size:10;not null;uniqueIndex:branch_account" json:"account_number"`

	Balance int         `json:"balance"`
	Type    AccountType `json:"type"`

	Branch           Branch            `gorm:"foreignKey:AccBankCode,AccBranchNumber;references:BranchBankCode,BranchNumber;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"branch"`
	CustomerAccounts []CustomerAccount `gorm:"foreignKey:CA_BankCode,CA_BranchNumber,CA_AccountNumber;references:AccBankCode,AccBranchNumber,AccountNumber;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null" json:"customer_accounts"`
}

type LoanType string

const (
	Business   LoanType = "business"
	Government LoanType = "government"
	Political  LoanType = "political"
)

type Loan struct {
	BaseModel `gorm:"embedded"`

	LoanBankCode     string `gorm:"size:10;not null;uniqueIndex:branch_loan" json:"loan_bank_code"`
	LoanBranchNumber string `gorm:"size:10;not null;uniqueIndex:branch_loan" json:"loan_branch_number"`
	LoanNumber       string `gorm:"size:10;not null;uniqueIndex:branch_loan" json:"loan_number"`

	Amount int      `gorm:"not null" json:"amount"`
	Type   LoanType `gorm:"type:enum('business','government','political');not null" json:"type"`

	Branch Branch `gorm:"foreignKey:LoanBankCode,LoanBranchNumber;references:BranchBankCode,BranchNumber;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"branch"`

	CustomerLoans []CustomerLoan `gorm:"foreignKey:CL_BankCode,CL_BranchNumber,CL_LoanNumber;references:LoanBankCode,LoanBranchNumber,LoanNumber;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"customer_loans"`
}

type Customer struct {
	BaseModel `gorm:"embedded"`

	Ssn     string `gorm:"size:10;uniqueIndex;not null" json:"ssn"`
	Name    string `gorm:"size:30;not null" json:"name"`
	Address string `gorm:"type:text;not null" json:"address"`
	Phone   string `gorm:"size:11;not null" json:"phone"`

	Accounts []CustomerAccount `gorm:"foreignKey:CA_CustomerSsn;references:Ssn;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"customer_accounts"`
	Loans    []CustomerLoan    `gorm:"foreignKey:CL_CustomerSsn;references:Ssn;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"customer_loans"`
}

type CustomerAccount struct {
	BaseModel `gorm:"embedded"`

	CA_BankCode      string `gorm:"size:10;not null;uniqueIndex:acc_cus" json:"ca_bank_code"`
	CA_AccountNumber string `gorm:"size:10;not null;uniqueIndex:acc_cus" json:"ca_account_number"`
	CA_BranchNumber  string `gorm:"size:10;not null;uniqueIndex:acc_cus" json:"ca_branch_number"`

	CA_CustomerSsn string `gorm:"size:10;not null;uniqueIndex:acc_cus" json:"ca_customer_ssn"`

	Customer Customer `gorm:"foreignKey:CA_CustomerSsn;references:Ssn;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null" json:"customer"`

	Account Account `gorm:"foreignKey:CA_BankCode,CA_BranchNumber,CA_AccountNumber;references:AccBankCode,AccBranchNumber,AccountNumber;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null" json:"account"`
}

type CustomerLoan struct {
	BaseModel `gorm:"embedded"`

	CL_BankCode     string `gorm:"size:10;not null;uniqueIndex:loan_cus" json:"cl_bank_code"`
	CL_BranchNumber string `gorm:"size:10;not null;uniqueIndex:loan_cus" json:"cl_branch_number"`
	CL_LoanNumber   string `gorm:"size:10;not null;uniqueIndex:loan_cus" json:"cl_loan_number"`

	CL_CustomerSsn string `gorm:"size:10;not null;uniqueIndex:loan_cus" json:"cl_customer_ssn"`

	Customer Customer `gorm:"foreignKey:CL_CustomerSsn;references:Ssn;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null" json:"customer"`

	Loan Loan `gorm:"foreignKey:CL_BankCode,CL_BranchNumber,CL_LoanNumber;references:LoanBankCode,LoanBranchNumber,LoanNumber;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null" json:"loan"`
}
