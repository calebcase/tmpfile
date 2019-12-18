package tmpfile

func ExampleNew() {
	f, err := New("", "example-")
	if err != nil {
		panic(err)
	}
	defer f.Close()
}
