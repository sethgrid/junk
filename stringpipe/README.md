#StringPipe

StringPipe is a tool I've used to help trim up a large string with newlines, like a log.

The idea is to be able to `StringPipe(str).Grep("thing").VGrep("not thing").CutOutColumns(1,2).Sort()`