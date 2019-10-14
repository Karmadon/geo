package geo

import (
	"github.com/sirupsen/logrus"
)

var (
	RtreeMinChildren = 25
	RtreeMaxChildren = 50
	pointLen         = 0.00001 //~1m
)

func NewRtree() *Rtree {
	return NewTree(RtreeMinChildren, RtreeMaxChildren)
}

func (r *Rtree) Intersections(q *Shape) []*Rnode {
	query := getRect(q)
	return r.intersections(query)
}

func (r *Rtree) intersections(q *Rect) []*Rnode {
	inodes := r.SearchIntersect(q)
	nodes := make([]*Rnode, len(inodes))
	for i, inode := range inodes {
		nodes[i] = inode.(*Rnode)
	}
	return nodes
}

func (r *Rtree) Contains(c Coordinate) []*Rnode {
	p := Point{c.X(), c.Y()}
	rect := p.ToRect(pointLen)
	return r.intersections(rect)
}

type Rnode struct {
	shape *Shape
	data  interface{}
}

func NewRnode(shape *Shape, data interface{}) *Rnode {
	return &Rnode{shape: shape, data: data}
}

func (n *Rnode) Feature() *Feature {
	return n.Value().(*Feature)
}

func (n *Rnode) Value() interface{} {
	return n.data
}

//implements rtree.Spatial
func (n *Rnode) Bounds() *Rect {
	return getRect(n.shape)
}

func getRect(s *Shape) *Rect {
	bbox := s.BoundingBox()
	p := Point{bbox.min.X(), bbox.min.Y()}
	d := bbox.max.Difference(bbox.min)
	rect, err := NewRect(p, [2]float64{d.X(), d.Y()})
	if err != nil {
		logrus.Fatal(err)
	}
	return &rect
}
