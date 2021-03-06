package day13

import (
	"fmt"
	"strconv"
	"strings"
)

type point struct {
	x, y       int
	isWall     bool
	visited    bool
	neighbours []*point
	depth      int
}

type maze struct {
	start  *point
	target *point
}

type pointQueue struct {
	points []*point
}

func newPointQueue() *pointQueue {
	return &pointQueue{
		points: []*point{},
	}
}

func (pq *pointQueue) enqueue(p *point) {
	pq.points = append(pq.points, p)
}

func (pq *pointQueue) dequeue() *point {
	if pq.len() == 0 {
		panic("Nothing in queue!")
	}

	point := pq.points[0]
	pq.points = pq.points[1:]

	return point
}

func (pq *pointQueue) len() int {
	return len(pq.points)
}

func solvePartOne(rows, cols, tX, tY, designerNum int) int {
	maze, _ := makeMaze(rows, cols, tX, tY, designerNum)
	return solveMaze(maze)
}

func solvePartTwo(rows, cols, tX, tY, designerNum int) int {
	maze, allPoints := makeMaze(rows, cols, tX, tY, designerNum)
	solveMaze(maze)

	total := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			p := allPoints[row][col]
			if p.isWall {
				fmt.Print("#")
			} else if p.visited && p.depth <= 50 {
				total++
				fmt.Print(p.depth % 10)
			} else if p.visited {
				fmt.Print("/")
			} else {
				fmt.Print(".")
			}

		}
		fmt.Print("\n")
	}

	return total
}

func makeMaze(rows, cols, tX, tY, designerNum int) (*maze, [][]*point) {
	m := maze{}

	points := [][]*point{}

	for row := 0; row < rows; row++ {
		points = append(points, []*point{})
		for col := 0; col < cols; col++ {
			p := point{
				x:          col,
				y:          row,
				neighbours: []*point{},
				depth:      1000,
			}
			isWall(&p, designerNum)
			points[row] = append(points[row], &p)

			// look for start and target
			if row == 1 && col == 1 {
				p.depth = 0
				m.start = &p
			}

			if row == tY && col == tX {
				m.target = &p
			}

			// connect with surrounding points
			if !p.isWall {
				if row-1 >= 0 && !points[row-1][col].isWall {
					p.neighbours = append(p.neighbours, points[row-1][col])
					points[row-1][col].neighbours = append(points[row-1][col].neighbours, &p)
				}

				if col-1 >= 0 && !points[row][col-1].isWall {
					p.neighbours = append(p.neighbours, points[row][col-1])
					points[row][col-1].neighbours = append(points[row][col-1].neighbours, &p)
				}
			}
		}
	}

	return &m, points
}

func isWall(p *point, designerNum int) {
	prod := p.x*p.x + 3*p.x + 2*p.x*p.y + p.y + p.y*p.y + designerNum

	// to string representation of binary value
	bin := strconv.FormatInt(int64(prod), 2)

	// find only the ones
	ones := len(strings.ReplaceAll(bin, "0", ""))

	// if even number of ones, it's open, otherwise it's a wall
	p.isWall = ones%2 == 1
}

func solveMaze(m *maze) int {
	// bfs
	queue := newPointQueue()

	queue.enqueue(m.start)

	for {
		p := queue.dequeue()
		p.visited = true

		if p.x == m.target.x && p.y == m.target.y {
			return p.depth
		}

		for _, neighbour := range p.neighbours {
			if neighbour.visited == false {
				neighbour.depth = p.depth + 1
				queue.enqueue(neighbour)
			}
		}
	}
}
