# Math Library

A comprehensive Go math library that provides various mathematical operations and utilities.

## Features

- Basic arithmetic operations (addition, subtraction, multiplication, division)
- Advanced mathematical functions (power, square root, logarithm)
- Statistical functions (mean, median, standard deviation)
- Geometric calculations (circle area, rectangle perimeter)
- Utility functions (rounding, degree/radian conversion, prime checking, GCD)

## Installation

```bash
go get github.com/pynewhere/math-lib
```

## Usage

```go
import "github.com/pynewhere/math-lib"

// Basic operations
result := basic.Add(10, 5)

// Advanced operations
power := advanced.Pow(2, 3)

// Statistical operations
mean := statistics.Mean([]float64{1, 2, 3, 4, 5})

// Geometric operations
area := geometry.CircleArea(5)

// Utility functions
rounded := utils.Round(3.14159, 2)
```

## Modules

- `basic`: Basic arithmetic operations
- `advanced`: Advanced mathematical functions
- `statistics`: Statistical calculations
- `geometry`: Geometric calculations
- `utils`: Utility functions

## License

MIT License 