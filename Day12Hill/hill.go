package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	area, startRow, startCol, err := readInput()
	if err != nil {
		log.Fatal(err)
	}
	rd1, err := round1(area, startRow, startCol)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Round 1: %d\n", rd1)

	rd2, err := round2(area, startRow, startCol)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Round 2: %d\n", rd2)

}

// S = 83, E=69 a = 97
func round1(area [][]Elevation, startRow int, startCol int) (int, error) {
	distances := make(map[key]int)
	queue := make([]Position, 0)

	queue = append(queue, Position{
		key: key{
			x: startCol,
			y: startRow,
		},
		distance: 0,
	})
	for {
		curr := queue[0]
		queue = queue[1:]
		currDistance := distances[key{
			x: curr.x,
			y: curr.y,
		}]
		x := curr.x + 1
		y := curr.y
		if isVaild(area, x, y) {
			if area[y][x].height <= area[curr.y][curr.x].height+1 {
				if area[y][x].isFinish {
					return currDistance + 1, nil
				}
				if oldDistance := distances[key{
					x: x,
					y: y,
				}]; oldDistance == 0 || oldDistance > currDistance+1 {
					distances[key{
						x: x,
						y: y,
					}] = currDistance + 1
					queue = append(queue, Position{
						key: key{
							x: x,
							y: y,
						},
						distance: currDistance + 1,
					})
				}
			}
		}
		x = curr.x - 1
		y = curr.y
		if isVaild(area, x, y) {
			if area[y][x].height <= area[curr.y][curr.x].height+1 {
				if area[y][x].isFinish {
					return currDistance + 1, nil
				}
				if oldDistance := distances[key{
					x: x,
					y: y,
				}]; oldDistance == 0 || oldDistance > currDistance+1 {
					distances[key{
						x: x,
						y: y,
					}] = currDistance + 1
					queue = append(queue, Position{
						key: key{
							x: x,
							y: y,
						},
						distance: currDistance + 1,
					})
				}
			}
		}
		x = curr.x
		y = curr.y + 1
		if isVaild(area, x, y) {
			if area[y][x].height <= area[curr.y][curr.x].height+1 {
				if area[y][x].isFinish {
					return currDistance + 1, nil
				}
				if oldDistance := distances[key{
					x: x,
					y: y,
				}]; oldDistance == 0 || oldDistance > currDistance+1 {
					distances[key{
						x: x,
						y: y,
					}] = currDistance + 1
					queue = append(queue, Position{
						key: key{
							x: x,
							y: y,
						},
						distance: currDistance + 1,
					})
				}
			}
		}
		x = curr.x
		y = curr.y - 1
		if isVaild(area, x, y) {
			if area[y][x].height <= area[curr.y][curr.x].height+1 {
				if area[y][x].isFinish {
					return currDistance + 1, nil
				}
				if oldDistance := distances[key{
					x: x,
					y: y,
				}]; oldDistance == 0 || oldDistance > currDistance+1 {
					distances[key{
						x: x,
						y: y,
					}] = currDistance + 1
					queue = append(queue, Position{
						key: key{
							x: x,
							y: y,
						},
						distance: currDistance + 1,
					})
				}
			}
		}
	}
}

// S = 83, E=69 a = 97
func round2(area [][]Elevation, startRow int, startCol int) (int, error) {
	distances := make(map[key]int)
	minDistance := 9999999999999
	queue := make([]Position, 0)

	queue = append(queue, Position{
		key: key{
			x: startCol,
			y: startRow,
		},
		distance: 0,
	})
	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		currDistance := distances[key{
			x: curr.x,
			y: curr.y,
		}]
		x := curr.x + 1
		y := curr.y
		if isVaild(area, x, y) {
			if area[y][x].height <= area[curr.y][curr.x].height+1 {
				if area[y][x].isFinish {
					if currDistance+1 < minDistance {
						minDistance = currDistance + 1
					}
				}
				if oldDistance := distances[key{
					x: x,
					y: y,
				}]; oldDistance == 0 || oldDistance > currDistance+1 {
					if point := area[y][x]; point.height == 97 && oldDistance == 0 {
						if !point.traveled {
							point.traveled = true
							area[y][x] = point
							queue = append(queue, Position{
								key: key{
									x: x,
									y: y,
								},
								distance: 0,
							})
						}
					} else {
						distances[key{
							x: x,
							y: y,
						}] = currDistance + 1
						queue = append(queue, Position{
							key: key{
								x: x,
								y: y,
							},
							distance: currDistance + 1,
						})
					}
				}
			}
		}
		x = curr.x - 1
		y = curr.y
		if isVaild(area, x, y) {
			if area[y][x].height <= area[curr.y][curr.x].height+1 {
				if area[y][x].isFinish {
					if currDistance+1 < minDistance {
						minDistance = currDistance + 1
					}
				}
				if oldDistance := distances[key{
					x: x,
					y: y,
				}]; oldDistance == 0 || oldDistance > currDistance+1 {
					if point := area[y][x]; point.height == 97 && oldDistance == 0 {
						if !point.traveled {
							point.traveled = true
							area[y][x] = point
							queue = append(queue, Position{
								key: key{
									x: x,
									y: y,
								},
								distance: 0,
							})
						}
					} else {
						distances[key{
							x: x,
							y: y,
						}] = currDistance + 1
						queue = append(queue, Position{
							key: key{
								x: x,
								y: y,
							},
							distance: currDistance + 1,
						})
					}
				}
			}
		}
		x = curr.x
		y = curr.y + 1
		if isVaild(area, x, y) {
			if area[y][x].height <= area[curr.y][curr.x].height+1 {
				if area[y][x].isFinish {
					if currDistance+1 < minDistance {
						minDistance = currDistance + 1
					}
				}
				if oldDistance := distances[key{
					x: x,
					y: y,
				}]; oldDistance == 0 || oldDistance > currDistance+1 {
					if point := area[y][x]; point.height == 97 && oldDistance == 0 {
						if !point.traveled {
							point.traveled = true
							area[y][x] = point
							queue = append(queue, Position{
								key: key{
									x: x,
									y: y,
								},
								distance: 0,
							})
						}
					} else {
						distances[key{
							x: x,
							y: y,
						}] = currDistance + 1
						queue = append(queue, Position{
							key: key{
								x: x,
								y: y,
							},
							distance: currDistance + 1,
						})
					}
				}
			}
		}
		x = curr.x
		y = curr.y - 1
		if isVaild(area, x, y) {
			if area[y][x].height <= area[curr.y][curr.x].height+1 {
				if area[y][x].isFinish {
					if currDistance+1 < minDistance {
						minDistance = currDistance + 1
					}
				}
				if oldDistance := distances[key{
					x: x,
					y: y,
				}]; oldDistance == 0 || oldDistance > currDistance+1 {
					if point := area[y][x]; point.height == 97 && oldDistance == 0 {
						if !point.traveled {
							point.traveled = true
							area[y][x] = point
							queue = append(queue, Position{
								key: key{
									x: x,
									y: y,
								},
								distance: 0,
							})
						}
					} else {
						distances[key{
							x: x,
							y: y,
						}] = currDistance + 1
						queue = append(queue, Position{
							key: key{
								x: x,
								y: y,
							},
							distance: currDistance + 1,
						})
					}
				}
			}
		}
	}
	return minDistance, nil
}

func readInput() ([][]Elevation, int, int, error) {
	res := make([][]Elevation, 0)
	f, err := os.Open("in.txt")

	if err != nil {
		return nil, 0, 0, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	j := 0
	startRow := 0
	startCol := 0
	for scanner.Scan() {
		text := scanner.Text()
		row := make([]Elevation, 0)
		for i, char := range text {
			if char == 'E' {
				row = append(row, Elevation{
					height:   int('z'),
					isFinish: true,
				})
			} else {
				val := int(char)
				if int(val) == 83 {
					startCol = i
					startRow = j
					val = 'z'
				}

				row = append(row, Elevation{
					height:   val,
					isFinish: false,
				})

			}

		}
		j++
		res = append(res, row)
	}
	return res, startRow, startCol, nil
}

type key struct {
	x, y int
}

type Position struct {
	key
	distance int
}

func isVaild(area [][]Elevation, x int, y int) bool {
	return x >= 0 && y >= 0 && x < len(area[0]) && y < len(area)
}

type Elevation struct {
	height   int
	isFinish bool
	traveled bool
}
