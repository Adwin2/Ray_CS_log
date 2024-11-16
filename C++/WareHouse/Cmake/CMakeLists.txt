cmake_minimum_required(VERSION 2.8) 

project(demo)

set (SRC_LIST
	./main.cpp
	./head.cpp)

add_executable(demo ${SRC_LIST})
