package calculations

import (
	"testing"
)

//also tests peek
func TestPush(t *testing.T) {

	calc := NewCalculation("1+2+3=6")

	c := GetInstance()
	c.Push(*calc)

	var testCalc Calculation = c.Peek()
	if testCalc.equation != "1+2+3=6" {
		t.Fatal("Failed to push onto Calculations")
	}
}

func TestPeek10(t *testing.T) {

	c := GetInstance()
	for i := 0; i < 10; i++ {
		calc := NewCalculation("1+2+3=6")
		c.Push(*calc)
	}

	calculations := c.Peek10()
	if len(calculations) != 10 {
		t.Fatal("Did not get 10 caulcations when testing Peek10()")
	}

	if calculations[9].equation == "" {
		t.Fatal("Failed to make a copy of the calculations for peek10")
	}
}
