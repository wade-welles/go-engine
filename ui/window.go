package ui

import (
	"image/color"

	"github.com/walesey/go-engine/renderer"
	vmath "github.com/walesey/go-engine/vectormath"
)

type Window struct {
	node, elementNode, background *renderer.Node
	backgroundBox                 *renderer.Geometry
	element                       Element
	size, position                vmath.Vector2
	mousePos                      vmath.Vector2
}

func (w *Window) Draw(renderer renderer.Renderer) {
	w.node.Draw(renderer)
}

func (w *Window) Destroy(renderer renderer.Renderer) {
	w.node.Destroy(renderer)
}

func (w *Window) Centre() vmath.Vector3 {
	return w.node.Centre()
}

func (w *Window) Optimize(geometry *renderer.Geometry, transform renderer.Transform) {
	w.node.Optimize(geometry, transform)
}

func (w *Window) SetScale(scale vmath.Vector3) {
	w.background.SetScale(scale)
	w.size = vmath.Vector2{scale.X, scale.Y}
	w.Render()
}

func (w *Window) SetTranslation(translation vmath.Vector3) {
	w.node.SetTranslation(translation)
	w.position = translation.ToVector2()
}

func (w *Window) SetOrientation(orientation vmath.Quaternion) {
	w.node.SetOrientation(orientation)
}

func (w *Window) SetBackgroundColor(r, g, b, a uint8) {
	w.backgroundBox.SetColor(color.NRGBA{r, g, b, a})
}

func (w *Window) SetElement(element Element) {
	if w.element != nil {
		w.elementNode.Remove(w.element.Spatial(), true)
	}
	w.element = element
	w.elementNode.Add(element.Spatial())
	w.Render()
}

func (w *Window) Render() {
	if w.element != nil {
		size := w.element.Render(w.size, vmath.Vector2{0, 0})
		if size.X > w.size.X {
			w.size.X = size.X
		}
		if size.Y > w.size.Y {
			w.size.Y = size.Y
		}
		w.background.SetScale(w.size.ToVector3())
	}
}

func (w *Window) ElementById(id string) Element {
	if w.element != nil {
		if w.element.GetId() == id {
			return w.element
		}
		container, ok := w.element.(*Container)
		if ok {
			return container.ElementById(id)
		}
	}
	return nil
}

func (w *Window) mouseMove(position vmath.Vector2) {
	w.mousePos = position.Subtract(w.position)
	// TODO: figure out what is causing this vertical error
	w.mousePos = w.mousePos.Add(vmath.Vector2{0, 2})
	w.mousePos = w.mousePos.Multiply(vmath.Vector2{1, 1.02})
	if w.element != nil {
		w.element.mouseMove(w.mousePos)
	}
}

func (w *Window) mouseClick(button int, release bool) {
	if w.element != nil {
		w.element.mouseClick(button, release, w.mousePos)
	}
}

func (w *Window) keyClick(key string, release bool) {
	if w.element != nil {
		w.element.keyClick(key, release)
	}
}

func NewWindow() *Window {
	node := renderer.CreateNode()
	elementNode := renderer.CreateNode()
	background := renderer.CreateNode()
	box := renderer.CreateBoxWithOffset(1, 1, 0, 0)
	box.SetColor(color.NRGBA{255, 255, 255, 255})
	mat := renderer.CreateMaterial()
	mat.LightingMode = renderer.MODE_UNLIT
	box.Material = mat
	background.Add(box)
	node.Add(background)
	node.Add(elementNode)
	return &Window{
		node:          node,
		backgroundBox: box,
		background:    background,
		elementNode:   elementNode,
		size:          vmath.Vector2{500, 1},
	}
}