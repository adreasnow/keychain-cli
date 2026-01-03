{
  # golangci-lint has historically required code changes to support new versions of
  # go so always use the latest specific go version that golangci-lint supports
  # rather than buildGoLatestModule.
  # This can be bumped when the release notes of golangci-lint detail support for
  # new version of go.
  buildGoModule,
  buildPackages,
  fetchFromGitHub,
  installShellFiles,
  lib,
  stdenv,
}:

buildGoModule (finalAttrs: {
  pname = "keys";
  version = "2026.1.13";

  src = fetchFromGitHub {
    owner = "adreasnow";
    repo = "keys";
    tag = "v${finalAttrs.version}";
    hash = "sha256-GLVvtsPRz9gzc3uOXz+SRqbs9KgZLjLtTF0KPDPsihQ=";
  };
  vendorHash = "sha256-diCC3gA2hnAlzH3E7syMrKD3yebU+ZsSnicWg8ZW6x8=";

  ldflags = [ "-s -w -X main.version=${finalAttrs.version}" ];

  nativeBuildInputs = [ installShellFiles ];

  subPackages = [ "cmd/keys" ];

  postInstall =
    let
      keychainCLIBin =
        if stdenv.buildPlatform.canExecute stdenv.hostPlatform then
          "$out"
        else
          lib.getBin buildPackages.keys;
    in
    ''
      ls -lah ${keychainCLIBin}/bin/
      installShellCompletion --cmd keys \
        --bash <(${keychainCLIBin}/bin/keys completion bash) \
        --fish <(${keychainCLIBin}/bin/keys completion fish) \
        --zsh <(${keychainCLIBin}/bin/keys completion zsh)
    '';

  meta = {
    description = "Lightweight Go wrapper around the apple keychain to act as a simple CLI tool for managing secrets.";
    homepage = "https://github.com/adreasnow/keys";
    mainProgram = "keys";
    license = lib.licenses.gpl3Plus;
    maintainers = with lib.maintainers; [
      adreasnow
    ];
  };
})
