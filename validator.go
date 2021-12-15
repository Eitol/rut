package rut

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrToShort       = fmt.Errorf("RUT is too short")
	ErrToLong        = fmt.Errorf("RUT is too long")
	ErrInvalidNumber = fmt.Errorf("RUT number is invalid")
	ErrInvalidVD     = fmt.Errorf("RUT validator digit is invalid")
)

func Validate(rut string) error {
	rut = cleanRut(rut)
	if len(rut) < 2 {
		return ErrToShort
	}
	if len(rut) > 10 {
		return ErrToLong
	}
	intVal, err := strconv.Atoi(rut[:len(rut)-1])
	if err != nil {
		return ErrInvalidNumber
	}
	calcVd := GenerateValidatorDigit(intVal)
	vd := rut[len(rut)-1]
	if vd == 'k' {
		vd = 'K'
	}
	if vd != calcVd {
		return ErrInvalidVD
	}
	return nil
}

func cleanRut(rut string) string {
	rut = strings.TrimSpace(rut)
	rut = strings.ReplaceAll(rut, ".", "")
	rut = strings.ReplaceAll(rut, "-", "")
	return rut
}

func GenerateValidatorDigit(rut int) uint8 {
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
