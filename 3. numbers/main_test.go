package main

import (
	"bytes"
	"testing"
)

func Test_printNum(t *testing.T) {
    bak := out
    defer func() { out = bak }()

	t.Run("Number Positive without separator", func(t *testing.T){
		out = new(bytes.Buffer)
		printNum("345")
		expected := "300\n40\n5\n"
		if out.(*bytes.Buffer).String() != expected {
			t.Errorf("printNum() = %v, want %v", out.(*bytes.Buffer).String(), expected)
		}	
	})

	t.Run("Number Positive with separator", func(t *testing.T){
		out = new(bytes.Buffer)
		printNum("1.459.440")
		expected := "1000000\n400000\n50000\n9000\n400\n40\n0\n"
		if out.(*bytes.Buffer).String() != expected {
			t.Errorf("printNum() = %v, want %v", out.(*bytes.Buffer).String(), expected)
		}	
	})

	t.Run("Number Positive 1 digit", func(t *testing.T){
		out = new(bytes.Buffer)
		printNum("1")
		expected := "1\n"
		if out.(*bytes.Buffer).String() != expected {
			t.Errorf("printNum() = %v, want %v", out.(*bytes.Buffer).String(), expected)
		}	
	})

	t.Run("Number Negative without separator", func(t *testing.T){
		out = new(bytes.Buffer)
		printNum("-345")
		expected := "-300\n-40\n-5\n"
		if out.(*bytes.Buffer).String() != expected {
			t.Errorf("printNum() = %v, want %v", out.(*bytes.Buffer).String(), expected)
		}	
	})

	t.Run("Number Negative with separator", func(t *testing.T){
		out = new(bytes.Buffer)
		printNum("-1.459.440")
		expected := "-1000000\n-400000\n-50000\n-9000\n-400\n-40\n-0\n"
		if out.(*bytes.Buffer).String() != expected{
			t.Errorf("printNum() = %v, want %v", out.(*bytes.Buffer).String(), expected)
		}	
	})

	t.Run("Number Negative 1 digit", func(t *testing.T){
		out = new(bytes.Buffer)
		printNum("-1")
		expected := "-1\n"
		if out.(*bytes.Buffer).String() != expected {
			t.Errorf("printNum() = %v, want %v", out.(*bytes.Buffer).String(), expected)
		}	
	})

}

