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

var glyphs = [][][]byte{
	{
		[]byte("       #   # #  # #   #  #   #  #    #  "),
		[]byte("       #   # #  # #  #####  #  # #   #  "),
		[]byte("       #       ######       #   #       "),
		[]byte("       #        # #  ###   #   ## #     "),
		[]byte("       #       #####    # #   # ##      "),
		[]byte("                # # ####  #  ##  #      "),
		[]byte("       #        # #   #  #   # ## #     "),
		[]byte("                                        "),
	},
	{
		[]byte("   #  #    # #                         #"),
		[]byte("  #    #    #    #                    # "),
		[]byte(" #      #  # #   #                    # "),
		[]byte(" #      #      #####     #####       #  "),
		[]byte(" #      #        #                  #   "),
		[]byte("  #    #         #     #            #   "),
		[]byte("   #  #               #         #  #    "),
		[]byte("                                        "),
	},
	{
		[]byte(" ###   #   ###  ###    #######  ## #####"),
		[]byte("#   # ##  #   ##   #  # ##     #   #   #"),
		[]byte("#  ##  #      #    # #  ##### #        #"),
		[]byte("# # #  #    ##   ## #   #    #####    # "),
		[]byte("##  #  #   #       ######    ##   #  #  "),
		[]byte("#   #  #  #   ##   #    ##   ##   #  #  "),
		[]byte(" ### ########## ###     # ###  ###   #  "),
		[]byte("                                        "),
	},
	{
		[]byte(" ###  ###              #      #     ### "),
		[]byte("#   ##   #            #        #   #   #"),
		[]byte("#   ##   #  #    #   #   #####  #      #"),
		[]byte(" ###  ####          #            #    # "),
		[]byte("#   #    #           #   #####  #    #  "),
		[]byte("#   #   #             #        #        "),
		[]byte(" ###  ##    #    #     #      #      #  "),
		[]byte("                #                       "),
	},
	{
		[]byte("      ### ####  ### #### ########## ####"),
		[]byte(" ### #   ##   ##   ##   ##    #    #    "),
		[]byte("#   ########## #    #   ####  ###  #  ##"),
		[]byte("# # ##   ##   ##    #   ##    #    #   #"),
		[]byte("# # ##   ##   ##    #   ##    #    #   #"),
		[]byte("# ####   ##   ##   ##   ##    #    #   #"),
		[]byte("#    #   #####  ### #### ######     ### "),
		[]byte(" ####                                   "),
	},
	{
		[]byte("#   # ###     ##   ##    #   ##   # ### "),
		[]byte("#   #  #      ##  # #    ## ####  ##   #"),
		[]byte("#####  #      ####  #    # # ## # ##   #"),
		[]byte("#   #  #      ##  # #    #   ##  ###   #"),
		[]byte("#   #  #      ##   ##    #   ##   ##   #"),
		[]byte("#   #  #  #   ##   ##    #   ##   ##   #"),
		[]byte("#   # ###  ### #   #######   ##   # ### "),
		[]byte("                                        "),
	},
	{
		[]byte("####  ### ####  ##########   ##   ##   #"),
		[]byte("#   ##   ##   ##      #  #   ##   ##   #"),
		[]byte("#### #   #####  ###   #  #   ##   ##   #"),
		[]byte("#    #   ##   #    #  #  #   ##   ##   #"),
		[]byte("#    # # ##   #    #  #  #   # # # # # #"),
		[]byte("#    #  # #   ##   #  #  #   # # # ## ##"),
		[]byte("#     ## ##   # ###   #   ###   #  #   #"),
		[]byte("                                        "),
	},
	{
		[]byte("#   ##   ###### ### #     ###   #       "),
		[]byte(" # #  # #     # #    #      #  # #      "),
		[]byte("  #    #     #  #    #      # #   #     "),
		[]byte(" # #   #    #   #     #     #           "),
		[]byte("#   #  #   #    #      #    #           "),
		[]byte("#   #  #  #     #      #    #           "),
		[]byte("#   #  #  ##### ###     # ###      #####"),
		[]byte("                                        "),
	},
	{
		[]byte(" #        #             #        ##     "),
		[]byte("  #       #             #       #       "),
		[]byte("      ### # ##  ###  ## # ###  #### ####"),
		[]byte("         ###  ##   ##  ###   #  #  #   #"),
		[]byte("      #####   ##    #   ######  #  #   #"),
		[]byte("     #   ##   ##   ##   ##      #   ####"),
		[]byte("      ########  ###  #### ####  #      #"),
		[]byte("                                   #### "),
	},
	{
		[]byte("#      #      ##     ##                 "),
		[]byte("#              #      #                 "),
		[]byte("# ##  ##      ##  #   #  ## # ####  ### "),
		[]byte("##  #  #      ## #    #  # # ##   ##   #"),
		[]byte("#   #  #      ###     #  # # ##   ##   #"),
		[]byte("#   #  #  #   ## #    #  #   ##   ##   #"),
		[]byte("#   #   ###   ##  #    ###   ##   # ### "),
		[]byte("           ###                          "),
	},
	{
		[]byte("                                        "),
		[]byte("                      #                 "),
		[]byte("# ##  ## ## ##  #### ### #   ##   ##   #"),
		[]byte("##  ##  ####  ##      #  #   ##   ##   #"),
		[]byte("#   ##   ##     ###   #  #   ##   ## # #"),
		[]byte("####  #####        #  #  #   # # # # # #"),
		[]byte("#        ##    ####    ## ####  #   ####"),
		[]byte("#        #                              "),
	},
	{
		[]byte("                  #   #   #        # # #"),
		[]byte("                 #    #    #        # # "),
		[]byte("#   ##   ######  #    #    #   #  ## # #"),
		[]byte(" # # #   #   # ##     #     ### # # # # "),
		[]byte("  #  #   #  #    #    #    #  #  # # # #"),
		[]byte(" # #  #### #     #    #    #        # # "),
		[]byte("#   #    ######   #   #   #        # # #"),
		[]byte("     ####                               "),
	},
}

func printLine(bs []byte, indexA int, writer io.Writer) {
	line := make([]byte, 0)

	for _, chr := range bs {

		ind := int(chr - ' ')

		if ind < 0 {
			ind = 0
		}

		line = append(line, glyphs[(ind / 8)][indexA][(ind%8*5):((ind%8*5)+5)]...)

		line = append(line, ' ')
	}

	str := string(line)
	str = strings.TrimRight(str, "\x00")
	fmt.Fprintln(writer, str)
}

// Banner converts little-text to large outputted to a writer
func Banner(str string, writer io.Writer) {
	bs := []byte(str)

	for indexA := 0; indexA < 8; indexA++ {
		printLine(bs, indexA, writer)
	}

	fmt.Fprintln(writer)
}
