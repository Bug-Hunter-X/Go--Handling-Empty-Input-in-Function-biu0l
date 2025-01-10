func processData(data []int) (int, error) {
    if len(data) == 0 {
        return 0, errors.New("data slice is empty")
    }

    sum := 0
    for _, v := range data {
        sum += v
    }

    return sum, nil
}

func main() {
    data := []int{1, 2, 3, 4, 5}
    sum, err := processData(data)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Sum:", sum)
    }

    emptyData := []int{}
    sum, err = processData(emptyData)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Sum:", sum)
    }
}