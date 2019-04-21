/* Copyright 2019 Freerware
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
 */

package work

type bestEffortUniter struct {
	parameters UnitParameters
}

// NewBestEffortUniter constructs a new best effort unit factory.
func NewBestEffortUniter(parameters UnitParameters) Uniter {
	return &bestEffortUniter{parameters: parameters}
}

// Unit constructs a new best effort work unit.
func (u *bestEffortUniter) Unit() (Unit, error) {
	return NewBestEffortUnit(u.parameters), nil
}