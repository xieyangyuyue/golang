package main

import "testing"

// 测试正常加法
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

// 测试负数加法
func TestAddNegative(t *testing.T) {
	result := Add(-1, -2)
	expected := -3
	if result != expected {
		t.Fatalf("Add(-1, -2) = %d; want %d", result, expected) // 使用 Fatalf 会立即终止测试
	}
}
