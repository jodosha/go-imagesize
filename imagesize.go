package imagesize

import(
  "bytes"
  "errors"
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
    case "jpeg":
      return getJpegSize(image)
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

func getJpegSize(image []byte) (size ImageSize, err error) {
  var i int

  /* Check for a valid JPEG file */
  if (image[i] == 0xFF && image[i+1] == 0xD8 && image[i+2] == 0xFF && image[i+3] == 0xE0) {
      i += 4

      /* Check for valid JPEG header (null terminated JFIF) */
      if (image[i+2] == 'J' && image[i+3] == 'F' && image[i+4] == 'I' && image[i+5] == 'F' && image[i+6] == 0x00) {

        var blockLength int = int(image[i]) * 256 + int(image[i+1])

        for( i < 1024 ) {
          i += blockLength

          if ( i >= 1024 )        { err = errors.New("JPEG out of range") }
          if ( image[i] != 0xFF ) { err = errors.New("Not at the beginning of a JPEG block") }

          if( image[i+1] == 0xC0 ) {
            size.Width  = int(image[i+7]) * 256 + int(image[i+8])
            size.Height = int(image[i+5]) * 256 + int(image[i+6])

            return
          } else {
            i += 2 /* Skip the block marker */
            blockLength = int(image[i]) * 256 + int(image[i+1])
          }
        }


      } else {
        err = errors.New("Invalid JFIF.")
      }

  } else {
    err = errors.New("Invalid JPEG.")
  }

  return
}
