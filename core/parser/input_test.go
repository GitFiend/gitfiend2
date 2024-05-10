package parser

import "testing"

func TestWhenInputIsCreatedExpectStartPosition0(t *testing.T) {
	input := Input{Code: "omg"}

	if input.Position != 0 {
		t.Errorf("Expected new Input object to start at postion 0, got: %d.", input.Position)
	}
}

func TestWhenAdvanceCalledOnInputIncrementsPosition(t *testing.T) {
	input := Input{Code: "omg"}

	if input.Position != 0 {
		t.Errorf("Expected new Input object to start at postion 0, got: %d.", input.Position)
	}

	input.Advance()

	if input.Position != 1 {
		t.Error("Expected i.Position to 1 after increment, got ", input.Position)
	}
}

func TestInput_Rest(t *testing.T) {
	input := Input{Code: "abcdefgHijKlmnop"}

	input.SetPosition(10)
	rest := input.Rest(2)

	if rest != "Kl" {
		t.Error("Expected Kl, got ", rest)
	}
}

func TestInput_End(t *testing.T) {
	input := Input{Code: ""}

	if !input.End() {
		t.Error("Expected input of length 0 to already be at the end")
	}
}
