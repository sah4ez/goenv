package goenv

import "testing"

func TestParseString(t *testing.T) {
	l := Loader{}

	line := "TEST=\"Value\""

	exp := Var{
		Name:  "TEST",
		Value: "Value",
		Type:  "string",
	}

	_, act, err := l.parseLine(line)

	if err != nil {
		t.Fatalf("Error: %e", err)
	}

	if act.Name != exp.Name {
		t.Fatalf("Name wrong: %s exptected: %s", act.Name, exp.Name)
	}
	if act.Value != exp.Value {
		t.Fatalf("Value wrong: %s exptected: %s", act.Value, exp.Value)
	}
	if act.Type != exp.Type {
		t.Fatalf("Type wrong: %s exptected: %s", act.Type, exp.Type)
	}
}

func TestParseInt(t *testing.T) {
	l := Loader{}

	line := "TEST  =  \"123\"  #  int"

	exp := Var{
		Name:  "TEST",
		Value: "123",
		Type:  "int",
	}

	_, act, err := l.parseLine(line)

	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	if act.Name != exp.Name {
		t.Fatalf("Name wrong: %s exptected: %s", act.Name, exp.Name)
	}
	if act.Value != exp.Value {
		t.Fatalf("Value wrong: %s exptected: %s", act.Value, exp.Value)
	}
	if act.Type != exp.Type {
		t.Fatalf("Type wrong: %s exptected: %s", act.Type, exp.Type)
	}
}

func TestParseInvalidType(t *testing.T) {
	l := Loader{}

	line := "TEST  =  \"123\"  #  invalidType"

	_, _, err := l.parseLine(line)

	if err == nil {
		t.Fatalf("Test should fail")
	}
}
