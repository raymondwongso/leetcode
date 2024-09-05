package main

import "fmt"

func main() {
	commands := []int{4, -1, -1, 6}
	obstacles := [][]int{
		{0, 0},
	}
	fmt.Println(robotSim(commands, obstacles))
}

// Time complexity: O(max(m,n))
// Space complexity: O(n)
// where m is len of commands, n is len of obstacles
func robotSim(commands []int, obstacles [][]int) int {
	set := map[string]struct{}{}

	for _, o := range obstacles {
		set[fmt.Sprintf("%d,%d", o[0], o[1])] = struct{}{}
	}

	dirs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	// default to face north
	res, x, y, xdelta, ydelta, face := 0, 0, 0, 0, 1, 0

	for _, c := range commands {
		if c == -1 {
			face = (face + 1) % 4
			xdelta = dirs[face][0]
			ydelta = dirs[face][1]
		} else if c == -2 {
			face = (face + 3) % 4
			xdelta = dirs[face][0]
			ydelta = dirs[face][1]
		} else {
			for range c {
				if _, ok := set[fmt.Sprintf("%d,%d", x+xdelta, y+ydelta)]; ok {
					break
				}
				x, y = x+xdelta, y+ydelta
			}

			xabs, yabs := x, y
			if xabs < 0 {
				xabs *= -1
			}
			if yabs < 0 {
				yabs *= -1
			}

			if res < (xabs*xabs + yabs*yabs) {
				res = (xabs*xabs + yabs*yabs)
			}
		}
	}

	return res
}
