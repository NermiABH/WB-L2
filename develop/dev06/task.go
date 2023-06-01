package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	flagFields    = flag.Int("f", 0, "Выбрать поля(колонки)")
	flagDelimiter = flag.String("d", "	", "Использовать другой разделитель")
	flagSeparated = flag.Bool("s", false, "только строки с разделителем")
)

func init() {
	flag.Parse()
}

func main() {
	if *flagFields <= 0 {
		log.Fatal("flag <= 0")
	}
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		splitedS := strings.Split(in.Text(), *flagDelimiter)
		if *flagSeparated && len(splitedS) < 2 {
			continue
		}
		if len(splitedS) > *flagFields {
			fmt.Println(splitedS[*flagFields-1])
		}
	}
}
