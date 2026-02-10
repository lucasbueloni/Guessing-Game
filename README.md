# Guessing Game em Go

Este foi o meu projeto de estreia no ecossistema **Go**, desenvolvido para consolidar o aprendizado da sintaxe e da biblioteca padrão da linguagem. É um jogo de adivinhação simples e interativo via terminal.

---

## O Jogo
O objetivo é descobrir um número secreto entre **0 e 100**.
- Você tem um limite de **10 tentativas**.
- O programa fornece feedback em tempo real (se o número secreto é maior ou menor que o seu palpite).
- Ao final, o jogo exibe o histórico de todos os números que você chutou.

---

## O que foi utilizado

Neste projeto, apliquei conceitos fundamentais de Go que demonstram a robustez da linguagem:

* **Manipulação de I/O com Buffer:** Uso do pacote `bufio` e `os.Stdin` para uma leitura de entrada mais eficiente no terminal.
* **Geração Aleatória Moderna:** Implementação utilizando o pacote `math/rand/v2` para gerar o número secreto.
* **Tratamento de Strings e Conversão:** Uso de `strings.TrimSpace` para limpar inputs e `strconv.ParseInt` para converter texto em números, incluindo tratamento básico de erros.
* **Estruturas de Dados:** Uso de **Arrays** e **Slices** para armazenar e exibir o histórico de chutes do jogador.
* **Controle de Fluxo:** Lógica estruturada com `for range` e `switch` cases limpos.

---

## Como testar

1. Certifique-se de ter o **Go 1.22+** instalado (devido ao uso de `rand/v2`).
2. Clone este repositório:
3. Execute o comando: go run main.go
