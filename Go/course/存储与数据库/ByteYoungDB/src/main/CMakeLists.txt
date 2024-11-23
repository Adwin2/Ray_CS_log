cmake_minimum_required(VERSION 2.8)

project(MyProject)

aux_source_directory(. SRC_LIST)

set(SRC_LIST
	main.cpp
	parser.cpp
	executor.cpp
	)
