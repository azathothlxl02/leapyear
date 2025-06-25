package main

import "testing"

func TestLeapYear_ReturnFalseForAnyYear(t *testing.T) {
    result := leapyear(2001)
    expected := false

    if result != expected {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}