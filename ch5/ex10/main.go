// Copyright © 2016 shoarai

// The toposort program prints the nodes of a DAG in topological order.
package main

import "fmt"

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool)

	visitAll = func(items map[string]bool) {
		for k := range items {
			if !seen[k] {
				seen[k] = true
				visitAll(slice2map(m[k]))
				order = append(order, k)
			}
		}
	}

	keys := make(map[string]bool)
	for key := range m {
		keys[key] = false
	}

	visitAll(keys)
	return order
}

func slice2map(s []string) map[string]bool {
	m := make(map[string]bool)
	for _, v := range s {
		m[v] = false
	}
	return m
}
