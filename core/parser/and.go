package parser

type And1Result[A any] struct {
	A A
}

func And1[A any](a Parser[A]) Parser[And1Result[A]] {
	return func(in *Input) Result[And1Result[A]] {
		start := in.Position
		resA := a(in)
		fail := Result[And1Result[A]]{Failed: true}

		if !resA.Failed {

			value := And1Result[A]{
				A: resA.Value,
			}
			return Result[And1Result[A]]{Value: value}
		}

		in.SetPosition(start)
		return fail
	}
}

type And2Result[
	A any,
	B any,
] struct {
	A A
	B B
}

func And2[A any, B any](a Parser[A], b Parser[B]) Parser[And2Result[A, B]] {
	return func(in *Input) Result[And2Result[A, B]] {
		fail := Result[And2Result[A, B]]{Failed: true}
		start := in.Position

		resA := a(in)
		if !resA.Failed {
			resB := b(in)
			if !resB.Failed {
				value := And2Result[A, B]{
					A: resA.Value,
					B: resB.Value,
				}
				return Result[And2Result[A, B]]{Value: value}
			}
		}

		in.SetPosition(start)
		return fail
	}
}

type And3Result[
	A any,
	B any,
	C any,
] struct {
	A A
	B B
	C C
}

func And3[A any, B any, C any](
	a Parser[A], b Parser[B], c Parser[C],
) Parser[And3Result[A, B, C]] {
	return func(in *Input) Result[And3Result[A, B, C]] {
		start := in.Position
		resA := a(in)

		if !resA.Failed {
			resB := b(in)
			if !resB.Failed {
				resC := c(in)
				if !resC.Failed {
					return Result[And3Result[A, B, C]]{
						Value: And3Result[A, B, C]{
							A: resA.Value,
							B: resB.Value,
							C: resC.Value,
						},
					}
				}
			}
		}
		in.SetPosition(start)
		return Result[And3Result[A, B, C]]{Failed: true}
	}
}

type And4Result[
	A any,
	B any,
	C any,
	D any,
] struct {
	A A
	B B
	C C
	D D
}

func And4[
	A any,
	B any,
	C any,
	D any,
](
	a Parser[A],
	b Parser[B],
	c Parser[C],
	d Parser[D],
) Parser[And4Result[A, B, C, D]] {
	return func(in *Input) Result[And4Result[A, B, C, D]] {
		start := in.Position
		resA := a(in)

		if !resA.Failed {
			resB := b(in)

			if !resB.Failed {
				resC := c(in)

				if !resC.Failed {
					resD := d(in)

					if !resD.Failed {
						return Result[And4Result[A, B, C, D]]{
							Value: And4Result[A, B, C, D]{
								A: resA.Value,
								B: resB.Value,
								C: resC.Value,
								D: resD.Value,
							},
						}
					}
				}
			}
		}

		in.SetPosition(start)
		return Result[And4Result[A, B, C, D]]{Failed: true}
	}
}

type And5Result[
	A any,
	B any,
	C any,
	D any,
	E any,
] struct {
	A A
	B B
	C C
	D D
	E E
}

func And5[
	A any,
	B any,
	C any,
	D any,
	E any,
](
	a Parser[A],
	b Parser[B],
	c Parser[C],
	d Parser[D],
	e Parser[E],
) Parser[And5Result[A, B, C, D, E]] {
	return func(in *Input) Result[And5Result[A, B, C, D, E]] {
		start := in.Position
		resA := a(in)

		if !resA.Failed {
			resB := b(in)

			if !resB.Failed {
				resC := c(in)

				if !resC.Failed {
					resD := d(in)

					if !resD.Failed {
						resE := e(in)

						if !resE.Failed {
							return Result[And5Result[A, B, C, D, E]]{
								Value: And5Result[A, B, C, D, E]{
									A: resA.Value,
									B: resB.Value,
									C: resC.Value,
									D: resD.Value,
									E: resE.Value,
								},
							}
						}
					}
				}
			}
		}

		in.SetPosition(start)
		return Result[And5Result[A, B, C, D, E]]{Failed: true}
	}
}

type And6Result[
	A any,
	B any,
	C any,
	D any,
	E any,
	F any,
] struct {
	A A
	B B
	C C
	D D
	E E
	F F
}

func And6[
	A any,
	B any,
	C any,
	D any,
	E any,
	F any,
](
	a Parser[A],
	b Parser[B],
	c Parser[C],
	d Parser[D],
	e Parser[E],
	f Parser[F],
) Parser[And6Result[A, B, C, D, E, F]] {
	return func(in *Input) Result[And6Result[A, B, C, D, E, F]] {
		start := in.Position
		resA := a(in)

		if !resA.Failed {
			resB := b(in)

			if !resB.Failed {
				resC := c(in)

				if !resC.Failed {
					resD := d(in)

					if !resD.Failed {
						resE := e(in)

						if !resE.Failed {
							resF := f(in)

							if !resF.Failed {
								return Result[And6Result[A, B, C, D, E, F]]{
									Value: And6Result[A, B, C, D, E, F]{
										A: resA.Value,
										B: resB.Value,
										C: resC.Value,
										D: resD.Value,
										E: resE.Value,
										F: resF.Value,
									},
								}
							}
						}
					}
				}
			}
		}

		in.SetPosition(start)
		return Result[And6Result[A, B, C, D, E, F]]{Failed: true}
	}
}

type And7Result[
	A any,
	B any,
	C any,
	D any,
	E any,
	F any,
	G any,
] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
}

func And7[
	A any,
	B any,
	C any,
	D any,
	E any,
	F any,
	G any,
](
	a Parser[A],
	b Parser[B],
	c Parser[C],
	d Parser[D],
	e Parser[E],
	f Parser[F],
	g Parser[G],
) Parser[And7Result[A, B, C, D, E, F, G]] {
	return func(in *Input) Result[And7Result[A, B, C, D, E, F, G]] {
		start := in.Position
		resA := a(in)

		if !resA.Failed {
			resB := b(in)

			if !resB.Failed {
				resC := c(in)

				if !resC.Failed {
					resD := d(in)

					if !resD.Failed {
						resE := e(in)

						if !resE.Failed {
							resF := f(in)

							if !resF.Failed {
								resG := g(in)

								if !resG.Failed {
									return Result[And7Result[A, B, C, D, E, F, G]]{
										Value: And7Result[A, B, C, D, E, F, G]{
											A: resA.Value,
											B: resB.Value,
											C: resC.Value,
											D: resD.Value,
											E: resE.Value,
											F: resF.Value,
											G: resG.Value,
										},
									}
								}
							}
						}
					}
				}
			}
		}

		in.SetPosition(start)
		return Result[And7Result[A, B, C, D, E, F, G]]{Failed: true}
	}
}

type And8Result[
	A any,
	B any,
	C any,
	D any,
	E any,
	F any,
	G any,
	H any,
] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
}

func And8[
	A any,
	B any,
	C any,
	D any,
	E any,
	F any,
	G any,
	H any,
](
	a Parser[A],
	b Parser[B],
	c Parser[C],
	d Parser[D],
	e Parser[E],
	f Parser[F],
	g Parser[G],
	h Parser[H],
) Parser[And8Result[A, B, C, D, E, F, G, H]] {
	return func(in *Input) Result[And8Result[A, B, C, D, E, F, G, H]] {
		start := in.Position
		resA := a(in)

		if !resA.Failed {
			resB := b(in)

			if !resB.Failed {
				resC := c(in)

				if !resC.Failed {
					resD := d(in)

					if !resD.Failed {
						resE := e(in)

						if !resE.Failed {
							resF := f(in)

							if !resF.Failed {
								resG := g(in)

								if !resG.Failed {
									resH := h(in)

									if !resH.Failed {
										return Result[And8Result[A, B, C, D, E, F, G, H]]{
											Value: And8Result[A, B, C, D, E, F, G, H]{
												A: resA.Value,
												B: resB.Value,
												C: resC.Value,
												D: resD.Value,
												E: resE.Value,
												F: resF.Value,
												G: resG.Value,
												H: resH.Value,
											},
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
		return Result[And8Result[A, B, C, D, E, F, G, H]]{Failed: true}
	}
}

type And9Result[
	A any,
	B any,
	C any,
	D any,
	E any,
	F any,
	G any,
	H any,
	I any,
] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
	I I
}

func And9[
	A any,
	B any,
	C any,
	D any,
	E any,
	F any,
	G any,
	H any,
	I any,
](
	a Parser[A],
	b Parser[B],
	c Parser[C],
	d Parser[D],
	e Parser[E],
	f Parser[F],
	g Parser[G],
	h Parser[H],
	i Parser[I],
) Parser[And9Result[A, B, C, D, E, F, G, H, I]] {
	return func(in *Input) Result[And9Result[A, B, C, D, E, F, G, H, I]] {
		start := in.Position
		resA := a(in)

		if !resA.Failed {
			resB := b(in)

			if !resB.Failed {
				resC := c(in)

				if !resC.Failed {
					resD := d(in)

					if !resD.Failed {
						resE := e(in)

						if !resE.Failed {
							resF := f(in)

							if !resF.Failed {
								resG := g(in)

								if !resG.Failed {
									resH := h(in)

									if !resH.Failed {
										resI := i(in)

										if !resI.Failed {
											return Result[And9Result[A, B, C, D, E, F, G, H, I]]{
												Value: And9Result[A, B, C, D, E, F, G, H, I]{
													A: resA.Value,
													B: resB.Value,
													C: resC.Value,
													D: resD.Value,
													E: resE.Value,
													F: resF.Value,
													G: resG.Value,
													H: resH.Value,
													I: resI.Value,
												},
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
		return Result[And9Result[A, B, C, D, E, F, G, H, I]]{Failed: true}
	}
}

type And10Result[
	A any,
	B any,
	C any,
	D any,
	E any,
	F any,
	G any,
	H any,
	I any,
	J any,
] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
	I I
	J J
}

func And10[
	A any,
	B any,
	C any,
	D any,
	E any,
	F any,
	G any,
	H any,
	I any,
	J any,
](
	a Parser[A],
	b Parser[B],
	c Parser[C],
	d Parser[D],
	e Parser[E],
	f Parser[F],
	g Parser[G],
	h Parser[H],
	i Parser[I],
	j Parser[J],
) Parser[And10Result[A, B, C, D, E, F, G, H, I, J]] {
	return func(in *Input) Result[And10Result[A, B, C, D, E, F, G, H, I, J]] {
		start := in.Position
		resA := a(in)

		if !resA.Failed {
			resB := b(in)

			if !resB.Failed {
				resC := c(in)

				if !resC.Failed {
					resD := d(in)

					if !resD.Failed {
						resE := e(in)

						if !resE.Failed {
							resF := f(in)

							if !resF.Failed {
								resG := g(in)

								if !resG.Failed {
									resH := h(in)

									if !resH.Failed {
										resI := i(in)

										if !resI.Failed {
											resJ := j(in)

											if !resJ.Failed {
												return Result[And10Result[A, B, C, D, E, F, G, H, I, J]]{
													Value: And10Result[A, B, C, D, E, F, G, H, I, J]{
														A: resA.Value,
														B: resB.Value,
														C: resC.Value,
														D: resD.Value,
														E: resE.Value,
														F: resF.Value,
														G: resG.Value,
														H: resH.Value,
														I: resI.Value,
														J: resJ.Value,
													},
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
		return Result[And10Result[A, B, C, D, E, F, G, H, I, J]]{Failed: true}
	}
}

type And11Result[
	A any,
	B any,
	C any,
	D any,
	E any,
	F any,
	G any,
	H any,
	I any,
	J any,
	K any,
] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
	I I
	J J
	K K
}

func And11[
	A any,
	B any,
	C any,
	D any,
	E any,
	F any,
	G any,
	H any,
	I any,
	J any,
	K any,
](
	a Parser[A],
	b Parser[B],
	c Parser[C],
	d Parser[D],
	e Parser[E],
	f Parser[F],
	g Parser[G],
	h Parser[H],
	i Parser[I],
	j Parser[J],
	k Parser[K],
) Parser[And11Result[A, B, C, D, E, F, G, H, I, J, K]] {
	return func(in *Input) Result[And11Result[A, B, C, D, E, F, G, H, I, J, K]] {
		start := in.Position
		resA := a(in)

		if !resA.Failed {
			resB := b(in)

			if !resB.Failed {
				resC := c(in)

				if !resC.Failed {
					resD := d(in)

					if !resD.Failed {
						resE := e(in)

						if !resE.Failed {
							resF := f(in)

							if !resF.Failed {
								resG := g(in)

								if !resG.Failed {
									resH := h(in)

									if !resH.Failed {
										resI := i(in)

										if !resI.Failed {
											resJ := j(in)

											if !resJ.Failed {
												resK := k(in)

												if !resK.Failed {
													return Result[And11Result[A, B, C, D, E, F, G, H, I, J, K]]{
														Value: And11Result[A, B, C, D, E, F, G, H, I, J, K]{
															A: resA.Value,
															B: resB.Value,
															C: resC.Value,
															D: resD.Value,
															E: resE.Value,
															F: resF.Value,
															G: resG.Value,
															H: resH.Value,
															I: resI.Value,
															J: resJ.Value,
															K: resK.Value,
														},
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
		return Result[And11Result[A, B, C, D, E, F, G, H, I, J, K]]{Failed: true}
	}
}

type And12Result[
	A any,
	B any,
	C any,
	D any,
	E any,
	F any,
	G any,
	H any,
	I any,
	J any,
	K any,
	L any,
] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
	I I
	J J
	K K
	L L
}

func And12[
	A any,
	B any,
	C any,
	D any,
	E any,
	F any,
	G any,
	H any,
	I any,
	J any,
	K any,
	L any,
](
	a Parser[A],
	b Parser[B],
	c Parser[C],
	d Parser[D],
	e Parser[E],
	f Parser[F],
	g Parser[G],
	h Parser[H],
	i Parser[I],
	j Parser[J],
	k Parser[K],
	l Parser[L],
) Parser[And12Result[A, B, C, D, E, F, G, H, I, J, K, L]] {
	return func(in *Input) Result[And12Result[A, B, C, D, E, F, G, H, I, J, K, L]] {
		start := in.Position
		resA := a(in)

		if !resA.Failed {
			resB := b(in)

			if !resB.Failed {
				resC := c(in)

				if !resC.Failed {
					resD := d(in)

					if !resD.Failed {
						resE := e(in)

						if !resE.Failed {
							resF := f(in)

							if !resF.Failed {
								resG := g(in)

								if !resG.Failed {
									resH := h(in)

									if !resH.Failed {
										resI := i(in)

										if !resI.Failed {
											resJ := j(in)

											if !resJ.Failed {
												resK := k(in)

												if !resK.Failed {
													resL := l(in)

													if !resL.Failed {
														return Result[And12Result[A, B, C, D, E, F, G, H, I, J, K, L]]{
															Value: And12Result[A, B, C, D, E, F, G, H, I, J, K, L]{
																A: resA.Value,
																B: resB.Value,
																C: resC.Value,
																D: resD.Value,
																E: resE.Value,
																F: resF.Value,
																G: resG.Value,
																H: resH.Value,
																I: resI.Value,
																J: resJ.Value,
																K: resK.Value,
																L: resL.Value,
															},
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
		return Result[And12Result[A, B, C, D, E, F, G, H, I, J, K, L]]{Failed: true}
	}
}

type And13Result[
	A any,
	B any,
	C any,
	D any,
	E any,
	F any,
	G any,
	H any,
	I any,
	J any,
	K any,
	L any,
	M any,
] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
	I I
	J J
	K K
	L L
	M M
}

func And13[
	A any,
	B any,
	C any,
	D any,
	E any,
	F any,
	G any,
	H any,
	I any,
	J any,
	K any,
	L any,
	M any,
](
	a Parser[A],
	b Parser[B],
	c Parser[C],
	d Parser[D],
	e Parser[E],
	f Parser[F],
	g Parser[G],
	h Parser[H],
	i Parser[I],
	j Parser[J],
	k Parser[K],
	l Parser[L],
	m Parser[M],
) Parser[And13Result[A, B, C, D, E, F, G, H, I, J, K, L, M]] {
	return func(in *Input) Result[And13Result[A, B, C, D, E, F, G, H, I, J, K, L, M]] {
		start := in.Position
		resA := a(in)

		if !resA.Failed {
			resB := b(in)

			if !resB.Failed {
				resC := c(in)

				if !resC.Failed {
					resD := d(in)

					if !resD.Failed {
						resE := e(in)

						if !resE.Failed {
							resF := f(in)

							if !resF.Failed {
								resG := g(in)

								if !resG.Failed {
									resH := h(in)

									if !resH.Failed {
										resI := i(in)

										if !resI.Failed {
											resJ := j(in)

											if !resJ.Failed {
												resK := k(in)

												if !resK.Failed {
													resL := l(in)

													if !resL.Failed {
														resM := m(in)

														if !resM.Failed {
															return Result[And13Result[A, B, C, D, E, F, G, H, I, J, K, L, M]]{
																Value: And13Result[A, B, C, D, E, F, G, H, I, J, K, L, M]{
																	A: resA.Value,
																	B: resB.Value,
																	C: resC.Value,
																	D: resD.Value,
																	E: resE.Value,
																	F: resF.Value,
																	G: resG.Value,
																	H: resH.Value,
																	I: resI.Value,
																	J: resJ.Value,
																	K: resK.Value,
																	L: resL.Value,
																	M: resM.Value,
																},
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
		return Result[And13Result[A, B, C, D, E, F, G, H, I, J, K, L, M]]{Failed: true}
	}
}

type And14Result[
	A any,
	B any,
	C any,
	D any,
	E any,
	F any,
	G any,
	H any,
	I any,
	J any,
	K any,
	L any,
	M any,
	N any,
] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
	I I
	J J
	K K
	L L
	M M
	N N
}

func And14[
	A any,
	B any,
	C any,
	D any,
	E any,
	F any,
	G any,
	H any,
	I any,
	J any,
	K any,
	L any,
	M any,
	N any,
](
	a Parser[A],
	b Parser[B],
	c Parser[C],
	d Parser[D],
	e Parser[E],
	f Parser[F],
	g Parser[G],
	h Parser[H],
	i Parser[I],
	j Parser[J],
	k Parser[K],
	l Parser[L],
	m Parser[M],
	n Parser[N],
) Parser[And14Result[A, B, C, D, E, F, G, H, I, J, K, L, M, N]] {
	return func(in *Input) Result[And14Result[A, B, C, D, E, F, G, H, I, J, K, L, M, N]] {
		start := in.Position
		resA := a(in)

		if !resA.Failed {
			resB := b(in)

			if !resB.Failed {
				resC := c(in)

				if !resC.Failed {
					resD := d(in)

					if !resD.Failed {
						resE := e(in)

						if !resE.Failed {
							resF := f(in)

							if !resF.Failed {
								resG := g(in)

								if !resG.Failed {
									resH := h(in)

									if !resH.Failed {
										resI := i(in)

										if !resI.Failed {
											resJ := j(in)

											if !resJ.Failed {
												resK := k(in)

												if !resK.Failed {
													resL := l(in)

													if !resL.Failed {
														resM := m(in)

														if !resM.Failed {
															resN := n(in)

															if !resN.Failed {
																return Result[And14Result[A, B, C, D, E, F, G, H, I, J, K, L, M, N]]{
																	Value: And14Result[A, B, C, D, E, F, G, H, I, J, K, L, M, N]{
																		A: resA.Value,
																		B: resB.Value,
																		C: resC.Value,
																		D: resD.Value,
																		E: resE.Value,
																		F: resF.Value,
																		G: resG.Value,
																		H: resH.Value,
																		I: resI.Value,
																		J: resJ.Value,
																		K: resK.Value,
																		L: resL.Value,
																		M: resM.Value,
																		N: resN.Value,
																	},
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
		return Result[And14Result[A, B, C, D, E, F, G, H, I, J, K, L, M, N]]{Failed: true}
	}
}

type And15Result[
	A any,
	B any,
	C any,
	D any,
	E any,
	F any,
	G any,
	H any,
	I any,
	J any,
	K any,
	L any,
	M any,
	N any,
	O any,
] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
	I I
	J J
	K K
	L L
	M M
	N N
	O O
}

func And15[
	A any,
	B any,
	C any,
	D any,
	E any,
	F any,
	G any,
	H any,
	I any,
	J any,
	K any,
	L any,
	M any,
	N any,
	O any,
](
	a Parser[A],
	b Parser[B],
	c Parser[C],
	d Parser[D],
	e Parser[E],
	f Parser[F],
	g Parser[G],
	h Parser[H],
	i Parser[I],
	j Parser[J],
	k Parser[K],
	l Parser[L],
	m Parser[M],
	n Parser[N],
	o Parser[O],
) Parser[And15Result[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O]] {
	return func(in *Input) Result[And15Result[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O]] {
		start := in.Position
		resA := a(in)

		if !resA.Failed {
			resB := b(in)

			if !resB.Failed {
				resC := c(in)

				if !resC.Failed {
					resD := d(in)

					if !resD.Failed {
						resE := e(in)

						if !resE.Failed {
							resF := f(in)

							if !resF.Failed {
								resG := g(in)

								if !resG.Failed {
									resH := h(in)

									if !resH.Failed {
										resI := i(in)

										if !resI.Failed {
											resJ := j(in)

											if !resJ.Failed {
												resK := k(in)

												if !resK.Failed {
													resL := l(in)

													if !resL.Failed {
														resM := m(in)

														if !resM.Failed {
															resN := n(in)

															if !resN.Failed {
																resO := o(in)

																if !resO.Failed {
																	return Result[And15Result[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O]]{
																		Value: And15Result[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O]{
																			A: resA.Value,
																			B: resB.Value,
																			C: resC.Value,
																			D: resD.Value,
																			E: resE.Value,
																			F: resF.Value,
																			G: resG.Value,
																			H: resH.Value,
																			I: resI.Value,
																			J: resJ.Value,
																			K: resK.Value,
																			L: resL.Value,
																			M: resM.Value,
																			N: resN.Value,
																			O: resO.Value,
																		},
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
		}

		in.SetPosition(start)
		return Result[And15Result[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O]]{Failed: true}
	}
}
