package banner_test

import (
	"testing"

	"github.com/tristanmorgan/sysvbanner/banner"
)

type bufferWriter struct {
	buf []string
}

func (writer *bufferWriter) Write(p []byte) (int, error) {
	writer.buf = append(writer.buf, string(p))
	return 0, nil
}

var table = []struct {
	name string
	code []string
}{
	{
		"ABCDEFGH", []string{
			" ###  ####   ###  ####  ##### #####  #### #   # \n",
			"#   # #   # #   # #   # #     #     #     #   # \n",
			"##### ####  #     #   # ###   ###   #  ## ##### \n",
			"#   # #   # #     #   # #     #     #   # #   # \n",
			"#   # #   # #     #   # #     #     #   # #   # \n",
			"#   # #   # #   # #   # #     #     #   # #   # \n",
			"#   # ####   ###  ####  ##### #      ###  #   # \n",
			"                                                \n",
			"\n",
		},
	}, {
		"IJKLMNOP", []string{
			" ###      # #   # #     #   # #   #  ###  ####  \n",
			"  #       # #  #  #     ## ## ##  # #   # #   # \n",
			"  #       # ###   #     # # # # # # #   # ####  \n",
			"  #       # #  #  #     #   # #  ## #   # #     \n",
			"  #       # #   # #     #   # #   # #   # #     \n",
			"  #   #   # #   # #     #   # #   # #   # #     \n",
			" ###   ###  #   # ##### #   # #   #  ###  #     \n",
			"                                                \n",
			"\n",
		},
	}, {
		"QRSTUVWX", []string{
			" ###  ####   #### ##### #   # #   # #   # #   # \n",
			"#   # #   # #       #   #   # #   # #   #  # #  \n",
			"#   # ####   ###    #   #   # #   # #   #   #   \n",
			"#   # #   #     #   #   #   # #   # #   #  # #  \n",
			"# # # #   #     #   #   #   #  # #  # # # #   # \n",
			"#  #  #   # #   #   #   #   #  # #  ## ## #   # \n",
			" ## # #   #  ###    #    ###    #   #   # #   # \n",
			"                                                \n",
			"\n",
		},
	}, {
		"YZabcdef", []string{
			"#   # #####       #               #          ## \n",
			" # #      #       #               #         #   \n",
			"  #      #   ###  # ##   ###   ## #  ###   #### \n",
			"  #     #       # ##  # #   # #  ## #   #   #   \n",
			"  #    #     #### #   # #     #   # #####   #   \n",
			"  #   #     #   # #   # #   # #   # #       #   \n",
			"  #   #####  #### ####   ###   ####  ####   #   \n",
			"                                                \n",
			"\n",
		},
	}, {
		"ghijklmn", []string{
			"      #       #       # #      ##               \n",
			"      #                 #       #               \n",
			" #### # ##   ##       # #  #    #   ## #  ####  \n",
			"#   # ##  #   #       # # #     #   # # # #   # \n",
			"#   # #   #   #       # ##      #   # # # #   # \n",
			" #### #   #   #   #   # # #     #   #   # #   # \n",
			"    # #   #    ## #   # #  #     ## #   # #   # \n",
			"####               ###                          \n",
			"\n",
		},
	}, {
		"opqrstuv", []string{
			"                                                \n",
			"                                #               \n",
			" ###  # ##   ## # # ##   ####  ###  #   # #   # \n",
			"#   # ##  # #  ## ##  # #       #   #   # #   # \n",
			"#   # #   # #   # #      ###    #   #   # #   # \n",
			"#   # ####   #### #         #   #   #   #  # #  \n",
			" ###  #         # #     ####     ##  ####   #   \n",
			"      #         #                               \n",
			"\n",
		},
	}, {
		"wxyz0123", []string{
			"                         ###    #    ###   ###  \n",
			"                        #   #  ##   #   # #   # \n",
			"#   # #   # #   # ##### #  ##   #       #     # \n",
			"#   #  # #  #   #    #  # # #   #     ##    ##  \n",
			"# # #   #   #   #   #   ##  #   #    #        # \n",
			"# # #  # #   ####  #    #   #   #   #   # #   # \n",
			" #### #   #     # #####  ###  ##### #####  ###  \n",
			"            ####                                \n",
			"\n",
		},
	},
}

func TestBanner(t *testing.T) {
	for _, v := range table {
		t.Run(v.name, func(t *testing.T) {
			writer := &bufferWriter{}
			banner.Banner(v.name, writer)

			for i, actual := range writer.buf {
				if v.code[i] != actual {
					t.Errorf("expected: '%v'\ngot: '%v'\n", v.code[i], actual)
				}
			}
		})
	}
}
