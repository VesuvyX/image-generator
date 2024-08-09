package img

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
	"strconv"

	"github.com/VesuvyX/2-image-generator-server/pkg/colors"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

const (
	imgColorDefault = "E5E5E5"
	msgColorDefault = "AAAAAA"
	imgWDefault     = 300
	imgHDefault     = 300

	fontSizeDefault         = 0
	fontFileDefault         = "Lifehack-Medium.ttf"
	dpiDefault      float64 = 72
	hintingDefault          = "none"
)

func GenerateFavicon() (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)

	m := image.NewRGBA(image.Rect(0, 0, 16, 16))
	clr := color.RGBA{B: 0, A: 0} // Черный, прозрачность 0
	draw.Draw(m, m.Bounds(), &image.Uniform{C: clr}, image.Point{}, draw.Src)

	var img image.Image = m
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		return nil, err
	}
	return buffer, nil
}

func Generate(ulrPart []string) (*bytes.Buffer, error) {

	var (
		err      error
		imgColor = imgColorDefault
		msgColor = msgColorDefault
		imgW     = imgWDefault
		imgH     = imgHDefault
		fontSize = fontSizeDefault
	)

	msg := ""

	for i, val := range ulrPart {
		switch i {
		case 1:
			if val != "" {
				imgW, err = strconv.Atoi(val)
				if err != nil {
					return nil, err
				}
			}
		case 2:
			if val != "" {
				imgH, err = strconv.Atoi(val)
				if err != nil {
					return nil, err
				}
			}
		case 3:
			if val != "" {
				imgColor = val
			}
		case 4:
			if val != "" {
				msg = val
			}
		case 5:
			if val != "" {
				msgColor = val
			}
		case 6:
			fontSize, err = strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
		}
	}

	// Если ширина и высота есть, а текст пустой или просто текст пустой,
	// то пишем ширину и высоту на картинке
	if ((imgW > 0 || imgH > 0) && msg == "") || msg == "" {
		msg = fmt.Sprintf("%d x %d", imgW, imgH)
	}

	if fontSize == 0 {
		fontSize = imgW / 9
		if imgH < imgW {
			fontSize = imgH / 5
		}
	}

	hx := colors.Hex(imgColor)
	rgb, err := hx.ToRGB()
	if err != nil {
		return nil, err
	}

	m := image.NewRGBA(image.Rect(0, 0, imgW, imgH))
	imgRgba := color.RGBA{R: rgb.Red, G: rgb.Green, B: rgb.Blue, A: 10}
	draw.Draw(m, m.Bounds(), &image.Uniform{C: imgRgba}, image.Point{}, draw.Src)

	addLabel(m, imgW, imgH, msg, fontSize, colors.Hex(msgColor))

	var img image.Image = m
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	return buffer, nil
}

func addLabel(img *image.RGBA, imgW, imgH int, msg string, msgFontSize int, msgColor colors.Hex) {
	var (
		fontFile = fontFileDefault
		dpi      = dpiDefault
		hinting  = hintingDefault
	)
	h := font.HintingNone
	switch hinting {
	case "full":
		h = font.HintingFull
	}
	fontBytes, err := os.ReadFile(fontFile)
	if err != nil {
		log.Println(err)
		return
	}
	fnt, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}
	rgb, err := colors.Hex2RGB(msgColor) // Берем цвета rgb и парсим msg
	if err != nil {
		log.Println(err)
		return
	}
	clr := color.Color(color.RGBA{R: rgb.Red, G: rgb.Green, B: rgb.Blue, A: 255}) // Поставляем все для структуры стандартного пакета color
	d := &font.Drawer{
		Dst: img,
		Src: image.NewUniform(clr),
		// Параметры шрифта
		Face: truetype.NewFace(fnt, &truetype.Options{
			Size:    float64(msgFontSize),
			DPI:     dpi,
			Hinting: h,
		}),
	}

	y := imgH/2 + msgFontSize/2 - 12
	// Позиционируем текст на картинке
	d.Dot = fixed.Point26_6{
		X: (fixed.I(imgW) - d.MeasureString(msg)) / 2, // По середине
		Y: fixed.I(y),                                 // +- по середине
	}
	d.DrawString(msg) // Пишем в картинку
}
