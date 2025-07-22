package mode

import (
	"testing"
)

func TestMode(t *testing.T) {
	t.Run("prefix parsing", func(t *testing.T) {
		t.Run("should parse /quick prefix", func(t *testing.T) {
			// Arrange
			input := "/quick buy milk"

			// Act
			mode, task, hasPrefix := ParseModePrefix(input)

			// Assert
			if mode != QuickMode {
				t.Errorf("expected mode %v, got %v", QuickMode, mode)
			}
			if task != "buy milk" {
				t.Errorf("expected task %q, got %q", "buy milk", task)
			}
			if !hasPrefix {
				t.Errorf("expected hasPrefix to be true, got false")
			}
		})

		t.Run("should parse /q shortcut", func(t *testing.T) {
			// Arrange
			input := "/q buy milk"

			// Act
			mode, task, hasPrefix := ParseModePrefix(input)

			// Assert
			if mode != QuickMode {
				t.Errorf("expected mode %v, got %v", QuickMode, mode)
			}
			if task != "buy milk" {
				t.Errorf("expected task %q, got %q", "buy milk", task)
			}
			if !hasPrefix {
				t.Errorf("expected hasPrefix to be true, got false")
			}
		})

		t.Run("should parse /talk prefix", func(t *testing.T) {
			// Arrange
			input := "/talk prepare for meeting"

			// Act
			mode, task, hasPrefix := ParseModePrefix(input)

			// Assert
			if mode != TalkMode {
				t.Errorf("expected mode %v, got %v", TalkMode, mode)
			}
			if task != "prepare for meeting" {
				t.Errorf("expected task %q, got %q", "prepare for meeting", task)
			}
			if !hasPrefix {
				t.Errorf("expected hasPrefix to be true, got false")
			}
		})

		t.Run("should parse /t shortcut", func(t *testing.T) {
			// Arrange
			input := "/t prepare for meeting"

			// Act
			mode, task, hasPrefix := ParseModePrefix(input)

			// Assert
			if mode != TalkMode {
				t.Errorf("expected mode %v, got %v", TalkMode, mode)
			}
			if task != "prepare for meeting" {
				t.Errorf("expected task %q, got %q", "prepare for meeting", task)
			}
			if !hasPrefix {
				t.Errorf("expected hasPrefix to be true, got false")
			}
		})

		t.Run("should parse /plan prefix", func(t *testing.T) {
			// Arrange
			input := "/plan organize conference"

			// Act
			mode, task, hasPrefix := ParseModePrefix(input)

			// Assert
			if mode != PlanningMode {
				t.Errorf("expected mode %v, got %v", PlanningMode, mode)
			}
			if task != "organize conference" {
				t.Errorf("expected task %q, got %q", "organize conference", task)
			}
			if !hasPrefix {
				t.Errorf("expected hasPrefix to be true, got false")
			}
		})

		t.Run("should parse /p shortcut", func(t *testing.T) {
			// Arrange
			input := "/p organize conference"

			// Act
			mode, task, hasPrefix := ParseModePrefix(input)

			// Assert
			if mode != PlanningMode {
				t.Errorf("expected mode %v, got %v", PlanningMode, mode)
			}
			if task != "organize conference" {
				t.Errorf("expected task %q, got %q", "organize conference", task)
			}
			if !hasPrefix {
				t.Errorf("expected hasPrefix to be true, got false")
			}
		})

		t.Run("should return no prefix for regular input", func(t *testing.T) {
			// Arrange
			input := "buy milk"

			// Act
			mode, task, hasPrefix := ParseModePrefix(input)

			// Assert
			if mode != QuickMode {
				t.Errorf("expected mode %v, got %v", QuickMode, mode)
			}
			if task != "buy milk" {
				t.Errorf("expected task %q, got %q", "buy milk", task)
			}
			if hasPrefix {
				t.Errorf("expected hasPrefix to be false, got true")
			}
		})
	})
}
