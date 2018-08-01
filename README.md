# Binary-Only Packageの作成から使用

Go1.7 から Binary-Only Package が導入されています。Binary-Only Package はパッケージのコンパイルに使用されるソースコードを含まずにバイナリ形式でパッケージを配布することを可能にするものです。このリポジトリは パッケージをバイナリにしてzipにして配布、そして使うという一連の流れを試した結果です。

## 作り方
まずは Root に hello.go を作る。このコードは実際には配布されません。build時に使うのみです。

```go
package hello

import "fmt"

// Hello - greet
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
```

ここから Binary-Only package を作る際には二つのファイルを作る必要があります。特殊なコメントを付与したソースコードと、他のバイナリーパッケージです。

特殊なコメントに関してはドキュメントで言及されている。

> the package must be distributed with a source file not excluded by build constraints and containing a "//go:binary-only-package" comment.

https://tip.golang.org/pkg/go/build/#hdr-Binary_Only_Packages

つまり`//go:binary-only-package`というコメントをソースコードに残せということらしい。今回の例は `src/github.com/po3rin/hello/hello.go` を下記のように作った。GitHubのアカウント名は適宜書き換えてください(po3rinの部分)

```go
// package hello is sample of binary-only package

//go:binary-only-package

package hello
```

そしてバイナリパッケージに build し、`pkg/darwin_amd64/github.com/po3rin/` に配置する。

```bash
$ go build -o pkg/darwin_amd64/github.com/po3rin/hello.a -x
```

そして, `src` と `pkg` をzipで固めます。

```bash
$ zip -r hello.zip src/* pkg/*
```

以上で Bainary-Only なパッケージができた。最終的には下記のような構成になる。

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

これで hello.zip だけ配ればパッケージとして機能してくれる。

## Bainar-Only Package を使う

このバイナリパッケージを使う際には、ただGOPATHで展開するだけでOK

```bash
$ unzip hello.zip -d $GOPATH/
```

あとは実際にこのパッケージを使った処理を書いてみる

```go
package main

import (
	"fmt"

	"github.com/po3rin/hello"
)

func main() {
	greet := hello.Hello("taro")
	fmt.Println(greet)
}
```

そして動作確認

```go
go run main.go
Hello, taro!
```

## 任意のパッケージが Binary-Only Package かどうかの確認

下記のコードでGOPATH以下にunzipしたパッケージがBinary-Only Packageになっているか確認できます。

```go
package main

import (
	"fmt"
	"go/build"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("error: Getwd methods")
		os.Exit(1)
	}
	path := "github.com/po3rin/hello"
	p, err := build.Import(path, dir, build.ImportComment)
	if err != nil {
		fmt.Println("no target package")
		os.Exit(1)
	}
	if p.BinaryOnly {
		fmt.Println("this package is binary only package")
	}
}
```

`go/build`パッケージのImportでパッケージ情報を所得し、BinaryOnlyフィールドで確認しています。
以上でBinary-Only Packageを作る、使う、確認するの一連の流れができました。
