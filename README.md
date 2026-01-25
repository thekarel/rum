# Rum

> Is it npm test or was it pnpm test:watch?!

TUI to list, search, run or copy package.json scripts.

Useful in monorepos, in packages with many commands or when jumping between projects.

Will use the correct package manager: npm, pnpm, yarn or bun.

# Usage

```sh
rum # List scripts in the current folder in a TUI
rum ./modules/thing/ # List scripts in the folder
rum /code/project/package.json # List scripts in the package.json file
run -l [optional path] # List scripts and quit (non-interactive)
rum -h # Show help
```

In the TUI

- press `Enter` to run the selected command
- `c` to copy the "run" command to the clipboard (e.g. to copy "npm run lint:fix")
- `C` to copy the actual script to the clipboard (e.g. to copy "oxlint --fix")

The clipboard might not be supported on all platforms.

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

# About

Repo: https://github.com/thekarel/rum

Issues: https://github.com/thekarel/rum/issues

# Gif

![Demo](https://raw.githubusercontent.com/thekarel/rum/main/rum.gif)
