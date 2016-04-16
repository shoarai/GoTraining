// Copyright © 2016 shoarai

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"os"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"linear algebra": {"calculus"}, // circulation
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},

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
	ss, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "topoSort: %v\n", err)
		os.Exit(1)
	}
	for i, course := range ss {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool) bool

	visitAll = func(items map[string]bool) bool {
		for k, v := range items {
			if v {
				return false
			}
			v = true
			if !seen[k] {
				seen[k] = true
				visitAll(slice2map(m[k]))
				order = append(order, k)
			}
		}
		return true
	}

	keys := make(map[string]bool)
	for key := range m {
		keys[key] = false
	}

	if !visitAll(keys) {
		return nil, fmt.Errorf("error: circulation in topological sort")
	}
	return order, nil
}

func slice2map(s []string) map[string]bool {
	m := make(map[string]bool)
	for _, v := range s {
		m[v] = false
	}
	return m
}
