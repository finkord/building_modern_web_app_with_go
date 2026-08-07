@echo off
go build -o bookings.exe .\cmd\web
if %errorlevel% neq 0 (
    echo Build failed!
    exit /b %errorlevel%
)
bookings.exe