package rut

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrToShort          = fmt.Errorf("RUT is too short")
	ErrToLong           = fmt.Errorf("RUT is too long")
	ErrInvalidRunNumber = fmt.Errorf("RUN number is invalid")
	ErrInvalidVD        = fmt.Errorf("RUT validator digit is invalid")
)

type Rut struct {
	Rut     string
	Run     int
	VD      uint8
	IsValid bool
}

func Parse(rut string) (*Rut, error) {
	rut = cleanRut(rut)

	err := validateRutLen(rut)
	if err != nil {
		return nil, err
	}

	run, err := extractRUN(rut, err)
	if err != nil {
		return nil, err
	}

	computedVd := GenerateValidationDigit(run)
	vd := rut[len(rut)-1]
	if vd == 'k' {
		vd = 'K'
	}
	if vd != computedVd {
		return nil, ErrInvalidVD
	}
	return &Rut{
		Rut:     rut,
		Run:     run,
		VD:      computedVd,
		IsValid: true,
	}, nil
}

func extractRUN(rut string, err error) (int, error) {
	run, err := strconv.Atoi(rut[:len(rut)-1])
	if err != nil {
		return 0, ErrInvalidRunNumber
	}
	return run, nil
}

func validateRutLen(rut string) error {
	if len(rut) < 2 {
		return ErrToShort
	}
	if len(rut) > 10 {
		return ErrToLong
	}
	return nil
}

func Validate(rut string) error {
	_, err := Parse(rut)
	return err
}

func cleanRut(rut string) string {
	rut = strings.TrimSpace(rut)
	rut = strings.ReplaceAll(rut, ".", "")
	rut = strings.ReplaceAll(rut, "-", "")
	return rut
}

func GenerateValidationDigit(rut int) uint8 {
	sum := 0
	factor := 2
	for ; rut != 0; rut /= 10 {
		sum += rut % 10 * factor
		if factor == 7 {
			factor = 2
		} else {
			factor++
		}
	}

	if val := 11 - (sum % 11); val == 11 {
		return '0'
	} else if val == 10 {
		return 'K'
	} else {
		return uint8(val) + 48
	}
}
