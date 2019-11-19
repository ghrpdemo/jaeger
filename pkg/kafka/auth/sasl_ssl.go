// Copyright (c) 2019 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"errors"
	"crypto/tls"
	"github.com/Shopify/sarama"
)

// SASLSSLConfig defines configurations required for SASL Plain authentication (Refer: https://kafka.apache.org/documentation/#security_sasl_plain)
type SASLSSLConfig struct {
	UserName string
	Password string
}

func setSASLSSLConfiguration(config *SASLSSLConfig, saramaConfig *sarama.Config) error {
	if len(config.UserName) == 0 || len(config.Password) == 0 {
		return errors.New("invalid username/password supplied for SASL_SSL authentication. username/password cannot be empty")
	}
	saramaConfig.Net.SASL.Enable = true
//	saramaConfig.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	saramaConfig.Net.SASL.User = config.UserName
	saramaConfig.Net.SASL.Password = config.Password
	saramaConfig.Net.SASL.handshake = true
	
	saramaConfig.Net.TLS.Enable = true
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ClientAuth: 0,
	}
	saramaConfig.Net.TLS.Config = tlsConfig
	return nil
}
