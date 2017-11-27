package dao

import (
	"os"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"time"
)

type SlideshowDAO interface {
	Run()
}

type Slide struct {
	File      string
	SlideTime time.Duration
}

type fileSlideShow struct {
	outputFile string
	slides     []Slide
}

func NewSlideShow(slides []Slide, outputFile string) SlideshowDAO {
	dao := fileSlideShow{
		slides:     slides,
		outputFile: outputFile,
	}
	return dao
}

func (dao fileSlideShow) Run() {
	for true {
		for _, slide := range dao.slides {
			file, err := os.Open(slide.File)
			if err != nil {
				log.Infof("Unable to read file %s, received error %v", slide.File, err)
				continue
			}

			byteRepr, err := ioutil.ReadAll(file)
			if err != nil {
				log.Info("Error reading file %s, received error %v", slide.File, err)
			}
			file.Close()

			outFile, err := os.OpenFile(dao.outputFile, os.O_TRUNC|os.O_WRONLY, 0777)
			if err != nil {
				log.Infof("Error opening output file %s: %v", dao.outputFile, err)
			}
			_, err = outFile.Write(byteRepr)
			if err != nil {
				log.Infof("Error writing to output file %s: %v", dao.outputFile, err)
			}
			outFile.Close()

			time.Sleep(slide.SlideTime)
		}
	}
}