package game

import (
	"testing"
)

func checkList(expected map[Point]bool, actual []*Point, t *testing.T) bool {
	for _, p := range actual {
		if _, ok := expected[*p]; !ok {
			t.Error("Unexpected point: ", p)
			return false
		} else {
			expected[*p] = true
		}
	}
	for k, v := range expected {
		if !v {
			t.Error("Did not list point ", k)
			return false
		}
	}
	return true
}

func TestBoard(t *testing.T) {
	x := New(19)
	t.Log("\n", &x)
}

func Test2Neighbors(t *testing.T) {
	b := New(19)
	p := &Point{0, 0, '-'}
	expected := map[Point]bool{
		Point{0, 1, '-'}: false,
		Point{1, 0, '-'}: false,
	}
	ps := b.neighbors(p)
	if len(ps) != 2 {
		t.Error("Wrong amount of neighbors")
	}
	checkList(expected, ps, t)
}

func Test4Neighbors(t *testing.T) {
	b := New(19)
	p := &Point{1, 1, '-'}
	expected := map[Point]bool{
		Point{2, 1, '-'}: false,
		Point{1, 2, '-'}: false,
		Point{0, 1, '-'}: false,
		Point{1, 0, '-'}: false,
	}
	ps := b.neighbors(p)
	if len(ps) != 4 {
		t.Error("Wrong amount of neighbors")
	}
	checkList(expected, ps, t)
}

func TestFriends(t *testing.T) {
	b := New(19)
	p := &Point{1, 1, 'b'}
	b.state[0][0] = 'b'
	b.state[0][1] = 'b'
	b.state[1][0] = 'w'
	b.state[3][1] = 'b'
	expected := map[Point]bool{
		Point{0, 1, 'b'}: false,
		Point{1, 1, 'b'}: false,
		Point{0, 0, 'b'}: false,
	}
	actual := b.friends(p)
	checkList(expected, actual, t)
}

func TestAdjacentColorMap(t *testing.T) {
	b := New(19)
	b.state[0][0] = 'b'
	b.state[0][1] = 'b'
	b.state[1][1] = 'b'
	b.state[1][0] = 'w'
	b.state[3][1] = 'b'
	ps := []*Point{&Point{1, 1, 'b'}, &Point{0, 1, 'b'}}
	if acm := b.adjacentColorMap(ps); acm['b'] != 3 || acm['w'] != 1 || acm['-'] != 3 {
		t.Error(acm, "\n", b)
	}
}

func TestScore(t *testing.T) {
	b := New(19)
	b.state[1][0] = 'b'
	b.state[0][1] = 'b'
	score := b.Score()
	if score['b'] != 361 || score['w'] != 0 {
		t.Error("Score not right", score)
	}
	b.state[1][1] = 'w'
	score = b.Score()
	if score['b'] != 3 || score['w'] != 1 {
		t.Error("Score not right", score)
	}
}

func TestPlay(t *testing.T) {
	b := New(19)
	b.Play(0, 0, "white")
	b.Play(0, 1, "black")
	t.Log("\n", &b)
	b.Play(1, 0, "black")
	t.Log("\n", &b)
	b.Clear()
	b.Play(2, 2, "white")
	b.Play(1, 2, "black")
	b.Play(2, 3, "black")
	b.Play(2, 1, "black")
	b.Play(3, 2, "black")
	t.Log("\n", &b)
}

func TestCopy(t *testing.T) {
	b := New(19)
	b.Play(1, 0, "black")
	b.Play(2, 2, "white")
	b.Play(3, 2, "black")
	c := b.Copy()
	if c.String() != b.String() {
		t.Error("copies not identical")
	}

	c.Play(4, 4, "black")
	if c.String() == b.String() {
		t.Error("copies should not be identical")
	}
}

func TestUCT(t *testing.T) {
	b := New(19)
	uct(&b, "black")
}

/*func TestKomi(t *testing.T) {
	seed()
	for j := 9; j <= 19; j++ {
		b := New(j)
		sum := 0
		for i := 0; i < 1000; i++ {
			score := random_evaluate(&b, "black")
			sum += score['b'] - score['w']
		}
		t.Log(j, float32(sum)/1000)
	}
}*/
