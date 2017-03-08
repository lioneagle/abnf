@echo off
setlocal
if exist benchmark.bat goto ok
echo benchmark.bat must be run from its folder
goto end
: ok
call env.bat

if not exist test_temp mkdir test_temp
if exist .\test_temp\test.exe  del .\test_temp\test.exe

go test %1 -bench=. -cpuprofile=.\test_temp\cpu.prof %2
if not exist ./test_temp/cpu.prof goto end

go test %1 -bench=. -c

if not exist test.exe goto end
move test.exe .\test_temp\test.exe

:end
echo finished