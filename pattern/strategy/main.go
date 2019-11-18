package main

import (
	"github.com/Hexiaopi/go-patterns-example/strategy/bubblesort"
	"github.com/Hexiaopi/go-patterns-example/strategy/combsort"
	"github.com/Hexiaopi/go-patterns-example/strategy/heapsort"
	"github.com/Hexiaopi/go-patterns-example/strategy/insertsort"
	"github.com/Hexiaopi/go-patterns-example/strategy/mergesort"
	"github.com/Hexiaopi/go-patterns-example/strategy/pancakesort"
	"github.com/Hexiaopi/go-patterns-example/strategy/quicksort"
	"github.com/Hexiaopi/go-patterns-example/strategy/selectsort"
	"github.com/Hexiaopi/go-patterns-example/strategy/shellsort"

	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type SortAlgorithm interface {
	Sort(values []int)
}

var (
	infile    *string = flag.String("i", "unsorted.dat", "File contanins values for Sorting")
	outfile   *string = flag.String("o", "sorted.dat", "File to received sorted values")
	algorithm *string = flag.String("a", "bubblesort", "Sort Algorithm")
)

func readValues(infile string) (values []int, err error) {
	values = make([]int, 0)
	file, err := os.Open(infile)
	if err != nil {
		log.Println("Failed to open the input file ", infile)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)

	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			log.Println("A too long line,seems unexpected.")
			return
		}
		value, err1 := strconv.Atoi(string(line))
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		log.Println("Failed to create the output file ", outfile)
		return err
	}
	defer file.Close()
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func main() {
	flag.Parse()

	if infile != nil {
		log.Println("\t infile:", *infile, "\t outfile:", *outfile, "\t sort algorithm:", *algorithm)
	}
	var sortAlgorithm SortAlgorithm
	values, err := readValues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
		case "bubblesort":
			sortAlgorithm = bubblesort.BubbleSort{}
		case "combsort":
			sortAlgorithm = combsort.CombSort{}
		case "heapsort":
			sortAlgorithm = heapsort.HeapSort{}
		case "insertsort":
			sortAlgorithm = insertsort.InsertSort{}
		case "mergesort":
			sortAlgorithm = mergesort.MergeSort{Low: 0, High: len(values) - 1}
		case "pancakesort":
			sortAlgorithm = pancakesort.PancakeSort{}
		case "quicksort":
			sortAlgorithm = quicksort.QuickSort{}
		case "selectsort":
			sortAlgorithm = selectsort.SelectSort{}
		case "shellsort":
			sortAlgorithm = shellsort.ShellSort{}
		default:
			log.Println("Sort Algorithm", *algorithm, "is either unknown or unsupported.")
		}
		sortAlgorithm.Sort(values)
		t2 := time.Now()
		log.Println("The Sorting process costs", t2.Sub(t1), "to completed.")
		writeValues(values, *outfile)
	} else {
		log.Println(err)
	}
}
