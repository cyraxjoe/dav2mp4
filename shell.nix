{ nixpkgs ? <nixpkgs> }:
let
  dav2mp4 = import ./default.nix { pkgs = (import nixpkgs {}); };
in
with(import nixpkgs {}); mkShell {
  nativeBuildInputs = [ dav2mp4 ];
}

