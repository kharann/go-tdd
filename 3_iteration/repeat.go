package main

func Repeat(character string, repeatAmount int) string {
	var repeated string
	for i := 0; i < repeatAmount; i++ {
		repeated += character
	}
	return repeated
}
