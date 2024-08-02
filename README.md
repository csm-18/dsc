# dsc

_dsc_ short for Desktop Shortcut Creator (for Linux Mint)

dsc makes shortcut creation easy!

## Requirements:

1. os: Linux Mint.
2. golang version: go1.22.5 or above.

## To Build The Project:

1. clone this repo.
2. cd into dsc directory.
3. run: go build
4. dsc binary is created in the root directory(dsc/).
5. place the binary at a place of your choice.
6. add dsc binary path to PATH environment variable: \
   Add this line to .bashrc file:
   ```bash
   export PATH=$PATH:<full-path-to-dsc-binary>
   ```
7. open a new terminal window and run:

```bash
dsc version
```

8. on how to use, run:

```bash
dsc help
```
