package pyament

import (
	"testing"

	"github.com/google/uuid"
)

func TestParsePaymentID(t *testing.T) {
	// Define test cases
	tests := []struct {
		name    string
		idStr   string
		want    ID
		wantErr bool
	}{
		{
			name:    "valid UUID",
			idStr:   "01890a5a-f410-7c66-a7dc-7723a5ff72bb",
			want:    ID(uuid.MustParse("01890a5a-f410-7c66-a7dc-7723a5ff72bb")),
			wantErr: false,
		},
		{
			name:    "invalid UUID format",
			idStr:   "not-a-uuid",
			wantErr: true,
		},
		{
			name:    "empty string",
			idStr:   "",
			wantErr: true,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePaymentID(tt.idStr)

			// Check error expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePaymentID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// If no error expected, check the returned ID
			if !tt.wantErr {
				if got != tt.want {
					t.Errorf("ParsePaymentID() = %v, want %v", got, tt.want)
				}

				// Also check that the string representation matches
				if got.String() != tt.idStr {
					t.Errorf("ParsePaymentID().String() = %q, want %q", got.String(), tt.idStr)
				}
			}
		})
	}
}
