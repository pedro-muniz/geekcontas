package domain

import (
	"math/big"
	"time"

	"github.com/pedro-muniz/geekcontas/financial/domain/enum/statementType"
)

type Statement struct {
	Id            int
	PersonId      int
	ReleaseDate   time.Time
	PaymentDate   time.Time
	StatementType statementType.StatementType
	Description   string
	Value         *big.Float
	IsPaid        bool
	Installments  []int
}

func NewStatement(
	id int,
	releaseDate time.Time,
	statementType statementType.StatementType,
	description string,
	value *big.Float,

) *Statement {
	return &Statement{
		Id:            id,
		ReleaseDate:   releaseDate,
		StatementType: statementType,
		Description:   description,
		Value:         value,
	}
}
