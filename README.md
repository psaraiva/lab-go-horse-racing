# Lab: Go Corrida de cavalos

## Objetivo
Essa laboratório tem como objetivo demostrar o uso de Goroutines de um jeito simples, prático e divertido.

## Proposta
Cada cavalo na corrida usa uma Goroutine, incrementando o valor de "pontos" a cada turno (sleep). A atualização das posições de cada cavalo é gerado por uma Goroutine. Por fim tem um SELECT que aguarda uma das Goroutine sinalizar que chegou na quantidade limite de pontos por meio de um canal, então o jogo acaba.

## Parâmetros de jogo
- `DELAY_REFRESH_SCREEN` Quantidade de tempo para atualizar a tela do jogo.
- `SCORE_LIMIT` Quantidade máxima de "pontos" para terminar o jogo.
- `STEP_LIMIT` Quantidade máxima de "pontos" que o cavalo anda, quando maior o valor, mais rápido ele pode deslocar.
- `QUANTITY_HORSES` Quantidade de cavalos, mínimo 2 máximo 10.

Obs: Esses parâmetros estão disponíveis em constantes no arquivo `main.go`

## Execução
- `go run main.go`

## Tela
```bash
+----------------------------------------------------+
|.................................................H1 |
|.........................................H2         |
|....................................H3              |
|.............................................H4     |
|............................................H5      |
|..................................................H6|
|................................................H7  |
+----------------------------------------------------+
```
