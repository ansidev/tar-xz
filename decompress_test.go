package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

const testDataDir = "testdata"
const testFileHash string = "dc8897816b30ddee02e83f4d285916731d7940d8b6fc03aab4e71185649034fd"

func getSHA256Checksum(filePath string) (string, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", sha256.Sum256(buf)), nil
}

func removeDir(t *testing.T, dirPath string) {
	err := os.RemoveAll(dirPath)
	if err != nil {
		t.Errorf("Failed to remove the directory: %v", err)
	}
}

func compareFileHash(t *testing.T, actualDirPath string, expectedDirPath string, verifyPaths []string) {
	for _, path := range verifyPaths {
		actualFilePath := filepath.Join(actualDirPath, path)
		expectedFilePath := filepath.Join(expectedDirPath, path)

		actualFileHash, err := getSHA256Checksum(actualFilePath)
		if err != nil {
			t.Errorf("Failed to get SHA256 checksum of the actual file: %v", err)
		}

		expectedFileHash, err := getSHA256Checksum(expectedFilePath)
		if err != nil {
			t.Errorf("Failed to get SHA256 checksum of the expected file: %v", err)
		}

		if actualFileHash != expectedFileHash {
			t.Errorf("Mismatched checksum. Got: %v, want: %v", actualFileHash, expectedFileHash)
		}
	}
}

func testGetSHA256Checksum(t *testing.T) {
	sum, err := getSHA256Checksum(filepath.Join(testDataDir, "sample.tar.xz"))
	if err != nil {
		t.Errorf("Internal function getSHA256Checksum() error: %v", err)
	}

	if sum != testFileHash {
		t.Errorf("Unmatched test file checksum. Got: %v, want: %v", sum, testFileHash)
	}
}

func TestDecompress(t *testing.T) {
	type args struct {
		filePath    string
		destDir     string
		verifyPaths []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Decompress sample.tar.xz to ./sample_test",
			args: args{
				filePath: filepath.Join(testDataDir, "sample.tar.xz"),
				destDir:  filepath.Join(testDataDir, "sample_test"),
				verifyPaths: []string{
					filepath.Join("sample", "MIT_LICENSE"),
					filepath.Join("sample", "MIT_LICENSE.pdf"),
				},
			},
			wantErr: false,
		},
	}

	// Test getSHA256Checksum first
	testGetSHA256Checksum(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Remove the directory if exists before running the test case
			if _, err := os.Stat(tt.args.destDir); !os.IsNotExist(err) {
				removeDir(t, tt.args.destDir)
			}
			err := Decompress(tt.args.filePath, tt.args.destDir)

			// Remove the directory after running the test case
			defer removeDir(t, tt.args.destDir)

			if (err != nil) != tt.wantErr {
				t.Errorf("Decompress() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				compareFileHash(t, testDataDir, tt.args.destDir, tt.args.verifyPaths)
			}
		})
	}
}
