@echo off
echo Detecting OS...
:: Starting project for only Windows systems.

:: Start the server
cd server
if exist start_server.bat (
    call start_server.bat
) else (
    echo Error: Missing start_server.bat
    exit /b 1
)
cd ..

:: Start the client
cd client
if exist start_client.bat (
    call start_client.bat
) else (
    echo Error: Missing start_client.bat
    exit /b 1
)
cd ..
