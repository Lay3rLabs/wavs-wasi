lint:
    cargo fmt --all -- --check
    cargo fix --allow-dirty --allow-staged
    cargo clippy --all-targets -- -D warnings

wit-clean:
    rm -f wit-definitions/types/wavs:types@*.wasm
    rm -f wit-definitions/worker/wavs:worker@*.wasm

wit-build:
    just wit-clean
    cd wit-definitions/types && wkg wit build
    cd wit-definitions/worker && wkg wit build

wit-publish:
    wkg publish wit-definitions/types/wavs:types@0.5.0-alpha.3.wasm
    wkg publish wit-definitions/worker/wavs:worker@0.5.0-alpha.3.wasm