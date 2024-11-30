set(SQL_PARSER-TEST-SRC
	test.cpp
	)

add_executable(sql_parser_test ${SQL_PARSER_TEST_SRC})

target_link_libraries(sql-parser-test ${CMAKE_SOURCE_DIR}/sqlparsercpp/lib/libsqlparser.so)
