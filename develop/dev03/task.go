package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	flagK = flag.Int("k", 0, "указание колонки для сортировки")
	flagN = flag.Bool("n", false, "сортировать по числовому значению")
	flagR = flag.Bool("r", false, "сортировать в обратном порядке")
	flagU = flag.Bool("u", false, "не выводить повторяющиеся строки")
	flagC = flag.Bool("c", false, "проверять отсортированы ли данные")
	path  string
)

func main() {
	flag.Parse()
	raws, err := ScanFile(path)
	if err != nil {
		log.Fatal(err)
	}
	SortingStrings(raws)
	if err := PrintSlice(raws); err != nil {
		log.Fatal(err)
	}
}

func ScanFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	data, repetitions := make([]string, 0), make(map[string]struct{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := columnOfText(scanner.Text())
		if *flagU {
			if _, ok := repetitions[text]; ok {
				repetitions[text] = struct{}{}
			} else {
				continue
			}
		}
		data = append(data, text)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

func SortingStrings(s []string) {
	switch {
	case *flagC:
		var min string
		for _, v := range s {
			if min < v {
				min = v
			}
		}
		fmt.Print(fmt.Sprintf("%s", min))
		os.Exit(1)
	case *flagN:
		sort.Slice(s, func(i, j int) bool {
			vI, _ := strconv.Atoi(s[i])
			vJ, _ := strconv.Atoi(s[j])
			return vI < vJ
		})
	default:
		sort.Strings(s)
	}

}

func PrintSlice(s []string) error {
	out := bufio.NewWriter(os.Stdout)
	if *flagR {
		for _, v := range s {
			if _, err := out.WriteString(v); err != nil {
				return err
			}
		}
	} else {
		for i := len(s) - 1; i >= 0; i-- {
			if _, err := out.WriteString(s[i]); err != nil {
				return err
			}
		}
	}

	if err := out.Flush(); err != nil {
		return err
	}
	return nil
}

func columnOfText(s string) string {
	if *flagK == 0 {
		return s
	}
	columns := strings.Split(s, " ")
	if len(columns) <= *flagK {
		return ""
	}
	return columns[*flagK-1]
}
