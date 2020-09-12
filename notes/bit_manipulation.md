# Bit Manipulation

## Bitwise Operations

### Addition

Add the binary numbers as you would decimal numbers, and carry over
the 1s.

```
  11
  0110
+ 0010
------
  1000
```

### Substraction

Subtract normally. If you encounter `0 - 1`, carry a 1 from the next
place and change it to a nominal `2`.

```
  0110
- 0011
------

==>
     2
  0100
- 0011
------
     1

==>
    22
  0000
- 0011
------
  0011
```

### Multiplication

Multiply normally, then carry 1s when adding.

```
    1100
*   1111
--------
   11100
  111000
 1110000
+1100000
--------
10110100
```

### Bitwise AND (`&`)

Performs the AND operation on every single bit.

```
1101 & 1111 = 1101
1101 & 0000 = 0000
```

### Bitwise OR (`|`)

Performs the OR operation on every single bit.

```
1101 | 1111 = 1111
1101 | 0000 = 1101
```

### Bitwise XOR (`^`)

Performs the XOR (exclusive-OR) operation on every single bit.

```
1101 ^ 1111 = 0010
1101 ^ 0000 = 1101
```

### Bit Operations Shortcuts

`~` represents the NOT operation. `1s` represents a string of `1` bits.

```
x ^ 0s = x     |     x & 0s = 0     |     x | 0s = x
x ^ 1s = ~x    |     x & 1s = x     |     x | 1s = 1s
x ^ x  = 0s    |     x & x  = x     |     x | x  = x
```

### Two's Complement

#### Positive to Negative in Two's Complement

1. Flip all bits (XOR with 1s).
2. Prepend a 1 to the left side.

```
011 (3) => 100 + 1 = 101 => (prepend 1) 1101
```
