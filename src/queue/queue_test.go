package queue

import "testing"

// Note - all the meaningful tests are in internal/stackqueuebase_test

func TestQueue(t *testing.T) {
	q := NewQueue[string]()
	const want = "qwerty"
	q.Push(want)

	got, err := q.Pop()
	if err != nil {
		t.Fatalf("q.Pop() = err(%+v)", err)
	}

	if got != want {
		t.Errorf("q.Pop() = %q, want %q", got, want)
	}

	_, err = q.Pop()
	if err == nil {
		t.Errorf("q.Pop() = err(nil), want error")
	}
}
