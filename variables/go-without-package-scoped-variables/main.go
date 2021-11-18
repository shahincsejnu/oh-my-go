// Followed tuto : https://dave.cheney.net/2017/06/11/go-without-package-scoped-variables

package main

func main() {
	var _ int = 10
	var _ int = *new(int)

}
