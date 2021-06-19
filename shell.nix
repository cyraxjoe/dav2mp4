{ nixpkgs ? <nixpkgs> }:
let
  dav2mp4 = import ./default.nix { inherit nixpkgs; };
in
with(import nixpkgs {}); mkShell {
  nativeBuildInputs = [ dav2mp4 ];
}

