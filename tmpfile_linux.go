package tmpfile

import (
	"io/ioutil"
	"os"
	"syscall"
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

	f, err = os.OpenFile(dir+"/.", os.O_RDWR|oTMPFILE, 0600)
	if err != nil {
		if pathErr, ok := err.(*os.PathError); ok {
			if sysErr, ok := pathErr.Err.(syscall.Errno); ok {
				// receiving any of these errors suggests the underlying
				// filesystem does not support O_TMPFILE. There are other
				// situations where these errors could arise, but in all
				// cases, we at least know that it is safe to try again
				// without O_TMPFILE.
				if sysErr == syscall.EINVAL || sysErr == syscall.EISDIR || sysErr == syscall.EOPNOTSUPP {
					// try creating a tempfile the old-fashioned way
					f, err = ioutil.TempFile(dir, pattern)
					if err != nil {
						return f, err
					}

					err = os.Remove(f.Name())
				}
			}
		}
	}

	return f, err
}
