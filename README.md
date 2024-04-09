# design_pattern_go
オライリーのデザインパターンをGolangで実装する

## build
```bash
go build -o bin/cmd/design-pattern-go ./cmd/main.go
```

## Usage
|Chapter|パターン名|コマンド|説明|
|--|--|--|--|
|Chapter1|Strategyパターン|`design-pattern-go duck-simulator [duckname1 duckname2 ...]`| duckname には `MalladDuck`, `RubberDuck`, `ExtremeDuck` が入る |
