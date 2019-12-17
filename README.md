# Cross Platform Temporary Files

This library attempts to bridge the gap between the what is provided in
[ioutil.TempFile][ioutil.tempfile] and the best practice of ensuring temporary
files are ***always*** deleted when the application exits.

The normal way to do this on a POSIX system is to use the behavior of
[unlink][posix.unlink] to immediately remove the directory entry for the
temporary file. The OS then ensures that when all open file handles on the file
are close that the file resources are removed. Unfortunately, despite Go having
[os.Remove][os.remove] this does not work on Windows because on Windows it is
necessary to open the files with special flags
([FILE_SHARE_DELETE][windows.flags.share],
[FILE_FLAG_DELETE_ON_CLOSE][windows.flags.on_close]) to enable removing a file
that is open (and ioutil does not do this).

---

[ioutil.tempfile]: https://golang.org/pkg/io/ioutil/#TempFile
[os.remove]: https://golang.org/pkg/os/#Remove
[posix.unlink]: https://pubs.opengroup.org/onlinepubs/9699919799/functions/unlink.html
[windows.flags.on_close]: https://github.com/golang/sys/blob/master/windows/types_windows.go#L108
[windows.flags.share]: https://github.com/golang/sys/blob/master/windows/types_windows.go#L71
