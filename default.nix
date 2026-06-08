let
  pkgs = import
    (builtins.fetchTarball {
      url = "https://channels.nixos.org/nixos-25.11/nixexprs.tar.xz";
    })
    { };
in
{
  env = pkgs.buildEnv {
    name = "env";
    paths = [
      pkgs.go
    ];
  };
}
