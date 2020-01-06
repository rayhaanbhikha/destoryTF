package main

import "log"

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func contains(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
