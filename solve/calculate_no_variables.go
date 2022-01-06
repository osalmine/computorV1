/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   calculate_no_variables.go                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: osalmine <osalmine@student.hive.fi>        +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/01/06 22:39:13 by osalmine          #+#    #+#             */
/*   Updated: 2022/01/06 22:40:34 by osalmine         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package solve

import "github.com/osalmine/computorV1/utils"

func calculateNoVariables(computor Computor) {
	if computor.Cells == nil {
		utils.PrintOnOption(!computor.Options.Silent, "All real numbers are the solution")
	} else {
		utils.PrintOnOption(!computor.Options.Silent, "No solutions exist")
	}
}
