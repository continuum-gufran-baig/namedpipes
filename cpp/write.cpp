#include <iostream>
#include <windows.h>
using namespace std;

int main(int argc, const char** argv)
{
    wcout << "Connecting to pipe..." << endl;

    // Open the named pipe
    // Most of these parameters aren't very relevant for pipes.
    HANDLE pipe = CreateFile(
        L"\\\\.\\pipe\\MyPipe",
        GENERIC_READ | GENERIC_WRITE, // only need read access
        0,
        NULL,
        OPEN_EXISTING,
        FILE_FLAG_OVERLAPPED | SECURITY_SQOS_PRESENT | SECURITY_ANONYMOUS,
        NULL
    );

    if (pipe == INVALID_HANDLE_VALUE) {
        wcout << "Failed to connect to pipe." << GetLastError();
        // look up error code here using GetLastError()
        system("pause");
        return 1;
    }

    DWORD BytesWritten;
    wcout << "Writing to pipe..." << endl;

    if (WriteFile(pipe, "test!!!", 14, &BytesWritten, NULL) == 0) {
        printf("WriteFile failed with error %d\n", GetLastError());
        CloseHandle(pipe);
    }

    wcout << "Data has been sent to pipe..." << endl;

    return 0;
}
