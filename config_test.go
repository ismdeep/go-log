package log

import (
	"reflect"
	"testing"
)

func TestParseConfig(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "",
			args: args{
				dsn: "console://[stdout,file]?level=debug&trace_level=error&time_encoder=rfc3339&caller_encoder=full&file_path=/tmp/server.log",
			},
			want: &Config{
				Encoder: "console",
				Outputs: []string{"stdout", "file"},
				Parameters: struct {
					Level         string
					TraceLevel    string
					TimeEncoder   string
					CallerEncoder string
					FilePath      string
				}{
					Level:         "debug",
					TraceLevel:    "error",
					TimeEncoder:   "rfc3339",
					CallerEncoder: "full",
					FilePath:      "/tmp/server.log",
				},
			},
			wantErr: false,
		},
		{
			name: "",
			args: args{
				dsn: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "",
			args: args{
				dsn: "invalid://[stdout,file]?level=debug&trace_level=error",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "",
			args: args{
				dsn: "console://",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "",
			args: args{
				dsn: "console://stdout]",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "",
			args: args{
				dsn: "console://[invalid_pipe]",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "",
			args: args{
				dsn: "console://[stdout,file].",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "",
			args: args{
				dsn: "console://[stdout,file]?level=debug=2&trace_level=error",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseConfig(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
