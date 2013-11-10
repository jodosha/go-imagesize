package imagesize

import(
  "bytes"
  "encoding/binary"
)

type ImageSize struct {
  Width  int
  Height int
}

func GetSize(image []byte, imageType string) (ImageSize, error) {
  return getPngSize(image)
}

func getPngSize(image []byte) (size ImageSize, err error) {
  var width  int32
  var height int32

  err = binary.Read(bytes.NewBuffer(image[16:20]), binary.BigEndian, &width)
  err = binary.Read(bytes.NewBuffer(image[20:24]), binary.BigEndian, &height)

  size.Width  = int(width)
  size.Height = int(height)

  return
}
