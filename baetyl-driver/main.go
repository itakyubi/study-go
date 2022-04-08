package main

import dm "github.com/baetyl/baetyl-go/v2/dmcontext"

func main() {
	dm.Run(func(ctx dm.Context) error {
		d, err := newDriver(ctx)
		if err != nil {
			return err
		}
		d.start()
		defer d.stop()
		ctx.Wait()
		return nil
	})
}
