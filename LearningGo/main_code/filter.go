package main_code

type Predicate func(int) bool

func isEven(n int) bool {
	return n%2 == 0
}

func isOdd(n int) bool {
	return n%2 != 0
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func GreaterThan(x int) Predicate {
	return func(n int) bool {
		return n > x
	}
}

func MultipleOf(x int) Predicate {
	return func(n int) bool {
		return n%x == 0
	}
}

func LessThan(x int) Predicate {
	return func(n int) bool {
		return n < x
	}
}

func All(predicates ...Predicate) Predicate {
	return func(n int) bool {
		for _, p := range predicates {
			if !p(n) {
				return false
			}
		}
		return true
	}
}

func Any(predicates ...Predicate) Predicate {
	return func(n int) bool {
		for _, p := range predicates {
			if p(n) {
				return true
			}
		}
		return false
	}
}

func filter(numbers []int, p Predicate) []int {
	var result []int
	for _, num := range numbers {
		if p(num) {
			result = append(result, num)
		}
	}
	return result
}