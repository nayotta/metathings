package main

import (
	"strings"

	rpio "github.com/stianeikeland/go-rpio"
)

/*

The library use the raw BCM2835 pinouts, not the ports as they are mapped
on the output pins for the raspberry pi, and not the wiringPi convention.

            Rev 2 and 3 Raspberry Pi                        Rev 1 Raspberry Pi (legacy)
  +-----+---------+----------+---------+-----+      +-----+--------+----------+--------+-----+
  | BCM |   Name  | Physical | Name    | BCM |      | BCM | Name   | Physical | Name   | BCM |
  +-----+---------+----++----+---------+-----+      +-----+--------+----++----+--------+-----+
  |     |    3.3v |  1 || 2  | 5v      |     |      |     | 3.3v   |  1 ||  2 | 5v     |     |
  |   2 |   SDA 1 |  3 || 4  | 5v      |     |      |   0 | SDA    |  3 ||  4 | 5v     |     |
  |   3 |   SCL 1 |  5 || 6  | 0v      |     |      |   1 | SCL    |  5 ||  6 | 0v     |     |
  |   4 | GPIO  7 |  7 || 8  | TxD     | 14  |      |   4 | GPIO 7 |  7 ||  8 | TxD    |  14 |
  |     |      0v |  9 || 10 | RxD     | 15  |      |     | 0v     |  9 || 10 | RxD    |  15 |
  |  17 | GPIO  0 | 11 || 12 | GPIO  1 | 18  |      |  17 | GPIO 0 | 11 || 12 | GPIO 1 |  18 |
  |  27 | GPIO  2 | 13 || 14 | 0v      |     |      |  21 | GPIO 2 | 13 || 14 | 0v     |     |
  |  22 | GPIO  3 | 15 || 16 | GPIO  4 | 23  |      |  22 | GPIO 3 | 15 || 16 | GPIO 4 |  23 |
  |     |    3.3v | 17 || 18 | GPIO  5 | 24  |      |     | 3.3v   | 17 || 18 | GPIO 5 |  24 |
  |  10 |    MOSI | 19 || 20 | 0v      |     |      |  10 | MOSI   | 19 || 20 | 0v     |     |
  |   9 |    MISO | 21 || 22 | GPIO  6 | 25  |      |   9 | MISO   | 21 || 22 | GPIO 6 |  25 |
  |  11 |    SCLK | 23 || 24 | CE0     | 8   |      |  11 | SCLK   | 23 || 24 | CE0    |   8 |
  |     |      0v | 25 || 26 | CE1     | 7   |      |     | 0v     | 25 || 26 | CE1    |   7 |
  |   0 |   SDA 0 | 27 || 28 | SCL 0   | 1   |      +-----+--------+----++----+--------+-----+
  |   5 | GPIO 21 | 29 || 30 | 0v      |     |
  |   6 | GPIO 22 | 31 || 32 | GPIO 26 | 12  |
  |  13 | GPIO 23 | 33 || 34 | 0v      |     |
  |  19 | GPIO 24 | 35 || 36 | GPIO 27 | 16  |
  |  26 | GPIO 25 | 37 || 38 | GPIO 28 | 20  |
  |     |      0v | 39 || 40 | GPIO 29 | 21  |
  +-----+---------+----++----+---------+-----+

See the spec for full details of the BCM2835 controller:

https://www.raspberrypi.org/documentation/hardware/raspberrypi/bcm2835/BCM2835-ARM-Peripherals.pdf
and https://elinux.org/BCM2835_datasheet_errata - for errors in that spec

*/

var (
	pinModernMapper = map[int]rpio.Pin{
		3:  rpio.Pin(2),
		5:  rpio.Pin(3),
		7:  rpio.Pin(4),
		11: rpio.Pin(17),
		13: rpio.Pin(27),
		15: rpio.Pin(22),
		19: rpio.Pin(10),
		21: rpio.Pin(9),
		23: rpio.Pin(11),
		27: rpio.Pin(0),
		29: rpio.Pin(5),
		31: rpio.Pin(6),
		33: rpio.Pin(13),
		35: rpio.Pin(19),
		37: rpio.Pin(26),
		8:  rpio.Pin(14),
		10: rpio.Pin(15),
		12: rpio.Pin(18),
		16: rpio.Pin(23),
		18: rpio.Pin(24),
		22: rpio.Pin(25),
		24: rpio.Pin(8),
		26: rpio.Pin(7),
		28: rpio.Pin(1),
		32: rpio.Pin(12),
		36: rpio.Pin(16),
		38: rpio.Pin(20),
		40: rpio.Pin(21),
	}
	pinLegacyMapper = map[int]rpio.Pin{
		3:  rpio.Pin(0),
		5:  rpio.Pin(1),
		7:  rpio.Pin(4),
		11: rpio.Pin(17),
		13: rpio.Pin(21),
		15: rpio.Pin(22),
		19: rpio.Pin(10),
		21: rpio.Pin(9),
		23: rpio.Pin(11),
		8:  rpio.Pin(14),
		10: rpio.Pin(15),
		12: rpio.Pin(18),
		16: rpio.Pin(23),
		18: rpio.Pin(24),
		22: rpio.Pin(25),
		24: rpio.Pin(8),
	}
)

type PiModel int8

const (
	MODERN PiModel = iota
	LEGACY
)

var models = map[string]PiModel{
	"modern": MODERN,
	"legacy": LEGACY,
	"pi":     LEGACY,
	"pi1":    LEGACY,
	"pi2":    MODERN,
	"pi3":    MODERN,
	"pi0":    MODERN,
	"pi0w":   MODERN,
}

func Pin(model string, pin int) (rpio.Pin, error) {
	var mapper map[int]rpio.Pin

	m, ok := models[strings.ToLower(model)]
	if !ok {
		return rpio.Pin(0), ErrUnknownModel
	}

	if m == LEGACY {
		mapper = pinLegacyMapper
	} else if m == MODERN {
		mapper = pinModernMapper
	}

	p, ok := mapper[pin]
	if !ok {
		return rpio.Pin(0), ErrUnknownPin
	}

	return p, nil
}
