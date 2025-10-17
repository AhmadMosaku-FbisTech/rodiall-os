# Makefile for rodial-utils
# Builds both C static library and Go module

CC := $(TARGET_CC)
AR := $(TARGET_AR)
CFLAGS := -Wall -O2
LIBNAME := librodialutils.a

all: $(LIBNAME)

$(LIBNAME): cli_helper.o
	$(AR) rcs $@ $^

cli_helper.o: cli_helper.c cli_helper.h
	$(CC) $(CFLAGS) -c cli_helper.c -o cli_helper.o

clean:
	rm -f *.o *.a

# Optional: build Go wrapper (for testing)
go:
	go build -o test_rodial_utils cli_helper.go

.PHONY: all clean go
