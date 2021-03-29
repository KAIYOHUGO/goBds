@echo off

echo server start
:w
set /p a="wait input :"
if %a% neq stop goto w
echo server stop