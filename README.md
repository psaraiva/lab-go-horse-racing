[![Go Report Card](https://goreportcard.com/badge/github.com/psaraiva/lab-go-horse-racing)](https://goreportcard.com/report/github.com/psaraiva/lab-go-horse-racing)
![Codecov](https://img.shields.io/codecov/c/github/psaraiva/lab-go-horse-racing)

# Lab: Go Corrida de cavalos

## Objetivo
Essa laboratório tem como objetivo demostrar o uso de Goroutines de um jeito simples, prático e divertido.

## Como isso funciona?
Cada cavalo na corrida usa uma Goroutine, incrementando o valor de "pontos" a cada turno (sleep). A atualização das posições de cada cavalo é gerado por uma Goroutine. Por fim tem um SELECT aguarda uma das Goroutine sinalizar que chegou na pontuação alvo por meio de um canal, então o jogo acaba.

## Parâmetros de jogo
- `DELAY_HORSE_STEP` Quantidade de tempo que o cavalo espera a cada ciclo.
- `DELAY_REFRESH_SCREEN` Quantidade de tempo para atualizar a tela do jogo.
- `HORSE_LABEL` Rótulo de identificação do cavalo.
- `HORSE_MAX_SPEED` Velocidade máxima do cavalo.
- `HORSE_MIN_SPEED` Velocidade mínima do cavalo.
- `HORSES_QUANTITY` Quantidade de cavalos no jogo.
- `HORSES_MAX_QUANTITY` Quantidade máxima de cavalos.
- `SCORE_TARGET` Pontuação necessária para ganhar o jogo.
- `TIMEOUT_GAME` Tempo limite para acabar o jogo.

Obs: Esses parâmetros estão disponíveis em constantes no arquivo `main.go`

## Execução
- `go run main.go`

## Tela
```bash
   +-----------------------------------------------------------------------------+
H01|......................................................................H01    |
H02|..........................................................................H02|
H03|.................................................................H03         |
H04|...........................................................................H04|
H05|....................................................................H05      |
H06|.................................................................H06         |
H07|..............................................................................H07|
H08|....................................................................H08      |
H09|.......................................................................H09   |
H10|........................................................................H10  |
   +-----------------------------------------------------------------------------+
The horse winner is: H07 - Score 78
```
