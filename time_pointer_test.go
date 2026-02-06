package xb

import (
	"testing"
	"time"
)

// 测试 InsertBuilder 处理 *time.Time
func TestInsertBuilderTimePointer(t *testing.T) {
	// 测试1: nil 指针
	t.Run("nil pointer", func(t *testing.T) {
		builder := &InsertBuilder{}
		var nilTime *time.Time = nil
		builder.Set("last_active_at", nilTime)
		
		if len(builder.bbs) != 0 {
			t.Errorf("Expected 0 fields, got %d", len(builder.bbs))
		}
		t.Log("✓ nil *time.Time 被正确忽略")
	})
	
	// 测试2: 非 nil 指针
	t.Run("non-nil pointer", func(t *testing.T) {
		builder := &InsertBuilder{}
		now := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
		builder.Set("last_active_at", &now)
		
		if len(builder.bbs) != 1 {
			t.Fatalf("Expected 1 field, got %d", len(builder.bbs))
		}
		
		if builder.bbs[0].Key != "last_active_at" {
			t.Errorf("Expected key 'last_active_at', got '%s'", builder.bbs[0].Key)
		}
		
		expected := "2024-01-01 12:00:00"
		if builder.bbs[0].Value != expected {
			t.Errorf("Expected value '%s', got '%v'", expected, builder.bbs[0].Value)
		}
		t.Logf("✓ 非 nil *time.Time 被正确格式化: %v", builder.bbs[0].Value)
	})
	
	// 测试3: time.Time 值类型
	t.Run("value type", func(t *testing.T) {
		builder := &InsertBuilder{}
		now := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
		builder.Set("created_at", now)
		
		if len(builder.bbs) != 1 {
			t.Fatalf("Expected 1 field, got %d", len(builder.bbs))
		}
		
		expected := "2024-01-01 12:00:00"
		if builder.bbs[0].Value != expected {
			t.Errorf("Expected value '%s', got '%v'", expected, builder.bbs[0].Value)
		}
		t.Logf("✓ time.Time 值类型被正确格式化: %v", builder.bbs[0].Value)
	})
	
	// 测试4: 零值 time.Time
	t.Run("zero value", func(t *testing.T) {
		builder := &InsertBuilder{}
		var zeroTime time.Time
		builder.Set("created_at", zeroTime)
		
		if len(builder.bbs) != 1 {
			t.Fatalf("Expected 1 field, got %d", len(builder.bbs))
		}
		
		expected := "0001-01-01 00:00:00"
		if builder.bbs[0].Value != expected {
			t.Errorf("Expected value '%s', got '%v'", expected, builder.bbs[0].Value)
		}
		t.Logf("✓ 零值 time.Time 被格式化为: %v", builder.bbs[0].Value)
	})
}

// 测试 UpdateBuilder 处理 *time.Time
func TestUpdateBuilderTimePointer(t *testing.T) {
	// 测试1: nil 指针
	t.Run("nil pointer", func(t *testing.T) {
		builder := &UpdateBuilder{}
		var nilTime *time.Time = nil
		builder.Set("last_active_at", nilTime)
		
		if len(builder.bbs) != 0 {
			t.Errorf("Expected 0 fields, got %d", len(builder.bbs))
		}
		t.Log("✓ nil *time.Time 被正确忽略")
	})
	
	// 测试2: 非 nil 指针
	t.Run("non-nil pointer", func(t *testing.T) {
		builder := &UpdateBuilder{}
		now := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
		builder.Set("last_active_at", &now)
		
		if len(builder.bbs) != 1 {
			t.Fatalf("Expected 1 field, got %d", len(builder.bbs))
		}
		
		expected := "2024-01-01 12:00:00"
		if builder.bbs[0].Value != expected {
			t.Errorf("Expected value '%s', got '%v'", expected, builder.bbs[0].Value)
		}
		t.Logf("✓ 非 nil *time.Time 被正确格式化: %v", builder.bbs[0].Value)
	})
}
