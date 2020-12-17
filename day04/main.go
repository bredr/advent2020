package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n\n")
	for i, line := range lines {
		lines[i] = strings.ReplaceAll(line, "\n", " ")
		lines[i] = strings.Trim(lines[i], " ")
	}
	return lines, nil
}

func main() {
	data, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}
	valids := 0
	for _, row := range data {
		kvs := strings.Split(row, " ")
		var keys []string
		for _, kv := range kvs {
			keyValue := strings.Split(kv, ":")
			keys = append(keys, keyValue[0])
		}
		expectedKeys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
		invalid := false
		for _, ek := range expectedKeys {
			if !contains(keys, ek) {
				invalid = true
			}
		}
		if !invalid {
			valids++
		}
	}
	fmt.Printf("Found %d valid\n", valids)

	strictValids := 0
	for _, row := range data {
		kvs := strings.Split(row, " ")
		var keys []string
		passport := make(map[string]string)
		for _, kv := range kvs {
			keyValue := strings.Split(kv, ":")
			keys = append(keys, keyValue[0])
			if len(keyValue) == 2 {
				passport[keyValue[0]] = strings.Trim(keyValue[1], " \n")
			}
		}
		expectedKeys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
		invalid := false
		for _, ek := range expectedKeys {
			if !contains(keys, ek) {
				invalid = true
			}
		}
		if invalid {
			continue
		}
		if !validated(passport) {
			invalid = true
		}
		if !invalid {
			strictValids++
		}
	}
	fmt.Printf("Found %d with strict validation\n", strictValids)
}

func contains(arr []string, k string) bool {
	for _, a := range arr {
		if a == k {
			return true
		}
	}
	return false
}

func validated(p map[string]string) bool {
	for k, v := range p {
		switch k {
		case "byr":
			y, err := strconv.Atoi(v)
			if err != nil {
				return false
			}
			if y < 1920 || y > 2002 {
				return false
			}
		case "iyr":
			y, err := strconv.Atoi(v)
			if err != nil {
				return false
			}
			if y < 2010 || y > 2020 {
				return false
			}
		case "eyr":
			y, err := strconv.Atoi(v)
			if err != nil {
				return false
			}
			if y < 2020 || y > 2030 {
				return false
			}
		case "hgt":
			validH1 := regexp.MustCompile(`([0-9]{3})cm`)
			validH2 := regexp.MustCompile(`([0-9]{2})in`)
			if validH1.MatchString(v) {
				h, err := strconv.Atoi(strings.ReplaceAll(v, "cm", ""))
				if err != nil {
					fmt.Printf("Invalid %s:%s %s\n", k, v, err.Error())
					return false
				}
				if h < 150 || h > 193 {
					fmt.Printf("Invalid %s:%s %d\n", k, v, h)
					return false
				}
				continue
			} else if validH2.MatchString(v) {
				h, err := strconv.Atoi(strings.ReplaceAll(v, "in", ""))
				if err != nil {
					fmt.Printf("Invalid %s:%s\n", k, v)
					return false
				}
				if h < 59 || h > 76 {
					fmt.Printf("Invalid %s:%s %d\n", k, v, h)
					return false
				}
				continue
			} else {
				fmt.Printf("Invalid %s:%s\n", k, v)
				return false
			}
		case "hcl":
			validHCL1 := regexp.MustCompile(`^\#([0-9A-Fa-f]{6})$`)
			match := validHCL1.MatchString(v)
			if !match {
				fmt.Printf("Invalid %s:%s\n", k, v)
				return false
			}
		case "ecl":
			valid := regexp.MustCompile(`^[a-z]{3}$`)
			match := valid.MatchString(v)
			if !match {
				fmt.Printf("Invalid %s:%s\n", k, v)
				return false
			}
			if !contains([]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}, v) {
				fmt.Printf("Invalid %s:%s\n", k, v)
				return false
			}
		case "pid":
			validPID := regexp.MustCompile(`^[0-9]{9}$`)
			match := validPID.MatchString(v)
			if !match {
				fmt.Printf("Invalid %s:%s\n", k, v)
				return false
			}
		default:
			continue
		}
	}
	return true
}
