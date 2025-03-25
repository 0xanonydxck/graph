package graph

import (
	"slices"
	"strconv"
	"strings"

	"github.com/zyedidia/generic/queue"
)

type Lock [4]int

func NewLock(s string) Lock {
	parts := strings.Split(s, "")
	lock := Lock{}
	for i := range parts {
		lock[i], _ = strconv.Atoi(parts[i])
	}

	return lock
}

func OpenLock(targetStr string, trapStrs []string) int {
	target := NewLock(targetStr)
	traps := make([]Lock, 0)
	for _, s := range trapStrs {
		traps = append(traps, NewLock(s))
	}

	getNeighbors := func(lock Lock) []Lock {
		locks := make([]Lock, 0)
		for i := range 4 {
			temp := lock
			if temp[i] < 9 {
				temp := temp
				temp[i] += 1
				locks = append(locks, temp)
			}
			if temp[i] > 0 {
				temp := temp
				temp[i] -= 1
				locks = append(locks, temp)

			}
		}

		return locks
	}

	unlock := func(start Lock) int {
		visited := map[Lock]bool{start: true}
		q := queue.New[Lock]()
		q.Enqueue(start)
		level := 0

		for !q.Empty() {
			newQ := queue.New[Lock]()
			for !q.Empty() {
				current := q.Dequeue()
				if current[0] == target[0] &&
					current[1] == target[1] &&
					current[2] == target[2] &&
					current[3] == target[3] {
					return level
				}

				for _, lock := range getNeighbors(current) {
					if !visited[lock] {
						visited[lock] = true
						if !slices.Contains(traps, lock) {
							newQ.Enqueue(lock)
						}
					}
				}
			}

			q = newQ
			level += 1
		}

		return -1
	}

	return unlock(NewLock("0000"))
}
