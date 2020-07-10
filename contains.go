package main

func contains(s []bool, e bool) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
