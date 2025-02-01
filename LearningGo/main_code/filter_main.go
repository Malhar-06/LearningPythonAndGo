package main_code

import (
	"reflect"
	"testing"
)

func TestStory1(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}
	result := filter(input, isEven)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestStory1 failed. Expected %v, got %v", expected, result)
	}
}

func TestStory2(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{1, 3, 5, 7, 9}
	result := filter(input, isOdd)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestStory2 failed. Expected %v, got %v", expected, result)
	}
}

func TestStory3(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 3, 5, 7}
	result := filter(input, isPrime)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestStory3 failed. Expected %v, got %v", expected, result)
	}
}

func TestStory4(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{3, 5, 7}
	predicate := All(isOdd, isPrime)
	result := filter(input, predicate)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestStory4 failed. Expected %v, got %v", expected, result)
	}
}

func TestStory5(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expected := []int{10, 20}
	predicate := All(isEven, MultipleOf(5))
	result := filter(input, predicate)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestStory5 failed. Expected %v, got %v", expected, result)
	}
}

func TestStory6(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expected := []int{15}
	predicate := All(isOdd, MultipleOf(3), GreaterThan(10))
	result := filter(input, predicate)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestStory6 failed. Expected %v, got %v", expected, result)
	}
}

func TestStory7_Case1(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expected := []int{9, 15}
	predicate := All(isOdd, GreaterThan(5), MultipleOf(3))
	result := filter(input, predicate)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestStory7_Case1 failed. Expected %v, got %v", expected, result)
	}
}

func TestStory7_Case2(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expected := []int{6, 12}
	predicate := All(isEven, LessThan(15), MultipleOf(3))
	result := filter(input, predicate)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestStory7_Case2 failed. Expected %v, got %v", expected, result)
	}
}

func TestStory8_Case1(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expected := []int{2, 3, 5, 7, 10, 11, 13, 15, 16, 17, 18, 19, 20}
	predicate := Any(isPrime, GreaterThan(15), MultipleOf(5))
	result := filter(input, predicate)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestStory8_Case1 failed. Expected %v, got %v", expected, result)
	}
}

func TestStory8_Case2(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expected := []int{1, 2, 3, 4, 5, 6, 9, 12, 15, 18}
	predicate := Any(LessThan(6), MultipleOf(3))
	result := filter(input, predicate)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestStory8_Case2 failed. Expected %v, got %v", expected, result)
	}
}