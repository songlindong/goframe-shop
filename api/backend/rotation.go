package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RotationReq struct {
	g.Meta `path:"/rotation" tags:"Rotation" method:"get" summary:"You first rotation api"`
}
type RotationRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
