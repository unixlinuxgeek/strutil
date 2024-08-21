// Упражнение 4.5
//
// Напишите функцию, которая без выделения дополнительной памяти удаляет все смежные дубликаты в срезе []string.

package strutil

import (
	"slices"
	"strings"
)

// RemDup removing duplicate elements
func RemDup(s []string) []string {
	t := strings.Join(s, "")

	x := s[:0]
	for u, v := range s {
		y := strings.Count(t, v)
		if y == 1 {
			x = append(x, v)
		} else if !slices.Contains(x, s[u]) {
			x = append(x, s[u])
		}
	}
	return x
}
