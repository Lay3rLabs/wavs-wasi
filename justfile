lint:
    cargo fmt --all -- --check
    cargo fix --allow-dirty --allow-staged
    cargo clippy --all-targets -- -D warnings

wit-clean:
    rm -f wit-definitions/types/wavs:types@*.wasm
    rm -f wit-definitions/operator/wavs:operator@*.wasm

wit-build config="":
    just inner-wit-build "{{ if config != '' { ' --config ' + '../../' + config } else { '' } }}"

wit-publish config="":
    just inner-wit-publish "{{ if config != '' { ' --config ' + '../../' + config } else { '' } }}"

inner-wit-build config-arg:
    just wit-clean
    cd wit-definitions/types && wkg wit build{{config-arg}}
    cd wit-definitions/operator && wkg wit build{{config-arg}}
    cd wit-definitions/aggregator && wkg wit build{{config-arg}}

inner-wit-publish config-arg:
    cd wit-definitions/types && wkg publish wavs:types@*.wasm{{config-arg}}
    cd wit-definitions/operator && wkg publish wavs:operator@*.wasm{{config-arg}}
    cd wit-definitions/aggregator && wkg publish wavs:aggregator@*.wasm{{config-arg}}

cargo-check:
    cd packages/wavs-wasi-utils && cargo check --all-targets --all-features

# Update version in all necessary files (eg. just set-tag v0.6.0-alpha.7)
set-tag version:
    #!/usr/bin/env bash
    set -euo pipefail

    # Ensure version doesn't start with 'v' for file updates
    VERSION="{{version}}"
    if [[ "$VERSION" == v* ]]; then
        VERSION="${VERSION#v}"
    fi

    echo "Setting version to: ${VERSION}"

    if command -v gsed >/dev/null 2>&1; then
        SED_CMD="gsed -i"
    else
        SED_CMD="sed -i"
    fi

    # Cargo.toml
    $SED_CMD 's/^version = ".*"/version = "'"${VERSION}"'"/' Cargo.toml

    # all WIT packages
    find wit-definitions -name "*.wit" -type f | while read -r file; do
        $SED_CMD 's/^package wavs:\([^@]*\)@.*/package wavs:\1@'"${VERSION}"';/' "$file"
        $SED_CMD 's|use wavs:types/\([^@]*\)@[^[:space:]]*|use wavs:types/\1@'"${VERSION}"'|g' "$file"
    done

    echo "Version updated to ${VERSION} in all files"

# Create and push git tags for Rust and go module (eg. just push-tag v0.6.0-alpha.7)
push-tag version:
    #!/usr/bin/env bash
    set -euo pipefail

    # Ensure version starts with 'v'
    if [[ "{{version}}" != v* ]]; then
        TAG="v{{version}}"
    else
        TAG="{{version}}"
    fi

    GO_TAG="go/${TAG}"

    echo "Creating tags: ${TAG} and ${GO_TAG}"

    # check if main tag already exists
    if git rev-parse "${TAG}" >/dev/null 2>&1; then
        echo "Error: Tag ${TAG} already exists"
        exit 1
    fi

    # check if go tag already exists
    if git rev-parse "${GO_TAG}" >/dev/null 2>&1; then
        echo "Error: Tag ${GO_TAG} already exists"
        exit 1
    fi

    git tag "${TAG}" -m "Release ${TAG}"
    git tag "${GO_TAG}" -m "Go module release ${TAG}"

    echo "Pushing tags to origin..."
    git push origin "${TAG}"
    git push origin "${GO_TAG}"

    echo "Successfully created and pushed tags: ${TAG} and ${GO_TAG}"
