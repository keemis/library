// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validation

import (
	"testing"
)

type People struct {
	ID   int    `default:"3"`
	Sex  int    `default:"1"`
	Name string `default:"Tom"`
	Age  int64  `default:"18"`
	Num  uint32 `default:"20200324"`
	Has  bool   `default:"true"`
}

func TestSetDefault(t *testing.T) {
	u := People{ID: 28, Sex: 0, Name: "", Num: 0, Has: false}
	t.Logf("%+v", u)
	err := SetDefault(&u)
	t.Logf("err: %v", err)
	t.Logf("%+v", u)
}
