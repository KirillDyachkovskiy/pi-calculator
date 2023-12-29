# pi-calculator

## Cli

| Name       | Description          |  Type  | Default |
|:-----------|:---------------------|:------:|:-------:|
| iterations | Number of iterations |  uint  | 10_000  |
| threads    | Number of threads    |  uint  |    8    |

## Usage

1) Build a binary file

```bash
go build -o ./build pi-calculator//cmd 
```

2) Calculate the PI

```bash
./build -threads 8 -iterations 1000000
```
