package main

func main() {
	var input []int = []int{100000, 5, 7, 9, 44, 33, 4, 6, 882222226, 453, 244, 234}
	QuickSort(input)
	fmt.Println(input)
}

func QuickSort(input []int) {
	sort(0, len(input)-1, input)
}

func sort(start, end int, input []int) {
	if start >= end {
		return
	}
	flag := input[start]
	i := start
	j := end
	for i < j {
		for input[j] >= flag && i < j {
			j--
		}
		for input[i] <= flag && i < j {
			i++
		}
		if i < j {
			temp := input[i]
			input[i] = input[j]
			input[j] = temp
		}
	}
	input[start] = input[i]
	input[i] = flag
	sort(start, i-1, input)
	sort(i+1, end, input)
}
