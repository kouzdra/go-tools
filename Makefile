all: gocdefs
	./$<  "SC_" <../go-scintilla/scintilla/include/Scintilla.h

gocdefs: gocdefs.go
	go build gocdefs.go
