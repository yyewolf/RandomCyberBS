package values

type Mode string

const (
	Unset Mode = "unset"
	Dev   Mode = "dev"
	Prod  Mode = "prod"
)
