package print_package

import "errors"

func PrintSomething(kata string) string {
	if kata == "ABC" {
		return "HALO ABC"
	}

	return kata + kata
}

func PrintSomething2(kata, kata2 string) string {
	return kata + " " + kata2
}

func PrintSomething3(kata, kata2 string) (string, error) {
	if kata == "" {
		return "", errors.New("Kata must be valid")
	}

	if kata2 == "" {
		return "", errors.New("Kata2 must be valid")
	}

	return kata + " " + kata2, nil
}
