package gen

type Option struct {
	Prefix string `short:"p" description:"Add prefix to generated file name." default:""`
	Suffix string `short:"s" description:"Add suffix to generated file name." default:"_gen"`
}
