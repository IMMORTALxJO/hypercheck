package fs

import (
	"testing"
)

func TestPositive(t *testing.T) {
	tests := [][]string{}
	generated := map[string]string{}
	tests = append(tests, test_totalsize_cases_pos...)
	tests = append(tests, test_exists_cases_pos...)
	tests = append(tests, test_count_cases_pos...)
	tests = append(tests, test_regular_cases_pos...)
	tests = append(tests, test_size_cases_pos...)
	tests = append(tests, test_dir_cases_pos...)
	tests = append(tests, test_owner_cases_pos...)
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
	tests := [][]string{
		[]string{"./assets/file_1Kb", "size>10Mb,size==1Kb"},
	}
	generated := map[string]string{}
	tests = append(tests, test_totalsize_cases_neg...)
	tests = append(tests, test_exists_cases_neg...)
	tests = append(tests, test_count_cases_neg...)
	tests = append(tests, test_regular_cases_neg...)
	tests = append(tests, test_size_cases_neg...)
	tests = append(tests, test_dir_cases_neg...)
	tests = append(tests, test_owner_cases_neg...)
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
