set(BYTE_YOUNG_SRC
	main.cpp
	parser.cpp
	optimizer.cpp
	executor.cpp
	metadata.cpp
	trx.cpp
	storage.cpp
	util.cpp
	)

add_executable(bydb ${BYTE_YOUNG_SRC})

target_link_libraries(bydb ${CMAKE_SOURCE_DIR}/sqlparsercpp/lib/libsqlparser.so)
