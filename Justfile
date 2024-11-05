

build: copy-play-sdk-dlls
    nim c -d:release -o:output/dav2mp4  ./src/dav2mp4.nim

copy-play-sdk-dlls:
    cp vendor/windows/* output

build-installer:
    'C:\Program Files (x86)\Inno Setup 6\ISCC.exe' setup.iss