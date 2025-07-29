package main

import (
	"reflect"
	"testing"
)

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func TestMultiplesOfThree(t *testing.T) {
	if contains(fizzbuzz(2), "fizz") {
		t.Error("should not contain fizz")
	}
	if !contains(fizzbuzz(3), "fizz") {
		t.Error("should contain fizz")
	}
	expected := []string{"1", "2", "fizz", "4", "buzz", "fizz"}
	result := fizzbuzz(6)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v got %v", expected, result)
	}
}

func TestMultiplesOfFive(t *testing.T) {
	if contains(fizzbuzz(3), "buzz") {
		t.Error("should not contain buzz")
	}
	if !contains(fizzbuzz(5), "buzz") {
		t.Error("should contain buzz")
	}
	expected := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz"}
	result := fizzbuzz(10)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v got %v", expected, result)
	}
}

func TestMultiplesOfThreeAndFive(t *testing.T) {
	if contains(fizzbuzz(10), "fizzbuzz") {
		t.Error("should not contain fizzbuzz")
	}
	if !contains(fizzbuzz(15), "fizzbuzz") {
		t.Error("should contain fizzbuzz")
	}
	expected := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"}
	result := fizzbuzz(15)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v got %v", expected, result)
	}
}
