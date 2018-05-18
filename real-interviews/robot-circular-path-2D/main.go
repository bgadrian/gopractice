package robot_circular_path_2D

import (
	"log"
)

/*
Given a start position in a 2D axis coordinates, and a path of a robot with
L, R - turn left, right
G - move forward
Check if the path is circular
*/

const ( //used in pos.facing
	NORTH = iota
	EAST
	SOUTH
	WEST
)

type robot struct{ x, y, facing int } //robot position
type move func(position *robot)

var eq = func(a, b *robot) bool {
	return a.x == b.x && a.y == b.y
}

var moveFacing = map[int]move{
	NORTH: func(p *robot) { p.y++ },
	EAST:  func(p *robot) { p.x++ },
	SOUTH: func(p *robot) { p.y-- },
	WEST:  func(p *robot) { p.x-- },
}

var commands = map[rune]move{
	'L': func(p *robot) { p.facing = (p.facing + 4 - 1) % 4 },
	'R': func(p *robot) { p.facing = (p.facing + 4 + 1) % 4 },
	'G': func(p *robot) {
		mov, ok := moveFacing[p.facing]
		if ok == false {
			log.Panicf("unknown facing value '%v'", p.facing)
		}
		mov(p)
	},
}

/*
Checks only "normal" simple circular, as in the end position is the same as the start one.
*/
func normal(start *robot, moves []rune) bool {
	end := &robot{start.x, start.y, start.facing}

	for _, cmd := range moves {
		action, ok := commands[cmd]
		if ok == false {
			log.Panicf("unknown command '%v'", cmd)
		}
		action(end)

	}
	return eq(start, end)
}
