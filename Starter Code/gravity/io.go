package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseOrderedPair(line string) (OrderedPair, error) {
	// Replace the Unicode minus sign with a standard hyphen-minus
	line = strings.ReplaceAll(line, "−", "-")

	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		return OrderedPair{}, fmt.Errorf("invalid ordered pair")
	}
	x, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return OrderedPair{}, err
	}
	y, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return OrderedPair{}, err
	}
	return OrderedPair{x: x, y: y}, nil
}

func ParseRGB(line string) (uint8, uint8, uint8, error) {
	parts := strings.Split(line, ",")
	if len(parts) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid RGB format")
	}
	red, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, 0, err
	}
	green, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, 0, 0, err
	}
	blue, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil {
		return 0, 0, 0, err
	}
	return uint8(red), uint8(green), uint8(blue), nil
}

func ReadUniverse(filename string) (Universe, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Universe{}, err
	}
	defer file.Close()

	var universe Universe
	scanner := bufio.NewScanner(file)

	// Read the first line to get the width of the universe
	if scanner.Scan() {
		width, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {
			return Universe{}, fmt.Errorf("invalid universe width: %v", err)
		}
		universe.width = width
	} else {
		return Universe{}, fmt.Errorf("file is empty or missing width")
	}

	// Read the second line to get the gravitational constant
	if scanner.Scan() {
		g, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {
			return Universe{}, fmt.Errorf("invalid gravitational constant: %v", err)
		}
		universe.gravitationalConstant = g
	} else {
		return Universe{}, fmt.Errorf("file is empty or missing width")
	}

	var currentBody Body
	lineType := 0 // Keeps track of which data is expected next

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		switch lineType {
		case 0: // Expecting a body name, starting with '>'
			if strings.HasPrefix(line, ">") {
				// If there was a previous body, add it to the universe
				if currentBody.name != "" {
					universe.bodies = append(universe.bodies, currentBody)
				}
				// Start a new body
				currentBody = Body{}
				currentBody.name = strings.TrimSpace(line[1:]) // Remove the '>'
				lineType = 1
			} else {
				return Universe{}, fmt.Errorf("expected body name, got: %s", line)
			}

		case 1: // Expecting RGB values
			red, green, blue, err := ParseRGB(line)
			if err != nil {
				return Universe{}, fmt.Errorf("invalid RGB values: %v", err)
			}
			currentBody.red, currentBody.green, currentBody.blue = red, green, blue
			lineType = 2

		case 2: // Expecting mass
			mass, err := strconv.ParseFloat(line, 64)
			if err != nil {
				return Universe{}, fmt.Errorf("invalid mass: %v", err)
			}
			currentBody.mass = mass
			lineType = 3

		case 3: // Expecting radius
			radius, err := strconv.ParseFloat(line, 64)
			if err != nil {
				return Universe{}, fmt.Errorf("invalid radius: %v", err)
			}
			currentBody.radius = radius
			lineType = 4

		case 4: // Expecting position (OrderedPair)
			position, err := ParseOrderedPair(line)
			if err != nil {
				return Universe{}, fmt.Errorf("invalid position: %v", err)
			}
			currentBody.position = position
			lineType = 5

		case 5: // Expecting velocity (OrderedPair)
			velocity, err := ParseOrderedPair(line)
			if err != nil {
				return Universe{}, fmt.Errorf("invalid velocity: %v", err)
			}
			currentBody.velocity = velocity
			lineType = 0 // Ready for the next body
		}
	}

	// Add the last body, if there is one
	if currentBody.name != "" {
		universe.bodies = append(universe.bodies, currentBody)
	}

	if err := scanner.Err(); err != nil {
		return Universe{}, err
	}

	return universe, nil
}
