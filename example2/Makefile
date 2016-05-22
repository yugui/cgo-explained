CGO_SRCS=export_example.go main.go

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
GOPKG=github.com/yugui/cgo-explained/example2

.SUFFIXES: .go .dot .png

all: _cgo_.o

.c.o:
	gcc -c -o $@ -fPIC $<
.dot.png:
	dot -Tpng $< > $@ || (rm -f $@; exit 1)

clean:
	rm -f $(GENFILES) *.o example2.a example2 _cgo_import.go
$(GENFILES): $(CGO_SRCS) use_exported.h
	env CGO_LDFLAGS=-lm go tool cgo -objdir .  \
		-importpath $(GOPKG) $(CGO_SRCS)
	rm _cgo_.o

use_exported.o: use_exported.c use_exported.h
_cgo_.o: $(GENFILES_C:.c=.o) $(EXPORT_H) _cgo_main.o use_exported.o
	# gcc -shared -fPIC -o _cgo_.o $(GENFILES_C) _cgo_export.c _cgo_main.c -lm
	gcc -o _cgo_.o $(GENFILES_C:.c=.o) _cgo_main.o use_exported.o -lm
_all.o: $(GENFILES_C:.c=.o) $(EXPORT_H) use_exported.o
	gcc -fPIC -nostdlib -o _all.o -Wl,-r $(GENFILES_C:.c=.o) use_exported.o
_cgo_import.go: _cgo_.o
	go tool cgo -dynimport _cgo_.o -dynout _cgo_import.go \
		-dynpackage main

example2.a: $(GENFILES_GO) _cgo_import.go main.go _all.o
	go tool compile -pack -o example2.a $(GENFILES_GO) _cgo_import.go
	go tool pack r example2.a _all.o
example2: example2.a
	go tool link -buildmode exe -o example2 example2.a
