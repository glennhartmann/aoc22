{
  inputs = {
    nixpkgs.url = github:NixOS/nixpkgs/55070e598e0e03d1d116c49b9eff322ef07c6ac6; # go1.19.5
    flake-compat.url = "https://flakehub.com/f/edolstra/flake-compat/1.tar.gz";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs = { self, nixpkgs, flake-compat, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
        aoc22 = pkgs.buildGoModule {
          pname = "aoc22";
          version = "v0.0.4";
          src = builtins.path { path = ./.; name = "aoc22"; };
          vendorHash = "sha256-ZPRNYpdQnPeJ/6rr1sf6N2RcYU2rhGtMa5eTQGTPuvA=";
        };
      in
      {
        packages = {
          inherit aoc22;
          default = aoc22;
        };
      }
    );
}
