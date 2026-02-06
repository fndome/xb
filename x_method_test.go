package xb

import (
	"testing"
)

// Test structs with TableName methods
type TestUser struct {
	ID        int64  `db:"id,pk"`
	Name      string `db:"name"`
	UpdatedAt string `db:"updated_at"`
}

func (u *TestUser) TableName() string {
	return "t_users"
}

// TestUpdateBuilderX tests the X() method with and without parameters
func TestUpdateBuilderX(t *testing.T) {
	t.Run("X without parameters", func(t *testing.T) {
		sql, args := Of(&TestUser{}).
			Update(func(ub *UpdateBuilder) {
				ub.Set("name", "test").
					X("updated_at = CURRENT_TIMESTAMP")
			}).
			Eq("id", int64(1)).
			Build().
			SqlOfUpdate()

		expectedSQL := "UPDATE t_users SET name = ?, updated_at = CURRENT_TIMESTAMP  WHERE id = ?"
		if sql != expectedSQL {
			t.Errorf("SQL mismatch:\nExpected: %s\nGot:      %s", expectedSQL, sql)
		}

		expectedArgs := 2 // "test" and 1
		if len(args) != expectedArgs {
			t.Errorf("Args count mismatch: expected %d, got %d. Args: %v", expectedArgs, len(args), args)
		}

		if args[0] != "test" {
			t.Errorf("First arg should be 'test', got: %v", args[0])
		}
		if args[1] != int64(1) {
			t.Errorf("Second arg should be int64(1), got: %v (type: %T)", args[1], args[1])
		}
	})

	t.Run("X with parameters", func(t *testing.T) {
		sql, args := Of(&TestUser{}).
			Update(func(ub *UpdateBuilder) {
				ub.Set("name", "test").
					X("updated_at = ?", "2024-01-01 00:00:00")
			}).
			Eq("id", int64(1)).
			Build().
			SqlOfUpdate()

		expectedSQL := "UPDATE t_users SET name = ?, updated_at = ?  WHERE id = ?"
		if sql != expectedSQL {
			t.Errorf("SQL mismatch:\nExpected: %s\nGot:      %s", expectedSQL, sql)
		}

		expectedArgs := 3 // "test", "2024-01-01 00:00:00", and 1
		if len(args) != expectedArgs {
			t.Errorf("Args count mismatch: expected %d, got %d. Args: %v", expectedArgs, len(args), args)
		}
	})
}

// TestCondBuilderX tests the X() method in WHERE clauses
func TestCondBuilderX(t *testing.T) {
	t.Run("X in WHERE without parameters", func(t *testing.T) {
		sql, args, _ := Of(&TestUser{}).
			Eq("name", "test").
			X("created_at > DATE_SUB(NOW(), INTERVAL 7 DAY)").
			Build().
			SqlOfSelect()

		expectedSQL := "SELECT * FROM t_users WHERE name = ? AND created_at > DATE_SUB(NOW(), INTERVAL 7 DAY)"
		if sql != expectedSQL {
			t.Errorf("SQL mismatch:\nExpected: %s\nGot:      %s", expectedSQL, sql)
		}

		expectedArgs := 1 // only "test"
		if len(args) != expectedArgs {
			t.Errorf("Args count mismatch: expected %d, got %d. Args: %v", expectedArgs, len(args), args)
		}
	})

	t.Run("X in WHERE with parameters", func(t *testing.T) {
		sql, args, _ := Of(&TestUser{}).
			Eq("name", "test").
			X("created_at > ?", "2024-01-01").
			Build().
			SqlOfSelect()

		expectedSQL := "SELECT * FROM t_users WHERE name = ? AND created_at > ?"
		if sql != expectedSQL {
			t.Errorf("SQL mismatch:\nExpected: %s\nGot:      %s", expectedSQL, sql)
		}

		expectedArgs := 2 // "test" and "2024-01-01"
		if len(args) != expectedArgs {
			t.Errorf("Args count mismatch: expected %d, got %d. Args: %v", expectedArgs, len(args), args)
		}
	})
}

