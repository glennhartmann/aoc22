package stackqueuebase

import (
	"testing"

	"golang.org/x/exp/slices"
)

func setupQueue() *Base[string] {
	q := NewBase[string](Queue[string]{})
	q.PushN("a", "b", "3", "z")
	return q
}

func TestQueueSizeEmpty(t *testing.T) {
	type op int
	const (
		push op = iota
		pop
	)

	tests := []struct {
		name      string
		op        op
		n         int
		wantSize  int
		wantEmpty bool
	}{
		{
			name:      "default",
			op:        push,
			n:         0,
			wantSize:  4,
			wantEmpty: false,
		},
		{
			name:      "push",
			op:        push,
			n:         5,
			wantSize:  9,
			wantEmpty: false,
		},
		{
			name:      "pop",
			op:        pop,
			n:         3,
			wantSize:  1,
			wantEmpty: false,
		},
		{
			name:      "pop to 0",
			op:        pop,
			n:         4,
			wantSize:  0,
			wantEmpty: true,
		},
		{
			name:      "pop past 0",
			op:        pop,
			n:         10,
			wantSize:  0,
			wantEmpty: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q := setupQueue()

			for i := 0; i < test.n; i++ {
				if test.op == push {
					q.Push("asdf")
				} else {
					_, _ = q.Pop()
				}
			}

			gotSize := q.Size()
			if gotSize != test.wantSize {
				t.Errorf("q.Size() = %d, want %d", gotSize, test.wantSize)
			}

			gotEmpty := q.Empty()
			if gotEmpty != test.wantEmpty {
				t.Errorf("q.Empty() = %v, want %v", gotEmpty, test.wantEmpty)
			}
		})
	}
}

func TestQueuePushPeekPop(t *testing.T) {
	type op int
	const (
		push op = iota
		pop
		push1
		pop1
		push2
		pop2
		push5
		pop5
	)

	tests := []struct {
		name    string
		op      op
		want    []string
		wantErr bool
	}{
		{
			name: "push",
			op:   push,
		},
		{
			name:    "pop",
			op:      pop,
			want:    []string{"a"},
			wantErr: false,
		},
		{
			name:    "pop2",
			op:      pop2,
			want:    []string{"b", "3"},
			wantErr: false,
		},
		{
			name:    "pop5 - underflow (preserves state for next test)",
			op:      pop5,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "pop",
			op:      pop,
			want:    []string{"z"},
			wantErr: false,
		},
		{
			name: "push1",
			op:   push1,
		},
		{
			name: "push2",
			op:   push2,
		},
		{
			name: "push5",
			op:   push5,
		},
		{
			name:    "pop5",
			op:      pop5,
			want:    []string{"asdf", "asdf1", "asdf2", "asdf2", "asdf5"},
			wantErr: false,
		},
		{
			name: "push2",
			op:   push2,
		},
		{
			name:    "pop1",
			op:      pop1,
			want:    []string{"asdf5"},
			wantErr: false,
		},
		{
			name:    "pop1",
			op:      pop1,
			want:    []string{"asdf5"},
			wantErr: false,
		},
		{
			name:    "pop2",
			op:      pop2,
			want:    []string{"asdf5", "asdf5"},
			wantErr: false,
		},
		{
			name:    "pop",
			op:      pop,
			want:    []string{"asdf2"},
			wantErr: false,
		},
		{
			name:    "pop",
			op:      pop,
			want:    []string{"asdf2"},
			wantErr: false,
		},
		{
			name:    "pop - underflow",
			op:      pop,
			want:    []string{""},
			wantErr: true,
		},
	}

	q := setupQueue()
	for i, test := range tests {
		var gotPeek, gotPop []string
		var peekErr, popErr error

		switch test.op {
		case push:
			q.Push("asdf")
		case pop:
			innerGot, innerErr := q.Peek()
			gotPeek = []string{innerGot}
			peekErr = innerErr

			innerGot, innerErr = q.Pop()
			gotPop = []string{innerGot}
			popErr = innerErr
		case push1:
			q.PushN("asdf1")
		case pop1:
			gotPeek, peekErr = q.PeekN(1)
			gotPop, popErr = q.PopN(1)
		case push2:
			q.PushN("asdf2", "asdf2")
		case pop2:
			gotPeek, peekErr = q.PeekN(2)
			gotPop, popErr = q.PopN(2)
		case push5:
			q.PushN("asdf5", "asdf5", "asdf5", "asdf5", "asdf5")
		case pop5:
			gotPeek, peekErr = q.PeekN(5)
			gotPop, popErr = q.PopN(5)
		default:
			t.Fatalf("unknown op: %q", test.op)
		}

		if !slices.Equal(gotPeek, test.want) {
			t.Errorf("%d. %q: gotPeek %q, want %q", i, test.name, gotPeek, test.want)
		}

		if !slices.Equal(gotPop, test.want) {
			t.Errorf("%d. %q: gotPop %q, want %q", i, test.name, gotPop, test.want)
		}

		if (peekErr != nil) != test.wantErr {
			t.Errorf("%d. %q: got peekErr(%+v), want peekErr(%v)", i, test.name, peekErr, test.wantErr)
		}

		if (popErr != nil) != test.wantErr {
			t.Errorf("%d. %q: got popErr(%+v), want popErr(%v)", i, test.name, popErr, test.wantErr)
		}
	}
}

func TestQueueJoin(t *testing.T) {
	q := setupQueue()

	const sep = ", "
	got := q.Join(sep)
	const want1 = "a, b, 3, z"
	if got != want1 {
		t.Errorf("q.Join(%q) = %q, want %q", sep, got, want1)
	}

	q = NewBase[string](Queue[string]{})
	got = q.Join(sep)
	const want2 = ""
	if got != want2 {
		t.Errorf("q.Join(%q) = %q, want %q", sep, got, want2)
	}
}

func setupStack() *Base[string] {
	q := NewBase[string](Stack[string]{})
	q.PushN("a", "b", "3", "z")
	return q
}

func TestStackSizeEmpty(t *testing.T) {
	type op int
	const (
		push op = iota
		pop
	)

	tests := []struct {
		name      string
		op        op
		n         int
		wantSize  int
		wantEmpty bool
	}{
		{
			name:      "default",
			op:        push,
			n:         0,
			wantSize:  4,
			wantEmpty: false,
		},
		{
			name:      "push",
			op:        push,
			n:         5,
			wantSize:  9,
			wantEmpty: false,
		},
		{
			name:      "pop",
			op:        pop,
			n:         3,
			wantSize:  1,
			wantEmpty: false,
		},
		{
			name:      "pop to 0",
			op:        pop,
			n:         4,
			wantSize:  0,
			wantEmpty: true,
		},
		{
			name:      "pop past 0",
			op:        pop,
			n:         10,
			wantSize:  0,
			wantEmpty: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q := setupStack()

			for i := 0; i < test.n; i++ {
				if test.op == push {
					q.Push("asdf")
				} else {
					_, _ = q.Pop()
				}
			}

			gotSize := q.Size()
			if gotSize != test.wantSize {
				t.Errorf("q.Size() = %d, want %d", gotSize, test.wantSize)
			}

			gotEmpty := q.Empty()
			if gotEmpty != test.wantEmpty {
				t.Errorf("q.Empty() = %v, want %v", gotEmpty, test.wantEmpty)
			}
		})
	}
}

func TestStackPushPeekPop(t *testing.T) {
	type op int
	const (
		push op = iota
		pop
		push1
		pop1
		push2
		pop2
		push5
		pop5
	)

	tests := []struct {
		name    string
		op      op
		want    []string
		wantErr bool
	}{
		{
			name: "push",
			op:   push,
		},
		{
			name:    "pop",
			op:      pop,
			want:    []string{"asdf"},
			wantErr: false,
		},
		{
			name:    "pop2",
			op:      pop2,
			want:    []string{"z", "3"},
			wantErr: false,
		},
		{
			name:    "pop5 - underflow (preserves state for next test)",
			op:      pop5,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "pop",
			op:      pop,
			want:    []string{"b"},
			wantErr: false,
		},
		{
			name: "push1",
			op:   push1,
		},
		{
			name: "push2",
			op:   push2,
		},
		{
			name: "push5",
			op:   push5,
		},
		{
			name:    "pop5",
			op:      pop5,
			want:    []string{"asdf5", "asdf5", "asdf5", "asdf5", "asdf5"},
			wantErr: false,
		},
		{
			name: "push2",
			op:   push2,
		},
		{
			name:    "pop1",
			op:      pop1,
			want:    []string{"asdf2"},
			wantErr: false,
		},
		{
			name:    "pop1",
			op:      pop1,
			want:    []string{"asdf2"},
			wantErr: false,
		},
		{
			name:    "pop2",
			op:      pop2,
			want:    []string{"asdf2", "asdf2"},
			wantErr: false,
		},
		{
			name:    "pop",
			op:      pop,
			want:    []string{"asdf1"},
			wantErr: false,
		},
		{
			name:    "pop",
			op:      pop,
			want:    []string{"a"},
			wantErr: false,
		},
		{
			name:    "pop - underflow",
			op:      pop,
			want:    []string{""},
			wantErr: true,
		},
	}

	q := setupStack()
	for i, test := range tests {
		var gotPeek, gotPop []string
		var peekErr, popErr error

		switch test.op {
		case push:
			q.Push("asdf")
		case pop:
			innerGot, innerErr := q.Peek()
			gotPeek = []string{innerGot}
			peekErr = innerErr

			innerGot, innerErr = q.Pop()
			gotPop = []string{innerGot}
			popErr = innerErr
		case push1:
			q.PushN("asdf1")
		case pop1:
			gotPeek, peekErr = q.PeekN(1)
			gotPop, popErr = q.PopN(1)
		case push2:
			q.PushN("asdf2", "asdf2")
		case pop2:
			gotPeek, peekErr = q.PeekN(2)
			gotPop, popErr = q.PopN(2)
		case push5:
			q.PushN("asdf5", "asdf5", "asdf5", "asdf5", "asdf5")
		case pop5:
			gotPeek, peekErr = q.PeekN(5)
			gotPop, popErr = q.PopN(5)
		default:
			t.Fatalf("unknown op: %q", test.op)
		}

		if !slices.Equal(gotPeek, test.want) {
			t.Errorf("%d. %q: gotPeek %q, want %q", i, test.name, gotPeek, test.want)
		}

		if !slices.Equal(gotPop, test.want) {
			t.Errorf("%d. %q: gotPop %q, want %q", i, test.name, gotPop, test.want)
		}

		if (peekErr != nil) != test.wantErr {
			t.Errorf("%d. %q: got peekErr(%+v), want peekErr(%v)", i, test.name, peekErr, test.wantErr)
		}

		if (popErr != nil) != test.wantErr {
			t.Errorf("%d. %q: got popErr(%+v), want popErr(%v)", i, test.name, popErr, test.wantErr)
		}
	}
}

func TestStackJoin(t *testing.T) {
	q := setupStack()

	const sep = ", "
	got := q.Join(sep)
	const want1 = "z, 3, b, a"
	if got != want1 {
		t.Errorf("q.Join(%q) = %q, want %q", sep, got, want1)
	}

	q = NewBase[string](Stack[string]{})
	got = q.Join(sep)
	const want2 = ""
	if got != want2 {
		t.Errorf("q.Join(%q) = %q, want %q", sep, got, want2)
	}
}
