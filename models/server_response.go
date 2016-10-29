package models

import (
  "github.com/davidjfelix/videosvc/errors"
)

type ServerResponse struct{
  StatusCode int
  Message string
  Errors errors.ServerResponseErrors
  Location string
}
