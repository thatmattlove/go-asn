<div align="center">
<h3>
    <code>go-asn</code>
</h3>
<br/>
Autonomous System Number Utility for Go
<br/>
<br/>
    <a href="https://github.com/thatmattlove/go-asn/actions/workflows/test.yml">
        <img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/thatmattlove/go-asn/test.yml?style=for-the-badge">
    </a>
    <a href="https://app.codecov.io/gh/thatmattlove/go-asn">
        <img alt="Codecov" src="https://img.shields.io/codecov/c/github/thatmattlove/go-asn?style=for-the-badge">
    </a>
    <a href="https://github.com/thatmattlove/go-asn/releases">
        <img alt="GitHub release (latest SemVer)" src="https://img.shields.io/github/v/release/thatmattlove/go-asn?label=version&style=for-the-badge">
    </a>

</div>

## Installation

```console
go get github.com/thatmattlove/go-asn
```

## Usage

### Parsing

```go
a, err := asn.Parse("65000")
a, err := asn.FromDecimal("65000")
a, err := asn.FromASDot("64086.59904")
a, err := asn.FromUint64(65001)
a, err := asn.FromBytes(255, 255, 253, 232)
a := asn.From4Bytes(255, 255, 253, 232)
a := asn.From2Bytes(253, 232)
a := asn.FromUint32(65001)
a := asn.MustParse("65000")
a := asn.MustDecimal("65000")
a := asn.MustASDot("0.65000")
```

### Formatting

```go
a := asn.MustParse("65000")
a.Size()
// 2
a.ASPlain()
// 65000
a.ASDot()
// 65000
a.ASDotPlus()
// 0.65000
a.String()
// 65000
a.ByteString()
// {0,0,253,232}

a = asn.MustParse("4200000000")
a.Size()
// 4
a.ASPlain()
// 4200000000
a.ASDot()
// 64086.59904
a.ASDotPlus()
// 64086.59904
a.String()
// 4200000000
a.ByteString()
// {250,86,234,0}
```

### Comparison

```go
a := asn.MustParse("65000")
b := asn.MustParse("65001")
c := asn.MustParse("65002")
d := asn.MustParse("65000")
e := asn.MustParse("64512")
a.Equal(b)
// false
a.Equal(d)
// true
a.LessThan(b)
// true
a.LEqual(c)
// true
a.GreaterThan(e)
// true
a.GEqual(e)
// true
```

### Iteration

```go
start := asn.MustParse("65000")
end := asn.MustParse("65005")

for iter := start.Range(end); iter.Continue(); {
    next := iter.Next()
    fmt.Println(next.ASPlain())
}
// 65001
// 65002
// 65003
// 65004
// 65505

a := asn.MustParse("65000")
for iter := a.Iter(); iter.Continue(); {
    next := iter.Next()
    fmt.Println(next.ASPlain())
}
// 65001
// 65002
// ...
// 4294967294
```

![GitHub](https://img.shields.io/github/license/thatmattlove/go-asn?style=for-the-badge&color=black)