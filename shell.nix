{pkgs ? import <nixpkgs> {}}:
with pkgs;
  pkgs.mkShell {
    buildInputs = [
      go
      cargo
      rustc
      cargo
      rustfmt
      rust-analyzer
      clippy
      openssl
      pkg-config
      libiconv
      aoc-cli
    ];
  }
