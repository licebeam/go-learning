package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 97

	fmt.Println("maps", m)

	v1 := m["k1"]
	fmt.Println("key 1", v1)

	v2 := m["k2"]
	fmt.Println("key 2", v2)

	v3 := m["k3"]
	fmt.Println("key 3", v3)

	fmt.Println("len", len(m))

	delete(m, "k2")
	fmt.Println("map", m)

	clear(m)
	fmt.Println("map", m)

	_, prs := m["k2"]
	fmt.Println("prs", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

}
