package upload

import "testing"

/*
curl --location --request POST 'http://127.0.0.1:36677/upload' --header 'Content-Type: application/json' --data-raw '{
    "list": [
        "/mnt/c/Users/dengjiawen/Downloads/ddd1.md"
    ]
}'
*/

func TestUploaderImpl_Upload(t *testing.T) {
	u := NewUploader()
	type args struct {
		localPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				localPath: "/mnt/c/Users/dengjiawen/Downloads/ddd1.md",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := u.Upload(tt.args.localPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploaderImpl.Upload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("UploaderImpl.Upload() = %v", got)
		})
	}
}
