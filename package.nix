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
  pname = "keychain-cli";
  version = "2026.1.9";

  src = fetchFromGitHub {
    owner = "adreasnow";
    repo = "keychain-cli";
    tag = "v${finalAttrs.version}";
    hash = "sha256-OrSRh1bipzhpf466X2RVXCiimE/FN/FfiNDCPfthSa8=";
  };
  vendorHash = "sha256-diCC3gA2hnAlzH3E7syMrKD3yebU+ZsSnicWg8ZW6x8=";

  ldflags = [ "-s -w -X main.version=${finalAttrs.version}" ];

  nativeBuildInputs = [ installShellFiles ];

  postInstall =
    let
      keychainCLIBin =
        if stdenv.buildPlatform.canExecute stdenv.hostPlatform then
          "$out"
        else
          lib.getBin buildPackages.keychain-cli;
    in
    ''
      ls -lah ${keychainCLIBin}/bin/
      installShellCompletion --cmd keychain-cli \
        --bash <(${keychainCLIBin}/bin/keychain-cli completion bash) \
        --fish <(${keychainCLIBin}/bin/keychain-cli completion fish) \
        --zsh <(${keychainCLIBin}/bin/keychain-cli completion zsh)
    '';

  meta = {
    description = "Lightweight Go wrapper around the apple keychain to act as a simple CLI tool for managing secrets.";
    homepage = "https://github.com/adreasnow/keychain-cli";
    mainProgram = "keychain-cli";
    license = lib.licenses.gpl3Plus;
    maintainers = with lib.maintainers; [
      adreasnow
    ];
  };
})
