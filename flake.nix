{
  inputs = {
    nixpkgs = {
      url = "github:NixOS/nixpkgs/nixos-24.11";
    };
  };

  outputs = { self, nixpkgs }:
    let
      allSystems = [
        "x86_64-linux" # 64-bit Intel/AMD Linux
        "aarch64-linux" # 64-bit ARM Linux
        "x86_64-darwin" # 64-bit Intel macOS
        "aarch64-darwin" # 64-bit ARM macOS
      ];
      forAllSystems = f: nixpkgs.lib.genAttrs allSystems (system: f {
        inherit system;
        pkgs = import nixpkgs {
          inherit system;
        };
      });
    in
    {
      packages = forAllSystems ({ system, pkgs, ... }:
        {
          default = pkgs.buildGoModule {
            name = "hashdir";
            src = ./.;
            vendorHash = null;

            # Optional flags.
            CGO_ENABLED = 0;
            flags = [ "-trimpath" ];
            ldflags = [ "-s" "-w" "-extldflags -static" ];
          };
        });

      devShells = forAllSystems ({ system, pkgs }: {
        default = pkgs.mkShell {
          packages = with pkgs; [
            go
            gotools
          ];
        };
      });
    };
}
