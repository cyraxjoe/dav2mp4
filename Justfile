set windows-powershell := true

build: copy-play-sdk-dlls
    nim c -d:release -o:output/dav2mp4  ./src/dav2mp4.nim

copy-play-sdk-dlls:
    cp vendor\windows\* output
    
run *args:
    ./output/dav2mp4 {{ args }}