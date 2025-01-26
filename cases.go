package caseconv

import (
	"github.com/peyton-spencer/caseconv/bytcase"
	"github.com/peyton-spencer/caseconv/strcase"
)

type Case uint8

const (
	CaseSnake Case = iota
	CaseSnakeScreaming
	CaseCamel
	CaseKebab
	CaseKebabScreaming
)

func (c Case) StrCase(s string) string {
	switch c {
	case CaseSnake:
		return strcase.ToSnake(s)
	case CaseSnakeScreaming:
		return strcase.ToScreamingSnake(s)
	case CaseCamel:
		return strcase.ToCamel(s)
	case CaseKebab:
		return strcase.ToKebab(s)
	case CaseKebabScreaming:
		return strcase.ToScreamingKebab(s)
	}
	return s
}

func (c Case) BytCase(s []byte) []byte {
	switch c {
	case CaseSnake:
		return bytcase.ToSnake(s)
	case CaseSnakeScreaming:
		return bytcase.ToScreamingSnake(s)
	case CaseCamel:
		return bytcase.ToCamel(s)
	case CaseKebab:
		return bytcase.ToKebab(s)
	case CaseKebabScreaming:
		return bytcase.ToScreamingKebab(s)
	}
	return s
}
