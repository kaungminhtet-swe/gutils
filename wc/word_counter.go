package wc

import (
	`bufio`
	`errors`
	`io`
)

func CountLines(rd io.Reader) (int64, error) {
	if rd == nil {
		return 0, errors.New("nil reader")
	}

	br := bufio.NewReader(rd)
	var count int64
	for {
		_, _, err := br.ReadLine()
		if errors.Is(err, io.EOF) {
			return count, nil
		}

		if err != nil {
			return 0, err
		}

		count++
	}
}
