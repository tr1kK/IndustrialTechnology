package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// sumOfDigits()
	// tempConverter()
	// doubleArrayValues()
	// mergeStrings()
	// distanceBetweenTwoDots()
	// evenOddCheck()
	// checkingTheHighYear()
	// maxOfThreeNumbers()
	// ageCategory()
	// checkingDiv()
	// factorial()
	// fibonacciNumbers()
	// arrayReverse()
	// primeNumbers()
	sumOfArray()
}

// Сумма цифр числа
func sumOfDigits() int {
	var n int
	fmt.Print("Введите число: \n")
	fmt.Scan(&n)
	res := recursionSumOfDigits(n)
	fmt.Println("Сумма числа", n, " = ", recursionSumOfDigits(n))
	return res
}

func recursionSumOfDigits(num int) int {
	if num == 0 {
		return 0
	}
	return num%10 + recursionSumOfDigits(num/10)
}

// Преобразование температуры
func tempConverter() float64 {
	var value string
	var degrees, res float64
	fmt.Scan(&degrees, &value)

	if value == "(Celsius)" {
		res = (degrees * 9 / 5) + 32
		fmt.Print(res, " (Fahrenheit)")
	} else {
		res = (32 - degrees) * 5 / 9
		fmt.Print(res, " (Celsius)")
	}
	return res
}

// Удвоение каждого элемента массива

func doubleArrayValues() []float64 {
	var numbers []float64

	values, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	values = strings.TrimSpace(values)
	str_nums := strings.Split(values, " ")
	for _, nums := range str_nums {
		num, err := strconv.ParseFloat(nums, 64)
		if err != nil {
			continue
		}
		numbers = append(numbers, num*2)
	}
	fmt.Print(numbers)
	return numbers
}

// Объединение строк

func mergeStrings() string {
	var res_text string = ""

	values, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	values = strings.TrimSpace(values)
	string_values := strings.Split(values, ",")
	for _, text := range string_values {
		res_text = res_text + text
	}
	fmt.Print(res_text)
	return res_text
}

// Расчет расстояния между двумя точками

func distanceBetweenTwoDots() (res float64) {
	var nums []float64

	values, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	values = strings.TrimSpace(values)
	values = strings.ReplaceAll(values, "(", "")
	values = strings.ReplaceAll(values, ")", "")
	values = strings.ReplaceAll(values, "x1", "")
	values = strings.ReplaceAll(values, "x2", "")
	values = strings.ReplaceAll(values, "y1", "")
	values = strings.ReplaceAll(values, "y2", "")
	values = strings.ReplaceAll(values, "=", "")
	string_values := strings.Split(values, ", ")
	for _, str_nums := range string_values {
		num, err := strconv.ParseFloat(str_nums, 64)
		if err != nil {
			continue
		}
		nums = append(nums, num)
	}

	res = math.Pow(math.Pow(nums[2]-nums[0], 2)+math.Pow(nums[3]-nums[1], 2), 0.5)
	fmt.Print(res)
	return
}

// Проверка на четность/нечетность

func evenOddCheck() (res string) {
	var n int
	fmt.Scan(&n)
	if n%2 == 0 {
		fmt.Print("Чётное")
		res = "Чётное"
	} else {
		fmt.Print("Нечётное")
		res = "Нечётное"
	}
	return
}

// Проверка высокосного года

func checkingTheHighYear() (res string) {
	var n int
	fmt.Scan(&n)
	if n%4 != 0 || (n%100 == 0 && n%400 != 0) {
		fmt.Print("Невисокосный")
		res = "Невисокосный"
	} else {
		fmt.Print("Високосный")
		res = "Високосный"
	}
	return
}

// Определение наибольшего из трех чисел

func maxOfThreeNumbers() (res float64) {
	var n1, n2, n3 float64
	fmt.Scan(&n1, &n2, &n3)
	if n1 > n2 && n1 > n3 {
		res = n1
		fmt.Print(n1)
	} else if n2 > n1 && n2 > n3 {
		res = n2
		fmt.Print(n2)
	} else {
		res = n3
		fmt.Print(n3)
	}
	return
}

/*
	Категория возраста

Ребенок: 0-40
Подросток: 41-60
Взрослый: 61-80
Пожилой: 81-inf
*/
func ageCategory() (res string) {
	var n int
	fmt.Scan(&n)
	if n < 41 {
		res = "Ребенок"
		fmt.Print("Ребенок")
	} else if n < 61 {
		res = "Подросток"
		fmt.Print("Подросток")
	} else if n < 81 {
		res = "Взрослый"
		fmt.Print("Взрослый")
	} else {
		res = "Пожилой"
		fmt.Print("Пожилой")
	}
	return
}

// Проверка делимости на 3 и 5

func checkingDiv() (res string) {
	var n int
	fmt.Scan(&n)
	if n%3 == 0 && n%5 == 0 {
		res = "Делится"
		fmt.Print("Делится")
	} else {
		res = "Не делится"
		fmt.Print("Не делится")
	}
	return
}

// Факториал числа
func factorial() (res int) {
	var n int
	fmt.Scan(&n)
	res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	fmt.Print(res)
	return
}

// Числа Фибоначчи
func fibonacciNumbers() (res []int) {
	var n int
	fmt.Scan(&n)
	if n > 0 {
		res = append(res, 0)
		if n == 1 {
			fmt.Print(res)
		} else {
			res = append(res, 1)
			for i := 2; i < n; i++ {
				res = append(res, res[len(res)-1]+res[len(res)-2])
			}
			fmt.Print(res)
		}
	}
	return
}

// Реверс массива
func arrayReverse() (res []float64) {
	values, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	values = strings.TrimSpace(values)
	str_nums := strings.Split(values, ", ")
	for _, nums := range str_nums {
		num, err := strconv.ParseFloat(nums, 64)
		if err != nil {
			continue
		}
		res = append(res, num)
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	fmt.Print(res)
	return
}

// Поиск простых чисел

func primeNumbers() []int {
	var n int
	fmt.Scan(&n)
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if isPrime[p] {
			for i := p * p; i <= n; i += p {
				isPrime[i] = false
			}
		}
	}

	var primes []int
	for p := 2; p <= n; p++ {
		if isPrime[p] {
			primes = append(primes, p)
		}
	}
	fmt.Print(primes)
	return primes
}

// Сумма чисел в массиве
func sumOfArray() float64 {
	var res []float64
	var sum float64
	values, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	values = strings.TrimSpace(values)
	str_nums := strings.Split(values, ", ")
	for _, nums := range str_nums {
		num, err := strconv.ParseFloat(nums, 64)
		if err != nil {
			continue
		}
		res = append(res, num)
	}
	for i := 0; i < len(res); i++ {
		sum += res[i]
	}
	fmt.Print(sum)
	return sum
}
