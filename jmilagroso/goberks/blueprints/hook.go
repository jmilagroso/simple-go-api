// hook.go
// Hook struct
// Jay Milagroso <jmilagroso@quadx.xyz> / Jan 24 2019

package blueprints

// Hook struct
type Hook struct {
	Command string "json:`command`"
	Group   string "json:`group`"
	Detect  string "json:`detect`"
	Hook    string "json:`hook`"
	Key     string "json:`key`"
	Time    string "json:`time`"
	Id      string "json:`id`"
	Object  Result
}
