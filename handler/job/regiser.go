package main

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ebi-yade/periodichat/env"
)

var (
	registerPattern = regexp.MustCompile(`^\s*register\s+`)
	schedulePattern = regexp.MustCompile(`^\d+min$`)

	maxLen = mustAtoi(env.Or("MAX_LEN", "1000"))

	errFailedMatchingRegister = errors.New("failed to match the patten of registering a message")
	errMaxLen                 = fmt.Errorf("message must not be over %d", maxLen)
)

func register(ctx context.Context, cmd string) error {
	if !registerPattern.MatchString(cmd) {
		return fmt.Errorf(`%w: command="%s"`, errFailedMatchingRegister, cmd)
	}

	sequel := registerPattern.ReplaceAllString(cmd, "") // e.g. "15min hello world!"
	args := strings.Fields(sequel)
	if len(args) < 1 {
		return fmt.Errorf(`%w: command="%s"`, errFailedMatchingRegister, cmd)
	}
	div, err := strconv.Atoi(args[0][:len(args[0])-3])
	if err != nil {
		return fmt.Errorf(`%w: %s: command="%s"`, errFailedMatchingRegister, err.Error(), cmd)
	}
	if div == 0 || !schedulePattern.MatchString(args[0]) {
		return fmt.Errorf(`%w: command="%s"`, errFailedMatchingRegister, cmd)
	}
	msg := strings.TrimSpace(strings.TrimSpace(sequel)[len(args[0]):])
	if len(msg) > maxLen {
		return fmt.Errorf("%w, but was %d", errMaxLen, len(msg))
	}

	if err := registerToDynamoDB(ctx, div, msg); err != nil {
		return fmt.Errorf(`failed to register a message to DynamoDB: %w`, err)
	}

	return nil
}

func mustAtoi(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return num
}
