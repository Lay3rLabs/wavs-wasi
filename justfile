lint:
    cargo fmt --all -- --check
    cargo fix --allow-dirty --allow-staged
    cargo clippy --all-targets -- -D warnings

wit-clean:
    rm -f wit-definitions/types/wavs:types@*.wasm
    rm -f wit-definitions/worker/wavs:worker@*.wasm

wit-build config="":
    just inner-wit-build "{{ if config != '' { ' --config ' + '../../' + config } else { '' } }}"

wit-publish config="":
    just inner-wit-publish "{{ if config != '' { ' --config ' + '../../' + config } else { '' } }}"

inner-wit-build config-arg:
    just wit-clean
    cd wit-definitions/types && wkg wit build{{config-arg}}
    cd wit-definitions/worker && wkg wit build{{config-arg}}
    cd wit-definitions/aggregator && wkg wit build{{config-arg}}

inner-wit-publish config-arg:
    cd wit-definitions/types && wkg publish wavs:types@*.wasm{{config-arg}}
    cd wit-definitions/worker && wkg publish wavs:worker@*.wasm{{config-arg}}
    cd wit-definitions/aggregator && wkg publish wavs:aggregator@*.wasm{{config-arg}}

cargo-check:
    cd packages/wavs-wasi-utils && cargo check --all-targets --all-features