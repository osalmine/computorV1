# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2021/11/03 19:00:13 by osalmine          #+#    #+#              #
#    Updated: 2021/11/23 15:31:03 by osalmine         ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

NAME = computor

all: $(NAME)

$(NAME):
	go build -o $(NAME) cmd/main.go

clean:
fclean:
	rm -f $(NAME)

re: fclean all

test:
	go test ./test/solve -v
	go test ./test/parse -v

.PHONY: all clean fclean re test