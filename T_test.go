m
func TestAbs(t *testing.T) {
	tests := []struct {
		input  float64
		output float64
	}{
		{-1, 1},
		{0, 0},
		{1, 1},
	}

	for _, tt := range tests {
		got := Abs(tt.input)
		if got != tt.output {
			t.Errorf("Abs(%v) = %v; want %v", tt.input, got, tt.output)
		}
	}
}
