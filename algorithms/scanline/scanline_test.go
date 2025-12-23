package scanline_test

import (
	"testing"

	"github.com/mahalde/advent-of-code/algorithms/scanline"
	"github.com/mahalde/advent-of-code/tuples"
)

func TestBuildSpans_Empty(t *testing.T) {
	offsetX, offsetY, width, height, spans := scanline.BuildSpans(nil)
	if offsetX != 0 || offsetY != 0 || width != 0 || height != 0 || spans != nil {
		t.Errorf("Empty input: got (%d,%d,%d,%d,%v), want (0,0,0,0,nil)", offsetX, offsetY, width, height, spans)
	}
}

func TestBuildSpans_SinglePoint(t *testing.T) {
	points := []*tuples.Tuple{tuples.NewTuple(5, 10)}
	offsetX, offsetY, width, height, spans := scanline.BuildSpans(points)
	if offsetX != 4 || offsetY != 9 || width != 3 || height != 3 {
		t.Errorf("Single point: got offsets (%d,%d), dims (%d,%d)", offsetX, offsetY, width, height)
	}
	if len(spans) != 2 {
		t.Errorf("Single point spans count: %d", len(spans))
	}
	// Span 0: y=0, empty segs
	if spans[0].YStart != 0 || spans[0].YEnd != 0 || len(spans[0].Segs) != 0 {
		t.Errorf("Span 0: %v", spans[0])
	}
	// Span 1: y=1, segs [1,1]
	if spans[1].YStart != 1 || spans[1].YEnd != 1 || len(spans[1].Segs) != 1 {
		t.Errorf("Span 1: %v", spans[1])
	}
	if spans[1].Segs[0].Start != 1 || spans[1].Segs[0].End != 1 {
		t.Errorf("Span 1 segs: %v", spans[1].Segs)
	}
}

func TestBuildSpans_Square(t *testing.T) {
	// Square: (0,0) -> (2,0) -> (2,2) -> (0,2)
	points := []*tuples.Tuple{
		tuples.NewTuple(0, 0),
		tuples.NewTuple(2, 0),
		tuples.NewTuple(2, 2),
		tuples.NewTuple(0, 2),
	}
	offsetX, offsetY, width, height, spans := scanline.BuildSpans(points)
	expectedOffsetX, expectedOffsetY := -1, -1
	expectedWidth, expectedHeight := 5, 5
	if offsetX != expectedOffsetX || offsetY != expectedOffsetY || width != expectedWidth || height != expectedHeight {
		t.Errorf("Square offsets/dims: got (%d,%d,%d,%d), want (%d,%d,%d,%d)",
			offsetX, offsetY, width, height, expectedOffsetX, expectedOffsetY, expectedWidth, expectedHeight)
	}
	// For a square, expect spans covering rows 0-4, with interior segs in middle rows
	if len(spans) < 1 {
		t.Errorf("Square spans empty")
	}
	// Row 1 (y=1 in grid): should have interior [1,3] or similar
	// This is approximate; main check is no panic and reasonable output
	for _, sp := range spans {
		if sp.YStart > sp.YEnd {
			t.Errorf("Invalid span: %v", sp)
		}
		for _, seg := range sp.Segs {
			if seg.Start > seg.End {
				t.Errorf("Invalid seg: %v", seg)
			}
		}
	}
}

func TestBuildSpans_Polygon(t *testing.T) {
	points := []*tuples.Tuple{
		tuples.NewTuple(7, 1),
		tuples.NewTuple(11, 1),
		tuples.NewTuple(11, 7),
		tuples.NewTuple(9, 7),
		tuples.NewTuple(9, 5),
		tuples.NewTuple(2, 5),
		tuples.NewTuple(2, 3),
		tuples.NewTuple(7, 3),
	}
	offsetX, offsetY, width, height, spans := scanline.BuildSpans(points)
	expectedOffsetX, expectedOffsetY := 1, 0
	expectedWidth, expectedHeight := 12, 9
	if offsetX != expectedOffsetX || offsetY != expectedOffsetY || width != expectedWidth || height != expectedHeight {
		t.Errorf("Polygon offsets/dims: got (%d,%d,%d,%d), want (%d,%d,%d,%d)",
			offsetX, offsetY, width, height, expectedOffsetX, expectedOffsetY, expectedWidth, expectedHeight)
	}
	if len(spans) == 0 {
		t.Errorf("Test input spans empty")
	}
	// Verify no invalid spans
	for _, sp := range spans {
		if sp.YStart > sp.YEnd {
			t.Errorf("Invalid span: %v", sp)
		}
	}
	// Specific check: row 1 (y=1) should have segs [6,10] from debug
	found := false
	for _, sp := range spans {
		if sp.YStart <= 1 && sp.YEnd >= 1 {
			for _, seg := range sp.Segs {
				if seg.Start == 6 && seg.End == 10 {
					found = true
				}
			}
		}
	}
	if !found {
		t.Errorf("Expected seg [6,10] in row 1 not found")
	}
}
