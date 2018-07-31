# Sample Binary-Only Packages

Go1.7 から Binary-Only パッケージが導入された. このリポジトリは パッケージをバイナリにしてzipにして配布という一連の流れを試した結果です。

## How to create?
まずは Root に hello.go を作る。このコードは配布されません。

Binary-Only パッケージを作る際には二つのファイルを作る必要があります。特別なコメントを付与したソースコードと、他のバイナリーパッケージです。

特別なコメントに関してはドキュメントで言及されている。

> the package must be distributed with a source file not excluded by build constraints and containing a "//go:binary-only-package" comment.

https://tip.golang.org/pkg/go/build/#hdr-Binary_Only_Packages

つまり`//go:binary-only-package`というコメントをソースコードに残せということらしい。この例は今回 `src/github.com/tcnksm/hello/.` に配置した

まずはバイナリパッケージに build し、`pkg/darwin_amd64/github.com/tcnksm/ directory.` に配置する。

```
$ go build -o pkg/darwin_amd64/github.com/po3rin/hello.a -x
```

そして, `src` をzipし、`pkg` に配置します。

```
$ zip -r hello.zip src/* pkg/*
```

以上で Bainary-Only なパッケージができた。
最終的にはこうなる

```bash
.
├── README.md
├── hello.go
├── hello.zip
├── pkg
│   └── darwin_amd64
│       └── github.com
│           └── po3rin
│               └── hello.a
└── src
    └── github.com
        └── po3rin
            └── hello
                └── hello.go
```

## How to distribute?

このバイナリパッケージを使う際には、ただGOPATHで展開するだけでOK

```bash
$ unzip hello.zip -d $GOPATH/
```

## 参考
https://github.com/tcnksm/go-binary-only-package/blob/master/README.md