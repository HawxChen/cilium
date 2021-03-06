// Copyright 2016-2017 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package policy

import (
	. "gopkg.in/check.v1"
)

func (s *PolicyTestSuite) TestReservedID(c *C) {
	i1 := GetReservedID("host")
	c.Assert(i1, Equals, ID_HOST)
	c.Assert(i1.String(), Equals, "host")

	i2 := GetReservedID("world")
	c.Assert(i2, Equals, ID_WORLD)
	c.Assert(i2.String(), Equals, "world")

	c.Assert(GetReservedID("unknown"), Equals, ID_UNKNOWN)
	unknown := NumericIdentity(700)
	c.Assert(unknown.String(), Equals, "700")
}
