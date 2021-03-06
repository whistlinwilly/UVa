// UVa 661 - Blowing Fuses

package main

import (
	"fmt"
	"os"
)

type device struct {
	on          bool
	consumption int
}

func main() {
	in, _ := os.Open("661.in")
	defer in.Close()
	out, _ := os.Create("661.out")
	defer out.Close()

	var n, m, c, tmp, kase int
	for {
		if fmt.Fscanf(in, "%d%d%d", &n, &m, &c); n == 0 && m == 0 && c == 0 {
			break
		}
		devices := make([]device, n)
		for i := range devices {
			fmt.Fscanf(in, "%d", &tmp)
			devices[i] = device{false, tmp}
		}

		var total, max int
		var blown bool
		for i := 0; i < m; i++ {
			fmt.Fscanf(in, "%d", &tmp)
			if !devices[tmp-1].on {
				total += devices[tmp-1].consumption
				if total > c {
					blown = true
					break
				}
				if total > max {
					max = total
				}
			} else {
				total -= devices[tmp-1].consumption
			}
			devices[tmp-1].on = !devices[tmp-1].on
		}

		kase++
		fmt.Fprintf(out, "Sequence %d\n", kase)
		if blown {
			fmt.Fprintln(out, "Fuse was blown.\n")
		} else {
			fmt.Fprintln(out, "Fuse was not blown.")
			fmt.Fprintf(out, "Maximal power consumption was %d amperes.\n\n", max)
		}
	}
}
