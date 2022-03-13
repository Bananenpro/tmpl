package templates

const CHelloWorld = `#include <stdio.h>

int main()
{
    printf("Hello World\n");
}`

// '$' need to be replaced with '%' after format
const CMakefile = `CC=%s
CFLAGS=-W -Wall -g
NAME=%s
TARGET=bin/$(NAME)

SRCS := $(wildcard src/*.c)
HDRS := $(wildcard src/*.h)
OBJS := $(patsubst src/§.c,bin/§.o,$(SRCS))

$(TARGET): $(OBJS) $(HDRS) Makefile
	@mkdir -p bin
	$(CC) $(CFLAGS) $(OBJS) -o $(TARGET)

bin/§.o: src/§.c $(HDRS) Makefile
	@mkdir -p bin
	$(CC) $(CFLAGS) -c $< -o $@

clean:
	rm -f $(TARGET) $(OBJS)`

const CCMakeLists = `cmake_minimum_required(VERSION 3.21)

project(%s)

set(C_STANDARD 17)
set(CMAKE_C_FLAGS_DEBUG "${CMAKE_C_FLAGS_DEBUG} -O0")

set(SRC_DIR src)

set(SRC
    ${SRC_DIR}/main.c
)

add_executable(${PROJECT_NAME} ${SRC})`
