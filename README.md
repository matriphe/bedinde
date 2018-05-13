# Bedinde

[![Build Status](https://travis-ci.org/matriphe/bedinde.svg?branch=master)](https://travis-ci.org/matriphe/bedinde)

**Bedinde** is Go package that provide basic formatting and function.

[The word *bedinde* in Indonesian means *helper*](https://kbbi.kemdikbud.go.id/entri/bedinde), so basically this is a helpers package.

## Installation

```shell
go get -u github.com/matriphe/bedinde
```

## Function Lists

### InArray

Check if value is in array.

```go
a := []string{"a", "b", "c", "d"}

exists, index := bedinde.InArray("c", a) // return true, 2
exists, index := bedinde.InArray("e", a) // return false, -1
```

### FormatNumber

Format number using separator.

```go
r := bedinde.FormatNumber(12345, "id") // return "12.345"
r := bedinde.FormatNumber(12345, "en") // return "12,345"
```

### FormatPhone

Format phone number to E164 standard.

```go
r := bedinde.FormatPhone("081823456789", "id") // return "+6281823456789"
r := bedinde.FormatPhone("081823456789", "sg") // return "+65081823456789"
```

### FormatTimeDiff

Format time difference to its components.

```go
d, _ := time.ParseDuration("2h3m4s")
r := bedinde.FormatTimeDiff(d) // return TimeDiff{Day: 0, Hour: 2, Minute: 3, Second: 4}
```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.