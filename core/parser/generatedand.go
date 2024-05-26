package parser

type And1Result[
	T1 any,
] struct {
	R1 T1
}

func And1[T1 any](p1 Parser[T1]) Parser[And1Result[T1]] {
	return func(in *Input) (And1Result[T1], bool) {
		start := in.Position

		res1, ok1 := p1(in)
		if ok1 {
			return And1Result[T1]{
				R1: res1,
			}, true
		}
		in.SetPosition(start)
		return And1Result[T1]{}, false
	}
}

type And2Result[
	T1 any,
	T2 any,
] struct {
	R1 T1
	R2 T2
}

func And2[T1 any, T2 any](p1 Parser[T1], p2 Parser[T2]) Parser[And2Result[T1, T2]] {
	return func(in *Input) (And2Result[T1, T2], bool) {
		start := in.Position

		res1, ok1 := p1(in)
		if ok1 {
			res2, ok2 := p2(in)
			if ok2 {
				return And2Result[T1, T2]{
					R1: res1,
					R2: res2,
				}, true
			}
		}
		in.SetPosition(start)
		return And2Result[T1, T2]{}, false
	}
}

type And3Result[
	T1 any,
	T2 any,
	T3 any,
] struct {
	R1 T1
	R2 T2
	R3 T3
}

func And3[T1 any, T2 any, T3 any](p1 Parser[T1], p2 Parser[T2], p3 Parser[T3]) Parser[And3Result[T1, T2, T3]] {
	return func(in *Input) (And3Result[T1, T2, T3], bool) {
		start := in.Position

		res1, ok1 := p1(in)
		if ok1 {
			res2, ok2 := p2(in)
			if ok2 {
				res3, ok3 := p3(in)
				if ok3 {
					return And3Result[T1, T2, T3]{
						R1: res1,
						R2: res2,
						R3: res3,
					}, true
				}
			}
		}
		in.SetPosition(start)
		return And3Result[T1, T2, T3]{}, false
	}
}

type And4Result[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
] struct {
	R1 T1
	R2 T2
	R3 T3
	R4 T4
}

func And4[T1 any, T2 any, T3 any, T4 any](p1 Parser[T1], p2 Parser[T2], p3 Parser[T3], p4 Parser[T4]) Parser[And4Result[T1, T2, T3, T4]] {
	return func(in *Input) (And4Result[T1, T2, T3, T4], bool) {
		start := in.Position

		res1, ok1 := p1(in)
		if ok1 {
			res2, ok2 := p2(in)
			if ok2 {
				res3, ok3 := p3(in)
				if ok3 {
					res4, ok4 := p4(in)
					if ok4 {
						return And4Result[T1, T2, T3, T4]{
							R1: res1,
							R2: res2,
							R3: res3,
							R4: res4,
						}, true
					}
				}
			}
		}
		in.SetPosition(start)
		return And4Result[T1, T2, T3, T4]{}, false
	}
}

type And5Result[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
] struct {
	R1 T1
	R2 T2
	R3 T3
	R4 T4
	R5 T5
}

func And5[T1 any, T2 any, T3 any, T4 any, T5 any](p1 Parser[T1], p2 Parser[T2], p3 Parser[T3], p4 Parser[T4], p5 Parser[T5]) Parser[And5Result[T1, T2, T3, T4, T5]] {
	return func(in *Input) (And5Result[T1, T2, T3, T4, T5], bool) {
		start := in.Position

		res1, ok1 := p1(in)
		if ok1 {
			res2, ok2 := p2(in)
			if ok2 {
				res3, ok3 := p3(in)
				if ok3 {
					res4, ok4 := p4(in)
					if ok4 {
						res5, ok5 := p5(in)
						if ok5 {
							return And5Result[T1, T2, T3, T4, T5]{
								R1: res1,
								R2: res2,
								R3: res3,
								R4: res4,
								R5: res5,
							}, true
						}
					}
				}
			}
		}
		in.SetPosition(start)
		return And5Result[T1, T2, T3, T4, T5]{}, false
	}
}

type And6Result[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
	T6 any,
] struct {
	R1 T1
	R2 T2
	R3 T3
	R4 T4
	R5 T5
	R6 T6
}

func And6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](p1 Parser[T1], p2 Parser[T2], p3 Parser[T3], p4 Parser[T4], p5 Parser[T5], p6 Parser[T6]) Parser[And6Result[T1, T2, T3, T4, T5, T6]] {
	return func(in *Input) (And6Result[T1, T2, T3, T4, T5, T6], bool) {
		start := in.Position

		res1, ok1 := p1(in)
		if ok1 {
			res2, ok2 := p2(in)
			if ok2 {
				res3, ok3 := p3(in)
				if ok3 {
					res4, ok4 := p4(in)
					if ok4 {
						res5, ok5 := p5(in)
						if ok5 {
							res6, ok6 := p6(in)
							if ok6 {
								return And6Result[T1, T2, T3, T4, T5, T6]{
									R1: res1,
									R2: res2,
									R3: res3,
									R4: res4,
									R5: res5,
									R6: res6,
								}, true
							}
						}
					}
				}
			}
		}
		in.SetPosition(start)
		return And6Result[T1, T2, T3, T4, T5, T6]{}, false
	}
}

type And7Result[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
	T6 any,
	T7 any,
] struct {
	R1 T1
	R2 T2
	R3 T3
	R4 T4
	R5 T5
	R6 T6
	R7 T7
}

func And7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any](p1 Parser[T1], p2 Parser[T2], p3 Parser[T3], p4 Parser[T4], p5 Parser[T5], p6 Parser[T6], p7 Parser[T7]) Parser[And7Result[T1, T2, T3, T4, T5, T6, T7]] {
	return func(in *Input) (And7Result[T1, T2, T3, T4, T5, T6, T7], bool) {
		start := in.Position

		res1, ok1 := p1(in)
		if ok1 {
			res2, ok2 := p2(in)
			if ok2 {
				res3, ok3 := p3(in)
				if ok3 {
					res4, ok4 := p4(in)
					if ok4 {
						res5, ok5 := p5(in)
						if ok5 {
							res6, ok6 := p6(in)
							if ok6 {
								res7, ok7 := p7(in)
								if ok7 {
									return And7Result[T1, T2, T3, T4, T5, T6, T7]{
										R1: res1,
										R2: res2,
										R3: res3,
										R4: res4,
										R5: res5,
										R6: res6,
										R7: res7,
									}, true
								}
							}
						}
					}
				}
			}
		}
		in.SetPosition(start)
		return And7Result[T1, T2, T3, T4, T5, T6, T7]{}, false
	}
}

type And8Result[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
	T6 any,
	T7 any,
	T8 any,
] struct {
	R1 T1
	R2 T2
	R3 T3
	R4 T4
	R5 T5
	R6 T6
	R7 T7
	R8 T8
}

func And8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any](p1 Parser[T1], p2 Parser[T2], p3 Parser[T3], p4 Parser[T4], p5 Parser[T5], p6 Parser[T6], p7 Parser[T7], p8 Parser[T8]) Parser[And8Result[T1, T2, T3, T4, T5, T6, T7, T8]] {
	return func(in *Input) (And8Result[T1, T2, T3, T4, T5, T6, T7, T8], bool) {
		start := in.Position

		res1, ok1 := p1(in)
		if ok1 {
			res2, ok2 := p2(in)
			if ok2 {
				res3, ok3 := p3(in)
				if ok3 {
					res4, ok4 := p4(in)
					if ok4 {
						res5, ok5 := p5(in)
						if ok5 {
							res6, ok6 := p6(in)
							if ok6 {
								res7, ok7 := p7(in)
								if ok7 {
									res8, ok8 := p8(in)
									if ok8 {
										return And8Result[T1, T2, T3, T4, T5, T6, T7, T8]{
											R1: res1,
											R2: res2,
											R3: res3,
											R4: res4,
											R5: res5,
											R6: res6,
											R7: res7,
											R8: res8,
										}, true
									}
								}
							}
						}
					}
				}
			}
		}
		in.SetPosition(start)
		return And8Result[T1, T2, T3, T4, T5, T6, T7, T8]{}, false
	}
}

type And9Result[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
	T6 any,
	T7 any,
	T8 any,
	T9 any,
] struct {
	R1 T1
	R2 T2
	R3 T3
	R4 T4
	R5 T5
	R6 T6
	R7 T7
	R8 T8
	R9 T9
}

func And9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any](p1 Parser[T1], p2 Parser[T2], p3 Parser[T3], p4 Parser[T4], p5 Parser[T5], p6 Parser[T6], p7 Parser[T7], p8 Parser[T8], p9 Parser[T9]) Parser[And9Result[T1, T2, T3, T4, T5, T6, T7, T8, T9]] {
	return func(in *Input) (And9Result[T1, T2, T3, T4, T5, T6, T7, T8, T9], bool) {
		start := in.Position

		res1, ok1 := p1(in)
		if ok1 {
			res2, ok2 := p2(in)
			if ok2 {
				res3, ok3 := p3(in)
				if ok3 {
					res4, ok4 := p4(in)
					if ok4 {
						res5, ok5 := p5(in)
						if ok5 {
							res6, ok6 := p6(in)
							if ok6 {
								res7, ok7 := p7(in)
								if ok7 {
									res8, ok8 := p8(in)
									if ok8 {
										res9, ok9 := p9(in)
										if ok9 {
											return And9Result[T1, T2, T3, T4, T5, T6, T7, T8, T9]{
												R1: res1,
												R2: res2,
												R3: res3,
												R4: res4,
												R5: res5,
												R6: res6,
												R7: res7,
												R8: res8,
												R9: res9,
											}, true
										}
									}
								}
							}
						}
					}
				}
			}
		}
		in.SetPosition(start)
		return And9Result[T1, T2, T3, T4, T5, T6, T7, T8, T9]{}, false
	}
}

type And10Result[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
	T6 any,
	T7 any,
	T8 any,
	T9 any,
	T10 any,
] struct {
	R1  T1
	R2  T2
	R3  T3
	R4  T4
	R5  T5
	R6  T6
	R7  T7
	R8  T8
	R9  T9
	R10 T10
}

func And10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any](p1 Parser[T1], p2 Parser[T2], p3 Parser[T3], p4 Parser[T4], p5 Parser[T5], p6 Parser[T6], p7 Parser[T7], p8 Parser[T8], p9 Parser[T9], p10 Parser[T10]) Parser[And10Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]] {
	return func(in *Input) (And10Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10], bool) {
		start := in.Position

		res1, ok1 := p1(in)
		if ok1 {
			res2, ok2 := p2(in)
			if ok2 {
				res3, ok3 := p3(in)
				if ok3 {
					res4, ok4 := p4(in)
					if ok4 {
						res5, ok5 := p5(in)
						if ok5 {
							res6, ok6 := p6(in)
							if ok6 {
								res7, ok7 := p7(in)
								if ok7 {
									res8, ok8 := p8(in)
									if ok8 {
										res9, ok9 := p9(in)
										if ok9 {
											res10, ok10 := p10(in)
											if ok10 {
												return And10Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{
													R1:  res1,
													R2:  res2,
													R3:  res3,
													R4:  res4,
													R5:  res5,
													R6:  res6,
													R7:  res7,
													R8:  res8,
													R9:  res9,
													R10: res10,
												}, true
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		in.SetPosition(start)
		return And10Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{}, false
	}
}

type And11Result[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
	T6 any,
	T7 any,
	T8 any,
	T9 any,
	T10 any,
	T11 any,
] struct {
	R1  T1
	R2  T2
	R3  T3
	R4  T4
	R5  T5
	R6  T6
	R7  T7
	R8  T8
	R9  T9
	R10 T10
	R11 T11
}

func And11[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any](p1 Parser[T1], p2 Parser[T2], p3 Parser[T3], p4 Parser[T4], p5 Parser[T5], p6 Parser[T6], p7 Parser[T7], p8 Parser[T8], p9 Parser[T9], p10 Parser[T10], p11 Parser[T11]) Parser[And11Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]] {
	return func(in *Input) (And11Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11], bool) {
		start := in.Position

		res1, ok1 := p1(in)
		if ok1 {
			res2, ok2 := p2(in)
			if ok2 {
				res3, ok3 := p3(in)
				if ok3 {
					res4, ok4 := p4(in)
					if ok4 {
						res5, ok5 := p5(in)
						if ok5 {
							res6, ok6 := p6(in)
							if ok6 {
								res7, ok7 := p7(in)
								if ok7 {
									res8, ok8 := p8(in)
									if ok8 {
										res9, ok9 := p9(in)
										if ok9 {
											res10, ok10 := p10(in)
											if ok10 {
												res11, ok11 := p11(in)
												if ok11 {
													return And11Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]{
														R1:  res1,
														R2:  res2,
														R3:  res3,
														R4:  res4,
														R5:  res5,
														R6:  res6,
														R7:  res7,
														R8:  res8,
														R9:  res9,
														R10: res10,
														R11: res11,
													}, true
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		in.SetPosition(start)
		return And11Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]{}, false
	}
}

type And12Result[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
	T6 any,
	T7 any,
	T8 any,
	T9 any,
	T10 any,
	T11 any,
	T12 any,
] struct {
	R1  T1
	R2  T2
	R3  T3
	R4  T4
	R5  T5
	R6  T6
	R7  T7
	R8  T8
	R9  T9
	R10 T10
	R11 T11
	R12 T12
}

func And12[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any](p1 Parser[T1], p2 Parser[T2], p3 Parser[T3], p4 Parser[T4], p5 Parser[T5], p6 Parser[T6], p7 Parser[T7], p8 Parser[T8], p9 Parser[T9], p10 Parser[T10], p11 Parser[T11], p12 Parser[T12]) Parser[And12Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]] {
	return func(in *Input) (And12Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12], bool) {
		start := in.Position

		res1, ok1 := p1(in)
		if ok1 {
			res2, ok2 := p2(in)
			if ok2 {
				res3, ok3 := p3(in)
				if ok3 {
					res4, ok4 := p4(in)
					if ok4 {
						res5, ok5 := p5(in)
						if ok5 {
							res6, ok6 := p6(in)
							if ok6 {
								res7, ok7 := p7(in)
								if ok7 {
									res8, ok8 := p8(in)
									if ok8 {
										res9, ok9 := p9(in)
										if ok9 {
											res10, ok10 := p10(in)
											if ok10 {
												res11, ok11 := p11(in)
												if ok11 {
													res12, ok12 := p12(in)
													if ok12 {
														return And12Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]{
															R1:  res1,
															R2:  res2,
															R3:  res3,
															R4:  res4,
															R5:  res5,
															R6:  res6,
															R7:  res7,
															R8:  res8,
															R9:  res9,
															R10: res10,
															R11: res11,
															R12: res12,
														}, true
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		in.SetPosition(start)
		return And12Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]{}, false
	}
}

type And13Result[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
	T6 any,
	T7 any,
	T8 any,
	T9 any,
	T10 any,
	T11 any,
	T12 any,
	T13 any,
] struct {
	R1  T1
	R2  T2
	R3  T3
	R4  T4
	R5  T5
	R6  T6
	R7  T7
	R8  T8
	R9  T9
	R10 T10
	R11 T11
	R12 T12
	R13 T13
}

func And13[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any](p1 Parser[T1], p2 Parser[T2], p3 Parser[T3], p4 Parser[T4], p5 Parser[T5], p6 Parser[T6], p7 Parser[T7], p8 Parser[T8], p9 Parser[T9], p10 Parser[T10], p11 Parser[T11], p12 Parser[T12], p13 Parser[T13]) Parser[And13Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]] {
	return func(in *Input) (And13Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13], bool) {
		start := in.Position

		res1, ok1 := p1(in)
		if ok1 {
			res2, ok2 := p2(in)
			if ok2 {
				res3, ok3 := p3(in)
				if ok3 {
					res4, ok4 := p4(in)
					if ok4 {
						res5, ok5 := p5(in)
						if ok5 {
							res6, ok6 := p6(in)
							if ok6 {
								res7, ok7 := p7(in)
								if ok7 {
									res8, ok8 := p8(in)
									if ok8 {
										res9, ok9 := p9(in)
										if ok9 {
											res10, ok10 := p10(in)
											if ok10 {
												res11, ok11 := p11(in)
												if ok11 {
													res12, ok12 := p12(in)
													if ok12 {
														res13, ok13 := p13(in)
														if ok13 {
															return And13Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]{
																R1:  res1,
																R2:  res2,
																R3:  res3,
																R4:  res4,
																R5:  res5,
																R6:  res6,
																R7:  res7,
																R8:  res8,
																R9:  res9,
																R10: res10,
																R11: res11,
																R12: res12,
																R13: res13,
															}, true
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		in.SetPosition(start)
		return And13Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]{}, false
	}
}

type And14Result[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
	T6 any,
	T7 any,
	T8 any,
	T9 any,
	T10 any,
	T11 any,
	T12 any,
	T13 any,
	T14 any,
] struct {
	R1  T1
	R2  T2
	R3  T3
	R4  T4
	R5  T5
	R6  T6
	R7  T7
	R8  T8
	R9  T9
	R10 T10
	R11 T11
	R12 T12
	R13 T13
	R14 T14
}

func And14[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any](p1 Parser[T1], p2 Parser[T2], p3 Parser[T3], p4 Parser[T4], p5 Parser[T5], p6 Parser[T6], p7 Parser[T7], p8 Parser[T8], p9 Parser[T9], p10 Parser[T10], p11 Parser[T11], p12 Parser[T12], p13 Parser[T13], p14 Parser[T14]) Parser[And14Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]] {
	return func(in *Input) (And14Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14], bool) {
		start := in.Position

		res1, ok1 := p1(in)
		if ok1 {
			res2, ok2 := p2(in)
			if ok2 {
				res3, ok3 := p3(in)
				if ok3 {
					res4, ok4 := p4(in)
					if ok4 {
						res5, ok5 := p5(in)
						if ok5 {
							res6, ok6 := p6(in)
							if ok6 {
								res7, ok7 := p7(in)
								if ok7 {
									res8, ok8 := p8(in)
									if ok8 {
										res9, ok9 := p9(in)
										if ok9 {
											res10, ok10 := p10(in)
											if ok10 {
												res11, ok11 := p11(in)
												if ok11 {
													res12, ok12 := p12(in)
													if ok12 {
														res13, ok13 := p13(in)
														if ok13 {
															res14, ok14 := p14(in)
															if ok14 {
																return And14Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]{
																	R1:  res1,
																	R2:  res2,
																	R3:  res3,
																	R4:  res4,
																	R5:  res5,
																	R6:  res6,
																	R7:  res7,
																	R8:  res8,
																	R9:  res9,
																	R10: res10,
																	R11: res11,
																	R12: res12,
																	R13: res13,
																	R14: res14,
																}, true
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		in.SetPosition(start)
		return And14Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]{}, false
	}
}

type And15Result[
	T1 any,
	T2 any,
	T3 any,
	T4 any,
	T5 any,
	T6 any,
	T7 any,
	T8 any,
	T9 any,
	T10 any,
	T11 any,
	T12 any,
	T13 any,
	T14 any,
	T15 any,
] struct {
	R1  T1
	R2  T2
	R3  T3
	R4  T4
	R5  T5
	R6  T6
	R7  T7
	R8  T8
	R9  T9
	R10 T10
	R11 T11
	R12 T12
	R13 T13
	R14 T14
	R15 T15
}

func And15[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, T12 any, T13 any, T14 any, T15 any](p1 Parser[T1], p2 Parser[T2], p3 Parser[T3], p4 Parser[T4], p5 Parser[T5], p6 Parser[T6], p7 Parser[T7], p8 Parser[T8], p9 Parser[T9], p10 Parser[T10], p11 Parser[T11], p12 Parser[T12], p13 Parser[T13], p14 Parser[T14], p15 Parser[T15]) Parser[And15Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]] {
	return func(in *Input) (And15Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15], bool) {
		start := in.Position

		res1, ok1 := p1(in)
		if ok1 {
			res2, ok2 := p2(in)
			if ok2 {
				res3, ok3 := p3(in)
				if ok3 {
					res4, ok4 := p4(in)
					if ok4 {
						res5, ok5 := p5(in)
						if ok5 {
							res6, ok6 := p6(in)
							if ok6 {
								res7, ok7 := p7(in)
								if ok7 {
									res8, ok8 := p8(in)
									if ok8 {
										res9, ok9 := p9(in)
										if ok9 {
											res10, ok10 := p10(in)
											if ok10 {
												res11, ok11 := p11(in)
												if ok11 {
													res12, ok12 := p12(in)
													if ok12 {
														res13, ok13 := p13(in)
														if ok13 {
															res14, ok14 := p14(in)
															if ok14 {
																res15, ok15 := p15(in)
																if ok15 {
																	return And15Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]{
																		R1:  res1,
																		R2:  res2,
																		R3:  res3,
																		R4:  res4,
																		R5:  res5,
																		R6:  res6,
																		R7:  res7,
																		R8:  res8,
																		R9:  res9,
																		R10: res10,
																		R11: res11,
																		R12: res12,
																		R13: res13,
																		R14: res14,
																		R15: res15,
																	}, true
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		in.SetPosition(start)
		return And15Result[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]{}, false
	}
}
