package app

import (
	"encoding/json"
	"fmt"

	"github.com/c-m-hunt/club-tweeter/pkg/app/loaders"
)


func (sc *ScoreCache) Save() error {
	s3l := &loaders.S3Loader{
		Bucket: cfg.ScoreImgs.CachePath.Bucket,
		Key: sc.keyPath(),
	}

	scContent, err := json.Marshal(sc)

	err = s3l.Save(&scContent)
	if err != nil {
		return err
	}
	return nil
}

func (sc *ScoreCache) keyPath() string {
	return fmt.Sprintf("%s/%s.json", cfg.ScoreImgs.CachePath.Key, sc.MatchId)
}

func NewScoreCache(matchId string) *ScoreCache {
	sc := &ScoreCache{
		MatchId: matchId,
	}

	sc.Load()

	return sc
}

func (sc *ScoreCache) Load() error {
	s3l := &loaders.S3Loader{
		Bucket: cfg.ScoreImgs.CachePath.Bucket,
		Key: sc.keyPath(),
	}

	content, err := s3l.Load()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(*content, &sc); err != nil {
        panic(err)
    }

	return nil
}

func (sc *ScoreCache) Compare(oldSc *ScoreCache) (*ScoreCache, error) {
	return nil, nil
}

