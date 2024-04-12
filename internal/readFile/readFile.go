package readfile

import "os"

func ReadOneFile(dir string) ([]byte, error) {

	return os.ReadFile(dir)
}
