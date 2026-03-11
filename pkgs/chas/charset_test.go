package chas

import (
	"bytes"
	"testing"
)

func TestSearchString(t *testing.T) {
	tests := []struct {
		name    string
		charset string
		text    string
		want    string
		wantErr bool
	}{
		{
			name:    "Exact match",
			charset: "abc",
			text:    "abc\n",
			want:    "abc\n",
			wantErr: false,
		},
		{
			name:    "Multiple lines with one match",
			charset: "gold",
			text:    "silver\ngolden\nbronze\n",
			want:    "golden\n",
			wantErr: false,
		},
		{
			name:    "Missing one character",
			charset: "abc",
			text:    "ab\n",
			want:    "",
			wantErr: false,
		},
		{
			name:    "Unicode/Emoji support",
			charset: "🚀⭐",
			text:    "to the 🚀 moon ⭐\njust a 🚀\n",
			want:    "to the 🚀 moon ⭐\n",
			wantErr: false,
		},
		{
			name:    "Charset with duplicates",
			charset: "aaa",
			text:    "apple\n",
			want:    "apple\n", // 'a' exists in apple, fulfilling the set
			wantErr: false,
		},
		{
			name:    "Empty input",
			charset: "abc",
			text:    "",
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := SearchString(tt.charset, tt.text, &buf)

			if (err != nil) != tt.wantErr {
				t.Errorf("SearchString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got := buf.String(); got != tt.want {
				t.Errorf("SearchString() output = %q, want %q", got, tt.want)
			}
		})
	}
}
