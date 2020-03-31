package tmpfile

import (
	"os"
)

// oTMPFILE is a flag indicating that the open(2) syscall should create
// an unnamed temporary file. When used, the pathname argument to open
// should be a directory instead of a filename. An unnamed inode will be
// created in that directory's filesystem. Anything written to the
// resulting file will be lost when the last file descriptor is closed,
// unless the file is given a name.
//
// Usually known as O_TMPFILE, but that name is not used here to conform
// with Go naming conventions.
//
// oTMPFILE must be specified with one of os.O_RDWR or os.O_WRONLY and,
// optionally, os.O_EXCL. If os.O_EXCL is _not_ specified, then
// "golang.org/x/sys/unix".Linkat() can be used to link the temporary
// file into the filesystem, making it permanent.
//
// Available on Linux 3.11 and later.
const oTMPFILE = 0x410000

// New creates a new unnamed temporary file in the directory dir using
// oTMPFILE so that it never needs to be unlinked with os.Remove().
// The pattern argument is ignored and exists for interface compatibility.
//
// If dir is the empty string it will default to using os.TempDir() as the
// directory for storing the temporary files.
//
// The target directory dir must already exist or an error will result. New
// does not create or remove the directory dir.
func New(dir, pattern string) (f *os.File, err error) {
	if dir == "" {
		dir = os.TempDir()
	}
	return os.OpenFile(dir+"/.", os.O_RDWR|oTMPFILE, 0600)
}
