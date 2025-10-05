cmake_minimum_required(VERSION 3.10)

# Nama project
project(MyProject)

# Cari semua file .cpp secara rekursif
file(GLOB_RECURSE SRC_FILES 
    "${CMAKE_SOURCE_DIR}/pelajaran/*.cpp"
)

# Buat executable dari semua file cpp
add_executable(${PROJECT_NAME} ${SRC_FILES})
