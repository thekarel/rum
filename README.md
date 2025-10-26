# Rum

> Is it npm test or was it pnpm test:watch?!

TUI to list, filter and run package.json scripts.

Useful in monorepos, in packages with many commands or when jumping between projects.

Will use the correct package manager: npm, pnpm, yarn or bun.

```sh
rum --help
TUI to list, filter and run package.json scripts.

To list the scripts in the current folder:
  rum

You can also pass relative or absolute paths either to a folder or a file:
  rum ./modules/thing/
  rum /code/project/package.json

Usage:
  rum <path to folder or package.json> [flags]
```

# Install

[![Attestations](https://img.shields.io/badge/Attestations-00aa00)](https://github.com/thekarel/rum/attestations)

```sh
# If you have npm :-)
npm i -g @thekarel/rum

# If you have npm and love delays
npx @thekarel/rum

# If you have go https://go.dev/doc/install
go install github.com/thekarel/rum@latest

# if you have eget https://github.com/zyedidia/eget
eget thekarel/rum
```

# Gif

![Demo](https://raw.githubusercontent.com/thekarel/rum/main/rum.gif)
