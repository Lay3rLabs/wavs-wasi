lint:
    cargo fmt --all -- --check
    cargo fix --allow-dirty --allow-staged
    cargo clippy --all-targets -- -D warnings

wit-build:
    wkg wit build