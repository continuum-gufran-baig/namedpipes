#include <windows.h> 
#include <stdio.h>
#define BUFSIZE 1024
#define PIPE_TIMEOUT 5000

void main()
{
    BOOL flg;
    DWORD cbBytesRead;
    char szPipeIn[BUFSIZE];
    HANDLE hFile;

    hFile = CreateFile(L"\\\\.\\pipe\\MyPipe",
        GENERIC_READ,
        0,
        NULL,
        OPEN_EXISTING,
        0,
        NULL);

    if (hFile != INVALID_HANDLE_VALUE)
    {
        for (;;)
        {
            PeekNamedPipe(hFile, NULL, NULL, NULL, &cbBytesRead, NULL);
            if (cbBytesRead != 0)
            {
                flg = ReadFile(hFile, // handle to pipe 
                    szPipeIn, // buffer to receive data 
                    BUFSIZE, // size of buffer 
                    &cbBytesRead, // number of bytes read 
                    NULL); // not overlapped I/O 

                if (flg != FALSE)
                {
                    szPipeIn[cbBytesRead] = '\0';
                    printf("Data Received: %s\n", szPipeIn);
                    FlushFileBuffers(hFile);
                }
            }
        }
    }
}