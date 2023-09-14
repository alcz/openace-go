set OLDPATH=%OLDPATH%
set CGO_ENABLED=1
set CGO_CPPFLAGS=-IC:\HARBOUR\include
set CGO_LDFLAGS=-L./ openace32.dll -Wl,--enable-stdcall-fixup -Wl,-verbose
call path.cmd
go build
set PATH=%OLDPATH%
