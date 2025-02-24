package main

import (
	"fmt"
	"math"
	"net/http"

	"github.com/fogleman/gg"
	// Подключите библиотеку для рисования:
	// go mod download github.com/fogleman/gg@latest
)

// Параметры:
// - involvementPlan: Вовлеченность элемента Plan (от 0 до 100).
// - involvementImprove: Вовлеченность элемента Improve (от 0 до 100).
// - involvementEngage: Вовлеченность элемента Engage (от 0 до 100).
// - involvementDesignTransition: Вовлеченность элемента Design and Transition (от 0 до 100).
// - involvementObtainBuild: Вовлеченность элемента Obtain/Build (от 0 до 100).
// - involvementDeliverSupport: Вовлеченность элемента Deliver and Support (от 0 до 100).

const (
	width           = 1000
	height          = 600
	nodeWidth       = 150
	nodeHeight      = 70
	ovalWidth       = 55
	ovalHeight      = 50
	cubeWidth       = 150
	cubeHeight      = 100
	cubeDepth       = 30
	largeCubeWidth  = 188
	largeCubeHeight = 125
	largeCubeDepth  = 38
)

func main() {
	http.HandleFunc("/draw", drawDiagram)
	fmt.Println("Сервис запущен на http://localhost:8080/draw")
	http.ListenAndServe(":8080", nil)
}

func drawDiagram2(w http.ResponseWriter, r *http.Request) {

	// Извлечение параметров из URL involvementPlan := r.URL.Query().Get("involvementPlan")
	//involvementImprove := r.URL.Query().Get("involvementImprove")
	//involvementEngage := r.URL.Query().Get("involvementEngage")
	//involvementDesignTransition := r.URL.Query().Get("involvementDesignTransition")
	//involvementObtainBuild := r.URL.Query().Get("involvementObtainBuild")
	//involvementDeliverSupport := r.URL.Query().Get("involvementDeliverSupport")

	// Создание контекста для рисования
	const width, height = 800, 600
	dc := gg.NewContext(width, height)

	// Цветовая схема
	bgColor := "#f0f4f8"
	textColor := "#333333"
	elementColor := "#1976d2"
	//highlightColor := "#ffa000"

	// Заливка фона
	dc.SetHexColor(bgColor)
	dc.Clear()

	// Рисование элементов (пример для Plan)
	dc.SetHexColor(elementColor)
	dc.DrawRectangle(100, 100, 150, 100)
	dc.Fill()
	dc.SetHexColor(textColor)
	dc.DrawStringAnchored("Plan", 175, 150, 0.5, 0.5)

	// Добавьте остальные элементы с использованием параметров вовлеченности
	// Отправка изображения в ответ
	w.Header().Set("Content-Type", "image/png")
	dc.EncodePNG(w)
}

func drawDiagram(w http.ResponseWriter, r *http.Request) {

	// Создание контекста для рисования
	// const width, height = 800, 600
	dc := gg.NewContext(width, height)

	// Цветовая схема
	bgColor := "#f0f4f8"
	//textColor := "#333333"
	//elementColor := "#1976d2"
	//highlightColor := "#ffa000"

	// Заливка фона
	dc.SetHexColor(bgColor)
	dc.Clear()

	// Центры для фигур
	positions := map[string][2]float64{
		"Demand":           {100, height / 2},
		"Value":            {width - 100, height / 2},
		"Plan":             {width / 2, 100},
		"Improve":          {width / 2, height - 100},
		"Engage":           {300, height / 2},
		"DesignTransition": {width / 2, 300},
		"ObtainBuild":      {width / 2, 300},
		"DeliverSupport":   {width / 2, 300},
		"ProductNServices": {width - 300, height / 2},
	}

	// Рисуем стрелки
	//drawArrow(dc, positions["Demand"], positions["Engage"])
	//drawArrow(dc, positions["Engage"], positions["DesignTransition"])
	//drawArrow(dc, positions["DesignTransition"], positions["ObtainBuild"])
	//drawArrow(dc, positions["ObtainBuild"], positions["DeliverSupport"])
	//drawArrow(dc, positions["DeliverSupport"], positions["Value"])
	//drawArrow(dc, positions["Improve"], positions["Plan"])

	// Рисуем фигуры
	drawOval(dc, positions["Demand"], "Demand", 0.7, 0.7, 0.9)

	drawRectangle(dc, positions["Plan"], "Plan", 0.5, 0.7, 0.9)       // лента
	drawRectangle(dc, positions["Improve"], "Improve", 0.7, 0.9, 0.6) // лента
	//drawRectangle(dc, positions["Engage"], "Engage", 0.9, 0.7, 0.5)

	// Рисуем куб "Engage"
	drawCube(dc, positions["Engage"], "Engage", 0.9, 0.7, 0.5)

	drawTopSideCube(dc, positions["DesignTransition"], "Design\nand\nTransition", 0.8, 0.6, 0.8)
	drawLeftSideCube(dc, positions["ObtainBuild"], "Obtain/\nBuild", 0.9, 0.8, 0.4)
	drawRightSideCube(dc, positions["DeliverSupport"], "Deliver\nand\nSupport", 0.9, 0.5, 0.6)
	//drawRectangle(dc, positions["DesignTransition"], "Design\nand\nTransition", 0.8, 0.6, 0.8)
	//drawRectangle(dc, positions["ObtainBuild"], "Obtain/\nBuild", 0.9, 0.8, 0.4)
	//drawRectangle(dc, positions["DeliverSupport"], "Deliver\nand\nSupport", 0.9, 0.5, 0.6)

	// Рисуем куб "Product & Services"
	drawCube(dc, positions["ProductNServices"], "Product\nand\nServices", 0.9, 0.7, 0.5)

	drawOval(dc, positions["Value"], "Value", 0.6, 0.9, 0.6)

	// Добавьте остальные элементы с использованием параметров вовлеченности
	// Отправка изображения в ответ
	w.Header().Set("Content-Type", "image/png")
	dc.EncodePNG(w)
}

// Рисует овал
func drawOval(dc *gg.Context, pos [2]float64, text string, r, g, b float64) {
	x, y := pos[0], pos[1]

	dc.SetRGB(r, g, b)
	dc.DrawEllipse(x, y, ovalWidth, ovalHeight)
	dc.Fill()
	dc.SetRGB(0, 0, 0)
	dc.DrawEllipse(x, y, ovalWidth, ovalHeight)
	dc.Stroke()
	dc.DrawStringAnchored(text, x, y, 0.5, 0.5)
}

// Рисует прямоугольник
func drawRectangle(dc *gg.Context, pos [2]float64, text string, r, g, b float64) {
	x, y := pos[0], pos[1]

	dc.SetRGB(r, g, b)
	dc.DrawRectangle(x-nodeWidth/2, y-nodeHeight/2, nodeWidth, nodeHeight)
	dc.Fill()
	dc.SetRGB(0, 0, 0)
	dc.DrawRectangle(x-nodeWidth/2, y-nodeHeight/2, nodeWidth, nodeHeight)
	dc.Stroke()
	dc.DrawStringWrapped(text, x, y, 0.5, 0.5, nodeWidth-10, 1.5, gg.AlignCenter)
}

// Рисует стрелку между двумя точками
func drawArrow(dc *gg.Context, start, end [2]float64) {
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(2)

	x1, y1 := start[0], start[1]
	x2, y2 := end[0], end[1]

	dc.DrawLine(x1, y1, x2, y2)
	dc.Stroke()

	// Рисуем наконечник стрелки
	arrowHead(dc, x1, y1, x2, y2)
}

// Добавляет наконечник к стрелке
func arrowHead(dc *gg.Context, x1, y1, x2, y2 float64) {

	angle := math.Atan2(y2-y1, x2-x1)

	arrowSize := 10.0
	x3 := x2 - arrowSize*math.Cos(angle+gg.Radians(30))
	y3 := y2 - arrowSize*math.Sin(angle+gg.Radians(30))
	x4 := x2 - arrowSize*math.Cos(angle-gg.Radians(30))
	y4 := y2 - arrowSize*math.Sin(angle-gg.Radians(30))

	dc.NewSubPath()
	dc.MoveTo(x2, y2)
	dc.LineTo(x3, y3)
	dc.LineTo(x4, y4)
	dc.ClosePath()
	dc.Fill()
}

func drawLeftSideCube(dc *gg.Context, pos [2]float64, label string, r, g, b float64) {
	centerX, centerY := pos[0], pos[1]

	var halfWidth float64 = largeCubeWidth / 2
	var halfHeight float64 = largeCubeHeight / 2

	x1 := centerX - halfWidth
	x2 := centerX

	y2 := centerY - halfHeight
	y3 := y2 + largeCubeDepth
	y4 := centerY + halfHeight
	y5 := y4 + largeCubeDepth

	// Левая передняя грань
	leftSide := []struct{ x, y float64 }{
		{x1, y2}, // Левый верхний
		{x2, y3}, // Правый верхний
		{x2, y5}, // Правый нижний
		{x1, y4}, // Левый нижний
	}

	// Рисуем левую боковую грань
	dc.SetRGB(r-0.2, g-0.2, b-0.2) // Еще темнее
	dc.MoveTo(leftSide[0].x, leftSide[0].y)
	for _, p := range leftSide[1:] {
		dc.LineTo(p.x, p.y)
	}
	dc.ClosePath()
	dc.Fill()

	// Обводка грани TODO вынести
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(2)

	for _, face := range [][]struct{ x, y float64 }{leftSide} {
		dc.MoveTo(face[0].x, face[0].y)
		for _, p := range face[1:] {
			dc.LineTo(p.x, p.y)
		}
		dc.ClosePath()
		dc.Stroke()
	}

	// Добавляем текст
	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored(label, centerX, centerY, 0.5, 0.5)
}

// Процедура для рисования куба
func drawTopSideCube(dc *gg.Context, pos [2]float64, label string, r, g, b float64) {
	centerX, centerY := pos[0], pos[1]

	var halfWidth float64 = largeCubeWidth / 2
	var halfHeight float64 = largeCubeHeight / 2

	x1 := centerX - halfWidth
	x2 := centerX
	x3 := centerX + halfWidth

	y2 := centerY - halfHeight
	y3 := y2 + largeCubeDepth

	// Верхняя грань
	top := []struct{ x, y float64 }{
		{x1, y2},                  // Левый
		{x2, y2 - largeCubeDepth}, // Верхний
		{x3, y2},                  // Правый
		{x2, y3},                  // Нижний
	}

	// Рисуем верхнюю грань
	dc.SetRGB(r-0.1, g-0.1, b-0.1) // Темнее основного цвета
	dc.MoveTo(top[0].x, top[0].y)
	for _, p := range top[1:] {
		dc.LineTo(p.x, p.y)
	}
	dc.ClosePath()
	dc.Fill()

	// Обводка куба
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(2)

	for _, face := range [][]struct{ x, y float64 }{top} {
		dc.MoveTo(face[0].x, face[0].y)
		for _, p := range face[1:] {
			dc.LineTo(p.x, p.y)
		}
		dc.ClosePath()
		dc.Stroke()
	}

	// Добавляем текст
	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored(label, centerX, centerY, 0.5, 0.5)
}

// Процедура для рисования куба
func drawRightSideCube(dc *gg.Context, pos [2]float64, label string, r, g, b float64) {
	centerX, centerY := pos[0], pos[1]

	var halfWidth float64 = largeCubeWidth / 2
	var halfHeight float64 = largeCubeHeight / 2

	x2 := centerX
	x3 := centerX + halfWidth

	y2 := centerY - halfHeight
	y3 := y2 + largeCubeDepth
	y4 := centerY + halfHeight
	y5 := y4 + largeCubeDepth

	// Правая передняя грань
	rightSide := []struct{ x, y float64 }{
		{x2, y3}, // Верхний левый
		{x3, y2}, // Верхний правый
		{x3, y4}, // Нижний правый
		{x2, y5}, // Нижний левый
	}

	// Рисуем правую боковую грань
	dc.SetRGB(r, g, b) // Основной цвет
	dc.MoveTo(rightSide[0].x, rightSide[0].y)
	for _, p := range rightSide[1:] {
		dc.LineTo(p.x, p.y)
	}
	dc.ClosePath()
	dc.Fill()

	// Обводка куба
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(2)

	for _, face := range [][]struct{ x, y float64 }{rightSide} {
		dc.MoveTo(face[0].x, face[0].y)
		for _, p := range face[1:] {
			dc.LineTo(p.x, p.y)
		}
		dc.ClosePath()
		dc.Stroke()
	}

	// Добавляем текст
	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored(label, centerX, centerY, 0.5, 0.5)
}

// Процедура для рисования куба
func drawCube(dc *gg.Context, pos [2]float64, label string, r, g, b float64) {
	centerX, centerY := pos[0], pos[1]

	var halfWidth float64 = cubeWidth / 2
	var halfHeight float64 = cubeHeight / 2

	x1 := centerX - halfWidth
	x2 := centerX
	x3 := centerX + halfWidth

	y2 := centerY - halfHeight
	y3 := y2 + cubeDepth
	y4 := centerY + halfHeight
	y5 := y4 + cubeDepth

	// Верхняя грань
	top := []struct{ x, y float64 }{
		{x1, y2},             // Левый
		{x2, y2 - cubeDepth}, // Верхний
		{x3, y2},             // Правый
		{x2, y3},             // Нижний
	}

	// Левая передняя грань
	leftSide := []struct{ x, y float64 }{
		{x1, y2}, // Левый верхний
		{x2, y3}, // Правый верхний
		{x2, y5}, // Правый нижний
		{x1, y4}, // Левый нижний
	}

	// Правая передняя грань
	rightSide := []struct{ x, y float64 }{
		{x2, y3}, // Верхний левый
		{x3, y2}, // Верхний правый
		{x3, y4}, // Нижний правый
		{x2, y5}, // Нижний левый
	}

	// Рисуем верхнюю грань
	dc.SetRGB(r-0.1, g-0.1, b-0.1) // Темнее основного цвета
	dc.MoveTo(top[0].x, top[0].y)
	for _, p := range top[1:] {
		dc.LineTo(p.x, p.y)
	}
	dc.ClosePath()
	dc.Fill()

	// Рисуем левую боковую грань
	dc.SetRGB(r-0.2, g-0.2, b-0.2) // Еще темнее
	dc.MoveTo(leftSide[0].x, leftSide[0].y)
	for _, p := range leftSide[1:] {
		dc.LineTo(p.x, p.y)
	}
	dc.ClosePath()
	dc.Fill()

	// Рисуем правую боковую грань
	dc.SetRGB(r, g, b) // Основной цвет
	dc.MoveTo(rightSide[0].x, rightSide[0].y)
	for _, p := range rightSide[1:] {
		dc.LineTo(p.x, p.y)
	}
	dc.ClosePath()
	dc.Fill()

	// Обводка куба
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(2)

	for _, face := range [][]struct{ x, y float64 }{top, leftSide, rightSide} {
		dc.MoveTo(face[0].x, face[0].y)
		for _, p := range face[1:] {
			dc.LineTo(p.x, p.y)
		}
		dc.ClosePath()
		dc.Stroke()
	}

	// Добавляем текст
	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored(label, centerX, centerY, 0.5, 0.5)
}
