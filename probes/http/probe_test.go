package http

import (
	"testing"
)

func TestPositive(t *testing.T) {
	tests := [][]string{}
	generated := map[string]string{}
	tests = append(tests, test_online_cases_pos...)
	tests = append(tests, test_code_cases_pos...)
	tests = append(tests, test_length_cases_pos...)
	for _, v := range tests {
		glob := v[0]
		_, ok := generated[glob]
		if ok {
			generated[glob] += "," + v[1]
		} else {
			generated[glob] = v[1]
		}
	}
	for glob, opts := range generated {
		tests = append(tests, []string{glob, opts})
	}
	for _, test := range tests {
		result := Probe(test[0], test[1])
		if !result {
			t.Errorf("Probe('%s','%s') Failed", test[0], test[1])
		}
	}
}

func TestNegative(t *testing.T) {
	tests := [][]string{}
	generated := map[string]string{}
	tests = append(tests, test_online_cases_neg...)
	tests = append(tests, test_code_cases_neg...)
	tests = append(tests, test_length_cases_neg...)
	for _, v := range tests {
		glob := v[0]
		_, ok := generated[glob]
		if ok {
			generated[glob] += "," + v[1]
		} else {
			generated[glob] = v[1]
		}
	}
	for glob, opts := range generated {
		tests = append(tests, []string{glob, opts})
	}
	for _, test := range tests {
		result := Probe(test[0], test[1])
		if result {
			t.Errorf("Probe('%s','%s') Not failed", test[0], test[1])
		}
	}
}
