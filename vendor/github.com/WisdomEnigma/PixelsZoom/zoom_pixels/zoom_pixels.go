package zoom_pixels

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"reflect"
)

// Set Image dimensions Dimension {Width = 200, Height = 200}
var width, height int = 200, 200

// const (
// 	Red int = iota << 1
// 	Green
// 	Blue
// 	AlphaValue
// )

// type RawImage struct {
// 	r, g, b, a uint32
// }

// decodeRawIamge hold image file pixels data in a form of vector
var decodeRawImage image.Image

// deltaPixels allow to hold changes in pixels after applied sub_ops
var deltaPixels Pixel_Diff

// deltaPixelsOps allow to hold change in pixels value after applied divsion
var deltaPixelsOps Pixel_Diff

// deltaPixelsAddOps allow to hold change in pixels value after applied addition
var deltaPixelsAddOps Pixel_Diff

// sorted picture pixels values
var zoom []Pixel_Diff

// Inverse specified this variables serve in inverse process (reverse)
var deltaInverse Pixel_Diff
var deltaInverseOps Pixel_Diff
var deltaInverseSubOps Pixel_Diff
var invZoom []Pixel_Diff

// Zoom_KTimes @Parameters( Level of Zoom, File Object)
// Zoom provide entry point for zoom calculation

func Zoom_KTime(value int, file *os.File) {

	// initialization of attributes
	zoom = make([]Pixel_Diff, GetImage().Bounds().Max.X*3)

	for i := 0; i < GetImage().Bounds().Max.X; i++ {
		for j := 0; j < GetImage().Bounds().Max.Y; j++ {

			// list hold Initial Pixel Value
			list := Pixels_Info{
				Value: GetImage().At(i, j),
			}

			// another hold next Pixel Value
			another_list := Pixels_Info{
				Value: GetImage().At(i, j+1),
			}

			// convert Pixel Value to RGBA color
			colorUnit_r, colorUnit_g, colorUnit_b, colorUnit_a := list.Value.RGBA()
			colorUnit_r0, colorUnit_g0, colorUnit_b0, colorUnit_a0 := another_list.Value.RGBA()

			// if RGBA color is white then ignore
			if reflect.DeepEqual(colorUnit_r, uint32(0)) && reflect.DeepEqual(colorUnit_g, uint32(0)) && reflect.DeepEqual(colorUnit_b, uint32(0)) && reflect.DeepEqual(colorUnit_a, uint32(0)) {
				continue
			}

			if reflect.DeepEqual(colorUnit_a0, uint32(0)) && reflect.DeepEqual(colorUnit_r0, uint32(0)) && reflect.DeepEqual(colorUnit_g0, uint32(0)) && reflect.DeepEqual(colorUnit_b0, uint32(0)) {
				continue
			}

			// Applied Substract Operation on RGBA color values
			deltaPixels = Substract(colorUnit_r, colorUnit_g, colorUnit_b, colorUnit_a, colorUnit_r0, colorUnit_g0, colorUnit_b0, colorUnit_a0)

			// set k value
			k = value

			// Applied Division Operation on RGBA color values
			deltaPixelsOps = Division(deltaPixels)

			// Applied Add Operation on RGBA color values
			deltaPixelsAddOps = Add(deltaPixels)

			// log.Println("Pixels Difference", deltaPixels, "Divison:", deltaPixelsOps, "Add:", deltaPixelsAddOps)

			// collect piceces in vector
			p, q, r := NewImage(deltaPixels, deltaPixelsOps, deltaPixelsAddOps)

			// sort pixel values of collected piceces
			u, v := Is_Sort(p, q, r)

			// store in memory stack for futher operation
			zoom = append(zoom, u, v)
		}

	}

	// copy pixel function scale up image value in vector
	newPicture := copy_pixel()

	// scale up values transform into image
	ZoomPicture(file, newPicture)

	// k always zero which means restart in next echo
	k = 0
}

// Zoom Level
var k int = 0

// var img *image.RGBA

// Pixels Info hold picture pixels coordinates {x, y}
type Pixels_Info struct {

	// Pixel value based on pixel coordinates
	Value color.Color
}

// Pixel Difference hold rgba color values for an image
type Pixel_Diff struct {

	// rgba  color values
	r, g, b, a uint32
}

// Substraction is not usually allowed on pixels values;because pixels values exist in vector form.
// In order to preserve, we use rgba color value
func Substract(r, g, b, a uint32, r0, g0, b0, a0 uint32) Pixel_Diff {

	sub_r := r - r0
	sub_g := g - g0
	sub_b := b - b0
	sub_a := a - a0

	// check whether value is white or not
	if r > r0 && g > g0 && b > b0 && a > a0 {
		return Pixel_Diff{r: sub_r, g: sub_g, b: sub_b, a: sub_a}
	}

	sub_r = r0 - r
	sub_g = g0 - g
	sub_b = b0 - b
	sub_a = a0 - a
	return Pixel_Diff{r: sub_r, g: sub_g, b: sub_b, a: sub_a}
}

// Divsion is similar to Substraction
func Division(p Pixel_Diff) Pixel_Diff {

	// check whether zoom value is zero then return empty pixel value
	if k == 0 {
		log.Fatalln("Divison operation is not allowed")
		return Pixel_Diff{}
	}

	dr := int(p.r) / k
	dg := int(p.g) / k
	db := int(p.b) / k
	da := int(p.a) / k

	return Pixel_Diff{r: uint32(dr), g: uint32(dg), b: uint32(db), a: uint32(da)}

}

// Add is similar to substract and divsion operation
func Add(p Pixel_Diff) Pixel_Diff {

	da := int(p.r) + k
	dg := int(p.g) + k
	db := int(p.b) + k
	dr := int(p.r) + k

	return Pixel_Diff{r: uint32(dr), g: uint32(dg), b: uint32(db), a: uint32(da)}
}

// var avatar_zoom *image.Paletted

// copy pixels will copy generated pixels data into new image
func copy_pixels(_zoom []Pixel_Diff) *image.Paletted {

	var pictureColor []color.Color

	for i := range _zoom {
		pictureColor = []color.Color{
			color.RGBA64{uint16(_zoom[i].r), uint16(_zoom[i].g), uint16(_zoom[i].b), uint16(_zoom[i].a)},
		}
	}

	return image.NewPaletted(image.Rect(0, 0, width*k, height*k), pictureColor)
}

func copy_pixel() *image.Paletted {

	var pictureColor []color.Color

	for i := range zoom {
		pictureColor = []color.Color{
			color.RGBA64{uint16(zoom[i].r), uint16(zoom[i].g), uint16(zoom[i].b), uint16(zoom[i].a)},
		}
	}

	return image.NewPaletted(image.Rect(0, 0, width*k, height*k), pictureColor)
}

// Set image hold picture pixels values
func SetImage(im image.Image) { decodeRawImage = im }

// return picture pixels value
func GetImage() image.Image { return decodeRawImage }

func NewImage(p, q, r Pixel_Diff) (Pixel_Diff, Pixel_Diff, Pixel_Diff) {

	return p, q, r
}

// shuffling is not allowed to pixels value; Instead of pixels values we use rgba color values
// shuffling shuffle rgba color values
func shuffle(s Pixel_Diff, t Pixel_Diff) (Pixel_Diff, Pixel_Diff) {

	temp := Pixel_Diff{}

	temp.r, temp.g, temp.b, temp.a = s.r, s.g, s.b, s.a
	s.r, s.g, s.b, s.a = t.r, t.g, t.b, t.a
	t.r, t.g, t.b, t.a = temp.r, temp.g, temp.b, temp.a

	return t, s
}

func Is_Sort(p, q, r Pixel_Diff) (Pixel_Diff, Pixel_Diff) {

	var u, v Pixel_Diff

	// log.Println("Q:", q, " P:", p)

	// if rgba color is different then shuffle otherwise return empty pixel diff (pixel valeu)
	if q.r > p.r && q.g > p.g && q.b > p.b && q.a > p.a {

		u, v = shuffle(p, q)
		// log.Println("Pixels sorts between p & q:", u, v)
		return u, v
	}

	// log.Println("R:", r, " P:", p)
	if r.r > p.r && r.g > p.g && r.b > p.b && r.a > p.a {

		u, v = shuffle(p, r)
		// log.Println("Pixels sorts between r & p:", u, v)
		return u, v
	}

	// log.Println("Q:", q, " R:", r)
	if r.r > q.r && r.g > q.g && r.b > q.b && r.a > q.a {

		u, v = shuffle(q, r)
		// log.Println("Pixels sorts between q & r:", u, v)
		return u, v
	}

	return Pixel_Diff{}, Pixel_Diff{}
}

// Zoom Picture allow to create a new picture based on the given pixel values
func ZoomPicture(file *os.File, newPicture *image.Paletted) {

	// compression operation should be performed before translation
	encoder := png.Encoder{CompressionLevel: png.BestCompression}
	err := encoder.Encode(file, newPicture)

	if err != nil {
		log.Fatalln("picture encode error:", err)
		return
	}
}

// Zoom Out Pixel is inverse process of image zoom in
func ZoomOutPixels(file *os.File, level int) {

	for i := 0; i < GetImage().Bounds().Max.X; i++ {
		for j := 0; j < GetImage().Bounds().Max.Y; j++ {

			first := Pixels_Info{
				Value: GetImage().At(i, j),
			}

			second := Pixels_Info{
				Value: GetImage().At(i, j),
			}

			// convert Pixel Value to RGBA color
			colorUnit_r, colorUnit_g, colorUnit_b, colorUnit_a := first.Value.RGBA()
			colorUnit_r0, colorUnit_g0, colorUnit_b0, colorUnit_a0 := second.Value.RGBA()

			// if RGBA color is white then ignore
			if reflect.DeepEqual(colorUnit_r, uint32(0)) && reflect.DeepEqual(colorUnit_g, uint32(0)) && reflect.DeepEqual(colorUnit_b, uint32(0)) && reflect.DeepEqual(colorUnit_a, uint32(0)) {
				continue
			}

			if reflect.DeepEqual(colorUnit_a0, uint32(0)) && reflect.DeepEqual(colorUnit_r0, uint32(0)) && reflect.DeepEqual(colorUnit_g0, uint32(0)) && reflect.DeepEqual(colorUnit_b0, uint32(0)) {
				continue
			}

			// Reverse zoom in process
			deltaInverse = InverseAddition(colorUnit_r, colorUnit_g, colorUnit_b, colorUnit_a, colorUnit_r0, colorUnit_g0, colorUnit_b0, colorUnit_a0)
			deltaInverseOps = InverseMultiplacate(deltaInverse)
			deltaInverseSubOps = InverseSubstract(deltaInverse)

			k = level

			p, q, r := NewImage(deltaInverse, deltaInverseOps, deltaInverseSubOps)

			// sort the pixels in reverse
			u, v := Is_Sort(p, q, r)

			// hold reverse pixels in reverse array called invZoom
			invZoom = append(invZoom, u, v)

		}
	}

	// create new image by clone the previous image data
	newPicture := copy_pixels(invZoom)

	// Picture is ready now
	ZoomPicture(file, newPicture)

	k = 0

}

func InverseAddition(r, g, b, a uint32, r0, g0, b0, a0 uint32) Pixel_Diff {

	add_r := r + r0
	add_g := g + g0
	add_b := b + b0
	add_a := a + a0

	if r0 > r && g0 > g && b0 > b && a0 > a {
		return Pixel_Diff{r: uint32(add_r), g: uint32(add_g), b: uint32(add_b), a: uint32(add_a)}
	}

	add_r = r0 + r
	add_g = g0 + g
	add_b = b0 + b
	add_a = a0 + a

	return Pixel_Diff{r: uint32(add_r), g: uint32(add_g), b: uint32(add_b), a: uint32(add_a)}
}

func InverseMultiplacate(p Pixel_Diff) Pixel_Diff {

	op_r := int(p.r) * k
	op_g := int(p.g) * k
	op_b := int(p.b) * k
	op_a := int(p.a) * k

	return Pixel_Diff{r: uint32(op_r), g: uint32(op_g), b: uint32(op_b), a: uint32(op_a)}
}

func InverseSubstract(p Pixel_Diff) Pixel_Diff {

	sub_r := int(p.r) - k
	sub_g := int(p.g) - k
	sub_b := int(p.b) - k
	sub_a := int(p.a) - k

	return Pixel_Diff{r: uint32(sub_r), g: uint32(sub_g), b: uint32(sub_b), a: uint32(sub_a)}
}
