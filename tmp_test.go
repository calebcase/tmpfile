package tmp

import (
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	var err error

	for _, tc := range []struct {
		Name    string
		TmpDir  string
		TmpName string
	}{
		{"basic", "", "tmp-test-"},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			var f *os.File

			f, err = File(tc.TmpDir, tc.TmpName)
			if err != nil {
				t.Error(err)
			}

			t.Logf("Created temporary file: %q\n", f.Name())

			_, err = os.Stat(f.Name())
			if err == nil || !os.IsNotExist(err) {
				t.Errorf("Temporary file exists, but it should have been unlinked already: %+v\n", err)
			}
		})
	}
}
