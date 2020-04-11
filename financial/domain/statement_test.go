package domain

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	statementType "github.com/pedro-muniz/geekcontas/financial/domain/enum/statementType"
)

func TestConstructor(t *testing.T) {
	d := NewStatement(
		1,
		time.Now(),
		statementType.Revenue,
		"Single test",
		new(big.Float).SetPrec(2).SetFloat64(10.0),
	)

	fmt.Printf("Statement Object -> %v\n", *d)
}
