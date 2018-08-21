package Q3_15

import "testing"

func TestPush(t *testing.T) {
	if !Push(2) {
		t.Log("stack is full\n")
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	status, _ := Pop()
	if !status {
		t.Log("stack is empty\n")
		t.Fail()
	}
}

func TestPop2(t *testing.T) {
	status, _ := Pop()
	if !status {
		t.Log("stack is empty\n")
		t.Fail()
	}
}

func TestPush2(t *testing.T) {
	Push(1)
	Push(2)
	Push(3)
	Push(4)
	Push(5)
	if !Push(6) {
		t.Log("stack is full\n")
		t.Fail()
	}
}
