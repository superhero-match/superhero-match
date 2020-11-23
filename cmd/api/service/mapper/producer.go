/*
  Copyright (C) 2019 - 2021 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package mapper

import (
	"github.com/superhero-match/superhero-match/cmd/api/model"
	producer "github.com/superhero-match/superhero-match/internal/producer/model"
)

// MapAPIMatchToProducer maps API Match model to Producer Match model.
func MapAPIMatchToProducer(apiMatch model.Match) producer.Match {
	return producer.Match{
		ID:                 apiMatch.ID,
		SuperheroID:        apiMatch.SuperheroID,
		MatchedSuperheroID: apiMatch.MatchedSuperheroID,
		CreatedAt:          apiMatch.CreatedAt,
	}
}
