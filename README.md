# hashdir

Calculate the SHA256 hashes of all files in a directory tree.

```
hashdir

Calculate the hash of all files in a directory tree. Empty directories are ignored.

Usage:

  hashdir

  hashdir ./dir
```

## Output

```
5656fafa00d4f294bcb606cf4f7d4fa877390e46f583e8b3c8744ace104a31d1 .gitignore
53c639ce9b5d2cb6f5dc334013fb62be8e3f7d86f406b2dfcd1695450cc73390 README.md
c20805e8aa8bfc24e48114a9a3937a0b567f6e886545a5a11369fcd419475795 flake.lock
728a84261c5519cb8ebe314de73eb17c90abf6338edf92c33018e2106c501873 flake.nix
4b0f39191c244ae63240e9787d967e2212a0a0af931f33e81a62a39d6ad386cc go.mod
eced0644e141c711d16133ad3be57b65f0c9f5d8a7ef216a9f24910b2754e3d3 main.go
4757cbc58eaebf7347f586f7e3e27e9f228dad66c29a46d4a5f928c61560b820 .

4757cbc58eaebf7347f586f7e3e27e9f228dad66c29a46d4a5f928c61560b820 .
```
