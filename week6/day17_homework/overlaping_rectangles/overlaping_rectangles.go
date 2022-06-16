package overlaping_rectangles

func FindOverlapingRectangles(A, B, C, D, E, F, G, H int) int {
	/// case 1 : (A between E and G) && (B between F & H)
	/// case 2 : (C between E and G) && (B between F & H)
	/// case 3 : (C between E and G) && (D between F & H)
	/// case 4 : (A between E and G) && (D between F & H)

	/// case 5 : (E between A and C) && (F between B & D)
	/// case 6 : (G between A and C) && (F between B & D)
	/// case 7 : (G between A and C) && (H between B & D)
	/// case 8 : (E between A and C) && (H between B & D)

	if A == C || B == D || E == G || F == H {
		return 0
	}
	if A >= G || E >= C {
		return 0
	}
	if B >= H || F >= D {
		return 0
	}
	return 1

}
