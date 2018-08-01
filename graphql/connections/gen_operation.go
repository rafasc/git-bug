// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package connections

import (
	"fmt"

	"github.com/MichaelMure/git-bug/bug"
	"github.com/MichaelMure/git-bug/graphql/models"
)

type BugOperationEdger func(value bug.Operation, offset int) Edge

type BugOperationConMaker func(
	edges []models.OperationEdge,
	nodes []bug.Operation,
	info models.PageInfo,
	totalCount int) (models.OperationConnection, error)

func BugOperationCon(source []bug.Operation, edger BugOperationEdger, conMaker BugOperationConMaker, input models.ConnectionInput) (models.OperationConnection, error) {
	var nodes []bug.Operation
	var edges []models.OperationEdge
	var pageInfo models.PageInfo

	emptyCon, _ := conMaker(edges, nodes, pageInfo, 0)

	offset := 0

	if input.After != nil {
		for i, value := range source {
			edge := edger(value, i)
			if edge.GetCursor() == *input.After {
				// remove all previous element including the "after" one
				source = source[i+1:]
				offset = i + 1
				break
			}
		}
	}

	if input.Before != nil {
		for i, value := range source {
			edge := edger(value, i+offset)

			if edge.GetCursor() == *input.Before {
				// remove all after element including the "before" one
				break
			}

			edges = append(edges, edge.(models.OperationEdge))
			nodes = append(nodes, value)
		}
	} else {
		edges = make([]models.OperationEdge, len(source))
		nodes = source

		for i, value := range source {
			edges[i] = edger(value, i+offset).(models.OperationEdge)
		}
	}

	if input.First != nil {
		if *input.First < 0 {
			return emptyCon, fmt.Errorf("first less than zero")
		}

		if len(edges) > *input.First {
			// Slice result to be of length first by removing edges from the end
			edges = edges[:*input.First]
			nodes = nodes[:*input.First]
			pageInfo.HasNextPage = true
		}
	}

	if input.Last != nil {
		if *input.Last < 0 {
			return emptyCon, fmt.Errorf("last less than zero")
		}

		if len(edges) > *input.Last {
			// Slice result to be of length last by removing edges from the start
			edges = edges[len(edges)-*input.Last:]
			nodes = nodes[len(nodes)-*input.Last:]
			pageInfo.HasPreviousPage = true
		}
	}

	return conMaker(edges, nodes, pageInfo, len(source))
}
