package decodingTree

import (
	"testing"
)

func TestBuildDecodingTree(t *testing.T) {
	tests := []struct {
		name    string
		ec      map[rune]string
		wantErr bool
	}{
		{
			name:    "valid codes",
			ec:      map[rune]string{'a': "0", 'b': "1"},
			wantErr: false,
		},
		{
			name:    "invalid code character",
			ec:      map[rune]string{'a': "2"},
			wantErr: true,
		},
		{
			name:    "conflicting codes",
			ec:      map[rune]string{'a': "0", 'b': "01"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := BuildDecodingTree(tt.ec)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildDecodingTree() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
