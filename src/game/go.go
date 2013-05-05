package game

var colors = map[string]byte{
	"black": 'b',
	"white": 'w',
	"empty": '-',
}

var enemies = map[string]string{
	"black": "white",
	"white": "black",
}

type Board struct {
	state [][]byte
	size  int
}

func New(size int) Board {
	state := make([][]byte, size)
	for i := range state {
		state[i] = make([]byte, size)
		for j := range state[i] {
			state[i][j] = colors["empty"]
		}
	}
	return Board{state, size}
}

func (board *Board) onBoard(p *Point) bool {
	return p.row >= 0 && p.col >= 0 && p.row < board.size && p.col < board.size
}

func (board *Board) getPoint(row int, col int) (*Point, bool) {
	ok := false
	p := &Point{row, col, '0'}
	if board.onBoard(p) {
		p.c = board.state[row][col]
		ok = true
	}
	return p, ok
}

func (board *Board) neighbors(p *Point) []*Point {
	nbrs := make([]*Point, 0, 4)
	if p, ok := (board.getPoint(p.row+1, p.col)); ok {
		nbrs = append(nbrs, p)
	}
	if p, ok := (board.getPoint(p.row, p.col+1)); ok {
		nbrs = append(nbrs, p)
	}
	if p, ok := (board.getPoint(p.row-1, p.col)); ok {
		nbrs = append(nbrs, p)
	}
	if p, ok := (board.getPoint(p.row, p.col-1)); ok {
		nbrs = append(nbrs, p)
	}
	return nbrs
}

func (board *Board) associates(p *Point, c byte) []*Point {
	friends := make([]*Point, 0, 1)
	ret := make([]*Point, 0, 1)
	friends = append(friends, p)
	if p.c == c {
		ret = append(ret, p)
	}
	seen := make(map[Point]bool)
	seen[*p] = true
	for len(friends) > 0 {
		friend := friends[0]
		friends = friends[1:]
		nbrs := board.neighbors(friend)
		for _, nbr := range nbrs {
			if !seen[*nbr] && nbr.c == c {
				friends = append(friends, nbr)
				ret = append(ret, nbr)
			}
			seen[*nbr] = true
		}
	}
	return ret
}

func (board *Board) friends(p *Point) []*Point {
	return board.associates(p, p.c)
}

func (board *Board) adjacentColorMap(group []*Point) map[byte]int {
	byColor := make(map[byte]map[Point]bool)
	for _, color := range colors {
		byColor[color] = make(map[Point]bool)
	}
	for _, p := range group {
		for _, nbr := range board.neighbors(p) {
			byColor[nbr.c][*nbr] = true
		}
	}
	countByColor := make(map[byte]int)
	for k, v := range byColor {
		countByColor[k] = len(v)
	}
	return countByColor
}

func (board *Board) liberties(p *Point) ([]*Point, int) {
	friends := board.friends(p)
	return friends, board.adjacentColorMap(friends)[colors["empty"]]
}

/* ----------- PUBLIC ------------*/

func (board *Board) Score() map[byte]int {
	visited := make([][]bool, board.size)
	for i := 0; i < board.size; i++ {
		visited[i] = make([]bool, board.size)
	}
	scores := make(map[byte]int)
	for i := 0; i < board.size; i++ {
		for j := 0; j < board.size; j++ {
			if !visited[i][j] {
				p, _ := board.getPoint(i, j)
				visited[i][j] = true
				if p.c == colors["empty"] {
					alsoEmpty := board.friends(p)
					for _, e := range alsoEmpty {
						visited[e.row][e.col] = true
					}
					adjacent := board.adjacentColorMap(alsoEmpty)
					if adjacent[colors["black"]] == 0 {
						scores[colors["white"]] += len(alsoEmpty)
					} else if adjacent[colors["white"]] == 0 {
						scores[colors["black"]] += len(alsoEmpty)
					}
				} else {
					scores[p.c] += 1
				}
			}
		}
	}
	return scores
}

func (board *Board) Clear() {
	for row := 0; row < board.size; row++ {
		for col := 0; col < board.size; col++ {
			board.state[row][col] = colors["empty"]
		}
	}
}

func (board *Board) Play(row int, col int, c string) bool {
	if _, ok := colors[c]; !ok {
		return false
	}
	if p, ok := board.getPoint(row, col); ok && p.c == colors["empty"] {
		p.c = colors[c]
		if _, l := board.liberties(p); l > 1 {
			board.state[row][col] = colors[c]
			seen := make(map[Point]bool)
			possibles := board.associates(p, colors[enemies[c]])
			for _, e := range possibles {
				seen[*e] = true
			}
			for k, _ := range seen {
				group, liberties := board.liberties(&k)
				killed := liberties == 0
				for _, g := range group {
					if killed {
						board.state[g.row][g.col] = colors["empty"]
					}
					delete(seen, *g)
				}
			}
			return true
		}
	}
	return false
}

func (board Board) String() string {
	s := ""
	for _, row := range board.state {
		s = s + string(row) + "\n"
	}
	return s
}
