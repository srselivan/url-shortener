package pudge

import "github.com/recoilme/pudge"

func New(filePath string) (*pudge.Db, error) {
	return pudge.Open(filePath, pudge.DefaultConfig)
}
