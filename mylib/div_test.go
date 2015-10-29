package mylib

import "testing" 

func TestDiv(t *testing.T) {
    result := div (4,2)
    if result != 2 {
        t.Fatal("4/2 != 2")
    }
    result = div (16,5)
    if result != 3.2 {
        t.Fatal("16 / 5 != 3.2")
    }
}
