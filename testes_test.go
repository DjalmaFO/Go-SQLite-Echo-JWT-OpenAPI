package main

import (
	"fmt"
	"testing"
)

func TestValidateName(t *testing.T) {
	u := new(User)

	validar := []Validar{
		{"Usuario", true},
		{"", false},
		{"usuario", false},
		{"Ana", true},
		{"An a", false},
		{"Usuario user", true},
	}

	for _, v := range validar {
		u.Name = fmt.Sprintf("%v", v.Valor)
		if v.Esperado != u.ValidateName() {
			t.Fail()
		}
	}
}

func TestValidatePassword(t *testing.T) {
	u := new(User)

	validar := []Validar{
		{"senha", false},
		{"senHa@123", true},
		{"@122#A", true},
		{"Teste#45", true},
		{"", false},
		{"A@9", false},
	}

	for _, v := range validar {
		u.Password = fmt.Sprintf("%v", v.Valor)
		if v.Esperado != u.ValidatePassword() {
			t.Fail()
		}
	}
}

func TestValidateLogin(t *testing.T) {
	u := new(User)

	validar := []Validar{
		{"usuario", false},
		{"User@123", true},
		{"@122#A", true},
		{"Teste#45", true},
		{"", false},
		{"      ", false},
		{"A@9", false},
	}

	for _, v := range validar {
		u.Login = fmt.Sprintf("%v", v.Valor)
		if v.Esperado != u.ValidateLogin() {
			t.Fail()
		}
	}
}
