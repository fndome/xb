// Copyright 2025 me.fndo.xb
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package xb

import (
	"testing"
	"time"
)

// TestInsertBuilderNilTimePointer tests that nil *time.Time doesn't cause panic
func TestInsertBuilderNilTimePointer(t *testing.T) {
	ib := &InsertBuilder{}
	
	var nilTime *time.Time = nil
	
	// This should not panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("InsertBuilder.Set panicked with nil *time.Time: %v", r)
		}
	}()
	
	ib.Set("created_at", nilTime)
	
	// Should have no items since nil is skipped
	if len(ib.bbs) != 0 {
		t.Errorf("Expected 0 items, got %d", len(ib.bbs))
	}
	
	t.Log("✓ InsertBuilder correctly handles nil *time.Time")
}

// TestUpdateBuilderNilTimePointer tests that nil *time.Time doesn't cause panic
func TestUpdateBuilderNilTimePointer(t *testing.T) {
	ub := &UpdateBuilder{}
	
	var nilTime *time.Time = nil
	
	// This should not panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("UpdateBuilder.Set panicked with nil *time.Time: %v", r)
		}
	}()
	
	ub.Set("updated_at", nilTime)
	
	// Should have no items since nil is skipped
	if len(ub.bbs) != 0 {
		t.Errorf("Expected 0 items, got %d", len(ub.bbs))
	}
	
	t.Log("✓ UpdateBuilder correctly handles nil *time.Time")
}

// TestInsertBuilderValidTimePointer tests that valid *time.Time works correctly
func TestInsertBuilderValidTimePointer(t *testing.T) {
	ib := &InsertBuilder{}
	
	now := time.Now()
	
	ib.Set("created_at", &now)
	
	// Should have 1 item
	if len(ib.bbs) != 1 {
		t.Errorf("Expected 1 item, got %d", len(ib.bbs))
	}
	
	// Value should be formatted string
	if ib.bbs[0].Key != "created_at" {
		t.Errorf("Expected key 'created_at', got '%s'", ib.bbs[0].Key)
	}
	
	if _, ok := ib.bbs[0].Value.(string); !ok {
		t.Errorf("Expected string value, got %T", ib.bbs[0].Value)
	}
	
	t.Logf("✓ InsertBuilder correctly handles valid *time.Time: %v", ib.bbs[0].Value)
}

// TestInsertBuilderTimeValue tests that time.Time value type works correctly
func TestInsertBuilderTimeValue(t *testing.T) {
	ib := &InsertBuilder{}
	
	now := time.Now()
	
	ib.Set("created_at", now)
	
	// Should have 1 item
	if len(ib.bbs) != 1 {
		t.Errorf("Expected 1 item, got %d", len(ib.bbs))
	}
	
	// Value should be formatted string
	if _, ok := ib.bbs[0].Value.(string); !ok {
		t.Errorf("Expected string value, got %T", ib.bbs[0].Value)
	}
	
	t.Logf("✓ InsertBuilder correctly handles time.Time value: %v", ib.bbs[0].Value)
}

// TestInsertBuilderNilValue tests that nil interface{} is handled correctly
func TestInsertBuilderNilValue(t *testing.T) {
	ib := &InsertBuilder{}
	
	ib.Set("nullable_field", nil)
	
	// Should have no items since nil is skipped
	if len(ib.bbs) != 0 {
		t.Errorf("Expected 0 items, got %d", len(ib.bbs))
	}
	
	t.Log("✓ InsertBuilder correctly handles nil value")
}

// TestUpdateBuilderNilValue tests that nil interface{} is handled correctly
func TestUpdateBuilderNilValue(t *testing.T) {
	ub := &UpdateBuilder{}
	
	ub.Set("nullable_field", nil)
	
	// Should have no items since nil is skipped
	if len(ub.bbs) != 0 {
		t.Errorf("Expected 0 items, got %d", len(ub.bbs))
	}
	
	t.Log("✓ UpdateBuilder correctly handles nil value")
}
