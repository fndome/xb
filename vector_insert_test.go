package xb

import (
	"strings"
	"testing"
)

type TestVectorTable struct {
	ID        int64  `db:"id,pk"`
	Embedding Vector `db:"embedding"`
}

func (tt *TestVectorTable) TableName() string {
	return "test_vectors"
}

// TestVectorInsert_NoJSONSerialization 验证 Vector 类型不会被 JSON 序列化
func TestVectorInsert_NoJSONSerialization(t *testing.T) {
	vec := Vector{0.1, 0.2, 0.3}
	
	sql, args := Of(&TestVectorTable{}).
		Insert(func(ib *InsertBuilder) {
			ib.Set("id", int64(1)).
				Set("embedding", vec)
		}).
		Build().
		SqlOfInsert()
	
	t.Logf("SQL: %s", sql)
	t.Logf("Args: %v", args)
	t.Logf("Args[1] type: %T", args[1])
	
	// 验证 args[1] 是 Vector 类型，而不是 string
	switch args[1].(type) {
	case Vector:
		t.Logf("✅ Vector 类型保持不变，driver.Valuer 会被调用")
	case string:
		// 如果是 string，检查是否被 JSON 序列化了
		str := args[1].(string)
		if strings.HasPrefix(str, "[") && strings.HasSuffix(str, "]") {
			t.Errorf("❌ Vector 被错误地 JSON 序列化为字符串: %s", str)
		} else {
			t.Errorf("❌ Vector 被转换为字符串: %s", str)
		}
	default:
		t.Errorf("❌ Vector 被转换为未知类型: %T", args[1])
	}
}

// TestVectorUpdate_NoJSONSerialization 验证 Update 中 Vector 类型不会被 JSON 序列化
func TestVectorUpdate_NoJSONSerialization(t *testing.T) {
	vec := Vector{0.1, 0.2, 0.3}
	
	sql, args := Of(&TestVectorTable{}).
		Update(func(ub *UpdateBuilder) {
			ub.Set("embedding", vec)
		}).
		Eq("id", int64(1)).
		Build().
		SqlOfUpdate()
	
	t.Logf("SQL: %s", sql)
	t.Logf("Args: %v", args)
	t.Logf("Args[0] type: %T", args[0])
	
	// 验证 args[0] 是 Vector 类型，而不是 string
	switch args[0].(type) {
	case Vector:
		t.Logf("✅ Vector 类型保持不变，driver.Valuer 会被调用")
	case string:
		str := args[0].(string)
		if strings.HasPrefix(str, "[") && strings.HasSuffix(str, "]") {
			t.Errorf("❌ Vector 被错误地 JSON 序列化为字符串: %s", str)
		} else {
			t.Errorf("❌ Vector 被转换为字符串: %s", str)
		}
	default:
		t.Errorf("❌ Vector 被转换为未知类型: %T", args[0])
	}
}
