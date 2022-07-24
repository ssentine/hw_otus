package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrOffsetIsNegative      = errors.New("offset is negative")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromFile, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer fromFile.Close()

	err = CheckOffeset(fromFile, offset)
	if err != nil {
		return err
	}

	toFile, err := os.Create(toPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer toFile.Close()

	err = CopyFile(toFile, fromFile, limit, offset)
	if err != nil {
		return err
	}
	return nil
}

func CheckOffeset(fd *os.File, offset int64) error {
	if offset < 0 {
		return ErrOffsetIsNegative
	}

	if FileSize(fd) < offset {
		return ErrOffsetExceedsFileSize
	}
	return nil
}

func CopyFile(toFile, fromFile *os.File, limit, offset int64) error {
	fs := FileSize(fromFile)
	if limit == 0 || limit > fs-offset {
		limit = fs - offset
	}

	reader := io.LimitReader(fromFile, limit)
	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(reader)
	defer bar.Finish()

	_, err := fromFile.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	_, err = io.Copy(toFile, barReader)
	if err != nil {
		return err
	}
	return nil
}

func FileSize(fd *os.File) int64 {
	info, _ := fd.Stat()
	return info.Size()
}
