package pets

type Pet interface {
	Feed(food string) string
	GivenAttention() string
	IsHungry() bool
}
