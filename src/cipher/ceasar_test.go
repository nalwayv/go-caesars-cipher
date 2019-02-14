package cipher

import "testing"

func TestCeasar(t *testing.T) {

	secret := "SERR PBQR PNZC"
	expected := "FREE CODE CAMP"

	secret2 := "GUR DHVPX OEBJA SBK WHZCF BIRE GUR YNML QBT."
	expected2 := "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG."

	secret3 := "serr cvmmn!"
	expected3 := "FREE PIZZA!"

	if Rot13(secret) != expected {
		t.Error("failed to decipher code")
	}

	if Rot13(secret2) != expected2 {
		t.Error("failed to decipher code")
	}

	if Rot13(secret3) != expected3 {
		t.Error("failed to decipher code")
	}
}
