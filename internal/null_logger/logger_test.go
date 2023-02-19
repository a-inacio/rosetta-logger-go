package null_logger

import "testing"

func TestNullLogger_Fatal(t *testing.T) {
	nullLogger := NullLogger{}
	// No error should occur
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected error!")
		}
	}()
	nullLogger.Fatal("This should panic")
}

func TestNullLogger_Fatalf(t *testing.T) {
	nullLogger := NullLogger{}
	// No error should occur
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected error")
		}
	}()
	nullLogger.Fatalf("This should panic %s", "formatted")
}
