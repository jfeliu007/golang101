package main
import "testing"

func TestSum(t *testing.T) {
    result := sum(1,1)
    if result != 2 {
        t.Fatal("1+1 != 2")
    }
}
