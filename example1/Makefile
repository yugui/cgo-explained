CGO_SRCS=import_example.go

GENFILES_GO=_cgo_gotypes.go \
	    $(CGO_SRCS:.go=.cgo1.go)
GENFILES_C=_cgo_export.c \
	   $(CGO_SRCS:.go=.cgo2.c)
EXPORT_H=_cgo_export.h
GENFILES= $(GENFILES_GO) \
	  $(GENFILES_C) \
	  $(EXPORT_H) \
	  _cgo_main.c \
	  _cgo_flags
GOPKG=github.com/yugui/cgo-explained/example1

.SUFFIXES: .go .dot .png

all: _cgo_.o

.c.o:
	gcc -c -o $@ -fPIC $<
.dot.png:
	dot -Tpng $< > $@ || (rm -f $@; exit 1)

clean:
	rm -f $(GENFILES) *.o example1.a example1 _cgo_import.go
$(GENFILES): $(CGO_SRCS)
	env CGO_LDFLAGS=-lm go tool cgo -objdir .  \
		-importpath $(GOPKG) $(CGO_SRCS)
	rm _cgo_.o
_cgo_.o: $(GENFILES_C:.c=.o) $(EXPORT_H) _cgo_main.o
	# gcc -shared -fPIC -o _cgo_.o $(GENFILES_C) _cgo_export.c _cgo_main.c -lm
	gcc -o _cgo_.o $(GENFILES_C:.c=.o) _cgo_main.o -lm
_all.o: $(GENFILES_C:.c=.o) $(EXPORT_H)
	gcc -fPIC -nostdlib -o _all.o -Wl,-r $(GENFILES_C:.c=.o)
_cgo_import.go: _cgo_.o
	go tool cgo -dynimport _cgo_.o -dynout _cgo_import.go \
		-dynpackage main

example1.a: $(GENFILES_GO) _cgo_import.go main.go _all.o
	go tool compile -pack -o example1.a $(GENFILES_GO) _cgo_import.go main.go
	go tool pack r example1.a _all.o
example1: example1.a
	go tool link -v -o example1  example1.a
run: example1
	./example1
