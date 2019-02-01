// redisclient.go
// Result struct
// Jay Milagroso <jmilagroso@quadx.xyz> / Jan 24 2019

package blueprints

// Result struct
type Result struct {
	Type        string    `json:type`
	Coordinates []float64 `json:coordinates`
}
