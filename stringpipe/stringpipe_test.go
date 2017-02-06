package stringpipe

import "testing"

func TestAll(t *testing.T) {
	input := `1234 5678 foo 1
1234 5679 foo 2
abcd 5680 foo 3
1234 5681 bar 1
1234 5682 bar 2
abcd 5683 bar 3
1234 5684 raz 1
1234 5685 raz 2
abcd 5686 raz 3
`
	t.Log("grep abcd - ", StringPipe(input).Grep("abcd"))
	t.Log("grep -v raz - ", StringPipe(input).Grep("abcd").VGrep("raz"))
	t.Log("cut columns - ", StringPipe(input).Grep("abcd").VGrep("raz").CutOutColumns(0, 1))
	t.Log("sort - ", StringPipe(input).Grep("abcd").VGrep("raz").CutOutColumns(0, 1).Sort())

	expected := "bar 3\nfoo 3"
	output := StringPipe(input).Grep("abcd").VGrep("raz").CutOutColumns(0, 1).Sort().String()
	if output != expected {
		t.Errorf("got %q, want %q", output, expected)
	}
}
