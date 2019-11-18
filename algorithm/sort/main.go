package main

import (
	"DataStructuresAndAlgorithm/sort/BubbleSort"
	"DataStructuresAndAlgorithm/sort/CombSort"
	"DataStructuresAndAlgorithm/sort/HeapSort"
	"DataStructuresAndAlgorithm/sort/InsertSort"
	"DataStructuresAndAlgorithm/sort/MergeSort"
	"DataStructuresAndAlgorithm/sort/PancakeSort"
	"DataStructuresAndAlgorithm/sort/QuickSort"
	"DataStructuresAndAlgorithm/sort/SelectSort"
	"DataStructuresAndAlgorithm/sort/ShellSort"
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

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
	values, err := readValues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
		case "bubblesort":
			BubbleSort.BubbleSort(values)
		case "combsort":
			CombSort.CombSort(values)
		case "heapsort":
			HeapSort.HeapSort(values)
		case "insertsort":
			InsertSort.InsertSort(values)
		case "mergesort":
			MergeSort.MergeSort(values, 0, len(values)-1)
		case "pancakesort":
			PancakeSort.PancakeSort(values)
		case "quicksort":
			QuickSort.QuickSort(values)
		case "selectsort":
			SelectSort.SelectSort(values)
		case "shellsort":
			ShellSort.ShellSort(values)
		default:
			log.Println("Sort Algorithm", *algorithm, "is either unknown or unsupported.")
		}
		t2 := time.Now()
		log.Println("The Sorting process costs", t2.Sub(t1), "to completed.")
		writeValues(values, *outfile)
	} else {
		log.Println(err)
	}
}
