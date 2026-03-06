@echo off
setlocal enabledelayedexpansion
chcp 65001 >nul

echo.
set user_text=Atualizacao
if not "%~1"=="" set user_text=%~1
:: echo %date% %time% - %user_text%> last-delivery.txt

:: echo.
:: echo -----
:: time /t
:: echo Initializing merge request
:: echo -----
:: echo.

git add .
git commit -m "%user_text%"
git push origin main

echo.
echo -----
time /t
echo Merge request finished
echo -----
echo.
