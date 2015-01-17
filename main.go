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
	"banner"
	"log"
	"os"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	for _, str := range os.Args[1:] {
		banner.Banner(&str, os.Stdout)
	}
}
