{
  # golangci-lint has historically required code changes to support new versions of
  # go so always use the latest specific go version that golangci-lint supports
  # rather than buildGoLatestModule.
  # This can be bumped when the release notes of golangci-lint detail support for
  # new version of go.
  buildGo125Module,
  buildPackages,
  fetchFromGitHub,
  installShellFiles,
  lib,
  stdenv,
}:

buildGo125Module (finalAttrs: {
  pname = "keychain-cli";
  version = "2026.1.2VENDOR_HASH";

  src = fetchFromGitHub {
    owner = "adreasnow";
    repo = "keychain-cli";
    tag = "v${finalAttrs.version}";
    hash = "sha256-H/y4OXZooqG+/6Jb8ZjWefFGPi9/uT/eKXdMCkXniq4=";
  };

  vendorHash = "sha256-BmcA7S8dNSPjvM/x/MXaL5SBbccvZYEhJUHi+Txdt6s=";

  subPackages = [ "keychain-cli" ];

  nativeBuildInputs = [ installShellFiles ];

  ldflags = [
    "-s"
    "-w"
    "-X main.version=${finalAttrs.version}"
    "-X main.commit=v${finalAttrs.version}"
    "-X main.date=1970-01-01T00:00:00Z"
  ];

  postInstall =
    let
      keychainCLIBin =
        if stdenv.buildPlatform.canExecute stdenv.hostPlatform then
          "$out"
        else
          lib.getBin buildPackages.keychain-cli;
    in
    ''
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
