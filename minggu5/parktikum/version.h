#ifndef VERSION_H
#define VERSION_H

#include <iostream>
#include <string>

// =====================
// ANSI COLOR CODES
// =====================
#define RESET   "\033[0m"
#define RED     "\033[31m"
#define GREEN   "\033[32m"
#define YELLOW  "\033[33m"
#define BLUE    "\033[34m"
#define CYAN    "\033[36m"
#define MAGENTA "\033[35m"
#define BOLD    "\033[1m"

// =====================
// APP INFO
// =====================
#define APP_NAME_DEFAULT  "@pandora"
#define APP_VERSION       "0.14.5(CBT)"
#define BUILD_DATE        __DATE__
#define BUILD_TIME        __TIME__

// =====================
// COMPILER INFO
// =====================
#if defined(__clang__)
    #define COMPILER_NAME "Clang/LLVM"
    #define COMPILER_VERSION __clang_version__

#elif defined(__GNUC__)
    #define COMPILER_NAME "GCC"
    #define COMPILER_VERSION __VERSION__

#elif defined(_MSC_VER)
    #define COMPILER_NAME "MSVC"
    #define STRINGIFY(x) #x
    #define TOSTRING(x) STRINGIFY(x)
    #define COMPILER_VERSION "MSVC " TOSTRING(_MSC_VER)

#elif defined(__MINGW64__)
    #define COMPILER_NAME "MinGW-w64"
    #define COMPILER_VERSION __VERSION__

#elif defined(__MINGW32__)
    #define COMPILER_NAME "MinGW"
    #define COMPILER_VERSION __VERSION__

#elif defined(__INTEL_COMPILER)
    #define COMPILER_NAME "Intel C/C++ (ICC/ICX)"
    #define STRINGIFY(x) #x
    #define TOSTRING(x) STRINGIFY(x)
    #define COMPILER_VERSION TOSTRING(__INTEL_COMPILER)

#elif defined(__NVCOMPILER)
    #define COMPILER_NAME "NVIDIA HPC/NVCC"
    #define COMPILER_VERSION "NVIDIA " __VERSION__

#elif defined(__arm__) && defined(__ARMCC_VERSION)
    #define COMPILER_NAME "ARM Compiler (armcc)"
    #define STRINGIFY(x) #x
    #define TOSTRING(x) STRINGIFY(x)
    #define COMPILER_VERSION TOSTRING(__ARMCC_VERSION)

#elif defined(__CC_ARM)
    #define COMPILER_NAME "Keil ARM Compiler"
    #define STRINGIFY(x) #x
    #define TOSTRING(x) STRINGIFY(x)
    #define COMPILER_VERSION TOSTRING(__ARMCC_VERSION)

#elif defined(__IAR_SYSTEMS_ICC__)
    #define COMPILER_NAME "IAR Embedded Workbench"
    #define STRINGIFY(x) #x
    #define TOSTRING(x) STRINGIFY(x)
    #define COMPILER_VERSION TOSTRING(__VER__)

#elif defined(__TI_COMPILER_VERSION__)
    #define COMPILER_NAME "TI ARM Compiler"
    #define STRINGIFY(x) #x
    #define TOSTRING(x) STRINGIFY(x)
    #define COMPILER_VERSION TOSTRING(__TI_COMPILER_VERSION__)

#elif defined(__EMSCRIPTEN__)
    #define COMPILER_NAME "Emscripten (WebAssembly)"
    #define COMPILER_VERSION __VERSION__

#elif defined(__SDCC)
    #define COMPILER_NAME "SDCC (Small Device C Compiler)"
    #define COMPILER_VERSION "SDCC"

#elif defined(__XTENSA__) || defined(ESP_PLATFORM)
    #define COMPILER_NAME "ESP-IDF Xtensa GCC"
    #ifdef __VERSION__
        #define COMPILER_VERSION __VERSION__
    #else
        #define COMPILER_VERSION "ESP-IDF GCC"
    #endif

#elif defined(__BORLANDC__)
    #define COMPILER_NAME "Borland C++"
    #define COMPILER_VERSION __BORLANDC__

#elif defined(__DMC__)
    #define COMPILER_NAME "Digital Mars C++"
    #define COMPILER_VERSION "Digital Mars"

#elif defined(__TURBOC__)
    #define COMPILER_NAME "Turbo C/C++"
    #define COMPILER_VERSION "TurboC"

#elif defined(__WATCOMC__)
    #define COMPILER_NAME "Open Watcom C/C++"
    #define COMPILER_VERSION __WATCOMC__

#else
    #define COMPILER_NAME "Unknown Compiler"
    #ifdef __VERSION__
        #define COMPILER_VERSION __VERSION__
    #else
        #define COMPILER_VERSION "N/A"
    #endif
#endif

// =====================
// ARCHITECTURE INFO
// =====================
#if defined(__x86_64__) || defined(_M_X64)
    #define ARCH_NAME "x86_64"
#elif defined(__i386__) || defined(_M_IX86)
    #define ARCH_NAME "x86"
#elif defined(__arm__) || defined(_M_ARM)
    #define ARCH_NAME "ARM"
#elif defined(__aarch64__)
    #define ARCH_NAME "ARM64"
#elif defined(__riscv)
    #define ARCH_NAME "RISC-V"
#else
    #define ARCH_NAME "Unknown Architecture"
#endif

// =====================
// OS DETECTION
// =====================
#if defined(_WIN32)
    #define OS_NAME "Windows"
#elif defined(__linux__)
    #define OS_NAME "Linux"
#elif defined(__APPLE__)
    #define OS_NAME "macOS"
#elif defined(__ANDROID__)
    #define OS_NAME "Android"
#elif defined(ESP_PLATFORM)
    #define OS_NAME "ESP32 / ESP-IDF"
#else
    #define OS_NAME "Unknown OS"
#endif

// =====================
// HELPER NAMESPACE
// =====================
namespace version {

inline void print_version(const std::string& program_name = APP_NAME_DEFAULT) {
    std::cout << BOLD << CYAN << program_name << RESET << " "
              << GREEN << "v" << APP_VERSION << RESET
              << " (" << YELLOW << BUILD_DATE << " " << BUILD_TIME << RESET << ")\n";

    std::cout << "Compiler: " << BLUE << COMPILER_NAME << RESET
              << " " << COMPILER_VERSION << "\n";

    std::cout << "Target: " << MAGENTA << ARCH_NAME << RESET << "\n";
    std::cout << "OS: " << CYAN << OS_NAME << RESET << "\n\n";
}

inline void print_help(const std::string& program_name = APP_NAME_DEFAULT) {
    using std::cout;
    using std::endl;

    cout << BOLD << YELLOW << "usage:" << RESET << " " << program_name
         << " [<options>] <command> [<args>]\n\n";

    cout << BOLD << CYAN << "options:" << RESET << "\n";
    cout << "    " << GREEN << "-h, --help" << RESET
         << "            Show this help message and exit\n";
    cout << "    " << GREEN << "-v, --version" << RESET
         << "         Show version information\n";

    cout << BOLD << CYAN << "commands:" << RESET << "\n";
    cout << "    " << MAGENTA << "soal1 [auto]" << RESET
         << "          Run soal1 (with optional 'auto' mode)\n";
    cout << "    " << MAGENTA << "soal2 [auto]" << RESET
         << "          Run soal2 (with optional 'auto' mode)\n\n";

    cout << BOLD << CYAN << "examples:" << RESET << "\n";
    cout << "    " << GREEN << program_name << RESET << " soal1\n";
    cout << "    " << GREEN << program_name << RESET << " soal1 auto\n";
    cout << "    " << GREEN << program_name << RESET << " soal2\n";
    cout << "    " << GREEN << program_name << RESET << " soal2 auto\n";
    cout << endl;
}
} // namespace helper

#endif // VERSION_H
