package golang_united_school_homework

import (
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	
	if len(b.shapes) == b.shapesCapacity {
		return fmt.Errorf("the box has no more capacity")
	}

	b.shapes = append(b.shapes, shape)

	return nil

}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	
	err := b.CheckIndex(i)
	if err != nil {
		return nil, err
	}

	return b.shapes[i], nil

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {

	err := b.CheckIndex(i)
	if err != nil {
		return nil, err
	}

	result := b.shapes[i]

	b.DelFromShapes(i)

	return result, nil

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	
	err := b.CheckIndex(i)
	if err != nil {
		return nil, err
	}

	result := b.shapes[i]

	b.shapes[i] = shape

	return result, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	
	var result float64

	for _, shape := range(b.shapes) {
		result += shape.CalcPerimeter()
	}

	return result

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	
	var result float64

	for _, shape := range(b.shapes) {
		result += shape.CalcArea()
	}

	return result

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {

	ok := b.FindCircle()
	if !ok {
		return fmt.Errorf("there are no circles in slice")
	}

	lastPos := len(b.shapes)-1
	
	for pos := 0; pos <= lastPos; {

		if _, ok := b.shapes[pos].(*Circle); ok {
			b.DelFromShapes(pos)
			lastPos--
			continue
		}
		
		pos++
	}

	return nil

}

func (b *box) CheckIndex(index int) (err error) {

	if index < 0 || index > len(b.shapes)-1 {
		return fmt.Errorf("index is outrange")
	}

	if b.shapes[index] == nil {
		return fmt.Errorf("shape with index %d doesn't exist", index)
	}

	return nil
}

func (b *box) DelFromShapes(index int) () {

	copy(b.shapes[index:], b.shapes[index+1:])
	b.shapes[len(b.shapes)-1] = nil
	b.shapes = b.shapes[:len(b.shapes)-1]

}

func (b *box) FindCircle() (ok bool) {

	for pos := range(b.shapes) {

		if _, ok := b.shapes[pos].(*Circle); ok {
			return true
		}
		
	}

	return
}
