cmake_minimum_required(VERSION 3.10)

project(SLL)

file(GLOB_RECURSE SRC_FILES
    "${CMAKE_SOURCE_DIR}/TR/*.cpp"
)

add_executable(${PROJECT_NAME} ${SRC_FILES})