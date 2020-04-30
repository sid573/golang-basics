package pack

import "os"

// OSCustomOpen for opening a file
func OSCustomOpen(f1 string, val []byte) (*os.File, error) {
	f, err := os.OpenFile(string(f1), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	_, err1 := f.Write(val)
	if err1 != nil {
		return nil, err1
	}

	return f, nil
}
