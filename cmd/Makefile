#
# Golang port of Open Chinese Convert
#
# Copyright 2018 Zhanliang Liu <liang@zliu.org>
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
#
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=gocc

PREFIX?=/usr/local
INSTALL_BIN=$(PREFIX)/bin
INSTALL_DATA=$(PREFIX)/share/gocc

build:
	$(GOGET)
	$(GOBUILD) -o $(BINARY_NAME)

clean:
	$(GOCLEAN)

install: build
	@mkdir -p $(INSTALL_BIN)
	@mkdir -p $(INSTALL_DATA)
	cp $(BINARY_NAME) $(INSTALL_BIN)
	cp -fr ../config ../dictionary $(INSTALL_DATA)
