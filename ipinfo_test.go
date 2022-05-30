package main

import "testing"

func TestIsValidIp(t *testing.T) {
	var table = []Table{
		{"0.0.0.0", true},
		{"192.168.0.0", true},
		{"255.256.100.2", false},
		{"-100.200.182.160", false},
		{"300.200.182.160", false},
		{"100.2333.182.160", false},
		{"150.18.182,26", false},
	}

	for _, v := range table {
		result := IsValidIp(v.ip)

		if result != v.isValid {
			t.Errorf("Expection is %t, result is %t", v.isValid, result)
		}
	}
}

type Table struct {
	ip      string
	isValid bool
}
