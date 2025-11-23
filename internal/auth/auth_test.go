package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name:    "missing Authorization header",
			headers: http.Header{},
			want:    "",
			wantErr: true,
		},
		{
			name:    "malformed Authorization header format",
			headers: http.Header{"Authorization": {"Bearer sometoken"}},
			want:    "",
			wantErr: true,
		},
		{
			name:    "ApiKey with no value",
			headers: http.Header{"Authorization": {"ApiKey"}},
			want:    "",
			wantErr: true,
		},
		{
			name:    "ApiKey with value",
			headers: http.Header{"Authorization": {"ApiKey correct-key"}},
			want:    "correct-key",
			wantErr: false,
		},
		{
			name:    "ApiKey with multiple spaces",
			headers: http.Header{"Authorization": {"ApiKey correct-key extra-ignored"}},
			want:    "correct-key",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
