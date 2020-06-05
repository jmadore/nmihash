package nmihash

import (
	"testing"
)

func TestCalculateHash(t *testing.T) {
	type args struct {
		hashKey    string
		terminalID string
		refNum     string
		amount     string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Valid Hash",
			args: args{
				hashKey:    "trVxrnoz22bvwvnV",
				terminalID: "99999999",
				refNum:     "0000000765",
				amount:     "1.23",
			},
			want:    "7PtU022473m+ntcZY2wt6pXzKWc=",
			wantErr: false,
		},
		{name: "Another Valid Hash",
			args: args{
				hashKey:    "trVxrnoz22bvwvnV",
				terminalID: "12345678",
				refNum:     "1234567890",
				amount:     "999.99",
			},
			want:    "5bkisjvitcyaNUIvbjEtc+N4oZA=",
			wantErr: false,
		},
		{name: "Expect err when hashkey blank",
			args: args{
				hashKey:    "",
				terminalID: "99999999",
				refNum:     "0000000765",
				amount:     "1.23",
			},
			want:    "",
			wantErr: true,
		},
		{name: "Expect err when terminalID blank",
			args: args{
				hashKey:    "trVxrnoz22bvwvnV",
				terminalID: "",
				refNum:     "0000000765",
				amount:     "1.23",
			},
			want:    "",
			wantErr: true,
		},
		{name: "Expect err when refNum blank",
			args: args{
				hashKey:    "trVxrnoz22bvwvnV",
				terminalID: "99999999",
				refNum:     "",
				amount:     "1.23",
			},
			want:    "",
			wantErr: true,
		},
		{name: "Expect err when amount blank",
			args: args{
				hashKey:    "trVxrnoz22bvwvnV",
				terminalID: "99999999",
				refNum:     "0000000765",
				amount:     "",
			},
			want:    "",
			wantErr: true,
		},
		{name: "Expect err when all blank",
			args: args{
				hashKey:    "",
				terminalID: "",
				refNum:     "",
				amount:     "",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateHash(tt.args.hashKey, tt.args.terminalID, tt.args.refNum, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
