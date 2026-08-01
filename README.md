# Minion Language

A fun minion-themed programming language interpreter built in Go! 🍌

## Installation

```bash
bash <(curl -fsSL https://github.com/c00lkiddpostshaxxs/minion/releases/download/v1.0.0/install.sh)
```

## Uninstallation
```bash
bash <(curl -fsSL https://github.com/c00lkiddpostshaxxs/minion/releases/download/v1.0.0/uninstall.sh)
```

## Usage

```bash
./minion <filename.minion>
```
or
```bash
minion <filename.minion>
```

## Keywords

| Minion | English | Usage |
|--------|---------|-------|
| `stampa` | print | `stampa(value)` |
| `despicable` | variable | `despicable x = 5` |
| `poka` | if | `poka x > 5 { ... }` |
| `nadaboba` | else | `} nadaboba { ... }` |
| `whileo` | while | `whileo x < 10 { ... }` |
| `por` | for | `por i = 0 todu 10 { ... }` |
| `todu` | to | (used with `por`) |
| `smash` | break | `smash` |
| `avanti` | continue | `avanti` |
| `workyo` | function | `workyo add(a, b) { ... }` |
| `tornadu` | return | `tornadu a + b` |
| `sidu` | true | `despicable flag = sidu` |
| `nono` | false | `despicable flag = nono` |
| `lista` | list/array | `despicable nums = [1, 2, 3]` |
| `numero` | convert to number | `numero("42")` |
| `stringo` | convert to string | `stringo(100)` |

## Boolean Operators

| Minion | English | Usage |
|--------|---------|-------|
| `e` | and | `poka x > 5 e y < 10 { ... }` |
| `ro` | or | `poka x == 5 ro y == 10 { ... }` |
| `nod` | not | `poka nod flag { ... }` |

## Built-in Functions

| Function | Description | Example |
|----------|-------------|---------|
| `stampa(x)` | Print value | `stampa(42)` |
| `numero(x)` | Convert to number | `numero("42")` |
| `stringo(x)` | Convert to string | `stringo(100)` |
| `lungo(lista)` | Get array length | `lungo([1, 2, 3])` |
| `suma(lista)` | Sum array elements | `suma([1, 2, 3])` |
| `maximo(lista)` | Get max value | `maximo([1, 5, 3])` |

## Comments

Use `#` for comments:

```minion
# This is a comment
despicable x = 5  # inline comment
```

## Examples

### Hello World

```minion
stampa(stringo(42))
```

### Variables and Math

```minion
despicable x = 10
despicable y = 5
stampa(x + y)
stampa(x * y)
```

### Arrays

```minion
despicable nums = [1, 2, 3, 4, 5]
stampa(nums[0])
stampa(lungo(nums))
stampa(suma(nums))
stampa(maximo(nums))
```

### If/Else

```minion
despicable age = 25
poka age >= 18 {
    stampa(1)
} nadaboba {
    stampa(0)
}
```

### While Loop

```minion
despicable i = 0
whileo i < 5 {
    stampa(i)
    despicable i = i + 1
}
```

### For Loop

```minion
por i = 0 todu 5 {
    stampa(i)
}
```

### Break and Continue

```minion
por i = 0 todu 10 {
    poka i == 5 {
        smash
    }
    stampa(i)
}
```

### Functions

```minion
workyo greet(name) {
    tornadu stringo(name)
}

stampa(greet(100))
```

### Boolean Logic

```minion
poka sidu e sidu {
    stampa(1)
}

poka nono ro sidu {
    stampa(2)
}

poka nod nono {
    stampa(3)
}
```

## Data Types

- **Numbers**: integers and floats (64-bit)
- **Strings**: text in quotes `"hello"`
- **Booleans**: `sidu` (true) and `nono` (false)
- **Arrays**: `[1, 2, 3]`

## Operators

### Arithmetic
- `+` Addition
- `-` Subtraction
- `*` Multiplication
- `/` Division

### Comparison
- `>` Greater than
- `<` Less than
- `==` Equal to

### Boolean
- `e` AND
- `ro` OR
- `nod` NOT

## Built by

[c00lkiddpostshaxxs](https://github.com/c00lkiddpostshaxxs) with a passion for minions! 🍌

Enjoy! 🎉
