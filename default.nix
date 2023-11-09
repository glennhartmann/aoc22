{ sources ? import ./nix/sources.nix
, pkgs ? import sources.nixpkgs { config = { }; overlays = [ ]; }
, go_1_19_5_pkgs ? import sources.go_1_19_5_pkgs { config = { }; overlays = [ ]; }
}:
let
  aoc22 = pkgs.callPackage ./aoc22.nix { buildGoModule = go_1_19_5_pkgs.buildGoModule; };
in
{
  inherit aoc22;
  shell = pkgs.mkShell {
    inputsFrom = [
      aoc22
    ];
  };
}
