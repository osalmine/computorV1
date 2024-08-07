# computorV1

> This project aims to make you code a simple equation solving program.

ComputorV1 solves equations up to the second degree.

Grade:

[![osalmine's 42 computorv1 Score: 119/100](https://badge42.coday.fr/api/v2/clysm3yyz5816201p4sqs880rw/project/3043747)](https://github.com/Coday-meric/badge42)

**SEE [SUBJECT](en.subject.pdf) FOR MORE DETAILS**

## Usage

This project is written with go, so you need that installed

After cloning, `make` compiles the executable `computorV1`

`make test` runs the tests

`make fclean` removes the executable

`make re` removes the executable and recompiles

`./computor -h` displays the help message

## How it works

The program splits the input string into "cells" which are monomials.

![monomials](https://user-images.githubusercontent.com/54369944/151806024-bcacb37b-f5ee-4cca-be0a-af1934292559.png)

It then combines all cells with the same exponent so that there is only one cell for each exponent.

![exponent-combination](https://user-images.githubusercontent.com/54369944/151806906-fe548ec9-3cc9-489b-bedd-ec33cde582c3.png)

After this, depending on the highest exponent, it solves the equation either with the linear or the quadratic formula.
