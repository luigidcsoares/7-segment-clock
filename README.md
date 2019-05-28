## Relógio com display de 7 segmentos

### Informações Gerais

O desenvolvimento foi realizado em ambiente Linux (Antergos), utilizando os seguintes itens:

* GO 1.12.5;
* Make 4.2.1.

Além disso, para movimentação do cursor no terminal e alteração da cor do
texto, foram utilizados os códigos de escape ANSI.

As cores aceitas até o momento são (entre parênteses está o valor a ser passado na execução):

* Preto (black);
* Vermelho (red);
* Verde (green);
* Amarelo (yellow);
* Azul (blue);
* Magenta (magenta);
* Ciano (cyan);
* Branco (white).

### Execução e Compilação

Para compilar, rode o seguinte comando:

```
$ make go-build
```

Para executar, rode o seguinte comando (a cor é opcional):

```
$ make run [color=<cor>]
```

É possível, também, realizar os dois comandos de uma só vez:

```
$ make [color=<cor>]
```
