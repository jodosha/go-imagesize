package imagesize

import(
  "testing"
  "io/ioutil"
)

func TestCalculatePngImageSize(t *testing.T) {
  img, _ := ioutil.ReadFile("fixtures/gopher.png")

  const imageType = "png"
  const width     = 176
  const height    = 240

  size, _ := GetSize(img, imageType)

  if size.Width != width {
    t.Errorf("Expected %v, got %v", width, size.Width)
  }

  if size.Height != height {
    t.Errorf("Expected %v, got %v", height, size.Height)
  }
}

func TestCalculateGifImageSize(t *testing.T) {
  img, _ := ioutil.ReadFile("fixtures/gopher-latte.gif")

  const imageType = "gif"
  const width     = 180
  const height    = 240

  size, _ := GetSize(img, imageType)

  if size.Width != width {
    t.Errorf("Expected %v, got %v", width, size.Width)
  }

  if size.Height != height {
    t.Errorf("Expected %v, got %v", height, size.Height)
  }
}

func TestCalculateJpegImageSize(t *testing.T) {
  img, _ := ioutil.ReadFile("fixtures/gopher.jpeg")

  const imageType = "jpeg"
  const width     = 612
  const height    = 612

  size, _ := GetSize(img, imageType)

  if size.Width != width {
    t.Errorf("Expected %v, got %v", width, size.Width)
  }

  if size.Height != height {
    t.Errorf("Expected %v, got %v", height, size.Height)
  }
}
