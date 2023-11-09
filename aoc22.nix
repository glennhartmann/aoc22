{ buildGoModule
}:
buildGoModule {
  pname = "aoc22";
  version = "v0.0.4";
  src = builtins.path { path = ./.; name = "aoc22"; };
  vendorHash = "sha256-ZPRNYpdQnPeJ/6rr1sf6N2RcYU2rhGtMa5eTQGTPuvA=";
}
