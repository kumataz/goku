# GOKU

A golang toolbox by Kumata, learning and accumulation.

## Directory

本项目除了记录学习研究内容以外，主要用`internal`目录收集自用的包作为kit-tool存档，在`cmd`项目主干中运行调试，其余目录为出于强迫症而建立的一个标准项目工程化结构。

```
.
├── api
├── cmd                                  // 包项目主干
│   └── goku
│       ├── goku
│       └── goku.go
├── configs
├── go.mod
├── go.sum
├── gotest                               // test目录
├── internal                             // 本项目自定义包的库
├── pkg                                  // 通用的外部引入的库
├── README.md
└── Reference.md
```

## Usage

```
git clone https://github.com/kumataahh/goku

cd goku/cmd/[XX] && go build
```