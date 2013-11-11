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
  switch imageType {
    case "png":
      return getPngSize(image)
    case "gif":
      return getGifSize(image)
  }

  return ImageSize{}, nil
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

func getGifSize(image []byte) (size ImageSize, err error) {
  var width  int16
  var height int16

  err = binary.Read(bytes.NewBuffer(image[6:8]),  binary.LittleEndian, &width)
  err = binary.Read(bytes.NewBuffer(image[8:10]), binary.LittleEndian, &height)

  size.Width  = int(width)
  size.Height = int(height)

  return
}
