{ pkgs ? (import <nixpkgs> {}) }:
let
  inherit(pkgs) stdenv makeWrapper;
in

stdenv.mkDerivation {
  name = "dav2mp4";
  src = ./.;
  nativeBuildInputs = with pkgs;
    [ autoPatchelfHook nim alsaLib
      libGL stdenv.cc.cc.lib ] ++ ( with xorg; [libX11 libXv]);
  buildInputs = [ makeWrapper ];
  buildPhase = ''
   export HOME=$TMPDIR
   nim c -d:release -o:dav2mp4  ./src/dav2mp4.nim
  '';

  installPhase = ''
    mkdir -p $out/bin
    mkdir -p $out/lib/playsdk_log
    cp vendor/linux/dhplay/libdhplay.so $out/lib
    cp dav2mp4 $out/bin
    wrapProgram $out/bin/dav2mp4 --set LD_LIBRARY_PATH "$out/lib"
  '';

}
