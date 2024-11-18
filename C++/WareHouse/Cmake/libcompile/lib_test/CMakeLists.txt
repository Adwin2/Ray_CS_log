aux_source_directory(. SRC_LIST)

add_library(head_shared SHARED ${SRC_LIST})
add_library(head_static STATIC ${SRC_LIST})

set_target_properties(head_shared PROPERTIES OUTPUT_NAME "head")
set_target_properties(head_static PROPERTIES OUTPUT_NAME "head")

set(LIBRARY_OUTPUT_PATH ${PROJECT_SOURCE_DIR}/lib)
