package blueprints

type Response struct {
	Id      string `json:id`
	Channel string `json:channel`
	Result  Result
}
