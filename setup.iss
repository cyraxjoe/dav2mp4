; Inno Setup Script for dav2mp4

[Setup]
AppName=dav2mp4
AppVersion=0.1
AppPublisher=Joel Rivera
AppPublisherURL=https://joel.mx/
AppSupportURL=https://github.com/cyraxjoe/dav2mp4
AppCopyright=© 2024 Joel Rivera
AppContact=Joel Rivera <rivera@joel.mx>
LicenseFile=LICENSE
DefaultDirName={commonpf}\dav2mp4
DefaultGroupName=dav2mp4
OutputBaseFilename=dav2mp4-setup
Compression=lzma
SolidCompression=yes
ArchitecturesInstallIn64BitMode=x64compatible
SetupIconFile=resources\dav2mp4_icon.ico 

[Files]
; Copies all files from the "output" directory into the installation directory
Source: "output\*"; DestDir: "{app}"; Flags: recursesubdirs ignoreversion

[Icons]
; Creates a shortcut for the application in the Start Menu
Name: "{group}\dav2mp4"; Filename: "{app}\dav2mp4.exe"

[Registry]
; Adds the installation directory to the PATH environment variable
Root: HKLM; Subkey: "SYSTEM\CurrentControlSet\Control\Session Manager\Environment"; \
    ValueName: "Path"; ValueType: expandsz; ValueData: "{olddata};{app}"; Flags: createvalueifdoesntexist 

