/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   prettyPrintCells.go                                :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/11/23 13:41:13 by osalmine          #+#    #+#             */
/*   Updated: 2021/11/23 14:10:59 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package utils

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/osalmine/computorV1"
)

func PrettyPrintCells(cells []computorV1.Cell) {
	cellJSON, err := json.MarshalIndent(cells, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Cells: %s\n", string(cellJSON))
}

func PrettyPrintRawCells(cells []computorV1.RawCell) {
	cellJSON, err := json.MarshalIndent(cells, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Raw cells: %s\n", string(cellJSON))
}

func PrettyStringCells(cells []computorV1.Cell) string {
	cellJSON, err := json.MarshalIndent(cells, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	return string(cellJSON)
}
