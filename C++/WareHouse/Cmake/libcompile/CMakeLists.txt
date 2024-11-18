cmake_minimum_required(VERSION 2.8)

project(MyProject)

option(MYDEBUG "enable debug compilation" OFF)

add_subdirectory(lib_test)
add_subdirectory(src)
