/*
 *
 * Copyright 2022 puzzleredisclient authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
package puzzleredisclient

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/dvaumoron/puzzlelogger"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

const redisKey = "redis"

func Create(logger *zap.Logger) *redis.Client {
	dbNum, err := strconv.Atoi(os.Getenv("REDIS_SERVER_DB"))
	if err != nil {
		logger.Fatal("Failed to parse REDIS_SERVER_DB")
	}

	redis.SetLogger(loggerWrapper{inner: logger})

	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_SERVER_ADDR"),
		Username: os.Getenv("REDIS_SERVER_USERNAME"),
		Password: os.Getenv("REDIS_SERVER_PASSWORD"),
		DB:       dbNum,
	})
}

type loggerWrapper struct {
	inner *zap.Logger
}

func (w loggerWrapper) Printf(ctx context.Context, msg string, args ...any) {
	fmt.Fprintf(puzzlelogger.InfoWrapper{Inner: w.inner, Lib: redisKey}, msg, args...)
}
