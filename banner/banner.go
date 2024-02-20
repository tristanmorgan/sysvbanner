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

// Package banner helps you print a banner
package banner

import (
	"fmt"
	"io"
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

func printLine(bs []byte, bsLen int, indexA int, writer io.Writer) {
	line := make([]byte, 80)
	var indexB int

	for indexB = 0; indexB < bsLen; indexB++ {
		ind := int(bs[indexB] - ' ')

		if ind < 0 {
			ind = 0
		}

		for indexC := 0; indexC < 7; indexC++ {
			line[indexB*8+indexC] = glyphs[(ind/8*7)+indexA][(ind%8*7)+indexC]
		}

		line[indexB*8+7] = ' '
	}

	for indexB = bsLen*8 - 1; indexB >= 0; indexB-- {
		if line[indexB] != ' ' {
			break
		}

		line[indexB] = ' '
	}

	str := string(line)
	str = strings.TrimRight(str, "\x00")
	fmt.Fprint(writer, str)
}

// Banner converts little-text to large outputted to a writer
func Banner(str string, writer io.Writer) {
	bs := []byte(str)
	bsLen0 := len(bs)

	for indexA := 0; indexA < 7; indexA++ {
		for offset := 0; offset < bsLen0; offset += 10 {
			bsLen := bsLen0 - offset

			if bsLen > 10 {
				bsLen = 10
			}

			printLine(bs[offset:], bsLen, indexA, writer)
		}

		fmt.Fprintln(writer)
	}

	fmt.Fprintln(writer)
}
