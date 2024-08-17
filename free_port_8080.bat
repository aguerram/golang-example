@echo off
FOR /F "tokens=5" %%A IN ('netstat -ano ^| findstr :8080') DO (
    echo Killing process with PID %%A
    taskkill /PID %%A /F
)
pause
