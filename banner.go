// The following is the original copyright:
// https://github.com/uffejakobsen/sysvbanner
// ------------------------------------------------------------------
/*****************************************************************
 *
 * SYSVbanner.c
 *
 * This is a PD version of the SYS V banner program (at least I think
 * it is compatible to SYS V) which I wrote to use with the clock
 * program written by:
 **     DCF, Inc.
 **     14623 North 49th Place
 **     Scottsdale, AZ 85254
 * and published in the net comp.sources.misc newsgroup in early July
 * since the BSD banner program works quite differently.
 *
 * There is no copyright or responsibility accepted for the use
 * of this software.
 *
 * Brian Wallis, brw@jim.odr.oz, 4 July 1988
 *
 *****************************************************************/

/* Changes by David Frey, david@eos.lugs.ch, 3 February 1997:
 * 1. protoized and indented, 2. changed @ character to #
 */
// ------------------------------------------------------------------
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var glyphs = []string{
	"         ###  ### ###  # #   ##### ###   #  ##     ###  ",
	"         ###  ### ###  # #  #  #  ## #  #  #  #    ###   ",
	"         ###   #   # ########  #   ### #    ##      #   ",
	"          #            # #   #####    #    ###     #    ",
	"                     #######   #  #  # ####   # #       ",
	"         ###           # #  #  #  # #  # ##    #        ",
	"         ###           # #   ##### #   ### #### #       ",

	"   ##    ##                                            #",
	"  #        #   #   #    #                             # ",
	" #          #   # #     #                            #  ",
	" #          # ### ### #####   ###   #####           #   ",
	" #          #   # #     #     ###           ###    #    ",
	"  #        #   #   #    #      #            ###   #     ",
	"   ##    ##                   #             ###  #      ",

	"  ###     #    #####  ##### #      ####### ##### #######",
	" #   #   ##   #     ##     ##    # #      #     ##    # ",
	"# #   # # #         #      ##    # #      #          #  ",
	"#  #  #   #    #####  ##### ####### ##### ######    #   ",
	"#   # #   #   #            #     #       ##     #  #    ",
	" #   #    #   #      #     #     # #     ##     #  #    ",
	"  ###   ##### ####### #####      #  #####  #####   #    ",

	" #####  #####    #     ###      #           #     ##### ",
	"#     ##     #  # #    ###     #             #   #     #",
	"#     ##     #   #            #     #####     #        #",
	" #####  ######         ###   #                 #     ## ",
	"#     #      #   #     ###    #     #####     #     #   ",
	"#     ##     #  # #     #      #             #          ",
	" #####  #####    #     #        #           #       #   ",

	" #####    #   ######  ##### ###### ############## ##### ",
	"#     #  # #  #     ##     ##     ##      #      #     #",
	"# ### # #   # #     ##      #     ##      #      #      ",
	"# # # ##     ####### #      #     ######  #####  #  ####",
	"# #### ########     ##      #     ##      #      #     #",
	"#     ##     ##     ##     ##     ##      #      #     #",
	" ##### #     #######  ##### ###### ########       ##### ",

	"#     #  ###        ##    # #      #     ##     ########",
	"#     #   #         ##   #  #      ##   ####    ##     #",
	"#     #   #         ##  #   #      # # # ## #   ##     #",
	"#######   #         ####    #      #  #  ##  #  ##     #",
	"#     #   #   #     ##  #   #      #     ##   # ##     #",
	"#     #   #   #     ##   #  #      #     ##    ###     #",
	"#     #  ###   ##### #    # ########     ##     ########",

	"######  ##### ######  ##### ########     ##     ##     #",
	"#     ##     ##     ##     #   #   #     ##     ##  #  #",
	"#     ##     ##     ##         #   #     ##     ##  #  #",
	"###### #     #######  #####    #   #     ##     ##  #  #",
	"#      #   # ##   #        #   #   #     # #   # #  #  #",
	"#      #    # #    # #     #   #   #     #  # #  #  #  #",
	"#       #### ##     # #####    #    #####    #    ## ## ",

	"#     ##     ######## ##### #       #####    #          ",
	" #   #  #   #      #  #      #          #   # #         ",
	"  # #    # #      #   #       #         #  #   #        ",
	"   #      #      #    #        #        #               ",
	"  # #     #     #     #         #       #               ",
	" #   #    #    #      #          #      #               ",
	"#     #   #   ####### #####       # #####        #######",

	"  ###                                                   ",
	"  ###     ##   #####   ####  #####  ###### ######  #### ",
	"   #     #  #  #    # #    # #    # #      #      #    #",
	"    #   #    # #####  #      #    # #####  #####  #     ",
	"        ###### #    # #      #    # #      #      #  ###",
	"        #    # #    # #    # #    # #      #      #    #",
	"        #    # #####   ####  #####  ###### #       #### ",

	"                                                        ",
	" #    #    #        # #    # #      #    # #    #  #### ",
	" #    #    #        # #   #  #      ##  ## ##   # #    #",
	" ######    #        # ####   #      # ## # # #  # #    #",
	" #    #    #        # #  #   #      #    # #  # # #    #",
	" #    #    #   #    # #   #  #      #    # #   ## #    #",
	" #    #    #    ####  #    # ###### #    # #    #  #### ",

	"                                                        ",
	" #####   ####  #####   ####   ##### #    # #    # #    #",
	" #    # #    # #    # #         #   #    # #    # #    #",
	" #    # #    # #    #  ####     #   #    # #    # #    #",
	" #####  #  # # #####       #    #   #    # #    # # ## #",
	" #      #   #  #   #  #    #    #   #    #  #  #  ##  ##",
	" #       ### # #    #  ####     #    ####    ##   #    #",

	"                       ###     #     ###   ##    # # # #",
	" #    #  #   # ###### #        #        # #  #  # # # # ",
	"  #  #    # #      #  #        #        #     ## # # # #",
	"   ##      #      #  ##                 ##        # # # ",
	"   ##      #     #    #        #        #        # # # #",
	"  #  #     #    #     #        #        #         # # # ",
	" #    #    #   ######  ###     #     ###         # # # #",
}

func printLine(bs []byte, bs_len int, index_a int) {
	line := make([]byte, 80)
	var index_b int

	for index_b = 0; index_b < bs_len; index_b++ {
		ind := int(bs[index_b] - ' ')

		if ind < 0 {
			ind = 0
		}

		for index_c := 0; index_c < 7; index_c++ {
			line[index_b*8+index_c] = glyphs[(ind/8*7)+index_a][(ind%8*7)+index_c]
		}

		line[index_b*8+7] = ' '
	}

	for index_b = bs_len*8 - 1; index_b >= 0; index_b-- {
		if line[index_b] != ' ' {
			break
		}

		line[index_b] = ' '
	}

	str := string(line)
	str = strings.TrimRight(str, "\x00")
	fmt.Print(str)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	for _, str := range os.Args[1:] {
		bs := []byte(str)
		bs_len0 := len(bs)

		for index_a := 0; index_a < 7; index_a++ {
			for offset := 0; offset < bs_len0; offset += 10 {
				bs_len := bs_len0 - offset

				if bs_len > 10 {
					bs_len = 10
				}

				printLine(bs[offset:], bs_len, index_a)
			}

			fmt.Println()
		}
	}
}
