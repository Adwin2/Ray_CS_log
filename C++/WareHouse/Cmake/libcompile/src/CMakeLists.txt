include_directories(../lib_test)

link_directories(${PROJECT_SOURCE_DIR}/lib)

add_compile_options(-std=c++11 -Wall)

add_executable(main main.cpp)

if(MYDEBUG)
	add_executable(main2 main2.cpp)
	target_link_libraries(main2 head)
else()
	message(STATUS "Currently is not debug mode")
endif()

target_link_libraries(main head)

set (EXECUTABLE_OUTPUT_PATH ${PROJECT_SOURCE_DIR}/bin)
