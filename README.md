# GOKU

A golang toolbox by Kumata, learning and accumulation.

## Directory

本项目主要用`internal`目录收集自定义包并存档，在`cmd`项目主干中运行调试，其余目录为出于强迫症而建立的一个项目工程化结构。

```
.
├── api
├── cmd                                  // 1. 包项目主干
│   └── goku
│       ├── goku
│       └── goku.go
├── configs
├── go.mod
├── go.sum
├── gotest
│   ├── try.go
│   └── try_test.go
├── internal                             // 2. GOKU - 本项目收集自定义包的库
│   ├── basics
│   │   ├── array.slice.go               // array and slice analysis
│   │   ├── error.go                     // error example
│   │   └── map.go                       // map example
│   ├── crypto
│   │   ├── aes_example.go               // aes internal function
│   │   ├── README.md
│   │   ├── sha256_reDesign.go           // sha256 internal function, src code rewriting
│   │   ├── sha3_256_example.go          // SHA3-256 and Keccak256 example
│   │   └── sha3_256_useByLinuxCMD.go    // use Keccak256 by Linux command
│   ├── embedded
│   │   └── port.go                      // use serial port
│   └── gokupkg
│       ├── file.go                      // file op
│       ├── flag.demo.go                 // get cmd'parameter and init variable
│       ├── format.go                    // data op
│       ├── http.go                      // httpserver demo
│       ├── linuxCmd.go                  // et cmd'info and return worked return
│       └── timeout.go                   // set loop timeout
├── pkg
├── README.md
└── Reference.md
```

## Usage

```
git clone https://github.com/kumataahh/goku

cd goku/cmd/[XX] && go build
```