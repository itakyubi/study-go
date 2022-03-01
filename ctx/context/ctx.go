package context

type Ctx struct {
}

func (ctx *Ctx) Run(f func(inputMsg string) string) {
	// handle msg
	go func() {
		for {
			input := "111"
			output := f(input)
			if output == "" {
				continue
			}
		}
	}()

	// send output to next (only for input and process cell)
	go func() {
		for {
			println("1121")
		}
	}()

}
