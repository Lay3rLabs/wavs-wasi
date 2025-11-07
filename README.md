# wavs-wasi
wasi-utils and wit files for easy MIT-license import

## Release Workflow

This project uses [just](https://github.com/casey/just) for managing releases. Here's the typical workflow:

1. **Make your changes** - Develop features, fix bugs, etc.
2. **Set the version** - When ready for release, update version numbers across all files:
   ```bash
   just set-version v0.6.0-alpha.7
   ```
3. **Create PR and merge** - Get your changes reviewed and merged to main
4. **Push tags** - After merge, create and push git tags:
   ```bash
   just push-tag v0.6.0-alpha.7
   ```

This will automatically update version numbers in `Cargo.toml` and all WIT package definitions, then create both standard (`v0.6.0-alpha.7`) and Go module (`go/v0.6.0-alpha.7`) tags.
